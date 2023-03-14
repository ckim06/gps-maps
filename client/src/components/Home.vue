
<script setup>
import { ref, onMounted ,computed} from 'vue'
import { useAppStore } from '@/store/app';
import GoogleMap from './GoogleMap.vue'
import DeviceList from './DeviceList.vue'
const store = useAppStore();
let dataReady = ref(false)
onMounted(async () => {
  await store.fetchDevices();
  dataReady.value = true
})
const getDevices = computed(() => {
  return store.getDevices
})
const getMarkers = computed(() => {
  return store.getMarkers
})

</script>
<template>
  <v-card>
    <v-layout>
      <v-navigation-drawer theme="dark" rail permanent>
        <v-list-item nav prepend-avatar="https://randomuser.me/api/portraits/women/75.jpg"></v-list-item>

        <v-divider></v-divider>

        <v-list density="compact" nav>
          <v-list-item prepend-icon="mdi-view-dashboard" value="dashboard"></v-list-item>

          <v-list-item prepend-icon="mdi-forum" value="messages"></v-list-item>
        </v-list>
      </v-navigation-drawer>
      <DeviceList :v-if="dataReady" :devices="getDevices"></DeviceList>

      <v-main>
        <GoogleMap :markers="getMarkers">
        </GoogleMap>
      </v-main>
    </v-layout>
  </v-card>
</template>
