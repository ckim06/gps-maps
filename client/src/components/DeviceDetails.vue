
<script setup>
import { useAppStore } from '@/store/app'
import EditableField from './EditableField.vue'
import { ref } from 'vue'
import {
    isValidPhoneNumber,
    AsYouType
} from 'libphonenumber-js'
const store = useAppStore()

const props = defineProps({
    device: Object,
})
const emit = defineEmits(['save'])

const phoneNumber = ref(props.device.device_ui_settings.phone_number)

// I debated moving this to home.vue so that all data interactions are in the 
// same place.  However, since this is embedded in the map, I think this should 
// stay here.
const onSaveDeviceNotes = async (value) => {
    const device = props.device
    device.device_ui_settings.notes = value
    const resp = await store.saveDevice(device)
}
const phoneRules = [
    (value) => !!value || 'Enter a phone number',
    (value) => {
        return isValidPhoneNumber(value, 'US') || 'Enter a valid US phone number'
    }]

const phoneKeyup = (e) => {
    phoneNumber.value = new AsYouType('US').input(e)
}

</script>
<template>
    <v-row class="ma-0">
        <v-card :title="device.display_name" width="400" :prepend-avatar="device.user_avatar"
            :subtitle="device.latest_accurate_device_point.device_state.drive_status">
            <v-card-text>
                <EditableField title="Phone Number" emptyMessage="Enter Phone Number" :rules="phoneRules"
                    :value="phoneNumber" :asYouType="phoneKeyup" @save="onSaveDeviceNotes">
                </EditableField>
                <v-spacer class="mb-2"></v-spacer>
                <EditableField title="Notes" emptyMessage="Enter Notes" :value="device.device_ui_settings.notes"
                    @save="onSaveDeviceNotes">
                </EditableField>
            </v-card-text>
        </v-card>
    </v-row>
</template>