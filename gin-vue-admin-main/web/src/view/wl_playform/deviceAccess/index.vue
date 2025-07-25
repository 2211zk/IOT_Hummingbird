<template>
  <div>
    <!-- 页面标题和描述 -->
    <div class="mb-6">
      <div class="text-sm text-gray-500 mb-2">首页 / 设备接入</div>
      <h1 class="text-2xl font-bold text-gray-900 mb-2">设备接入</h1>
      <p class="text-gray-600">管理物联网平台中的产品和设备，支持产品创建、设备添加、批量导入等功能</p>
    </div>

    <!-- 功能卡片区域 -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <!-- 产品管理卡片 -->
      <div class="bg-white rounded-lg shadow-sm hover:shadow-md transition-shadow duration-200 border border-gray-100">
        <div class="p-6">
          <div class="flex items-center justify-between mb-4">
            <div class="flex items-center space-x-3">
              <div class="w-10 h-10 bg-blue-100 rounded-lg flex items-center justify-center">
                <el-icon class="text-blue-600 text-xl"><Box /></el-icon>
              </div>
              <div>
                <h3 class="text-lg font-semibold text-gray-900">产品管理</h3>
                <p class="text-sm text-gray-500">管理物联网产品</p>
              </div>
            </div>
            <el-button type="primary" size="small" @click="goToProducts">进入</el-button>
          </div>
          
          <div class="space-y-3">
            <div class="flex items-center justify-between text-sm">
              <span class="text-gray-600">产品总数</span>
              <span class="font-semibold text-gray-900">{{ productStats.total }}</span>
            </div>
            <div class="flex items-center justify-between text-sm">
              <span class="text-gray-600">在线产品</span>
              <span class="font-semibold text-green-600">{{ productStats.online }}</span>
            </div>
            <div class="flex items-center justify-between text-sm">
              <span class="text-gray-600">离线产品</span>
              <span class="font-semibold text-red-600">{{ productStats.offline }}</span>
            </div>
          </div>

          <div class="mt-4 pt-4 border-t border-gray-100">
            <div class="flex space-x-2">
              <el-button type="primary" link size="small" @click="createProduct">
                <el-icon style="margin-right: 4px"><Plus /></el-icon>创建产品
              </el-button>
              <el-button type="primary" link size="small" @click="viewProducts">
                <el-icon style="margin-right: 4px"><View /></el-icon>查看产品
              </el-button>
            </div>
          </div>
        </div>
      </div>

      <!-- 设备管理卡片 -->
      <div class="bg-white rounded-lg shadow-sm hover:shadow-md transition-shadow duration-200 border border-gray-100">
        <div class="p-6">
          <div class="flex items-center justify-between mb-4">
            <div class="flex items-center space-x-3">
              <div class="w-10 h-10 bg-green-100 rounded-lg flex items-center justify-center">
                <el-icon class="text-green-600 text-xl"><Monitor /></el-icon>
              </div>
              <div>
                <h3 class="text-lg font-semibold text-gray-900">设备管理</h3>
                <p class="text-sm text-gray-500">管理物联网设备</p>
              </div>
            </div>
            <el-button type="primary" size="small" @click="goToDevices">进入</el-button>
          </div>
          
          <div class="space-y-3">
            <div class="flex items-center justify-between text-sm">
              <span class="text-gray-600">设备总数</span>
              <span class="font-semibold text-gray-900">{{ deviceStats.total }}</span>
            </div>
            <div class="flex items-center justify-between text-sm">
              <span class="text-gray-600">在线设备</span>
              <span class="font-semibold text-green-600">{{ deviceStats.online }}</span>
            </div>
            <div class="flex items-center justify-between text-sm">
              <span class="text-gray-600">离线设备</span>
              <span class="font-semibold text-red-600">{{ deviceStats.offline }}</span>
            </div>
          </div>

          <div class="mt-4 pt-4 border-t border-gray-100">
            <div class="flex space-x-2">
              <el-button type="primary" link size="small" @click="createDevice">
                <el-icon style="margin-right: 4px"><Plus /></el-icon>添加设备
              </el-button>
              <el-button type="primary" link size="small" @click="viewDevices">
                <el-icon style="margin-right: 4px"><View /></el-icon>查看设备
              </el-button>
            </div>
          </div>
        </div>
      </div>

      <!-- 驱动管理卡片 -->
      <div class="bg-white rounded-lg shadow-sm hover:shadow-md transition-shadow duration-200 border border-gray-100">
        <div class="p-6">
          <div class="flex items-center justify-between mb-4">
            <div class="flex items-center space-x-3">
              <div class="w-10 h-10 bg-purple-100 rounded-lg flex items-center justify-center">
                <el-icon class="text-purple-600 text-xl"><Connection /></el-icon>
              </div>
              <div>
                <h3 class="text-lg font-semibold text-gray-900">驱动管理</h3>
                <p class="text-sm text-gray-500">管理设备驱动</p>
              </div>
            </div>
            <el-button type="primary" size="small" @click="goToDrivers">进入</el-button>
          </div>
          
          <div class="space-y-3">
            <div class="flex items-center justify-between text-sm">
              <span class="text-gray-600">驱动总数</span>
              <span class="font-semibold text-gray-900">{{ driverStats.total }}</span>
            </div>
            <div class="flex items-center justify-between text-sm">
              <span class="text-gray-600">已绑定</span>
              <span class="font-semibold text-green-600">{{ driverStats.bound }}</span>
            </div>
            <div class="flex items-center justify-between text-sm">
              <span class="text-gray-600">未绑定</span>
              <span class="font-semibold text-orange-600">{{ driverStats.unbound }}</span>
            </div>
          </div>

          <div class="mt-4 pt-4 border-t border-gray-100">
            <div class="flex space-x-2">
              <el-button type="primary" link size="small" @click="createDriver">
                <el-icon style="margin-right: 4px"><Plus /></el-icon>创建驱动
              </el-button>
              <el-button type="primary" link size="small" @click="viewDrivers">
                <el-icon style="margin-right: 4px"><View /></el-icon>查看驱动
              </el-button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 快速操作区域 -->
    <div class="mt-8">
      <h2 class="text-lg font-semibold text-gray-900 mb-4">快速操作</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
        <el-button type="primary" @click="createProduct" class="h-16">
          <el-icon style="margin-right: 8px"><Plus /></el-icon>
          创建产品
        </el-button>
        <el-button type="success" @click="createDevice" class="h-16">
          <el-icon style="margin-right: 8px"><Plus /></el-icon>
          添加设备
        </el-button>
        <el-button type="warning" @click="batchImport" class="h-16">
          <el-icon style="margin-right: 8px"><Upload /></el-icon>
          批量导入
        </el-button>
        <el-button type="info" @click="viewStatistics" class="h-16">
          <el-icon style="margin-right: 8px"><DataAnalysis /></el-icon>
          查看统计
        </el-button>
      </div>
    </div>

    <!-- 最近活动 -->
    <div class="mt-8">
      <h2 class="text-lg font-semibold text-gray-900 mb-4">最近活动</h2>
      <div class="bg-white rounded-lg shadow-sm border border-gray-100">
        <div class="p-4">
          <div v-if="recentActivities.length === 0" class="text-center text-gray-500 py-8">
            暂无最近活动
          </div>
          <div v-else class="space-y-3">
            <div v-for="activity in recentActivities" :key="activity.id" class="flex items-center space-x-3 p-3 hover:bg-gray-50 rounded">
              <div class="w-2 h-2 rounded-full" :class="getActivityColor(activity.type)"></div>
              <div class="flex-1">
                <div class="text-sm text-gray-900">{{ activity.message }}</div>
                <div class="text-xs text-gray-500">{{ formatDate(activity.time) }}</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { 
  Box, 
  Monitor, 
  Connection, 
  Plus, 
  View, 
  Upload, 
  DataAnalysis 
} from '@element-plus/icons-vue'
import { formatDate } from '@/utils/format'

defineOptions({
  name: 'DeviceAccess'
})

const router = useRouter()

// 统计数据
const productStats = ref({
  total: 0,
  online: 0,
  offline: 0
})

const deviceStats = ref({
  total: 0,
  online: 0,
  offline: 0
})

// 最近活动
const recentActivities = ref([])

// 获取活动颜色
const getActivityColor = (type) => {
  const colors = {
    'create': 'bg-green-500',
    'update': 'bg-blue-500',
    'delete': 'bg-red-500',
    'online': 'bg-green-500',
    'offline': 'bg-orange-500'
  }
  return colors[type] || 'bg-gray-500'
}

// 页面跳转函数 - 处理卡片中的"进入"按钮点击事件
const goToProducts = () => {
  router.push('/wl_playform/wlProducts')
}

const goToDevices = () => {
  router.push('/wl_playform/wlEquipment')
}

// 快速操作函数 - 处理页面底部快速操作按钮的点击事件
const createProduct = () => {
  router.push('/wl_playform/wlProducts')
}

const createDevice = () => {
  router.push('/wl_playform/wlEquipment')
}

const viewProducts = () => {
  router.push('/wl_playform/wlProducts')
}

const viewDevices = () => {
  router.push('/wl_playform/wlEquipment')
}

const batchImport = () => {
  router.push('/wl_playform/wlEquipment')
}

const viewStatistics = () => {
  ElMessage({
    type: 'info',
    message: '统计功能待实现'
  })
}

// 获取统计数据 - 模拟从后端API获取设备接入模块的统计数据
// 注意：当前使用模拟数据，实际项目中应调用真实的API接口
const getStats = async () => {
  // 这里应该调用实际的API获取统计数据
  productStats.value = {
    total: 12,
    online: 8,
    offline: 4
  }
  
  deviceStats.value = {
    total: 156,
    online: 89,
    offline: 67
  }
  
  // 移除 driverStats 相关内容
}

// 获取最近活动 - 模拟从后端API获取设备接入模块的最近活动记录
// 注意：当前使用模拟数据，实际项目中应调用真实的API接口
const getRecentActivities = async () => {
  // 这里应该调用实际的API获取最近活动
  recentActivities.value = [
    {
      id: 1,
      type: 'create',
      message: '创建了新设备：智能温湿度传感器',
      time: new Date(Date.now() - 1000 * 60 * 30) // 30分钟前
    },
    {
      id: 2,
      type: 'online',
      message: '设备 RTU-001 上线',
      time: new Date(Date.now() - 1000 * 60 * 60) // 1小时前
    },
    {
      id: 3,
      type: 'create',
      message: '创建了新产品：智能家居控制器',
      time: new Date(Date.now() - 1000 * 60 * 60 * 2) // 2小时前
    },
    {
      id: 4,
      type: 'offline',
      message: '设备 摄像头-002 离线',
      time: new Date(Date.now() - 1000 * 60 * 60 * 3) // 3小时前
    }
  ]
}

onMounted(() => {
  getStats()
  getRecentActivities()
})
</script>

<style scoped>
.el-button {
  font-weight: 500;
}
</style> 