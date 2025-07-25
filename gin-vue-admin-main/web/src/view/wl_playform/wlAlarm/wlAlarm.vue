
<template>
  <div>
    <!-- 统计图和统计区块 -->
    <div class="alarm-statistics-wrapper">
      <div class="alarm-pie-chart">
        <v-chart :option="alarmPieOption" autoresize style="height: 180px; width: 220px;" />
      </div>
      <div class="alarm-statistics-info">
        <div class="alarm-stat-title">告警数（7日）</div>
        <div class="alarm-stat-total">{{ alarmTotal }}</div>
        <div class="alarm-stat-legend">
          <div class="alarm-stat-legend-item" v-for="item in alarmPieData" :key="item.name">
            <span class="alarm-dot" :style="{ background: item.color }"></span>
            <span class="alarm-type">{{ item.name }}</span>
            <span class="alarm-percent">{{ item.percent }}%</span>
            <span class="alarm-count">{{ item.value }}</span>
          </div>
        </div>
      </div>
    </div>
    <!-- 新增按钮 -->
    <div style="margin-bottom: 12px;">
      <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
    </div>
    <!-- 查询区 -->
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAtRange">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>

      <el-date-picker
            v-model="searchInfo.createdAtRange"
            class="w-[380px]"
            type="datetimerange"
            range-separator="至"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
          />
       </el-form-item>
      

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <!-- 卡片式告警列表 -->
    <div class="alarm-card-list">
      <div v-if="tableData.length === 0" class="alarm-card-empty">暂无数据</div>
      <div v-for="item in tableData" :key="item.id" class="alarm-card">
        <div class="alarm-card-level" :class="'level-' + getLevelClass(item)">
          {{ filterDict(item.alarmLevel, alarm_levelOptions) }}
        </div>
        <div class="alarm-card-main">
          <div class="alarm-card-header">
            <span class="alarm-card-title">{{ item.alarmType || '未知类型' }}</span>
            <span class="alarm-card-id">ID: {{ item.id }}</span>
            <span class="alarm-card-status">{{ filterDict(item.alarmStatus, alarm_statusOptions) }}</span>
          </div>
          <div class="alarm-card-desc">{{ item.alarmContent || '无描述' }}</div>
          <div class="alarm-card-meta">
            <span>规则名称：{{ item.ruleName || '-' }}</span>
            <span>触发时间：{{ formatDate(item.createTime) }}</span>
            <span>处理时间：{{ formatDate(item.handleTime) }}</span>
            <span>处理结果：{{ item.handleRemark || '-' }}</span>
          </div>
        </div>
        <div class="alarm-card-actions">
          <el-button type="primary" link class="alarm-btn" @click="openProcessDialog(item)">处理</el-button>
          <el-button type="primary" link class="alarm-btn" @click="ignoreAlarm(item)">忽略</el-button>
          <el-button type="primary" link class="alarm-btn" @click="定位Func(item)">数据定位</el-button>
        </div>
      </div>
    </div>
    <!-- 分页 -->
    <div class="gva-pagination">
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
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="设备ID:" prop="deviceId">
    <el-input v-model.number="formData.deviceId" :clearable="true" placeholder="请输入设备ID" />
</el-form-item>
            <el-form-item label="告警类型:" prop="alarmType">
    <el-input v-model="formData.alarmType" :clearable="true" placeholder="请输入告警类型" />
</el-form-item>
            <el-form-item label="告警级别:" prop="alarmLevel">
    <el-select v-model="formData.alarmLevel" placeholder="请选择告警级别" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in alarm_levelOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
            <el-form-item label="告警状态:" prop="alarmStatus">
    <el-select v-model="formData.alarmStatus" placeholder="请选择告警状态" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in alarm_statusOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
            <el-form-item label="告警内容描述:" prop="alarmContent">
    <el-input v-model="formData.alarmContent" :clearable="true" placeholder="请输入告警内容描述" />
</el-form-item>
            <el-form-item label="告警相关数据:" prop="alarmData">
    <el-input v-model="formData.alarmData" :clearable="true" placeholder="请输入告警相关数据" />
</el-form-item>
            <el-form-item label="告警创建时间:" prop="createTime">
    <el-date-picker v-model="formData.createTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
            <el-form-item label="告警更新时间:" prop="updateTime">
    <el-date-picker v-model="formData.updateTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
            <el-form-item label="处理时间:" prop="handleTime">
    <el-date-picker v-model="formData.handleTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
            <el-form-item label="处理人:" prop="handleUser">
    <el-input v-model="formData.handleUser" :clearable="true" placeholder="请输入处理人" />
</el-form-item>
            <el-form-item label="处理备注:" prop="handleRemark">
    <el-input v-model="formData.handleRemark" :clearable="true" placeholder="请输入处理备注" />
</el-form-item>
            <el-form-item label="创建者:" prop="createdBy">
    <el-input v-model.number="formData.createdBy" :clearable="true" placeholder="请输入创建者" />
</el-form-item>
            <el-form-item label="更新者:" prop="updatedBy">
    <el-input v-model.number="formData.updatedBy" :clearable="true" placeholder="请输入更新者" />
</el-form-item>
            <el-form-item label="删除者:" prop="deletedBy">
    <el-input v-model.number="formData.deletedBy" :clearable="true" placeholder="请输入删除者" />
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="设备ID">
    {{ detailFrom.deviceId }}
</el-descriptions-item>
                    <el-descriptions-item label="告警类型">
    {{ detailFrom.alarmType }}
</el-descriptions-item>
                    <el-descriptions-item label="告警级别">
    {{ detailFrom.alarmLevel }}
</el-descriptions-item>
                    <el-descriptions-item label="告警状态">
    {{ detailFrom.alarmStatus }}
</el-descriptions-item>
                    <el-descriptions-item label="告警内容描述">
    {{ detailFrom.alarmContent }}
</el-descriptions-item>
                    <el-descriptions-item label="告警相关数据">
    {{ detailFrom.alarmData }}
</el-descriptions-item>
                    <el-descriptions-item label="告警创建时间">
    {{ detailFrom.createTime }}
</el-descriptions-item>
                    <el-descriptions-item label="告警更新时间">
    {{ detailFrom.updateTime }}
</el-descriptions-item>
                    <el-descriptions-item label="处理时间">
    {{ detailFrom.handleTime }}
</el-descriptions-item>
                    <el-descriptions-item label="处理人">
    {{ detailFrom.handleUser }}
</el-descriptions-item>
                    <el-descriptions-item label="处理备注">
    {{ detailFrom.handleRemark }}
</el-descriptions-item>
                    <el-descriptions-item label="创建者">
    {{ detailFrom.createdBy }}
</el-descriptions-item>
                    <el-descriptions-item label="更新者">
    {{ detailFrom.updatedBy }}
</el-descriptions-item>
                    <el-descriptions-item label="删除者">
    {{ detailFrom.deletedBy }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

    <!-- 处理弹窗 -->
    <el-dialog v-model="processDialogVisible" title="处理告警" width="400px" :before-close="closeProcessDialog">
      <el-form :model="processForm" label-width="80px">
        <el-form-item label="处理结果">
          <el-input v-model="processForm.handleRemark" type="textarea" placeholder="请输入处理结果" rows="3" />
        </el-form-item>
        <el-form-item label="处理人">
          <el-input v-model="processForm.handleUser" placeholder="请输入处理人" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="closeProcessDialog">取消</el-button>
        <el-button type="primary" @click="confirmProcess">确定</el-button>
      </template>
    </el-dialog>

  </div>
</template>

<script setup>
import {
  createWlAlarm,
  deleteWlAlarm,
  deleteWlAlarmByIds,
  updateWlAlarm,
  findWlAlarm,
  getWlAlarmList
} from '@/api/wl_playform/wlAlarm'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, computed, watch } from 'vue'
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'
import { useAppStore } from "@/pinia"
import VChart from 'vue-echarts'
import { use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { PieChart } from 'echarts/charts'
import { TooltipComponent, LegendComponent } from 'echarts/components'
import { useRouter } from 'vue-router'
const router = useRouter()

use([CanvasRenderer, PieChart, TooltipComponent, LegendComponent])

// 告警级别与颜色映射（低、中、高、紧急）
const alarmLevelMap = [
  { name: '低', value: 'low', color: '#67C23A' },
  { name: '中', value: 'medium', color: '#409EFF' },
  { name: '高', value: 'high', color: '#e6a23c' },
  { name: '紧急', value: 'urgent', color: '#f56c6c' }
]

// 统计当前页各类告警数量
const alarmPieData = computed(() => {
  const counts = { low: 0, medium: 0, high: 0, urgent: 0 }
  tableData.value.forEach(item => {
    const levelName = filterDict(item.alarmLevel, alarm_levelOptions.value)
    const level = alarmLevelMap.find(l => l.name === levelName)
    if (level) counts[level.value]++
  })
  const total = Object.values(counts).reduce((a, b) => a + b, 0)
  return alarmLevelMap.map(l => ({
    name: l.name,
    value: counts[l.value],
    color: l.color,
    percent: total ? ((counts[l.value] / total) * 100).toFixed(2) : '0.00'
  }))
})

const alarmTotal = computed(() => alarmPieData.value.reduce((a, b) => a + b.value, 0))

// ECharts 配置
const alarmPieOption = computed(() => ({
  tooltip: {
    trigger: 'item',
    formatter: '{b}: {c} ({d}%)'
  },
  legend: {
    show: false
  },
  series: [
    {
      name: '告警级别',
      type: 'pie',
      radius: ['60%', '80%'],
      avoidLabelOverlap: false,
      itemStyle: {
        borderRadius: 8,
        borderColor: '#fff',
        borderWidth: 2
      },
      label: {
        show: false
      },
      data: alarmPieData.value.map(item => ({ value: item.value, name: item.name, itemStyle: { color: item.color } }))
    }
  ]
}))

defineOptions({
    name: 'WlAlarm'
})
// 按钮权限实例化
    const btnAuth = useBtnAuth()

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const alarm_statusOptions = ref([])
const alarm_levelOptions = ref([])
const formData = ref({
            deviceId: undefined,
            alarmType: '',
            alarmLevel: '',
            alarmStatus: '',
            alarmContent: '',
            alarmData: '',
            createTime: new Date(),
            updateTime: new Date(),
            handleTime: new Date(),
            handleUser: '',
            handleRemark: '',
            createdBy: undefined,
            updatedBy: undefined,
            deletedBy: undefined,
        })



// 验证规则
const rule = reactive({
               deviceId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
    CreatedAt:"CreatedAt",
    ID:"ID",
            createTime: 'create_time',
  }

  let sort = sortMap[prop]
  if(!sort){
   sort = prop.replace(/[A-Z]/g, match => `_${match.toLowerCase()}`)
  }

  searchInfo.value.sort = sort
  searchInfo.value.order = order
  getTableData()
}
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getWlAlarmList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    alarm_statusOptions.value = await getDictFunc('alarm_status')
    alarm_levelOptions.value = await getDictFunc('alarm_level')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteWlAlarmFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteWlAlarmByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateWlAlarmFunc = async(row) => {
    const res = await findWlAlarm({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteWlAlarmFunc = async (row) => {
    const res = await deleteWlAlarm({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        deviceId: undefined,
        alarmType: '',
        alarmLevel: '',
        alarmStatus: '',
        alarmContent: '',
        alarmData: '',
        createTime: new Date(),
        updateTime: new Date(),
        handleTime: new Date(),
        handleUser: '',
        handleRemark: '',
        createdBy: undefined,
        updatedBy: undefined,
        deletedBy: undefined,
        }
}
// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              let res
              switch (type.value) {
                case 'create':
                  res = await createWlAlarm(formData.value)
                  break
                case 'update':
                  res = await updateWlAlarm(formData.value)
                  break
                default:
                  res = await createWlAlarm(formData.value)
                  break
              }
              btnLoading.value = false
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}

const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findWlAlarm({ ID: row.ID })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}

// 获取告警等级class
const getLevelClass = (item) => {
  const levelName = filterDict(item.alarmLevel, alarm_levelOptions.value)
  switch (levelName) {
    case '低': return 'low'
    case '中': return 'medium'
    case '高': return 'high'
    case '紧急': return 'urgent'
    default: return 'low'
  }
}

// 处理弹窗相关
const processDialogVisible = ref(false)
const processForm = ref({ id: null, handleRemark: '', handleUser: '' })
let processRow = null

const openProcessDialog = (row) => {
  processRow = row
  processForm.value = {
    id: row.id,
    handleRemark: row.handleRemark || '',
    handleUser: row.handleUser || ''
  }
  processDialogVisible.value = true
}
const closeProcessDialog = () => {
  processDialogVisible.value = false
  processForm.value = { id: null, handleRemark: '', handleUser: '' }
  processRow = null
}
const confirmProcess = async () => {
  if (!processRow) return
  const updateData = {
    ...processRow,
    alarmStatus: '2', // 用字符串类型，假设'2'为“已处理”
    handleRemark: processForm.value.handleRemark,
    handleUser: processForm.value.handleUser,
    handleTime: new Date()
  }
  const res = await updateWlAlarm(updateData)
  if (res.code === 0) {
    ElMessage.success('处理成功')
    closeProcessDialog()
    getTableData()
  }
}

const ignoreAlarm = async (row) => {
  const updateData = {
    ...row,
    alarmStatus: '1' // 假设'1'为“忽略”状态
    // 不修改handleRemark
  }
  const res = await updateWlAlarm(updateData)
  if (res.code === 0) {
    ElMessage.success('已忽略')
    getTableData()
  }
}

const 定位Func = (row) => {
  router.push({
    name: 'wlProducts',
    query: { id: row.productsId, showLocateDialog: 1 }
  })
}

</script>

<style scoped>
.alarm-statistics-wrapper {
  display: flex;
  align-items: center;
  background: #fff;
  border-radius: 10px;
  box-shadow: 0 2px 8px #e3e6f0;
  padding: 18px 32px 18px 18px;
  margin-bottom: 18px;
  gap: 32px;
}
.alarm-pie-chart {
  min-width: 220px;
  display: flex;
  align-items: center;
  justify-content: center;
}
.alarm-statistics-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 8px;
}
.alarm-stat-title {
  font-size: 18px;
  font-weight: bold;
  color: #333;
  margin-bottom: 8px;
}
.alarm-stat-total {
  font-size: 32px;
  font-weight: bold;
  color: #409EFF;
  margin-bottom: 8px;
}
.alarm-stat-legend {
  display: flex;
  gap: 32px;
  margin-top: 8px;
}
.alarm-stat-legend-item {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 15px;
}
.alarm-dot {
  display: inline-block;
  width: 12px;
  height: 12px;
  border-radius: 50%;
  margin-right: 4px;
}
.alarm-type {
  min-width: 36px;
  font-weight: 500;
}
.alarm-percent {
  color: #888;
  min-width: 48px;
}
.alarm-count {
  color: #333;
  min-width: 24px;
  font-weight: bold;
}
.alarm-card-list {
  display: flex;
  flex-direction: column;
  gap: 18px;
  margin: 18px 0;
}
.alarm-card-empty {
  text-align: center;
  color: #bbb;
  font-size: 18px;
  padding: 40px 0;
}
.alarm-card {
  display: flex;
  align-items: flex-start;
  background: #fff;
  border-radius: 10px;
  box-shadow: 0 2px 8px #e3e6f0;
  padding: 18px 24px;
  gap: 24px;
  border-left: 6px solid #e3e6f0;
  transition: box-shadow 0.2s, border-color 0.2s;
}
.alarm-card:hover {
  box-shadow: 0 4px 16px #d1d9e6;
  border-left: 6px solid #409EFF;
}
.alarm-card-level {
  min-width: 48px;
  height: 48px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  font-weight: bold;
  color: #fff;
  margin-right: 18px;
}
.level-low { background: #67C23A; }
.level-medium { background: #409EFF; }
.level-high { background: #e6a23c; }
.level-urgent { background: #f56c6c; }
.alarm-card-main {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 8px;
}
.alarm-card-header {
  display: flex;
  align-items: center;
  gap: 18px;
  font-size: 16px;
  font-weight: 500;
}
.alarm-card-title {
  color: #333;
}
.alarm-card-id {
  color: #bbb;
  font-size: 13px;
}
.alarm-card-status {
  color: #409EFF;
  font-size: 14px;
}
.alarm-card-desc {
  color: #666;
  font-size: 15px;
  margin: 2px 0 4px 0;
}
.alarm-card-meta {
  color: #999;
  font-size: 13px;
  display: flex;
  flex-wrap: wrap;
  gap: 18px;
}
.alarm-card-actions {
  display: flex;
  flex-direction: column;
  gap: 10px;
  min-width: 80px;
  align-items: flex-end;
  margin-left: 18px;
}
.alarm-btn {
  font-size: 15px;
  color: #409EFF;
  font-weight: 500;
  padding: 0 8px;
  border-radius: 6px;
  transition: background 0.2s, color 0.2s;
}
.alarm-btn:hover {
  background: #ecf5ff;
  color: #1d7dfa;
}
</style>
