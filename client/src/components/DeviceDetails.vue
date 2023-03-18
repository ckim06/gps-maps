
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
const phoneRules = [
    (value)=> !! value || 'Enter a phone number', 
    (value) => {
    const phoneRE = /^[0-9]{3}-[0-9]{3}-[0-9]{4}$/im;
    return phoneRE.test(value) || 'Enter in this format 800-555-1212'
}]

</script>
<template>
    <v-row class="ma-0">
        <v-card :title="device.display_name" width="400" :prepend-avatar="device.user_avatar"
            :subtitle="device.latest_accurate_device_point.device_state.drive_status">
            <v-card-text>
                <EditableField title="Phone Number" emptyMessage="Enter Phone Number" :rules="phoneRules"
                    :value="device.device_ui_settings.phone_number" @save="onSaveDeviceNotes">
                </EditableField>
                <v-spacer class="mb-2"></v-spacer>
                <EditableField title="Notes" emptyMessage="Enter Notes" :value="device.device_ui_settings.notes"
                    @save="onSaveDeviceNotes">
                </EditableField>
            </v-card-text>
        </v-card>
    </v-row>
</template>