<script setup>
import { ref, onMounted, watch, nextTick } from 'vue'
import DeviceDetails from './DeviceDetails.vue'

const props = defineProps({
    markers: Object,
    deviceFromList: Object
})
let mapObj;
const gmap = ref(0)
let openWindows = ref({})
onMounted(() => {
    gmap.value.$mapPromise.then((mapObject) => {
        mapObj = mapObject
        const bounds = new google.maps.LatLngBounds();
        props.markers.forEach((marker) => {
            bounds.extend(marker.latest_accurate_device_point)
        })
        gmap.value.fitBounds(bounds)

    })
})
const close = (deviceId) => {
    openWindows.value[deviceId] = false
}
const markerClick = (marker) => {
    if (marker) {
        mapObj.setCenter(marker.latest_accurate_device_point)
        mapObj.setZoom(10)
        nextTick(() => {
            openWindows.value[marker.device_id] = true
        });
    }
}
watch(() => props.deviceFromList, (currentValue, oldValue) => {
    openWindows.value[currentValue.device_id] = null
    markerClick(currentValue)
})

</script>
<template>
    <GMapMap :options="{
        zoomControl: true,
        mapTypeControl: false,
        scaleControl: true,
        streetViewControl: false,
        rotateControl: false,
        fullscreenControl: false
    }" :center="{ lat: 0, lng: 0 }" :zoom="7" ref="gmap" class="h-screen v-screen">
        <GMapMarker :key="index" v-for="(marker, index) in markers" :position="marker.latest_accurate_device_point"
             @click="markerClick(marker)" @closeclick="close(marker.device_id)">
            <GMapInfoWindow :closeclick="true" @closeclick="close(marker.device_id)" :opened="openWindows[marker.device_id] ?? false">
                <div>
                    <DeviceDetails :device="marker"></DeviceDetails>
                </div>
            </GMapInfoWindow>
        </GMapMarker>
    </GMapMap>
</template>