<template>
  <el-dialog
    v-model="visible"
    title="选择标准品类"
    width="80%"
    :before-close="handleClose"
    :close-on-click-modal="false"
    class="category-modal"
  >
    <div class="category-selector-content">
      <div class="left-panel">
        <SearchBar @search="handleSearch" />
        <CategoryTable
          :data="categoryList"
          :loading="loading"
          :total="total"
          :current-page="searchParams.page"
          :page-size="searchParams.pageSize"
          @select="handleSelect"
          @selection-change="handleSelectionChange"
          @page-change="handlePageChange"
        />
      </div>
      <div class="right-panel">
        <SelectedList
          :selected="selectedCategories"
          :max-selections="maxSelections"
          @remove="handleRemove"
          @clear-all="handleClearAll"
        />
      </div>
    </div>
    
    <template #footer>
      <ActionButtons
        :selected-count="selectedCategories.length"
        :disabled="loading"
        @confirm="handleConfirm"
        @cancel="handleCancel"
      />
    </template>
  </el-dialog>
</template>

<script setup>
import { ref, reactive, watch, nextTick } from 'vue'
import { ElMessage } from 'element-plus'
import SearchBar from './SearchBar.vue'
import CategoryTable from './CategoryTable.vue'
import SelectedList from './SelectedList.vue'
import ActionButtons from './ActionButtons.vue'
import { getStandardCategoryList } from '@/api/standardCategory'
import { handleError, withErrorHandling, withRetry } from './utils/errorHandler'
import { useLoading, delayedLoading } from './utils/loadingManager'

defineOptions({
  name: 'CategoryModal'
})

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  multiple: {
    type: Boolean,
    default: true
  },
  maxSelections: {
    type: Number,
    default: 0
  },
  initialSelected: {
    type: Array,
    default: () => []
  }
})

const emits = defineEmits(['update:modelValue', 'confirm', 'cancel'])

// 弹框显示状态
const visible = ref(false)

// 使用加载状态管理
const { loading, withLoading } = useLoading('categoryList')

// 延迟加载管理
const delayedLoadingManager = delayedLoading(200)

// 品类列表数据
const categoryList = ref([])
const total = ref(0)

// 搜索参数
const searchParams = reactive({
  page: 1,
  pageSize: 10,
  keyword: '',
  category: '',
  status: null
})

// 已选择的品类
const selectedCategories = ref([])

// 监听弹框显示状态
watch(() => props.modelValue, (newVal) => {
  visible.value = newVal
  if (newVal) {
    // 弹框打开时初始化数据
    initializeData()
  }
})

watch(visible, (newVal) => {
  emits('update:modelValue', newVal)
})

// 初始化数据
const initializeData = () => {
  // 重置搜索参数
  Object.assign(searchParams, {
    page: 1,
    pageSize: 10,
    keyword: '',
    category: '',
    status: null
  })
  
  // 初始化已选择的品类
  selectedCategories.value = [...props.initialSelected]
  
  // 获取品类列表
  fetchCategoryList()
}

// 获取品类列表
const fetchCategoryList = async () => {
  try {
    const response = await withLoading(
      () => withRetry(() => getStandardCategoryList(searchParams), 2, 1000),
      '正在获取品类列表...'
    )
    
    if (response.code === 0) {
      categoryList.value = response.data.list || []
      total.value = response.data.total || 0
    } else {
      throw new Error(response.msg || '获取品类列表失败')
    }
  } catch (error) {
    handleError(error, '获取品类列表失败')
    categoryList.value = []
    total.value = 0
  }
}

// 处理搜索
const handleSearch = (searchData) => {
  Object.assign(searchParams, searchData, { page: 1 })
  fetchCategoryList()
}

// 处理分页变化
const handlePageChange = ({ page, pageSize }) => {
  searchParams.page = page
  searchParams.pageSize = pageSize
  fetchCategoryList()
}

// 处理单个选择
const handleSelect = (category) => {
  if (isAlreadySelected(category)) {
    ElMessage.warning('该品类已经被选择')
    return
  }

  if (!props.multiple) {
    // 单选模式
    selectedCategories.value = [category]
  } else {
    // 多选模式
    if (props.maxSelections && selectedCategories.value.length >= props.maxSelections) {
      ElMessage.warning(`最多只能选择 ${props.maxSelections} 个品类`)
      return
    }
    selectedCategories.value.push(category)
  }
}

// 处理表格多选
const handleSelectionChange = (selection) => {
  if (!props.multiple) return

  // 过滤掉已经选择的品类，避免重复
  const newSelections = selection.filter(item => !isAlreadySelected(item))
  
  if (props.maxSelections) {
    const availableSlots = props.maxSelections - selectedCategories.value.length
    if (newSelections.length > availableSlots) {
      ElMessage.warning(`最多只能选择 ${props.maxSelections} 个品类`)
      return
    }
  }

  selectedCategories.value.push(...newSelections)
}

// 检查是否已经选择
const isAlreadySelected = (category) => {
  return selectedCategories.value.some(item => item.id === category.id)
}

// 移除选择
const handleRemove = (category) => {
  const index = selectedCategories.value.findIndex(item => item.id === category.id)
  if (index > -1) {
    selectedCategories.value.splice(index, 1)
  }
}

// 清空所有选择
const handleClearAll = () => {
  selectedCategories.value = []
}

// 处理确定
const handleConfirm = () => {
  if (selectedCategories.value.length === 0) {
    ElMessage.warning('请至少选择一个品类')
    return
  }

  emits('confirm', [...selectedCategories.value])
  visible.value = false
}

// 处理取消
const handleCancel = () => {
  emits('cancel')
  visible.value = false
}

// 处理关闭
const handleClose = (done) => {
  handleCancel()
  done()
}
</script>

<style lang="scss" scoped>
.category-modal {
  :deep(.el-dialog) {
    max-width: 1200px;
    min-width: 800px;
  }

  :deep(.el-dialog__body) {
    padding: 20px;
  }

  .category-selector-content {
    display: flex;
    height: 500px;
    gap: 20px;

    .left-panel {
      flex: 2;
      display: flex;
      flex-direction: column;
      min-width: 0; // 防止flex子项溢出

      .search-bar {
        flex-shrink: 0;
      }

      .category-table {
        flex: 1;
        overflow: hidden;
      }
    }

    .right-panel {
      flex: 1;
      border-left: 1px solid #ebeef5;
      padding-left: 20px;
      min-width: 300px;

      .selected-list {
        height: 100%;
      }
    }
  }
}

// 响应式设计
@media (max-width: 1024px) {
  .category-modal {
    :deep(.el-dialog) {
      width: 95% !important;
      min-width: auto;
      margin: 5vh auto;
    }

    .category-selector-content {
      height: 60vh;
      flex-direction: column;
      gap: 16px;

      .left-panel {
        flex: none;
        height: 60%;
      }

      .right-panel {
        flex: none;
        height: 40%;
        border-left: none;
        border-top: 1px solid #ebeef5;
        padding-left: 0;
        padding-top: 16px;
        min-width: auto;
      }
    }
  }
}

@media (max-width: 768px) {
  .category-modal {
    :deep(.el-dialog) {
      width: 98% !important;
      margin: 2vh auto;
    }

    .category-selector-content {
      height: 70vh;
    }
  }
}
</style>