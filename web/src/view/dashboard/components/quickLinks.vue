<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->
<template>
  <div class="mt-8 w-full">
    <!-- 系统管理功能 -->
    <div class="mb-8">
      <h3 class="text-sm font-medium text-gray-600 mb-4">系统管理</h3>
      <div class="grid grid-cols-2 md:grid-cols-3 3xl:grid-cols-4">
        <div
          v-for="(item, index) in shortcuts"
          :key="index"
          class="flex flex-col items-center mb-3 group cursor-pointer"
          @click="toPath(item)"
        >
          <div
            class="w-8 h-8 rounded bg-gray-200 dark:bg-slate-500 flex items-center justify-center group-hover:bg-blue-400 group-hover:text-white"
          >
            <el-icon><component :is="item.icon" /></el-icon>
          </div>
          <div class="text-xs mt-2 text-gray-700 dark:text-gray-300">
            {{ item.title }}
          </div>
        </div>
      </div>
    </div>

    <!-- 业务功能 -->
    <div>
      <h3 class="text-sm font-medium text-gray-600 mb-4">业务功能</h3>
      <div class="grid grid-cols-2 md:grid-cols-3 3xl:grid-cols-5">
        <div
          v-for="(item, index) in businessFunctions"
          :key="index"
          class="flex flex-col items-center mb-3 group cursor-pointer p-3 rounded-lg border border-gray-200 hover:border-blue-400 transition-all hover:shadow-md"
          @click="item.isNewTab ? openLink(item) : toPath(item)"
        >
          <div
            class="w-10 h-10 rounded bg-blue-100 flex items-center justify-center group-hover:bg-blue-500 group-hover:text-white mb-2"
          >
            <el-icon class="text-lg"><component :is="item.icon" /></el-icon>
          </div>
          <div class="text-sm font-medium text-gray-700 dark:text-gray-300">
            {{ item.title }}
          </div>
          <div class="text-xs text-gray-500 dark:text-gray-400">
            {{ item.desc }}
          </div>
        </div>
      </div>
    </div>

    <!-- 最近访问 -->
    <div class="mt-8">
      <h3 class="text-sm font-medium text-gray-600 mb-4">最近访问</h3>
      <div class="grid grid-cols-2 md:grid-cols-3 3xl:grid-cols-4">
        <div
          v-for="(item, index) in recentVisits"
          :key="index"
          class="flex flex-col items-center mb-3 group cursor-pointer"
          @click="openLink(item)"
        >
          <div
            class="w-8 h-8 rounded bg-gray-200 dark:bg-slate-500 flex items-center justify-center group-hover:bg-blue-400 group-hover:text-white"
          >
            <el-icon><component :is="item.icon" /></el-icon>
          </div>
          <div class="text-xs mt-2 text-gray-700 dark:text-gray-300">
            {{ item.title }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup>
  import {
    Menu,
    Link,
    User,
    Service,
    Document,
    Reading,
    Files,
    Memo,
    Shop,
    Check,
    Setting,
    List
  } from '@element-plus/icons-vue'
  import { useRouter } from 'vue-router'
  const router = useRouter()

  const toPath = (item) => {
    router.push({ name: item.path })
  }

  const openLink = (item) => {
    window.open(item.path, '_blank')
  }
  const shortcuts = [
    {
      icon: Menu,
      title: '菜单管理',
      path: 'menu'
    },
    {
      icon: Link,
      title: 'API管理',
      path: 'api'
    },
    {
      icon: Service,
      title: '角色管理',
      path: 'authority'
    },
    {
      icon: User,
      title: '用户管理',
      path: 'user'
    },
    {
      icon: Files,
      title: '自动化包',
      path: 'autoPkg'
    },
    {
      icon: Memo,
      title: '自动代码',
      path: 'autoCode'
    }
  ]

  // 业务功能区域
  const businessFunctions = [
    {
      icon: Document,
      title: '我的询价',
      desc: '销售录入',
      path: '/inquiry-demo?menu=my-inquiry',
      isNewTab: true
    },
    {
      icon: Shop,
      title: '待报价项目',
      desc: '采购报价',
      path: '/inquiry-demo?menu=pending-quote',
      isNewTab: true
    },
    {
      icon: List,
      title: '询报价台账',
      desc: '全部记录',
      path: '/inquiry-demo?menu=inquiry-ledger',
      isNewTab: true
    },
    {
      icon: Check,
      title: '支付申请',
      desc: '吧盛支付',
      path: '/inquiry-demo?menu=payment-approval',
      isNewTab: true
    },
    {
      icon: Setting,
      title: '协议价管理',
      desc: '管理员',
      path: '/inquiry-demo?menu=agreement-price',
      isNewTab: true
    }
  ]

  
  const recentVisits = [
    {
      icon: Reading,
      title: '帮助文档',
      path: 'https://www.bashengkeji.com/docs'
    },
    {
      icon: Document,
      title: '关于我们',
      path: 'https://www.bashengkeji.com'
    }
  ]
</script>

<style scoped lang="scss"></style>
