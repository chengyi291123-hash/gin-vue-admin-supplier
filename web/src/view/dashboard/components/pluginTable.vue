<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->

<template>
  <div>
    <el-table :data="tableData" stripe style="width: 100%">
      <el-table-column prop="ranking" label="排名" width="80" align="center" />
      <el-table-column prop="title" label="插件标题" show-overflow-tooltip>
        <template #default="{ row }">
          <a class="text-active" :href="row.link" target="_blank">{{
            row.title
          }}</a>
        </template>
      </el-table-column>
      <el-table-column prop="click_num" label="关注度" width="100" />
      <el-table-column prop="hot" label="热度值" width="100" />
    </el-table>
  </div>
</template>

<script setup>
  const tableData = [
    {
      ranking: 1,
      title: '用户管理：管理系统用户，分配角色权限。',
      click_num: 523,
      hot: 263,
      link: '#'
    },
    {
      ranking: 2,
      title: '角色管理：配置系统角色，管理角色权限。',
      click_num: 416,
      hot: 223,
      link: '#'
    },
    {
      ranking: 3,
      title: '菜单管理：配置系统菜单，控制菜单显示。',
      click_num: 337,
      hot: 176,
      link: '#'
    },
    {
      ranking: 4,
      title: 'API管理：管理系统API接口，控制接口权限。',
      click_num: 292,
      hot: 145,
      link: '#'
    },
    {
      ranking: 5,
      title: '操作日志：记录系统操作日志，便于追溯。',
      click_num: 173,
      hot: 110,
      link: '#'
    }
  ]
</script>

<style scoped lang="scss"></style>
