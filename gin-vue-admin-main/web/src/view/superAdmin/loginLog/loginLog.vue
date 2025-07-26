<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="用户名">
          <el-input v-model="searchInfo.userName" placeholder="请输入用户名" clearable />
        </el-form-item>
        <el-form-item label="登录IP">
          <el-input v-model="searchInfo.loginAddress" placeholder="请输入登录IP" clearable />
        </el-form-item>
        <el-form-item label="登录地点">
          <el-input v-model="searchInfo.loginLocation" placeholder="请输入登录地点" clearable />
        </el-form-item>
        <el-form-item label="登录状态">
          <el-select v-model="searchInfo.loginStatus" placeholder="请选择登录状态" clearable>
            <el-option label="成功" value="成功" />
            <el-option label="失败" value="失败" />
          </el-select>
        </el-form-item>
        <el-form-item label="时间范围">
          <el-date-picker
            v-model="timeRange"
            type="datetimerange"
            range-separator="至"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
            format="YYYY-MM-DD HH:mm:ss"
            value-format="YYYY-MM-DD HH:mm:ss"
            @change="handleTimeRangeChange"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="Search" @click="onSubmit">查询</el-button>
          <el-button icon="Refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="Download" @click="exportData">导出Excel</el-button>
        <el-button type="warning" icon="Delete" @click="showCleanDialog">清理日志</el-button>
        <el-button type="info" icon="DataAnalysis" @click="showStatistics">统计信息</el-button>
        <el-button 
          type="danger" 
          icon="Delete" 
          :disabled="multipleSelection.length === 0"
          @click="batchDelete"
        >
          批量删除 ({{ multipleSelection.length }})
        </el-button>
      </div>
      
      <el-table 
        ref="multipleTable"
        :data="tableData" 
        @selection-change="handleSelectionChange"
        row-key="ID"
        v-loading="loading"
        :row-class-name="tableRowClassName"
        :empty-text="tableData.length === 0 && !loading ? '暂无登录日志数据' : ''"
        style="width: 100%"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="访问编号" min-width="100">
          <template #default="scope">
            {{ scope.row.accessNumber || scope.row.ID || '-' }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="用户名称" min-width="120" prop="userName" />
        <el-table-column align="left" label="登录地址" min-width="140" prop="loginAddress" />
        <el-table-column align="left" label="登录地点" min-width="150" prop="loginLocation" />
        <el-table-column align="left" label="浏览器" min-width="120" prop="browser" />
        <el-table-column align="left" label="操作系统" min-width="120" prop="operatingSystem" />
        <el-table-column align="left" label="登录状态" min-width="100">
          <template #default="scope">
            <el-tag :type="scope.row.loginStatus === '成功' ? 'success' : 'danger'">
              {{ scope.row.loginStatus }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作信息" min-width="150" prop="operationalInfo" />
        <el-table-column align="left" label="登录时间" min-width="180">
          <template #default="scope">
            {{ formatTime(scope.row.loginTime) }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="操作" min-width="100" fixed="right">
          <template #default="scope">
            <el-button type="primary" size="small" @click="showDetail(scope.row)">
              详情
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <!-- 详情弹窗 -->
    <el-dialog v-model="detailVisible" title="登录日志详情" width="700px">
      <el-descriptions :column="2" border>
        <el-descriptions-item label="访问编号">{{ currentDetail.accessNumber || '-' }}</el-descriptions-item>
        <el-descriptions-item label="用户名称">{{ currentDetail.userName || '-' }}</el-descriptions-item>
        <el-descriptions-item label="登录地址">{{ currentDetail.loginAddress || '-' }}</el-descriptions-item>
        <el-descriptions-item label="登录地点">{{ currentDetail.loginLocation || '-' }}</el-descriptions-item>
        <el-descriptions-item label="浏览器">{{ currentDetail.browser || '-' }}</el-descriptions-item>
        <el-descriptions-item label="操作系统">{{ currentDetail.operatingSystem || '-' }}</el-descriptions-item>
        <el-descriptions-item label="登录状态">
          <el-tag :type="currentDetail.loginStatus === '成功' ? 'success' : 'danger'">
            {{ currentDetail.loginStatus || '-' }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="登录时间">{{ formatTime(currentDetail.loginTime) }}</el-descriptions-item>
        <el-descriptions-item label="操作信息" :span="2">{{ currentDetail.operationalInfo || '-' }}</el-descriptions-item>
        <el-descriptions-item label="记录ID" :span="2">{{ currentDetail.ID || currentDetail.id || '-' }}</el-descriptions-item>
      </el-descriptions>
      <template #footer>
        <el-button @click="detailVisible = false">关闭</el-button>
      </template>
    </el-dialog>

    <!-- 清理日志弹窗 -->
    <el-dialog v-model="cleanVisible" title="清理过期日志" width="500px">
      <el-form :model="cleanForm" label-width="120px">
        <el-form-item label="保留天数">
          <el-input-number v-model="cleanForm.days" :min="1" :max="365" />
          <div class="el-form-item__description">将删除 {{ cleanForm.days }} 天前的所有登录日志</div>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="cleanVisible = false">取消</el-button>
        <el-button type="danger" @click="confirmClean">确认清理</el-button>
      </template>
    </el-dialog>

    <!-- 统计信息弹窗 -->
    <el-dialog v-model="statisticsVisible" title="登录统计信息" width="800px">
      <el-row :gutter="20">
        <el-col :span="6">
          <el-statistic title="总登录次数" :value="statistics.totalLogins" />
        </el-col>
        <el-col :span="6">
          <el-statistic title="成功登录" :value="statistics.successLogins" />
        </el-col>
        <el-col :span="6">
          <el-statistic title="失败登录" :value="statistics.failedLogins" />
        </el-col>
        <el-col :span="6">
          <el-statistic title="独立用户" :value="statistics.uniqueUsers" />
        </el-col>
      </el-row>
      <el-divider />
      <el-row>
        <el-col :span="12">
          <h4>热门登录IP</h4>
          <el-table :data="topIPs" size="small">
            <el-table-column prop="login_address" label="IP地址" />
            <el-table-column prop="login_location" label="地理位置" />
            <el-table-column prop="login_count" label="登录次数" />
          </el-table>
        </el-col>
        <el-col :span="12">
          <h4>最近登录记录</h4>
          <el-table :data="recentLogs" size="small">
            <el-table-column prop="userName" label="用户名" />
            <el-table-column prop="loginAddress" label="IP地址" />
            <el-table-column prop="loginStatus" label="状态">
              <template #default="scope">
                <el-tag :type="scope.row.loginStatus === '成功' ? 'success' : 'danger'" size="small">
                  {{ scope.row.loginStatus }}
                </el-tag>
              </template>
            </el-table-column>
          </el-table>
        </el-col>
      </el-row>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { 
  getLoginLogList, 
  getLoginLogDetail, 
  exportLoginLog, 
  cleanExpiredLogs,
  getLoginStatistics,
  getTopLoginIPs,
  getRecentLoginLogs,
  deleteLoginLogByIds
} from '@/api/loginLog'

// 响应式数据
const tableData = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const timeRange = ref([])
const loading = ref(false)

const searchInfo = reactive({
  userName: '',
  loginAddress: '',
  loginLocation: '',
  loginStatus: '',
  startTime: '',
  endTime: ''
})

const detailVisible = ref(false)
const currentDetail = ref({})

const cleanVisible = ref(false)
const cleanForm = reactive({
  days: 90
})

const statisticsVisible = ref(false)
const statistics = ref({})
const topIPs = ref([])
const recentLogs = ref([])

// 多选相关
const multipleSelection = ref([])
const multipleTable = ref()

// 方法
const getTableData = async () => {
  loading.value = true
  try {
    const table = await getLoginLogList({
      page: page.value,
      pageSize: pageSize.value,
      ...searchInfo
    })
    if (table.code === 0) {
      tableData.value = table.data.list || []
      total.value = table.data.total || 0
      page.value = table.data.page || 1
      pageSize.value = table.data.pageSize || 10
      
      // 调试信息：检查数据结构
      if (tableData.value.length > 0) {
        console.log('登录日志数据示例:', tableData.value[0])
      }
    } else {
      ElMessage.error(table.msg || '获取数据失败')
      tableData.value = []
      total.value = 0
    }
  } catch (error) {
    console.error('获取登录日志失败:', error)
    ElMessage.error('获取数据失败')
    tableData.value = []
    total.value = 0
  } finally {
    loading.value = false
  }
}

const onSubmit = () => {
  page.value = 1
  getTableData()
}

const onReset = () => {
  Object.keys(searchInfo).forEach(key => {
    searchInfo[key] = ''
  })
  timeRange.value = []
  page.value = 1
  getTableData()
}

const handleTimeRangeChange = (val) => {
  if (val && val.length === 2) {
    searchInfo.startTime = val[0]
    searchInfo.endTime = val[1]
  } else {
    searchInfo.startTime = ''
    searchInfo.endTime = ''
  }
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const showDetail = async (row) => {
  // 确保row有ID字段，如果没有则使用id字段
  const id = row.ID || row.id
  if (!id) {
    ElMessage.error('无法获取日志详情：ID不存在')
    return
  }
  
  try {
    const res = await getLoginLogDetail(id)
    if (res.code === 0) {
      currentDetail.value = res.data.loginLog
      detailVisible.value = true
    } else {
      ElMessage.error(res.msg || '获取详情失败')
    }
  } catch (error) {
    ElMessage.error('获取详情失败')
  }
}

// 处理多选变化
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 批量删除
const batchDelete = async () => {
  if (multipleSelection.value.length === 0) {
    ElMessage.warning('请选择要删除的记录')
    return
  }

  try {
    await ElMessageBox.confirm(
      `确定要删除选中的 ${multipleSelection.value.length} 条记录吗？此操作不可恢复！`,
      '警告',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      }
    )
    
    const ids = multipleSelection.value.map(item => item.ID || item.id)
    const res = await deleteLoginLogByIds({ IDs: ids })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      getTableData()
      multipleSelection.value = []
    } else {
      ElMessage.error(res.msg || '删除失败')
    }
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('删除失败')
    }
  }
}

const exportData = async () => {
  try {
    const res = await exportLoginLog({
      ...searchInfo,
      page: 1,
      pageSize: 999999
    })
    
    const blob = new Blob([res], { 
      type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet' 
    })
    const url = window.URL.createObjectURL(blob)
    const link = document.createElement('a')
    link.href = url
    link.download = `登录日志_${new Date().toISOString().slice(0, 10)}.xlsx`
    link.click()
    window.URL.revokeObjectURL(url)
    
    ElMessage.success('导出成功')
  } catch (error) {
    ElMessage.error('导出失败')
  }
}

const showCleanDialog = () => {
  cleanVisible.value = true
}

const confirmClean = async () => {
  try {
    await ElMessageBox.confirm(
      `确定要清理 ${cleanForm.days} 天前的登录日志吗？此操作不可恢复！`,
      '警告',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      }
    )
    
    const res = await cleanExpiredLogs({ days: cleanForm.days })
    if (res.code === 0) {
      ElMessage.success(`成功清理了 ${res.data.deletedCount} 条记录`)
      cleanVisible.value = false
      getTableData()
    }
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('清理失败')
    }
  }
}

const showStatistics = async () => {
  try {
    const [statsRes, ipsRes, logsRes] = await Promise.all([
      getLoginStatistics(7),
      getTopLoginIPs(10, 7),
      getRecentLoginLogs(24, 10)
    ])
    
    if (statsRes.code === 0) {
      statistics.value = statsRes.data
    }
    if (ipsRes.code === 0) {
      topIPs.value = ipsRes.data
    }
    if (logsRes.code === 0) {
      recentLogs.value = logsRes.data
    }
    
    statisticsVisible.value = true
  } catch (error) {
    ElMessage.error('获取统计信息失败')
  }
}

const formatTime = (time) => {
  if (!time) return ''
  return new Date(time).toLocaleString('zh-CN')
}

const tableRowClassName = ({ row }) => {
  if (row.loginStatus === '成功') {
    return 'success-row'
  } else if (row.loginStatus === '失败') {
    return 'error-row'
  }
  return ''
}

// 生命周期
onMounted(() => {
  getTableData()
})
</script>

<style scoped>
.el-form-item__description {
  font-size: 12px;
  color: #909399;
  margin-top: 4px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .gva-search-box .el-form {
    flex-direction: column;
  }
  
  .gva-search-box .el-form-item {
    width: 100%;
    margin-right: 0;
  }
  
  .gva-btn-list {
    flex-direction: column;
    gap: 8px;
  }
  
  .gva-btn-list .el-button {
    width: 100%;
  }
  
  .el-table {
    font-size: 12px;
  }
  
  .el-table .el-table__cell {
    padding: 8px 4px;
  }
}

@media (max-width: 480px) {
  .gva-table-box {
    overflow-x: auto;
  }
  
  .el-table {
    min-width: 800px;
  }
  
  .el-dialog {
    width: 95% !important;
    margin: 0 auto;
  }
  
  .el-descriptions {
    font-size: 12px;
  }
}

/* 表格样式优化 */
.el-table .el-table__row:hover {
  background-color: #f5f7fa;
  cursor: pointer;
}

.el-table .success-row {
  background-color: #f0f9ff;
}

.el-table .error-row {
  background-color: #fef2f2;
}

/* 统计卡片样式 */
.el-statistic {
  text-align: center;
  padding: 20px;
  border: 1px solid #e4e7ed;
  border-radius: 4px;
  background-color: #fff;
}

.el-statistic .el-statistic__head {
  font-size: 14px;
  color: #606266;
  margin-bottom: 8px;
}

.el-statistic .el-statistic__content {
  font-size: 24px;
  font-weight: bold;
  color: #303133;
}

/* 加载状态 */
.loading-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 200px;
}

/* 空数据状态 */
.empty-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 200px;
  color: #909399;
}

.empty-container .empty-icon {
  font-size: 48px;
  margin-bottom: 16px;
}

.empty-container .empty-text {
  font-size: 14px;
}

/* 按钮组样式 */
.gva-btn-list {
  display: flex;
  gap: 12px;
  margin-bottom: 16px;
  flex-wrap: wrap;
}

/* 搜索框样式 */
.gva-search-box {
  background-color: #fff;
  padding: 20px;
  border-radius: 4px;
  margin-bottom: 16px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.gva-search-box .el-form-item {
  margin-bottom: 16px;
}

/* 表格容器样式 */
.gva-table-box {
  background-color: #fff;
  border-radius: 4px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  overflow: hidden;
}

/* 分页样式 */
.gva-pagination {
  padding: 20px;
  display: flex;
  justify-content: flex-end;
  background-color: #fafafa;
  border-top: 1px solid #e4e7ed;
}

/* 标签样式优化 */
.el-tag {
  font-weight: 500;
}

.el-tag.el-tag--success {
  background-color: #f0f9ff;
  border-color: #bfdbfe;
  color: #1e40af;
}

.el-tag.el-tag--danger {
  background-color: #fef2f2;
  border-color: #fecaca;
  color: #dc2626;
}

/* 对话框样式优化 */
.el-dialog__header {
  background-color: #f8f9fa;
  border-bottom: 1px solid #e9ecef;
}

.el-dialog__body {
  padding: 24px;
}

/* 描述列表样式 */
.el-descriptions {
  margin-top: 16px;
}

.el-descriptions .el-descriptions__label {
  font-weight: 600;
  color: #606266;
}

.el-descriptions .el-descriptions__content {
  color: #303133;
}

/* 表单样式优化 */
.el-form-item__label {
  font-weight: 500;
  color: #606266;
}

.el-input, .el-select, .el-date-editor {
  width: 100%;
}

/* 动画效果 */
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.3s;
}

.fade-enter-from, .fade-leave-to {
  opacity: 0;
}

/* 滚动条样式 */
::-webkit-scrollbar {
  width: 6px;
  height: 6px;
}

::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 3px;
}

::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 3px;
}

::-webkit-scrollbar-thumb:hover {
  background: #a8a8a8;
}
</style>