import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)

// 注册全局组件
import ComponentA from './components/ComponentA.vue'
app.component('ComponentA', ComponentA)

app.mount('#app')
