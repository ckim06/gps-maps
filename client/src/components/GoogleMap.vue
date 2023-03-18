<script setup>
import { ref, onMounted, watch, computed, nextTick } from 'vue'
import DeviceDetails from './DeviceDetails.vue'

const props = defineProps({
    markers: Object,
    selectedDevice: Object,
    showFitCenterBtn: Boolean
})
let mapObj
const gmap = ref(0)
const openWindows = ref({})
const showFitCenterBtnLocal = ref(false);

const emit = defineEmits(['markerOver'])

const getMarkers = computed(() => {
    return props.markers.filter((m) => !m.device_ui_settings.is_hidden)
})
onMounted(() => {
    showFitCenterBtnLocal.value = props.showFitCenterBtn;
    gmap.value.$mapPromise.then((mapObject) => {
        mapObj = mapObject
        fitMarkerBounds()
    })
})

const fitMarkerBounds = () => {
    if (mapObj) {
        const bounds = new google.maps.LatLngBounds()
        props.markers
            .filter((m) => !m.device_ui_settings.is_hidden)
            .forEach((marker) => {
                bounds.extend(marker.latest_accurate_device_point)
            })
        gmap.value.fitBounds(bounds)
        showFitCenterBtnLocal.value = false
    }
}
const close = (deviceId) => {
    openWindows.value[deviceId] = false
}
const markerClick = (marker) => {
    if (marker) {
        // close open windows
        Object.keys(openWindows.value).forEach((deviceId)=>{
            if(marker.device_id !== deviceId){
                openWindows.value[deviceId] = false
            }
        })

        mapObj.setCenter(marker.latest_accurate_device_point)
        mapObj.setZoom(10)
        nextTick(() => {
            openWindows.value[marker.device_id] = true
        });
    }
}
watch(() => props.selectedDevice, (currentValue) => {
    openWindows.value[currentValue.device_id] = null
    markerClick(currentValue)
})
watch(() => props.showFitCenterBtn, (currentValue) => {
    showFitCenterBtnLocal.value = true
})

</script>
<template>
    <v-snackbar v-model="showFitCenterBtnLocal" timeout="5000">
        <v-btn variant="text" @click="fitMarkerBounds" class="text-decoration-underline">Recenter around visible devices? </v-btn>

        <template v-slot:actions>
            <v-btn icon="mdi-close" variant="text" @click="showFitCenterBtnLocal = false">

            </v-btn>
        </template>
    </v-snackbar>
    <GMapMap :options="{
        zoomControl: true,
        mapTypeControl: false,
        scaleControl: true,
        streetViewControl: false,
        rotateControl: false,
        fullscreenControl: false
    }" :center="{ lat: 0, lng: 0 }" :zoom="7" ref="gmap" class="map">
        <GMapMarker :key="index" v-for="(marker, index) in getMarkers" :position="marker.latest_accurate_device_point"
            @click="markerClick(marker)" @closeclick="close(marker.device_id)" @mouseover="emit('markerOver', marker)">
            <GMapInfoWindow :closeclick="true" @closeclick="close(marker.device_id)"
                :opened="openWindows[marker.device_id] ?? false">
                <div>
                    <DeviceDetails :device="marker"></DeviceDetails>
                </div>
            </GMapInfoWindow>
        </GMapMarker>
    </GMapMap>
</template>
<style scoped>
.map {
    height: 100%;
    width: 100%;
}
</style>