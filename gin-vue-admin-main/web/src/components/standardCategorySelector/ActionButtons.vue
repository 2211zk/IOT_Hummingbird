<template>
  <div class="action-buttons">
    <el-button @click="handleCancel">
      取消
    </el-button>
    <el-button
      type="primary"
      :disabled="confirmDisabled"
      @click="handleConfirm"
    >
      确定 {{ selectedCount > 0 ? `(${selectedCount})` : '' }}
    </el-button>
  </div>
</template>

<script setup>
import { computed } from 'vue'

defineOptions({
  name: 'ActionButtons'
})

const props = defineProps({
  selectedCount: {
    type: Number,
    default: 0
  },
  disabled: {
    type: Boolean,
    default: false
  },
  loading: {
    type: Boolean,
    default: false
  }
})

const emits = defineEmits(['confirm', 'cancel'])

// 确定按钮是否禁用
const confirmDisabled = computed(() => {
  return props.disabled || props.loading || props.selectedCount === 0
})

// 处理确定
const handleConfirm = () => {
  if (!confirmDisabled.value) {
    emits('confirm')
  }
}

// 处理取消
const handleCancel = () => {
  emits('cancel')
}
</script>

<style lang="scss" scoped>
.action-buttons {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  padding-top: 16px;
  border-top: 1px solid #ebeef5;
}
</style>