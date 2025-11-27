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
    path: '/supplier-demo',
    name: 'SupplierDemo',
    meta: {
      title: '供应商管理演示',
      client: true
    },
    component: () => import('@/view/supplier/index.vue')
  },
  {
    path: '/supplier-approval',
    name: 'SupplierApproval',
    meta: {
      title: '供应商审批',
      client: true
    },
    component: () => import('@/view/supplier/approval/index.vue')
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
