<template>
  <div class="sort-manager">
    <div class="manager-header">
      <h4>排序管理</h4>
      <div class="header-actions">
        <el-button
          type="text"
          size="small"
          @click="autoSort"
          :loading="autoSortLoading"
        >
          自动排序
        </el-button>
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
    </div>
    
    <div class="manager-content">
      <!-- 排序说明 -->
      <div class="sort-info">
        <el-alert
          title="排序说明"
          type="info"
          :closable="false"
          show-icon
        >
          <template #default>
            <p>• 排序值越小，显示位置越靠前</p>
            <p>• 同级部门按排序值升序排列</p>
            <p>• 排序值相同时按创建时间排序</p>
          </template>
        </el-alert>
      </div>
      
      <!-- 批量排序 -->
      <div class="batch-sort">
        <div class="batch-header">
          <span>批量排序</span>
        </div>
        <div class="batch-controls">
          <el-input-number
            v-model="batchSortValue"
            :min="0"
            :max="9999"
            placeholder="排序值"
            style="width: 120px"
          />
          <el-button
            type="primary"
            size="small"
            @click="handleBatchSort"
            :disabled="selectedDepartments.length === 0"
            :loading="batchLoading"
          >
            批量设置 ({{ selectedDepartments.length }})
          </el-button>
          <el-button
            type="success"
            size="small"
            @click="handleSequentialSort"
            :disabled="selectedDepartments.length === 0"
            :loading="batchLoading"
          >
            顺序排序
          </el-button>
        </div>
      </div>
      
      <!-- 部门列表 -->
      <div class="department-list">
        <div class="list-header">
          <span>部门列表</span>
          <div class="header-controls">
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
        </div>
        
        <div class="list-content">
          <draggable
            v-model="departments"
            @end="handleDragEnd"
            item-key="id"
            class="drag-list"
          >
            <template #item="{ element: department, index }">
              <div
                class="department-item"
                :class="{ selected: selectedDepartments.includes(department.id) }"
                @click="toggleSelection(department.id)"
              >
                <div class="drag-handle">
                  <el-icon><Rank /></el-icon>
                </div>
                
                <div class="department-info">
                  <div class="department-name">{{ department.name }}</div>
                  <div class="department-level">{{ getDepartmentLevel(department) }}</div>
                </div>
                
                <div class="sort-controls">
                  <el-input-number
                    v-model="department.sort"
                    :min="0"
                    :max="9999"
                    size="small"
                    controls-position="right"
                    @change="handleSortChange(department)"
                    @click.stop
                    style="width: 100px"
                  />
                </div>
                
                <div class="quick-actions">
                  <el-button
                    type="text"
                    size="small"
                    @click.stop="moveUp(index)"
                    :disabled="index === 0"
                    icon="ArrowUp"
                  />
                  <el-button
                    type="text"
                    size="small"
                    @click.stop="moveDown(index)"
                    :disabled="index === departments.length - 1"
                    icon="ArrowDown"
                  />
                </div>
              </div>
            </template>
          </draggable>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Rank, ArrowUp, ArrowDown } from '@element-plus/icons-vue'
import draggable from 'vuedraggable'
import { getDepartmentList, updateDepartment } from '@/api/wlDepartment'

// 响应式数据
const loading = ref(false)
const batchLoading = ref(false)
const autoSortLoading = ref(false)
const departments = ref([])
const selectedDepartments = ref([])
const batchSortValue = ref(0)

// 计算属性
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
      treeMode: false // 使用平铺模式
    })
    
    if (response.code === 0) {
      // 按当前排序值排序
      departments.value = (response.data.list || []).sort((a, b) => {
        if (a.sort !== b.sort) {
          return (a.sort || 0) - (b.sort || 0)
        }
        return new Date(a.createdAt) - new Date(b.createdAt)
      })
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

// 批量排序
const handleBatchSort = async () => {
  try {
    await ElMessageBox.confirm(
      `确定要将选中的 ${selectedDepartments.value.length} 个部门的排序值设置为 ${batchSortValue.value} 吗？`,
      '批量排序确认',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'info'
      }
    )
    
    await batchUpdateSort(batchSortValue.value)
  } catch (error) {
    if (error !== 'cancel') {
      console.error('批量排序失败:', error)
    }
  }
}

// 顺序排序
const handleSequentialSort = async () => {
  try {
    await ElMessageBox.confirm(
      `确定要对选中的 ${selectedDepartments.value.length} 个部门进行顺序排序吗？\n排序值将从 ${batchSortValue.value} 开始递增。`,
      '顺序排序确认',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'info'
      }
    )
    
    await sequentialSort()
  } catch (error) {
    if (error !== 'cancel') {
      console.error('顺序排序失败:', error)
    }
  }
}

// 批量更新排序
const batchUpdateSort = async (sortValue) => {
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
          status: department.status,
          sort: sortValue,
          deviceIds: department.deviceIds || []
        })
        
        if (response.code === 0) {
          department.sort = sortValue
          successCount++
        } else {
          failCount++
        }
      } catch (error) {
        failCount++
      }
    }
    
    if (successCount > 0) {
      ElMessage.success(`成功更新 ${successCount} 个部门的排序`)
      await getData() // 重新获取数据以更新排序
    }
    
    if (failCount > 0) {
      ElMessage.warning(`${failCount} 个部门排序更新失败`)
    }
    
    clearSelection()
  } catch (error) {
    console.error('批量更新排序失败:', error)
    ElMessage.error('批量更新排序失败')
  } finally {
    batchLoading.value = false
  }
}

// 顺序排序
const sequentialSort = async () => {
  try {
    batchLoading.value = true
    let successCount = 0
    let failCount = 0
    let currentSort = batchSortValue.value
    
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
          status: department.status,
          sort: currentSort,
          deviceIds: department.deviceIds || []
        })
        
        if (response.code === 0) {
          department.sort = currentSort
          currentSort += 10 // 递增10，留出调整空间
          successCount++
        } else {
          failCount++
        }
      } catch (error) {
        failCount++
      }
    }
    
    if (successCount > 0) {
      ElMessage.success(`成功设置 ${successCount} 个部门的顺序排序`)
      await getData() // 重新获取数据以更新排序
    }
    
    if (failCount > 0) {
      ElMessage.warning(`${failCount} 个部门排序设置失败`)
    }
    
    clearSelection()
  } catch (error) {
    console.error('顺序排序失败:', error)
    ElMessage.error('顺序排序失败')
  } finally {
    batchLoading.value = false
  }
}

// 自动排序
const autoSort = async () => {
  try {
    await ElMessageBox.confirm(
      '确定要对所有部门进行自动排序吗？\n系统将按照部门层级和名称自动分配排序值。',
      '自动排序确认',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    
    autoSortLoading.value = true
    
    // 实现自动排序逻辑
    const sortedDepartments = [...departments.value].sort((a, b) => {
      // 先按层级排序，再按名称排序
      const levelA = getDepartmentLevelNumber(a)
      const levelB = getDepartmentLevelNumber(b)
      
      if (levelA !== levelB) {
        return levelA - levelB
      }
      
      return a.name.localeCompare(b.name)
    })
    
    let successCount = 0
    let failCount = 0
    
    for (let i = 0; i < sortedDepartments.length; i++) {
      const department = sortedDepartments[i]
      const newSort = (i + 1) * 10
      
      try {
        const response = await updateDepartment({
          id: department.id,
          name: department.name,
          parentId: department.parentId,
          leader: department.leader,
          phone: department.phone,
          email: department.email,
          status: department.status,
          sort: newSort,
          deviceIds: department.deviceIds || []
        })
        
        if (response.code === 0) {
          department.sort = newSort
          successCount++
        } else {
          failCount++
        }
      } catch (error) {
        failCount++
      }
    }
    
    if (successCount > 0) {
      ElMessage.success(`自动排序完成，成功更新 ${successCount} 个部门`)
      await getData()
    }
    
    if (failCount > 0) {
      ElMessage.warning(`${failCount} 个部门自动排序失败`)
    }
  } catch (error) {
    if (error !== 'cancel') {
      console.error('自动排序失败:', error)
      ElMessage.error('自动排序失败')
    }
  } finally {
    autoSortLoading.value = false
  }
}

// 单个排序变更
const handleSortChange = async (department) => {
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
      ElMessage.success('排序更新成功')
      // 延迟刷新以显示新的排序
      setTimeout(() => {
        getData()
      }, 500)
    } else {
      ElMessage.error(response.msg || '排序更新失败')
    }
  } catch (error) {
    console.error('排序更新失败:', error)
    ElMessage.error('排序更新失败')
  }
}

// 向上移动
const moveUp = (index) => {
  if (index > 0) {
    const temp = departments.value[index]
    departments.value[index] = departments.value[index - 1]
    departments.value[index - 1] = temp
    
    // 更新排序值
    updatePositionSort(index - 1, index)
  }
}

// 向下移动
const moveDown = (index) => {
  if (index < departments.value.length - 1) {
    const temp = departments.value[index]
    departments.value[index] = departments.value[index + 1]
    departments.value[index + 1] = temp
    
    // 更新排序值
    updatePositionSort(index, index + 1)
  }
}

// 拖拽结束
const handleDragEnd = (event) => {
  const { oldIndex, newIndex } = event
  if (oldIndex !== newIndex) {
    updatePositionSort(Math.min(oldIndex, newIndex), Math.max(oldIndex, newIndex))
  }
}

// 更新位置排序
const updatePositionSort = async (startIndex, endIndex) => {
  try {
    for (let i = startIndex; i <= endIndex; i++) {
      const department = departments.value[i]
      const newSort = (i + 1) * 10
      
      if (department.sort !== newSort) {
        await updateDepartment({
          id: department.id,
          name: department.name,
          parentId: department.parentId,
          leader: department.leader,
          phone: department.phone,
          email: department.email,
          status: department.status,
          sort: newSort,
          deviceIds: department.deviceIds || []
        })
        
        department.sort = newSort
      }
    }
    
    ElMessage.success('排序更新成功')
  } catch (error) {
    console.error('更新位置排序失败:', error)
    ElMessage.error('更新位置排序失败')
    getData() // 重新获取数据
  }
}

// 获取部门层级
const getDepartmentLevel = (department) => {
  return department.parentId ? `子部门 (上级: ${department.parentId})` : '顶级部门'
}

// 获取部门层级数字
const getDepartmentLevelNumber = (department) => {
  // 简化实现，实际应该递归计算层级深度
  return department.parentId ? 2 : 1
}

// 组件挂载时获取数据
onMounted(() => {
  getData()
})
</script>

<style scoped>
.sort-manager {
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

.header-actions {
  display: flex;
  gap: 8px;
}

.manager-content {
  padding: 20px;
}

.sort-info {
  margin-bottom: 24px;
}

.batch-sort {
  margin-bottom: 24px;
}

.batch-header {
  font-size: 14px;
  color: #606266;
  font-weight: 500;
  margin-bottom: 12px;
}

.batch-controls {
  display: flex;
  align-items: center;
  gap: 12px;
}

.list-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 14px;
  color: #606266;
  font-weight: 500;
  margin-bottom: 12px;
}

.list-content {
  max-height: 400px;
  overflow-y: auto;
}

.drag-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.department-item {
  display: flex;
  align-items: center;
  padding: 12px;
  border: 1px solid #e4e7ed;
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.2s;
  gap: 12px;
}

.department-item:hover {
  border-color: #409eff;
  background-color: #f0f9ff;
}

.department-item.selected {
  border-color: #409eff;
  background-color: #e6f7ff;
}

.drag-handle {
  cursor: move;
  color: #909399;
  font-size: 16px;
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

.department-level {
  font-size: 12px;
  color: #909399;
}

.sort-controls {
  margin-right: 12px;
}

.quick-actions {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

:deep(.el-alert__content) {
  padding: 0;
}

:deep(.el-alert__content p) {
  margin: 2px 0;
  font-size: 13px;
}
</style>