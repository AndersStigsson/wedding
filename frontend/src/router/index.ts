import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Home from '../views/Home.vue'
import Toastmasters from '../components/ToastMasters.vue'
import Couple from '../components/Couple.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/toastmasters',
    name: 'Toastmasters',
    component: Toastmasters
  },
  {
    path: '/couple',
    name: 'Couple',
    component: Couple
  },
  {
    path: '/:catchAll(.*)',
    name: 'Home',
    component: Home
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
