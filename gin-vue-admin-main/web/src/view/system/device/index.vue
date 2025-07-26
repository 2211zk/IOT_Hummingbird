<template>
  <div class="device-management">
    <!-- 页面头部 -->
    <div class="page-header">
      <div class="header-content">
        <h2 class="page-title">设备管理</h2>
        <p class="page-description">设备管理用于维护系统中的设备信息，支持设备的增删改查操作。</p>
      </div>
    </div>

    <!-- 搜索区域 -->
    <div class="search-container">
      <el-form :model="searchForm" :inline="true" class="search-form">
        <el-form-item>
          <el-input
            v-model="searchForm.deviceName"
            placeholder="请输入设备名称"
            clearable
            style="width: 200px"
            @keyup.enter="handleSearch"
          />
        </el-form-item>
        <el-form-item>
          <el-input
            v-model="searchForm.productName"
            placeholder="请输入产品名称"
            clearable
            style="width: 200px"
            @keyup.enter="handleSearch"
          />
        </el-form-item>
        <el-form-item>
          <el-select
            v-model="searchForm.status"
            placeholder="请选择状态"
            clearable
            style="width: 120px"
          >
            <el-option label="启用" value="启用" />
            <el-option label="禁用" value="禁用" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch" :loading="loading">
            查询
          </el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- 操作区域 -->
    <div class="action-container">
      <el-button type="primary" @click="handleAdd" icon="Plus">
        新增设备
      </el-button>
    </div>

    <!-- 表格区域 -->
    <div class="table-container">
      <el-table
        :data="tableData"
        v-loading="loading"
        style="width: 100%"
      >
        <el-table-column prop="deviceName" label="设备名称" min-width="150" />
        <el-table-column prop="productName" label="产品名称" min-width="150">
          <template #default="{ row }">
            <span>{{ row.productName || '-' }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">
              {{ row.status }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createdAt" label="创建时间" width="160">
          <template #default="{ row }">
            <span>{{ formatDateTime(row.createdAt) }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button
              type="primary"
              link
              size="small"
              @click="handleEdit(row)"
            >
              编辑
            </el-button>
            <el-button
              type="danger"
              link
              size="small"
              @click="handleDelete(row)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <!-- 分页 -->
    <div class="pagination-container">
      <el-pagination
        v-model:current-page="pagination.page"
        v-model:page-size="pagination.pageSize"
        :page-sizes="[10, 20, 50, 100]"
        :total="pagination.total"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>

    <!-- 设备编辑弹窗 -->
    <DeviceDialog
      v-model="dialogVisible"
      :mode="dialogMode"
      :device-id="currentDeviceId"
      @success="handleDialogSuccess"
    />
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getDeviceList, deleteDevice } from '@/api/device'
import { formatDateTime } from '@/utils/format'
import DeviceDialog from './components/DeviceDialog.vue'

// 响应式数据
const loading = ref(false)
const tableData = ref([])

// 弹窗相关
const dialogVisible = ref(false)
const dialogMode = ref('create')
const currentDeviceId = ref(null)

// 搜索表单
const searchForm = reactive({
  deviceName: '',
  productName: '',
  status: ''
})

// 分页数据
const pagination = reactive({
  page: 1,
  pageSize: 20,
  total: 0
})

// 获取表格数据
const getTableData = async () => {
  try {
    loading.value = true
    const params = {
      ...searchForm,
      page: pagination.page,
      pageSize: pagination.pageSize
    }
    
    const response = await getDeviceList(params)
    if (response.code === 0) {
      tableData.value = response.data.list || []
      pagination.total = response.data.total || 0
      pagination.page = response.data.page || 1
      pagination.pageSize = response.data.pageSize || 20
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
  getTableData()
}

// 重置
const handleReset = () => {
  searchForm.deviceName = ''
  searchForm.productName = ''
  searchForm.status = ''
  pagination.page = 1
  getTableData()
}

// 新增
const handleAdd = () => {
  dialogMode.value = 'create'
  currentDeviceId.value = null
  dialogVisible.value = true
}

// 编辑
const handleEdit = (row) => {
  dialogMode.value = 'edit'
  currentDeviceId.value = row.id
  dialogVisible.value = true
}

// 删除
const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm(
      `确定要删除设备"${row.deviceName}"吗？`,
      '删除确认',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    
    const response = await deleteDevice({ id: row.id })
    if (response.code === 0) {
      ElMessage.success('删除成功')
      getTableData()
    } else {
      ElMessage.error(response.msg || '删除失败')
    }
  } catch (error) {
    if (error !== 'cancel') {
      console.error('删除设备失败:', error)
      ElMessage.error('删除失败')
    }
  }
}

// 分页处理
const handleSizeChange = (size) => {
  pagination.pageSize = size
  pagination.page = 1
  getTableData()
}

const handleCurrentChange = (page) => {
  pagination.page = page
  getTableData()
}

// 弹窗成功回调
const handleDialogSuccess = () => {
  getTableData()
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

// 组件挂载时获取数据
onMounted(() => {
  getTableData()
})
</script>

<style scoped>
.device-management {
  padding: 20px;
  background: #f5f7fa;
  min-height: 100vh;
}

.page-header {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 8px;
  padding: 30px;
  margin-bottom: 20px;
  color: white;
}

.header-content {
  max-width: 1200px;
}

.page-title {
  font-size: 28px;
  font-weight: 600;
  margin: 0 0 10px 0;
}

.page-description {
  font-size: 14px;
  opacity: 0.9;
  margin: 0;
  line-height: 1.5;
}

.search-container {
  background: white;
  border-radius: 8px;
  padding: 20px;
  margin-bottom: 20px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.search-form {
  margin: 0;
}

.action-container {
  background: white;
  border-radius: 8px;
  padding: 16px 20px;
  margin-bottom: 20px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.table-container {
  background: white;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.pagination-container {
  background: white;
  border-radius: 8px;
  padding: 16px 20px;
  margin-top: 20px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  text-align: right;
}

:deep(.el-table) {
  border-radius: 4px;
}

:deep(.el-table__header) {
  background-color: #fafafa;
}

:deep(.el-table__row:hover) {
  background-color: #f5f7fa;
}

:deep(.el-button + .el-button) {
  margin-left: 8px;
}

:deep(.el-tag) {
  border-radius: 12px;
}
</style>