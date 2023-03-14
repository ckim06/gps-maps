<script setup>
import { ref, onMounted } from 'vue'
import DeviceDetails from './DeviceDetails.vue'

const props = defineProps({
    markers: Object,
})
let mapObj;
const gmap = ref(0)
let openWindow = ref(0)
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

const markerClick = (marker) => {
    mapObj.setCenter(marker.latest_accurate_device_point)
    mapObj.setZoom(15)
    openWindow.value = marker.device_id
}


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
        <GMapMarker v-for="marker in markers" :position="marker.latest_accurate_device_point" @click="markerClick(marker)">
            <GMapInfoWindow :opened="openWindow === marker.device_id">
                <div>
                    <DeviceDetails :device="marker"></DeviceDetails>
                </div>
            </GMapInfoWindow>
        </GMapMarker>
    </GMapMap>
</template>