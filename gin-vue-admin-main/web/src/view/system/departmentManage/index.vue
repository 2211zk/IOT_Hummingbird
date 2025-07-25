<template>
  <div>
    <el-card>
      <div style="margin-bottom: 16px;">
        <el-form :inline="true" :model="searchForm">
          <el-form-item label="部门名称">
            <el-input v-model="searchForm.departmentName" placeholder="请输入部门名称" clearable />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="handleSearch">查询</el-button>
            <el-button @click="handleReset">重置</el-button>
          </el-form-item>
        </el-form>
      </div>
      <el-button type="primary" @click="handleAdd" style="margin-bottom: 12px;">新增</el-button>
      <el-table :data="depList" border style="width: 100%;">
        <el-table-column prop="departmentName" label="部门名称" />
        <el-table-column prop="leader" label="负责人" />
        <el-table-column prop="phone" label="电话" />
        <el-table-column prop="email" label="邮箱" />
        <el-table-column prop="status" label="状态" />
        <el-table-column prop="sort" label="排序" />
        <el-table-column prop="createdAt" label="创建时间" />
        <el-table-column label="操作">
          <template #default="scope">
            <el-button size="mini" @click="handleView(scope.row)">查看</el-button>
            <el-button size="mini" @click="handleEdit(scope.row)">编辑</el-button>
            <el-button size="mini" type="danger" @click="handleDelete(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        style="margin-top: 16px; text-align: right;"
        background
        layout="total, sizes, prev, pager, next, jumper"
        :total="total"
        :page-size="searchForm.pageSize"
        :current-page="searchForm.page"
        @size-change="handleSizeChange"
        @current-change="handlePageChange"
      />
    </el-card>
    <!-- 新增/编辑弹窗 -->
    <el-dialog :title="dialogTitle" v-model="dialogVisible" width="500px">
      <el-form :model="form" :rules="rules" ref="formRef" label-width="90px">
        <el-form-item label="部门名称" prop="departmentName">
          <el-input v-model="form.departmentName" />
        </el-form-item>
        <el-form-item label="负责人" prop="leader">
          <el-input v-model="form.leader" />
        </el-form-item>
        <el-form-item label="电话" prop="phone">
          <el-input v-model="form.phone" />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="form.email" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="form.status">
            <el-option label="启用" value="启用" />
            <el-option label="禁用" value="禁用" />
          </el-select>
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number v-model="form.sort" :min="0" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleDialogOk">确定</el-button>
      </template>
    </el-dialog>
    <!-- 查看弹窗 -->
    <el-dialog title="查看部门" v-model="viewDialogVisible" width="500px" :show-close="true">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="部门名称">{{ viewRow.departmentName }}</el-descriptions-item>
        <el-descriptions-item label="负责人">{{ viewRow.leader }}</el-descriptions-item>
        <el-descriptions-item label="电话">{{ viewRow.phone }}</el-descriptions-item>
        <el-descriptions-item label="邮箱">{{ viewRow.email }}</el-descriptions-item>
        <el-descriptions-item label="状态">{{ viewRow.status }}</el-descriptions-item>
        <el-descriptions-item label="排序">{{ viewRow.sort }}</el-descriptions-item>
        <el-descriptions-item label="创建时间">{{ viewRow.createdAt }}</el-descriptions-item>
      </el-descriptions>
      <template #footer>
        <el-button @click="viewDialogVisible = false">关闭</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted, reactive } from 'vue'
import { getWlDepartmentList, addWlDepartment, updateWlDepartment, deleteWlDepartment } from '@/api/wlDepartment'
import { ElMessage, ElMessageBox } from 'element-plus'

const depList = ref([])
const total = ref(0)
const searchForm = ref({
  page: 1,
  pageSize: 10,
  departmentName: ''
})

const fetchData = async () => {
  const { data } = await getWlDepartmentList(searchForm.value)
  depList.value = data.list || []
  total.value = data.total || 0
}

const handleSearch = () => {
  searchForm.value.page = 1
  fetchData()
}

const handleReset = () => {
  searchForm.value = {
    page: 1,
    pageSize: 10,
    departmentName: ''
  }
  fetchData()
}

const handleSizeChange = (size) => {
  searchForm.value.pageSize = size
  fetchData()
}

const handlePageChange = (page) => {
  searchForm.value.page = page
  fetchData()
}

// 弹窗相关
const dialogVisible = ref(false)
const dialogTitle = ref('')
const form = reactive({
  id: undefined,
  departmentName: '',
  leader: '',
  phone: '',
  email: '',
  status: '启用',
  sort: 0
})
const rules = {
  departmentName: [{ required: true, message: '请输入部门名称', trigger: 'blur' }]
}
const formRef = ref()

const handleAdd = () => {
  dialogTitle.value = '新增部门'
  Object.assign(form, { id: undefined, departmentName: '', leader: '', phone: '', email: '', status: '启用', sort: 0 })
  dialogVisible.value = true
}
const handleEdit = (row) => {
  dialogTitle.value = '编辑部门'
  Object.assign(form, row)
  dialogVisible.value = true
}
const handleDialogOk = () => {
  formRef.value.validate(async (valid) => {
    if (!valid) return
    if (form.id) {
      await updateWlDepartment(form)
      ElMessage.success('编辑成功')
    } else {
      await addWlDepartment(form)
      ElMessage.success('新增成功')
    }
    dialogVisible.value = false
    fetchData()
  })
}
const handleDelete = (row) => {
  ElMessageBox.confirm('确定要删除该部门吗？', '提示', { type: 'warning' })
    .then(async () => {
      await deleteWlDepartment({ id: row.id })
      ElMessage.success('删除成功')
      fetchData()
    })
}
// 查看弹窗
const viewDialogVisible = ref(false)
const viewRow = ref({})
const handleView = (row) => {
  viewRow.value = { ...row }
  viewDialogVisible.value = true
}

onMounted(fetchData)
</script> 