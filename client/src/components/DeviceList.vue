
<script setup>
import { computed } from 'vue'
const props = defineProps({
    devices: Object,
})
const emit = defineEmits(['listClick'])


const getOnlineDevices = computed(() => {
    return props.devices.filter((d) => d.latest_accurate_device_point.lat !== 0 && d.latest_accurate_device_point.lat !== 0)
})
const getOfflineDevices = computed(() => {
    return props.devices.filter((d) => d.latest_accurate_device_point.lat === 0 && d.latest_accurate_device_point.lat === 0)
})


</script>
<template>
    <v-navigation-drawer permanent>
        <v-list>
            <v-list-item
                :class="[d.latest_accurate_device_point.device_state.drive_status === 'off' ? 'stripe-red' : 'stripe-green']"
                v-for="d in getOnlineDevices" :prepend-avatar="d.user_avatar" :title="d.display_name" :value="d.device_id"
                @click="emit('listClick', device)">

                <template v-slot:append>
                    <v-btn icon="mdi-eye" variant="plain"></v-btn>
                </template>
            </v-list-item>
            <v-list-subheader>Offline devices</v-list-subheader>
            <v-list-item disabled v-for="device in getOfflineDevices" :prepend-avatar="device.user_avatar"
                :title="device.display_name" :value="device.device_id" @click="emit('listClick', device)">

            </v-list-item>
        </v-list>
    </v-navigation-drawer>
</template>
<style scoped>
.v-list {
    padding: 0;
}

.stripe-red {
    border-left: red 5px solid;
}

.stripe-green {
    border-left: green 5px solid;
}</style>