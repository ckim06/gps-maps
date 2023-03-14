
<script setup>
import { ref, onMounted } from 'vue'

const props = defineProps({
    value: String,
    emptyMessage:String
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
    <div class="textarea-clickable" v-if="!editable" @click="edit">
        <span :v-if="localValue" >{{ localValue  }}</span>
        <span class="font-italic" :v-if="!localValue" >{{ emptyMessage  }}</span>
        <v-btn variant="plain" icon="mdi-pencil"></v-btn>
    </div>

    <v-textarea v-if="editable" v-model="localValue" variant="solo">
        <template v-slot:append>
            <v-btn variant="plain" icon="mdi-check" @click="save"></v-btn>
            <v-btn icon="mdi-cancel" variant="plain" @click="cancel"></v-btn>
        </template>
    </v-textarea>
</template>
<style scoped>
.textarea-clickable{
    cursor: pointer;
}
</style>