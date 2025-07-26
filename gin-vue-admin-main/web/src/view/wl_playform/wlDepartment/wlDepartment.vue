
<template>
  <div class="department-manage-page">
    <!-- 面包屑导航 -->
    <el-breadcrumb separator="/" class="breadcrumb-bar">
      <el-breadcrumb-item>首页</el-breadcrumb-item>
      <el-breadcrumb-item>部门管理</el-breadcrumb-item>
    </el-breadcrumb>
    <!-- 筛选区 -->
    <div class="filter-bar">
      <el-date-picker
        v-model="searchInfo.createdAtRange"
        type="daterange"
        range-separator="至"
        start-placeholder="开始时间"
        end-placeholder="结束时间"
        class="date-picker"
        style="width: 260px; margin-right: 12px;"
      />
      <el-button type="primary" @click="onSubmit">查询</el-button>
      <el-button @click="onReset">重置</el-button>
      <el-button link @click="showAllQuery = !showAllQuery">{{ showAllQuery ? '收起' : '展开' }}</el-button>
    </div>
    <!-- 新增按钮 -->
    <div class="action-bar">
      <el-button type="primary" @click="openDialog" icon="plus">新增</el-button>
    </div>
    <!-- 表格 -->
    <el-table :data="tableData" style="width: 100%;" v-loading="loading" empty-text="暂无数据">
      <el-table-column prop="createdAt" label="日期" width="160">
        <template #default="scope">{{ formatDate(scope.row.createdAt) }}</template>
      </el-table-column>
      <el-table-column prop="departmentName" label="部门名称" min-width="120" />
      <el-table-column prop="leader" label="负责人" min-width="100" />
      <el-table-column prop="phone" label="电话" min-width="120" />
      <el-table-column prop="email" label="邮箱" min-width="160" />
      <el-table-column prop="status" label="状态" min-width="80" />
      <el-table-column prop="sort" label="排序" min-width="80" />
      <el-table-column label="操作" width="140">
        <template #default="scope">
          <el-button type="primary" link @click="updateWlDepartmentFunc(scope.row)">编辑</el-button>
          <el-button type="danger" link @click="deleteRow(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <!-- 分页 -->
    <div class="pagination-bar">
      <el-pagination
        layout="total, sizes, prev, pager, next, jumper"
        :current-page="page"
        :page-size="pageSize"
        :page-sizes="[10, 30, 50, 100]"
        :total="total"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
      />
    </div>
    <!-- 新增/编辑弹窗 -->
    <el-drawer v-model="dialogFormVisible" title="新增" size="500px" direction="rtl" :show-close="false" :before-close="closeDialog">
      <el-form :model="formData" ref="elFormRef" label-width="90px" label-position="top" :rules="rule">
        <el-form-item label="部门名称" prop="departmentName">
          <el-input v-model="formData.departmentName" placeholder="请输入部门名称" />
        </el-form-item>
        <el-form-item label="上级部门" prop="parentId">
          <el-tree-select
            v-model="formData.parentId"
            :data="departmentTree"
            :props="{ label: 'departmentName', value: 'id', children: 'children' }"
            placeholder="请选择上级部门"
            clearable
            filterable
            check-strictly
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="负责人" prop="leader">
          <el-input v-model="formData.leader" placeholder="请输入负责人" />
        </el-form-item>
        <el-form-item label="电话" prop="phone">
          <el-input v-model="formData.phone" placeholder="请输入电话" />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="formData.email" placeholder="请输入邮箱" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-input v-model="formData.status" placeholder="请输入状态" />
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input v-model.number="formData.sort" placeholder="请输入排序" />
        </el-form-item>
        <el-form-item label="设备分配" prop="deviceIds">
          <el-select
            v-model="formData.deviceIds"
            multiple
            filterable
            remote
            reserve-keyword
            placeholder="请选择设备"
            :remote-method="fetchDeviceList"
            :loading="deviceLoading"
            style="width: 100%"
          >
            <el-option
              v-for="item in deviceList"
              :key="item.id"
              :label="item.name + (item.productName ? '（' + item.productName + '）' : '')"
              :value="item.id"
            />
          </el-select>
        </el-form-item>
        <div class="drawer-footer">
          <el-button type="primary" @click="enterDialog">确定</el-button>
          <el-button @click="closeDialog">取消</el-button>
        </div>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import { createWlDepartment, updateWlDepartment, deleteWlDepartment, getWlDepartmentList, getDevicesByDepartment } from '@/api/wl_playform/wlDepartment'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { formatDate } from '@/utils/format'

const showAllQuery = ref(false)
const loading = ref(false)
const dialogFormVisible = ref(false)
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const tableData = ref([])
const searchInfo = ref({ createdAtRange: [] })
const elFormRef = ref()
const formData = ref({
  departmentName: '',
  parentId: undefined,
  leader: '',
  phone: '',
  email: '',
  status: '',
  sort: undefined,
  deviceIds: []
})
const rule = reactive({
  departmentName: [{ required: true, message: '请输入部门名称', trigger: 'blur' }]
})
const departmentTree = ref([])
const deviceList = ref([])
const deviceLoading = ref(false)

const getTableData = async () => {
  loading.value = true
  const res = await getWlDepartmentList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (res.code === 0) {
    tableData.value = res.data.list
    total.value = res.data.total
    page.value = res.data.page
    pageSize.value = res.data.pageSize
  }
  loading.value = false
}
const onReset = () => {
  searchInfo.value = { createdAtRange: [] }
  getTableData()
}
const onSubmit = () => {
  getTableData()
}
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}
const openDialog = () => {
  dialogFormVisible.value = true
  fetchDepartmentTree()
  fetchDeviceList()
}
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    departmentName: '',
    parentId: undefined,
    leader: '',
    phone: '',
    email: '',
    status: '',
    sort: undefined,
    deviceIds: []
  }
}
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    if (formData.value.id) {
      res = await updateWlDepartment(formData.value)
    } else {
      res = await createWlDepartment(formData.value)
    }
    if (res.code === 0) {
      ElMessage.success('操作成功')
      closeDialog()
      getTableData()
    }
  })
}
const updateWlDepartmentFunc = (row) => {
  formData.value = { ...row, deviceIds: row.deviceIds || [] }
  dialogFormVisible.value = true
  fetchDepartmentTree()
  fetchDeviceList()
}
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await deleteWlDepartment({ ID: row.id })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      getTableData()
    }
  })
}
// 部门树
const fetchDepartmentTree = async () => {
  const res = await getWlDepartmentList({ page: 1, pageSize: 9999 })
  if (res.code === 0) {
    departmentTree.value = buildTree(res.data.list)
  }
}
function buildTree(list, parentId = null) {
  return list.filter(item => item.parentId === parentId).map(item => ({
    ...item,
    children: buildTree(list, item.id)
  }))
}
// 设备列表
const fetchDeviceList = async (query) => {
  deviceLoading.value = true
  const res = await getDevicesByDepartment({ departmentId: 0 })
  if (res.code === 0) {
    deviceList.value = res.data || []
  }
  deviceLoading.value = false
}
getTableData()
</script>

<style scoped>
.department-manage-page {
  background: #fff;
  padding: 24px;
}
.breadcrumb-bar {
  margin-bottom: 16px;
}
.filter-bar {
  display: flex;
  align-items: center;
  margin-bottom: 16px;
}
.action-bar {
  margin-bottom: 16px;
}
.pagination-bar {
  margin-top: 16px;
  text-align: right;
}
.drawer-footer {
  display: flex;
  justify-content: flex-end;
  margin-top: 24px;
}
</style>
