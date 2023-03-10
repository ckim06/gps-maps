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
        .map((device) => device.latest_accurate_device_point)
    },
  },
  actions: {
    async fetchDevices() {
      try {
        const data = await axios.get('http://localhost:8080/devices')
        this.devices = data.data
      }
      catch (error) {
        console.log(error)
      }
    }
  },

})


