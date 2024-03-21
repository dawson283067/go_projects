<template>
  <div class="content">
    <!-- Login Form表单 -->

    <a-form class="login-form" :rules="formRules" :model="form" @submit="handleSubmit">
      <div class="title">欢迎登录博客系统</div>
      <a-form-item hide-label field="username" label="">
        <a-input v-model="form.username" placeholder="请输入用户名...">
          <template #prefix>
            <icon-user />
          </template>
        </a-input>
      </a-form-item>
      <a-form-item hide-label field="password" label="">
        <a-input-password v-model="form.password" placeholder="请输入用户密码...">
          <template #prefix>
            <icon-lock />
          </template>
        </a-input-password>
      </a-form-item>
      <a-form-item hide-label field="is_member">
        <a-checkbox v-model="form.is_member"> 记住 </a-checkbox>
      </a-form-item>
      <a-form-item hide-label>
        <a-button type="primary" html-type="submit" style="width: 100%">登录</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const form = ref({
  username: '',
  password: '',
  is_member: false
})

const handleSubmit = (data) => {
  // 调用后端api，对接后端
  console.log(data)
}

//
const formRules = {
  username: [
    {
      required: true,
      message: '请输入用户名'
    }
  ],
  password: [
    {
      required: true,
      message: '请输入密码'
    },
    {
      minLength: 6,
      message: '密码不能少于6位'
    }
  ]
}
</script>

<style lang="css" scoped>
.content {
  height: 100%;
  width: 100%;
  display: flex;
  justify-content: center;
}

.login-form {
  height: 100%;
  width: 500px;
  display: flex;
  justify-content: center;
}

.title {
  display: flex;
  justify-content: center;
  font-size: 20px;
  font-weight: 600;
  margin-bottom: 12px;
}
</style>
