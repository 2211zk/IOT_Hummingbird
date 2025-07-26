import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    redirect: '/login'
  },
  {
    path: '/init',
    name: 'Init',
    component: () => import('@/view/init/index.vue')
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/view/login/index.vue')
  },
  {
    path: '/scanUpload',
    name: 'ScanUpload',
    meta: {
      title: '扫码上传',
      client: true
    },
    component: () => import('@/view/example/upload/scanUpload.vue')
  },
  {
    path: '/system/department',
    name: 'DepartmentManagement',
    meta: { 
      title: '部门管理',
      icon: 'office-building',
      keepAlive: true
    },
    component: () => import('@/view/system/department/index.vue')
  },
  {
    path: '/system/device',
    name: 'DeviceManagement',
    meta: { 
      title: '设备管理',
      icon: 'monitor',
      keepAlive: true
    },
    component: () => import('@/view/system/device/index.vue')
  },
  {
    path: '/:catchAll(.*)',
    meta: {
      closeTab: true
    },
    component: () => import('@/view/error/index.vue')
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
