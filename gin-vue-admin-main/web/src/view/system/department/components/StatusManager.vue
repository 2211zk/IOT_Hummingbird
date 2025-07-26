<template>
  <div class="status-manager">
    <div class="manager-header">
      <h4>状态管理</h4>
      <el-button
        type="text"
        size="small"
        @click="refreshData"
        :loading="loading"
        icon="Refresh"
      >
        刷新
      </el-button>
    </div>
    
    <div class="manager-content">
      <!-- 批量操作 -->
      <div class="batch-operations">
        <div class="operation-header">
          <span>批量操作</span>
        </div>
        <div class="operation-buttons">
          <el-button
            type="success"
            size="small"
            @click="handleBatchEnable"
            :disabled="selectedDepartments.length === 0"
            :loading="batchLoading"
          >
            批量启用 ({{ selectedDepartments.length }})
          </el-button>
          <el-button
            type="warning"
            size="small"
            @click="handleBatchDisable"
            :disabled="selectedDepartments.length === 0"
            :loading="batchLoading"
          >
            批量禁用 ({{ selectedDepartments.length }})
          </el-button>
        </div>
      </div>
      
      <!-- 状态统计 -->
      <div class="status-stats">
        <div class="stats-header">
          <span>状态统计</span>
        </div>
        <div class="stats-items">
          <div class="stat-item">
            <div class="stat-icon enabled">
              <el-icon><CircleCheck /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ stats.enabled }}</div>
              <div class="stat-label">启用</div>
            </div>
          </div>
          <div class="stat-item">
            <div class="stat-icon disabled">
              <el-icon><CircleClose /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ stats.disabled }}</div>
              <div class="stat-label">禁用</div>
            </div>
          </div>
        </div>
        <div class="stats-progress">
          <div class="progress-label">
            <span>启用率</span>
            <span>{{ enabledRate }}%</span>
          </div>
          <el-progress
            :percentage="enabledRate"
            :color="getProgressColor(enabledRate)"
            :stroke-width="6"
          />
        </div>
      </div>
      
      <!-- 部门列表 -->
      <div class="department-list">
        <div class="list-header">
          <span>部门列表</span>
          <el-button
            type="text"
            size="small"
            @click="selectAll"
            v-if="!allSelected"
          >
            全选
          </el-button>
          <el-button
            type="text"
            size="small"
            @click="clearSelection"
            v-else
          >
            取消全选
          </el-button>
        </div>
        <div class="list-content">
          <div
            v-for="department in departments"
            :key="department.id"
            class="department-item"
            :class="{ selected: selectedDepartments.includes(department.id) }"
            @click="toggleSelection(department.id)"
          >
            <div class="department-info">
              <div class="department-name">{{ department.name }}</div>
              <div class="department-path">{{ getDepartmentPath(department) }}</div>
            </div>
            <div class="department-status">
              <el-switch
                v-model="department.status"
                active-text="启用"
                inactive-text="禁用"
                active-value="启用"
                inactive-value="禁用"
                @change="handleStatusChange(department)"
                @click.stop
                size="small"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { CircleCheck, CircleClose } from '@element-plus/icons-vue'
import { getDepartmentList, updateDepartment } from '@/api/wlDepartment'

// Props
const props = defineProps({
  modelValue: {
    type: Array,
    default: () => []
  }
})

// Emits
const emit = defineEmits(['update:modelValue', 'change'])

// 响应式数据
const loading = ref(false)
const batchLoading = ref(false)
const departments = ref([])
const selectedDepartments = ref([])

// 统计数据
const stats = reactive({
  enabled: 0,
  disabled: 0,
  total: 0
})

// 计算属性
const enabledRate = computed(() => {
  if (stats.total === 0) return 0
  return Math.round((stats.enabled / stats.total) * 100)
})

const allSelected = computed(() => {
  return departments.value.length > 0 && selectedDepartments.value.length === departments.value.length
})

// 获取数据
const getData = async () => {
  try {
    loading.value = true
    const response = await getDepartmentList({
      page: 1,
      pageSize: 1000,
      treeMode: false // 使用平铺模式获取所有部门
    })
    
    if (response.code === 0) {
      departments.value = response.data.list || []
      updateStats()
    } else {
      ElMessage.error(response.msg || '获取部门列表失败')
    }
  } catch (error) {
    console.error('获取部门列表失败:', error)
    ElMessage.error('获取部门列表失败')
  } finally {
    loading.value = false
  }
}

// 更新统计
const updateStats = () => {
  stats.total = departments.value.length
  stats.enabled = departments.value.filter(d => d.status === '启用').length
  stats.disabled = departments.value.filter(d => d.status === '禁用').length
}

// 刷新数据
const refreshData = () => {
  getData()
}

// 选择切换
const toggleSelection = (departmentId) => {
  const index = selectedDepartments.value.indexOf(departmentId)
  if (index > -1) {
    selectedDepartments.value.splice(index, 1)
  } else {
    selectedDepartments.value.push(departmentId)
  }
}

// 全选
const selectAll = () => {
  selectedDepartments.value = departments.value.map(d => d.id)
}

// 清除选择
const clearSelection = () => {
  selectedDepartments.value = []
}

// 批量启用
const handleBatchEnable = async () => {
  try {
    await ElMessageBox.confirm(
      `确定要启用选中的 ${selectedDepartments.value.length} 个部门吗？`,
      '批量启用确认',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'info'
      }
    )
    
    await batchUpdateStatus('启用')
  } catch (error) {
    if (error !== 'cancel') {
      console.error('批量启用失败:', error)
    }
  }
}

// 批量禁用
const handleBatchDisable = async () => {
  try {
    await ElMessageBox.confirm(
      `确定要禁用选中的 ${selectedDepartments.value.length} 个部门吗？\n注意：禁用部门会同时禁用其所有子部门。`,
      '批量禁用确认',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    
    await batchUpdateStatus('禁用')
  } catch (error) {
    if (error !== 'cancel') {
      console.error('批量禁用失败:', error)
    }
  }
}

// 批量更新状态
const batchUpdateStatus = async (status) => {
  try {
    batchLoading.value = true
    let successCount = 0
    let failCount = 0
    
    for (const departmentId of selectedDepartments.value) {
      const department = departments.value.find(d => d.id === departmentId)
      if (!department) continue
      
      try {
        const response = await updateDepartment({
          id: department.id,
          name: department.name,
          parentId: department.parentId,
          leader: department.leader,
          phone: department.phone,
          email: department.email,
          status: status,
          sort: department.sort,
          deviceIds: department.deviceIds || []
        })
        
        if (response.code === 0) {
          department.status = status
          successCount++
        } else {
          failCount++
        }
      } catch (error) {
        failCount++
      }
    }
    
    if (successCount > 0) {
      ElMessage.success(`成功${status} ${successCount} 个部门`)
      updateStats()
      emit('change')
    }
    
    if (failCount > 0) {
      ElMessage.warning(`${failCount} 个部门${status}失败`)
    }
    
    clearSelection()
  } catch (error) {
    console.error('批量更新状态失败:', error)
    ElMessage.error('批量更新状态失败')
  } finally {
    batchLoading.value = false
  }
}

// 单个状态切换
const handleStatusChange = async (department) => {
  try {
    const response = await updateDepartment({
      id: department.id,
      name: department.name,
      parentId: department.parentId,
      leader: department.leader,
      phone: department.phone,
      email: department.email,
      status: department.status,
      sort: department.sort,
      deviceIds: department.deviceIds || []
    })
    
    if (response.code === 0) {
      ElMessage.success(`部门已${department.status}`)
      updateStats()
      emit('change')
    } else {
      // 恢复原状态
      department.status = department.status === '启用' ? '禁用' : '启用'
      ElMessage.error(response.msg || '状态更新失败')
    }
  } catch (error) {
    // 恢复原状态
    department.status = department.status === '启用' ? '禁用' : '启用'
    console.error('状态更新失败:', error)
    ElMessage.error('状态更新失败')
  }
}

// 获取部门路径
const getDepartmentPath = (department) => {
  // 这里可以实现部门路径的获取逻辑
  // 暂时返回简单的父部门信息
  return department.parentId ? `上级部门ID: ${department.parentId}` : '顶级部门'
}

// 获取进度条颜色
const getProgressColor = (percentage) => {
  if (percentage >= 80) return '#67c23a'
  if (percentage >= 60) return '#e6a23c'
  if (percentage >= 40) return '#f56c6c'
  return '#909399'
}

// 组件挂载时获取数据
onMounted(() => {
  getData()
})
</script>

<style scoped>
.status-manager {
  border: 1px solid #e4e7ed;
  border-radius: 4px;
  overflow: hidden;
}

.manager-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 20px;
  background: #fafafa;
  border-bottom: 1px solid #e4e7ed;
}

.manager-header h4 {
  margin: 0;
  font-size: 16px;
  font-weight: 600;
  color: #303133;
}

.manager-content {
  padding: 20px;
}

.batch-operations,
.status-stats,
.department-list {
  margin-bottom: 24px;
}

.operation-header,
.stats-header,
.list-header {
  font-size: 14px;
  color: #606266;
  font-weight: 500;
  margin-bottom: 12px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.operation-buttons {
  display: flex;
  gap: 12px;
}

.stats-items {
  display: flex;
  gap: 20px;
  margin-bottom: 16px;
}

.stat-item {
  display: flex;
  align-items: center;
  gap: 8px;
}

.stat-icon {
  width: 24px;
  height: 24px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
  color: white;
}

.stat-icon.enabled {
  background: #67c23a;
}

.stat-icon.disabled {
  background: #f56c6c;
}

.stat-info {
  display: flex;
  flex-direction: column;
}

.stat-value {
  font-size: 18px;
  font-weight: 600;
  color: #303133;
  line-height: 1;
}

.stat-label {
  font-size: 12px;
  color: #909399;
}

.stats-progress {
  margin-top: 12px;
}

.progress-label {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
  font-size: 14px;
  color: #606266;
}

.list-content {
  max-height: 300px;
  overflow-y: auto;
}

.department-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px;
  border: 1px solid #e4e7ed;
  border-radius: 4px;
  margin-bottom: 8px;
  cursor: pointer;
  transition: all 0.2s;
}

.department-item:hover {
  border-color: #409eff;
  background-color: #f0f9ff;
}

.department-item.selected {
  border-color: #409eff;
  background-color: #e6f7ff;
}

.department-info {
  flex: 1;
}

.department-name {
  font-size: 14px;
  font-weight: 500;
  color: #303133;
  margin-bottom: 4px;
}

.department-path {
  font-size: 12px;
  color: #909399;
}

.department-status {
  margin-left: 12px;
}

:deep(.el-progress-bar__outer) {
  border-radius: 3px;
}

:deep(.el-progress-bar__inner) {
  border-radius: 3px;
}
</style>