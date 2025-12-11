<template>
  <div class="login-container">
    <!-- 动态背景 -->
    <div class="bg-layer">
      <div class="grid-lines"></div>
      <div class="floating-shapes">
        <div class="shape shape-1"></div>
        <div class="shape shape-2"></div>
        <div class="shape shape-3"></div>
        <div class="shape shape-4"></div>
        <div class="shape shape-5"></div>
      </div>
      <div class="glow-orb glow-1"></div>
      <div class="glow-orb glow-2"></div>
      <div class="glow-orb glow-3"></div>
    </div>

    <!-- 主内容区 -->
    <div class="main-content">
      <!-- 左侧品牌区 -->
      <div class="brand-area">
        <div class="brand-wrapper">
          <!-- Logo -->
          <div class="logo-container">
            <div class="logo-ring">
              <div class="ring-inner"></div>
            </div>
            <div class="logo-core">
              <span class="logo-char">B</span>
              <span class="logo-char">S</span>
            </div>
          </div>

          <!-- 品牌名称 -->
          <div class="brand-text">
            <h1 class="brand-name">
              <span class="char" style="--i: 0">叭</span>
              <span class="char" style="--i: 1">盛</span>
              <span class="char" style="--i: 2">科</span>
              <span class="char" style="--i: 3">技</span>
            </h1>
            <div class="brand-en">BASHENG TECHNOLOGY</div>
          </div>

          <!-- 装饰线 -->
          <div class="deco-line">
            <div class="line-pulse"></div>
          </div>

          <!-- Slogan -->
          <p class="slogan">
            <span class="slogan-word">智能驱动</span>
            <span class="slogan-dot"></span>
            <span class="slogan-word">数据赋能</span>
            <span class="slogan-dot"></span>
            <span class="slogan-word">安全护航</span>
          </p>

          <!-- 特性卡片 -->
          <div class="feature-cards">
            <div class="feature-card" style="--delay: 0">
              <div class="card-icon">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
                  <path d="M9 3H5a2 2 0 0 0-2 2v4m6-6h10a2 2 0 0 1 2 2v4M9 3v18m0 0h10a2 2 0 0 0 2-2v-4M9 21H5a2 2 0 0 1-2-2v-4m0-6v6m18-6v6"/>
                </svg>
              </div>
              <div class="card-content">
                <span class="card-title">智慧管理</span>
                <span class="card-desc">全流程数字化</span>
              </div>
            </div>

            <div class="feature-card" style="--delay: 1">
              <div class="card-icon">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
                  <path d="M3 3v18h18"/>
                  <path d="M18 9l-5 5-2-2-4 4"/>
                </svg>
              </div>
              <div class="card-content">
                <span class="card-title">数据洞察</span>
                <span class="card-desc">实时分析决策</span>
              </div>
            </div>

            <div class="feature-card" style="--delay: 2">
              <div class="card-icon">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
                  <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/>
                  <path d="M9 12l2 2 4-4"/>
                </svg>
              </div>
              <div class="card-content">
                <span class="card-title">安全可靠</span>
                <span class="card-desc">企业级防护</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 右侧登录区 -->
      <div class="login-area">
        <div class="login-panel">
          <!-- 面板装饰 -->
          <div class="panel-corner tl"></div>
          <div class="panel-corner tr"></div>
          <div class="panel-corner bl"></div>
          <div class="panel-corner br"></div>

          <!-- 登录头部 -->
          <div class="login-header">
            <h2 class="login-title">系统登录</h2>
            <p class="login-subtitle">System Login</p>
          </div>

          <!-- 登录表单 -->
          <el-form
            ref="loginForm"
            :model="loginFormData"
            :rules="rules"
            :validate-on-rule-change="false"
            class="login-form"
            @keyup.enter="submitForm"
          >
            <el-form-item prop="username" class="form-item">
              <div class="input-wrapper">
                <div class="input-icon">
                  <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
                    <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/>
                    <circle cx="12" cy="7" r="4"/>
                  </svg>
                </div>
                <el-input
                  v-model="loginFormData.username"
                  placeholder="请输入用户名"
                  class="custom-input"
                />
              </div>
            </el-form-item>

            <el-form-item prop="password" class="form-item">
              <div class="input-wrapper">
                <div class="input-icon">
                  <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
                    <rect x="3" y="11" width="18" height="11" rx="2" ry="2"/>
                    <path d="M7 11V7a5 5 0 0 1 10 0v4"/>
                  </svg>
                </div>
                <el-input
                  v-model="loginFormData.password"
                  type="password"
                  placeholder="请输入密码"
                  show-password
                  class="custom-input"
                />
              </div>
            </el-form-item>

            <el-form-item v-if="loginFormData.openCaptcha" prop="captcha" class="form-item">
              <div class="captcha-row">
                <div class="input-wrapper captcha-input">
                  <div class="input-icon">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
                      <path d="M12 2L2 7l10 5 10-5-10-5z"/>
                      <path d="M2 17l10 5 10-5"/>
                      <path d="M2 12l10 5 10-5"/>
                    </svg>
                  </div>
                  <el-input
                    v-model="loginFormData.captcha"
                    placeholder="验证码"
                    class="custom-input"
                  />
                </div>
                <div class="captcha-img" @click="loginVerify()">
                  <img v-if="picPath" :src="picPath" alt="验证码" />
                  <span v-else class="captcha-loading">加载中</span>
                </div>
              </div>
            </el-form-item>

            <el-form-item class="form-item submit-item">
              <button type="button" class="submit-btn" @click="submitForm">
                <span class="btn-text">登 录</span>
                <span class="btn-icon">
                  <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M5 12h14M12 5l7 7-7 7"/>
                  </svg>
                </span>
                <div class="btn-shine"></div>
              </button>
            </el-form-item>
          </el-form>

          <!-- 底部操作 -->
          <div class="login-footer">
            <button class="footer-link" @click="checkInit">
              <span class="link-dot"></span>
              <span>系统初始化</span>
            </button>
            <button class="footer-link" @click="router.push({ name: 'SupplierDemo' })">
              <span class="link-dot"></span>
              <span>查看演示</span>
            </button>
          </div>

          <!-- 版权信息 -->
          <div class="copyright">
            <span>© 2024 叭盛科技 · Basheng Technology</span>
          </div>
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
// 引入字体
@import url('https://fonts.googleapis.com/css2?family=Orbitron:wght@400;500;600;700&family=Rajdhani:wght@300;400;500;600;700&display=swap');

// 颜色变量
:root {
  --bg-deep: #0a0e17;
  --bg-dark: #0d1320;
  --bg-card: rgba(15, 23, 42, 0.8);
  --border-color: rgba(56, 189, 248, 0.15);
  --border-glow: rgba(56, 189, 248, 0.4);
  --accent-primary: #38bdf8;
  --accent-secondary: #818cf8;
  --accent-tertiary: #22d3ee;
  --text-primary: #f1f5f9;
  --text-secondary: #94a3b8;
  --text-muted: #64748b;
  --glow-cyan: rgba(34, 211, 238, 0.5);
  --glow-blue: rgba(56, 189, 248, 0.5);
  --glow-purple: rgba(129, 140, 248, 0.4);
}

.login-container {
  position: relative;
  width: 100vw;
  height: 100vh;
  background: var(--bg-deep);
  overflow: hidden;
  font-family: 'Rajdhani', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}

// 背景层
.bg-layer {
  position: absolute;
  inset: 0;
  overflow: hidden;
}

// 网格线背景
.grid-lines {
  position: absolute;
  inset: 0;
  background-image:
    linear-gradient(rgba(56, 189, 248, 0.03) 1px, transparent 1px),
    linear-gradient(90deg, rgba(56, 189, 248, 0.03) 1px, transparent 1px);
  background-size: 60px 60px;
  animation: gridMove 20s linear infinite;
}

@keyframes gridMove {
  0% { transform: translate(0, 0); }
  100% { transform: translate(60px, 60px); }
}

// 浮动几何形状
.floating-shapes {
  position: absolute;
  inset: 0;
  pointer-events: none;
}

.shape {
  position: absolute;
  border: 1px solid;
  opacity: 0.15;
  animation: float 20s ease-in-out infinite;

  &.shape-1 {
    top: 15%;
    left: 10%;
    width: 120px;
    height: 120px;
    border-color: var(--accent-primary);
    transform: rotate(45deg);
    animation-delay: 0s;
  }

  &.shape-2 {
    top: 60%;
    left: 5%;
    width: 80px;
    height: 80px;
    border-color: var(--accent-secondary);
    border-radius: 50%;
    animation-delay: -5s;
  }

  &.shape-3 {
    top: 25%;
    right: 35%;
    width: 60px;
    height: 60px;
    border-color: var(--accent-tertiary);
    animation-delay: -10s;
  }

  &.shape-4 {
    bottom: 20%;
    left: 25%;
    width: 100px;
    height: 100px;
    border-color: var(--accent-primary);
    border-radius: 50%;
    animation-delay: -7s;
  }

  &.shape-5 {
    top: 70%;
    right: 45%;
    width: 50px;
    height: 50px;
    border-color: var(--accent-secondary);
    transform: rotate(30deg);
    animation-delay: -3s;
  }
}

@keyframes float {
  0%, 100% { transform: translateY(0) rotate(0deg); }
  25% { transform: translateY(-20px) rotate(5deg); }
  50% { transform: translateY(0) rotate(0deg); }
  75% { transform: translateY(20px) rotate(-5deg); }
}

// 光晕效果
.glow-orb {
  position: absolute;
  border-radius: 50%;
  filter: blur(80px);
  opacity: 0.4;
  animation: orbPulse 8s ease-in-out infinite;

  &.glow-1 {
    top: 10%;
    left: 20%;
    width: 400px;
    height: 400px;
    background: var(--glow-blue);
    animation-delay: 0s;
  }

  &.glow-2 {
    bottom: 10%;
    left: 30%;
    width: 300px;
    height: 300px;
    background: var(--glow-purple);
    animation-delay: -3s;
  }

  &.glow-3 {
    top: 40%;
    right: 20%;
    width: 350px;
    height: 350px;
    background: var(--glow-cyan);
    animation-delay: -5s;
  }
}

@keyframes orbPulse {
  0%, 100% { transform: scale(1); opacity: 0.3; }
  50% { transform: scale(1.1); opacity: 0.5; }
}

// 主内容
.main-content {
  position: relative;
  z-index: 1;
  display: flex;
  height: 100%;
}

// 左侧品牌区
.brand-area {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 60px;
}

.brand-wrapper {
  max-width: 500px;
}

// Logo
.logo-container {
  position: relative;
  width: 100px;
  height: 100px;
  margin-bottom: 40px;
}

.logo-ring {
  position: absolute;
  inset: 0;
  border: 2px solid var(--accent-primary);
  border-radius: 24px;
  transform: rotate(45deg);
  animation: ringRotate 10s linear infinite;

  .ring-inner {
    position: absolute;
    inset: 8px;
    border: 1px solid var(--accent-tertiary);
    border-radius: 18px;
    opacity: 0.5;
  }
}

@keyframes ringRotate {
  0% { transform: rotate(45deg); }
  100% { transform: rotate(405deg); }
}

.logo-core {
  position: absolute;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 2px;
}

.logo-char {
  font-family: 'Orbitron', monospace;
  font-size: 28px;
  font-weight: 700;
  color: var(--text-primary);
  text-shadow: 0 0 20px var(--glow-cyan);
}

// 品牌文字
.brand-text {
  margin-bottom: 32px;
}

.brand-name {
  font-size: 48px;
  font-weight: 700;
  color: var(--text-primary);
  margin: 0 0 8px 0;
  letter-spacing: 8px;
  display: flex;
  gap: 4px;
}

.char {
  display: inline-block;
  animation: charReveal 0.8s ease-out forwards;
  animation-delay: calc(var(--i) * 0.1s);
  opacity: 0;
  transform: translateY(20px);
}

@keyframes charReveal {
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.brand-en {
  font-family: 'Orbitron', monospace;
  font-size: 14px;
  font-weight: 500;
  color: var(--accent-primary);
  letter-spacing: 4px;
  opacity: 0;
  animation: fadeIn 0.8s ease-out 0.5s forwards;
}

@keyframes fadeIn {
  to { opacity: 1; }
}

// 装饰线
.deco-line {
  position: relative;
  width: 100%;
  height: 2px;
  background: linear-gradient(90deg, var(--accent-primary), var(--accent-tertiary), transparent);
  margin-bottom: 32px;
  opacity: 0;
  animation: fadeIn 0.8s ease-out 0.6s forwards;
}

.line-pulse {
  position: absolute;
  left: 0;
  top: -2px;
  width: 6px;
  height: 6px;
  background: var(--accent-primary);
  border-radius: 50%;
  box-shadow: 0 0 10px var(--glow-cyan), 0 0 20px var(--glow-cyan);
  animation: linePulse 3s ease-in-out infinite;
}

@keyframes linePulse {
  0%, 100% { left: 0; }
  50% { left: calc(100% - 6px); }
}

// Slogan
.slogan {
  display: flex;
  align-items: center;
  gap: 16px;
  margin: 0 0 48px 0;
  opacity: 0;
  animation: fadeIn 0.8s ease-out 0.7s forwards;
}

.slogan-word {
  font-size: 16px;
  font-weight: 500;
  color: var(--text-secondary);
  letter-spacing: 2px;
}

.slogan-dot {
  width: 4px;
  height: 4px;
  background: var(--accent-primary);
  border-radius: 50%;
  box-shadow: 0 0 8px var(--glow-cyan);
}

// 特性卡片
.feature-cards {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.feature-card {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 16px 20px;
  background: rgba(15, 23, 42, 0.5);
  border: 1px solid var(--border-color);
  border-radius: 12px;
  backdrop-filter: blur(10px);
  transition: all 0.3s ease;
  opacity: 0;
  transform: translateX(-20px);
  animation: cardSlide 0.6s ease-out forwards;
  animation-delay: calc(0.8s + var(--delay) * 0.15s);

  &:hover {
    border-color: var(--border-glow);
    background: rgba(15, 23, 42, 0.7);
    transform: translateX(8px);
    box-shadow: 0 0 30px rgba(56, 189, 248, 0.1);
  }
}

@keyframes cardSlide {
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

.card-icon {
  width: 44px;
  height: 44px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, rgba(56, 189, 248, 0.15), rgba(129, 140, 248, 0.1));
  border-radius: 10px;
  color: var(--accent-primary);

  svg {
    width: 22px;
    height: 22px;
  }
}

.card-content {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.card-title {
  font-size: 15px;
  font-weight: 600;
  color: var(--text-primary);
}

.card-desc {
  font-size: 12px;
  color: var(--text-muted);
}

// 右侧登录区
.login-area {
  width: 520px;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 60px;
  background: rgba(10, 14, 23, 0.6);
  backdrop-filter: blur(20px);
  border-left: 1px solid var(--border-color);
}

// 登录面板
.login-panel {
  position: relative;
  width: 100%;
  max-width: 380px;
  padding: 48px 40px;
  background: var(--bg-card);
  border: 1px solid var(--border-color);
  border-radius: 20px;
  backdrop-filter: blur(20px);
  animation: panelFade 0.8s ease-out;
}

@keyframes panelFade {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

// 面板角落装饰
.panel-corner {
  position: absolute;
  width: 20px;
  height: 20px;
  border-color: var(--accent-primary);
  border-style: solid;
  border-width: 0;

  &.tl {
    top: -1px;
    left: -1px;
    border-top-width: 2px;
    border-left-width: 2px;
    border-radius: 20px 0 0 0;
  }

  &.tr {
    top: -1px;
    right: -1px;
    border-top-width: 2px;
    border-right-width: 2px;
    border-radius: 0 20px 0 0;
  }

  &.bl {
    bottom: -1px;
    left: -1px;
    border-bottom-width: 2px;
    border-left-width: 2px;
    border-radius: 0 0 0 20px;
  }

  &.br {
    bottom: -1px;
    right: -1px;
    border-bottom-width: 2px;
    border-right-width: 2px;
    border-radius: 0 0 20px 0;
  }
}

// 登录头部
.login-header {
  text-align: center;
  margin-bottom: 40px;
}

.login-title {
  font-size: 28px;
  font-weight: 700;
  color: var(--text-primary);
  margin: 0 0 8px 0;
  letter-spacing: 4px;
}

.login-subtitle {
  font-family: 'Orbitron', monospace;
  font-size: 12px;
  color: var(--text-muted);
  letter-spacing: 3px;
  margin: 0;
}

// 表单样式
.login-form {
  :deep(.el-form-item) {
    margin-bottom: 24px;
  }

  :deep(.el-form-item__error) {
    font-size: 12px;
    color: #f87171;
    padding-top: 4px;
  }
}

.form-item {
  margin-bottom: 24px;
}

.input-wrapper {
  position: relative;
  display: flex;
  align-items: center;
  background: rgba(15, 23, 42, 0.6);
  border: 1px solid var(--border-color);
  border-radius: 12px;
  transition: all 0.3s ease;
  overflow: hidden;

  &:hover {
    border-color: rgba(56, 189, 248, 0.3);
  }

  &:focus-within {
    border-color: var(--accent-primary);
    box-shadow: 0 0 0 3px rgba(56, 189, 248, 0.1), 0 0 20px rgba(56, 189, 248, 0.1);
  }
}

.input-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 50px;
  height: 50px;
  color: var(--text-muted);
  flex-shrink: 0;

  svg {
    width: 20px;
    height: 20px;
  }
}

.custom-input {
  flex: 1;

  :deep(.el-input__wrapper) {
    background: transparent;
    box-shadow: none;
    padding: 0 16px 0 0;

    &.is-focus {
      box-shadow: none;
    }
  }

  :deep(.el-input__inner) {
    height: 50px;
    font-size: 15px;
    color: var(--text-primary);
    font-family: inherit;

    &::placeholder {
      color: var(--text-muted);
    }
  }

  :deep(.el-input__suffix) {
    color: var(--text-muted);
  }
}

// 验证码行
.captcha-row {
  display: flex;
  gap: 12px;
}

.captcha-input {
  flex: 1;
}

.captcha-img {
  width: 120px;
  height: 52px;
  border-radius: 12px;
  overflow: hidden;
  cursor: pointer;
  background: rgba(15, 23, 42, 0.8);
  border: 1px solid var(--border-color);
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s ease;

  &:hover {
    border-color: var(--accent-primary);
  }

  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }

  .captcha-loading {
    font-size: 12px;
    color: var(--text-muted);
  }
}

// 提交按钮
.submit-item {
  margin-top: 32px;
  margin-bottom: 0;
}

.submit-btn {
  position: relative;
  width: 100%;
  height: 52px;
  background: linear-gradient(135deg, var(--accent-primary), var(--accent-secondary));
  border: none;
  border-radius: 12px;
  color: white;
  font-family: inherit;
  font-size: 16px;
  font-weight: 600;
  letter-spacing: 4px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  overflow: hidden;
  transition: all 0.3s ease;

  &:hover {
    transform: translateY(-2px);
    box-shadow: 0 10px 40px rgba(56, 189, 248, 0.3);

    .btn-icon {
      transform: translateX(4px);
    }

    .btn-shine {
      left: 100%;
    }
  }

  &:active {
    transform: translateY(0);
  }
}

.btn-text {
  position: relative;
  z-index: 1;
}

.btn-icon {
  position: relative;
  z-index: 1;
  display: flex;
  transition: transform 0.3s ease;

  svg {
    width: 18px;
    height: 18px;
  }
}

.btn-shine {
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255,255,255,0.2), transparent);
  transition: left 0.5s ease;
}

// 底部链接
.login-footer {
  display: flex;
  justify-content: center;
  gap: 32px;
  margin-top: 28px;
}

.footer-link {
  display: flex;
  align-items: center;
  gap: 8px;
  background: none;
  border: none;
  color: var(--text-muted);
  font-family: inherit;
  font-size: 13px;
  cursor: pointer;
  transition: all 0.3s ease;
  padding: 4px 0;

  &:hover {
    color: var(--accent-primary);

    .link-dot {
      background: var(--accent-primary);
      box-shadow: 0 0 8px var(--glow-cyan);
    }
  }
}

.link-dot {
  width: 4px;
  height: 4px;
  background: var(--text-muted);
  border-radius: 50%;
  transition: all 0.3s ease;
}

// 版权
.copyright {
  text-align: center;
  margin-top: 32px;
  padding-top: 24px;
  border-top: 1px solid var(--border-color);

  span {
    font-size: 11px;
    color: var(--text-muted);
    letter-spacing: 1px;
  }
}

// 响应式
@media (max-width: 1024px) {
  .main-content {
    flex-direction: column;
  }

  .brand-area {
    flex: none;
    padding: 40px 24px;
  }

  .brand-wrapper {
    max-width: 100%;
    text-align: center;
  }

  .logo-container {
    margin: 0 auto 32px;
  }

  .brand-name {
    justify-content: center;
    font-size: 36px;
  }

  .slogan {
    justify-content: center;
    margin-bottom: 32px;
  }

  .feature-cards {
    display: none;
  }

  .deco-line {
    max-width: 200px;
    margin: 0 auto 24px;
  }

  .login-area {
    width: 100%;
    flex: 1;
    padding: 24px;
    border-left: none;
    border-top: 1px solid var(--border-color);
  }

  .login-panel {
    padding: 32px 24px;
  }
}

@media (max-width: 480px) {
  .brand-name {
    font-size: 28px;
    letter-spacing: 4px;
  }

  .brand-en {
    font-size: 11px;
  }

  .slogan-word {
    font-size: 13px;
  }

  .login-title {
    font-size: 24px;
  }
}
</style>
