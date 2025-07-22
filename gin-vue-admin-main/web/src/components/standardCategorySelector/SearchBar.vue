<template>
  <div class="search-bar">
    <el-row :gutter="16">
      <el-col :span="8">
        <el-input
          v-model="searchForm.keyword"
          placeholder="请输入品类名称或编码"
          clearable
          @input="handleSearch"
          @clear="handleSearch"
        >
          <template #prefix>
            <el-icon><Search /></el-icon>
          </template>
        </el-input>
      </el-col>
      <el-col :span="6">
        <el-select
          v-model="searchForm.category"
          placeholder="选择类别"
          clearable
          @change="handleSearch"
        >
          <el-option
            v-for="category in categories"
            :key="category"
            :label="category"
            :value="category"
          />
        </el-select>
      </el-col>
      <el-col :span="6">
        <el-select
          v-model="searchForm.status"
          placeholder="选择状态"
          clearable
          @change="handleSearch"
        >
          <el-option label="启用" :value="1" />
          <el-option label="禁用" :value="0" />
        </el-select>
      </el-col>
      <el-col :span="4">
        <el-button type="primary" @click="handleSearch">
          <el-icon><Search /></el-icon>
          搜索
        </el-button>
        <el-button @click="handleReset">
          <el-icon><Refresh /></el-icon>
          重置
        </el-button>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { Search, Refresh } from '@element-plus/icons-vue'
import { getStandardCategoryCategories } from '@/api/standardCategory'
import { ElMessage } from 'element-plus'
import { debounce } from 'lodash-es'
import { handleError, withRetry } from './utils/errorHandler'
import { useLoading } from './utils/loadingManager'

defineOptions({
  name: 'SearchBar'
})

const emits = defineEmits(['search'])

// 搜索表单
const searchForm = reactive({
  keyword: '',
  category: '',
  status: null
})

// 类别选项
const categories = ref([])

// 使用加载状态管理
const { loading: categoriesLoading } = useLoading('categories')

// 防抖搜索
const debouncedSearch = debounce(() => {
  emits('search', { ...searchForm })
}, 300)

// 处理搜索
const handleSearch = () => {
  debouncedSearch()
}

// 重置搜索
const handleReset = () => {
  searchForm.keyword = ''
  searchForm.category = ''
  searchForm.status = null
  emits('search', { ...searchForm })
}

// 获取类别列表
const fetchCategories = async () => {
  try {
    const response = await withRetry(
      () => getStandardCategoryCategories(),
      2,
      500
    )
    
    if (response.code === 0) {
      categories.value = response.data || []
    } else {
      throw new Error(response.msg || '获取类别列表失败')
    }
  } catch (error) {
    handleError(error, '获取类别列表失败')
    // 设置默认类别选项
    categories.value = []
  }
}

onMounted(() => {
  fetchCategories()
})
</script>

<style lang="scss" scoped>
.search-bar {
  padding: 16px;
  background: #f8f9fa;
  border-radius: 4px;
  margin-bottom: 16px;

  .el-button + .el-button {
    margin-left: 8px;
  }
}
</style>