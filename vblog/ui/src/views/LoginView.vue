<template>
  <div>
    <a-form :model="form" @submit="handleSubmit">
      <a-form-item field="username" label="用户名">
        <a-input v-model="form.username" placeholder="请输入用户名..." />
      </a-form-item>
      <a-form-item field="password" label="密码">
        <a-input-password v-model="form.password" placeholder="请输入用户密码..." />
      </a-form-item>
      <a-form-item field="is_member">
        <a-checkbox v-model="form.is_member"> 记住 </a-checkbox>
      </a-form-item>
      <a-form-item>
        <a-button html-type="submit">登录</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup>
import { ref } from 'vue';

const form = ref({
    username: '',
    password: '',
    is_member: false,
})

const handleSubmit = (data) => {
    console.log(data);
}
</script>

<style lang="css" scoped></style>
