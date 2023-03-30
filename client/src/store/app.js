// Utilities
import { defineStore } from 'pinia'
import axios from 'axios'
export const useAppStore = defineStore('app', {
  state: () => ({
    devices: []
  }),
  getters: {
    getDevices(state) {
      return state.devices
    },
    getMarkers(state) {
      return state.devices
        .filter((device) => device.latest_accurate_device_point.lat !== 0 && device.latest_accurate_device_point.lng !== 0)
    },
  },
  actions: {
    updateLocations(locationData) {
      this.devices.forEach((device)=>{
        device.latest_accurate_device_point = locationData[device.device_id]
      });
    },
    async fetchDevices() {
      try {
        const data = await axios.get('http://localhost:8080/devices/')
        this.devices = data.data
      }
      catch (error) {
        console.log(error)
      }
    },
    async saveDevice(device) {
      try {
        const data = await axios.put(`http://localhost:8080/devices/${device.device_id}`, device)
        
      }
      catch (error) {
        console.log(error)
      }
    },
  }
})


