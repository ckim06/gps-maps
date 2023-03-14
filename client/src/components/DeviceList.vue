
<script setup>
const props = defineProps({
    devices: Object,
})

</script>
<template>
    <v-navigation-drawer permanent>
        <v-list>
            <v-list-item v-for="device in devices" :title="device.display_name" :value="device.device_id"></v-list-item>
        </v-list>
    </v-navigation-drawer>
</template>
