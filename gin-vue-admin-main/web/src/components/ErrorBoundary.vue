<template>
  <div class="error-boundary">
    <slot v-if="!hasError" />
    
    <!-- 错误显示 -->
    <div v-else class="error-display">
      <div class="error-content">
        <div class="error-icon">
          <el-icon size="48" color="#f56c6c">
            <Warning />
          </el-icon>
        </div>
        
        <div class="error-info">
          <h3 class="error-title">{{ errorTitle }}</h3>
          <p class="error-message">{{ errorMessage }}</p>
          
          <div class="error-actions">
            <el-button type="primary" @click="handleRetry">
              重试
            </el-button>
            <el-button @click="handleReset">
              重置
            </el-button>
            <el-button 
              v-if="showDetails" 
              type="info" 
              @click="toggleDetails"
            >
              {{ showErrorDetails ? '隐藏详情' : '显示详情' }}
            </el-button>
          </div>
          
          <!-- 错误详情 -->
          <div v-if="showErrorDetails" class="error-details">
            <el-collapse>
              <el-collapse-item title="错误详情" name="details">
                <pre class="error-stack">{{ errorDetails }}</pre>
              </el-collapse-item>
            </el-collapse>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onErrorCaptured, watch } from 'vue'
import { Warning } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'

const props = defineProps({
  // 是否显示详情按钮
  showDetails: {
    type: Boolean,
    default: false
  },
  // 自定义错误标题
  errorTitle: {
    type: String,
    default: '出现了一些问题'
  },
  // 自定义错误消息
  customErrorMessage: {
    type: String,
    default: ''
  },
  // 重试回调
  onRetry: {
    type: Function,
    default: null
  },
  // 重置回调
  onReset: {
    type: Function,
    default: null
  }
})

const emit = defineEmits(['error', 'retry', 'reset'])

// 响应式数据
const hasError = ref(false)
const errorMessage = ref('')
const errorDetails = ref('')
const showErrorDetails = ref(false)

// 捕获错误
onErrorCaptured((error, instance, info) => {
  console.error('Error captured by ErrorBoundary:', error)
  
  hasError.value = true
  errorMessage.value = props.customErrorMessage || error.message || '发生未知错误'
  errorDetails.value = `${error.stack}\n\nComponent Info: ${info}`
  
  // 发送错误事件
  emit('error', { error, instance, info })
  
  // 阻止错误继续传播
  return false
})

// 处理重试
const handleRetry = () => {
  if (props.onRetry) {
    props.onRetry()
  } else {
    // 默认重试逻辑：重置错误状态
    hasError.value = false
    errorMessage.value = ''
    errorDetails.value = ''
    showErrorDetails.value = false
  }
  
  emit('retry')
}

// 处理重置
const handleReset = () => {
  if (props.onReset) {
    props.onReset()
  } else {
    // 默认重置逻辑
    hasError.value = false
    errorMessage.value = ''
    errorDetails.value = ''
    showErrorDetails.value = false
    
    // 刷新页面
    window.location.reload()
  }
  
  emit('reset')
}

// 切换详情显示
const toggleDetails = () => {
  showErrorDetails.value = !showErrorDetails.value
}

// 监听外部错误状态变化
watch(() => props.customErrorMessage, (newMessage) => {
  if (newMessage) {
    hasError.value = true
    errorMessage.value = newMessage
  }
})

// 暴露方法给父组件
defineExpose({
  setError: (error, message) => {
    hasError.value = true
    errorMessage.value = message || error.message || '发生错误'
    errorDetails.value = error.stack || error.toString()
  },
  clearError: () => {
    hasError.value = false
    errorMessage.value = ''
    errorDetails.value = ''
    showErrorDetails.value = false
  }
})
</script>

<style scoped>
.error-boundary {
  width: 100%;
  height: 100%;
}

.error-display {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 400px;
  padding: 40px 20px;
  background: #fafafa;
  border-radius: 8px;
}

.error-content {
  text-align: center;
  max-width: 500px;
}

.error-icon {
  margin-bottom: 20px;
}

.error-info {
  color: #606266;
}

.error-title {
  font-size: 20px;
  font-weight: 600;
  margin: 0 0 12px 0;
  color: #303133;
}

.error-message {
  font-size: 14px;
  line-height: 1.6;
  margin: 0 0 24px 0;
  color: #606266;
}

.error-actions {
  display: flex;
  gap: 12px;
  justify-content: center;
  margin-bottom: 20px;
}

.error-details {
  text-align: left;
  margin-top: 20px;
}

.error-stack {
  background: #f5f5f5;
  padding: 12px;
  border-radius: 4px;
  font-size: 12px;
  line-height: 1.4;
  color: #666;
  white-space: pre-wrap;
  word-break: break-all;
  max-height: 200px;
  overflow-y: auto;
}

:deep(.el-collapse-item__header) {
  font-size: 14px;
  font-weight: 500;
}

:deep(.el-collapse-item__content) {
  padding: 12px 0;
}
</style>