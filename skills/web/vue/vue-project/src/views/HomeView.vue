...
<template>
    <h1>输入框 View --- Model </h1>
    <br>
    <div>
        <!-- 单向修改 -->
        <!-- 用户名：<input type="text" @change="userInputChanged" :value="username"><br>
        密码：<input type="password" @change="passwordInputChanged" :value="password"><br> -->
        <!-- v-model 指令 负责视图和数据的双向绑定 -->
        用户名：<input type="text" v-model="username"> <br>
        密码：<input type="password" v-model="password"> <br>
        登录：<button>提交</button>
    </div>

    <h1>输入框变量展示 Model --- View </h1>
    <br>
    <div>
        用户名：{{ username }} <br>
        密码：{{ password }}
    </div>
</template>

<script setup>
import { getCurrentInstance, ref } from 'vue';

// 1. 页面的数据
var username = ref('admin')
var password = ref('123456')

// 2. 页面数据发生变化
// 用了 v-model 以下就不需要了。这是单向修改。
// var userInputChanged = (e) => {
//     // console.log(e.target.value)
//     username.value = e.target.value
// }
// var passwordInputChanged = (e) => {
//     password.value = e.target.value
// }

// 获取当前组件的view model实例, vue instance
const vm = getCurrentInstance()
console.log(vm);


// Vue实例生命周期的演示
// 以库的形式来使用vue实例提供的API
import {
  onBeforeMount,
  onMounted,
  onBeforeUpdate,
  onUpdated,
  onBeforeUnmount,
  onUnmounted,
} from "vue";

onBeforeMount(() => {
  console.log("before mount");
});
onMounted(() => {
  console.log("mounted");
});
onBeforeUpdate(() => {
  console.log("before update");
});
onUpdated(() => {
  console.log("on updated");
});
onBeforeUnmount(() => {
  console.log("before unmount");
});
onUnmounted(() => {
  console.log("unmounted");
});

</script>

<style lang="scss" scoped>

</style>