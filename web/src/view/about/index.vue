<template>
  <div class="mt-2">
    <div class="flex flex-col md:flex-row gap-4">
      <div class="w-full md:w-1/2">
        <el-card class="min-w-96">
          <template #header>
            <el-divider>叭盛科技</el-divider>
          </template>
          <div>
            <div class="w-full flex items-center justify-center">
              <img
                class="org-img dom-center"
                src="@/assets/logo.png"
                alt="叭盛科技"
              />
            </div>
            <div class="w-full flex items-center justify-center mt-4">
              <p class="text-gray-600">叭盛科技管理平台</p>
            </div>
          </div>
        </el-card>
      </div>
      <div class="w-full md:w-1/2">
        <el-card>
          <template #header>
            <div>系统信息</div>
          </template>
          <div class="p-4">
            <p class="mb-2">版本：v2.8.6</p>
            <p class="mb-2">技术栈：Vue 3 + Golang</p>
            <p class="mb-2">版权所有：叭盛科技</p>
          </div>
        </el-card>
      </div>
    </div>
  </div>
</template>

<script setup>
  defineOptions({
    name: 'About'
  })
</script>

<style scoped>
  .org-img {
    height: 150px;
    width: 150px;
  }

  .dom-center {
    margin-left: 50%;
    transform: translateX(-50%);
  }
</style>
