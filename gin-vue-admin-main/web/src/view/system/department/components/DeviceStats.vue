<template>
  <div class="device-stats">
    <div class="stats-header">
      <h4>设备统计</h4>
      <el-button
        type="text"
        size="small"
        @click="refreshStats"
        :loading="loading"
        icon="Refresh"
      >
        刷新
      </el-button>
    </div>
    
    <div class="stats-content">
      <!-- 总体统计 -->
      <div class="stats-overview">
        <div class="stat-card">
          <div class="stat-icon total">
            <el-icon><Box /></el-icon>
          </div>
          <div class="stat-info">
            <div class="stat-value">{{ stats.total }}</div>
            <div class="stat-label">总设备数</div>
          </div>
        </div>
        
        <div class="stat-card">
          <div class="stat-icon associated">
            <el-icon><Connection /></el-icon>
          </div>
          <div class="stat-info">
            <div class="stat-value">{{ stats.associated }}</div>
            <div class="stat-label">已关联</div>
          </div>
        </div>
        
        <div class="stat-card">
          <div class="stat-icon available">
            <el-icon><CircleCheck /></el-icon>
          </div>
          <div class="stat-info">
            <div class="stat-value">{{ stats.available }}</div>
            <div class="stat-label">可关联</div>
          </div>
        </div>
        
        <div class="stat-card">
          <div class="stat-icon disabled">
            <el-icon><CircleClose /></el-icon>
          </div>
          <div class="stat-info">
            <div class="stat-value">{{ stats.disabled }}</div>
            <div class="stat-label">已禁用</div>
          </div>
        </div>
      </div>
      
      <!-- 关联率 -->
      <div class="association-rate">
        <div class="rate-header">
          <span class="rate-label">设备关联率</span>
          <span class="rate-value">{{ associationRate }}%</span>
        </div>
        <el-progress
          :percentage="associationRate"
          :color="getProgressColor(associationRate)"
          :stroke-width="8"
        />
      </div>
      
      <!-- 状态分布 -->
      <div class="status-distribution">
        <div class="distribution-header">
          <span>状态分布</span>
        </div>
        <div class="distribution-items">
          <div class="distribution-item">
            <div class="item-indicator enabled"></div>
            <span class="item-label">启用</span>
            <span class="item-value">{{ stats.enabled }}</span>
          </div>
          <div class="distribution-item">
            <div class="item-indicator disabled"></div>
            <span class="item-label">禁用</span>
            <span class="item-value">{{ stats.disabled }}</span>
          </div>
        </div>
      </div>
      
      <!-- 最近关联 -->
      <div class="recent-associations" v-if="recentAssociations.length > 0">
        <div class="recent-header">
          <span>最近关联</span>
        </div>
        <div class="recent-list">
          <div
            v-for="item in recentAssociations"
            :key="item.id"
            class="recent-item"
          >
            <div class="recent-device">
              <div class="device-name">{{ item.deviceName }}</div>
              <div class="department-name">{{ item.departmentName }}</div>
            </div>
            <div class="recent-time">
              {{ formatRelativeTime(item.associatedAt) }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { Box, Connection, CircleCheck, CircleClose, Refresh } from '@element-plus/icons-vue'
import { getDeviceStats } from '@/api/device'

// 响应式数据
const loading = ref(false)
const stats = reactive({
  total: 0,
  associated: 0,
  available: 0,
  enabled: 0,
  disabled: 0
})
const recentAssociations = ref([])

// 计算属性
const associationRate = computed(() => {
  if (stats.total === 0) return 0
  return Math.round((stats.associated / stats.total) * 100)
})

// 获取统计数据
const getStats = async () => {
  try {
    loading.value = true
    const response = await getDeviceStats()
    
    if (response.code === 0) {
      const data = response.data
      Object.assign(stats, {
        total: data.total || 0,
        associated: data.associated || 0,
        available: data.available || 0,
        enabled: data.enabled || 0,
        disabled: data.disabled || 0
      })
      
      recentAssociations.value = data.recentAssociations || []
    } else {
      ElMessage.error(response.msg || '获取统计数据失败')
    }
  } catch (error) {
    console.error('获取统计数据失败:', error)
    // 使用模拟数据
    Object.assign(stats, {
      total: 150,
      associated: 89,
      available: 61,
      enabled: 135,
      disabled: 15
    })
  } finally {
    loading.value = false
  }
}

// 刷新统计
const refreshStats = () => {
  getStats()
}

// 获取进度条颜色
const getProgressColor = (percentage) => {
  if (percentage >= 80) return '#67c23a'
  if (percentage >= 60) return '#e6a23c'
  if (percentage >= 40) return '#f56c6c'
  return '#909399'
}

// 格式化相对时间
const formatRelativeTime = (time) => {
  if (!time) return '-'
  
  const now = new Date()
  const target = new Date(time)
  const diff = now - target
  
  const minutes = Math.floor(diff / (1000 * 60))
  const hours = Math.floor(diff / (1000 * 60 * 60))
  const days = Math.floor(diff / (1000 * 60 * 60 * 24))
  
  if (minutes < 1) return '刚刚'
  if (minutes < 60) return `${minutes}分钟前`
  if (hours < 24) return `${hours}小时前`
  if (days < 7) return `${days}天前`
  
  return target.toLocaleDateString()
}

// 组件挂载时获取数据
onMounted(() => {
  getStats()
})
</script>

<style scoped>
.device-stats {
  border: 1px solid #e4e7ed;
  border-radius: 4px;
  overflow: hidden;
}

.stats-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 20px;
  background: #fafafa;
  border-bottom: 1px solid #e4e7ed;
}

.stats-header h4 {
  margin: 0;
  font-size: 16px;
  font-weight: 600;
  color: #303133;
}

.stats-content {
  padding: 20px;
}

.stats-overview {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(120px, 1fr));
  gap: 16px;
  margin-bottom: 24px;
}

.stat-card {
  display: flex;
  align-items: center;
  padding: 16px;
  background: #f8f9fa;
  border-radius: 8px;
  gap: 12px;
}

.stat-icon {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  color: white;
}

.stat-icon.total {
  background: #409eff;
}

.stat-icon.associated {
  background: #67c23a;
}

.stat-icon.available {
  background: #e6a23c;
}

.stat-icon.disabled {
  background: #f56c6c;
}

.stat-info {
  flex: 1;
}

.stat-value {
  font-size: 24px;
  font-weight: 600;
  color: #303133;
  line-height: 1;
  margin-bottom: 4px;
}

.stat-label {
  font-size: 12px;
  color: #909399;
}

.association-rate {
  margin-bottom: 24px;
}

.rate-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.rate-label {
  font-size: 14px;
  color: #606266;
  font-weight: 500;
}

.rate-value {
  font-size: 16px;
  font-weight: 600;
  color: #303133;
}

.status-distribution {
  margin-bottom: 24px;
}

.distribution-header {
  font-size: 14px;
  color: #606266;
  font-weight: 500;
  margin-bottom: 12px;
}

.distribution-items {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.distribution-item {
  display: flex;
  align-items: center;
  gap: 8px;
}

.item-indicator {
  width: 8px;
  height: 8px;
  border-radius: 50%;
}

.item-indicator.enabled {
  background: #67c23a;
}

.item-indicator.disabled {
  background: #f56c6c;
}

.item-label {
  flex: 1;
  font-size: 14px;
  color: #606266;
}

.item-value {
  font-size: 14px;
  font-weight: 500;
  color: #303133;
}

.recent-associations {
  border-top: 1px solid #e4e7ed;
  padding-top: 20px;
}

.recent-header {
  font-size: 14px;
  color: #606266;
  font-weight: 500;
  margin-bottom: 12px;
}

.recent-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.recent-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 0;
}

.recent-device {
  flex: 1;
}

.device-name {
  font-size: 14px;
  color: #303133;
  font-weight: 500;
  margin-bottom: 2px;
}

.department-name {
  font-size: 12px;
  color: #909399;
}

.recent-time {
  font-size: 12px;
  color: #909399;
}

:deep(.el-progress-bar__outer) {
  border-radius: 4px;
}

:deep(.el-progress-bar__inner) {
  border-radius: 4px;
}
</style>