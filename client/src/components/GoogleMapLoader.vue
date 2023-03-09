<template>
    asdf
    <div id="map"></div>
</template>
<script >
import { Loader } from '@googlemaps/js-api-loader';

export default {

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
        this.loader = new Loader({
            apiKey: this.apiKey,
            version: "weekly",
            libraries: ["places"]
        });
        this.initializeMap()
    },

    methods: {
        async initializeMap() {
            try {
                const google = await this.loader.load();
                this.map = new google.maps.Map(document.getElementById("map"), this.mapConfig)
                this.markers.forEach((location) => {
                    const marker = new google.maps.Marker({
                        position: {lat:location.lat, lng:location.lng},
                        map: this.map,
                    });
                    marker.setMap(this.map);
                    this.map.
                })
            }
            catch (e) {
                console.error(e)
            }
        }
    }
}
</script>
<style>
#map {
    width: 200px;
    height: 200px;
}
</style>