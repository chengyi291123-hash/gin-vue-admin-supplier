<template>
  <div class="login-container">
    <!-- 动态背景 -->
    <div class="bg-layer">
      <!-- 渐变背景 -->
      <div class="gradient-bg"></div>
      <!-- 装饰圆形 -->
      <div class="deco-circle circle-1"></div>
      <div class="deco-circle circle-2"></div>
      <div class="deco-circle circle-3"></div>
      <!-- 流动线条 -->
      <svg class="flow-lines" viewBox="0 0 1920 1080" preserveAspectRatio="xMidYMid slice">
        <path class="flow-path path-1" d="M-100,200 Q400,100 600,300 T1000,250 T1400,350 T2000,200" />
        <path class="flow-path path-2" d="M-100,400 Q300,500 500,350 T900,450 T1300,300 T2000,400" />
        <path class="flow-path path-3" d="M-100,600 Q400,700 700,550 T1100,650 T1500,500 T2000,600" />
      </svg>
      <!-- 浮动粒子 -->
      <div class="particles">
        <span v-for="i in 20" :key="i" class="particle" :style="getParticleStyle(i)"></span>
      </div>
    </div>

    <!-- 主内容区 -->
    <div class="main-content">
      <!-- 左侧品牌区 -->
      <div class="brand-area">
        <div class="brand-wrapper">
          <!-- Logo -->
          <div class="logo-container">
            <div class="logo-bg">
              <div class="logo-ring"></div>
            </div>
            <div class="logo-text">
              <span class="logo-b">B</span>
              <span class="logo-s">S</span>
            </div>
          </div>

          <!-- 品牌名称 -->
          <div class="brand-text">
            <h1 class="brand-name">
              <span class="char" v-for="(char, index) in '叭盛科技'" :key="index" :style="`--i: ${index}`">{{ char }}</span>
            </h1>
            <div class="brand-en">BASHENG TECHNOLOGY</div>
          </div>

          <!-- 分隔线 -->
          <div class="divider">
            <div class="divider-line"></div>
            <div class="divider-dot"></div>
            <div class="divider-line"></div>
          </div>

          <!-- Slogan -->
          <p class="slogan">
            <span class="slogan-item">
              <span class="slogan-icon">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <circle cx="12" cy="12" r="3"/>
                  <path d="M12 1v2m0 18v2M4.22 4.22l1.42 1.42m12.72 12.72l1.42 1.42M1 12h2m18 0h2M4.22 19.78l1.42-1.42M18.36 5.64l1.42-1.42"/>
                </svg>
              </span>
              <span class="slogan-text">智能驱动</span>
            </span>
            <span class="slogan-item">
              <span class="slogan-icon">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"/>
                  <polyline points="7.5 4.21 12 6.81 16.5 4.21"/>
                  <polyline points="7.5 19.79 7.5 14.6 3 12"/>
                  <polyline points="21 12 16.5 14.6 16.5 19.79"/>
                  <polyline points="3.27 6.96 12 12.01 20.73 6.96"/>
                  <line x1="12" y1="22.08" x2="12" y2="12"/>
                </svg>
              </span>
              <span class="slogan-text">数据赋能</span>
            </span>
            <span class="slogan-item">
              <span class="slogan-icon">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/>
                </svg>
              </span>
              <span class="slogan-text">安全护航</span>
            </span>
          </p>

          <!-- 特性卡片 -->
          <div class="feature-cards">
            <div class="feature-card" v-for="(feature, index) in features" :key="index" :style="`--delay: ${index}`">
              <div class="card-icon" :style="`background: ${feature.gradient}`">
                <component :is="feature.icon" />
              </div>
              <div class="card-info">
                <span class="card-title">{{ feature.title }}</span>
                <span class="card-desc">{{ feature.desc }}</span>
              </div>
              <div class="card-arrow">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M5 12h14M12 5l7 7-7 7"/>
                </svg>
              </div>
            </div>
          </div>

          <!-- 底部装饰 -->
          <div class="brand-footer">
            <div class="stats">
              <div class="stat-item">
                <span class="stat-value">99.9%</span>
                <span class="stat-label">系统稳定性</span>
              </div>
              <div class="stat-divider"></div>
              <div class="stat-item">
                <span class="stat-value">500+</span>
                <span class="stat-label">企业客户</span>
              </div>
              <div class="stat-divider"></div>
              <div class="stat-item">
                <span class="stat-value">24/7</span>
                <span class="stat-label">技术支持</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 右侧登录区 -->
      <div class="login-area">
        <div class="login-card">
          <!-- 卡片光效 -->
          <div class="card-glow"></div>

          <!-- 登录头部 -->
          <div class="login-header">
            <div class="welcome-icon">
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
                <path d="M15 3h4a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-4"/>
                <polyline points="10 17 15 12 10 7"/>
                <line x1="15" y1="12" x2="3" y2="12"/>
              </svg>
            </div>
            <h2 class="login-title">欢迎回来</h2>
            <p class="login-subtitle">Welcome Back</p>
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
              <div class="input-group">
                <label class="input-label">用户名</label>
                <div class="input-wrapper">
                  <span class="input-icon">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                      <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/>
                      <circle cx="12" cy="7" r="4"/>
                    </svg>
                  </span>
                  <el-input
                    v-model="loginFormData.username"
                    placeholder="请输入用户名"
                    class="custom-input"
                  />
                </div>
              </div>
            </el-form-item>

            <el-form-item prop="password" class="form-item">
              <div class="input-group">
                <label class="input-label">密码</label>
                <div class="input-wrapper">
                  <span class="input-icon">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                      <rect x="3" y="11" width="18" height="11" rx="2" ry="2"/>
                      <path d="M7 11V7a5 5 0 0 1 10 0v4"/>
                    </svg>
                  </span>
                  <el-input
                    v-model="loginFormData.password"
                    type="password"
                    placeholder="请输入密码"
                    show-password
                    class="custom-input"
                  />
                </div>
              </div>
            </el-form-item>

            <el-form-item v-if="loginFormData.openCaptcha" prop="captcha" class="form-item">
              <div class="input-group">
                <label class="input-label">验证码</label>
                <div class="captcha-row">
                  <div class="input-wrapper captcha-input">
                    <span class="input-icon">
                      <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <rect x="3" y="3" width="7" height="7"/>
                        <rect x="14" y="3" width="7" height="7"/>
                        <rect x="14" y="14" width="7" height="7"/>
                        <rect x="3" y="14" width="7" height="7"/>
                      </svg>
                    </span>
                    <el-input
                      v-model="loginFormData.captcha"
                      placeholder="请输入验证码"
                      class="custom-input"
                    />
                  </div>
                  <div class="captcha-box" @click="loginVerify()">
                    <img v-if="picPath" :src="picPath" alt="验证码" />
                    <span v-else class="captcha-loading">
                      <svg class="loading-spin" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M21 12a9 9 0 1 1-6.219-8.56"/>
                      </svg>
                    </span>
                  </div>
                </div>
              </div>
            </el-form-item>

            <el-form-item class="form-item submit-item">
              <button type="button" class="login-btn" @click="submitForm">
                <span class="btn-content">
                  <span class="btn-text">立即登录</span>
                  <span class="btn-arrow">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                      <line x1="5" y1="12" x2="19" y2="12"/>
                      <polyline points="12 5 19 12 12 19"/>
                    </svg>
                  </span>
                </span>
                <div class="btn-bg"></div>
              </button>
            </el-form-item>
          </el-form>

          <!-- 快捷操作 -->
          <div class="login-actions">
            <button class="action-btn" @click="checkInit">
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M12 20h9"/>
                <path d="M16.5 3.5a2.121 2.121 0 0 1 3 3L7 19l-4 1 1-4L16.5 3.5z"/>
              </svg>
              <span>系统初始化</span>
            </button>
            <button class="action-btn" @click="router.push({ name: 'SupplierDemo' })">
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <polygon points="5 3 19 12 5 21 5 3"/>
              </svg>
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
import { reactive, ref, h } from 'vue'
import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/pinia/modules/user'

defineOptions({
  name: 'Login'
})

const router = useRouter()
const captchaRequiredLength = ref(6)

// 特性数据
const features = [
  {
    title: '智慧管理',
    desc: '全流程数字化管控',
    gradient: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
    icon: {
      render() {
        return h('svg', { viewBox: '0 0 24 24', fill: 'none', stroke: 'currentColor', 'stroke-width': '2' }, [
          h('rect', { x: '3', y: '3', width: '18', height: '18', rx: '2', ry: '2' }),
          h('line', { x1: '3', y1: '9', x2: '21', y2: '9' }),
          h('line', { x1: '9', y1: '21', x2: '9', y2: '9' })
        ])
      }
    }
  },
  {
    title: '数据洞察',
    desc: '实时分析决策支持',
    gradient: 'linear-gradient(135deg, #f093fb 0%, #f5576c 100%)',
    icon: {
      render() {
        return h('svg', { viewBox: '0 0 24 24', fill: 'none', stroke: 'currentColor', 'stroke-width': '2' }, [
          h('line', { x1: '18', y1: '20', x2: '18', y2: '10' }),
          h('line', { x1: '12', y1: '20', x2: '12', y2: '4' }),
          h('line', { x1: '6', y1: '20', x2: '6', y2: '14' })
        ])
      }
    }
  },
  {
    title: '安全可靠',
    desc: '企业级安全防护',
    gradient: 'linear-gradient(135deg, #4facfe 0%, #00f2fe 100%)',
    icon: {
      render() {
        return h('svg', { viewBox: '0 0 24 24', fill: 'none', stroke: 'currentColor', 'stroke-width': '2' }, [
          h('path', { d: 'M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z' }),
          h('polyline', { points: '9 12 11 14 15 10' })
        ])
      }
    }
  }
]

// 粒子样式生成
const getParticleStyle = (i) => {
  const size = Math.random() * 6 + 2
  const left = Math.random() * 100
  const delay = Math.random() * 20
  const duration = Math.random() * 10 + 15
  return {
    width: `${size}px`,
    height: `${size}px`,
    left: `${left}%`,
    animationDelay: `${delay}s`,
    animationDuration: `${duration}s`
  }
}

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
@import url('https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@400;500;600;700;800&family=Noto+Sans+SC:wght@400;500;600;700&display=swap');

// 颜色变量 - 明亮现代风格
$primary: #4f46e5;
$primary-light: #818cf8;
$primary-dark: #3730a3;
$secondary: #06b6d4;
$accent: #f472b6;
$success: #10b981;

$bg-start: #f8fafc;
$bg-end: #e0e7ff;
$card-bg: rgba(255, 255, 255, 0.85);
$text-primary: #1e293b;
$text-secondary: #64748b;
$text-muted: #94a3b8;
$border-color: rgba(148, 163, 184, 0.2);

.login-container {
  position: relative;
  width: 100vw;
  height: 100vh;
  overflow: hidden;
  font-family: 'Plus Jakarta Sans', 'Noto Sans SC', -apple-system, sans-serif;
}

// 背景层
.bg-layer {
  position: absolute;
  inset: 0;
  overflow: hidden;
}

.gradient-bg {
  position: absolute;
  inset: 0;
  background: linear-gradient(135deg, $bg-start 0%, $bg-end 50%, #fae8ff 100%);
}

// 装饰圆形
.deco-circle {
  position: absolute;
  border-radius: 50%;
  filter: blur(60px);
  opacity: 0.6;
  animation: floatCircle 20s ease-in-out infinite;

  &.circle-1 {
    width: 600px;
    height: 600px;
    top: -200px;
    right: -100px;
    background: linear-gradient(135deg, rgba(129, 140, 248, 0.4), rgba(167, 139, 250, 0.3));
    animation-delay: 0s;
  }

  &.circle-2 {
    width: 400px;
    height: 400px;
    bottom: -100px;
    left: 10%;
    background: linear-gradient(135deg, rgba(6, 182, 212, 0.3), rgba(16, 185, 129, 0.2));
    animation-delay: -7s;
  }

  &.circle-3 {
    width: 300px;
    height: 300px;
    top: 40%;
    left: 30%;
    background: linear-gradient(135deg, rgba(244, 114, 182, 0.25), rgba(251, 146, 60, 0.2));
    animation-delay: -14s;
  }
}

@keyframes floatCircle {
  0%, 100% { transform: translate(0, 0) scale(1); }
  33% { transform: translate(30px, -30px) scale(1.05); }
  66% { transform: translate(-20px, 20px) scale(0.95); }
}

// 流动线条
.flow-lines {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
}

.flow-path {
  fill: none;
  stroke-width: 2;
  stroke-linecap: round;
  opacity: 0.15;

  &.path-1 {
    stroke: $primary;
    stroke-dasharray: 20 10;
    animation: flowLine 15s linear infinite;
  }

  &.path-2 {
    stroke: $secondary;
    stroke-dasharray: 15 15;
    animation: flowLine 20s linear infinite reverse;
  }

  &.path-3 {
    stroke: $accent;
    stroke-dasharray: 25 5;
    animation: flowLine 18s linear infinite;
  }
}

@keyframes flowLine {
  0% { stroke-dashoffset: 0; }
  100% { stroke-dashoffset: 1000; }
}

// 浮动粒子
.particles {
  position: absolute;
  inset: 0;
  pointer-events: none;
}

.particle {
  position: absolute;
  bottom: -10px;
  background: linear-gradient(135deg, $primary-light, $secondary);
  border-radius: 50%;
  opacity: 0.4;
  animation: particleFloat linear infinite;
}

@keyframes particleFloat {
  0% {
    transform: translateY(0) rotate(0deg);
    opacity: 0;
  }
  10% {
    opacity: 0.4;
  }
  90% {
    opacity: 0.4;
  }
  100% {
    transform: translateY(-100vh) rotate(720deg);
    opacity: 0;
  }
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
  padding: 60px 80px;
}

.brand-wrapper {
  max-width: 520px;
}

// Logo
.logo-container {
  position: relative;
  width: 88px;
  height: 88px;
  margin-bottom: 32px;
}

.logo-bg {
  position: absolute;
  inset: 0;
  background: linear-gradient(135deg, $primary 0%, $primary-light 100%);
  border-radius: 24px;
  transform: rotate(-10deg);
  box-shadow: 0 20px 40px -10px rgba(79, 70, 229, 0.4);
  animation: logoFloat 6s ease-in-out infinite;
}

.logo-ring {
  position: absolute;
  inset: -8px;
  border: 2px dashed rgba(255, 255, 255, 0.3);
  border-radius: 30px;
  animation: ringRotate 20s linear infinite;
}

@keyframes logoFloat {
  0%, 100% { transform: rotate(-10deg) translateY(0); }
  50% { transform: rotate(-10deg) translateY(-8px); }
}

@keyframes ringRotate {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.logo-text {
  position: absolute;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 2px;
  font-family: 'Plus Jakarta Sans', sans-serif;
  font-size: 32px;
  font-weight: 800;
  color: white;
  text-shadow: 0 2px 10px rgba(0, 0, 0, 0.2);
}

// 品牌文字
.brand-text {
  margin-bottom: 24px;
}

.brand-name {
  font-size: 52px;
  font-weight: 700;
  color: $text-primary;
  margin: 0 0 8px 0;
  letter-spacing: 4px;
  display: flex;
}

.char {
  display: inline-block;
  animation: charBounce 0.6s ease-out forwards;
  animation-delay: calc(var(--i) * 0.08s);
  opacity: 0;
  transform: translateY(30px);
}

@keyframes charBounce {
  0% { opacity: 0; transform: translateY(30px); }
  60% { transform: translateY(-5px); }
  100% { opacity: 1; transform: translateY(0); }
}

.brand-en {
  font-family: 'Plus Jakarta Sans', sans-serif;
  font-size: 14px;
  font-weight: 600;
  color: $primary;
  letter-spacing: 3px;
  opacity: 0;
  animation: fadeSlide 0.6s ease-out 0.4s forwards;
}

@keyframes fadeSlide {
  from { opacity: 0; transform: translateX(-10px); }
  to { opacity: 1; transform: translateX(0); }
}

// 分隔线
.divider {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 24px;
  opacity: 0;
  animation: fadeSlide 0.6s ease-out 0.5s forwards;
}

.divider-line {
  flex: 1;
  height: 2px;
  background: linear-gradient(90deg, $primary, transparent);
  border-radius: 1px;

  &:last-child {
    background: linear-gradient(90deg, transparent, $secondary);
  }
}

.divider-dot {
  width: 8px;
  height: 8px;
  background: linear-gradient(135deg, $primary, $secondary);
  border-radius: 50%;
  animation: dotPulse 2s ease-in-out infinite;
}

@keyframes dotPulse {
  0%, 100% { transform: scale(1); opacity: 1; }
  50% { transform: scale(1.2); opacity: 0.7; }
}

// Slogan
.slogan {
  display: flex;
  gap: 32px;
  margin: 0 0 40px 0;
  opacity: 0;
  animation: fadeSlide 0.6s ease-out 0.6s forwards;
}

.slogan-item {
  display: flex;
  align-items: center;
  gap: 8px;
}

.slogan-icon {
  width: 20px;
  height: 20px;
  color: $primary;

  svg {
    width: 100%;
    height: 100%;
  }
}

.slogan-text {
  font-size: 14px;
  font-weight: 500;
  color: $text-secondary;
}

// 特性卡片
.feature-cards {
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-bottom: 40px;
}

.feature-card {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 16px 20px;
  background: $card-bg;
  border: 1px solid $border-color;
  border-radius: 16px;
  backdrop-filter: blur(20px);
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  opacity: 0;
  transform: translateX(-20px);
  animation: cardReveal 0.5s ease-out forwards;
  animation-delay: calc(0.7s + var(--delay) * 0.1s);

  &:hover {
    transform: translateX(8px);
    box-shadow: 0 10px 40px -10px rgba(79, 70, 229, 0.2);
    border-color: rgba(79, 70, 229, 0.3);

    .card-arrow {
      transform: translateX(4px);
      opacity: 1;
    }
  }
}

@keyframes cardReveal {
  to { opacity: 1; transform: translateX(0); }
}

.card-icon {
  width: 48px;
  height: 48px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 12px;
  color: white;
  flex-shrink: 0;

  svg {
    width: 24px;
    height: 24px;
  }
}

.card-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.card-title {
  font-size: 15px;
  font-weight: 600;
  color: $text-primary;
}

.card-desc {
  font-size: 13px;
  color: $text-muted;
}

.card-arrow {
  width: 20px;
  height: 20px;
  color: $primary;
  opacity: 0;
  transform: translateX(-4px);
  transition: all 0.3s ease;

  svg {
    width: 100%;
    height: 100%;
  }
}

// 底部统计
.brand-footer {
  opacity: 0;
  animation: fadeSlide 0.6s ease-out 1s forwards;
}

.stats {
  display: flex;
  align-items: center;
  gap: 24px;
  padding: 20px 24px;
  background: $card-bg;
  border: 1px solid $border-color;
  border-radius: 16px;
  backdrop-filter: blur(20px);
}

.stat-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
}

.stat-value {
  font-size: 24px;
  font-weight: 700;
  color: $primary;
}

.stat-label {
  font-size: 12px;
  color: $text-muted;
}

.stat-divider {
  width: 1px;
  height: 32px;
  background: $border-color;
}

// 右侧登录区
.login-area {
  width: 500px;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 40px;
  background: linear-gradient(180deg, rgba(255,255,255,0.5) 0%, rgba(255,255,255,0.8) 100%);
  backdrop-filter: blur(20px);
}

// 登录卡片
.login-card {
  position: relative;
  width: 100%;
  max-width: 400px;
  padding: 40px;
  background: white;
  border-radius: 24px;
  box-shadow:
    0 25px 50px -12px rgba(0, 0, 0, 0.1),
    0 0 0 1px rgba(0, 0, 0, 0.02);
  animation: cardFloat 0.8s ease-out;
}

@keyframes cardFloat {
  from { opacity: 0; transform: translateY(30px); }
  to { opacity: 1; transform: translateY(0); }
}

.card-glow {
  position: absolute;
  top: -2px;
  left: -2px;
  right: -2px;
  bottom: -2px;
  background: linear-gradient(135deg, $primary, $secondary, $accent);
  border-radius: 26px;
  opacity: 0;
  z-index: -1;
  transition: opacity 0.3s ease;
}

.login-card:hover .card-glow {
  opacity: 0.1;
}

// 登录头部
.login-header {
  text-align: center;
  margin-bottom: 32px;
}

.welcome-icon {
  width: 56px;
  height: 56px;
  margin: 0 auto 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, rgba(79, 70, 229, 0.1), rgba(6, 182, 212, 0.1));
  border-radius: 16px;
  color: $primary;

  svg {
    width: 28px;
    height: 28px;
  }
}

.login-title {
  font-size: 28px;
  font-weight: 700;
  color: $text-primary;
  margin: 0 0 4px 0;
}

.login-subtitle {
  font-family: 'Plus Jakarta Sans', sans-serif;
  font-size: 13px;
  font-weight: 500;
  color: $text-muted;
  letter-spacing: 2px;
  margin: 0;
}

// 表单样式
.login-form {
  :deep(.el-form-item) {
    margin-bottom: 20px;
  }

  :deep(.el-form-item__error) {
    font-size: 12px;
    color: #ef4444;
    padding-top: 4px;
  }
}

.form-item {
  margin-bottom: 20px;
}

.input-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.input-label {
  font-size: 13px;
  font-weight: 600;
  color: $text-secondary;
}

.input-wrapper {
  position: relative;
  display: flex;
  align-items: center;
  background: #f8fafc;
  border: 2px solid transparent;
  border-radius: 12px;
  transition: all 0.3s ease;

  &:hover {
    background: #f1f5f9;
  }

  &:focus-within {
    background: white;
    border-color: $primary;
    box-shadow: 0 0 0 4px rgba(79, 70, 229, 0.1);
  }
}

.input-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 48px;
  height: 48px;
  color: $text-muted;
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
    height: 48px;
    font-size: 15px;
    color: $text-primary;
    font-family: inherit;

    &::placeholder {
      color: $text-muted;
    }
  }

  :deep(.el-input__suffix) {
    color: $text-muted;
  }
}

// 验证码
.captcha-row {
  display: flex;
  gap: 12px;
}

.captcha-input {
  flex: 1;
}

.captcha-box {
  width: 120px;
  height: 52px;
  border-radius: 12px;
  overflow: hidden;
  cursor: pointer;
  background: #f8fafc;
  border: 2px solid transparent;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s ease;

  &:hover {
    border-color: $primary;
  }

  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
}

.captcha-loading {
  color: $text-muted;
}

.loading-spin {
  width: 24px;
  height: 24px;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

// 登录按钮
.submit-item {
  margin-top: 28px;
  margin-bottom: 0;
}

.login-btn {
  position: relative;
  width: 100%;
  height: 52px;
  background: transparent;
  border: none;
  border-radius: 12px;
  cursor: pointer;
  overflow: hidden;
}

.btn-content {
  position: relative;
  z-index: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  height: 100%;
  color: white;
}

.btn-text {
  font-family: inherit;
  font-size: 16px;
  font-weight: 600;
  letter-spacing: 2px;
}

.btn-arrow {
  display: flex;
  transition: transform 0.3s ease;

  svg {
    width: 20px;
    height: 20px;
  }
}

.btn-bg {
  position: absolute;
  inset: 0;
  background: linear-gradient(135deg, $primary 0%, $primary-dark 100%);
  transition: all 0.3s ease;
}

.login-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 10px 30px -5px rgba(79, 70, 229, 0.5);

  .btn-arrow {
    transform: translateX(4px);
  }

  .btn-bg {
    background: linear-gradient(135deg, $primary-light 0%, $primary 100%);
  }
}

.login-btn:active {
  transform: translateY(0);
}

// 快捷操作
.login-actions {
  display: flex;
  justify-content: center;
  gap: 24px;
  margin-top: 24px;
  padding-top: 24px;
  border-top: 1px solid $border-color;
}

.action-btn {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 8px 12px;
  background: transparent;
  border: none;
  border-radius: 8px;
  color: $text-muted;
  font-family: inherit;
  font-size: 13px;
  cursor: pointer;
  transition: all 0.3s ease;

  svg {
    width: 16px;
    height: 16px;
  }

  &:hover {
    color: $primary;
    background: rgba(79, 70, 229, 0.05);
  }
}

// 版权
.copyright {
  text-align: center;
  margin-top: 24px;

  span {
    font-size: 11px;
    color: $text-muted;
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
    margin: 0 auto 24px;
  }

  .brand-name {
    justify-content: center;
    font-size: 36px;
  }

  .divider {
    max-width: 200px;
    margin: 0 auto 24px;
  }

  .slogan {
    justify-content: center;
    flex-wrap: wrap;
    gap: 16px;
    margin-bottom: 24px;
  }

  .feature-cards,
  .brand-footer {
    display: none;
  }

  .login-area {
    width: 100%;
    flex: 1;
    padding: 24px;
  }

  .login-card {
    padding: 32px 24px;
  }
}

@media (max-width: 480px) {
  .brand-name {
    font-size: 28px;
    letter-spacing: 2px;
  }

  .brand-en {
    font-size: 11px;
  }

  .slogan-text {
    font-size: 12px;
  }

  .login-title {
    font-size: 24px;
  }
}
</style>
