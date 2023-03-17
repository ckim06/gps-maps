
<script setup>
import { useAppStore } from '@/store/app'
import EditableField from './EditableField.vue'
const store = useAppStore()

const props = defineProps({
    device: Object,
})
const emit = defineEmits(['save'])

const onSaveDeviceNotes = async (value) => {
    const device = props.device
    device.device_ui_settings.notes = value
    const resp = await store.saveDevice(device)
}

</script>
<template>
    <v-card :title="device.display_name" width="400" :prepend-avatar="device.user_avatar"
        :subtitle="device.latest_accurate_device_point.device_state.drive_status">
        <v-card-text>
            <span class="text-body-2">Notes</span>
           
            <EditableField emptyMessage="Enter Notes" :value="device.device_ui_settings.notes" @save="onSaveDeviceNotes">
            </EditableField>
           <p></p>
            <span class="text-body-2">Phone Number</span>
            
            <EditableField emptyMessage="Enter Phone Number" :value="device.device_ui_settings.phone_number" @save="onSaveDeviceNotes">
            </EditableField>
        </v-card-text>
    </v-card>
</template>