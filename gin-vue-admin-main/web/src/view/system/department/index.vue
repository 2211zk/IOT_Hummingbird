<template>
  <div class="department-management">
    <!-- 页面头部 -->
    <div class="page-header">
      <div class="header-content">
        <h2 class="page-title">部门管理</h2>
        <p class="page-description">部门管理可进行组织架构管理，新增部门、联络等功能操作，用于管理平台组织。</p>
      </div>
    </div>

    <!-- 搜索区域 -->
    <div class="search-container">
      <el-form :model="searchForm" :inline="true" class="search-form">
        <el-form-item>
          <el-input
            v-model="searchForm.name"
            placeholder="请输入部门名称"
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

    <!-- 操作区域 -->
    <div class="action-container">
      <el-button type="primary" @click="handleAdd" icon="Plus">
        新增
      </el-button>
    </div>

    <!-- 表格区域 -->
    <div class="table-container">
      <el-table
        :data="tableData"
        v-loading="loading"
        row-key="id"
        :tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
        :expand-row-keys="expandedKeys"
        @expand-change="handleExpandChange"
        style="width: 100%"
      >
        <el-table-column prop="name" label="部门名称" min-width="200">
          <template #default="{ row }">
            <div class="department-name">
              <span>{{ row.name }}</span>
            </div>
          </template>
        </el-table-column>
        
        <el-table-column prop="leader" label="负责人" width="120">
          <template #default="{ row }">
            <span>{{ row.leader || '-' }}</span>
          </template>
        </el-table-column>
        
        <el-table-column prop="phone" label="电话" width="140">
          <template #default="{ row }">
            <span>{{ row.phone || '-' }}</span>
          </template>
        </el-table-column>
        
        <el-table-column prop="email" label="邮箱" width="180">
          <template #default="{ row }">
            <span>{{ row.email || '-' }}</span>
          </template>
        </el-table-column>
        
        <el-table-column prop="status" label="状态" width="120">
          <template #default="{ row }">
            <el-switch
              v-model="row.status"
              active-text="启用"
              inactive-text="禁用"
              active-value="启用"
              inactive-value="禁用"
              @change="handleStatusChange(row)"
              :loading="row.statusLoading"
            />
          </template>
        </el-table-column>
        
        <el-table-column prop="sort" label="排序" width="100">
          <template #default="{ row }">
            <el-input-number
              v-model="row.sort"
              :min="0"
              :max="9999"
              size="small"
              controls-position="right"
              @change="handleSortChange(row)"
              :loading="row.sortLoading"
              style="width: 80px"
            />
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
              @click="handleView(row)"
            >
              查看
            </el-button>
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
    <div class="pagination-container" v-if="!searchForm.treeMode">
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

    <!-- 部门编辑弹窗 -->
    <DepartmentDialog
      v-model="dialogVisible"
      :mode="dialogMode"
      :department-id="currentDepartmentId"
      @success="handleDialogSuccess"
    />
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getDepartmentList, deleteDepartment, updateDepartment, getDepartmentDevices } from '@/api/wlDepartment'
import { formatDateTime } from '@/utils/format'
import { useLoading, useError } from '@/utils/loadingManager'
import { debounceApiCall } from '@/utils/apiError'
import { useNotification } from '@/utils/userNotification'
import { usePerformanceMonitor } from '@/utils/performanceMonitor'
import DepartmentDialog from './components/DepartmentDialog.vue'

// 响应式数据
const tableData = ref([])
const expandedKeys = ref([])

// 使用加载和错误管理
const { getLoading, loadingStates } = useLoading()
const { handleError } = useError()
const { notify, departmentNotifications } = useNotification()
const { startMeasure, endMeasure } = usePerformanceMonitor()

// 各种加载状态
const loading = computed(() => getLoading('departmentList'))
const treeLoading = ref(false)
const searchLoading = ref(false)

// 搜索表单
const searchForm = reactive({
  name: '',
  status: '',
  treeMode: true // 默认使用树形模式
})

// 分页数据
const pagination = reactive({
  page: 1,
  pageSize: 20,
  total: 0
})

// 获取表格数据
const getTableData = async () => {
  startMeasure('getDepartmentList')
  
  try {
    const params = {
      ...searchForm,
      page: pagination.page,
      pageSize: pagination.pageSize
    }
    
    const response = await getDepartmentList(params)
    if (response.code === 0) {
      if (searchForm.treeMode) {
        // 树形模式
        tableData.value = response.data.list || []
        pagination.total = response.data.total || 0
      } else {
        // 平铺模式
        tableData.value = response.data.list || []
        pagination.total = response.data.total || 0
        pagination.page = response.data.page || 1
        pagination.pageSize = response.data.pageSize || 20
      }
    } else {
      notify.error(response.msg || '获取部门列表失败')
    }
  } catch (error) {
    if (error.code === 'NETWORK_ERROR') {
      notify.error('网络连接失败，请检查网络设置', { notification: true })
    } else {
      notify.error('获取部门列表失败，请稍后重试', { notification: true })
    }
  } finally {
    endMeasure('getDepartmentList')
  }
}

// 防抖搜索
const debouncedSearch = debounceApiCall(getTableData, 300)

// 搜索
const handleSearch = () => {
  pagination.page = 1
  debouncedSearch()
}

// 重置
const handleReset = () => {
  searchForm.name = ''
  searchForm.status = ''
  pagination.page = 1
  getTableData()
}

// 弹窗相关
const dialogVisible = ref(false)
const dialogMode = ref('create')
const currentDepartmentId = ref(null)

// 新增
const handleAdd = () => {
  dialogMode.value = 'create'
  currentDepartmentId.value = null
  dialogVisible.value = true
}

// 查看
const handleView = (row) => {
  dialogMode.value = 'view'
  currentDepartmentId.value = row.id
  dialogVisible.value = true
}

// 编辑
const handleEdit = (row) => {
  dialogMode.value = 'edit'
  currentDepartmentId.value = row.id
  dialogVisible.value = true
}

// 弹窗成功回调
const handleDialogSuccess = () => {
  getTableData()
}

// 状态切换
const handleStatusChange = async (row) => {
  try {
    // 如果要禁用部门，先确认
    if (row.status === '禁用') {
      const confirmed = await departmentNotifications.confirmDisable(row.name)
      if (!confirmed) {
        row.status = '启用' // 恢复原状态
        return
      }
    }
    
    // 添加加载状态
    row.statusLoading = true
    
    const response = await updateDepartment({
      id: row.id,
      name: row.name,
      parentId: row.parentId,
      leader: row.leader,
      phone: row.phone,
      email: row.email,
      status: row.status,
      sort: row.sort,
      deviceIds: row.deviceIds || []
    })
    
    if (response.code === 0) {
      if (row.status === '启用') {
        departmentNotifications.departmentEnabled(row.name)
      } else {
        departmentNotifications.departmentDisabled(row.name)
      }
      
      // 如果禁用部门，需要检查是否有子部门需要同时禁用
      if (row.status === '禁用') {
        await handleDisableChildren(row)
      }
      
      // 刷新数据以获取最新状态
      getTableData()
    } else {
      // 恢复原状态
      row.status = row.status === '启用' ? '禁用' : '启用'
      notify.error(response.msg || '状态更新失败')
    }
  } catch (error) {
    // 恢复原状态
    row.status = row.status === '启用' ? '禁用' : '启用'
    notify.error('状态更新失败，请稍后重试')
  } finally {
    row.statusLoading = false
  }
}

// 禁用子部门
const handleDisableChildren = async (parentRow) => {
  if (parentRow.children && parentRow.children.length > 0) {
    for (const child of parentRow.children) {
      if (child.status === '启用') {
        try {
          await updateDepartment({
            id: child.id,
            name: child.name,
            parentId: child.parentId,
            leader: child.leader,
            phone: child.phone,
            email: child.email,
            status: '禁用',
            sort: child.sort,
            deviceIds: child.deviceIds || []
          })
          
          // 递归处理子部门的子部门
          await handleDisableChildren(child)
        } catch (error) {
          console.error(`禁用子部门 ${child.name} 失败:`, error)
        }
      }
    }
  }
}

// 排序变更
const handleSortChange = async (row) => {
  try {
    // 添加加载状态
    row.sortLoading = true
    
    const response = await updateDepartment({
      id: row.id,
      name: row.name,
      parentId: row.parentId,
      leader: row.leader,
      phone: row.phone,
      email: row.email,
      status: row.status,
      sort: row.sort,
      deviceIds: row.deviceIds || []
    })
    
    if (response.code === 0) {
      ElMessage.success('排序更新成功')
      // 延迟刷新以显示新的排序
      setTimeout(() => {
        getTableData()
      }, 500)
    } else {
      ElMessage.error(response.msg || '排序更新失败')
    }
  } catch (error) {
    console.error('排序更新失败:', error)
    ElMessage.error('排序更新失败')
  } finally {
    row.sortLoading = false
  }
}

// 删除
const handleDelete = async (row) => {
  try {
    // 检查删除权限
    if (!checkDeletePermission(row)) {
      return
    }
    
    // 检查子部门和设备关联
    const hasChildren = await checkHasChildren(row.id)
    const hasDevices = await checkHasDevices(row.id)
    
    // 使用专用的确认对话框
    const confirmed = await departmentNotifications.confirmDelete(row.name, hasChildren, hasDevices)
    if (!confirmed) {
      return
    }
    
    const response = await deleteDepartment({ id: row.id })
    if (response.code === 0) {
      departmentNotifications.departmentDeleted(row.name)
      getTableData()
    } else {
      if (response.code === 7003) {
        departmentNotifications.cannotDeleteWithChildren(row.name)
      } else {
        notify.error(response.msg || '删除失败')
      }
    }
  } catch (error) {
    if (error !== 'cancel') {
      notify.error('删除操作失败，请稍后重试')
    }
  }
}

// 检查删除权限
const checkDeletePermission = (row) => {
  // 检查用户权限（这里可以根据实际权限系统实现）
  const userRole = getUserRole() // 假设有获取用户角色的函数
  
  // 超级管理员可以删除任何部门
  if (userRole === 'super_admin') {
    return true
  }
  
  // 管理员可以删除非顶级部门
  if (userRole === 'admin' && row.parentId) {
    return true
  }
  
  // 普通用户不能删除部门
  if (userRole === 'user') {
    departmentNotifications.noPermissionToDelete()
    return false
  }
  
  // 不能删除顶级部门（除非是超级管理员）
  if (!row.parentId && userRole !== 'super_admin') {
    notify.error('不能删除顶级部门')
    return false
  }
  
  return true
}

// 检查是否有子部门
const checkHasChildren = async (departmentId) => {
  try {
    const response = await getDepartmentList({
      page: 1,
      pageSize: 1,
      parentId: departmentId
    })
    
    return response.code === 0 && response.data.total > 0
  } catch (error) {
    console.error('检查子部门失败:', error)
    return false
  }
}

// 检查是否有关联设备
const checkHasDevices = async (departmentId) => {
  try {
    const response = await getDepartmentDevices({
      departmentId,
      page: 1,
      pageSize: 1
    })
    
    return response.code === 0 && response.data.total > 0
  } catch (error) {
    console.error('检查关联设备失败:', error)
    return false
  }
}

// 获取用户角色（模拟函数，实际应该从用户状态或权限系统获取）
const getUserRole = () => {
  // 这里应该从实际的用户状态管理中获取
  // 暂时返回管理员角色用于演示
  return 'admin'
}

// 展开/收起处理
const handleExpandChange = (row, expanded) => {
  if (expanded) {
    if (!expandedKeys.value.includes(row.id)) {
      expandedKeys.value.push(row.id)
    }
  } else {
    const index = expandedKeys.value.indexOf(row.id)
    if (index > -1) {
      expandedKeys.value.splice(index, 1)
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
.department-management {
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

.department-name {
  display: flex;
  align-items: center;
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