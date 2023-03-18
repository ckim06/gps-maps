
<script setup>
import { ref, onMounted } from 'vue'

const props = defineProps({
    value: String,
    emptyMessage: String,
    rules: Object,
    title: String
})

const emit = defineEmits(['save'])
const editable = ref(false)
const localValue = ref('')
const loading = ref(false)
const form = ref(0)

onMounted(() => {
    localValue.value = props.value
})

const edit = () => {
    editable.value = true
}
const save = async () => {
    const { valid } = await form.value.validate()
    if (valid) {
        loading.value = true
        emit("save", localValue.value)
        loading.value = false
        editable.value = false
    }
}
const cancel = () => {
    editable.value = false
}

</script>
<template>
    <div class="font-weight-bold">{{ title }}: &nbsp;</div>
    <div class="textarea-clickable" v-if="!editable" @click="edit">
        <span v-if="localValue">{{ localValue }}</span>
        <span class="font-italic" v-if="!localValue">{{ emptyMessage }}</span>
        <v-btn variant="plain" class="pl-2" size="sm" icon="mdi-pencil"></v-btn>
    </div>
    <div class="editable-field-wrapper" v-if="editable">
        <v-form ref="form" @submit.prevent="save">
            <v-text-field :rules="rules" :loading="loading" v-model="localValue" variant="solo" density="compact">
                <template v-slot:append>
                    <v-btn variant="plain" icon="mdi-check" type="submit"></v-btn>
                    <v-btn icon="mdi-cancel" variant="plain" @click="cancel"></v-btn>
                </template>
            </v-text-field>
        </v-form>
    </div>
</template>
<style scoped>
.textarea-clickable {
    cursor: pointer;
}
</style>