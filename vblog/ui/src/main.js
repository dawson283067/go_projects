import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)

// 引入UI组件库
import ArcoVue from '@arco-design/web-vue'
// Icon是一个独立的组件库，需要单独引入
import ArcoVueIcon from '@arco-design/web-vue/es/icon'
// 同时引入arco.design的样式（该样式会覆盖之前的样式）
import '@arco-design/web-vue/dist/arco.css'
app.use(ArcoVue)
app.use(ArcoVueIcon)

app.mount('#app')
