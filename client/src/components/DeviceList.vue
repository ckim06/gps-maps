
<script setup>
const props = defineProps({
    devices: Object,
})
const emit = defineEmits(['listClick'])

</script>
<template>
    <v-navigation-drawer permanent>
        <v-list>
            <v-list-item v-for="device in devices" :title="device.display_name" :value="device.device_id" @click="emit('listClick', device)"></v-list-item>
        </v-list>
    </v-navigation-drawer>
</template>
