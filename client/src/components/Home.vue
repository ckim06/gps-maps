
<script setup>
import { ref, onMounted, computed } from 'vue'
import { useAppStore } from '@/store/app'
import GoogleMap from './GoogleMap.vue'
import DeviceList from './DeviceList.vue'
import { useDisplay } from 'vuetify'

const { mdAndDown } = useDisplay()
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
  <v-layout>
    <v-system-bar>
      <v-icon icon="mdi-menu"></v-icon>
    </v-system-bar>
    <v-theme-provider theme="dark" v-if="mdAndDown">
      <v-bottom-navigation>
        <v-btn @click="selectedRail = 'devices'"><v-icon>mdi-devices</v-icon>DEVICES</v-btn>
        <v-btn @click="selectedRail = 'history'"><v-icon>mdi-history</v-icon>HISTORY</v-btn>
        <v-btn @click="selectedRail = 'places'"><v-icon>mdi-map-marker</v-icon>PLACES</v-btn>
        <v-btn @click="selectedRail = 'alerts'"><v-icon>mdi-bell-outline</v-icon>ALERTS</v-btn>
        <v-btn @click="selectedRail = 'reports'"><v-icon>mdi-clipboard</v-icon>REPORTS</v-btn>
      </v-bottom-navigation>
    </v-theme-provider>
    <v-navigation-drawer theme="dark" rail class="primary-rail" v-if="!mdAndDown">
      <v-divider></v-divider>
      <v-list nav>
        <v-list-item prepend-icon="mdi-devices" title="devices" @click="selectedRail = 'devices'"> <v-tooltip
            activator="parent" location="end">DEVICES</v-tooltip></v-list-item>
        <v-list-item prepend-icon="mdi-history" @click="selectedRail = 'history'">
          <v-tooltip activator="parent" location="end">HISTORY</v-tooltip>
        </v-list-item>
        <v-list-item prepend-icon="mdi-map-marker" @click="selectedRail = 'places'">
          <v-tooltip activator="parent" location="end">PLACES</v-tooltip>
        </v-list-item>
        <v-list-item prepend-icon="mdi-bell-outline" @click="selectedRail = 'alerts'">
          <v-tooltip activator="parent" location="end">ALERTS</v-tooltip>
        </v-list-item>
        <v-list-item prepend-icon="mdi-clipboard" @click="selectedRail = 'reports'">
          <v-tooltip activator="parent" location="end">REPORTS</v-tooltip>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-navigation-drawer class="secondary-nav-drawer" app :location="mdAndDown ? 'bottom' : 'start'">
      <v-list lines="two" v-if="selectedRail === 'devices'">
        <DeviceList :devices="getDevices" @listClick="onListClick" :highlightedDevice="highlightedDevice"
          @toggleVisibility="onToggleVisibility"></DeviceList>
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
</template>
<style lang="scss" scoped>
// cleans up a gap at the top of the navigation list
.secondary-nav-drawer .v-list {
  padding: 0;
}
// Show list items in a grid instead of list
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
</style>