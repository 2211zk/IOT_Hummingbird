<template>
  <div class="standard-category-selector">
    <!-- 触发按钮 -->
    <div class="selector-trigger">
      <el-button
        :disabled="disabled"
        @click="openModal"
      >
        <el-icon><Plus /></el-icon>
        {{ placeholder }}
      </el-button>
      
      <!-- 已选择的品类显示 -->
      <div v-if="selectedCategories.length > 0" class="selected-display">
        <el-tag
          v-for="category in displayCategories"
          :key="category.id"
          closable
          :disable-transitions="false"
          @close="handleRemoveCategory(category)"
        >
          {{ category.name }}
        </el-tag>
        
        <el-tag v-if="hasMore" type="info">
          +{{ selectedCategories.length - maxDisplay }}
        </el-tag>
      </div>
    </div>

    <!-- 弹出模态框 -->
    <CategoryModal
      v-model="modalVisible"
      :multiple="multiple"
      :max-selections="maxSelections"
      :initial-selected="selectedCategories"
      @confirm="handleConfirm"
      @cancel="handleCancel"
    />
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { Plus } from '@element-plus/icons-vue'
import CategoryModal from './CategoryModal.vue'

defineOptions({
  name: 'StandardCategorySelector'
})

const props = defineProps({
  modelValue: {
    type: Array,
    default: () => []
  },
  multiple: {
    type: Boolean,
    default: true
  },
  placeholder: {
    type: String,
    default: '选择标准品类'
  },
  disabled: {
    type: Boolean,
    default: false
  },
  maxSelections: {
    type: Number,
    default: 0
  },
  maxDisplay: {
    type: Number,
    default: 3
  }
})

const emits = defineEmits(['update:modelValue', 'change', 'confirm'])

// 弹框显示状态
const modalVisible = ref(false)

// 已选择的品类
const selectedCategories = ref([...props.modelValue])

// 监听外部值变化
watch(() => props.modelValue, (newVal) => {
  selectedCategories.value = [...newVal]
}, { deep: true })

// 显示的品类（限制显示数量）
const displayCategories = computed(() => {
  return selectedCategories.value.slice(0, props.maxDisplay)
})

// 是否有更多品类未显示
const hasMore = computed(() => {
  return selectedCategories.value.length > props.maxDisplay
})

// 打开弹框
const openModal = () => {
  if (props.disabled) return
  modalVisible.value = true
}

// 处理确认选择
const handleConfirm = (categories) => {
  selectedCategories.value = [...categories]
  emits('update:modelValue', selectedCategories.value)
  emits('change', selectedCategories.value)
  emits('confirm', selectedCategories.value)
}

// 处理取消
const handleCancel = () => {
  // 取消时不做任何操作，保持原有选择
}

// 移除单个品类
const handleRemoveCategory = (category) => {
  const index = selectedCategories.value.findIndex(item => item.id === category.id)
  if (index > -1) {
    selectedCategories.value.splice(index, 1)
    emits('update:modelValue', selectedCategories.value)
    emits('change', selectedCategories.value)
  }
}

// 暴露方法给父组件
defineExpose({
  openModal,
  clearSelection: () => {
    selectedCategories.value = []
    emits('update:modelValue', selectedCategories.value)
    emits('change', selectedCategories.value)
  },
  getSelectedCategories: () => selectedCategories.value
})
</script>

<style lang="scss" scoped>
@import './styles.scss';

.standard-category-selector {
  .selector-trigger {
    .selected-display {
      margin-top: 8px;
      display: flex;
      flex-wrap: wrap;
      gap: 6px;

      .el-tag {
        margin: 0;
      }
    }
  }
}
</style>