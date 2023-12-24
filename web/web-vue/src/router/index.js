import { createRouter, createWebHashHistory } from 'vue-router'
import AppLayout from '@/layout/AppLayout.vue'

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: '/',
      component: AppLayout,
      children: [
        {
          path: '/',
          name: 'dashboard',
          component: () => import('@/views/Empty.vue')
        },
        {
          path: '/env/tool',
          name: 'tool',
          component: () => import('@/views/env/Tool.vue')
      },
      ]
    }
  ]
})

export default router