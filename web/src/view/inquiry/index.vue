<template>
  <div class="inquiry-container">
    <!-- Sidebar -->
    <div class="aside">
      <div class="logo">
        <span>Gin-Vue-Admin 演示</span>
      </div>
      <el-menu
        :default-openeds="['1', '2', '3', '4']"
        :default-active="currentMenu"
        background-color="#2b3649"
        text-color="#bfcbd9"
        active-text-color="#409EFF"
        @select="handleSelect"
      >
        <el-sub-menu index="1">
          <template #title>
            <el-icon><Shop /></el-icon>
            <span>供应商管理</span>
          </template>
          <el-menu-item index="supplier-apply">供应商准入申请</el-menu-item>
          <el-menu-item index="supplier-list">供应商档案库</el-menu-item>
          <el-menu-item index="supplier-approval">供应商审批</el-menu-item>
          <el-menu-item index="supplier-change">供应商变更</el-menu-item>
        </el-sub-menu>

        <el-sub-menu index="2">
          <template #title>
            <el-icon><Document /></el-icon>
            <span>询价管理</span>
          </template>
          <el-menu-item index="my-inquiry">我的询价（销售）</el-menu-item>
          <el-menu-item index="pending-quote">待报价项目（采购）</el-menu-item>
          <el-menu-item index="inquiry-ledger">全部询报价台账</el-menu-item>
          <el-menu-item index="bid-projects">中标项目管理</el-menu-item>
        </el-sub-menu>

        <el-sub-menu index="3">
          <template #title>
            <el-icon><Check /></el-icon>
            <span>采购审批</span>
          </template>
          <el-menu-item index="payment-approval">吧盛支付申请</el-menu-item>
          <el-menu-item index="order-approval">采购订单审批</el-menu-item>
        </el-sub-menu>

        <el-sub-menu index="4" v-if="isAdmin">
          <template #title>
            <el-icon><Setting /></el-icon>
            <span>协议价管理</span>
          </template>
          <el-menu-item index="agreement-price">协议价管理</el-menu-item>
        </el-sub-menu>
      </el-menu>
    </div>

    <!-- Main Layout -->
    <div style="flex: 1; display: flex; flex-direction: column; overflow: hidden;">
      <!-- Header -->
      <div class="header">
        <div class="breadcrumb">
          首页 / {{ getCurrentCategory() }} / {{ getCurrentTitle() }}
        </div>
        <div style="display: flex; align-items: center;">
          <el-avatar :size="32" src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"></el-avatar>
        </div>
      </div>

      <!-- Content -->
      <div class="main">
        <div class="content-card">
          <!-- 我的询价（销售） -->
          <div v-if="currentMenu === 'my-inquiry'">
            <div class="form-section-title">询价需求录入</div>

            <el-form :model="inquiryForm" label-width="120px" style="max-width: 800px">
              <el-form-item label="项目名称">
                <el-input v-model="inquiryForm.project_name" placeholder="请输入项目名称"></el-input>
              </el-form-item>
              <el-form-item label="客户名称">
                <el-input v-model="inquiryForm.customer_name" placeholder="请输入客户名称"></el-input>
              </el-form-item>

              <el-form-item label="物料信息">
                <el-table :data="inquiryForm.items" border style="width: 100%">
                  <el-table-column label="物料编码" width="150">
                    <template #default="scope">
                      <el-input v-model="scope.row.material_code" placeholder="输入编码"></el-input>
                    </template>
                  </el-table-column>
                  <el-table-column label="物料名称" width="200">
                    <template #default="scope">
                      <el-input v-model="scope.row.material_name" placeholder="输入名称"></el-input>
                    </template>
                  </el-table-column>
                  <el-table-column label="规格型号" width="150">
                    <template #default="scope">
                      <el-input v-model="scope.row.specification" placeholder="输入规格"></el-input>
                    </template>
                  </el-table-column>
                  <el-table-column label="数量" width="120">
                    <template #default="scope">
                      <el-input-number v-model="scope.row.quantity" :min="1" size="small"></el-input-number>
                    </template>
                  </el-table-column>
                  <el-table-column label="单位" width="100">
                    <template #default="scope">
                      <el-input v-model="scope.row.unit" placeholder="单位"></el-input>
                    </template>
                  </el-table-column>
                  <el-table-column label="操作" width="100">
                    <template #default="scope">
                      <el-button type="danger" size="small" @click="removeItem(scope.$index)">删除</el-button>
                    </template>
                  </el-table-column>
                </el-table>
                <el-button type="primary" size="small" @click="addItem" style="margin-top: 10px;">+ 添加物料</el-button>
              </el-form-item>

              <el-form-item label="期望交期">
                <el-date-picker
                  v-model="inquiryForm.expected_delivery"
                  type="date"
                  placeholder="选择日期"
                  style="width: 100%"
                ></el-date-picker>
              </el-form-item>

              <el-form-item label="技术要求">
                <el-input
                  v-model="inquiryForm.technical_requirements"
                  type="textarea"
                  :rows="3"
                  placeholder="请输入技术要求"
                ></el-input>
              </el-form-item>

              <el-form-item label="技术附件">
                <el-upload
                  action="#"
                  :auto-upload="false"
                  :on-change="handleTechFile"
                  :file-list="techFiles"
                  multiple
                >
                  <el-button type="primary">上传附件</el-button>
                  <template #tip>
                    <div class="el-upload__tip">支持上传技术图纸、要求文档等</div>
                  </template>
                </el-upload>
              </el-form-item>

              <el-form-item label="录入人">
                <el-input v-model="currentUser" disabled></el-input>
              </el-form-item>

              <el-form-item label="录入时间">
                <el-input v-model="currentTime" disabled></el-input>
              </el-form-item>

              <el-form-item>
                <el-button type="primary" @click="submitInquiry">单条提交</el-button>
                <el-button type="success" @click="showImportDialog">批量导入</el-button>
              </el-form-item>
            </el-form>

            <!-- 批量导入对话框 -->
            <el-dialog v-model="importDialogVisible" title="批量导入询价" width="500px">
              <el-upload
                class="upload-demo"
                drag
                action="#"
                :auto-upload="false"
                :on-change="handleExcelFile"
                accept=".xlsx,.xls"
              >
                <el-icon class="el-icon--upload"><upload-filled /></el-icon>
                <div class="el-upload__text">
                  将Excel文件拖到此处，或<em>点击上传</em>
                </div>
                <template #tip>
                  <div class="el-upload__tip">
                    请按照模板格式上传Excel文件
                    <el-link type="primary" @click="downloadTemplate">下载模板</el-link>
                  </div>
                </template>
              </el-upload>
              <template #footer>
                <div class="dialog-footer">
                  <el-button @click="importDialogVisible = false">取消</el-button>
                  <el-button type="primary" @click="importExcel">确认导入</el-button>
                </div>
              </template>
            </el-dialog>
          </div>

          <!-- 待报价项目（采购） -->
          <div v-else-if="currentMenu === 'pending-quote'">
            <div class="form-section-title">待报价项目</div>
            <el-alert title="这是采购人员的主要工作台" type="info" show-icon style="margin-bottom: 20px;"></el-alert>
            <p>待报价项目列表功能开发中...</p>
          </div>

          <!-- 全部询报价台账 -->
          <div v-else-if="currentMenu === 'inquiry-ledger'">
            <div class="form-section-title">全部询报价台账</div>
            <p>询报价台账功能开发中...</p>
          </div>

          <!-- 中标项目管理 -->
          <div v-else-if="currentMenu === 'bid-projects'">
            <div class="form-section-title">中标项目管理</div>
            <p>中标项目管理功能开发中...</p>
          </div>

          <!-- 吧盛支付申请 -->
          <div v-else-if="currentMenu === 'payment-approval'">
            <div class="form-section-title">吧盛支付申请</div>
            <p>支付申请功能开发中...</p>
          </div>

          <!-- 采购订单审批 -->
          <div v-else-if="currentMenu === 'order-approval'">
            <div class="form-section-title">采购订单审批</div>
            <p>订单审批功能开发中...</p>
          </div>

          <!-- 协议价管理 -->
          <div v-else-if="currentMenu === 'agreement-price'">
            <div class="form-section-title">协议价管理</div>
            <p>协议价管理功能开发中...</p>
          </div>

          <!-- 供应商相关页面跳转 -->
          <div v-else>
            <div style="text-align: center; padding: 50px;">
              <p>正在跳转到供应商管理页面...</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { Shop, Document, Check, Setting, UploadFilled } from '@element-plus/icons-vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/pinia/modules/user'

const router = useRouter()
const userStore = useUserStore()
const currentMenu = ref('my-inquiry')
const isAdmin = ref(true)
const importDialogVisible = ref(false)
const techFiles = ref([])

// 当前用户和时间 - 从 userStore 获取真实用户信息
const currentUser = computed(() => userStore.userInfo?.nickName || userStore.userInfo?.userName || '未登录用户')
const currentTime = computed(() => new Date().toLocaleString())

// 询价表单
const inquiryForm = reactive({
  project_name: '',
  customer_name: '',
  expected_delivery: '',
  technical_requirements: '',
  items: [
    {
      material_code: '',
      material_name: '',
      specification: '',
      quantity: 1,
      unit: '个'
    }
  ]
})

// 获取当前菜单标题
const getCurrentTitle = () => {
  const map = {
    'my-inquiry': '我的询价',
    'pending-quote': '待报价项目',
    'inquiry-ledger': '全部询报价台账',
    'bid-projects': '中标项目管理',
    'payment-approval': '吧盛支付申请',
    'order-approval': '采购订单审批',
    'agreement-price': '协议价管理'
  }
  return map[currentMenu.value] || ''
}

// 获取当前分类
const getCurrentCategory = () => {
  const map = {
    'my-inquiry': '询价管理',
    'pending-quote': '询价管理',
    'inquiry-ledger': '询价管理',
    'bid-projects': '询价管理',
    'payment-approval': '采购审批',
    'order-approval': '采购审批',
    'agreement-price': '协议价管理'
  }
  return map[currentMenu.value] || ''
}

// 处理菜单选择
const handleSelect = (key) => {
  // 如果是供应商管理相关，跳转到原供应商页面
  if (['supplier-apply', 'supplier-list', 'supplier-approval', 'supplier-change'].includes(key)) {
    router.push({ name: 'SupplierDemo', query: { menu: key } })
    return
  }
  currentMenu.value = key
}

// 添加物料
const addItem = () => {
  inquiryForm.items.push({
    material_code: '',
    material_name: '',
    specification: '',
    quantity: 1,
    unit: '个'
  })
}

// 删除物料
const removeItem = (index) => {
  inquiryForm.items.splice(index, 1)
}

// 处理技术文件上传
const handleTechFile = (file) => {
  techFiles.value.push({
    name: file.name,
    raw: file
  })
}

// 显示导入对话框
const showImportDialog = () => {
  importDialogVisible.value = true
}

// 处理Excel文件
const handleExcelFile = (file) => {
  console.log('Excel file:', file)
}

// 下载模板
const downloadTemplate = () => {
  ElMessage.success('正在下载模板...')
}

// 导入Excel
const importExcel = () => {
  ElMessage.success('Excel导入成功')
  importDialogVisible.value = false
}

// 提交询价
const submitInquiry = () => {
  // 验证必填项
  if (!inquiryForm.project_name || !inquiryForm.customer_name) {
    ElMessage.error('请填写项目名称和客户名称')
    return
  }

  ElMessage.success('询价提交成功')
}

onMounted(() => {
  // 获取URL参数，如果有menu参数则跳转到相应页面
  const urlParams = new URLSearchParams(window.location.search)
  const menu = urlParams.get('menu')
  if (menu) {
    currentMenu.value = menu
  }
})
</script>

<style lang="scss" scoped>
.inquiry-container {
  height: 100vh;
  display: flex;
  background-color: #f5f7fa;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";
}

.aside {
  background-color: #001529;
  color: #fff;
  transition: width 0.3s;
  width: 220px;
  display: flex;
  flex-direction: column;
  box-shadow: 2px 0 6px rgba(0,21,41,.35);
}

.logo {
  height: 64px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  font-weight: 600;
  background-color: #002140;
  color: #fff;
}

.el-menu {
  border-right: none;
}

.header {
  background-color: #fff;
  border-bottom: 1px solid #e6e6e6;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 24px;
  height: 64px;
  box-shadow: 0 1px 4px rgba(0,21,41,.08);
}

.breadcrumb {
  font-size: 14px;
  color: #595959;
}

.main {
  background-color: #f0f2f5;
  padding: 24px;
  flex: 1;
  overflow-y: auto;
}

.content-card {
  background-color: #fff;
  padding: 32px;
  border-radius: 4px;
  min-height: calc(100vh - 150px);
  box-shadow: 0 2px 4px 0 rgba(0,0,0,0.05);
}

.form-section-title {
  font-size: 18px;
  font-weight: 500;
  color: #262626;
  margin-bottom: 24px;
  padding-bottom: 12px;
  border-bottom: 1px solid #f0f0f0;
  display: flex;
  align-items: center;
}

.form-section-title::before {
  content: '';
  display: block;
  width: 4px;
  height: 16px;
  background-color: #1890ff;
  margin-right: 8px;
  border-radius: 2px;
}
</style>