
<script setup>
import { useAppStore } from '@/store/app';
import EditableField from './EditableField.vue'
const store = useAppStore();

const props = defineProps({
    device: Object,
})
const emit = defineEmits(['save'])

const onSaveDeviceNotes = (value) => {
    const device = props.device
    device.device_ui_settings.notes = value
    const resp = store.saveDevice(device);

}

</script>
<template>
    <v-card :title="device.display_name" width="400" :subtitle="device.latest_accurate_device_point.drive_status">
        <v-card-text>
            <EditableField :value="device.device_ui_settings.notes" @save="onSaveDeviceNotes"></EditableField>
        </v-card-text>
    </v-card>
</template>