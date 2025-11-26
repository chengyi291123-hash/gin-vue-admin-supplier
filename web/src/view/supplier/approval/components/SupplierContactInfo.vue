<template>
  <el-table :data="contacts" border style="width: 100%">
    <el-table-column prop="name" label="姓名" width="120" />
    <el-table-column prop="department" label="部门" width="120" />
    <el-table-column prop="position" label="职位" width="120" />
    <el-table-column prop="mobile" label="手机号码" width="150" />
    <el-table-column prop="telephone" label="固定电话" width="150" />
    <el-table-column prop="email" label="邮箱地址" width="200" />
    <el-table-column prop="wechat" label="微信号" width="150" />
    <el-table-column prop="fax" label="传真" width="150" />
    <el-table-column label="默认联系人" width="120" align="center">
      <template #default="{ row }">
        <el-tag v-if="row.is_default" type="success" size="small">是</el-tag>
        <span v-else>-</span>
      </template>
    </el-table-column>
    <el-table-column label="状态" width="100" align="center">
      <template #default="{ row }">
        <el-tag :type="row.status === '启用' ? 'success' : 'info'" size="small">
          {{ row.status || '启用' }}
        </el-tag>
      </template>
    </el-table-column>
    <el-table-column prop="address" label="联系地址" min-width="200" />
    <el-table-column prop="remark" label="备注" width="150" />
  </el-table>

  <div v-if="!contacts || contacts.length === 0" class="empty-state">
    <el-empty description="暂无联系人信息" />
  </div>
</template>

<script setup>
import { defineProps } from 'vue'

const props = defineProps({
  contacts: {
    type: Array,
    default: () => []
  }
})
</script>

<style lang="scss" scoped>
.empty-state {
  padding: 40px 0;
  text-align: center;
}
</style>