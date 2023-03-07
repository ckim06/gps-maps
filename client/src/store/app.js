// Utilities
import { defineStore } from 'pinia'
import axios from 'axios'
export const useAppStore = defineStore('app', {
  state: () => ({
    devices:[]
  }),
  getters: {
    getDevices(state){
        return state.devices
      }
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


