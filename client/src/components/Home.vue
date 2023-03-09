<template>
  <v-container class="fill-height">
    <v-responsive class="d-flex align-center text-center fill-height">


      <div v-for="device in getDevices">
        {{ device.device_id }}
      </div>

      <GoogleMapLoader :mapConfig="{
    center: {
      lat: 0,
      lng: 0
    },
    zoom: 4
  }" apiKey="AIzaSyCao8Qkxu-lc3_Xywqg3kLfI2umRAwkmVs" :markers="getMarkers"></GoogleMapLoader>

      <div class="py-14" />

      <v-row class="d-flex align-center justify-center">
        <v-col cols="auto">

        </v-col>

        <v-col cols="auto">

        </v-col>

        <v-col cols="auto">

        </v-col>
      </v-row>
    </v-responsive>
  </v-container>
</template>

<script >

import { useAppStore } from '@/store/app';
import GoogleMapLoader from './GoogleMapLoader.vue'
const store = useAppStore();

export default {
  components: {
    GoogleMapLoader
  },
  mounted() {
    store.fetchDevices();
  },
  computed: {
    getDevices() {
      return store.getDevices
    },
  getMarkers(){
    return store.getMarkers
  }
  }

}


</script>
