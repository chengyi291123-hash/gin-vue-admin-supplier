<template>
  <el-table :data="banks" border style="width: 100%">
    <el-table-column label="默认账户" width="100" align="center">
      <template #default="{ row }">
        <el-tag v-if="row.is_default" type="success" size="small">是</el-tag>
        <span v-else>-</span>
      </template>
    </el-table-column>
    <el-table-column prop="account_name" label="账户名称" width="200" />
    <el-table-column prop="account_number" label="银行账号" width="200" />
    <el-table-column prop="bank_name" label="开户银行" width="250" />
    <el-table-column prop="bank_code" label="支付系统行号" width="180" />
    <el-table-column prop="branch_name" label="开户支行" width="200" />
    <el-table-column prop="account_type" label="账户类型" width="120">
      <template #default="{ row }">
        <el-tag>{{ getAccountTypeLabel(row.account_type) }}</el-tag>
      </template>
    </el-table-column>
    <el-table-column prop="currency" label="币种" width="100">
      <template #default="{ row }">
        <el-tag>{{ row.currency || 'CNY' }}</el-tag>
      </template>
    </el-table-column>
    <el-table-column label="状态" width="100" align="center">
      <template #default="{ row }">
        <el-tag :type="row.status === '启用' ? 'success' : 'info'" size="small">
          {{ row.status || '启用' }}
        </el-tag>
      </template>
    </el-table-column>
    <el-table-column prop="remark" label="备注" min-width="150" />
  </el-table>

  <div v-if="!banks || banks.length === 0" class="empty-state">
    <el-empty description="暂无银行账户信息" />
  </div>
</template>

<script setup>
import { defineProps } from 'vue'

const props = defineProps({
  banks: {
    type: Array,
    default: () => []
  }
})

const getAccountTypeLabel = (type) => {
  const typeMap = {
    'basic': '基本存款账户',
    'general': '一般存款账户',
    'special': '专用存款账户',
    'temporary': '临时存款账户'
  }
  return typeMap[type] || type || '一般存款账户'
}
</script>

<style lang="scss" scoped>
.empty-state {
  padding: 40px 0;
  text-align: center;
}
</style>