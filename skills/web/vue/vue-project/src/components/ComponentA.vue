<script setup>
import { ref } from 'vue';

const pros = defineProps({
    countP: {type: Number, default: 0}
})

// const count = ref(0);
const count = ref(pros.countP)
</script>

<template>
    <button @click="count++"> You clicked me {{count}} times.</button>
</template>

<style scoped>

</style>