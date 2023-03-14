
<script setup>
import { ref, onMounted } from 'vue'

const props = defineProps({
    value: String,
})
const emit = defineEmits(['save'])
let editable = ref(false)
let localValue = ref('')

onMounted(() => {
    localValue.value = props.value
})


const save = () => {
    emit("save", localValue.value)
    editable.value = false
}
const cancel = () => {
    editable.value = false
}

</script>
<template>
    <span>{{ localValue }}</span>
    <v-text-field v-model="localValue" variant="solo">
        <template v-slot:append>
            <v-btn variant="plain" icon="mdi-checkbox-marked" @click="save"></v-btn>

            <v-btn icon="mdi-cancel" variant="plain" @click="cancel"></v-btn>
        </template>
    </v-text-field>
</template>