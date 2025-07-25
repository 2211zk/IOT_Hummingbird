<template>
  <div>
    <el-card>
      <div style="margin-bottom: 16px;">
        <el-form :inline="true" :model="searchForm">
          <el-form-item label="用户名">
            <el-input v-model="searchForm.userName" placeholder="请输入用户名" clearable />
          </el-form-item>
          <el-form-item label="手机号">
            <el-input v-model="searchForm.mobile" placeholder="请输入手机号" clearable />
          </el-form-item>
          <el-form-item label="状态">
            <el-select v-model="searchForm.status" placeholder="请选择状态" clearable>
              <el-option label="正常" value="正常" />
              <el-option label="禁用" value="禁用" />
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="handleSearch">查询</el-button>
            <el-button @click="handleReset">重置</el-button>
          </el-form-item>
        </el-form>
      </div>
      <el-table :data="userList" border style="width: 100%;">
        <el-table-column prop="userName" label="用户名" />
        <el-table-column prop="userNickname" label="昵称" />
        <el-table-column prop="mobile" label="手机" />
        <el-table-column prop="department" label="部门" />
        <el-table-column prop="userStatus" label="状态" />
        <el-table-column prop="creationTime" label="创建时间" />
        <el-table-column label="操作">
          <template #default="scope">
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
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getWlUserList } from '@/api/wlUser'

const userList = ref([])
const total = ref(0)
const searchForm = ref({
  page: 1,
  pageSize: 10,
  userName: '',
  mobile: '',
  status: '',
  department: ''
})

const fetchData = async () => {
  const { data } = await getWlUserList(searchForm.value)
  userList.value = data.list || []
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
    userName: '',
    mobile: '',
    status: '',
    department: ''
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

const handleEdit = (row) => {
  // TODO: 打开编辑弹窗
  alert('编辑功能待实现: ' + row.userName)
}

const handleDelete = (row) => {
  // TODO: 删除逻辑
  alert('删除功能待实现: ' + row.userName)
}

onMounted(fetchData)
</script> 