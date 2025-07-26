<template>
  <el-dialog
    v-model="visible"
    title="选择设备"
    width="1000px"
    :close-on-click-modal="false"
    @close="handleClose"
  >
    <div class="device-selector">
      <!-- 搜索区域 -->
      <div class="search-container">
        <el-form :model="searchForm" :inline="true">
          <el-form-item label="设备名称">
            <el-input
              v-model="searchForm.deviceName"
              placeholder="请输入设备名称"
              clearable
              style="width: 200px"
              @keyup.enter="handleSearch"
            />
          </el-form-item>
          <el-form-item label="产品名称">
            <el-input
              v-model="searchForm.productName"
              placeholder="请输入产品名称"
              clearable
              style="width: 200px"
              @keyup.enter="handleSearch"
            />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="handleSearch" :loading="loading">
              查询
            </el-button>
            <el-button @click="handleReset">重置</el-button>
          </el-form-item>
        </el-form>
      </div>
      
      <!-- 设备选择内容 -->
      <div class="selector-content">
        <div class="available-devices">
          <div class="section-header">
            <h4>待选设备</h4>
            <span class="device-count">共 {{ pagination.total }} 个设备</span>
          </div>
          
          <el-table
            ref="availableTableRef"
            :data="availableDevices"
            v-loading="loading"
            @selection-change="handleSelectionChange"
            height="400"
            style="width: 100%"
          >
            <el-table-column type="selection" width="55" />
            <el-table-column prop="deviceName" label="设备名称" min-width="150" />
            <el-table-column prop="productName" label="产品名称" min-width="150">
              <template #default="{ row }">
                <span>{{ row.productName || '-' }}</span>
              </template>
            </el-table-column>
            <el-table-column prop="status" label="状态" width="80">
              <template #default="{ row }">
                <el-tag :type="getStatusType(row.status)">
                  {{ row.status }}
                </el-tag>
              </template>
            </el-table-column>
          </el-table>
          
          <!-- 分页 -->
          <div class="pagination-container">
            <el-pagination
              v-model:current-page="pagination.page"
              v-model:page-size="pagination.pageSize"
              :page-sizes="[10, 20, 50]"
              :total="pagination.total"
              layout="total, sizes, prev, pager, next"
              @size-change="handleSizeChange"
              @current-change="handleCurrentChange"
            />
          </div>
        </div>
        
        <div class="transfer-buttons">
          <el-button
            type="primary"
            :disabled="currentSelection.length === 0"
            @click="addToSelected"
            icon="Right"
          >
            添加
          </el-button>
          <el-button
            :disabled="tempSelectedDevices.length === 0"
            @click="removeFromSelected"
            icon="Left"
          >
            移除
          </el-button>
        </div>
        
        <div class="selected-devices">
          <div class="section-header">
            <h4>已选设备</h4>
            <span class="device-count">已选择 {{ tempSelectedDevices.length }} 个设备</span>
          </div>
          
          <div class="selected-list">
            <div v-if="tempSelectedDevices.length === 0" class="empty-selected">
              <el-empty description="暂无选择设备" :image-size="60" />
            </div>
            <div v-else class="device-items">
              <div
                v-for="device in tempSelectedDevices"
                :key="device.id"
                class="device-item"
                :class="{ active: selectedForRemoval.includes(device.id) }"
                @click="toggleRemovalSelection(device.id)"
              >
                <div class="device-info">
                  <div class="device-name">{{ device.deviceName }}</div>
                  <div class="product-name">{{ device.productName || '-' }}</div>
                </div>
                <el-button
                  type="danger"
                  size="small"
                  text
                  @click.stop="removeDevice(device.id)"
                  icon="Close"
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="handleClose">取消</el-button>
        <el-button type="primary" @click="handleConfirm">
          确定 ({{ tempSelectedDevices.length }})
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import { ref, reactive, computed, watch, nextTick } from 'vue'
import { ElMessage } from 'element-plus'
import { getAvailableDevices } from '@/api/wlDepartment'

// Props
const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  selectedDevices: {
    type: Array,
    default: () => []
  },
  departmentId: {
    type: Number,
    default: null
  }
})

// Emits
const emit = defineEmits(['update:modelValue', 'confirm'])

// 响应式数据
const loading = ref(false)
const availableTableRef = ref()
const availableDevices = ref([])
const currentSelection = ref([])
const tempSelectedDevices = ref([])
const selectedForRemoval = ref([])

// 搜索表单
const searchForm = reactive({
  deviceName: '',
  productName: ''
})

// 分页数据
const pagination = reactive({
  page: 1,
  pageSize: 10,
  total: 0
})

// 计算属性
const visible = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

// 监听弹窗显示状态
watch(visible, (newVal) => {
  if (newVal) {
    initializeData()
    loadAvailableDevices()
  } else {
    resetData()
  }
})

// 初始化数据
const initializeData = () => {
  tempSelectedDevices.value = [...props.selectedDevices]
  selectedForRemoval.value = []
  currentSelection.value = []
}

// 重置数据
const resetData = () => {
  searchForm.deviceName = ''
  searchForm.productName = ''
  pagination.page = 1
  pagination.pageSize = 10
  pagination.total = 0
  availableDevices.value = []
  currentSelection.value = []
  selectedForRemoval.value = []
}

// 加载可用设备
const loadAvailableDevices = async () => {
  try {
    loading.value = true
    const params = {
      page: pagination.page,
      pageSize: pagination.pageSize,
      deviceName: searchForm.deviceName,
      productName: searchForm.productName,
      departmentId: props.departmentId // 排除已关联当前部门的设备
    }
    
    const response = await getAvailableDevices(params)
    if (response.code === 0) {
      availableDevices.value = response.data.list || []
      pagination.total = response.data.total || 0
      
      // 清除之前的选择状态
      nextTick(() => {
        availableTableRef.value?.clearSelection()
      })
    } else {
      ElMessage.error(response.msg || '获取设备列表失败')
    }
  } catch (error) {
    console.error('获取设备列表失败:', error)
    ElMessage.error('获取设备列表失败')
  } finally {
    loading.value = false
  }
}

// 搜索
const handleSearch = () => {
  pagination.page = 1
  loadAvailableDevices()
}

// 重置搜索
const handleReset = () => {
  searchForm.deviceName = ''
  searchForm.productName = ''
  pagination.page = 1
  loadAvailableDevices()
}

// 分页处理
const handleSizeChange = (size) => {
  pagination.pageSize = size
  pagination.page = 1
  loadAvailableDevices()
}

const handleCurrentChange = (page) => {
  pagination.page = page
  loadAvailableDevices()
}

// 表格选择变化
const handleSelectionChange = (selection) => {
  currentSelection.value = selection
}

// 添加到已选
const addToSelected = () => {
  const newDevices = currentSelection.value.filter(device => 
    !tempSelectedDevices.value.some(selected => selected.id === device.id)
  )
  
  tempSelectedDevices.value.push(...newDevices)
  
  // 清除表格选择
  availableTableRef.value?.clearSelection()
  currentSelection.value = []
  
  ElMessage.success(`已添加 ${newDevices.length} 个设备`)
}

// 从已选中移除
const removeFromSelected = () => {
  if (selectedForRemoval.value.length === 0) {
    ElMessage.warning('请先选择要移除的设备')
    return
  }
  
  tempSelectedDevices.value = tempSelectedDevices.value.filter(device =>
    !selectedForRemoval.value.includes(device.id)
  )
  
  ElMessage.success(`已移除 ${selectedForRemoval.value.length} 个设备`)
  selectedForRemoval.value = []
}

// 切换移除选择状态
const toggleRemovalSelection = (deviceId) => {
  const index = selectedForRemoval.value.indexOf(deviceId)
  if (index > -1) {
    selectedForRemoval.value.splice(index, 1)
  } else {
    selectedForRemoval.value.push(deviceId)
  }
}

// 移除单个设备
const removeDevice = (deviceId) => {
  tempSelectedDevices.value = tempSelectedDevices.value.filter(device => device.id !== deviceId)
  selectedForRemoval.value = selectedForRemoval.value.filter(id => id !== deviceId)
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

// 确认选择
const handleConfirm = () => {
  emit('confirm', tempSelectedDevices.value)
  handleClose()
}

// 关闭弹窗
const handleClose = () => {
  visible.value = false
}
</script>

<style scoped>
.device-selector {
  height: 600px;
  display: flex;
  flex-direction: column;
}

.search-container {
  margin-bottom: 20px;
  padding: 16px;
  background: #f5f7fa;
  border-radius: 4px;
}

.selector-content {
  flex: 1;
  display: flex;
  gap: 16px;
  min-height: 0;
}

.available-devices,
.selected-devices {
  flex: 1;
  display: flex;
  flex-direction: column;
  border: 1px solid #e4e7ed;
  border-radius: 4px;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background: #fafafa;
  border-bottom: 1px solid #e4e7ed;
}

.section-header h4 {
  margin: 0;
  font-size: 14px;
  font-weight: 600;
  color: #303133;
}

.device-count {
  font-size: 12px;
  color: #909399;
}

.pagination-container {
  padding: 12px 16px;
  border-top: 1px solid #e4e7ed;
  background: #fafafa;
}

.transfer-buttons {
  display: flex;
  flex-direction: column;
  justify-content: center;
  gap: 12px;
  padding: 0 8px;
}

.selected-list {
  flex: 1;
  padding: 16px;
  overflow-y: auto;
}

.empty-selected {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 200px;
}

.device-items {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.device-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px;
  border: 1px solid #e4e7ed;
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.2s;
}

.device-item:hover {
  border-color: #409eff;
  background-color: #f0f9ff;
}

.device-item.active {
  border-color: #409eff;
  background-color: #e6f7ff;
}

.device-info {
  flex: 1;
}

.device-name {
  font-size: 14px;
  font-weight: 500;
  color: #303133;
  margin-bottom: 4px;
}

.product-name {
  font-size: 12px;
  color: #909399;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

:deep(.el-table) {
  border: none;
}

:deep(.el-table__header) {
  background-color: #fafafa;
}

:deep(.el-pagination) {
  justify-content: center;
}
</style>