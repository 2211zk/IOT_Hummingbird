
<template>
  <div>
    <!-- 页面标题和描述 -->
    <div class="mb-6">
      <div class="text-sm text-gray-500 mb-2">首页 / 设备接入 / 设备管理</div>
      <h1 class="text-2xl font-bold text-gray-900 mb-2">设备管理</h1>
      <p class="text-gray-600">物理设备要连接到平台,需要先在平台创建设备(支持单个或批量导入创建)</p>
    </div>

    <!-- 搜索和筛选区域 -->
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="产品范围" prop="productScope">
          <el-select v-model="searchInfo.productScope" placeholder="产品范围(全部)" style="width: 150px">
            <el-option label="全部" value="" />
            <el-option label="标准产品" value="standard" />
            <el-option label="自定义产品" value="custom" />
          </el-select>
        </el-form-item>
        
        <el-form-item label="产品类型" prop="productType">
          <el-select v-model="searchInfo.productType" placeholder="产品类型(全部)" style="width: 150px">
            <el-option label="全部" value="" />
            <el-option label="直连设备" value="direct" />
            <el-option label="网关设备" value="gateway" />
          </el-select>
        </el-form-item>
        
        <el-form-item label="状态" prop="status">
          <el-select v-model="searchInfo.status" placeholder="状态(全部)" style="width: 150px">
            <el-option label="全部" value="" />
            <el-option label="在线" value="online" />
            <el-option label="离线" value="offline" />
          </el-select>
        </el-form-item>
        
        <el-form-item label="设备名称" prop="eqName">
          <el-input v-model="searchInfo.eqName" placeholder="请输入设备名称" style="width: 200px" />
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
        <el-button v-auth="btnAuth.add" type="primary" icon="plus" @click="openDialog()">+ 添加设备</el-button>
        <el-button icon="refresh" style="margin-left: 10px;" @click="getTableData">刷新</el-button>
        <ExportTemplate v-auth="btnAuth.exportTemplate" template-id="wl_playform_WlEquipment" />
        <ExportExcel v-auth="btnAuth.exportExcel" template-id="wl_playform_WlEquipment" filterDeleted/>
        <ImportExcel v-auth="btnAuth.importExcel" template-id="wl_playform_WlEquipment" @on-success="getTableData" />
      </div>

      <!-- 设备列表表格 -->
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="设备ID" prop="ID" width="120" />
        
        <el-table-column align="left" label="设备名称" prop="eqName" width="150" />

        <el-table-column align="left" label="所属产品" prop="productsId" width="120">
          <template #default="scope">
            {{ getProductName(scope.row.productsId) }}
          </template>
        </el-table-column>

        <el-table-column align="left" label="设备状态" prop="status" width="100">
          <template #default="scope">
            <div class="flex items-center">
              <div class="w-2 h-2 rounded-full mr-2" :class="scope.row.status === 'online' ? 'bg-green-500' : 'bg-red-500'"></div>
              <span :class="scope.row.status === 'online' ? 'text-green-600' : 'text-red-600'">
                {{ scope.row.status === 'online' ? '在线' : '离线' }}
              </span>
            </div>
          </template>
        </el-table-column>

        <el-table-column align="left" label="创建时间" prop="CreatedAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>

        <el-table-column align="left" label="上线时间" prop="onlineTime" width="180">
          <template #default="scope">
            {{ scope.row.status === 'online' ? formatDate(scope.row.onlineTime) : '-' }}
          </template>
        </el-table-column>

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
          <template #default="scope">
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button" @click="getDetails(scope.row)">详情</el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button" @click="updateWlEquipmentFunc(scope.row)">编辑</el-button>
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

    <!-- 新增/编辑抽屉 -->
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{type==='create'?'创建设备':'编辑设备'}}</span>
          <div>
            <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>

      <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
        <!-- 添加设备方式 -->
        <el-form-item label="添加设备方式" prop="addMethod">
          <el-radio-group v-model="formData.addMethod" class="w-full">
            <el-radio label="single">单个设备</el-radio>
            <el-radio label="batch">批量添加</el-radio>
          </el-radio-group>
          <div v-if="formData.addMethod === 'batch'" class="text-xs text-gray-500 mt-2">
            批量添加设备可在批次列表中查询相关记录 格式.xlsx 最大2M,单次500个设备 
            <el-button type="text" size="small" @click="downloadTemplate">模板下载</el-button>
          </div>
        </el-form-item>

        <!-- 单个设备表单 -->
        <div v-if="formData.addMethod === 'single'">
          <!-- 设备名称 -->
          <el-form-item label="设备名称" prop="eqName" required>
            <el-input v-model="formData.eqName" :clearable="true" placeholder="请输入设备名称" />
          </el-form-item>
          
          <!-- 设备唯一标识 -->
          <el-form-item label="设备唯一标识" prop="eqLogotype" required>
            <div class="flex items-center">
              <el-input v-model="formData.eqLogotype" :clearable="true" placeholder="请输入设备唯一标识" class="flex-1" />
              <el-button type="text" icon="refresh" @click="generateUUID" class="ml-2">刷新</el-button>
            </div>
          </el-form-item>
          
          <!-- 所属产品 -->
          <el-form-item label="所属产品" prop="productsId" required>
            <el-select v-model="formData.productsId" placeholder="请选择所属产品" style="width:100%" filterable :clearable="true">
              <el-option v-for="product in productOptions" :key="product.ID" :label="product.prName" :value="product.ID" />
            </el-select>
          </el-form-item>
          
<<<<<<< HEAD
          <!-- 关联驱动 -->
          <el-form-item label="关联驱动" prop="driveId">
            <el-select v-model="formData.driveId" placeholder="请选择关联驱动" style="width:100%" filterable :clearable="true">
              <el-option v-for="driver in driverOptions" :key="driver.ID" :label="`${driver.driverId || driver.ID} - ${driver.driverName}`" :value="driver.ID" />
            </el-select>
          </el-form-item>
          
=======
>>>>>>> kai
          <!-- 设备坐标 -->
          <el-form-item label="设备坐标" prop="eqCoordinate">
            <el-input 
              v-model="formData.eqCoordinate" 
              placeholder="请选择设备坐标" 
              readonly 
              @click="openMapDialog"
              style="width: 100%; cursor: pointer;"
            />
          </el-form-item>
          
          <!-- 设备详细地址 -->
          <el-form-item label="设备详细地址" prop="eqAddress">
            <el-input v-model="formData.eqAddress" :clearable="true" placeholder="请输入设备详细地址" />
          </el-form-item>
          
          <!-- 设备描述 -->
          <el-form-item label="设备描述" prop="eqInfo">
            <el-input v-model="formData.eqInfo" :clearable="true" placeholder="请输入产品描述" type="textarea" :rows="3" />
          </el-form-item>
        </div>

        <!-- 批量添加表单 -->
        <div v-if="formData.addMethod === 'batch'">
          <!-- 所属产品 -->
          <el-form-item label="所属产品" prop="batchProductsId" required>
            <el-select v-model="formData.batchProductsId" placeholder="请选择所属产品" style="width:100%" filterable :clearable="true">
              <el-option v-for="product in productOptions" :key="product.ID" :label="product.prName" :value="product.ID" />
            </el-select>
          </el-form-item>
          
<<<<<<< HEAD
          <!-- 关联驱动 -->
          <el-form-item label="关联驱动" prop="batchDriveId" required>
            <el-select v-model="formData.batchDriveId" placeholder="请选择关联驱动" style="width:100%" filterable :clearable="true">
              <el-option v-for="driver in driverOptions" :key="driver.ID" :label="`${driver.driverId || driver.ID} - ${driver.driverName}`" :value="driver.ID" />
            </el-select>
          </el-form-item>
          
=======
>>>>>>> kai
          <!-- 上传设备表 -->
          <el-form-item label="上传设备表" prop="batchFile" required>
            <el-upload
              ref="uploadRef"
              :auto-upload="false"
              :on-change="handleFileChange"
              :on-remove="handleFileRemove"
              :file-list="fileList"
              accept=".xlsx,.xls"
              :limit="1"
              drag
            >
              <el-icon class="el-icon--upload"><upload-filled /></el-icon>
              <div class="el-upload__text">
                将文件拖到此处，或<em>点击上传</em>
              </div>
              <template #tip>
                <div class="el-upload__tip">
                  只能上传 xlsx/xls 文件，且不超过 2MB
                </div>
              </template>
            </el-upload>
          </el-form-item>
        </div>
      </el-form>
    </el-drawer>

    <!-- 详情抽屉 -->
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="设备详情">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="设备ID">
          {{ detailFrom.ID }}
        </el-descriptions-item>
        <el-descriptions-item label="设备名称">
          {{ detailFrom.eqName }}
        </el-descriptions-item>
        <el-descriptions-item label="设备唯一标识">
          {{ detailFrom.eqLogotype }}
        </el-descriptions-item>
        <el-descriptions-item label="所属产品">
          {{ getProductName(detailFrom.productsId) }}
        </el-descriptions-item>
        <el-descriptions-item label="设备状态">
          <div class="flex items-center">
            <div class="w-2 h-2 rounded-full mr-2" :class="detailFrom.status === 'online' ? 'bg-green-500' : 'bg-red-500'"></div>
            <span :class="detailFrom.status === 'online' ? 'text-green-600' : 'text-red-600'">
              {{ detailFrom.status === 'online' ? '在线' : '离线' }}
            </span>
          </div>
        </el-descriptions-item>
        <el-descriptions-item label="设备坐标">
          {{ getCoordinateName(detailFrom.eqCoordinate) || '-' }}
        </el-descriptions-item>
        <el-descriptions-item label="设备详细地址">
          {{ detailFrom.eqAddress || '-' }}
        </el-descriptions-item>
        <el-descriptions-item label="设备描述">
          {{ detailFrom.eqInfo || '-' }}
        </el-descriptions-item>
        <el-descriptions-item label="创建时间">
          {{ formatDate(detailFrom.CreatedAt) }}
        </el-descriptions-item>
        <el-descriptions-item label="上线时间">
          {{ detailFrom.status === 'online' ? formatDate(detailFrom.onlineTime) : '-' }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>

    <!-- 地图选点对话框 -->
    <el-dialog v-model="mapDialogVisible" title="地图选点" width="80%" :before-close="closeMapDialog">
      <div class="map-container">
        <!-- 搜索输入框 -->
        <div class="mb-4">
          <el-input 
            v-model="mapSearchKeyword" 
            placeholder="请输入地点搜索" 
            @keyup.enter="searchLocation"
            style="width: 300px"
          >
            <template #append>
              <el-button @click="searchLocation">搜索</el-button>
            </template>
          </el-input>
        </div>

        <!-- 坐标输入框 -->
        <div class="mb-4 flex gap-4">
          <el-input v-model="selectedLongitude" placeholder="经度" style="width: 200px" />
          <el-input v-model="selectedLatitude" placeholder="纬度" style="width: 200px" />
        </div>

        <!-- 地图容器 -->
        <div class="map-wrapper" style="height: 400px; border: 1px solid #dcdfe6; border-radius: 4px;">
          <div id="map-container" style="width: 100%; height: 100%; position: relative;"></div>
        </div>

        <!-- 详细地址 -->
        <div class="mt-4">
          <el-input v-model="selectedAddress" placeholder="详细地址" />
        </div>
      </div>

      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeMapDialog">取消</el-button>
          <el-button type="primary" @click="confirmMapSelection">确定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createWlEquipment,
  deleteWlEquipment,
  deleteWlEquipmentByIds,
  updateWlEquipment,
  findWlEquipment,
  getWlEquipmentList
} from '@/api/wl_playform/wlEquipment'

// 导入产品API
import { getWlProductsList } from '@/api/wl_playform/wlProducts'
// 导入驱动API
import { getWlDriversList } from '@/api/wl_driver/wlDrivers'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, nextTick, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { UploadFilled } from '@element-plus/icons-vue'
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
    name: 'WlEquipment'
})
// 按钮权限实例化
const btnAuth = useBtnAuth()

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()
const route = useRoute()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 产品选项
const productOptions = ref([])

const formData = ref({
            eqName: '',
            eqLogotype: '',
            productsId: undefined,
            eqCoordinate: '',
            eqAddress: '',
            eqInfo: '',
            status: 'offline',
            onlineTime: null,
            addMethod: 'single',
            // 批量添加字段
            batchProductsId: undefined,
            batchFile: null,
        })

// 文件上传相关
const uploadRef = ref()
const fileList = ref([])

// 地图相关变量
const mapDialogVisible = ref(false)
const mapSearchKeyword = ref('')
const selectedLongitude = ref('')
const selectedLatitude = ref('')
const selectedAddress = ref('')
const mapInstance = ref(null)
const mapMarker = ref(null)
const baiduMapLoaded = ref(false)

// 生成UUID
const generateUUID = () => {
  const uuid = 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
    const r = Math.random() * 16 | 0
    const v = c === 'x' ? r : (r & 0x3 | 0x8)
    return v.toString(16)
  })
  formData.value.eqLogotype = uuid
}

// 下载模板
const downloadTemplate = () => {
  // 创建一个示例Excel模板
  const templateData = [
    ['设备名称', '设备唯一标识', '设备坐标', '设备详细地址', '设备描述'],
    ['示例设备1', 'uuid-1', '北京市朝阳区', '北京市朝阳区xxx街道xxx号', '示例设备描述1'],
    ['示例设备2', 'uuid-2', '上海市浦东新区', '上海市浦东新区xxx街道xxx号', '示例设备描述2'],
  ]
  
  // 这里应该调用实际的下载API
  ElMessage({
    type: 'success',
    message: '模板下载功能待实现'
  })
}

// 文件变化处理
const handleFileChange = (file) => {
  // 检查文件大小（2MB限制）
  const isLt2M = file.size / 1024 / 1024 < 2
  if (!isLt2M) {
    ElMessage.error('上传文件大小不能超过 2MB!')
    return false
  }
  
  // 检查文件类型
  const isExcel = file.type === 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet' || 
                  file.type === 'application/vnd.ms-excel'
  if (!isExcel) {
    ElMessage.error('只能上传 Excel 文件!')
    return false
  }
  
  formData.value.batchFile = file.raw
  return true
}

// 文件移除处理
const handleFileRemove = () => {
  formData.value.batchFile = null
}

// 验证规则
const rule = reactive({
               eqName : [{
                   required: true,
                   message: '请输入设备名称',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               eqLogotype : [{
                   required: true,
                   message: '请输入设备唯一标识',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               productsId : [{
                   required: true,
                   message: '请选择所属产品',
                   trigger: ['change'],
               }],
               addMethod : [{
                   required: true,
                   message: '请选择添加设备方式',
                   trigger: ['change'],
               }],
               // 批量添加验证规则
               batchProductsId : [{
                   required: true,
                   message: '请选择所属产品',
                   trigger: ['change'],
               }],
               batchFile : [{
                   required: true,
                   message: '请上传设备表文件',
                   trigger: ['change'],
               }],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({
    productScope: '',
    productType: '',
    status: '',
    eqName: '',
})

// 重置
const onReset = () => {
  searchInfo.value = {
    productScope: '',
    productType: '',
    status: '',
    eqName: '',
  }
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

// 查询 - 从后端API获取设备列表数据
// 功能：支持分页查询、搜索过滤，更新表格数据和分页信息
const getTableData = async() => {
  const table = await getWlEquipmentList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取产品名称 - 根据产品ID查找对应的产品名称
// 功能：在表格中显示产品名称而不是产品ID，提升用户体验
const getProductName = (productId) => {
  const product = productOptions.value.find(p => p.ID === productId)
  return product ? product.prName : '-'
}

<<<<<<< HEAD
// 获取驱动名称 - 根据驱动ID查找对应的驱动名称
// 功能：在表格中显示驱动ID和名称，提升用户体验
const getDriverName = (driverId) => {
  const driver = driverOptions.value.find(d => d.ID === driverId)
  return driver ? `${driver.driverId || driver.ID} - ${driver.driverName}` : '-'
}

// 批量驱动绑定 - 为选中的设备批量绑定驱动
// 功能：检查选中状态，提示用户选择设备，调用批量绑定API
const batchBindDriver = () => {
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要绑定的设备'
    })
    return
  }
  ElMessage({
    type: 'success',
    message: '批量驱动绑定功能待实现'
  })
}

// 批量驱动解绑 - 为选中的设备批量解绑驱动
// 功能：检查选中状态，提示用户选择设备，调用批量解绑API
const batchUnbindDriver = () => {
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要解绑的设备'
    })
    return
  }
  ElMessage({
    type: 'success',
    message: '批量驱动解绑功能待实现'
  })
}

=======
>>>>>>> kai
// 获取产品选项 - 从后端API获取所有产品列表
// 功能：为设备创建/编辑表单中的产品下拉选择框提供选项数据
const getProductOptions = async () => {
  try {
    const res = await getWlProductsList({ page: 1, pageSize: 1000 }) // 获取所有产品
    if (res.code === 0) {
      productOptions.value = res.data.list || []
    }
  } catch (error) {
    console.error('获取产品列表失败:', error)
    ElMessage.error('获取产品列表失败')
  }
}

// 获取驱动选项 - 从后端API获取所有驱动列表
// 功能：为设备创建/编辑表单中的驱动下拉选择框提供选项数据
const getDriverOptions = async () => {
  try {
    const res = await getWlDriversList({ page: 1, pageSize: 1000 }) // 获取所有驱动
    if (res.code === 0) {
      // 过滤掉异常的驱动ID（大于1000的ID）
      const validDrivers = res.data.list.filter(driver => driver.ID <= 1000)
      driverOptions.value = validDrivers
      console.log('获取到的驱动列表:', driverOptions.value)
      
      // 检查驱动ID是否合理
      driverOptions.value.forEach(driver => {
        if (driver.ID > 1000) {
          console.warn('发现异常的驱动ID:', driver.ID, '驱动名称:', driver.driverName)
        }
      })
      
      if (validDrivers.length < res.data.list.length) {
        console.warn(`过滤掉了 ${res.data.list.length - validDrivers.length} 个异常的驱动ID`)
      }
    }
  } catch (error) {
    console.error('获取驱动列表失败:', error)
    ElMessage.error('获取驱动列表失败')
  }
}

// 获取需要的字典 - 初始化页面所需的选项数据
// 功能：在页面加载时获取产品列表等选项数据
const setOptions = async () =>{
    // 获取所有产品列表
    await getProductOptions()
<<<<<<< HEAD
    // 获取所有驱动列表
    await getDriverOptions()
=======
>>>>>>> kai
}

// 初始化产品选项
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
            deleteWlEquipmentFunc(row)
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
      const res = await deleteWlEquipmentByIds({ IDs })
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
const updateWlEquipmentFunc = async(row) => {
    const res = await findWlEquipment({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
        // 刷新产品列表和驱动列表
        getProductOptions()
        getDriverOptions()
    }
}

// 删除行
const deleteWlEquipmentFunc = async (row) => {
    const res = await deleteWlEquipment({ ID: row.ID })
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
    // 自动生成UUID
    generateUUID()
    // 刷新产品列表和驱动列表
    getProductOptions()
    getDriverOptions()
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        eqName: '',
        eqLogotype: '',
        productsId: undefined,
        eqCoordinate: '',
        eqAddress: '',
        eqInfo: '',
        status: 'offline',
        onlineTime: null,
        addMethod: 'single',
        // 批量添加字段
        batchProductsId: undefined,
        batchFile: null,
        }
    // 清空文件列表
    fileList.value = []
    if (uploadRef.value) {
        uploadRef.value.clearFiles()
    }
}

// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              
             // 确保driveId是数字类型
             const submitData = { ...formData.value }
             
             // 调试信息
             console.log('原始formData:', formData.value)
             console.log('driveId类型:', typeof formData.value.driveId)
             console.log('driveId值:', formData.value.driveId)
             
             if (submitData.driveId && typeof submitData.driveId === 'string') {
               submitData.driveId = parseInt(submitData.driveId)
             }
             if (submitData.productsId && typeof submitData.productsId === 'string') {
               submitData.productsId = parseInt(submitData.productsId)
             }
             
             // 验证driveId是否在合理范围内
             if (submitData.driveId && submitData.driveId > 1000) {
               ElMessage.error('选择的驱动ID异常，请重新选择驱动')
               btnLoading.value = false
               return
             }
             
             console.log('转换后的submitData:', submitData)
              
             if (formData.value.addMethod === 'single') {
               // 单个设备添加
               let res
               switch (type.value) {
                 case 'create':
                   res = await createWlEquipment(submitData)
                   break
                 case 'update':
                   res = await updateWlEquipment(submitData)
                   break
                 default:
                   res = await createWlEquipment(submitData)
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
             } else {
               // 批量设备添加
               if (!formData.value.batchFile) {
                 ElMessage.error('请上传设备表文件')
                 btnLoading.value = false
                 return
               }
               
               // 这里应该调用批量上传API
               ElMessage({
                 type: 'success',
                 message: '批量添加功能待实现，文件已上传'
               })
               btnLoading.value = false
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
  const res = await findWlEquipment({ ID: row.ID })
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

// 获取坐标名称
const getCoordinateName = (code) => {
  const coordinateMap = {
    'beijing_chaoyang': '北京市朝阳区',
    'shanghai_pudong': '上海市浦东新区',
    'guangzhou_tianhe': '广州市天河区',
    'shenzhen_nanshan': '深圳市南山区',
  }
  return coordinateMap[code] || '-'
}

// 打开地图选点对话框
const openMapDialog = () => {
  console.log('打开地图对话框')
  mapDialogVisible.value = true
  // 在下一个tick中初始化地图
  nextTick(() => {
    console.log('开始初始化地图')
    initMap()
  })
}

// 关闭地图选点对话框
const closeMapDialog = () => {
  mapDialogVisible.value = false
  // 清理地图实例
  if (mapInstance.value && typeof mapInstance.value.destroy === 'function') {
    try {
      mapInstance.value.destroy()
    } catch (error) {
      console.warn('地图销毁失败:', error)
    }
    mapInstance.value = null
  }
  mapMarker.value = null
}

// 加载百度地图API
const loadBaiduMapAPI = () => {
  return new Promise((resolve, reject) => {
    if (window.BMap) {
      baiduMapLoaded.value = true
      resolve()
      return
    }

    const script = document.createElement('script')
    script.src = 'https://api.map.baidu.com/api?v=3.0&ak=UQK3TvBBpH45jQAfmnnC5BBWCLXniODu&callback=initBaiduMap'
    script.onerror = reject
    script.onload = () => {
      window.initBaiduMap = () => {
        baiduMapLoaded.value = true
        resolve()
      }
    }
    document.head.appendChild(script)
  })
}

// 初始化地图
const initMap = async () => {
  console.log('初始化地图...')
  
  try {
    // 加载百度地图API
    await loadBaiduMapAPI()
    
    // 等待DOM更新完成
    nextTick(() => {
      const mapContainer = document.getElementById('map-container')
      if (mapContainer && window.BMap) {
        console.log('地图容器找到，初始化百度地图')
        
        // 清空容器
        mapContainer.innerHTML = ''
        
        try {
          // 创建百度地图实例
          const map = new window.BMap.Map(mapContainer)
          mapInstance.value = map
          
          // 设置地图中心点（北京）
          const point = new window.BMap.Point(116.397428, 39.90923)
          map.centerAndZoom(point, 15)
          
          // 添加地图控件
          map.addControl(new window.BMap.NavigationControl())
          map.addControl(new window.BMap.ScaleControl())
          map.addControl(new window.BMap.OverviewMapControl())
          map.addControl(new window.BMap.MapTypeControl())
          
          // 添加点击事件
          map.addEventListener('click', handleBaiduMapClick)
          
          // 设置默认坐标
          selectedLongitude.value = '116.397428'
          selectedLatitude.value = '39.90923'
          selectedAddress.value = '北京市东城区王府井大街'
          
          // 显示默认标记
          showBaiduMapMarker(point, '北京市东城区王府井大街')
        } catch (mapError) {
          console.error('百度地图初始化失败:', mapError)
          initSimulatedMap()
        }
        
      } else {
        console.error('地图容器未找到或百度地图API未加载')
        // 如果百度地图加载失败，使用模拟地图
        initSimulatedMap()
      }
    })
  } catch (error) {
    console.error('百度地图加载失败:', error)
    // 使用模拟地图作为备选
    initSimulatedMap()
  }
}

// 处理百度地图点击
const handleBaiduMapClick = (event) => {
  try {
    const point = event.point
    const lng = point.lng
    const lat = point.lat
    
    selectedLongitude.value = lng.toString()
    selectedLatitude.value = lat.toString()
    
    // 根据坐标获取地址
    const geoc = new window.BMap.Geocoder()
    geoc.getLocation(point, (result) => {
      if (result) {
        selectedAddress.value = result.address
        showBaiduMapMarker(point, result.address)
      } else {
        selectedAddress.value = `北京市 (${lng}, ${lat})`
        showBaiduMapMarker(point, `北京市 (${lng}, ${lat})`)
      }
    })
  } catch (error) {
    console.error('百度地图点击处理失败:', error)
    // 使用模拟坐标
    selectedLongitude.value = '116.397428'
    selectedLatitude.value = '39.90923'
    selectedAddress.value = '北京市东城区王府井大街'
  }
}

// 显示百度地图标记
const showBaiduMapMarker = (point, address) => {
  try {
    const map = mapInstance.value
    if (!map) return
    
    // 清除之前的标记
    if (mapMarker.value) {
      try {
        map.removeOverlay(mapMarker.value)
      } catch (error) {
        console.warn('清除地图标记失败:', error)
      }
    }
    
    // 创建新标记
    const marker = new window.BMap.Marker(point)
    map.addOverlay(marker)
    mapMarker.value = marker
    
    // 添加信息窗口
    const infoWindow = new window.BMap.InfoWindow(address, {
      width: 200,
      height: 100,
      title: '选择的位置'
    })
    marker.addEventListener('click', () => {
      map.openInfoWindow(infoWindow, point)
    })
  } catch (error) {
    console.error('显示百度地图标记失败:', error)
  }
}

// 搜索位置（百度地图）
const searchLocation = async () => {
  if (!mapSearchKeyword.value.trim()) {
    ElMessage.warning('请输入搜索关键词')
    return
  }
  
  const map = mapInstance.value
  if (!map) {
    ElMessage.error('地图未初始化')
    return
  }
  
  try {
    // 使用百度地图搜索
    const local = new window.BMap.LocalSearch(map, {
      onSearchComplete: (results) => {
        try {
          if (local.getStatus() === window.BMAP_STATUS_SUCCESS) {
            const result = results.getPoi(0)
            if (result) {
              const point = result.point
              selectedLongitude.value = point.lng.toString()
              selectedLatitude.value = point.lat.toString()
              selectedAddress.value = result.address
              showBaiduMapMarker(point, result.address)
              
              // 将地图中心移动到搜索结果
              map.panTo(point)
              
              ElMessage.success(`找到: ${result.title}`)
            } else {
              ElMessage.warning('未找到相关位置')
            }
          } else {
            ElMessage.error('搜索失败')
          }
        } catch (error) {
          console.error('搜索回调处理失败:', error)
          ElMessage.error('搜索处理失败')
        }
      }
    })
    
    local.search(mapSearchKeyword.value)
  } catch (error) {
    console.error('百度地图搜索失败:', error)
    ElMessage.error('搜索功能暂时不可用')
  }
}

// 模拟地图初始化（备选方案）
const initSimulatedMap = () => {
  console.log('使用模拟地图')
  const mapContainer = document.getElementById('map-container')
  if (mapContainer) {
    try {
      mapContainer.addEventListener('click', handleMapClick)
      
      // 设置默认坐标
      selectedLongitude.value = '116.397428'
      selectedLatitude.value = '39.90923'
      selectedAddress.value = '北京市东城区王府井大街'
      
      // 显示默认标记
      showMapMarker(50, 50)
    } catch (error) {
      console.error('模拟地图初始化失败:', error)
    }
  }
}

// 处理模拟地图点击
const handleMapClick = (event) => {
  try {
    // 获取点击位置相对于地图容器的坐标
    const mapContainer = document.getElementById('map-container')
    if (!mapContainer) return
    
    const rect = mapContainer.getBoundingClientRect()
    const x = event.clientX - rect.left
    const y = event.clientY - rect.top
    
    // 计算点击位置在地图中的百分比
    const percentX = (x / rect.width * 100).toFixed(2)
    const percentY = (y / rect.height * 100).toFixed(2)
    
    // 模拟坐标计算
    const lng = (116.3 + (x / rect.width) * 0.2).toFixed(6)
    const lat = (39.8 + (y / rect.height) * 0.2).toFixed(6)
    
    selectedLongitude.value = lng
    selectedLatitude.value = lat
    selectedAddress.value = `北京市 (${lng}, ${lat})`
    
    // 显示标记
    showMapMarker(percentX, percentY)
  } catch (error) {
    console.error('模拟地图点击处理失败:', error)
  }
}

// 显示模拟地图标记
const showMapMarker = (x, y) => {
  try {
    console.log(`显示标记: ${x}%, ${y}%`)
    
    // 获取或创建标记元素
    let marker = document.getElementById('map-marker')
    if (!marker) {
      const mapContainer = document.getElementById('map-container')
      if (!mapContainer) return
      
      marker = document.createElement('div')
      marker.id = 'map-marker'
      marker.className = 'map-marker'
      mapContainer.appendChild(marker)
    }
    
    // 设置标记位置
    marker.style.left = x + '%'
    marker.style.top = y + '%'
    marker.style.display = 'block'
  } catch (error) {
    console.error('显示模拟地图标记失败:', error)
  }
}

// 确认地图选择
const confirmMapSelection = () => {
  if (!selectedLongitude.value || !selectedLatitude.value) {
    ElMessage.warning('请选择坐标位置')
    return
  }
  
  // 将选择的坐标设置到表单
  formData.value.eqCoordinate = `${selectedLongitude.value},${selectedLatitude.value}`
  
  // 关闭对话框
  closeMapDialog()
  
  ElMessage.success('坐标选择成功')
}

// 处理路由参数 - 从产品管理页面跳转过来时自动设置产品ID
onMounted(() => {
  // 检查是否有路由参数
  if (route.query.productId) {
    console.log('从产品管理页面跳转过来，产品ID:', route.query.productId)
    console.log('产品名称:', route.query.productName)
    console.log('产品编号:', route.query.productNumber)
    
    // 自动设置产品ID到搜索条件
    searchInfo.value.productsId = route.query.productId
    
    // 重新加载数据
    getTableData()
  }
})
</script>

<style>
.map-container {
  padding: 20px;
}

.map-wrapper {
  position: relative;
  cursor: crosshair;
  overflow: hidden;
}

#map-container {
  position: relative;
  width: 100%;
  height: 100%;
}

.map-marker {
  position: absolute;
  width: 20px;
  height: 20px;
  background: #409eff;
  border: 2px solid white;
  border-radius: 50%;
  transform: translate(-50%, -50%);
  cursor: pointer;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.3);
  z-index: 10;
}

.map-marker::after {
  content: '';
  position: absolute;
  top: 20px;
  left: 50%;
  transform: translateX(-50%);
  width: 0;
  height: 0;
  border-left: 6px solid transparent;
  border-right: 6px solid transparent;
  border-top: 8px solid #409eff;
}
</style>
