package main

import (
	_ "embed"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
	"sync"
)

//go:embed keys.ini
var apiKey string
var (
	listDevicesRe  = regexp.MustCompile(`^\/devices[\/]$`)
	updateDeviceRe = regexp.MustCompile(`^\/devices\/(.+)$`)
)

type Location struct {
	Lat          json.Number `json:"lat"`
	Lng          json.Number `json:"lng"`
	Device_state DeviceState `json:"device_state"`
}
type DeviceResponse struct {
	Result_list []Device
}
type Device struct {
	Device_id                    string         `json:"device_id"`
	Display_name                 string         `json:"display_name"`
	Latest_accurate_device_point Location       `json:"latest_accurate_device_point"`
	Device_ui_settings           DeviceSettings `json:"device_ui_settings"`
}
type DeviceSettings struct {
	Is_visible  bool   `json:"is_visible"`
	Notes       string `json:"notes"`
	PhoneNumber string `json:"phone_number"`
}
type DeviceState struct {
	Drive_status string `json:"drive_status"`
}
type datastore struct {
	m map[string]Device
	*sync.RWMutex
}

type deviceHandler struct {
	store *datastore
}

func (h *deviceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	switch {
	case r.Method == http.MethodGet && listDevicesRe.MatchString(r.URL.Path):
		w.Header().Set("content-type", "application/json")
		h.List(w, r)
		return
	case r.Method == http.MethodPut && updateDeviceRe.MatchString(r.URL.Path):
		h.Update(w, r)
		return
	case r.Method == http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	default:
		notFound(w, r)
		return
	}
}

func (h *deviceHandler) List(w http.ResponseWriter, r *http.Request) {
	h.store.RLock()
	devices := make([]Device, 0, len(h.store.m))
	for _, v := range h.store.m {
		devices = append(devices, v)
	}
	h.store.RUnlock()
	jsonBytes, err := json.Marshal(devices)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *deviceHandler) Update(w http.ResponseWriter, r *http.Request) {
	matches := updateDeviceRe.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		notFound(w, r)
		return
	}

	h.store.RLock()
	_, ok := h.store.m[matches[1]]
	h.store.RUnlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("device not found"))
		return
	}

	var postDevice Device
	if err := json.NewDecoder(r.Body).Decode(&postDevice); err != nil {
		internalServerError(w, r)
		return
	}

	h.store.RLock()
	h.store.m[matches[1]] = postDevice
	h.store.RUnlock()

	w.WriteHeader(http.StatusOK)

}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}

func UnmarshalDevices(r *http.Response) []Device {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var deviceResponse DeviceResponse
	err = json.Unmarshal(body, &deviceResponse)
	if err != nil {
		panic(err)
	}

	return deviceResponse.Result_list
}

func main() {
	mux := http.NewServeMux()
	devicesResponse, err := http.Get("https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key=" + apiKey)

	if err != nil {
		log.Fatal(err)
	}

	defer devicesResponse.Body.Close()

	devices := UnmarshalDevices(devicesResponse)
	var mappedDevices = make(map[string]Device)
	for _, device := range devices {
		mappedDevices[device.Device_id] = device
	}

	deviceH := &deviceHandler{
		store: &datastore{
			m:       mappedDevices,
			RWMutex: &sync.RWMutex{},
		},
	}
	mux.Handle("/devices", deviceH)
	mux.Handle("/devices/", deviceH)

	http.ListenAndServe("localhost:8080", mux)
}
