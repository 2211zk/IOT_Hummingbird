<template>
  <div class="category-table">
    <el-table
      v-loading="loading"
      :data="data"
      height="400"
      stripe
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="55" />
      <el-table-column prop="code" label="品类编码" width="120" />
      <el-table-column prop="name" label="品类名称" min-width="150" />
      <el-table-column prop="category" label="所属类别" width="120" />
      <el-table-column prop="description" label="描述" min-width="200" show-overflow-tooltip />
      <el-table-column prop="status" label="状态" width="80">
        <template #default="{ row }">
          <el-tag :type="row.status === 1 ? 'success' : 'danger'">
            {{ row.status === 1 ? '启用' : '禁用' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="100" fixed="right">
        <template #default="{ row }">
          <el-button
            type="primary"
            size="small"
            :disabled="row.status === 0"
            @click="handleSelect(row)"
          >
            选择
          </el-button>
        </template>
      </el-table-column>
      
      <template #empty>
        <el-empty description="暂无数据" />
      </template>
    </el-table>

    <!-- 分页 -->
    <div class="pagination-container">
      <el-pagination
        v-model:current-page="currentPage"
        v-model:page-size="pageSize"
        :total="total"
        :page-sizes="[10, 20, 50, 100]"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { ElMessage } from 'element-plus'

defineOptions({
  name: 'CategoryTable'
})

const props = defineProps({
  data: {
    type: Array,
    default: () => []
  },
  loading: {
    type: Boolean,
    default: false
  },
  total: {
    type: Number,
    default: 0
  },
  currentPage: {
    type: Number,
    default: 1
  },
  pageSize: {
    type: Number,
    default: 10
  }
})

const emits = defineEmits(['select', 'selection-change', 'page-change'])

// 处理单个选择
const handleSelect = (row) => {
  if (row.status === 0) {
    ElMessage.warning('该品类已禁用，无法选择')
    return
  }
  emits('select', row)
}

// 处理多选
const handleSelectionChange = (selection) => {
  // 过滤掉禁用的品类
  const validSelection = selection.filter(item => item.status === 1)
  if (validSelection.length !== selection.length) {
    ElMessage.warning('已自动过滤禁用的品类')
  }
  emits('selection-change', validSelection)
}

// 处理页码变化
const handleCurrentChange = (page) => {
  emits('page-change', { page, pageSize: props.pageSize })
}

// 处理页大小变化
const handleSizeChange = (size) => {
  emits('page-change', { page: 1, pageSize: size })
}
</script>

<style lang="scss" scoped>
.category-table {
  .pagination-container {
    display: flex;
    justify-content: center;
    margin-top: 16px;
  }
}
</style>