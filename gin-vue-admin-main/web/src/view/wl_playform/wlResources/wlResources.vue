
<template>
  <div class="resource-management-page">
    <!-- 页面标题和描述 -->
    <div class="page-header">
      <h2>资源管理</h2>
      <p class="description">
        平台提供多种消息通信中间件资源，资源可作为规则引擎的消息目的地，通过创建资源快速将数据推送至应用平台。
      </p>
    </div>

    <!-- 资源类型标签页 -->
    <div class="resource-tabs">
      <el-tabs v-model="activeResourceType" @tab-click="handleResourceTypeChange">
        <el-tab-pane label="HTTP推送" name="http">
          <template #label>
            <span>HTTP推送</span>
          </template>
        </el-tab-pane>
        <el-tab-pane label="消息对队列MQTT" name="mqtt">
          <template #label>
            <span>消息对队列MQTT</span>
          </template>
        </el-tab-pane>
        <el-tab-pane label="消息队列Kafka" name="kafka">
          <template #label>
            <span>消息队列Kafka</span>
          </template>
        </el-tab-pane>
        <el-tab-pane label="InfluxDB" name="influxdb">
          <template #label>
            <span>InfluxDB</span>
          </template>
        </el-tab-pane>
        <el-tab-pane label="TDengine" name="tdengine">
          <template #label>
            <span>TDengine</span>
          </template>
        </el-tab-pane>
      </el-tabs>
    </div>

    <!-- 资源列表表格 -->
    <div class="resource-table">
      <div class="gva-btn-list">
        <el-button v-auth="btnAuth.add" type="primary" icon="plus" @click="openDialog()">新增</el-button>
        <el-button v-auth="btnAuth.batchDelete" icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
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
        <el-table-column sortable align="left" label="实例名称" prop="instanceName" width="200" />
        <el-table-column sortable align="left" label="创建时间" prop="CreatedAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="验证状态" prop="verificationStatus" width="120">
          <template #default="scope">
            <el-tag :type="getStatusType(scope.row.verificationStatus)">
              {{ scope.row.verificationStatus || '未验证' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
          <template #default="scope">
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button" @click="getDetails(scope.row)">
              <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看
            </el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button" @click="updateWlResourcesFunc(scope.row)">编辑</el-button>
            <el-button 
              v-auth="btnAuth.verify"
              type="primary" 
              link 
              :loading="scope.row.verifying"
              @click="handleVerification(scope.row)"
            >
              验证
            </el-button>
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

    <!-- 添加实例对话框 -->
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === 'create' ? '添加实例' : '编辑实例' }}</span>
          <div>
            <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>

      <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
        <!-- 实例名称 -->
        <el-form-item label="实例名称:" prop="instanceName">
          <el-input v-model="formData.instanceName" :clearable="true" placeholder="1-32位字符,支持中文、英文、数字及特殊字符_-,必须以英文或中文字符开头" />
        </el-form-item>

        <!-- 动态表单字段 -->
        <template v-if="activeResourceType === 'http'">
          <el-form-item label="URL:" prop="url">
            <el-input 
              v-model="formData.url" 
              :clearable="true" 
              placeholder="请输入需要推送的服务地址,以http://或https://开头"
              class="http-input"
            />
          </el-form-item>
          
          <el-form-item label="HTTP method:" prop="httpMethod">
            <el-select 
              v-model="formData.httpMethod" 
              placeholder="请选择" 
              style="width:100%"
              class="http-select"
            >
              <el-option label="GET" value="GET" />
              <el-option label="POST" value="POST" />
              <el-option label="PUT" value="PUT" />
              <el-option label="DELETE" value="DELETE" />
            </el-select>
          </el-form-item>
          
          <el-form-item label="Body type:" prop="bodyType">
            <el-select 
              v-model="formData.bodyType" 
              placeholder="请选择" 
              style="width:100%"
              class="http-select"
            >
              <el-option label="JSON" value="json" />
              <el-option label="XML" value="xml" />
              <el-option label="Form" value="form" />
            </el-select>
          </el-form-item>
          
          <el-form-item label="Timeout(ms):" prop="timeout">
            <el-input 
              v-model="formData.timeout" 
              :clearable="true" 
              placeholder="请输入超时时间（毫秒）"
              class="http-input"
            />
          </el-form-item>
          
          <el-form-item label="HTTP headers:" prop="headers">
            <div class="headers-container">
              <div class="headers-header">
                <div class="headers-title-section">
                  <el-icon class="header-icon"><Connection /></el-icon>
                  <span class="headers-title">请求头配置</span>
                  <span class="headers-count">({{ formData.headers.length }} 项)</span>
                </div>
                <el-button 
                  type="primary" 
                  size="small" 
                  icon="Plus" 
                  @click="addHeader"
                  class="add-header-btn"
                >
                  添加请求头
                </el-button>
              </div>
              
              <div class="headers-list">
                <div 
                  v-for="(header, index) in formData.headers" 
                  :key="index" 
                  class="header-item"
                  :class="{ 
                    'header-item-active': header.key || header.value,
                    'header-item-filled': header.key && header.value 
                  }"
                >
                  <div class="header-index">
                    <span class="index-number">{{ index + 1 }}</span>
                  </div>
                  
                  <div class="header-inputs">
                    <div class="input-group">
                      <label class="input-label">Key</label>
                      <el-input 
                        v-model="header.key" 
                        placeholder="请输入请求头名称" 
                        class="header-key-input"
                        clearable
                      />
                    </div>
                    
                    <div class="input-group">
                      <label class="input-label">Value</label>
                      <el-input 
                        v-model="header.value" 
                        placeholder="请输入请求头值" 
                        class="header-value-input"
                        clearable
                      />
                    </div>
                  </div>
                  
                  <div class="header-actions">
                    <el-button 
                      type="danger" 
                      size="small" 
                      icon="Delete" 
                      @click="removeHeader(index)"
                      class="remove-header-btn"
                      :disabled="formData.headers.length === 1"
                      title="删除此请求头"
                    />
                  </div>
                </div>
              </div>
              
              <div v-if="formData.headers.length === 0" class="empty-headers">
                <el-empty description="暂无请求头配置" :image-size="60">
                  <template #image>
                    <el-icon class="empty-icon"><Connection /></el-icon>
                  </template>
                </el-empty>
              </div>
              
              <div class="headers-footer">
                <span class="footer-tip">💡 提示：请求头用于设置HTTP请求的额外参数，如Content-Type、Authorization等</span>
              </div>
            </div>
          </el-form-item>
        </template>

        <template v-if="activeResourceType === 'mqtt'">
          <el-form-item label="MQTT broker address:" prop="brokerAddress">
            <el-input v-model="formData.brokerAddress" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="MQTT topic:" prop="topic">
            <el-input v-model="formData.topic" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="MQTT client:" prop="client">
            <el-input v-model="formData.client" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="MQTT protocol version:" prop="protocolVersion">
            <el-select v-model="formData.protocolVersion" placeholder="请选择" style="width:100%">
              <el-option label="3.1" value="3.1" />
              <el-option label="3.1.1" value="3.1.1" />
              <el-option label="5.0" value="5.0" />
            </el-select>
          </el-form-item>
          <el-form-item label="QoS:" prop="qos">
            <el-select v-model="formData.qos" placeholder="请选择" style="width:100%">
              <el-option label="0" :value="0" />
              <el-option label="1" :value="1" />
              <el-option label="2" :value="2" />
            </el-select>
          </el-form-item>
          <el-form-item label="Username:" prop="username">
            <el-input v-model="formData.username" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="Password:" prop="password">
            <el-input v-model="formData.password" type="password" :clearable="true" placeholder="请输入" show-password />
          </el-form-item>
        </template>

        <template v-if="activeResourceType === 'kafka'">
          <el-form-item label="Kafka brokers:" prop="brokers">
            <el-input v-model="formData.brokers" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="Topic:" prop="topic">
            <el-input v-model="formData.topic" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="SaslAuthType:" prop="saslAuthType">
            <el-select v-model="formData.saslAuthType" placeholder="请选择" style="width:100%">
              <el-option label="PLAIN" value="PLAIN" />
              <el-option label="SCRAM-SHA-256" value="SCRAM-SHA-256" />
              <el-option label="SCRAM-SHA-512" value="SCRAM-SHA-512" />
            </el-select>
          </el-form-item>
          <el-form-item label="SaslUserName:" prop="saslUserName">
            <el-input v-model="formData.saslUserName" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="SaslPassword:" prop="saslPassword">
            <el-input v-model="formData.saslPassword" type="password" :clearable="true" placeholder="请输入" show-password />
          </el-form-item>
        </template>

        <template v-if="activeResourceType === 'influxdb'">
          <el-form-item label="Host:" prop="host">
            <el-input v-model="formData.host" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="Port:" prop="port">
            <el-input v-model="formData.port" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="User:" prop="user">
            <el-input v-model="formData.user" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="Password:" prop="password">
            <el-input v-model="formData.password" type="password" :clearable="true" placeholder="请输入" show-password />
          </el-form-item>
          <el-form-item label="Measurement:" prop="measurement">
            <el-input v-model="formData.measurement" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="Databasename:" prop="databasename">
            <el-input v-model="formData.databasename" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="Tagkey:" prop="tagkey">
            <el-input v-model="formData.tagkey" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="Tagvalue:" prop="tagvalue">
            <el-input v-model="formData.tagvalue" :clearable="true" placeholder="请输入" />
          </el-form-item>
        </template>

        <template v-if="activeResourceType === 'tdengine'">
          <el-form-item label="Host:" prop="host">
            <el-input v-model="formData.host" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="Port:" prop="port">
            <el-input v-model="formData.port" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="User:" prop="user">
            <el-input v-model="formData.user" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="Password:" prop="password">
            <el-input v-model="formData.password" type="password" :clearable="true" placeholder="请输入" show-password />
          </el-form-item>
          <el-form-item label="Database:" prop="database">
            <el-input v-model="formData.database" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="Table:" prop="table">
            <el-input v-model="formData.table" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="Fields:" prop="fields">
            <el-input v-model="formData.fields" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="ProvideTs:" prop="provideTs">
            <el-select v-model="formData.provideTs" placeholder="请选择" style="width:100%">
              <el-option label="true" :value="true" />
              <el-option label="false" :value="false" />
            </el-select>
          </el-form-item>
          <el-form-item label="TsFieldName:" prop="tsFieldName">
            <el-input v-model="formData.tsFieldName" :clearable="true" placeholder="请输入" />
          </el-form-item>
        </template>
      </el-form>
    </el-drawer>

    <!-- 详情查看对话框 -->
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看详情">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="实例名称">{{ detailFrom.instanceName }}</el-descriptions-item>
        <el-descriptions-item label="创建时间">{{ formatDate(detailFrom.CreatedAt) }}</el-descriptions-item>
        <el-descriptions-item label="验证状态">{{ detailFrom.verificationStatus || '未验证' }}</el-descriptions-item>
        <el-descriptions-item label="资源类型">{{ getResourceTypeName(activeResourceType) }}</el-descriptions-item>
      </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createWlResourcesWithTransaction,
  deleteWlResources,
  deleteWlResourcesByIds,
  updateWlResources,
  findWlResources,
  getWlResourcesList,
  verifyWlResources
} from '@/api/wl_playform/wlResources'

import { getDictFunc, formatDate, formatBoolean, filterDict, filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useBtnAuth } from '@/utils/btnAuth'
import { useAppStore } from "@/pinia"
import { Connection } from '@element-plus/icons-vue'

defineOptions({
  name: 'WlResources'
})

// 按钮权限实例化
const btnAuth = useBtnAuth()
const appStore = useAppStore()

// 提交按钮loading
const btnLoading = ref(false)

// 当前激活的资源类型
const activeResourceType = ref('http')

// 表单数据
const formData = ref({
  instanceName: '',
  // HTTP相关字段
  url: '',
  httpMethod: '',
  bodyType: '',
  timeout: '',
  headers: [{ key: '', value: '' }],
  // MQTT相关字段
  brokerAddress: '',
  topic: '',
  client: '',
  protocolVersion: '',
  qos: 0,
  username: '',
  password: '',
  // Kafka相关字段
  brokers: '',
  saslAuthType: '',
  saslUserName: '',
  saslPassword: '',
  // InfluxDB相关字段
  host: '',
  port: '',
  user: '',
  measurement: '',
  databasename: '',
  tagkey: '',
  tagvalue: '',
  // TDengine相关字段
  database: '',
  table: '',
  fields: '',
  provideTs: false,
  tsFieldName: ''
})

// 验证规则
const rule = reactive({
  instanceName: [{
    required: true,
    message: '请填写实例名称',
    trigger: ['input', 'blur'],
  }]
})

const elFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 查询
const getTableData = async () => {
  const table = await getWlResourcesList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 多选数据
const multipleSelection = ref([])
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
    deleteWlResourcesFunc(row)
  })
}

// 批量删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const ids = multipleSelection.value.map(item => item.ID)
    if (ids.length === 0) {
      ElMessage({
        type: 'warning',
        message: '请选择要删除的数据'
      })
      return
    }
    const res = await deleteWlResourcesByIds({ ids })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === ids.length && page.value > 1) {
        page.value--
      }
      await getTableData()
    }
  })
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateWlResourcesFunc = async (row) => {
  const res = await findWlResources({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteWlResourcesFunc = async (row) => {
  const res = await deleteWlResources({ ID: row.ID })
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

// 验证处理
const handleVerification = async (row) => {
  // 设置验证loading状态
  row.verifying = true
  
  try {
    // 模拟验证过程
    await new Promise(resolve => setTimeout(resolve, 2000))
    
    // 调用验证API
    const res = await verifyWlResources({ ID: row.ID })
    
    if (res.code === 0) {
      // 更新验证状态
      row.verificationStatus = '验证成功'
      ElMessage({
        type: 'success',
        message: '验证成功'
      })
    } else {
      // 验证失败
      row.verificationStatus = '验证失败'
      ElMessage({
        type: 'error',
        message: '验证失败'
      })
    }
  } catch (error) {
    // 验证出错
    row.verificationStatus = '验证失败'
    ElMessage({
      type: 'error',
      message: '验证失败'
    })
  } finally {
    row.verifying = false
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

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  resetFormData()
}

// 重置表单数据
const resetFormData = () => {
  formData.value = {
    instanceName: '',
    url: '',
    httpMethod: '',
    bodyType: '',
    timeout: '',
    headers: [{ key: '', value: '' }],
    brokerAddress: '',
    topic: '',
    client: '',
    protocolVersion: '',
    qos: 0,
    username: '',
    password: '',
    brokers: '',
    saslAuthType: '',
    saslUserName: '',
    saslPassword: '',
    host: '',
    port: '',
    user: '',
    measurement: '',
    databasename: '',
    tagkey: '',
    tagvalue: '',
    database: '',
    table: '',
    fields: '',
    provideTs: false,
    tsFieldName: ''
  }
}

// 弹窗确定
const enterDialog = async () => {
  btnLoading.value = true
  try {
    const valid = await elFormRef.value?.validate()
    if (!valid) {
      btnLoading.value = false
      return
    }
    
    // 构建资源数据
    const resourceData = buildResourceData()
    
    // 生成请求ID防止重复提交
    const requestId = Date.now().toString() + '_' + Math.random().toString(36).substr(2, 9)
    
    let res
    switch (type.value) {
      case 'create':
        res = await createWlResourcesWithTransaction({
          ...formData.value,
          resourceType: activeResourceType.value,
          resourceData: resourceData,
          requestId: requestId
        })
        break
      case 'update':
        res = await updateWlResources({
          ...formData.value,
          resourceType: activeResourceType.value,
          resourceData: resourceData
        })
        break
      default:
        res = await createWlResourcesWithTransaction({
          ...formData.value,
          resourceType: activeResourceType.value,
          resourceData: resourceData,
          requestId: requestId
        })
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
    console.error('提交失败:', error)
    ElMessage({
      type: 'error',
      message: '提交失败，请重试'
    })
  } finally {
    btnLoading.value = false
  }
}

// 构建资源数据
const buildResourceData = () => {
  const data = {
    resourceType: activeResourceType.value,
    instanceName: formData.value.instanceName
  }

  switch (activeResourceType.value) {
    case 'http':
      data.url = formData.value.url
      data.httpMethod = formData.value.httpMethod
      data.bodyType = formData.value.bodyType
      data.timeout = formData.value.timeout
      data.headers = formData.value.headers.filter(h => h.key && h.value)
      break
    case 'mqtt':
      data.brokerAddress = formData.value.brokerAddress
      data.topic = formData.value.topic
      data.client = formData.value.client
      data.protocolVersion = formData.value.protocolVersion
      data.qos = formData.value.qos
      data.username = formData.value.username
      data.password = formData.value.password
      break
    case 'kafka':
      data.brokers = formData.value.brokers
      data.topic = formData.value.topic
      data.saslAuthType = formData.value.saslAuthType
      data.saslUserName = formData.value.saslUserName
      data.saslPassword = formData.value.saslPassword
      break
    case 'influxdb':
      data.host = formData.value.host
      data.port = formData.value.port
      data.user = formData.value.user
      data.password = formData.value.password
      data.measurement = formData.value.measurement
      data.databasename = formData.value.databasename
      data.tagkey = formData.value.tagkey
      data.tagvalue = formData.value.tagvalue
      break
    case 'tdengine':
      data.host = formData.value.host
      data.port = formData.value.port
      data.user = formData.value.user
      data.password = formData.value.password
      data.database = formData.value.database
      data.table = formData.value.table
      data.fields = formData.value.fields
      data.provideTs = formData.value.provideTs
      data.tsFieldName = formData.value.tsFieldName
      break
  }

  return data
}

// 处理资源类型变化
const handleResourceTypeChange = () => {
  resetFormData()
}

// 添加HTTP header
const addHeader = () => {
  formData.value.headers.push({ key: '', value: '' })
}

// 移除HTTP header
const removeHeader = (index) => {
  formData.value.headers.splice(index, 1)
}

// 获取状态类型
const getStatusType = (status) => {
  if (status === '验证成功') return 'success'
  if (status === '验证失败') return 'danger'
  return 'info'
}

// 获取资源类型名称
const getResourceTypeName = (type) => {
  const typeMap = {
    'http': 'HTTP推送',
    'mqtt': '消息对队列MQTT',
    'kafka': '消息队列Kafka',
    'influxdb': 'InfluxDB',
    'tdengine': 'TDengine'
  }
  return typeMap[type] || '未知类型'
}

const detailFrom = ref({})
const detailShow = ref(false)

// 打开详情
const getDetails = async (row) => {
  const res = await findWlResources({ ID: row.ID })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}

// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}

// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}
</script>

<style scoped>
.resource-management-page {
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

.resource-tabs {
  margin-bottom: 20px;
  background: #0f172a;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.4);
  border: 1px solid #334155;
}

.action-buttons {
  margin-bottom: 20px;
  background: #0f172a;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.4);
  border: 1px solid #334155;
}

.resource-table {
  background: #0f172a;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.4);
  border: 1px solid #334155;
}

/* 自定义表格头部样式 */
.resource-table :deep(.el-table__header-wrapper) {
  background: #0f172a;
}

.resource-table :deep(.el-table__header) {
  background: #0f172a;
}

.resource-table :deep(.el-table__header th) {
  background: #0f172a !important;
  color: #fff !important;
  border-bottom: 1px solid #334155;
}

.resource-table :deep(.el-table__header th .cell) {
  color: #fff !important;
  font-weight: 600;
}

/* 表格行样式 */
.resource-table :deep(.el-table__body tr) {
  background: #0f172a;
  color: #fff;
}

.resource-table :deep(.el-table__body tr:hover) {
  background: #1e293b;
}

.resource-table :deep(.el-table__body td) {
  background: #0f172a;
  color: #fff;
  border-bottom: 1px solid #334155;
}

/* 表格边框样式 */
.resource-table :deep(.el-table) {
  border: 1px solid #334155;
  border-radius: 8px;
  overflow: hidden;
  background: #0f172a;
}

.resource-table :deep(.el-table__border-line) {
  background: #334155;
}

/* 表格背景覆盖 */
.resource-table :deep(.el-table__body-wrapper) {
  background: #0f172a;
}

.resource-table :deep(.el-table__fixed-header-wrapper) {
  background: #0f172a;
}

.resource-table :deep(.el-table__fixed-body-wrapper) {
  background: #0f172a;
}

/* 标签页样式优化 */
.resource-tabs :deep(.el-tabs__header) {
  margin-bottom: 0;
}

.resource-tabs :deep(.el-tabs__nav-wrap) {
  padding: 0;
}

.resource-tabs :deep(.el-tabs__item) {
  color: #cbd5e1;
  font-weight: 500;
}

.resource-tabs :deep(.el-tabs__item.is-active) {
  color: #409eff;
  font-weight: 600;
}

.resource-tabs :deep(.el-tabs__active-bar) {
  background-color: #409eff;
}

.resource-tabs :deep(.el-tabs__content) {
  color: #fff;
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
  background: #1e293b;
  color: #fff;
  border: 1px solid #334155;
}

.gva-pagination :deep(.el-pagination .el-pager li.is-active) {
  background: #409eff;
  color: #fff;
}

.gva-pagination :deep(.el-pagination .btn-prev),
.gva-pagination :deep(.el-pagination .btn-next) {
  background: #1e293b;
  color: #fff;
  border: 1px solid #334155;
}

/* 空数据样式 */
.resource-table :deep(.el-table__empty-block) {
  background: #0f172a;
}

.resource-table :deep(.el-table__empty-text) {
  color: #cbd5e1;
}

/* 按钮样式调整 */
.action-buttons :deep(.el-button) {
  border: 1px solid #334155;
}

.action-buttons :deep(.el-button--primary) {
  background: #409eff;
  border-color: #409eff;
}

/* 表格选择框样式 */
.resource-table :deep(.el-table__fixed-header-wrapper) {
  background: #1e293b;
}

.resource-table :deep(.el-table__fixed-body-wrapper) {
  background: #1e293b;
}

.header-item {
  display: flex;
  align-items: center;
  padding: 8px 10px;
  background: #334155;
  border: 1px solid #475569;
  border-radius: 4px;
  margin-bottom: 8px;
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
}

.header-item::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 3px;
  height: 100%;
  background: #409eff;
  transform: scaleY(0);
  transition: transform 0.3s ease;
}

.header-item:hover {
  background: #475569;
  transform: translateX(2px);
  box-shadow: 0 2px 8px rgba(64, 158, 255, 0.2);
}

.header-item:hover::before {
  transform: scaleY(1);
}

.header-item-active {
  background: #1e3a8a;
  border-color: #409eff;
  box-shadow: 0 0 0 1px #409eff;
}

.header-item-filled {
  background: linear-gradient(135deg, #1e3a8a 0%, #1e40af 100%);
  border-color: #409eff;
  box-shadow: 0 0 0 1px #409eff, 0 2px 8px rgba(64, 158, 255, 0.3);
}

.header-item-filled::before {
  transform: scaleY(1);
}

.header-index {
  width: 30px;
  text-align: center;
  margin-right: 10px;
  color: #cbd5e1;
  font-size: 14px;
  font-weight: 600;
}

.index-number {
  display: inline-block;
  width: 24px;
  height: 24px;
  line-height: 24px;
  text-align: center;
  background: #475569;
  border-radius: 50%;
  color: #fff;
  font-size: 12px;
  font-weight: 600;
}

.header-item-filled .index-number {
  background: #409eff;
  color: #fff;
}

.header-inputs {
  display: flex;
  align-items: center;
  flex-grow: 1;
  gap: 10px;
}

.input-group {
  display: flex;
  flex-direction: column;
  flex: 1;
}

.input-label {
  font-size: 12px;
  color: #cbd5e1;
  margin-bottom: 4px;
  white-space: nowrap;
  font-weight: 500;
}

.header-key-input,
.header-value-input {
  width: 100%;
  background: #334155;
  border: 1px solid #475569;
  color: #fff;
  transition: all 0.3s ease;
}

.header-key-input:focus,
.header-value-input:focus {
  border-color: #409eff;
  box-shadow: 0 0 0 2px rgba(64, 158, 255, 0.2);
  background: #475569;
}

.header-item-filled .header-key-input,
.header-item-filled .header-value-input {
  background: #1e293b;
  border-color: #409eff;
}

.header-actions {
  display: flex;
  align-items: center;
  margin-left: 10px;
}

.remove-header-btn {
  background: #f56c6c;
  border-color: #f56c6c;
  color: #fff;
  font-size: 12px;
  padding: 6px 8px;
  border-radius: 4px;
  transition: all 0.3s ease;
  opacity: 0.8;
}

.remove-header-btn:hover {
  background: #f78989;
  border-color: #f78989;
  opacity: 1;
  transform: scale(1.05);
}

.remove-header-btn:disabled {
  opacity: 0.4;
  cursor: not-allowed;
}

.headers-container {
  background: #1e293b;
  border: 1px solid #334155;
  border-radius: 8px;
  padding: 16px;
  margin-top: 10px;
  margin-bottom: 10px;
  transition: all 0.3s ease;
}

.headers-container:hover {
  border-color: #475569;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
}

.headers-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  padding-bottom: 12px;
  border-bottom: 1px solid #334155;
}

.headers-title-section {
  display: flex;
  align-items: center;
  margin-right: 10px;
}

.header-icon {
  color: #409eff;
  margin-right: 8px;
  font-size: 18px;
  animation: pulse 2s infinite;
}

@keyframes pulse {
  0%, 100% {
    opacity: 1;
  }
  50% {
    opacity: 0.7;
  }
}

.headers-title {
  font-size: 14px;
  font-weight: 600;
  color: #fff;
}

.headers-count {
  font-size: 12px;
  color: #cbd5e1;
  margin-left: 8px;
  padding: 2px 6px;
  background: #334155;
  border-radius: 10px;
}

.add-header-btn {
  background: #409eff;
  border-color: #409eff;
  color: #fff;
  font-size: 12px;
  padding: 6px 12px;
  border-radius: 6px;
  transition: all 0.3s ease;
  font-weight: 500;
}

.add-header-btn:hover {
  background: #66b1ff;
  border-color: #66b1ff;
  transform: translateY(-1px);
  box-shadow: 0 4px 8px rgba(64, 158, 255, 0.3);
}

.headers-list {
  max-height: 300px;
  overflow-y: auto;
  margin-bottom: 16px;
  padding-right: 4px;
}

.headers-list::-webkit-scrollbar {
  width: 6px;
}

.headers-list::-webkit-scrollbar-track {
  background: #334155;
  border-radius: 3px;
}

.headers-list::-webkit-scrollbar-thumb {
  background: #475569;
  border-radius: 3px;
}

.headers-list::-webkit-scrollbar-thumb:hover {
  background: #64748b;
}

.empty-headers {
  padding: 30px 20px;
  text-align: center;
  color: #cbd5e1;
}

.empty-icon {
  color: #64748b;
  font-size: 60px;
  margin-bottom: 10px;
  opacity: 0.6;
}

.headers-footer {
  margin-top: 16px;
  padding-top: 12px;
  border-top: 1px dashed #334155;
  color: #cbd5e1;
  font-size: 12px;
}

.footer-tip {
  color: #94a3b8;
  font-size: 12px;
  line-height: 1.4;
  display: flex;
  align-items: center;
  justify-content: center;
}

/* 抽屉样式调整 */
:deep(.el-drawer) {
  background: #1e293b;
  color: #fff;
}

:deep(.el-drawer__header) {
  background: #1e293b;
  color: #fff;
  border-bottom: 1px solid #334155;
}

:deep(.el-drawer__body) {
  background: #1e293b;
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

/* HTTP输入框样式 */
.http-input :deep(.el-input__wrapper) {
  background: #334155;
  border: 1px solid #475569;
  transition: all 0.3s ease;
}

.http-input :deep(.el-input__wrapper:hover) {
  border-color: #64748b;
}

.http-input :deep(.el-input__wrapper.is-focus) {
  border-color: #409eff;
  box-shadow: 0 0 0 2px rgba(64, 158, 255, 0.2);
}

.http-input :deep(.el-input__inner) {
  background: #334155;
  color: #fff;
  border: none;
}

.http-select :deep(.el-input__wrapper) {
  background: #334155;
  border: 1px solid #475569;
  transition: all 0.3s ease;
}

.http-select :deep(.el-input__wrapper:hover) {
  border-color: #64748b;
}

.http-select :deep(.el-input__wrapper.is-focus) {
  border-color: #409eff;
  box-shadow: 0 0 0 2px rgba(64, 158, 255, 0.2);
}

.http-select :deep(.el-input__inner) {
  background: #334155;
  color: #fff;
  border: none;
}
</style>
