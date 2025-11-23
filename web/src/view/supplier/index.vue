<template>
  <div class="supplier-container">
    <!-- Sidebar -->
    <div class="aside">
      <div class="logo">
        <span>Gin-Vue-Admin 演示</span>
      </div>
      <el-menu
        :default-openeds="['1']"
        :default-active="currentMenu"
        background-color="#2b3649"
        text-color="#bfcbd9"
        active-text-color="#409EFF"
        @select="handleSelect"
      >
        <el-sub-menu index="1">
          <template #title>
            <el-icon><Menu /></el-icon>
            <span>供应商管理</span>
          </template>
          <el-menu-item index="apply">供应商申请</el-menu-item>
          <el-menu-item index="change">供应商变更</el-menu-item>
          <el-menu-item index="temp">临时供应商</el-menu-item>
          <el-menu-item index="qualified">合格供应商</el-menu-item>
          <el-menu-item index="blacklist">供应商黑名单</el-menu-item>
        </el-sub-menu>
      </el-menu>
    </div>

    <!-- Main Layout -->
    <div style="flex: 1; display: flex; flex-direction: column; overflow: hidden;">
      <!-- Header -->
      <div class="header">
        <div class="breadcrumb">
          首页 / 供应商管理 / {{ getCurrentTitle() }}
        </div>
        <div style="display: flex; align-items: center;">
          <el-button type="primary" size="small" @click="saveData" v-if="currentMenu === 'apply'" style="margin-right: 15px;">保存申请</el-button>
          <el-avatar :size="32" src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"></el-avatar>
        </div>
      </div>

      <!-- Content -->
      <div class="main">
        <div class="content-card">

          <!-- VIEW 1: 供应商申请 -->
          <div v-if="currentMenu === 'apply'">
            <el-alert title="提示：临时供方只需填写红色必填项，直接提交；申请合格供方需填写全部信息并审批。" type="info" show-icon style="margin-bottom: 20px;"></el-alert>

            <el-tabs v-model="activeTab">
              <el-tab-pane label="基本信息" name="basic">
                <div class="form-section-title">基本信息</div>
                <el-form :model="form" label-width="150px" style="max-width: 900px">
                  <el-form-item label="企业名称">
                    <el-input v-model="form.enterprise_name" placeholder="请输入企业全称"></el-input>
                  </el-form-item>
                  <el-form-item label="统一社会信用代码" v-if="!isTempAdd">
                    <el-input v-model="form.credit_code"></el-input>
                  </el-form-item>
                  <el-row>
                    <el-col :span="12">
                      <el-form-item label="供应商类型">
                        <template #label><span class="red-star">*</span>供应商类型</template>
                        <el-select v-model="form.entry_type" placeholder="请选择" style="width: 100%">
                          <el-option label="生产商" value="manufacturing"></el-option>
                          <el-option label="贸易商" value="trading"></el-option>
                        </el-select>
                      </el-form-item>
                    </el-col>
                    <el-col :span="12">
                      <el-form-item label="供应品类">
                        <template #label><span class="red-star">*</span>供应品类</template>
                        <el-select v-model="form.category" placeholder="请选择分类" style="width: 100%">
                          <el-option label="原材料" value="raw"></el-option>
                          <el-option label="电子元器件" value="electronic"></el-option>
                          <el-option label="办公用品" value="office"></el-option>
                        </el-select>
                      </el-form-item>
                    </el-col>
                  </el-row>
                  <el-row>
                    <el-col :span="12" v-if="!isTempAdd">
                      <el-form-item label="合作区域">
                        <el-input v-model="form.region"></el-input>
                      </el-form-item>
                    </el-col>
                    <el-col :span="12">
                      <el-form-item label="合作行业">
                        <template #label><span class="red-star">*</span>合作行业</template>
                        <el-select v-model="form.industry" placeholder="请选择" style="width: 100%">
                          <el-option label="核电" value="nuclear"></el-option>
                          <el-option label="石化" value="petrochemical"></el-option>
                          <el-option label="军工" value="military"></el-option>
                          <el-option label="其他" value="other"></el-option>
                        </el-select>
                      </el-form-item>
                    </el-col>
                  </el-row>
                  <el-form-item label="合作品牌">
                    <template #label><span class="red-star">*</span>合作品牌</template>
                    <el-input v-model="form.brand" placeholder="请输入品牌"></el-input>
                  </el-form-item>

                  <el-divider></el-divider>
                  <div class="form-section-title">联系人与采购</div>
                  <el-row>
                    <el-col :span="12">
                      <el-form-item label="联系人">
                        <template #label><span class="red-star">*</span>联系人</template>
                        <el-input v-model="form.contact_person"></el-input>
                      </el-form-item>
                    </el-col>
                    <el-col :span="12">
                      <el-form-item label="联系电话">
                        <template #label><span class="red-star">*</span>联系电话</template>
                        <el-input v-model="form.mobile"></el-input>
                      </el-form-item>
                    </el-col>
                  </el-row>
                  <el-form-item label="采购员" v-if="!isTempAdd">
                    <template #label><span class="red-star">*</span>采购员</template>
                    <el-input v-model="form.purchaser" placeholder="默认当前登录用户"></el-input>
                  </el-form-item>

                  <div v-if="!isTempAdd">
                    <el-divider></el-divider>
                    <div class="form-section-title">财务账户信息</div>
                    <el-form-item label="开户行">
                      <el-input v-model="form.bank_name"></el-input>
                    </el-form-item>
                    <el-form-item label="开户行名称 (支行)">
                      <el-input v-model="form.branch_name"></el-input>
                    </el-form-item>
                    <el-form-item label="银行账号">
                      <el-input v-model="form.bank_account"></el-input>
                    </el-form-item>
                    <el-form-item label="结算方式">
                      <el-select v-model="form.settlement" placeholder="请选择" style="width: 100%">
                        <el-option label="月结30天" value="m30"></el-option>
                        <el-option label="预付" value="prepay"></el-option>
                        <el-option label="货到付款" value="cod"></el-option>
                      </el-select>
                    </el-form-item>
                  </div>

                  <el-form-item>
                    <el-button>上一步</el-button>
                    <el-button type="primary" @click="nextTab('qual')">下一步：资质证书</el-button>
                  </el-form-item>
                </el-form>
              </el-tab-pane>

              <el-tab-pane label="资质证书" name="qual">
                <div class="form-section-title">资质文件上传</div>

                <el-form label-width="150px">
                  <el-form-item label="供应商类型">
                    <el-select v-model="form.entry_type" placeholder="请选择供应商类型" style="width: 100%">
                      <el-option label="生产商" value="manufacturing"></el-option>
                      <el-option label="贸易商" value="trading"></el-option>
                    </el-select>
                  </el-form-item>
                  <el-form-item label="选择证书类型">
                    <el-select v-model="currentCertType" placeholder="请选择要上传的证书类型" style="width: 100%;">
                      <el-option v-for="cert in certificateOptions" :key="cert" :label="cert" :value="cert"></el-option>
                    </el-select>
                  </el-form-item>

                  <el-form-item label="文件上传">
                    <el-upload class="upload-demo" drag action="#" :auto-upload="false" :on-change="handleFileChange">
                      <el-icon class="el-icon--upload"><upload-filled /></el-icon>
                      <div class="el-upload__text">将文件拖到此处，或 <em>点击上传</em></div>
                    </el-upload>
                  </el-form-item>

                  <!-- Uploaded List -->
                  <div style="margin-left: 150px; margin-bottom: 20px;">
                    <div v-for="(file, index) in uploadedFiles" :key="index" style="display: flex; align-items: center; margin-bottom: 5px; background: #f5f7fa; padding: 5px 10px; border-radius: 4px;">
                      <el-icon style="margin-right: 5px;"><document /></el-icon>
                      <span style="flex: 1;">{{ file.type }} - {{ file.name }}</span>
                      <el-icon style="cursor: pointer; color: #f56c6c;" @click="removeFile(index)"><delete /></el-icon>
                    </div>
                  </div>

                  <el-form-item>
                    <el-button @click="activeTab = 'basic'">上一步</el-button>
                    <el-button type="success" @click="saveData">提交申请</el-button>
                  </el-form-item>
                </el-form>
              </el-tab-pane>
            </el-tabs>
          </div>

          <!-- VIEW 2: 供应商变更 (NEW) -->
          <div v-else-if="currentMenu === 'change'">
            <div class="form-section-title">供应商信息变更</div>
            <el-alert title="采购员筛选供应商进行变更，变更需采购部长审批。临时供方可直接修改。" type="warning" style="margin-bottom: 20px;" :closable="false"></el-alert>

            <div style="margin-bottom: 20px;">
              <el-form :inline="true" :model="searchForm" class="demo-form-inline">
                <el-form-item label="企业名称">
                  <el-input v-model="searchForm.name" placeholder="请输入名称"></el-input>
                </el-form-item>
                <el-form-item label="状态">
                  <el-select v-model="searchForm.status" placeholder="全部" style="width: 150px;">
                    <el-option label="合格" value="qualified"></el-option>
                    <el-option label="临时" value="temp"></el-option>
                  </el-select>
                </el-form-item>
                <el-form-item>
                  <el-button type="primary" @click="onSearch">筛选</el-button>
                </el-form-item>
              </el-form>
            </div>

            <el-table :data="tableData" border style="width: 100%">
              <el-table-column prop="id" label="ID" width="60"></el-table-column>
              <el-table-column prop="name" label="企业名称" min-width="200"></el-table-column>
              <el-table-column prop="category" label="品类" width="120"></el-table-column>
              <el-table-column prop="status" label="当前状态" width="100">
                <template #default="scope">
                  <el-tag :type="getStatusType(scope.row.status)">{{ scope.row.status }}</el-tag>
                </template>
              </el-table-column>
              <el-table-column label="操作" width="150">
                <template #default="scope">
                  <el-button type="primary" size="small" plain icon="Edit" @click="openChangeDialog(scope.row)">变更信息</el-button>
                </template>
              </el-table-column>
            </el-table>

            <!-- 变更弹窗 -->
            <el-dialog v-model="changeDialogVisible" title="变更供应商信息" width="600px">
              <el-form :model="changeForm" label-width="120px">
                <el-form-item label="企业名称">
                  <el-input v-model="changeForm.name" disabled></el-input>
                </el-form-item>
                <el-form-item label="联系人">
                  <el-input v-model="changeForm.contact"></el-input>
                </el-form-item>
                <el-form-item label="联系电话">
                  <el-input v-model="changeForm.mobile"></el-input>
                </el-form-item>
                <el-form-item label="开户行">
                  <el-input v-model="changeForm.bank"></el-input>
                </el-form-item>
                <el-form-item label="变更原因">
                  <el-input type="textarea" v-model="changeForm.reason"></el-input>
                </el-form-item>
                <el-form-item label="转入黑名单">
                  <el-switch v-model="changeForm.toBlacklist" active-text="是" inactive-text="否"></el-switch>
                  <div style="font-size: 12px; color: #999;">* 若开启，将发起拉黑流程</div>
                </el-form-item>
              </el-form>
              <template #footer>
                <span class="dialog-footer">
                  <el-button @click="changeDialogVisible = false">取消</el-button>
                  <el-button type="primary" @click="submitChange">{{ changeForm.status === '临时' ? '确定' : '提交变更' }}</el-button>
                </span>
              </template>
            </el-dialog>
          </div>

          <!-- VIEW: 列表页面 (临时/合格/黑名单) -->
          <div v-else>
            <div style="margin-bottom: 20px; display: flex; justify-content: space-between;">
              <div>
                <el-input placeholder="搜索企业名称/联系人" style="width: 250px; margin-right: 10px;" prefix-icon="Search" v-model="searchForm.name"></el-input>
                <el-button type="primary" @click="onSearch">搜索</el-button>
              </div>
              <div>
                <el-button type="success" plain @click="exportExcel">导出 Excel</el-button>
                <el-button type="primary" v-if="currentMenu === 'temp'" @click="currentMenu = 'apply'; isTempAdd = true;">+ 新增</el-button>
                <el-button type="primary" v-if="currentMenu === 'qualified'" @click="currentMenu = 'apply'; isTempAdd = false;">+ 新增</el-button>
              </div>
            </div>

            <el-table :data="tableData" border style="width: 100%">
              <el-table-column prop="id" label="序号" width="60" fixed></el-table-column>
              <el-table-column prop="name" label="企业名称" min-width="200" fixed></el-table-column>
              <el-table-column prop="code" label="统一社会信用代码" width="180" v-if="currentMenu === 'qualified'"></el-table-column>
              <el-table-column prop="entry_type" label="供应商类型" width="110"></el-table-column>
              <el-table-column prop="category" label="供应品类" width="100"></el-table-column>
              <el-table-column prop="region" label="合作区域" width="100" v-if="currentMenu === 'qualified'"></el-table-column>
              <el-table-column prop="industry" label="合作行业" width="100"></el-table-column>
              <el-table-column prop="brand" label="合作品牌" width="100"></el-table-column>
              <el-table-column prop="contact" label="联系人" width="100"></el-table-column>
              <el-table-column prop="mobile" label="联系电话" width="120"></el-table-column>
              <el-table-column prop="bank_name" label="开户行" width="120" v-if="currentMenu === 'qualified'"></el-table-column>
              <el-table-column prop="branch_name" label="开户行名称" width="150" v-if="currentMenu === 'qualified'"></el-table-column>
              <el-table-column prop="settlement" label="结算方式" width="100" v-if="currentMenu === 'qualified'"></el-table-column>
              <el-table-column prop="purchaser" label="采购员" width="100" v-if="currentMenu === 'qualified'"></el-table-column>
              <el-table-column label="操作" width="100" fixed="right" v-if="currentMenu !== 'blacklist'">
                <template #default>
                  <el-button link type="primary" size="small">详情</el-button>
                </template>
              </el-table-column>
            </el-table>

            <div style="margin-top: 20px; text-align: right;">
              <el-pagination background layout="prev, pager, next" :total="tableData.length"></el-pagination>
            </div>
          </div>

        </div>
      </div>
    </div>

    <div class="demo-tips" v-if="showTips">
      <h4>功能更新说明</h4>
      <p>1. 侧边栏已更新为5项，包含<b>"供应商变更"</b>。</p>
      <p>2. "供应商申请"已更新字段（红色星号为必填）。</p>
      <p>3. "供应商变更"页支持筛选和信息修改/拉黑。</p>
      <el-button size="small" @click="showTips = false">知道了</el-button>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { Menu, UploadFilled, Document, Delete, Edit, Search } from '@element-plus/icons-vue'
import service from '@/utils/request' // Use the project's configured axios instance

const currentMenu = ref('apply')
const activeTab = ref('basic')
const showTips = ref(true)
const changeDialogVisible = ref(false)

// Certificate Logic
const currentCertType = ref('')
const uploadedFiles = ref([])

const form = reactive({
  enterprise_name: '杭州森森休闲用品有限公司',
  credit_code: '913301020743241989',
  entry_type: 'manufacturing',
  category: 'raw',
  region: '华东区',
  industry: 'petrochemical',
  brand: '森森',
  contact_person: '俞志勇',
  mobile: '19846768980',
  purchaser: 'Admin',
  email: '19846768980@126.com',
  settlement: 'm30',
  bank_name: '',
  branch_name: '',
  bank_account: ''
})

const searchForm = reactive({
  name: '',
  status: ''
})

const changeForm = reactive({
  id: '',
  name: '',
  contact: '',
  mobile: '',
  bank: '',
  reason: '',
  status: '',
  toBlacklist: false
})

const isTempAdd = ref(false)
const tableData = ref([])

const certificateOptions = computed(() => {
  if (form.entry_type === 'manufacturing') {
    return [
      '营业执照',
      '制造/设计许可证',
      'ISO 9001质量管理体系证书',
      '14001 职业健康安全管理体系证书',
      '18001 环境管理体系认证证书',
      '质量手册/质保大纲',
      '外部审查报告',
      '生产设备清单',
      '检测设备清单',
      '计量器具清单',
      '设计人员名单',
      '检验人员名单',
      '特种工艺人员名单及资质证书',
      '专职质保人员名单',
      '近三年公司业绩清单',
      '其他文件'
    ]
  } else {
    return [
      '营业执照',
      '品牌代理证书',
      'ISO 9001质量管理体系证书',
      '14001 职业健康安全管理体系证书',
      '18001 环境管理体系认证证书',
      '近三年公司业绩清单',
      '其他文件'
    ]
  }
})

const fetchData = async () => {
  let apiStatus = ''
  if (currentMenu.value === 'qualified') apiStatus = 'qualified'
  else if (currentMenu.value === 'temp') apiStatus = 'temp'
  else if (currentMenu.value === 'blacklist') apiStatus = 'blacklist'
  
  const params = {
    page: 1,
    pageSize: 100,
  }
  if (apiStatus) params.status = apiStatus
  if (searchForm.name) params.enterprise_name = searchForm.name
  if (searchForm.status) params.status = searchForm.status

  try {
    const res = await service({
      url: '/suppliers',
      method: 'get',
      params
    })
    if (res.code === 0) {
        tableData.value = res.data.list.map(item => ({
            ...item,
            id: item.ID,
            name: item.enterprise_name,
            code: item.credit_code,
            contact: item.contact_person
        }))
    }
  } catch (error) {
    console.error(error)
    ElMessage.error('获取数据失败')
  }
}

onMounted(() => {
  fetchData()
})

const handleSelect = (key) => {
  currentMenu.value = key
  if (key === 'apply') isTempAdd.value = false
  fetchData()
}

const nextTab = (tabName) => {
  activeTab.value = tabName
}

const saveData = async () => {
  try {
    const payload = { ...form }
    if (isTempAdd.value) payload.status = 'temp'
    else payload.status = 'qualified'
    
    const res = await service({
        url: '/suppliers',
        method: 'post',
        data: payload
    })
    
    if (res.code === 0) {
        ElMessage.success('提交申请成功！')
        fetchData()
    }
  } catch(e) {
    ElMessage.error('提交失败')
  }
}

const onSearch = () => {
  fetchData()
}

const exportExcel = () => {
  ElMessage.success('正在导出 Excel...')
}

const openChangeDialog = (row) => {
  changeForm.id = row.id
  changeForm.name = row.name
  changeForm.contact = row.contact
  changeForm.mobile = row.mobile
  changeForm.bank = row.bank_name
  changeForm.reason = ''
  changeForm.status = row.status
  changeForm.toBlacklist = false
  changeDialogVisible.value = true
}

const submitChange = async () => {
  const payload = {
    ID: changeForm.id,
    contact_person: changeForm.contact,
    mobile: changeForm.mobile,
    bank_name: changeForm.bank,
  }
  // For demo, we are using ID to update
  if (changeForm.toBlacklist) {
    payload.status = 'blacklist'
  }
  
  try {
    const res = await service({
        url: `/suppliers/${changeForm.id}`,
        method: 'put',
        data: payload
    })

    if (res.code === 0) {
        if (changeForm.toBlacklist) {
            ElMessage.warning('已转入黑名单')
        } else {
            ElMessage.success('变更成功')
        }
        changeDialogVisible.value = false
        fetchData()
    }
  } catch(e) {
    ElMessage.error('变更失败')
  }
}

const getCurrentTitle = () => {
  const map = {
    'apply': '供应商申请',
    'change': '供应商变更',
    'temp': '临时供应商',
    'qualified': '合格供应商',
    'blacklist': '供应商黑名单'
  }
  return map[currentMenu.value] || ''
}

const getStatusType = (status) => {
  if (status === 'qualified') return 'success'
  if (status === 'temp') return 'warning'
  if (status === 'blacklist') return 'danger'
  return 'info'
}

const handleFileChange = (uploadFile) => {
  if (!currentCertType.value) {
    ElMessage.warning('请先选择证书类型！')
    return
  }
  uploadedFiles.value.push({
    type: currentCertType.value,
    name: uploadFile.name
  })
  ElMessage.success(`已添加 ${currentCertType.value}`)
}

const removeFile = (index) => {
  uploadedFiles.value.splice(index, 1)
}
</script>

<style scoped>
.supplier-container {
  height: 100vh;
  display: flex;
  background-color: #f0f2f5;
}

.aside {
  background-color: #2b3649;
  color: #fff;
  transition: width 0.3s;
  width: 220px;
  display: flex;
  flex-direction: column;
}
.logo {
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  font-weight: bold;
  background-color: #2b3649;
  border-bottom: 1px solid #454d5e;
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
  padding: 0 20px;
  height: 60px;
  box-shadow: 0 1px 4px rgba(0,21,41,.08);
}
.breadcrumb {
  font-size: 14px;
  color: #97a8be;
}

.main {
  background-color: #f0f2f5;
  padding: 20px;
  flex: 1;
  overflow-y: auto;
}
.content-card {
  background-color: #fff;
  padding: 20px;
  border-radius: 4px;
  min-height: calc(100vh - 140px);
}

.form-section-title {
  font-size: 16px;
  font-weight: bold;
  margin-bottom: 20px;
  padding-left: 10px;
  border-left: 4px solid #409EFF;
}
.demo-tips {
  position: fixed;
  bottom: 20px;
  right: 20px;
  background: rgba(0,0,0,0.7);
  color: white;
  padding: 15px;
  border-radius: 8px;
  max-width: 300px;
  z-index: 9999;
  font-size: 14px;
}

.red-star {
  color: #f56c6c;
  margin-right: 4px;
}
</style>
