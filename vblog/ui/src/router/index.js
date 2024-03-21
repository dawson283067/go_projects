import { createRouter, createWebHistory } from 'vue-router'
// import HomeView from '../views/HomeView.vue'
import BackendLayout from '../views/backend/BackendLayout.vue'
import FrontendLayout from '../views/frontend/FrontendLayout.vue'

const router = createRouter({
  // /backend/page
  // /backend#page1
  // https://cn.vitejs.dev/guide/env-and-mode.html
  // 这是纯前端逻辑
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'LoginView',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      // 这是懒加载
      component: () => import('../views/LoginView.vue')
    },
    {
      path: '/backend',
      name: 'BackendLayout',
      component: BackendLayout
    },
    {
      path: '/frontend',
      name: 'FrontendLayout',
      component: FrontendLayout
    }
  ]
})

export default router
