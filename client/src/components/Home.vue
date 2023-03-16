
<script setup>
import { ref, onMounted, computed } from 'vue'
import { useAppStore } from '@/store/app'
import GoogleMap from './GoogleMap.vue'
import DeviceList from './DeviceList.vue'
const store = useAppStore()
let dataReady = ref(false)
let clickedDevice = ref(null)
let selectedRail = ref('devices')
let showFitCenterBtn = ref(false)
let highlightedDevice = ref(null)
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

const onMarkerOver = (device) => {
  highlightedDevice.value = device
}

const onToggleVisibility = async (device) => {
  device.device_ui_settings.is_hidden = !device.device_ui_settings.is_hidden
  const resp = await store.saveDevice(device)
  showFitCenterBtn.value = !showFitCenterBtn.value
}
</script>
<template>
  <v-card>
    <v-layout>
      <v-system-bar>
        <v-icon icon="mdi-menu"></v-icon>
      </v-system-bar>
      <v-navigation-drawer theme="dark" rail class="primary-rail" location="bottom">
        <v-divider></v-divider>
        <v-list nav>
          <v-list-item prepend-icon="mdi-devices" @click="selectedRail = 'devices'" value="Devices"></v-list-item>
          <v-list-item prepend-icon="mdi-history" @click="selectedRail = 'history'" value="History"></v-list-item>
          <v-list-item prepend-icon="mdi-map-marker" @click="selectedRail = 'places'" value="Places"></v-list-item>
          <v-list-item prepend-icon="mdi-bell-outline" @click="selectedRail = 'alerts'" value="Alerts"></v-list-item>
          <v-list-item prepend-icon="mdi-clipboard" @click="selectedRail = 'reports'" value="Reports"></v-list-item>
        </v-list>
      </v-navigation-drawer>
      <v-navigation-drawer class="secondary-nav-drawer" location="bottom">
        <v-list v-if="selectedRail === 'devices'">
          <DeviceList :devices="getDevices" @listClick="onListClick" :highlightedDevice="highlightedDevice" @toggleVisibility="onToggleVisibility"></DeviceList>
        </v-list>

        <v-list v-if="selectedRail === 'history'">
          <h3>History</h3>Out of scope
        </v-list>
        <v-list v-if="selectedRail === 'places'">
          <h3>Places</h3>Out of scope
        </v-list>
        <v-list v-if="selectedRail === 'alerts'">
          <h3>Alerts</h3>Out of scope
        </v-list>
        <v-list v-if="selectedRail === 'reports'">
          <h3>Reports</h3>Out of scope
        </v-list>

      </v-navigation-drawer>


      <v-main class="h-screen v-screen">
        <GoogleMap :markers="getMarkers" :selectedDevice="clickedDevice" :showFitCenterBtn="showFitCenterBtn"
          @markerOver="onMarkerOver">
        </GoogleMap>
      </v-main>
    </v-layout>
  </v-card>
</template>
<style lang="scss" scoped>
.secondary-nav-drawer .v-list {
  padding: 0;
}

:deep(.secondary-nav-drawer.v-navigation-drawer--bottom) {
  .v-list {
    display: flex;
    flex-wrap: wrap;

    .v-list-item {
      flex-basis: 50%;

    }

    .v-list-subheader {
      flex: 0 0 100%;
      padding: 20px 0;
    }
  }
}

:deep(.primary-rail.v-navigation-drawer--bottom) {
  .v-list {
    display: flex;
    flex-wrap: wrap;

    .v-list-item {
      flex-basis: 20%;
      justify-content: center;
    }
  }
}
</style>