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
          path: '/project/list',
          name: 'project',
          component: () => import('@/views/project/List.vue')
        },
        {
          path: '/env/tool',
          name: 'tool',
          component: () => import('@/views/env/Tool.vue')
        },{
          path: '/env/crud',
          name: 'crud',
          component: () => import('@/views/env/Crud.vue')
        },
      ]
    }
  ]
})

export default router