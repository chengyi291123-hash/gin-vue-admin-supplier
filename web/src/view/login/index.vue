<template>
  <div class="login-container">
    <!-- 左侧品牌区域 -->
    <div class="brand-section">
      <div class="brand-content">
        <div class="logo-wrapper">
          <div class="logo-icon">
            <span class="logo-text">叭盛</span>
          </div>
        </div>
        <h1 class="brand-title">叭盛科技</h1>
        <p class="brand-subtitle">BASHENG TECHNOLOGY</p>
        <div class="brand-divider"></div>
        <p class="brand-slogan">智能 · 高效 · 安全</p>

        <div class="features">
          <div class="feature-item">
            <div class="feature-icon">
              <el-icon><Monitor /></el-icon>
            </div>
            <span>系统管理</span>
          </div>
          <div class="feature-item">
            <div class="feature-icon">
              <el-icon><DataAnalysis /></el-icon>
            </div>
            <span>数据分析</span>
          </div>
          <div class="feature-item">
            <div class="feature-icon">
              <el-icon><Lock /></el-icon>
            </div>
            <span>安全可靠</span>
          </div>
        </div>
      </div>
    </div>

    <!-- 右侧登录区域 -->
    <div class="login-section">
      <div class="login-card">
        <div class="login-header">
          <h2>欢迎登录</h2>
          <p>Welcome back</p>
        </div>

        <el-form
          ref="loginForm"
          :model="loginFormData"
          :rules="rules"
          :validate-on-rule-change="false"
          class="login-form"
          @keyup.enter="submitForm"
        >
          <el-form-item prop="username">
            <el-input
              v-model="loginFormData.username"
              placeholder="请输入用户名"
              :prefix-icon="User"
              size="large"
            />
          </el-form-item>

          <el-form-item prop="password">
            <el-input
              v-model="loginFormData.password"
              type="password"
              placeholder="请输入密码"
              :prefix-icon="Lock"
              show-password
              size="large"
            />
          </el-form-item>

          <el-form-item v-if="loginFormData.openCaptcha" prop="captcha">
            <div class="captcha-wrapper">
              <el-input
                v-model="loginFormData.captcha"
                placeholder="请输入验证码"
                :prefix-icon="Key"
                size="large"
              />
              <div class="captcha-img" @click="loginVerify()">
                <img v-if="picPath" :src="picPath" alt="验证码" />
              </div>
            </div>
          </el-form-item>

          <el-form-item>
            <el-button
              type="primary"
              size="large"
              class="login-btn"
              @click="submitForm"
            >
              登 录
            </el-button>
          </el-form-item>

          <div class="login-actions">
            <el-button text size="small" @click="checkInit">
              系统初始化
            </el-button>
            <el-button text size="small" @click="router.push({ name: 'SupplierDemo' })">
              查看演示
            </el-button>
          </div>
        </el-form>

        <div class="login-footer">
          <span>© 2024 叭盛科技 · 保留所有权利</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { captcha } from '@/api/user'
import { checkDB } from '@/api/initdb'
import { reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/pinia/modules/user'
import { User, Lock, Key, Monitor, DataAnalysis } from '@element-plus/icons-vue'

defineOptions({
  name: 'Login'
})

const router = useRouter()
const captchaRequiredLength = ref(6)

const checkUsername = (rule, value, callback) => {
  if (value.length < 5) {
    return callback(new Error('请输入正确的用户名'))
  }
  callback()
}

const checkPassword = (rule, value, callback) => {
  if (value.length < 6) {
    return callback(new Error('请输入正确的密码'))
  }
  callback()
}

const checkCaptcha = (rule, value, callback) => {
  if (!loginFormData.openCaptcha) {
    return callback()
  }
  const sanitizedValue = (value || '').replace(/\s+/g, '')
  if (!sanitizedValue) {
    return callback(new Error('请输入验证码'))
  }
  if (!/^\d+$/.test(sanitizedValue)) {
    return callback(new Error('验证码须为数字'))
  }
  if (sanitizedValue.length < captchaRequiredLength.value) {
    return callback(new Error(`请输入至少${captchaRequiredLength.value}位数字验证码`))
  }
  if (sanitizedValue !== value) {
    loginFormData.captcha = sanitizedValue
  }
  callback()
}

const loginVerify = async () => {
  const ele = await captcha()
  captchaRequiredLength.value = Number(ele.data?.captchaLength) || 0
  picPath.value = ele.data?.picPath
  loginFormData.captchaId = ele.data?.captchaId
  loginFormData.openCaptcha = ele.data?.openCaptcha
}
loginVerify()

const loginForm = ref(null)
const picPath = ref('')
const loginFormData = reactive({
  username: 'admin',
  password: '',
  captcha: '',
  captchaId: '',
  openCaptcha: false
})

const rules = reactive({
  username: [{ validator: checkUsername, trigger: 'blur' }],
  password: [{ validator: checkPassword, trigger: 'blur' }],
  captcha: [{ validator: checkCaptcha, trigger: 'blur' }]
})

const userStore = useUserStore()

const login = async () => {
  return await userStore.LoginIn(loginFormData)
}

const submitForm = () => {
  loginForm.value.validate(async (v) => {
    if (!v) {
      ElMessage({
        type: 'error',
        message: '请正确填写登录信息',
        showClose: true
      })
      return false
    }
    const flag = await login()
    if (!flag) {
      await loginVerify()
      return false
    }
    return true
  })
}

const checkInit = async () => {
  const res = await checkDB()
  if (res.code === 0) {
    if (res.data?.needInit) {
      userStore.NeedInit()
      await router.push({ name: 'Init' })
    } else {
      ElMessage({
        type: 'info',
        message: '已配置数据库信息，无法初始化'
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.login-container {
  display: flex;
  width: 100vw;
  height: 100vh;
  background: #f8f9fa;
}

// 左侧品牌区域
.brand-section {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  position: relative;
  overflow: hidden;

  &::before {
    content: '';
    position: absolute;
    top: -50%;
    left: -50%;
    width: 200%;
    height: 200%;
    background: radial-gradient(circle, rgba(255,255,255,0.1) 0%, transparent 60%);
    animation: pulse 15s ease-in-out infinite;
  }

  &::after {
    content: '';
    position: absolute;
    bottom: -30%;
    right: -30%;
    width: 80%;
    height: 80%;
    background: radial-gradient(circle, rgba(255,255,255,0.08) 0%, transparent 50%);
  }
}

@keyframes pulse {
  0%, 100% { transform: scale(1); opacity: 0.5; }
  50% { transform: scale(1.1); opacity: 0.8; }
}

.brand-content {
  position: relative;
  z-index: 1;
  text-align: center;
  color: #fff;
  padding: 48px;
}

.logo-wrapper {
  margin-bottom: 32px;
}

.logo-icon {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 96px;
  height: 96px;
  background: rgba(255, 255, 255, 0.2);
  backdrop-filter: blur(10px);
  border-radius: 24px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  transition: transform 0.3s ease, box-shadow 0.3s ease;

  &:hover {
    transform: translateY(-4px);
    box-shadow: 0 12px 40px rgba(0, 0, 0, 0.15);
  }
}

.logo-text {
  font-size: 36px;
  font-weight: 700;
  color: #fff;
  letter-spacing: 2px;
}

.brand-title {
  font-size: 32px;
  font-weight: 600;
  margin: 0 0 8px 0;
  letter-spacing: 4px;
}

.brand-subtitle {
  font-size: 14px;
  opacity: 0.8;
  letter-spacing: 3px;
  margin: 0;
}

.brand-divider {
  width: 48px;
  height: 2px;
  background: rgba(255, 255, 255, 0.4);
  margin: 32px auto;
  border-radius: 1px;
}

.brand-slogan {
  font-size: 16px;
  opacity: 0.9;
  margin: 0 0 48px 0;
  letter-spacing: 8px;
}

.features {
  display: flex;
  gap: 32px;
  justify-content: center;
}

.feature-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 12px;
  opacity: 0.9;
  transition: opacity 0.3s ease, transform 0.3s ease;

  &:hover {
    opacity: 1;
    transform: translateY(-2px);
  }

  span {
    font-size: 12px;
    letter-spacing: 1px;
  }
}

.feature-icon {
  width: 48px;
  height: 48px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(255, 255, 255, 0.15);
  border-radius: 12px;
  font-size: 20px;
  backdrop-filter: blur(5px);
}

// 右侧登录区域
.login-section {
  width: 480px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #fff;
  box-shadow: -20px 0 60px rgba(0, 0, 0, 0.05);
}

.login-card {
  width: 100%;
  max-width: 360px;
  padding: 48px;
}

.login-header {
  text-align: center;
  margin-bottom: 40px;

  h2 {
    font-size: 24px;
    font-weight: 600;
    color: #1a1a2e;
    margin: 0 0 8px 0;
  }

  p {
    font-size: 14px;
    color: #999;
    margin: 0;
  }
}

.login-form {
  :deep(.el-form-item) {
    margin-bottom: 24px;
  }

  :deep(.el-input__wrapper) {
    border-radius: 8px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
    border: 1px solid #eee;
    transition: all 0.3s ease;

    &:hover {
      border-color: #667eea;
    }

    &.is-focus {
      border-color: #667eea;
      box-shadow: 0 4px 12px rgba(102, 126, 234, 0.15);
    }
  }

  :deep(.el-input__inner) {
    height: 48px;
    font-size: 14px;
  }
}

.captcha-wrapper {
  display: flex;
  gap: 16px;
  width: 100%;

  .el-input {
    flex: 1;
  }

  .captcha-img {
    width: 120px;
    height: 48px;
    border-radius: 8px;
    overflow: hidden;
    cursor: pointer;
    background: #f5f5f5;
    transition: opacity 0.3s ease;

    &:hover {
      opacity: 0.8;
    }

    img {
      width: 100%;
      height: 100%;
      object-fit: cover;
    }
  }
}

.login-btn {
  width: 100%;
  height: 48px;
  font-size: 16px;
  font-weight: 500;
  border-radius: 8px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border: none;
  box-shadow: 0 4px 16px rgba(102, 126, 234, 0.3);
  transition: all 0.3s ease;

  &:hover {
    transform: translateY(-2px);
    box-shadow: 0 6px 20px rgba(102, 126, 234, 0.4);
  }

  &:active {
    transform: translateY(0);
  }
}

.login-actions {
  display: flex;
  justify-content: center;
  gap: 24px;
  margin-top: -8px;

  :deep(.el-button) {
    color: #999;
    font-size: 13px;

    &:hover {
      color: #667eea;
    }
  }
}

.login-footer {
  text-align: center;
  margin-top: 48px;
  padding-top: 24px;
  border-top: 1px solid #f0f0f0;

  span {
    font-size: 12px;
    color: #bbb;
  }
}

// 响应式
@media (max-width: 968px) {
  .login-container {
    flex-direction: column;
  }

  .brand-section {
    flex: none;
    padding: 48px 24px;
  }

  .brand-content {
    padding: 24px;
  }

  .logo-icon {
    width: 72px;
    height: 72px;
    border-radius: 18px;
  }

  .logo-text {
    font-size: 28px;
  }

  .brand-title {
    font-size: 24px;
  }

  .features {
    display: none;
  }

  .brand-divider,
  .brand-slogan {
    display: none;
  }

  .login-section {
    width: 100%;
    flex: 1;
  }

  .login-card {
    padding: 32px 24px;
  }
}
</style>
