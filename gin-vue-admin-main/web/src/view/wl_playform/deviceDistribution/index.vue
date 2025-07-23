<template>
  <div>
    <!-- 页面标题和描述 -->
    <div class="mb-6">
      <div class="text-sm text-gray-500 mb-2">首页 / 设备分布</div>
      <h1 class="text-2xl font-bold text-gray-900 mb-2">设备分布</h1>
      <p class="text-gray-600">查看和管理设备的分布情况</p>
    </div>

    <!-- 统计卡片 -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <el-card class="stat-card">
        <div class="flex items-center">
          <el-icon class="text-blue-500 text-2xl mr-3">
            <Monitor />
          </el-icon>
          <div>
            <div class="text-sm text-gray-500">总设备数</div>
            <div class="text-2xl font-bold">{{ deviceStats.total || 0 }}</div>
          </div>
        </div>
      </el-card>
      
      <el-card class="stat-card">
        <div class="flex items-center">
          <el-icon class="text-green-500 text-2xl mr-3">
            <CircleCheck />
          </el-icon>
          <div>
            <div class="text-sm text-gray-500">在线设备</div>
            <div class="text-2xl font-bold">{{ deviceStats.online || 0 }}</div>
          </div>
        </div>
      </el-card>
      
      <el-card class="stat-card">
        <div class="flex items-center">
          <el-icon class="text-red-500 text-2xl mr-3">
            <CircleClose />
          </el-icon>
          <div>
            <div class="text-sm text-gray-500">离线设备</div>
            <div class="text-2xl font-bold">{{ deviceStats.offline || 0 }}</div>
          </div>
        </div>
      </el-card>
      
      <el-card class="stat-card">
        <div class="flex items-center">
          <el-icon class="text-purple-500 text-2xl mr-3">
            <Location />
          </el-icon>
          <div>
            <div class="text-sm text-gray-500">分布区域</div>
            <div class="text-2xl font-bold">{{ deviceStats.regions || 0 }}</div>
          </div>
        </div>
      </el-card>
    </div>

    <!-- 功能区域 -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <!-- 设备地图卡片 -->
      <el-card class="feature-card">
        <template #header>
          <div class="flex items-center">
            <el-icon class="text-blue-500 mr-2">
              <Location />
            </el-icon>
            <span>设备地图</span>
          </div>
        </template>
        <div class="text-gray-600 mb-4">
          在地图上查看所有设备的分布情况和位置信息
        </div>
        <el-button type="primary" @click="goToDeviceMap">
          查看设备地图
        </el-button>
      </el-card>

      <!-- 设备统计卡片 -->
      <el-card class="feature-card">
        <template #header>
          <div class="flex items-center">
            <el-icon class="text-green-500 mr-2">
              <PieChart />
            </el-icon>
            <span>设备统计</span>
          </div>
        </template>
        <div class="text-gray-600 mb-4">
          查看设备的详细统计信息和分析报告
        </div>
        <el-button type="success" @click="goToDeviceStats">
          查看统计报告
        </el-button>
      </el-card>
    </div>

    <!-- 最近设备列表 -->
    <el-card class="mt-6">
      <template #header>
        <div class="flex items-center justify-between">
          <span>最近设备</span>
          <el-button type="primary" link @click="goToDeviceManagement">
            查看全部
          </el-button>
        </div>
      </template>
      
      <el-table :data="recentDevices" style="width: 100%">
        <el-table-column prop="ID" label="设备ID" width="100" />
        <el-table-column prop="eqName" label="设备名称" width="150" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.status === 'online' ? 'success' : 'danger'">
              {{ scope.row.status === 'online' ? '在线' : '离线' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="eqCoordinate" label="坐标" width="200" />
        <el-table-column prop="CreatedAt" label="创建时间" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.CreatedAt) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150">
          <template #default="scope">
            <el-button type="primary" link @click="viewDeviceDetail(scope.row)">
              详情
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { Monitor, CircleCheck, CircleClose, Location, PieChart } from '@element-plus/icons-vue'
import { getDeviceMapData } from '@/api/wl_playform/deviceMap'
import { formatDate } from '@/utils/format'

defineOptions({
  name: 'DeviceDistribution'
})

const router = useRouter()

// 响应式数据
const deviceStats = ref({
  total: 0,
  online: 0,
  offline: 0,
  regions: 0
})

const recentDevices = ref([])

// 加载设备统计数据
const loadDeviceStats = async () => {
  try {
    const response = await getDeviceMapData({ page: 1, pageSize: 1000 })
    
    if (response.code === 0) {
      const devices = response.data.list || []
      
      // 计算统计数据
      deviceStats.value = {
        total: devices.length,
        online: devices.filter(d => d.status === 'online').length,
        offline: devices.filter(d => d.status === 'offline').length,
        regions: new Set(devices.filter(d => d.eqCoordinate).map(d => {
          const coords = d.eqCoordinate.split(',')
          return `${Math.floor(coords[0])},${Math.floor(coords[1])}`
        })).size
      }
      
      // 获取最近设备（前10个）
      recentDevices.value = devices.slice(0, 10)
    }
  } catch (error) {
    console.error('加载设备统计失败:', error)
    ElMessage.error('加载设备统计失败')
  }
}

// 跳转到设备地图
const goToDeviceMap = () => {
  router.push({ name: 'deviceMap' })
}

// 跳转到设备统计
const goToDeviceStats = () => {
  ElMessage.info('设备统计功能开发中...')
}

// 跳转到设备管理
const goToDeviceManagement = () => {
  router.push({ name: 'wlEquipment' })
}

// 查看设备详情
const viewDeviceDetail = (device) => {
  ElMessage.info(`查看设备详情: ${device.eqName}`)
}

// 组件挂载
onMounted(() => {
  loadDeviceStats()
})
</script>

<style scoped>
.stat-card {
  transition: all 0.3s ease;
}

.stat-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.feature-card {
  transition: all 0.3s ease;
}

.feature-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}
</style> 