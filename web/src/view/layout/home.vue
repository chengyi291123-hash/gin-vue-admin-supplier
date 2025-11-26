<template>
  <div class="home-container">
    <!-- 顶部统计卡片 -->
    <div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-7 py-2 gap-4 md:gap-2 gva-container2">
      <gva-card custom-class="col-span-1 lg:col-span-2">
        <gva-chart :type="1" title="访问人数" />
      </gva-card>
      <gva-card custom-class="col-span-1 lg:col-span-2">
        <gva-chart :type="2" title="新增客户" />
      </gva-card>
      <gva-card custom-class="col-span-1 lg:col-span-2">
        <gva-chart :type="3" title="解决数量" />
      </gva-card>
      <gva-card
        title="快捷功能"
        show-action
        custom-class="col-start-1 md:col-start-3 lg:col-start-7 row-span-2"
      >
        <gva-quick-link />
      </gva-card>
    </div>

    <!-- 新增功能卡片区域 -->
    <div class="feature-cards mt-6">
      <el-row :gutter="20">
        <!-- 询价管理 -->
        <el-col :xs="24" :sm="24" :md="8" :lg="8" :xl="8">
          <el-card class="feature-card" shadow="hover">
            <template #header>
              <div class="card-header">
                <span class="card-title">
                  <el-icon><Document /></el-icon>
                  询价管理
                </span>
              </div>
            </template>
            <div class="card-content">
              <el-row :gutter="16">
                <el-col :span="12" v-for="item in inquiryMenus" :key="item.key">
                  <div class="menu-item" @click="goToPage(item.key)">
                    <el-icon><component :is="item.icon" /></el-icon>
                    <div class="menu-text">
                      <div class="menu-title">{{ item.title }}</div>
                      <div class="menu-desc">{{ item.desc }}</div>
                    </div>
                  </div>
                </el-col>
              </el-row>
            </div>
          </el-card>
        </el-col>

        <!-- 采购审批 -->
        <el-col :xs="24" :sm="24" :md="8" :lg="8" :xl="8">
          <el-card class="feature-card" shadow="hover">
            <template #header>
              <div class="card-header">
                <span class="card-title">
                  <el-icon><Check /></el-icon>
                  采购审批
                </span>
              </div>
            </template>
            <div class="card-content">
              <el-row :gutter="16">
                <el-col :span="12" v-for="item in purchaseMenus" :key="item.key">
                  <div class="menu-item" @click="goToPage(item.key)">
                    <el-icon><component :is="item.icon" /></el-icon>
                    <div class="menu-text">
                      <div class="menu-title">{{ item.title }}</div>
                      <div class="menu-desc">{{ item.desc }}</div>
                    </div>
                  </div>
                </el-col>
              </el-row>
            </div>
          </el-card>
        </el-col>

        <!-- 协议价管理 -->
        <el-col :xs="24" :sm="24" :md="8" :lg="8" :xl="8">
          <el-card class="feature-card" shadow="hover">
            <template #header>
              <div class="card-header">
                <span class="card-title">
                  <el-icon><Setting /></el-icon>
                  协议价管理
                </span>
              </div>
            </template>
            <div class="card-content">
              <el-row :gutter="16">
                <el-col :span="12" v-for="item in agreementMenus" :key="item.key">
                  <div class="menu-item" @click="goToPage(item.key)">
                    <el-icon><component :is="item.icon" /></el-icon>
                    <div class="menu-text">
                      <div class="menu-title">{{ item.title }}</div>
                      <div class="menu-desc">{{ item.desc }}</div>
                    </div>
                  </div>
                </el-col>
              </el-row>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>

    <!-- 保留原有内容 -->
    <div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-7 py-2 gap-4 md:gap-2 gva-container2 mt-6">
      <gva-card
        title="内容数据"
        custom-class="col-span-1 md:col-span-2 md:row-start-2 lg:col-span-6 col-start-1 row-span-2"
      >
        <gva-chart :type="4" />
      </gva-card>
      <gva-card
        title="文档"
        show-action
        custom-class="md:row-start-8 md:col-start-3 lg:row-start-3 lg:col-start-7"
      >
        <gva-wiki />
      </gva-card>
      <gva-card
        title="最新更新"
        custom-class="col-span-1 md:col-span-3 row-span-2"
      >
        <gva-table />
      </gva-card>
      <gva-card
        title="最新插件"
        custom-class="col-span-1 md:col-span-3 row-span-2"
      >
        <gva-plugin-table />
      </gva-card>
      <gva-card title="公告" show-action custom-class="col-span-1 lg:col-start-7">
        <gva-notice />
      </gva-card>
      <gva-card
        without-padding
        custom-class="overflow-hidden lg:h-40 col-span-1 md:col-start-2 md:col-span-1 lg:col-start-7"
      >
        <gva-banner />
      </gva-card>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import {
  GvaPluginTable,
  GvaTable,
  GvaChart,
  GvaWiki,
  GvaNotice,
  GvaQuickLink,
  GvaCard,
  GvaBanner
} from '@/view/dashboard/components'
import {
  Document,
  Shop,
  Check,
  Setting,
  List,
  Trophy,
  Money
} from '@element-plus/icons-vue'

defineOptions({
  name: 'Home'
})

// 询价管理菜单
const inquiryMenus = ref([
  {
    key: 'my-inquiry',
    title: '我的询价',
    desc: '销售录入',
    icon: Document
  },
  {
    key: 'pending-quote',
    title: '待报价项目',
    desc: '采购报价',
    icon: Shop
  },
  {
    key: 'inquiry-ledger',
    title: '询报价台账',
    desc: '全部记录',
    icon: List
  },
  {
    key: 'bid-projects',
    title: '中标项目',
    desc: '管理中标',
    icon: Trophy
  }
])

// 采购审批菜单
const purchaseMenus = ref([
  {
    key: 'payment-approval',
    title: '支付申请',
    desc: '吧盛支付',
    icon: Money
  },
  {
    key: 'order-approval',
    title: '订单审批',
    desc: '采购订单',
    icon: Check
  }
])

// 协议价管理菜单
const agreementMenus = ref([
  {
    key: 'agreement-price',
    title: '协议价管理',
    desc: '管理员',
    icon: Setting
  }
])

// 跳转到页面
const goToPage = (menu) => {
  window.open(`/#/inquiry-demo?menu=${menu}`, '_blank')
}
</script>

<style lang="scss" scoped>
.home-container {
  padding: 20px;
}

.feature-cards {
  .feature-card {
    height: 100%;

    .card-header {
      display: flex;
      align-items: center;

      .card-title {
        font-size: 18px;
        font-weight: 600;
        display: flex;
        align-items: center;
        gap: 8px;

        .el-icon {
          font-size: 20px;
          color: #409eff;
        }
      }
    }

    .card-content {
      padding: 10px 0;

      .menu-item {
        display: flex;
        align-items: center;
        padding: 15px;
        margin-bottom: 10px;
        border: 1px solid #e6e6e6;
        border-radius: 8px;
        cursor: pointer;
        transition: all 0.3s;

        &:hover {
          border-color: #409eff;
          background-color: #f0f9ff;
          transform: translateY(-2px);
        }

        .el-icon {
          font-size: 24px;
          color: #409eff;
          margin-right: 12px;
        }

        .menu-text {
          .menu-title {
            font-size: 14px;
            font-weight: 500;
            color: #303133;
            margin-bottom: 4px;
          }

          .menu-desc {
            font-size: 12px;
            color: #909399;
          }
        }
      }
    }
  }
}

</style>