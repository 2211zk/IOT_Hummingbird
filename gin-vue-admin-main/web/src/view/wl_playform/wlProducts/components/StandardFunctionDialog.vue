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
        <el-table-column prop="functionType" label="功能类型" min-width="150">
          <template #default="scope">
            <el-tag type="primary" size="small">{{ scope.row.functionType || scope.row.function_type }}</el-tag>
            <el-tag type="info" size="small" style="margin-left: 4px;">系统</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="functionName" label="功能名称" min-width="200">
          <template #default="scope">
            {{ scope.row.functionName || scope.row.function_name }}
          </template>
        </el-table-column>
        <el-table-column prop="identifier" label="标识符" min-width="180" />
        <el-table-column prop="dataType" label="数据类型" min-width="120">
          <template #default="scope">
            {{ scope.row.dataType || scope.row.data_type }}
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





// 确认选择
const confirmSelection = () => {
  // 由于移除了操作列，这里可以关闭对话框
  handleClose()
}

// 关闭对话框
const handleClose = () => {
  visible.value = false
  searchInfo.functionName = ''
  searchInfo.functionType = ''
  page.value = 1
}
</script>

<style scoped>
/* 对话框整体样式 */
:deep(.el-dialog) {
  border-radius: 12px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3);
  background: linear-gradient(135deg, #2c3e50 0%, #34495e 100%);
}

:deep(.el-dialog__header) {
  background: transparent;
  color: #ecf0f1;
  border-radius: 12px 12px 0 0;
  padding: 20px 24px;
  border-bottom: none;
}

:deep(.el-dialog__title) {
  color: #ecf0f1;
  font-weight: 600;
  font-size: 18px;
}

:deep(.el-dialog__body) {
  padding: 24px;
  background: #34495e;
}

:deep(.el-dialog__footer) {
  background: #2c3e50;
  border-radius: 0 0 12px 12px;
  padding: 16px 24px;
}

/* 搜索区域样式 */
.mb-4 {
  background: linear-gradient(135deg, #34495e 0%, #2c3e50 100%);
  padding: 20px;
  border-radius: 8px;
  margin-bottom: 20px;
  border: 1px solid #465c71;
}

/* 表格样式 */
:deep(.el-table) {
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.3);
  background: #34495e;
}

:deep(.el-table__header) {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

:deep(.el-table__header th) {
  background: transparent;
  color: white;
  font-weight: 600;
  border-bottom: none;
}

:deep(.el-table__body tr:hover) {
  background: linear-gradient(135deg, #465c71 0%, #34495e 100%);
}

:deep(.el-table__body td) {
  border-bottom: 1px solid #465c71;
  background: #34495e;
  color: #ecf0f1;
}

:deep(.el-table__body tr) {
  background: #34495e;
  color: #ecf0f1;
}

/* 标签样式 */
:deep(.el-tag--primary) {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border: none;
  color: white;
  font-weight: 500;
}

:deep(.el-tag--info) {
  background: linear-gradient(135deg, #6c757d 0%, #495057 100%);
  border: none;
  color: white;
  font-weight: 500;
}

/* 按钮样式 */
:deep(.el-button--primary) {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border: none;
  border-radius: 6px;
  font-weight: 500;
  transition: all 0.3s ease;
}

:deep(.el-button--primary:hover) {
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.4);
}

:deep(.el-button) {
  border-radius: 6px;
  font-weight: 500;
  transition: all 0.3s ease;
}

:deep(.el-button:hover) {
  transform: translateY(-1px);
}

/* 输入框样式 */
:deep(.el-input__wrapper) {
  border-radius: 6px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
  transition: all 0.3s ease;
  background: #2c3e50;
  border: 1px solid #465c71;
}

:deep(.el-input__wrapper:hover) {
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.4);
  border-color: #667eea;
}

:deep(.el-input__wrapper.is-focus) {
  box-shadow: 0 0 0 2px rgba(102, 126, 234, 0.3);
  border-color: #667eea;
}

:deep(.el-input__inner) {
  color: #ecf0f1;
  background: transparent;
}

:deep(.el-input__inner::placeholder) {
  color: #95a5a6;
}

/* 分页样式 */
:deep(.el-pagination) {
  margin-top: 20px;
  text-align: center;
  color: #ecf0f1;
}

:deep(.el-pagination .el-pager li) {
  border-radius: 6px;
  transition: all 0.3s ease;
  background: #2c3e50;
  color: #ecf0f1;
  border: 1px solid #465c71;
}

:deep(.el-pagination .el-pager li:hover) {
  background: #465c71;
  color: #ecf0f1;
}

:deep(.el-pagination .el-pager li.is-active) {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border-color: #667eea;
}

:deep(.el-pagination .btn-prev),
:deep(.el-pagination .btn-next) {
  background: #2c3e50;
  color: #ecf0f1;
  border: 1px solid #465c71;
}

:deep(.el-pagination .btn-prev:hover),
:deep(.el-pagination .btn-next:hover) {
  background: #465c71;
  color: #ecf0f1;
}

/* 对话框底部样式 */
.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

/* 加载状态样式 */
:deep(.el-loading-mask) {
  background: rgba(255, 255, 255, 0.9);
  border-radius: 8px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  :deep(.el-dialog) {
    width: 95% !important;
    margin: 5vh auto;
  }
  
  .mb-4 {
    padding: 16px;
  }
  
  :deep(.el-dialog__body) {
    padding: 16px;
  }
}
</style> 