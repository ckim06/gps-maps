<template>
    <GMapMap :center="{ lat: 1, lng: 1 }" :zoom="7" ref="gmap" class="h-screen v-screen">
        <GMapMarker v-for="marker in markers" :position="marker" >
            <GMapInfoWindow>
        <div>I am in info window {{ marker.lat }}
        </div>
      </GMapInfoWindow>
        </GMapMarker>
    </GMapMap>
</template>
<script >

export default {
    gmap: {},
    props: {
        mapConfig: Object,
        apiKey: String,
        markers: Object
    },

    data() {
        return {
            loader: null,
            map: null
        }
    },

    mounted() {
        this.$refs.gmap.$mapPromise.then((mapObject) => {
            mapObject.addListener('tilesloaded', ()=>{
                this.centerMap(this.markers, mapObject);
            })
           
        });
    },

    methods: {
        centerMap: (markers, gmap) => {
            const bounds = new google.maps.LatLngBounds();
            markers.forEach((location) => {
                bounds.extend(location)
            })
            gmap.fitBounds(bounds);
        }
    }
}
</script>
