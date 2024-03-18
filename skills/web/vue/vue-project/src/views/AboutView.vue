<template>
    <div>
        <button @click="count++">Your clicked me {{count}} times.</button>
        <!-- <div>{{ name.split('').reverse().join('')}}</div> -->
        <div v-bind:id="name">{{ reverseName }}</div>
        <button :style="buttonStyle" :disabled="isButtonDisabled" v-on:click="clickButton">button</button>

        <!-- view -> inputMessage -->
        <input v-model="inputMessage" placeholder="edit me"/>
        <!-- inputMessage -> view -->
        {{inputMessage}}

        <ul>
            <li v-for="(value) in names" :key="value">{{value}}</li>
        </ul>
        <ComponentA :countP="20"></ComponentA>
    </div>
</template>

<script setup>
import {ref,computed} from "vue";

// 引入组件：HTML <ComponentA></ComponentA> <component-a></component-a>
// import ComponentA from './ComponentA.vue'

const count = ref(0)
count.value = 20
const name = ref('abcdefghijk')

// 一个计算属性
const reverseName = computed(() => {return name.value.split('').reverse().join('')})

const buttonClass = ref({"red": true, "blue": false})
const buttonStyle = ref({"color":"red"})
const isButtonDisabled = ref(false)
const clickButton = () => {
    isButtonDisabled.value = !isButtonDisabled.value
    buttonClass.value.red = false
    buttonClass.value.blue = true
    buttonStyle.value.color = "blue"
}

const inputMessage = ref('默认值')

const names = ref(['张三','李四'])
</script>

<style lang="css" scoped>
.red {
    color: red;
}

.blue {
    color: blue;
}
</style>