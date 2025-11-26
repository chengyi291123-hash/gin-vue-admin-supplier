<template>
  <el-descriptions :column="2" border>
    <el-descriptions-item label="供应商编号">{{ supplier.code }}</el-descriptions-item>
    <el-descriptions-item label="供应商名称">{{ supplier.name }}</el-descriptions-item>
    <el-descriptions-item label="供应商类型">
      <el-tag>{{ supplier.type }}</el-tag>
    </el-descriptions-item>
    <el-descriptions-item label="统一社会信用代码">{{ supplier.unified_social_credit_code }}</el-descriptions-item>
    <el-descriptions-item label="法定代表人">{{ supplier.legal_representative }}</el-descriptions-item>
    <el-descriptions-item label="注册资本">{{ supplier.registered_capital }}</el-descriptions-item>
    <el-descriptions-item label="成立日期">{{ supplier.established_date || '-' }}</el-descriptions-item>
    <el-descriptions-item label="企业类型">{{ supplier.enterprise_type || '-' }}</el-descriptions-item>
    <el-descriptions-item label="经营期限">{{ supplier.business_term || '-' }}</el-descriptions-item>
    <el-descriptions-item label="登记机关">{{ supplier.registration_authority || '-' }}</el-descriptions-item>
    <el-descriptions-item label="经营范围" :span="2">
      <div class="business-scope">{{ supplier.business_scope }}</div>
    </el-descriptions-item>
    <el-descriptions-item label="注册地址" :span="2">{{ supplier.registered_address }}</el-descriptions-item>
    <el-descriptions-item label="通讯地址" :span="2">{{ supplier.communication_address || '-' }}</el-descriptions-item>
    <el-descriptions-item label="供应商状态">
      <el-tag :type="getStatusType(supplier.status)">{{ supplier.status }}</el-tag>
    </el-descriptions-item>
    <el-descriptions-item label="创建时间">{{ supplier.created_time }}</el-descriptions-item>
    <el-descriptions-item label="备注" :span="2">{{ supplier.remark || '-' }}</el-descriptions-item>
  </el-descriptions>

  <div class="attachments" v-if="supplier.attachments && supplier.attachments.length > 0">
    <h4>附件信息</h4>
    <el-table :data="supplier.attachments" border style="width: 100%">
      <el-table-column prop="name" label="附件名称" />
      <el-table-column prop="type" label="附件类型" width="120" />
      <el-table-column label="操作" width="150">
        <template #default="{ row }">
          <el-button type="primary" size="small" @click="handlePreview(row)">预览</el-button>
          <el-button type="success" size="small" @click="handleDownload(row)">下载</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script setup>
import { defineProps } from 'vue'

const props = defineProps({
  supplier: {
    type: Object,
    required: true
  }
})

const getStatusType = (status) => {
  const statusMap = {
    '待审批': 'warning',
    '已通过': 'success',
    '已驳回': 'danger',
    '已注销': 'info'
  }
  return statusMap[status] || 'info'
}

const handlePreview = (attachment) => {
  // TODO: 实现附件预览功能
  console.log('预览附件:', attachment)
}

const handleDownload = (attachment) => {
  // TODO: 实现附件下载功能
  console.log('下载附件:', attachment)
}
</script>

<style lang="scss" scoped>
.business-scope {
  white-space: pre-wrap;
  line-height: 1.6;
}

.attachments {
  margin-top: 20px;

  h4 {
    margin-bottom: 10px;
    color: #606266;
  }
}
</style>