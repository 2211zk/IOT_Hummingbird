
<template>
  <div>
    <!-- 页面标题和描述 -->
    <div class="mb-6">
      <div class="text-sm text-gray-500 mb-2">首页 / 设备接入 / 产品管理</div>
      <h1 class="text-2xl font-bold text-gray-900 mb-2">产品管理</h1>
      <p class="text-gray-600">在物联网平台中，某一类具有相同能力或特征的设备的合集被称为一款产品。</p>
    </div>

    <!-- 搜索区域 -->
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="产品名称" prop="prName">
          <el-input v-model="searchInfo.prName" placeholder="请输入产品名称" />
        </el-form-item>
        
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- 操作按钮区域 -->
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">+ 创建产品</el-button>
        <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
      </div>
    </div>

    <!-- 表格区域 -->
    <div class="gva-table">
      <el-table :data="tableData" @selection-change="handleSelectionChange" row-key="ID">
        <el-table-column type="selection" width="55" />
        <el-table-column label="产品ID" prop="ID" width="120" />
        <el-table-column label="产品名称" prop="prName" width="150" />
        <!-- 产品编号列 - 显示格式化的产品编号并提供复制功能 -->
        <el-table-column label="产品编号" prop="ID" width="150">
          <template #default="scope">
            <div class="flex items-center">
              <span>{{ generateProductNumber(scope.row.ID) }}</span>
              <el-button type="text" icon="CopyDocument" size="small" @click="copyToClipboard(generateProductNumber(scope.row.ID))" />
            </div>
          </template>
        </el-table-column>
        <el-table-column label="产品协议" prop="accessProtocol" width="120">
          <template #default="scope">
            {{ formatProtocol(scope.row.accessProtocol) }}
          </template>
        </el-table-column>
        <el-table-column label="类型" prop="nodeType" width="120">
          <template #default="scope">
            {{ formatNodeType(scope.row.nodeType) }}
          </template>
        </el-table-column>
        <el-table-column label="创建时间" prop="CreatedAt" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.CreatedAt) }}
          </template>
        </el-table-column>
        <el-table-column label="状态" prop="status" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.status === '已发布' ? 'success' : 'warning'">
              {{ scope.row.status || '未发布' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="300">
          <template #default="scope">
            <el-button type="primary" link icon="view" @click="openDetail(scope.row)">详情</el-button>
            <el-button type="primary" link icon="edit" @click="updateWlProductsFunc(scope.row)">编辑</el-button>
            <el-button type="primary" link icon="setting" @click="manageDevices(scope.row)">管理设备</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <!-- 分页区域 -->
    <div class="gva-pagination">
      <el-pagination
        layout="total, sizes, prev, pager, next, jumper"
        :current-page="page"
        :page-size="pageSize"
        :page-sizes="[20, 50, 100, 200]"
        :total="total"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
        background
      />
    </div>

    <!-- 新增/编辑抽屉 -->
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{type==='create'?'创建产品':'编辑产品'}}</span>
          <div>
            <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>

      <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
        <!-- 产品名称 -->
        <el-form-item label="产品名称" prop="prName" required>
          <el-input v-model="formData.prName" :clearable="true" placeholder="请输入产品名称" />
        </el-form-item>

        <!-- 所属品类 -->
        <el-form-item label="所属品类" prop="prCategory" required>
          <el-radio-group v-model="formData.prCategory" class="w-full">
            <el-radio label="标准品类">标准品类</el-radio>
            <el-radio label="自定义品类">自定义品类</el-radio>
          </el-radio-group>
          <div v-if="formData.prCategory === '标准品类'" class="text-xs text-gray-500 mt-1">
            标准品类为系统预定义的设备类型，自定义品类为用户自定义的设备类型
          </div>
          <div v-else-if="formData.prCategory === '自定义品类'" class="text-xs text-gray-500 mt-1">
            选择自定义品类时,需从系统生成的品类列表中选择一个品类
          </div>
        </el-form-item>

        <!-- 选择标准品类 - 仅在选择标准品类时显示 -->
        <!-- 使用只读输入框配合选择按钮，点击后弹出品类选择弹窗 -->
        <el-form-item v-if="formData.prCategory === '标准品类'" label="选择标准品类" prop="standardQuality" required>
          <el-input 
            v-model="formData.standardQualityName" 
            placeholder="请选择" 
            readonly 
            @click="openCategoryDialog"
            style="width:100%"
          >
            <template #append>
              <el-button @click="openCategoryDialog">选择</el-button>
            </template>
          </el-input>
        </el-form-item>

        <!-- 节点类型 -->
        <el-form-item label="节点类型" prop="nodeType" required>
          <el-select v-model="formData.nodeType" placeholder="请选择" style="width:100%" filterable :clearable="true">
            <el-option label="直连设备" value="直连设备" />
            <el-option label="网关设备" value="网关设备" />
            <el-option label="监控设备" value="监控设备" />
            <el-option label="传感器设备" value="传感器设备" />
            <el-option label="控制器设备" value="控制器设备" />
          </el-select>
        </el-form-item>

        <!-- 接入协议 -->
        <el-form-item label="接入协议" prop="accessProtocol" required>
          <el-select v-model="formData.accessProtocol" placeholder="请选择" style="width:100%" filterable :clearable="true">
            <el-option label="MQTT" value="MQTT" />
            <el-option label="ModbusRTU" value="ModbusRTU" />
            <el-option label="ModbusTCP" value="ModbusTCP" />
            <el-option label="GB28181" value="GB28181" />
            <el-option label="TCP" value="TCP" />
            <el-option label="HTTP" value="HTTP" />
            <el-option label="CoAP" value="CoAP" />
          </el-select>
        </el-form-item>

        <!-- 数据格式 -->
        <el-form-item label="数据格式" prop="dataFormat" required>
          <el-select v-model="formData.dataFormat" placeholder="请选择" style="width:100%" filterable :clearable="true">
            <el-option label="标准物模型" value="标准物模型" />
            <el-option label="自定义格式" value="自定义格式" />
            <el-option label="JSON格式" value="JSON格式" />
            <el-option label="XML格式" value="XML格式" />
            <el-option label="二进制格式" value="二进制格式" />
          </el-select>
        </el-form-item>

        <!-- 网络类型 -->
        <el-form-item label="网络类型" prop="networkType" required>
          <el-radio-group v-model="formData.networkType" class="w-full">
            <el-radio label="以太网">以太网</el-radio>
            <el-radio label="蜂窝">蜂窝</el-radio>
            <el-radio label="WIFI">WIFI</el-radio>
            <el-radio label="NB">NB</el-radio>
            <el-radio label="其他">其他</el-radio>
          </el-radio-group>
        </el-form-item>

        <!-- 工厂 -->
        <el-form-item label="工厂" prop="factory">
          <el-input v-model="formData.factory" :clearable="true" placeholder="请输入" />
        </el-form-item>

        <!-- 产品描述 -->
        <el-form-item label="产品描述" prop="prInfo">
          <el-input v-model="formData.prInfo" :clearable="true" placeholder="请输入产品描述" type="textarea" :rows="3" />
        </el-form-item>
      </el-form>
    </el-drawer>

    <!-- 选择品类弹窗 - 嵌套在创建产品弹窗中的品类选择界面 -->
    <!-- 功能：提供品类搜索、列表展示、分页和选择功能 -->
    <el-dialog v-model="categoryDialogVisible" title="选择品类" width="60%" :before-close="closeCategoryDialog">
      <!-- 品类搜索区域 -->
      <div class="mb-4">
        <el-form :inline="true" :model="categorySearchInfo" class="demo-form-inline">
          <el-form-item label="输入品类名称">
            <el-input v-model="categorySearchInfo.name" placeholder="输入品类名称" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="searchCategories">查询</el-button>
            <el-button @click="resetCategorySearch">重置</el-button>
          </el-form-item>
        </el-form>
      </div>

      <!-- 品类列表表格 -->
      <el-table :data="categoryList" @row-click="selectCategory" style="width: 100%">
        <el-table-column prop="name" label="品类名称" />
        <el-table-column prop="scene" label="所属场景" />
        <el-table-column label="操作" width="200">
          <template #default="scope">
            <el-button type="primary" link @click="selectCategory(scope.row)">选择</el-button>
            <el-button type="primary" link @click="viewStandardModel(scope.row)">查看标准物模型</el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 品类分页控制 -->
      <div class="mt-4">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="categoryPage"
          :page-size="categoryPageSize"
          :page-sizes="[10, 20, 50, 100]"
          :total="categoryTotal"
          @current-change="handleCategoryPageChange"
          @size-change="handleCategorySizeChange"
          background
        />
      </div>

      <template #footer>
        <span class="dialog-footer">
          <el-button @click="closeCategoryDialog">取消</el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 标准功能定义对话框 -->
    <StandardFunctionDialog 
      v-model="standardFunctionDialogVisible"
      :category-id="selectedCategoryId"
      :category-name="selectedCategoryName"
      @select="handleFunctionSelect"
    />

    <!-- 详情抽屉 -->
    <el-drawer destroy-on-close :size="'80%'" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="产品详情">
      <!-- 产品详情头部 -->
      <div class="mb-6">
        <div class="text-sm text-gray-500 mb-2">首页 / 产品管理 / 产品详情 [{{ detailFrom.prName }}]</div>
        <h1 class="text-2xl font-bold text-gray-900 mb-2">产品详情</h1>
        <p class="text-gray-600 mb-4">产品物模型是设备在数字世界中的抽象描述，包含属性、事件和服务。</p>
        <div class="flex items-center gap-2">
          <span class="text-lg font-semibold">{{ detailFrom.prName }}</span>
          <el-tag :type="detailFrom.status === '已发布' ? 'success' : 'warning'">
            {{ detailFrom.status || '未发布' }}
          </el-tag>
        </div>
      </div>

      <!-- 产品基本信息 -->
      <div class="mb-6">
        <h2 class="text-lg font-semibold mb-4">产品基本信息</h2>
        <el-descriptions :column="2" border>
          <el-descriptions-item label="产品ID">
            {{ detailFrom.ID }}
          </el-descriptions-item>
          <el-descriptions-item label="产品名称">
            {{ detailFrom.prName }}
          </el-descriptions-item>
          <el-descriptions-item label="产品标识">
            {{ detailFrom.ID ? `ct${detailFrom.ID.toString().padStart(8, '0')}` : '-' }}
          </el-descriptions-item>
          <el-descriptions-item label="产品类型">
            {{ formatNodeType(detailFrom.nodeType) }}
          </el-descriptions-item>
          <el-descriptions-item label="网络类型">
            {{ detailFrom.networkType }}
          </el-descriptions-item>
          <el-descriptions-item label="数据类型">
            {{ detailFrom.dataFormat }}
          </el-descriptions-item>
          <el-descriptions-item label="工厂名称">
            {{ detailFrom.factory || '-' }}
          </el-descriptions-item>
          <el-descriptions-item label="产品状态">
            <el-tag :type="detailFrom.status === '已发布' ? 'success' : 'warning'">
              {{ detailFrom.status || '未发布' }}
            </el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="是否为IPC产品">
            否
          </el-descriptions-item>
          <el-descriptions-item label="产品协议">
            {{ formatProtocol(detailFrom.accessProtocol) }}
          </el-descriptions-item>
          <el-descriptions-item label="创建时间">
            {{ formatDate(detailFrom.CreatedAt) }}
          </el-descriptions-item>
          <el-descriptions-item label="更新时间">
            {{ formatDate(detailFrom.UpdatedAt) }}
          </el-descriptions-item>
          <el-descriptions-item label="描述" :span="2">
            {{ detailFrom.prInfo || '-' }}
          </el-descriptions-item>
        </el-descriptions>
      </div>
    </el-drawer>
  </div>
</template>

<script setup>
import { 
  createWlProducts, 
  deleteWlProducts, 
  deleteWlProductsByIds, 
  updateWlProducts, 
  findWlProducts, 
  getWlProductsList
} from '@/api/wl_playform/wlProducts'

import { getWlCategoryList, getWlCategoryPublic } from '@/api/wl_playform/wlCategory'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { Box } from '@element-plus/icons-vue'
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'
import { useAppStore } from "@/pinia"

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'
// 标准功能定义组件
import StandardFunctionDialog from './components/StandardFunctionDialog.vue'

defineOptions({
    name: 'WlProducts'
})
// 按钮权限实例化
const btnAuth = useBtnAuth()

const appStore = useAppStore()
const router = useRouter()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const Node_typeOptions = ref([])
const access_protocolOptions = ref([])
const Network_typeOptions = ref([])
const CategoryOptions = ref([])
const data_formatOptions = ref([])

// 表单数据
const formData = ref({
    prName: '',
    prCategory: '',
    standardQuality: null, // 标准品类选择，使用数字类型
    standardQualityName: '', // 标准品类名称显示
    nodeType: '',
    accessProtocol: '',
    dataFormat: '',
    networkType: '',
    factory: '',
    prInfo: '',
})

// 验证规则
const rule = reactive({
    prName : [{
        required: true,
        message: '请输入产品名称',
        trigger: ['input','blur'],
    },
    {
        whitespace: true,
        message: '不能只输入空格',
        trigger: ['input', 'blur'],
    }
    ],
    prCategory : [{
        required: true,
        message: '请选择所属品类',
        trigger: ['change'],
    }],
    standardQuality : [{
        required: true,
        message: '请选择标准品类',
        trigger: ['change'],
    }],
    nodeType : [{
        required: true,
        message: '请选择节点类型',
        trigger: ['change'],
    }],
    accessProtocol : [{
        required: true,
        message: '请选择接入协议',
        trigger: ['change'],
    }],
    dataFormat : [{
        required: true,
        message: '请选择数据格式',
        trigger: ['change'],
    }],
    networkType : [{
        required: true,
        message: '请选择网络类型',
        trigger: ['change'],
    }],
})

// 品类选择相关 - 控制选择品类弹窗的状态和数据
// 功能：管理品类选择弹窗的显示、品类列表数据、搜索条件和分页信息
const categoryDialogVisible = ref(false)
const categoryList = ref([])
const categorySearchInfo = ref({ name: '' })
const categoryPage = ref(1)
const categoryPageSize = ref(10)
const categoryTotal = ref(0)

// 标准功能定义相关
const standardFunctionDialogVisible = ref(false)
const selectedCategoryId = ref(null)
const selectedCategoryName = ref('')

// 基础数据
const elFormRef = ref()
const elSearchFormRef = ref()
const dialogFormVisible = ref(false)
const type = ref('')
const btnLoading = ref(false)
const detailShow = ref(false)
const detailFrom = ref({})

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(50) // 增加默认每页显示数量，显示更多数据
const tableData = ref([])
const searchInfo = ref({})

// 多选数据
const multipleSelection = ref([])

// 获取表格数据 - 从后端API获取产品列表数据
// 功能：支持分页查询、搜索过滤，并处理不同的数据格式
const getTableData = async() => {
  try {
    // 构建请求参数，确保分页参数正确传递
    const params = {
      page: page.value,
      pageSize: pageSize.value,
      ...searchInfo.value
    }
    
    const table = await getWlProductsList(params)
    
    if (table.code === 0) {
      // 确保数据是数组格式 - 处理不同API返回的数据结构
      if (Array.isArray(table.data.list)) {
        tableData.value = table.data.list
      } else if (table.data.list && Array.isArray(table.data.list.list)) {
        // 如果数据在嵌套的list中 - 兼容嵌套数据结构
        tableData.value = table.data.list.list
      } else if (Array.isArray(table.data)) {
        // 如果数据直接在data中
        tableData.value = table.data
      } else {
        tableData.value = []
      }
      
      // 确保分页信息正确设置
      total.value = table.data.total || table.data.totalCount || 0
      page.value = table.data.page || table.data.currentPage || 1
      pageSize.value = table.data.pageSize || table.data.pageSize || 50
    } else {
      console.error('API调用失败:', table)
      ElMessage.error(`API调用失败: ${table.msg || '未知错误'}`)
    }
  } catch (error) {
    console.error('获取数据时出错:', error)
    ElMessage.error('获取数据时出错')
  }
}



// 删除行 - 删除单个产品记录
// 功能：弹出确认对话框，用户确认后删除指定产品
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteWlProductsFunc(row)
  })
}

// 多选删除 - 批量删除选中的产品记录
// 功能：检查选中状态，收集选中项的ID，调用批量删除API
const onDelete = async() => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
  multipleSelection.value.map(item => {
    ids.push(item.ID)
  })
  const res = await deleteWlProductsByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功!'
    })
    // 如果当前页数据全部被删除且不是第一页，则跳转到上一页
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    await getTableData()
  }
}

// 搜索 - 提交搜索条件并重新获取数据
// 功能：重置到第一页，保持每页显示数量，执行搜索查询
const onSubmit = () => {
  page.value = 1
  pageSize.value = 50 // 使用更大的默认页面大小
  getTableData()
}

// 重置 - 清空搜索条件并重新获取数据
// 功能：清空所有搜索条件，重置到第一页，获取所有数据
const onReset = () => {
  searchInfo.value = {}
  page.value = 1
  pageSize.value = 50
  getTableData()
}

// 分页处理函数
// 改变每页显示数量
const handleSizeChange = (val) => {
  pageSize.value = val
  page.value = 1 // 重置到第一页
  getTableData()
}

// 改变当前页码
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    Node_typeOptions.value = await getDictFunc('Node_type')
    access_protocolOptions.value = await getDictFunc('access_protocol')
    Network_typeOptions.value = await getDictFunc('Network_type')
    CategoryOptions.value = await getDictFunc('Category')
    data_formatOptions.value = await getDictFunc('data_format')
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 表格选择处理
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 生成产品编号
const generateProductNumber = (id) => {
  // 确保id是字符串类型
  const idStr = String(id)
  return `WL-${idStr.padStart(8, '0')}`
}

// 复制到剪贴板 - 复制产品编号到系统剪贴板
// 功能：使用浏览器Clipboard API复制文本，提供用户友好的成功/失败提示
const copyToClipboard = (text) => {
  navigator.clipboard.writeText(text).then(() => {
    ElMessage({
      type: 'success',
      message: '已复制到剪贴板'
    })
  }).catch(() => {
    ElMessage({
      type: 'error',
      message: '复制失败'
    })
  })
}

// 格式化产品协议
const formatProtocol = (value) => {
  // 直接映射数字到协议名称
  const protocolMap = {
    1: 'MQTT',
    2: 'ModbusRTU', 
    3: 'ModbusTCP',
    4: 'GB28181',
    5: 'TCP',
    6: 'HTTP',
    7: 'CoAP',
    '1': 'MQTT',
    '2': 'ModbusRTU', 
    '3': 'ModbusTCP',
    '4': 'GB28181',
    '5': 'TCP',
    '6': 'HTTP',
    '7': 'CoAP'
  }
  
  // 如果字典已加载，优先使用字典
  if (access_protocolOptions.value && access_protocolOptions.value.length > 0) {
    const dict = access_protocolOptions.value.find(item => item.value === value)
    if (dict) return dict.label
  }
  
  // 否则使用直接映射
  return protocolMap[value] || value
}

// 格式化节点类型
const formatNodeType = (value) => {
  // 直接映射数字到节点类型名称
  const nodeTypeMap = {
    1: '直连设备',
    2: '网关设备',
    3: '监控设备',
    4: '传感器设备',
    5: '控制器设备',
    '1': '直连设备',
    '2': '网关设备',
    '3': '监控设备',
    '4': '传感器设备',
    '5': '控制器设备'
  }

  // 如果字典已加载，优先使用字典
  if (Node_typeOptions.value && Node_typeOptions.value.length > 0) {
    const dict = Node_typeOptions.value.find(item => item.value === value)
    if (dict) return dict.label
  }

  // 否则使用直接映射
  return nodeTypeMap[value] || value
}

// 更新行
const updateWlProductsFunc = async(row) => {
    const res = await findWlProducts({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}

// 删除行
const deleteWlProductsFunc = async(row) => {
    const res = await deleteWlProducts({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
            type: 'success',
            message: '删除成功!'
        })
        if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        await getTableData()
    }
}

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        prName: '',
        prCategory: '',
        standardQuality: null,
        standardQualityName: '',
        nodeType: '',
        accessProtocol: '',
        dataFormat: '',
        networkType: '',
        factory: '',
        prInfo: '',
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
                  res = await createWlProducts(formData.value)
                  break
                case 'update':
                  res = await updateWlProducts(formData.value)
                  break
                default:
                  res = await createWlProducts(formData.value)
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

// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}

// 打开详情
const openDetail = (row) => {
  detailFrom.value = row
  detailShow.value = true
}

// 关闭详情
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}

// 管理设备 - 跳转到设备管理页面
const manageDevices = (row) => {
  console.log('管理设备 - 产品数据:', row)
  
  // 构建跳转参数
  const params = {
    productId: row.ID,
    productName: row.productName || row.ProductName,
    productNumber: row.productNumber || row.ProductNumber
  }
  
  // 使用路由跳转到设备管理页面
  router.push({
    name: 'wlEquipment',
    query: params
  })
}

// 品类选择相关函数 - 处理品类选择弹窗的各种操作
// 打开品类选择弹窗
const openCategoryDialog = () => {
  categoryDialogVisible.value = true
  loadCategoryList()
}

// 关闭品类选择弹窗
const closeCategoryDialog = () => {
  categoryDialogVisible.value = false
  categorySearchInfo.value = { name: '' }
}

// 加载品类列表数据 - 从后端API获取品类数据
// 功能：调用真实的API接口获取品类数据，支持分页和搜索
const loadCategoryList = async () => {
  try {
    console.log('开始加载品类数据...')
    
    // 构建请求参数
    const params = {
      page: categoryPage.value,
      pageSize: categoryPageSize.value,
      name: categorySearchInfo.value.name
    }
    
    console.log('品类API请求参数:', params)
    
    // 调用公开API（无需权限）
    const response = await getWlCategoryPublic(params)
    
    console.log('品类API响应:', response)
    
    if (response.code === 0) {
      // 处理API返回的数据
      let dataList = []
      if (Array.isArray(response.data.list)) {
        dataList = response.data.list
      } else if (response.data.list && Array.isArray(response.data.list.list)) {
        dataList = response.data.list.list
      } else if (Array.isArray(response.data)) {
        dataList = response.data
      }
      
      // 转换字段名以适配前端显示
      categoryList.value = dataList.map(item => ({
        id: item.ID || item.id, // 确保ID字段正确
        name: item.caName || item.name, // 支持两种字段名
        key: item.caKey || item.key,
        scenario: item.caScenario || item.scenario,
        created_at: item.created_at,
        updated_at: item.updated_at
      }))
      
      categoryTotal.value = response.data.total || response.data.totalCount || 0
      categoryPage.value = response.data.page || response.data.currentPage || 1
      categoryPageSize.value = response.data.pageSize || 10
      
      console.log('品类数据加载成功:', categoryList.value.length, '条')
      console.log('品类数据示例:', categoryList.value.slice(0, 3))
    } else {
      console.log('品类API调用失败:', response.msg)
      ElMessage.error(`品类数据加载失败: ${response.msg || '未知错误'}`)
      categoryList.value = []
      categoryTotal.value = 0
    }
    
  } catch (error) {
    console.error('加载品类数据失败:', error)
    ElMessage.error('加载品类数据失败')
    categoryList.value = []
    categoryTotal.value = 0
  }
}

// 搜索品类 - 根据搜索条件过滤品类列表
const searchCategories = () => {
  categoryPage.value = 1 // 重置到第一页
  loadCategoryList()
}

// 重置品类搜索 - 清空搜索条件并重新加载品类列表
const resetCategorySearch = () => {
  categorySearchInfo.value = { name: '' }
  categoryPage.value = 1 // 重置到第一页
  loadCategoryList()
}

// 选择品类 - 用户点击选择某个品类后的处理
// 功能：将选中的品类信息填充到表单中，并关闭弹窗
const selectCategory = (row) => {
  formData.value.standardQuality = row.id
  formData.value.standardQualityName = row.name
  closeCategoryDialog()
}

// 查看标准物模型 - 查看指定品类的标准物模型定义
// 功能：显示品类的标准功能定义信息
const viewStandardModel = (row) => {
  console.log('查看标准物模型 - 品类数据:', row)
  console.log('品类数据所有字段:', Object.keys(row))
  
  // 尝试多种可能的ID字段
  const categoryId = row.ID || row.id || row.CaId || row.caId
  const categoryName = row.caName || row.name || row.CaName
  
  selectedCategoryId.value = categoryId
  selectedCategoryName.value = categoryName
  
  console.log('设置品类ID:', selectedCategoryId.value, '类型:', typeof selectedCategoryId.value)
  console.log('设置品类名称:', selectedCategoryName.value)
  
  standardFunctionDialogVisible.value = true
}

// 处理功能选择 - 当用户在标准功能定义对话框中选择功能时的处理
// 功能：接收用户选择的功能定义，并可以进行后续处理
const handleFunctionSelect = (functionData) => {
  const functionName = functionData.functionName || functionData.function_name
  const identifier = functionData.identifier
  
  ElMessage({
    type: 'success',
    message: `已选择功能: ${functionName} (${identifier})`
  })
  
  // 这里可以根据需要将选择的功能添加到产品定义中
  // 例如：将功能添加到产品的物模型列表中
  console.log('选择的功能:', functionData)
}

// 处理品类分页变化 - 当用户切换品类列表页码时的处理
const handleCategoryPageChange = (val) => {
  categoryPage.value = val
  loadCategoryList()
}

// 处理品类分页大小变化 - 当用户改变每页显示数量时的处理
const handleCategorySizeChange = (val) => {
  categoryPageSize.value = val
  categoryPage.value = 1 // 重置到第一页
  loadCategoryList()
}



// 初始化
getTableData()

// 初始化字典
const initDict = async () => {
  Node_typeOptions.value = await getDictFunc('Node_type')
  access_protocolOptions.value = await getDictFunc('access_protocol')
  Network_typeOptions.value = await getDictFunc('Network_type')
  CategoryOptions.value = await getDictFunc('Category')
  data_formatOptions.value = await getDictFunc('data_format')
}

initDict()
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
