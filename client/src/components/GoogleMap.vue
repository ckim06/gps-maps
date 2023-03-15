<script setup>
import { ref, onMounted, watch, nextTick } from 'vue'
import DeviceDetails from './DeviceDetails.vue'

const props = defineProps({
    markers: Object,
    deviceFromList: Object
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
    if (!marker) {
        openWindow.value = null;
    } else {
        mapObj.setCenter(marker.latest_accurate_device_point)
        mapObj.setZoom(10)
        nextTick(() => {
            openWindow.value = marker.device_id
        });

    }

}
watch(() => props.deviceFromList, (currentValue, oldValue) => {
    openWindow.value = null
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
            @closeclick="markerClick()" @click="markerClick(marker)">
            <GMapInfoWindow :closeclick="true" @closeclick="markerClick()" :opened="openWindow === marker.device_id">
                <div>
                    <DeviceDetails :device="marker"></DeviceDetails>
                </div>
            </GMapInfoWindow>
        </GMapMarker>
    </GMapMap>
</template>