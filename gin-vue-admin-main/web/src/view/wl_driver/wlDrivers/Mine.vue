<template>
  <div class="mine-driver-container">
    <!-- 顶部搜索栏 -->
    <el-card class="search-card">
      <el-form :inline="true" :model="searchForm" class="search-form">
        <el-form-item label="驱动名称">
          <el-input v-model="searchForm.name" placeholder="请输入驱动名称" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
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
        <el-table-column prop="id" label="驱动编号" width="110" align="center" />
        <el-table-column prop="name" label="驱动名称" min-width="180" />
        <el-table-column prop="version" label="版本" width="90" align="center" />
        <el-table-column prop="type" label="驱动类型" width="100" align="center">
          <template #default="scope">
            <el-tag :type="scope.row.type === '官方' ? 'success' : 'info'" effect="light">
              {{ scope.row.type }}
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
        <el-table-column prop="createdAt" label="创建时间" min-width="170" align="center" />
        <el-table-column label="操作" width="180" align="center">
          <template #default="scope">
            <el-link type="primary" @click="handleOperate(scope.row)">操作 √</el-link>
            <span style="color:#dcdcdc;"> / </span>
            <el-link type="primary" @click="handleLog(scope.row)">日志</el-link>
            <span style="color:#dcdcdc;"> / </span>
            <el-link type="primary" @click="handleDelete(scope.row)">删除</el-link>
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
  </div>
</template>

<script setup>
import { ref } from 'vue'

// 搜索表单
const searchForm = ref({
  name: ''
})

// 表格数据（可替换为接口数据）
const tableData = ref([
  { id: '31747985', name: 'GB281协议驱动-31747985', version: '2.7', type: '官方', status: '停止', createdAt: '2025-07-22 10:22:40' },
  { id: '98017279', name: 'MQTT协议驱动-98017279', version: '2.7', type: '官方', status: '运行中', createdAt: '2025-07-22 10:22:40' },
  { id: '48654311', name: 'mqtt-ca-48654311', version: '2.0-ca', type: '自定义', status: '停止', createdAt: '2025-07-22 10:22:40' },
  { id: '30593142', name: 'mqtt试驱动2.7版本-30593142', version: '2.7', type: '自定义', status: '停止', createdAt: '2025-07-22 10:22:40' },
  { id: '53703706', name: 'TCP协议驱动-53703706', version: '2.7', type: '官方', status: '停止', createdAt: '2025-07-22 10:22:40' },
  { id: '96475139', name: 'HTTP协议驱动-96475139', version: '2.7', type: '官方', status: '停止', createdAt: '2025-07-22 10:22:40' },
  { id: '82218226', name: 'MODBUS TCP协议驱动-82218226', version: '2.7', type: '官方', status: '停止', createdAt: '2025-07-22 10:22:40' },
  { id: '48849713', name: 'rtu驱动-48849713', version: '2.0', type: '自定义', status: '停止', createdAt: '2025-07-22 10:22:40' },
])

// 分页相关
const total = ref(8)
const pageSize = ref(10)
const currentPage = ref(1)

const handleSearch = () => {
  // TODO: 搜索逻辑
}
const handleReset = () => {
  searchForm.value.name = ''
  // TODO: 重置逻辑
}
const handleOperate = (row) => {
  // TODO: 操作逻辑
}
const handleLog = (row) => {
  // TODO: 日志逻辑
}
const handleDelete = (row) => {
  // TODO: 删除逻辑
}
const handleSizeChange = (size) => {
  pageSize.value = size
  // TODO: 分页逻辑
}
const handlePageChange = (page) => {
  currentPage.value = page
  // TODO: 分页逻辑
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