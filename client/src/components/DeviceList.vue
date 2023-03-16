
<script setup>

import { computed } from 'vue'
const props = defineProps({
    devices: Object,
})
const emit = defineEmits(['listClick', 'toggleVisiblity'])

const getOnlineDevices = computed(() => {
    return props.devices.filter((d) => d.latest_accurate_device_point.lat !== 0 && d.latest_accurate_device_point.lat !== 0)
})
const getOfflineDevices = computed(() => {
    return props.devices.filter((d) => d.latest_accurate_device_point.lat === 0 && d.latest_accurate_device_point.lat === 0)
})


</script>
<template>
    <v-navigation-drawer app permanent>
        <v-list>
            <v-list-item :class="{
                'stripe-red': d.latest_accurate_device_point.device_state.drive_status === 'off',
                'stripe-yellow': d.latest_accurate_device_point.device_state.drive_status === 'idle',
                'stripe-green': d.latest_accurate_device_point.device_state.drive_status !== 'off' && d.latest_accurate_device_point.device_state.drive_status !== 'idle'
            }" v-for="d in getOnlineDevices" :prepend-avatar="d.user_avatar" :value="d.device_id"
                @click="emit('listClick', d)" :subtitle="d.latest_accurate_device_point.device_state.drive_status">
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
        </v-list>
    </v-navigation-drawer>
</template>
<style scoped>
.v-list {
    padding: 0;
}

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