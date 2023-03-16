
<script setup>
import { ref, onMounted, computed } from 'vue'
import { useAppStore } from '@/store/app'
import GoogleMap from './GoogleMap.vue'
import DeviceList from './DeviceList.vue'
const store = useAppStore()
let dataReady = ref(false)
let clickedDevice = ref(null)

onMounted(async () => {
  await store.fetchDevices()
  dataReady.value = true
})
const getDevices = computed(() => {
  return store.getDevices
})
const getMarkers = computed(() => {
  return store.getMarkers
})

const onListClick = (device) => {
  clickedDevice.value = device
}

const  onToggleVisibility = async (device) => {
  device.device_ui_settings.is_hidden = !device.device_ui_settings.is_hidden
  const resp = await store.saveDevice(device)
}
</script>
<template>
  <v-card>
    <v-layout>
      <v-system-bar>
        <v-icon icon="mdi-menu" class="ms-2"></v-icon>
      </v-system-bar>
      <v-navigation-drawer theme="dark" rail permanent>
        <v-divider></v-divider>
        <v-list nav>
          <v-list-item prepend-icon="mdi-devices" value="Devices"></v-list-item>
          <v-list-item prepend-icon="mdi-history" value="History"></v-list-item>
          <v-list-item prepend-icon="mdi-map-marker" value="Places"></v-list-item>
          <v-list-item prepend-icon="mdi-bell-outline" value="Alerts"></v-list-item>
          <v-list-item prepend-icon="mdi-clipboard" value="Reports"></v-list-item>
        </v-list>
      </v-navigation-drawer>
      <DeviceList :v-if="dataReady" :devices="getDevices" @listClick="onListClick" @toggleVisibility="onToggleVisibility"></DeviceList>
      <v-main>
        <GoogleMap :markers="getMarkers" :deviceFromList="clickedDevice">
        </GoogleMap>
      </v-main>
    </v-layout>
  </v-card>
</template>
