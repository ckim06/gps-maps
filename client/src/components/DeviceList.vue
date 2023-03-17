
<script setup>

import { computed, ref, watch } from 'vue'
const props = defineProps({
    devices: Object,
    highlightedDevice: Object
})
let highlightedDeviceLocal = ref(null)
const emit = defineEmits(['listClick', 'toggleVisibility'])

const getOnlineDevices = computed(() => {
    return props.devices.filter((d) => d.latest_accurate_device_point.lat !== 0 && d.latest_accurate_device_point.lat !== 0)
})
const getOfflineDevices = computed(() => {
    return props.devices.filter((d) => d.latest_accurate_device_point.lat === 0 && d.latest_accurate_device_point.lat === 0)
})
watch(() => props.highlightedDevice, (currentValue) => {
    highlightedDeviceLocal.value = currentValue

    setTimeout(() => {
        highlightedDeviceLocal.value = null
    }, 500);
})


</script>
<template>
    <v-list-item :active="d.device_id === highlightedDeviceLocal?.device_id" :class="{
        'stripe-red': d.latest_accurate_device_point.device_state.drive_status === 'off',
        'stripe-yellow': d.latest_accurate_device_point.device_state.drive_status === 'idle',
        'stripe-green': d.latest_accurate_device_point.device_state.drive_status !== 'off' && d.latest_accurate_device_point.device_state.drive_status !== 'idle',

    }" v-for="d in getOnlineDevices" :prepend-avatar="d.user_avatar" :value="d.device_id" @click="emit('listClick', d)"
        :subtitle="d.latest_accurate_device_point.device_state.drive_status">
        <template v-slot:title>
            {{ d.display_name }}

            <v-btn :icon="d.device_ui_settings.is_hidden ? 'mdi-eye-off' : 'mdi-eye'" size="small"
                @click.stop="emit('toggleVisibility', d)" variant="plain"></v-btn>
        </template>
    </v-list-item>
    <v-list-subheader>Offline devices</v-list-subheader>
    <v-list-item disabled v-for="d in getOfflineDevices" :prepend-avatar="d.user_avatar" :title="d.display_name"
        :value="d.device_id">

    </v-list-item>
</template>
<style scoped>
.stripe-red {
    border-left: #bb1e10 5px solid;
}

.stripe-yellow {
    border-left: #f7b500 5px solid;
}

.stripe-green {
    border-left: #57E964 5px solid;
}
</style>
