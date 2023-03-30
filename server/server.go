package main

import (
	_ "embed"
	"encoding/json"
	"time"

	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"

	"regexp"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
)

//go:embed keys.ini
var apiKey string
var devicesURL = "https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key=" + apiKey
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
	User_avatar                  string         `json:"user_avatar"`
	Latest_accurate_device_point Location       `json:"latest_accurate_device_point"`
	Device_ui_settings           DeviceSettings `json:"device_ui_settings"`
}
type DeviceSettings struct {
	Is_hidden   bool   `json:"is_hidden"`
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
	case r.URL.RequestURI() == "/ws":
		h.serveWs(w, r)
		return
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

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func (h *deviceHandler) serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer ws.Close()
	for {
		for range time.Tick(time.Minute / 2) {
			deviceLocations := make(map[string]Location)
			devices := getDevices()

			for _, d := range devices {
				h.store.RLock()
				storeDevice := h.store.m[d.Device_id]
				storeDevice.Latest_accurate_device_point = d.Latest_accurate_device_point
				h.store.m[d.Device_id] = storeDevice
				h.store.RUnlock()

				deviceLocations[d.Device_id] = storeDevice.Latest_accurate_device_point
			}
			jsonBytes, err := json.Marshal(deviceLocations)
			if err != nil {
				internalServerError(w, r)
				return
			}
			ws.WriteMessage(1, jsonBytes)
		}
	}

}

func getDevices() []Device {
	devicesResponse, err := http.Get(devicesURL)

	if err != nil {
		log.Fatal(err)
	}

	defer devicesResponse.Body.Close()
	return UnmarshalDevices(devicesResponse)

}

func main() {

	mux := http.NewServeMux()
	devices := getDevices()
	mappedDevices := make(map[string]Device)

	genders := []string{
		"women",
		"men",
	}
	for _, device := range devices {
		gender := genders[rand.Intn(2)]
		device.User_avatar = "https://randomuser.me/api/portraits/" + gender + "/" + strconv.Itoa(rand.Intn(100)) + ".jpg"
		mappedDevices[device.Device_id] = device
	}

	deviceH := &deviceHandler{
		store: &datastore{
			m:       mappedDevices,
			RWMutex: &sync.RWMutex{},
		},
	}
	mux.Handle("/ws", deviceH)
	mux.Handle("/devices/", deviceH)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe("localhost:8080", mux)
}
