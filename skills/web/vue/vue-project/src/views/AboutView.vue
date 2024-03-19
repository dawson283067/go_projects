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
        <!-- countP -->
        <!-- props: countP -->
        <!-- countChanged: func -->
        <!-- v-model:countP, countP是一个双向绑定属性 -->
        <!-- countP属性， -->
        <!-- countChangedHandler countP = e, 事件的名称：update:CountP，这个名字是vue约定俗成的-->
        <!-- 默认属性：modelValue ==> v-model:modelValue -->
        <ComponentA v-model="countP"></ComponentA><br><br>
        <span>鼠标位置：({{x}},{{y}})</span>

        <!-- 插槽 -->
        <LayoutView>
            <template #header>
                <div>Header</div>
            </template>
            <template #default>
                <div>Main</div>
            </template>
            <template #footer>
                <div>Footer</div>
            </template>
        </LayoutView>
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

// 父组件处理来自子组件触发的事件
const countP = ref(20)
// const countChangedHandler = (e) => {
//     countP.value = e
// }

// 使用组合式函数
import { useMouse } from "./mouse";
const {x,y} = useMouse()

// 引入带插槽的组件
import LayoutView from "./LayoutView.vue";
</script>

<style lang="css" scoped>
.red {
    color: red;
}

.blue {
    color: blue;
}
</style>