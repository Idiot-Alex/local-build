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
          path: '/project',
          name: 'project',
          component: () => import('@/views/Project.vue')
        },
        {
          path: '/build-plan',
          name: 'build-plan',
          component: () => import('@/views/BuildPlan.vue')
        },
        {
          path: '/tool',
          name: 'tool',
          component: () => import('@/views/Tool.vue')
        },
      ]
    }
  ]
})

export default router