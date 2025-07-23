<template>
  <el-dialog 
    v-model="visible" 
    title="标准功能定义" 
    width="80%" 
    :before-close="handleClose"
    destroy-on-close
  >
    <!-- 搜索区域 -->
    <div class="mb-4">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="功能名称">
          <el-input v-model="searchInfo.functionName" placeholder="请输入功能名称" />
        </el-form-item>
        <el-form-item label="功能类型">
          <el-select v-model="searchInfo.functionType" placeholder="请选择" clearable>
            <el-option label="属性" value="属性" />
            <el-option label="事件" value="事件" />
            <el-option label="服务" value="服务" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="searchFunctions">查询</el-button>
          <el-button @click="resetSearch">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

          <!-- 功能定义表格 -->
      <el-table :data="functionList" style="width: 100%" v-loading="loading">
        <el-table-column prop="functionType" label="功能类型" width="120">
          <template #default="scope">
            <el-tag type="primary" size="small">{{ scope.row.functionType || scope.row.function_type }}</el-tag>
            <el-tag type="info" size="small" style="margin-left: 4px;">系统</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="functionName" label="功能名称" width="200">
          <template #default="scope">
            {{ scope.row.functionName || scope.row.function_name }}
          </template>
        </el-table-column>
        <el-table-column prop="identifier" label="标识符" width="120" />
        <el-table-column prop="dataType" label="数据类型" width="100">
          <template #default="scope">
            {{ scope.row.dataType || scope.row.data_type }}
          </template>
        </el-table-column>
        <el-table-column label="描述" min-width="200">
          <template #default="scope">
            {{ getFunctionDescription(scope.row) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="scope">
            <el-button type="primary" link @click="selectFunction(scope.row)">选择</el-button>
            <el-button type="primary" link @click="viewDetail(scope.row)">详情</el-button>
          </template>
        </el-table-column>
      </el-table>

    <!-- 分页控制 -->
    <div class="mt-4">
      <el-pagination
        layout="total, sizes, prev, pager, next, jumper"
        :current-page="page"
        :page-size="pageSize"
        :page-sizes="[10, 20, 50, 100]"
        :total="total"
        @current-change="handlePageChange"
        @size-change="handleSizeChange"
        background
      />
    </div>

    <template #footer>
      <span class="dialog-footer">
        <el-button @click="handleClose">关闭</el-button>
        <el-button type="primary" @click="confirmSelection">确定</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup>
import { ref, reactive, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { getWlCaFunctionByCategory, getWlCaFunctionPublic } from '@/api/wl_playform/wlCaFunction'

// 定义props
const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  categoryId: {
    type: [String, Number],
    default: null
  },
  categoryName: {
    type: String,
    default: ''
  }
})

// 定义emits
const emit = defineEmits(['update:modelValue', 'select'])

// 响应式数据
const visible = ref(false)
const loading = ref(false)
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const functionList = ref([])
const selectedFunction = ref(null)

// 搜索条件
const searchInfo = reactive({
  functionName: '',
  functionType: ''
})

// 监听modelValue变化
watch(() => props.modelValue, (newVal) => {
  visible.value = newVal
  if (newVal) {
    loadFunctionList()
  }
})

// 监听visible变化
watch(visible, (newVal) => {
  emit('update:modelValue', newVal)
})

// 加载功能定义列表
const loadFunctionList = async () => {
  loading.value = true
  try {
    // 构建请求参数 - 发送后端支持的参数
    const params = {
      page: page.value,
      pageSize: pageSize.value,
      caId: props.categoryId,
      functionName: searchInfo.functionName,
      functionType: searchInfo.functionType
    }
    
    console.log('发送API请求参数:', params)
    console.log('品类ID:', props.categoryId, '类型:', typeof props.categoryId)
    console.log('品类名称:', props.categoryName)
    console.log('搜索条件:', searchInfo.value)
    
    // 检查品类ID是否有效
    if (!props.categoryId) {
      console.warn('警告: 品类ID为空，将显示所有功能定义')
      ElMessage.warning('品类ID为空，显示所有功能定义')
    }
    
    // 调用公开API（无需权限）
    const response = await getWlCaFunctionPublic(params)
    
    console.log('API响应:', response)
    
    if (response.code === 0) {
      // 处理API返回的数据
      if (Array.isArray(response.data.list)) {
        functionList.value = response.data.list
      } else if (response.data.list && Array.isArray(response.data.list.list)) {
        functionList.value = response.data.list.list
      } else if (Array.isArray(response.data)) {
        functionList.value = response.data
      } else {
        functionList.value = []
      }
      
      total.value = response.data.total || response.data.totalCount || 0
      page.value = response.data.page || response.data.currentPage || 1
      pageSize.value = response.data.pageSize || 10
      
      console.log('API调用成功，获取到数据:', functionList.value.length, '条')
    } else {
      // 如果API调用失败，显示错误信息
      console.log('API调用失败，错误信息:', response.msg)
      ElMessage.error(`API调用失败: ${response.msg || '未知错误'}`)
      functionList.value = []
      total.value = 0
    }
    
  } catch (error) {
    console.error('加载功能定义失败:', error)
    console.error('错误详情:', error.message)
    ElMessage.error('加载功能定义失败')
    functionList.value = []
    total.value = 0
  } finally {
    loading.value = false
  }
}



// 搜索功能
const searchFunctions = () => {
  page.value = 1
  loadFunctionList()
}

// 重置搜索
const resetSearch = () => {
  searchInfo.functionName = ''
  searchInfo.functionType = ''
  page.value = 1
  loadFunctionList()
}

// 分页处理
const handlePageChange = (val) => {
  page.value = val
  loadFunctionList()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  page.value = 1
  loadFunctionList()
}

// 获取功能描述
const getFunctionDescription = (row) => {
  const functionName = row.functionName || row.function_name
  const functionType = row.functionType || row.function_type
  const dataType = row.dataType || row.data_type
  
  if (functionType === '属性') {
    return `${functionName}的${dataType}类型数据`
  } else if (functionType === '事件') {
    return `${functionName}事件，数据类型为${dataType}`
  } else if (functionType === '服务') {
    return `${functionName}服务，支持${dataType}类型参数`
  }
  
  return `${functionName}功能定义`
}

// 选择功能
const selectFunction = (row) => {
  selectedFunction.value = row
  const functionName = row.functionName || row.function_name
  ElMessage.success(`已选择功能: ${functionName}`)
}

// 查看详情
const viewDetail = (row) => {
  const functionName = row.functionName || row.function_name
  ElMessage.info(`查看功能详情: ${functionName}`)
  // 这里可以打开详情弹窗或跳转到详情页面
}

// 确认选择
const confirmSelection = () => {
  if (selectedFunction.value) {
    emit('select', selectedFunction.value)
    handleClose()
  } else {
    ElMessage.warning('请先选择一个功能')
  }
}

// 关闭对话框
const handleClose = () => {
  visible.value = false
  selectedFunction.value = null
  searchInfo.functionName = ''
  searchInfo.functionType = ''
  page.value = 1
}
</script>

<style scoped>
.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}
</style> 