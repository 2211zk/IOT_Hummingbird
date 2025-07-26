<template>
  <div class="device-association">
    <!-- 关联设备列表 -->
    <div class="association-header">
      <h4>关联设备管理</h4>
      <el-button
        type="primary"
        size="small"
        @click="showSelector = true"
        icon="Plus"
      >
        添加设备
      </el-button>
    </div>
    
    <div class="association-content">
      <!-- 搜索区域 -->
      <div class="search-bar">
        <el-input
          v-model="searchKeyword"
          placeholder="搜索设备名称或产品名称"
          clearable
          style="width: 300px"
          @input="handleSearch"
        />
      </div>
      
      <!-- 设备列表 -->
      <div class="device-list">
        <div v-if="filteredDevices.length === 0" class="empty-devices">
          <el-empty description="暂无关联设备" :image-size="80">
            <el-button type="primary" @click="showSelector = true">
              添加设备
            </el-button>
          </el-empty>
        </div>
        
        <div v-else class="device-grid">
          <div
            v-for="device in paginatedDevices"
            :key="device.id"
            class="device-card"
          >
            <div class="device-header">
              <div class="device-info">
                <div class="device-name">{{ device.deviceName }}</div>
                <div class="product-name">{{ device.productName || '-' }}</div>
              </div>
              <div class="device-actions">
                <el-tag :type="getStatusType(device.status)" size="small">
                  {{ device.status }}
                </el-tag>
                <el-button
                  type="danger"
                  size="small"
                  text
                  @click="handleRemoveDevice(device)"
                  icon="Close"
                />
              </div>
            </div>
            
            <div class="device-details">
              <div class="detail-item">
                <span class="label">设备ID:</span>
                <span class="value">{{ device.id }}</span>
              </div>
              <div class="detail-item">
                <span class="label">关联时间:</span>
                <span class="value">{{ formatDateTime(device.associatedAt) }}</span>
              </div>
            </div>
          </div>
        </div>
        
        <!-- 分页 -->
        <div v-if="filteredDevices.length > pageSize" class="pagination-container">
          <el-pagination
            v-model:current-page="currentPage"
            :page-size="pageSize"
            :total="filteredDevices.length"
            layout="prev, pager, next"
            @current-change="handlePageChange"
          />
        </div>
      </div>
    </div>
    
    <!-- 设备选择器 -->
    <DeviceSelector
      v-model="showSelector"
      :selected-devices="associatedDevices"
      :department-id="departmentId"
      @confirm="handleDeviceSelect"
    />
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getDepartmentDevices, updateDepartment } from '@/api/wlDepartment'
import { formatDateTime } from '@/utils/format'
import DeviceSelector from './DeviceSelector.vue'

// Props
const props = defineProps({
  departmentId: {
    type: Number,
    required: true
  },
  modelValue: {
    type: Array,
    default: () => []
  }
})

// Emits
const emit = defineEmits(['update:modelValue', 'change'])

// 响应式数据
const showSelector = ref(false)
const searchKeyword = ref('')
const currentPage = ref(1)
const pageSize = ref(12)
const associatedDevices = ref([])

// 计算属性
const filteredDevices = computed(() => {
  if (!searchKeyword.value) {
    return associatedDevices.value
  }
  
  const keyword = searchKeyword.value.toLowerCase()
  return associatedDevices.value.filter(device => 
    device.deviceName.toLowerCase().includes(keyword) ||
    (device.productName && device.productName.toLowerCase().includes(keyword))
  )
})

const paginatedDevices = computed(() => {
  const start = (currentPage.value - 1) * pageSize.value
  const end = start + pageSize.value
  return filteredDevices.value.slice(start, end)
})

// 监听props变化
watch(() => props.modelValue, (newVal) => {
  associatedDevices.value = [...newVal]
}, { immediate: true })

watch(() => props.departmentId, (newVal) => {
  if (newVal) {
    loadDepartmentDevices()
  }
}, { immediate: true })

// 加载部门设备
const loadDepartmentDevices = async () => {
  if (!props.departmentId) return
  
  try {
    const response = await getDepartmentDevices({
      departmentId: props.departmentId,
      page: 1,
      pageSize: 1000 // 获取所有设备
    })
    
    if (response.code === 0) {
      associatedDevices.value = response.data.list || []
      emit('update:modelValue', associatedDevices.value)
    } else {
      ElMessage.error(response.msg || '获取关联设备失败')
    }
  } catch (error) {
    console.error('获取关联设备失败:', error)
    ElMessage.error('获取关联设备失败')
  }
}

// 搜索处理
const handleSearch = () => {
  currentPage.value = 1
}

// 分页处理
const handlePageChange = (page) => {
  currentPage.value = page
}

// 设备选择处理
const handleDeviceSelect = async (selectedDevices) => {
  try {
    // 更新部门的设备关联
    const deviceIds = selectedDevices.map(device => device.id)
    const response = await updateDepartment({
      id: props.departmentId,
      deviceIds
    })
    
    if (response.code === 0) {
      associatedDevices.value = selectedDevices
      emit('update:modelValue', associatedDevices.value)
      emit('change', selectedDevices)
      ElMessage.success('设备关联更新成功')
    } else {
      ElMessage.error(response.msg || '更新设备关联失败')
    }
  } catch (error) {
    console.error('更新设备关联失败:', error)
    ElMessage.error('更新设备关联失败')
  }
}

// 移除设备
const handleRemoveDevice = async (device) => {
  try {
    await ElMessageBox.confirm(
      `确定要移除设备"${device.deviceName}"的关联吗？`,
      '移除确认',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    
    const remainingDevices = associatedDevices.value.filter(d => d.id !== device.id)
    const deviceIds = remainingDevices.map(d => d.id)
    
    const response = await updateDepartment({
      id: props.departmentId,
      deviceIds
    })
    
    if (response.code === 0) {
      associatedDevices.value = remainingDevices
      emit('update:modelValue', associatedDevices.value)
      emit('change', remainingDevices)
      ElMessage.success('设备移除成功')
    } else {
      ElMessage.error(response.msg || '移除设备失败')
    }
  } catch (error) {
    if (error !== 'cancel') {
      console.error('移除设备失败:', error)
      ElMessage.error('移除设备失败')
    }
  }
}

// 获取状态类型
const getStatusType = (status) => {
  switch (status) {
    case '启用':
      return 'success'
    case '禁用':
      return 'danger'
    default:
      return 'info'
  }
}
</script>

<style scoped>
.device-association {
  border: 1px solid #e4e7ed;
  border-radius: 4px;
  overflow: hidden;
}

.association-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 20px;
  background: #fafafa;
  border-bottom: 1px solid #e4e7ed;
}

.association-header h4 {
  margin: 0;
  font-size: 16px;
  font-weight: 600;
  color: #303133;
}

.association-content {
  padding: 20px;
}

.search-bar {
  margin-bottom: 20px;
}

.empty-devices {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 200px;
}

.device-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 16px;
  margin-bottom: 20px;
}

.device-card {
  border: 1px solid #e4e7ed;
  border-radius: 8px;
  padding: 16px;
  background: white;
  transition: all 0.2s;
}

.device-card:hover {
  border-color: #409eff;
  box-shadow: 0 2px 8px rgba(64, 158, 255, 0.1);
}

.device-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 12px;
}

.device-info {
  flex: 1;
}

.device-name {
  font-size: 16px;
  font-weight: 600;
  color: #303133;
  margin-bottom: 4px;
}

.product-name {
  font-size: 14px;
  color: #909399;
}

.device-actions {
  display: flex;
  align-items: center;
  gap: 8px;
}

.device-details {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.detail-item {
  display: flex;
  justify-content: space-between;
  font-size: 14px;
}

.label {
  color: #909399;
  font-weight: 500;
}

.value {
  color: #606266;
}

.pagination-container {
  display: flex;
  justify-content: center;
  margin-top: 20px;
}

:deep(.el-empty) {
  padding: 40px 0;
}
</style>