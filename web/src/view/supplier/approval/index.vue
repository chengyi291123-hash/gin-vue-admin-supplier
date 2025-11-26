<template>
  <div class="supplier-approval">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>供应商审批列表</span>
          <el-button type="primary" @click="handleRefresh">刷新</el-button>
        </div>
      </template>

      <el-table v-loading="loading" :data="supplierList" style="width: 100%">
        <el-table-column type="expand">
          <template #default="{ row }">
            <div class="supplier-detail">
              <el-tabs v-model="activeTab" type="card">
                <el-tab-pane label="供应商基本信息" name="basic">
                  <supplier-basic-info :supplier="row" />
                </el-tab-pane>
                <el-tab-pane label="联系人信息" name="contact">
                  <supplier-contact-info :contacts="row.contacts" />
                </el-tab-pane>
                <el-tab-pane label="银行账户信息" name="bank">
                  <supplier-bank-info :banks="row.banks" />
                </el-tab-pane>
              </el-tabs>
            </div>
          </template>
        </el-table-column>

        <el-table-column prop="code" label="供应商编号" width="120" />
        <el-table-column prop="name" label="供应商名称" width="200" />
        <el-table-column prop="type" label="供应商类型" width="120">
          <template #default="{ row }">
            <el-tag>{{ row.type }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === '待审批' ? 'warning' : row.status === '已通过' ? 'success' : 'danger'">
              {{ row.status }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_time" label="创建时间" width="180" />
        <el-table-column label="操作" fixed="right" width="200">
          <template #default="{ row }">
            <el-button
              v-if="row.status === '待审批'"
              type="success"
              size="small"
              @click="handleApprove(row)"
            >
              通过
            </el-button>
            <el-button
              v-if="row.status === '待审批'"
              type="danger"
              size="small"
              @click="handleReject(row)"
            >
              驳回
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="pagination-container">
        <el-pagination
          v-model:current-page="currentPage"
          v-model:page-size="pageSize"
          :page-sizes="[10, 20, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>

    <!-- 审批对话框 -->
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="500px">
      <el-form ref="formRef" :model="form" label-width="80px">
        <el-form-item label="审批意见" prop="remark">
          <el-input
            v-model="form.remark"
            type="textarea"
            :rows="4"
            placeholder="请输入审批意见"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="submitApproval">确认</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import SupplierBasicInfo from './components/SupplierBasicInfo.vue'
import SupplierContactInfo from './components/SupplierContactInfo.vue'
import SupplierBankInfo from './components/SupplierBankInfo.vue'

const loading = ref(false)
const supplierList = ref([])
const activeTab = ref('basic')
const currentPage = ref(1)
const pageSize = ref(10)
const total = ref(0)
const dialogVisible = ref(false)
const dialogTitle = ref('')
const currentSupplier = ref(null)

const form = reactive({
  remark: ''
})

const getSupplierList = async () => {
  loading.value = true
  try {
    // TODO: 调用API获取供应商审批列表
    const mockData = {
      data: [
        {
          id: 1,
          code: 'GYS20231126001',
          name: '上海测试供应商1',
          type: '原材料供应商',
          status: '待审批',
          created_time: '2023-11-26 10:00:00',
          unified_social_credit_code: '91310000123456789A',
          legal_representative: '张三',
          registered_capital: '1000万元',
          business_scope: '建筑材料、装饰材料、金属材料、五金交电的销售',
          registered_address: '上海市浦东新区张江高科技园区',
          contacts: [
            {
              name: '李四',
              department: '采购部',
              position: '采购经理',
              mobile: '13800138000',
              email: 'lisi@example.com',
              is_default: true
            }
          ],
          banks: [
            {
              account_name: '上海测试供应商1',
              account_number: '1234567890123456789',
              bank_name: '中国银行股份有限公司上海浦东分行',
              bank_code: '104290000034',
              is_default: true
            }
          ]
        },
        {
          id: 2,
          code: 'GYS20231126002',
          name: '上海测试供应商2',
          type: '设备供应商',
          status: '待审批',
          created_time: '2023-11-26 11:00:00',
          unified_social_credit_code: '91310000987654321B',
          legal_representative: '王五',
          registered_capital: '500万元',
          business_scope: '机械设备、电气设备、自动化设备的销售及维修',
          registered_address: '上海市闵行区莘庄工业园区',
          contacts: [
            {
              name: '赵六',
              department: '销售部',
              position: '销售经理',
              mobile: '13900139000',
              email: 'zhaoliu@example.com',
              is_default: true
            }
          ],
          banks: [
            {
              account_name: '上海测试供应商2',
              account_number: '9876543210987654321',
              bank_name: '中国工商银行股份有限公司上海闵行支行',
              bank_code: '102290000045',
              is_default: true
            }
          ]
        }
      ],
      total: 2,
      page: 1,
      pageSize: 10
    }

    supplierList.value = mockData.data
    total.value = mockData.total
  } catch (error) {
    ElMessage.error('获取供应商列表失败')
  } finally {
    loading.value = false
  }
}

const handleApprove = (row) => {
  currentSupplier.value = row
  dialogTitle.value = '审批通过'
  form.remark = ''
  dialogVisible.value = true
}

const handleReject = (row) => {
  currentSupplier.value = row
  dialogTitle.value = '审批驳回'
  form.remark = ''
  dialogVisible.value = true
}

const submitApproval = async () => {
  if (!form.remark.trim()) {
    ElMessage.warning('请输入审批意见')
    return
  }

  try {
    const isApproved = dialogTitle.value === '审批通过'
    // TODO: 调用API提交审批结果
    ElMessage.success(isApproved ? '审批通过成功' : '审批驳回成功')
    dialogVisible.value = false
    getSupplierList()
  } catch (error) {
    ElMessage.error('操作失败')
  }
}

const handleRefresh = () => {
  getSupplierList()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  getSupplierList()
}

const handleCurrentChange = (val) => {
  currentPage.value = val
  getSupplierList()
}

onMounted(() => {
  getSupplierList()
})
</script>

<style lang="scss" scoped>
.supplier-approval {
  padding: 20px;

  .card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .supplier-detail {
    padding: 20px;
  }

  .pagination-container {
    margin-top: 20px;
    display: flex;
    justify-content: flex-end;
  }
}
</style>