<template>
  <div id="userLayout" class="w-full h-full relative overflow-hidden">
    <!-- 背景动画 -->
    <div class="absolute inset-0 bg-gradient-to-br from-blue-900 via-purple-900 to-indigo-900">
      <!-- 装饰网格 -->
      <div class="absolute inset-0 opacity-20">
        <div class="absolute inset-0 bg-pattern"></div>
      </div>
    </div>
    
    <!-- 浮动装饰元素 -->
    <div class="absolute top-20 left-20 w-32 h-32 bg-gradient-to-r from-blue-500 to-purple-500 rounded-full opacity-20 animate-pulse"></div>
    <div class="absolute bottom-20 right-20 w-24 h-24 bg-gradient-to-r from-purple-500 to-pink-500 rounded-full opacity-20 animate-pulse" style="animation-delay: 1s;"></div>
    <div class="absolute top-1/2 left-1/4 w-16 h-16 bg-gradient-to-r from-cyan-500 to-blue-500 rounded-full opacity-20 animate-pulse" style="animation-delay: 2s;"></div>
    
    <!-- 主要内容区域 -->
    <div class="relative z-10 flex items-center justify-center w-full h-full">
      <div class="w-full max-w-md mx-auto">
        <!-- 登录卡片 -->
        <div class="bg-white/10 backdrop-blur-lg rounded-2xl p-8 shadow-2xl border border-white/20">
          <!-- Logo区域 -->
          <div class="text-center mb-8">
            <div class="flex items-center justify-center mb-6">
              <img 
                class="w-48 h-12 object-contain logo-image" 
                :src="logoUrl" 
                alt="万联智控" 
                @error="handleLogoError"
              />
            </div>
            <div class="mb-6">
              <h1 class="text-3xl font-bold text-white mb-2">万联智控</h1>
              <p class="text-white/80 text-sm">智能设备管理平台</p>
            </div>
          </div>
          
          <!-- 登录表单 -->
          <el-form
            ref="loginForm"
            :model="loginFormData"
            :rules="rules"
            :validate-on-rule-change="false"
            @keyup.enter="submitForm"
            class="space-y-6"
          >
            <el-form-item prop="username">
              <el-input
                v-model="loginFormData.username"
                size="large"
                placeholder="请输入用户名"
                suffix-icon="user"
                class="modern-input"
              />
            </el-form-item>
            
            <el-form-item prop="password">
              <el-input
                v-model="loginFormData.password"
                show-password
                size="large"
                type="password"
                placeholder="请输入密码"
                class="modern-input"
              />
            </el-form-item>
            
            <el-form-item
              v-if="loginFormData.openCaptcha"
              prop="captcha"
            >
              <div class="flex w-full gap-4">
                <el-input
                  v-model="loginFormData.captcha"
                  placeholder="请输入验证码"
                  size="large"
                  class="flex-1 modern-input"
                />
                <div class="w-1/3 h-11 bg-white/20 backdrop-blur-sm rounded-lg border border-white/30 overflow-hidden">
                  <img
                    v-if="picPath"
                    class="w-full h-full object-cover cursor-pointer hover:opacity-80 transition-opacity"
                    :src="picPath"
                    alt="请输入验证码"
                    @click="loginVerify()"
                  />
                </div>
              </div>
            </el-form-item>
            
            <el-form-item>
              <el-button
                class="w-full h-12 bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700 text-white font-semibold rounded-lg shadow-lg hover:shadow-xl transition-all duration-300 transform hover:scale-105"
                size="large"
                @click="submitForm"
              >
                登 录
              </el-button>
            </el-form-item>
            
            <el-form-item>
              <el-button
                class="w-full h-12 bg-white/20 backdrop-blur-sm hover:bg-white/30 text-white font-semibold rounded-lg border border-white/30 shadow-lg hover:shadow-xl transition-all duration-300"
                size="large"
                @click="checkInit"
              >
                前往初始化
              </el-button>
            </el-form-item>
          </el-form>
        </div>
      </div>
    </div>

    <!-- 底部信息 -->
    <div class="absolute bottom-6 left-0 right-0 z-20">
      <div class="flex items-center justify-center gap-6">
        <a href="#" class="text-white/60 hover:text-white transition-colors">
          <i class="el-icon-document text-xl"></i>
        </a>
        <a href="#" class="text-white/60 hover:text-white transition-colors">
          <i class="el-icon-service text-xl"></i>
        </a>
        <a href="#" class="text-white/60 hover:text-white transition-colors">
          <i class="el-icon-link text-xl"></i>
        </a>
        <a href="#" class="text-white/60 hover:text-white transition-colors">
          <i class="el-icon-video-play text-xl"></i>
        </a>
      </div>
      <div class="text-center mt-4">
        <p class="text-white/40 text-sm">Powered by 万联智控团队</p>
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
  import logoUrl from '@/assets/logo_login.png'

  defineOptions({
    name: 'Login'
  })

  const router = useRouter()
  
  // 验证函数
  const checkUsername = (rule, value, callback) => {
    if (value.length < 5) {
      return callback(new Error('请输入正确的用户名'))
    } else {
      callback()
    }
  }
  
  const checkPassword = (rule, value, callback) => {
    if (value.length < 6) {
      return callback(new Error('请输入正确的密码'))
    } else {
      callback()
    }
  }

  // 获取验证码
  const loginVerify = async () => {
    const ele = await captcha()
    rules.captcha.push({
      max: ele.data.captchaLength,
      min: ele.data.captchaLength,
      message: `请输入${ele.data.captchaLength}位验证码`,
      trigger: 'blur'
    })
    picPath.value = ele.data.picPath
    loginFormData.captchaId = ele.data.captchaId
    loginFormData.openCaptcha = ele.data.openCaptcha
  }
  loginVerify()

  // 登录相关操作
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
    captcha: [
      {
        message: '验证码格式不正确',
        trigger: 'blur'
      }
    ]
  })

  const userStore = useUserStore()
  
  // Logo错误处理
  const handleLogoError = (event) => {
    console.log('Logo加载失败')
  }
  
  const login = async () => {
    return await userStore.LoginIn(loginFormData)
  }
  
  const submitForm = () => {
    loginForm.value.validate(async (v) => {
      if (!v) {
        // 未通过前端静态验证
        ElMessage({
          type: 'error',
          message: '请正确填写登录信息',
          showClose: true
        })
        await loginVerify()
        return false
      }

      // 通过验证，请求登陆
      const flag = await login()

      // 登陆失败，刷新验证码
      if (!flag) {
        await loginVerify()
        return false
      }

      // 登陆成功
      return true
    })
  }

  // 跳转初始化
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

<style scoped>
/* Logo样式 */
.logo-image {
  filter: drop-shadow(0 8px 20px rgba(102, 126, 234, 0.5));
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  animation: logoGlow 4s ease-in-out infinite alternate;
}

.logo-image:hover {
  transform: scale(1.05);
  filter: drop-shadow(0 12px 28px rgba(102, 126, 234, 0.7));
}

@keyframes logoGlow {
  0% {
    filter: drop-shadow(0 8px 20px rgba(102, 126, 234, 0.5));
  }
  100% {
    filter: drop-shadow(0 8px 20px rgba(118, 75, 162, 0.7));
  }
}

/* 现代化输入框样式 */
:deep(.modern-input .el-input__wrapper) {
  background: rgba(255, 255, 255, 0.1) !important;
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2) !important;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
}

:deep(.modern-input .el-input__wrapper:hover) {
  border-color: rgba(255, 255, 255, 0.4) !important;
  box-shadow: 0 6px 16px rgba(0, 0, 0, 0.15);
}

:deep(.modern-input .el-input__wrapper.is-focus) {
  border-color: rgba(102, 126, 234, 0.8) !important;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

:deep(.modern-input .el-input__inner) {
  color: white !important;
  font-weight: 500;
}

:deep(.modern-input .el-input__inner::placeholder) {
  color: rgba(255, 255, 255, 0.6) !important;
}

:deep(.modern-input .el-input__suffix) {
  color: rgba(255, 255, 255, 0.8) !important;
}

/* 表单验证样式 */
:deep(.el-form-item__error) {
  color: #ff6b6b !important;
  font-size: 12px;
  margin-top: 4px;
}

/* 按钮悬停效果 */
.el-button:hover {
  transform: translateY(-2px);
}

/* 背景动画 */
@keyframes float {
  0%, 100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-10px);
  }
}

.animate-float {
  animation: float 6s ease-in-out infinite;
}

/* 背景图案 */
.bg-pattern {
  background-image: radial-gradient(circle at 1px 1px, rgba(255,255,255,0.1) 1px, transparent 0);
  background-size: 40px 40px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .max-w-md {
    max-width: 90%;
  }
  
  .p-8 {
    padding: 1.5rem;
  }
}
</style>
