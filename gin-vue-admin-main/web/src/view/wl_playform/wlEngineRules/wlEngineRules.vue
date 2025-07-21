
<template>
  <div class="engine-rules-page">
    <!-- 页面头部 -->
    <div class="page-header">
      <h2>引擎规则</h2>
      <p class="description">
        规则引擎提供数据流转能力，可对接入平台的设备数据进行过滤转换，并将数据推送至用户指定的消息目的地。规则引擎流转规则需要配置消息源(推送消息类型)、条件过滤规则及消息目的地(推送方式)。
      </p>
    </div>

    <!-- 表格区域 -->
    <div class="rules-table">
      <div class="gva-btn-list">
        <el-button v-auth="btnAuth.add" type="primary" icon="plus" @click="openDialog()">新增</el-button>
        <el-button v-auth="btnAuth.batchDelete" icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        <ExportTemplate v-auth="btnAuth.exportTemplate" template-id="wl_playform_WlEngineRules" />
        <ExportExcel v-auth="btnAuth.exportExcel" template-id="wl_playform_WlEngineRules" filterDeleted/>
        <ImportExcel v-auth="btnAuth.importExcel" template-id="wl_playform_WlEngineRules" @on-success="getTableData" />
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        
        <el-table-column sortable align="left" label="日期" prop="CreatedAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
        <el-table-column align="left" label="规则名称" prop="ruleName" width="120" />

        <el-table-column align="left" label="规则描述" prop="ruleDescription" width="120" />

        <el-table-column align="left" label="消息源" prop="messageSource" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.messageSource,informationOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="查询字段" prop="queryField" width="120" />

        <el-table-column align="left" label="条件" prop="condition" width="120" />

        <el-table-column align="left" label="sql语句" prop="sqlStatement" width="120" />

        <el-table-column align="left" label="转换方法" prop="forwardingMethod" width="120" />

        <el-table-column align="left" label="使用资源id" prop="resourceId" width="120" />

        <el-table-column align="left" label="启用状态" prop="ruleStatus" width="120">
          <template #default="scope">
            <el-switch
              v-model="scope.row.ruleStatus"
              :active-value="'1'"
              :inactive-value="'0'"
              active-text="启用"
              inactive-text="禁用"
              @change="handleStatusChange(scope.row)"
              :loading="scope.row.statusLoading"
            />
          </template>
        </el-table-column>

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
          <template #default="scope">
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button" @click="updateWlEngineRulesFunc(scope.row)">编辑</el-button>
            <el-button v-auth="btnAuth.delete" type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
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
        <!-- 第一步：基本信息 -->
        <div class="form-step" :class="{ 'step-active': currentStep >= 1 }">
          <div class="step-header">
            <div class="step-number" :class="{ 'completed': currentStep > 1 }">1</div>
            <span class="step-title">基本信息</span>
          </div>
          
          <el-form-item label="规则名称:" prop="ruleName">
            <el-input 
              v-model="formData.ruleName" 
              :clearable="true" 
              placeholder="请输入规则名称"
              @input="handleRuleNameInput"
            />
          </el-form-item>
          
          <el-form-item label="规则描述:" prop="ruleDescription">
            <el-input 
              v-model="formData.ruleDescription" 
              type="textarea" 
              :rows="4"
              :clearable="true" 
              placeholder="请输入规则描述" 
              resize="vertical"
              maxlength="500"
              show-word-limit
            />
          </el-form-item>
        </div>

        <!-- 第二步：条件过滤 -->
        <div class="form-step" :class="{ 'step-active': currentStep >= 2 }" v-if="currentStep >= 2">
          <div class="step-header">
            <div class="step-number" :class="{ 'completed': currentStep > 2 }">2</div>
            <span class="step-title">条件过滤</span>
            <el-icon class="info-icon"><InfoFilled /></el-icon>
          </div>
          
          <div class="step-description">
            (SELECT [查询字段] FROM [消息源] WHERE [条件]),多个筛选项之间取交集
          </div>
          
          <div class="warning-box">
            <el-icon><Warning /></el-icon>
            <span>【重要提示:修改消息源和查询字段可能会导致输出的数据格式有变化】</span>
          </div>
          
          <el-form-item label="消息源:" prop="messageSource">
            <el-select 
              v-model="formData.messageSource" 
              placeholder="请选择" 
              style="width:100%"
              @change="handleMessageSourceChange"
            >
              <el-option label="消息总线" value="message_bus" />
              <el-option label="设备数据" value="device_data" />
              <el-option label="传感器数据" value="sensor_data" />
            </el-select>
          </el-form-item>
          
          <el-form-item label="查询字段:" prop="queryField">
            <el-input 
              v-model="formData.queryField" 
              :clearable="true" 
              placeholder="请输入查询字段"
              maxlength="300"
              show-word-limit
              @input="updateSqlStatement"
            />
          </el-form-item>
          
          <el-form-item label="条件:" prop="condition">
            <el-input 
              v-model="formData.condition" 
              :clearable="true" 
              placeholder="请输入"
              @input="updateSqlStatement"
            />
          </el-form-item>
          
          <el-form-item label="SQL语句展示:" prop="sqlStatement">
            <el-input 
              v-model="formData.sqlStatement" 
              type="textarea" 
              :rows="3"
              readonly
              placeholder="SQL语句将在这里显示"
            />
          </el-form-item>
        </div>

        <!-- 第三步：转发方式 -->
        <div class="form-step" :class="{ 'step-active': currentStep >= 3 }" v-if="currentStep >= 3">
          <div class="step-header">
            <div class="step-number" :class="{ 'completed': currentStep > 3 }">3</div>
            <span class="step-title">转发方式</span>
          </div>
          
          <el-form-item label="转发方式:" prop="forwardingMethod">
            <el-radio-group v-model="formData.forwardingMethod" @change="handleForwardingMethodChange">
              <el-radio label="HTTP推送">HTTP推送</el-radio>
              <el-radio label="消息对队列MQTT">消息对队列MQTT</el-radio>
              <el-radio label="消息队列Kafka">消息队列Kafka</el-radio>
              <el-radio label="InfluxDB">InfluxDB</el-radio>
              <el-radio label="TDengine">TDengine</el-radio>
            </el-radio-group>
          </el-form-item>
          
          <el-form-item label="使用资源:" prop="resourceId" v-if="formData.forwardingMethod">
            <el-select 
              v-model="formData.resourceId" 
              placeholder="请选择" 
              style="width:100%"
              filterable
            >
              <el-option 
                v-for="resource in availableResources" 
                :key="resource.id" 
                :label="resource.name" 
                :value="resource.id"
              />
            </el-select>
            <div class="form-hint">您也可以进入资源管理添加新的实例</div>
          </el-form-item>
        </div>
      </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="规则名称">
          {{ detailFrom.ruleName }}
        </el-descriptions-item>
        <el-descriptions-item label="规则描述">
          {{ detailFrom.ruleDescription }}
        </el-descriptions-item>
        <el-descriptions-item label="消息源">
          {{ detailFrom.messageSource }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createWlEngineRules,
  deleteWlEngineRules,
  deleteWlEngineRulesByIds,
  updateWlEngineRules,
  findWlEngineRules,
  getWlEngineRulesList
} from '@/api/wl_playform/wlEngineRules'

import {
  getWlResourcesList,
  findWlResources
} from '@/api/wl_playform/wlResources'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { Plus, Delete, InfoFilled, Warning } from '@element-plus/icons-vue'
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'
import { useAppStore } from "@/pinia"

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'


defineOptions({
    name: 'WlEngineRules'
})
// 按钮权限实例化
    const btnAuth = useBtnAuth()

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 动态表单显示控制
const showFilterCondition = ref(false)
const showConversionMethod = ref(false)
const showResourceSelect = ref(false)

// 资源加载状态
const resourceLoading = ref(false)

// 过滤条件数组
const filterConditions = ref([])

// MongoDB资源key
const resourceKey = ref('')

// 自动化生成的字典（可能为空）以及字段
const informationOptions = ref([])
const resourceOptions = ref([])

// 转换方法选项
const conversionMethods = ref([
  { label: 'HTTP请求', value: 'http' },
  { label: '数据库查询', value: 'database' },
  { label: '文件处理', value: 'file' },
  { label: '数据转换', value: 'transform' }
])

// 表单数据
const formData = ref({
  ruleName: '',
  ruleDescription: '',
  messageSource: '',
  queryField: '*',
  condition: '',
  sqlStatement: 'SELECT * FROM mqtt_stream',
  forwardingMethod: '',
  resourceId: undefined,
  ruleStatus: '0', // 默认禁用状态
})

// 当前步骤
const currentStep = ref(1)

// 可用资源列表
const availableResources = ref([
  { id: 1, name: 'HTTP推送实例1', type: 'HTTP推送' },
  { id: 2, name: 'MQTT实例1', type: '消息对队列MQTT' },
  { id: 3, name: 'Kafka实例1', type: '消息队列Kafka' },
  { id: 4, name: 'InfluxDB实例1', type: 'InfluxDB' },
  { id: 5, name: 'TDengine实例1', type: 'TDengine' },
])

// 处理规则名称输入
const handleRuleNameInput = () => {
  if (formData.value.ruleName && currentStep.value === 1) {
    currentStep.value = 2
  }
}

// 处理消息源变化
const handleMessageSourceChange = () => {
  if (formData.value.messageSource && currentStep.value === 2) {
    currentStep.value = 3
  }
  updateSqlStatement()
}

// 处理转发方式变化
const handleForwardingMethodChange = () => {
  // 根据选择的转发方式过滤可用资源
  if (formData.value.forwardingMethod) {
    // 这里可以根据选择的转发方式动态加载对应的资源
    console.log('选择的转发方式:', formData.value.forwardingMethod)
  }
}

// 更新SQL语句
const updateSqlStatement = () => {
  const queryField = formData.value.queryField || '*'
  const messageSource = formData.value.messageSource || 'mqtt_stream'
  const condition = formData.value.condition || ''
  
  let sql = `SELECT ${queryField} FROM ${messageSource}`
  if (condition) {
    sql += ` WHERE ${condition}`
  }
  
  formData.value.sqlStatement = sql
}


// 验证规则
const rule = reactive({
               ruleName : [{
                   required: true,
                   message: '请填写规则名称',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               messageSource : [{
                   required: true,
                   message: '请填写消息源',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               queryField : [{
                   required: true,
                   message: '请填写查询字段',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               forwardingMethod : [{
                   required: true,
                   message: '请选择转换方法',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               resourceId : [{
                   required: true,
                   message: '请选择资源',
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
  const table = await getWlEngineRulesList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
    
    // 数据加载完成后，设置初始化标志为false
    setTimeout(() => {
      isInitializing.value = false
    }, 100)
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    informationOptions.value = await getDictFunc('information')
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
            deleteWlEngineRulesFunc(row)
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
      const res = await deleteWlEngineRulesByIds({ IDs })
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
const updateWlEngineRulesFunc = async(row) => {
    const res = await findWlEngineRules({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        // 根据现有数据设置动态显示状态
        showFilterCondition.value = !!formData.value.ruleName
        showConversionMethod.value = !!formData.value.messageSource
        showResourceSelect.value = !!formData.value.forwardingMethod
        
        // 如果需要显示资源选择，加载资源选项
        if (showResourceSelect.value) {
          loadResourceOptions()
        }
        
        // 如果有资源ID，尝试加载过滤条件
        if (formData.value.resourceId) {
          loadFilterConditions(formData.value.resourceId)
        }
        
        dialogFormVisible.value = true
    }
}

// 加载过滤条件
const loadFilterConditions = async (resourceId) => {
  try {
    // 查询资源表获取resources_key
    const res = await findWlResources({ ID: resourceId })
    if (res.code === 0 && res.data && res.data.resourcesKey) {
      resourceKey.value = res.data.resourcesKey
      
      // 这里应该有一个API调用来从MongoDB获取过滤条件
      // 由于没有看到具体的API，这里模拟一个加载过程
      // 实际实现时，需要替换为真实的API调用
      
      // 模拟从MongoDB加载数据
      setTimeout(() => {
        // 假设从MongoDB获取到的数据格式如下
        const mockData = {
          conditions: [
            { key: 'name', value: 'test' },
            { key: 'age', value: '18' }
          ]
        }
        
        // 将数据赋值给filterConditions
        filterConditions.value = mockData.conditions || []
      }, 500)
    }
  } catch (error) {
    console.error('加载过滤条件失败:', error)
    ElMessage.error('加载过滤条件失败')
  }
}


// 删除行
const deleteWlEngineRulesFunc = async (row) => {
    const res = await deleteWlEngineRules({ ID: row.ID })
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

// 初始化标志
const isInitializing = ref(true)

// 状态切换处理
const handleStatusChange = async (row) => {
  // 如果是初始化阶段，不显示提示
  if (isInitializing.value) {
    return
  }
  
  // 设置loading状态
  row.statusLoading = true
  
  try {
    const res = await updateWlEngineRules({
      ID: row.ID,
      ruleName: row.ruleName,
      ruleDescription: row.ruleDescription,
      messageSource: row.messageSource,
      queryField: row.queryField,
      condition: row.condition,
      sqlStatement: row.sqlStatement,
      forwardingMethod: row.forwardingMethod,
      resourceId: row.resourceId,
      ruleStatus: row.ruleStatus
    })
    
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: row.ruleStatus === '1' ? '启用成功' : '禁用成功'
      })
    } else {
      // 如果更新失败，恢复原状态
      row.ruleStatus = row.ruleStatus === '1' ? '0' : '1'
      ElMessage({
        type: 'error',
        message: '状态更新失败'
      })
    }
  } catch (error) {
    // 如果出错，恢复原状态
    row.ruleStatus = row.ruleStatus === '1' ? '0' : '1'
    ElMessage({
      type: 'error',
      message: '状态更新失败'
    })
  } finally {
    row.statusLoading = false
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  resetFormData()
  dialogFormVisible.value = true
}

// 重置表单数据
const resetFormData = () => {
  formData.value = {
    ruleName: '',
    ruleDescription: '',
    messageSource: '',
    queryField: '*',
    condition: '',
    sqlStatement: 'SELECT * FROM mqtt_stream',
    forwardingMethod: '',
    resourceId: undefined,
    ruleStatus: '0', // 默认禁用状态
  }
  currentStep.value = 1
}

// 规则名称变化处理
const onRuleNameChange = (value) => {
  showFilterCondition.value = !!value && value.trim().length > 0
  if (!showFilterCondition.value) {
    // 重置后续字段
    showConversionMethod.value = false
    showResourceSelect.value = false
    formData.value.messageSource = ''
    formData.value.queryField = ''
    formData.value.condition = ''
    formData.value.forwardingMethod = ''
    formData.value.sqlStatement = ''
    formData.value.resourceId = undefined
    // 清空过滤条件
    filterConditions.value = []
  }
}

// 添加过滤条件
const addFilterCondition = () => {
  filterConditions.value.push({
    key: '',
    value: ''
  })
}

// 删除过滤条件
const removeFilterCondition = (index) => {
  filterConditions.value.splice(index, 1)
}

// 消息源变化处理
const onMessageSourceChange = (value) => {
  showConversionMethod.value = !!value
  if (!showConversionMethod.value) {
    // 重置后续字段
    showResourceSelect.value = false
    formData.value.forwardingMethod = ''
    formData.value.sqlStatement = ''
    formData.value.resourceId = undefined
  }
}

// 转换方法变化处理
const onConversionMethodChange = (value) => {
  showResourceSelect.value = !!value
  if (showResourceSelect.value) {
    // 加载资源选项
    loadResourceOptions()
  } else {
    formData.value.resourceId = undefined
  }
}

// 加载资源选项
const loadResourceOptions = async () => {
  resourceLoading.value = true
  try {
    const res = await getWlResourcesList({ page: 1, pageSize: 100 })
    if (res.code === 0 && res.data && res.data.list) {
      resourceOptions.value = res.data.list.map(item => ({
        label: item.resourceName || `资源${item.ID}`,
        value: item.ID
      }))
    }
  } catch (error) {
    console.error('加载资源选项失败:', error)
  } finally {
    resourceLoading.value = false
  }
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    // 重置动态显示状态
    showFilterCondition.value = false
    showConversionMethod.value = false
    showResourceSelect.value = false
    // 清空过滤条件
    filterConditions.value = []
    // 重置资源key
    resourceKey.value = ''
    formData.value = {
        ruleName: '',
        ruleDescription: '',
        messageSource: '',
        queryField: '',
        condition: '',
        sqlStatement: '',
        forwardingMethod: '',
        resourceId: undefined,
        }
}
// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
             
             try {
               // 1. 先将过滤条件保存到MongoDB
               const mongoData = {
                 conditions: filterConditions.value,
                 ruleName: formData.value.ruleName,
                 createdAt: new Date().toISOString()
               }
               
               // 这里应该有一个API调用来保存数据到MongoDB
               // 由于没有看到具体的API，这里模拟一个保存过程
               // 实际实现时，需要替换为真实的API调用
               
               // 模拟保存到MongoDB并获取key
               const mongoKey = await saveToMongoDB(mongoData)
               
               // 2. 将MongoDB返回的key保存到formData中
               formData.value.resourcesKey = mongoKey
               
               // 3. 保存规则数据
               let res
               switch (type.value) {
                 case 'create':
                   res = await createWlEngineRules(formData.value)
                   break
                 case 'update':
                   res = await updateWlEngineRules(formData.value)
                   break
                 default:
                   res = await createWlEngineRules(formData.value)
                   break
               }
               
               if (res.code === 0) {
                 ElMessage({
                   type: 'success',
                   message: '创建/更改成功'
                 })
                 closeDialog()
                 getTableData()
               }
             } catch (error) {
               console.error('保存数据失败:', error)
               ElMessage.error('保存数据失败')
             } finally {
               btnLoading.value = false
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
  const res = await findWlEngineRules({ ID: row.ID })
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

// 保存数据到MongoDB
const saveToMongoDB = async (data) => {
  // 这里应该是真实的API调用
  // 由于没有看到具体的API，这里模拟一个保存过程
  // 实际实现时，需要替换为真实的API调用
  
  return new Promise((resolve) => {
    setTimeout(() => {
      // 模拟MongoDB返回的key
      const key = 'mongo_' + Date.now()
      resolve(key)
    }, 500)
  })
}


</script>

<style scoped>
.engine-rules-page {
  background: #0f172a;
  min-height: 100vh;
  padding: 20px;
  color: #fff;
}

.page-header {
  margin-bottom: 20px;
  background: #0f172a;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.4);
  border: 1px solid #334155;
}

.page-header h2 {
  margin: 0 0 10px 0;
  font-size: 24px;
  font-weight: 600;
  color: #fff;
}

.description {
  margin: 0;
  color: #cbd5e1;
  line-height: 1.6;
}

.action-buttons {
  margin-bottom: 20px;
  background: #0f172a;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.4);
  border: 1px solid #334155;
}

.rules-table {
  background: #0f172a;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.4);
  border: 1px solid #334155;
}

/* 表格样式 */
.rules-table :deep(.el-table__header-wrapper) {
  background: #0f172a;
}

.rules-table :deep(.el-table__header) {
  background: #0f172a;
}

.rules-table :deep(.el-table__header th) {
  background: #0f172a !important;
  color: #fff !important;
  border-bottom: 1px solid #334155;
}

.rules-table :deep(.el-table__header th .cell) {
  color: #fff !important;
  font-weight: 600;
}

.rules-table :deep(.el-table__body tr) {
  background: #0f172a;
  color: #fff;
}

.rules-table :deep(.el-table__body tr:hover) {
  background: #1e293b;
}

.rules-table :deep(.el-table__body td) {
  background: #0f172a;
  color: #fff;
  border-bottom: 1px solid #334155;
}

.rules-table :deep(.el-table) {
  border: 1px solid #334155;
  border-radius: 8px;
  overflow: hidden;
  background: #0f172a;
}

.rules-table :deep(.el-table__border-line) {
  background: #334155;
}

.rules-table :deep(.el-table__body-wrapper) {
  background: #0f172a;
}

.rules-table :deep(.el-table__fixed-header-wrapper) {
  background: #0f172a;
}

.rules-table :deep(.el-table__fixed-body-wrapper) {
  background: #0f172a;
}

.rules-table :deep(.el-table__empty-block) {
  background: #0f172a;
}

.rules-table :deep(.el-table__empty-text) {
  color: #cbd5e1;
}

/* 分页样式 */
.gva-pagination {
  margin-top: 20px;
  text-align: right;
}

.gva-pagination :deep(.el-pagination) {
  color: #fff;
}

.gva-pagination :deep(.el-pagination .el-pager li) {
  background: #0f172a;
  color: #fff;
  border: 1px solid #334155;
}

.gva-pagination :deep(.el-pagination .el-pager li.is-active) {
  background: #409eff;
  color: #fff;
}

.gva-pagination :deep(.el-pagination .btn-prev),
.gva-pagination :deep(.el-pagination .btn-next) {
  background: #0f172a;
  color: #fff;
  border: 1px solid #334155;
}

/* 按钮样式调整 */
.action-buttons :deep(.el-button) {
  border: 1px solid #334155;
}

.action-buttons :deep(.el-button--primary) {
  background: #409eff;
  border-color: #409eff;
}

/* 搜索框样式 */
.gva-search-box {
  margin-bottom: 20px;
  background: #0f172a;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.4);
  border: 1px solid #334155;
}

/* 抽屉样式调整 */
:deep(.el-drawer) {
  background: #0f172a;
  color: #fff;
}

:deep(.el-drawer__header) {
  background: #0f172a;
  color: #fff;
  border-bottom: 1px solid #334155;
}

:deep(.el-drawer__body) {
  background: #0f172a;
  color: #fff;
}

/* 表单样式调整 */
:deep(.el-form-item__label) {
  color: #fff;
}

:deep(.el-input__wrapper) {
  background: #334155;
  border: 1px solid #475569;
}

:deep(.el-input__inner) {
  background: #334155;
  color: #fff;
  border: none;
}

:deep(.el-select .el-input__wrapper) {
  background: #334155;
  border: 1px solid #475569;
}

/* 表单步骤样式 */
.form-step {
  margin-bottom: 30px;
  opacity: 0.6;
  transition: all 0.3s ease;
}

.form-step.step-active {
  opacity: 1;
}

.step-header {
  display: flex;
  align-items: center;
  margin-bottom: 20px;
  padding-bottom: 10px;
  border-bottom: 1px solid #334155;
}

.step-number {
  width: 30px;
  height: 30px;
  border-radius: 50%;
  background: #475569;
  color: #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  margin-right: 10px;
  transition: all 0.3s ease;
}

.step-number.completed {
  background: #409eff;
}

.step-title {
  font-size: 16px;
  font-weight: 600;
  color: #fff;
  margin-right: 10px;
}

.info-icon {
  color: #409eff;
  font-size: 18px;
}

.step-description {
  color: #cbd5e1;
  font-size: 14px;
  margin-bottom: 15px;
  padding: 10px;
  background: #1e293b;
  border-radius: 4px;
  border-left: 3px solid #409eff;
}

.warning-box {
  display: flex;
  align-items: center;
  padding: 12px;
  background: #fef3c7;
  border: 1px solid #f59e0b;
  border-radius: 6px;
  margin-bottom: 20px;
  color: #92400e;
}

.warning-box .el-icon {
  color: #f59e0b;
  margin-right: 8px;
  font-size: 16px;
}

.form-hint {
  color: #94a3b8;
  font-size: 12px;
  margin-top: 5px;
}

/* 单选框样式 */
:deep(.el-radio) {
  color: #fff;
  margin-right: 20px;
  margin-bottom: 10px;
}

:deep(.el-radio__input.is-checked .el-radio__inner) {
  background: #409eff;
  border-color: #409eff;
}

:deep(.el-radio__input.is-checked + .el-radio__label) {
  color: #409eff;
}

/* 下拉框样式 */
:deep(.el-select-dropdown) {
  background: #1e293b;
  border: 1px solid #334155;
}

:deep(.el-select-dropdown__item) {
  color: #fff;
}

:deep(.el-select-dropdown__item:hover) {
  background: #334155;
}

:deep(.el-select-dropdown__item.selected) {
  background: #409eff;
  color: #fff;
}

/* 开关样式 */
:deep(.el-switch) {
  --el-switch-on-color: #409eff;
  --el-switch-off-color: #606266;
}

:deep(.el-switch__core) {
  border-color: #475569;
}

:deep(.el-switch.is-checked .el-switch__core) {
  border-color: #409eff;
}

:deep(.el-switch__label) {
  color: #fff;
}

:deep(.el-switch__label.is-active) {
  color: #409eff;
}
</style>
