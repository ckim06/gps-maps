
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

const edit = () => {
    editable.value = true
}
const save = () => {
    emit("save", localValue.value)
    editable.value = false
}
const cancel = () => {
    editable.value = false
}

</script>
<template>
    <div v-if="!editable">
        <span >{{ localValue }}</span>
        <v-btn variant="plain" icon="mdi-pencil" @click="edit"></v-btn>
    </div>

    <v-text-field v-if="editable" v-model="localValue" variant="solo">
        <template v-slot:append>
            <v-btn variant="plain" icon="mdi-check" @click="save"></v-btn>
            <v-btn icon="mdi-cancel" variant="plain" @click="cancel"></v-btn>
        </template>
    </v-text-field>
</template>