<script setup>
// import { ref } from 'vue';

// 定义组件的属性
const props = defineProps({
    modelValue: {type: Number, default: 0}
})

// 定义组件的事件
// 父组件可以监听事件，子组件可以抛出事件，相当于在子组件中定义了一个接口
const emits = defineEmits(['update:modelValue'])

// const count = ref(0);
// const count = ref(props.countP)

const handleClick = () => {
    // count.value++
    // 抛出countChanged事件：count.value
    emits('update:modelValue', props.modelValue + 1)
}

</script>

<template>
    <button @click="handleClick"> You clicked me {{modelValue}} times.</button>
</template>

<style scoped>

</style>