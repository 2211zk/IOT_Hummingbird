<template>
  <div class="mine-driver-container">
    <!-- 顶部搜索栏 -->
    <el-card class="search-card">
      <el-form :inline="true" :model="searchForm" class="search-form">
        <el-form-item label="驱动名称">
          <el-input v-model="searchForm.name" placeholder="请输入驱动名称" clearable />
        </el-form-item>
        <el-form-item label="驱动类型">
          <el-select v-model="searchForm.driverType" placeholder="请选择驱动类型" clearable>
            <el-option
              v-for="item in driverTypeOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
          <el-button type="success" @click="handleAdd">新增驱动</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 驱动列表表格 -->
    <el-card class="table-card">
      <el-table
        :data="tableData"
        stripe
        border
        class="driver-table"
        style="width: 100%"
      >
        <el-table-column prop="driverId" label="驱动编号" width="110" align="center" />
        <el-table-column prop="driverName" label="驱动名称" min-width="180" />
        <el-table-column prop="version" label="版本" width="90" align="center" />
        <el-table-column prop="driverType" label="驱动类型" width="100" align="center">
          <template #default="scope">
            <el-tag :type="getDriverTypeTagType(scope.row.driverType)" effect="light">
              {{ getDriverTypeLabel(scope.row.driverType) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100" align="center">
          <template #default="scope">
            <el-tag :type="scope.row.status === '运行中' ? 'success' : 'info'" effect="plain">
              <i :class="scope.row.status === '运行中' ? 'el-icon-success' : 'el-icon-error'" style="margin-right:4px;"></i>
              {{ scope.row.status }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createdTime" label="创建时间" min-width="170" align="center" />
        <el-table-column label="操作" width="200" align="center">
          <template #default="scope">
            <el-link type="primary" @click="handleEdit(scope.row)">编辑</el-link>
            <span style="color:#dcdcdc;"> / </span>
            <el-link type="primary" @click="handleOperate(scope.row)">操作</el-link>
            <span style="color:#dcdcdc;"> / </span>
            <el-link type="primary" @click="handleLog(scope.row)">日志</el-link>
            <span style="color:#dcdcdc;"> / </span>
            <el-link type="danger" @click="handleDelete(scope.row)">删除</el-link>
          </template>
        </el-table-column>
      </el-table>
      <!-- 分页器 -->
      <div class="pagination-wrapper">
        <el-pagination
          background
          layout="total, prev, pager, next, sizes"
          :total="total"
          :page-size="pageSize"
          :current-page="currentPage"
          @size-change="handleSizeChange"
          @current-change="handlePageChange"
          :page-sizes="[10, 20, 50, 100]"
        />
      </div>
    </el-card>

    <!-- 驱动表单对话框 -->
    <DriverForm
      v-model:visible="formVisible"
      :form-type="formType"
      :driver-data="currentDriver"
      @success="handleFormSuccess"
    />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getWlDriversList, deleteWlDrivers } from '@/api/wl_driver/wlDrivers'
import { getDict } from '@/utils/dictionary'
import DriverForm from './DriverForm.vue'
import { ElMessage, ElMessageBox } from 'element-plus'

// 搜索表单
const searchForm = ref({
  name: '',
  driverType: ''
})

// 字典选项
const driverTypeOptions = ref([])

// 表格数据
const tableData = ref([])
const total = ref(0)
const pageSize = ref(10)
const currentPage = ref(1)

// 表单相关
const formVisible = ref(false)
const formType = ref('add')
const currentDriver = ref({})

// 获取字典数据
const loadDictionaries = async () => {
  try {
    const driverTypeDict = await getDict('driver_type')
    driverTypeOptions.value = driverTypeDict || []
  } catch (error) {
    console.error('加载字典数据失败:', error)
  }
}

// 获取驱动类型标签
const getDriverTypeLabel = (value) => {
  const option = driverTypeOptions.value.find(item => item.value === value)
  return option ? option.label : value
}

// 获取驱动类型标签样式
const getDriverTypeTagType = (value) => {
  return value === 'official' ? 'success' : 'info'
}

const getTableData = async () => {
  const params = {
    page: currentPage.value,
    pageSize: pageSize.value,
    driverName: searchForm.value.name,
    driverType: searchForm.value.driverType
  }
  const res = await getWlDriversList(params)
  if (res.code === 0) {
    tableData.value = res.data.list
    total.value = res.data.total
  }
}

onMounted(async () => {
  await loadDictionaries()
  getTableData()
})

const handleSearch = () => {
  currentPage.value = 1
  getTableData()
}

const handleReset = () => {
  searchForm.value.name = ''
  searchForm.value.driverType = ''
  currentPage.value = 1
  getTableData()
}

// 新增驱动
const handleAdd = () => {
  formType.value = 'add'
  currentDriver.value = {}
  formVisible.value = true
}

// 编辑驱动
const handleEdit = (row) => {
  formType.value = 'edit'
  currentDriver.value = { ...row }
  formVisible.value = true
}

// 表单提交成功
const handleFormSuccess = () => {
  getTableData()
}

const handleOperate = (row) => {
  // TODO: 操作逻辑
}

const handleLog = (row) => {
  // TODO: 日志逻辑
}

const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm('确定要删除这个驱动吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    
    const res = await deleteWlDrivers({ ID: row.id })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      getTableData()
    } else {
      ElMessage.error(res.msg || '删除失败')
    }
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('删除失败')
    }
  }
}

const handleSizeChange = (size) => {
  pageSize.value = size
  getTableData()
}

const handlePageChange = (page) => {
  currentPage.value = page
  getTableData()
}
</script>

<style scoped>
.mine-driver-container {
  padding: 20px;
  background: #f5f7fa;
  min-height: 100vh;
}
.search-card {
  margin-bottom: 18px;
  border-radius: 10px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.04);
}
.search-form {
  display: flex;
  align-items: center;
}
.table-card {
  border-radius: 12px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.06);
  padding-bottom: 10px;
}
.driver-table {
  border-radius: 8px;
  overflow: hidden;
}
.el-table th, .el-table td {
  text-align: center;
}
.el-tag {
  font-size: 13px;
  border-radius: 6px;
  padding: 0 10px;
}
.pagination-wrapper {
  display: flex;
  justify-content: flex-end;
  margin-top: 18px;
}
.el-link {
  font-size: 14px;
  margin: 0 2px;
}
</style> 