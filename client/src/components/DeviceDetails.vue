
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
const phoneRules = [(value) => {
    const phoneRE = /^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$/im;
    return phoneRE.test(value) || 'Phone number is invalid'
}]

</script>
<template>
    <v-card :title="device.display_name" width="400" :prepend-avatar="device.user_avatar"
        :subtitle="device.latest_accurate_device_point.device_state.drive_status">
        <v-card-text>

            <EditableField title="Notes" emptyMessage="Enter Notes" :value="device.device_ui_settings.notes"
                @save="onSaveDeviceNotes">
            </EditableField>


            <EditableField title="Phone Number" emptyMessage="Enter Phone Number" :rules="phoneRules"
                :value="device.device_ui_settings.phone_number" @save="onSaveDeviceNotes">
            </EditableField>
        </v-card-text>
    </v-card>
</template>