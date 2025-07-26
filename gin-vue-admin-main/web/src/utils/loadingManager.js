import { ref, reactive } from 'vue'
import { ElMessage, ElNotification, ElLoading } from 'element-plus'

// 全局加载状态管理
class LoadingManager {
  constructor() {
    this.loadingStates = reactive({})
    this.globalLoading = ref(false)
    this.loadingInstances = new Map()
  }

  // 设置加载状态
  setLoading(key, loading = true, options = {}) {
    this.loadingStates[key] = loading
    
    // 更新全局加载状态
    this.updateGlobalLoading()
    
    // 如果需要全屏加载
    if (options.fullscreen && loading) {
      this.showFullscreenLoading(key, options)
    } else if (!loading && this.loadingInstances.has(key)) {
      this.hideFullscreenLoading(key)
    }
  }

  // 获取加载状态
  getLoading(key) {
    return this.loadingStates[key] || false
  }

  // 更新全局加载状态
  updateGlobalLoading() {
    this.globalLoading.value = Object.values(this.loadingStates).some(loading => loading)
  }

  // 显示全屏加载
  showFullscreenLoading(key, options = {}) {
    const loadingInstance = ElLoading.service({
      lock: true,
      text: options.text || '加载中...',
      background: options.background || 'rgba(0, 0, 0, 0.7)',
      spinner: options.spinner,
      customClass: options.customClass
    })
    
    this.loadingInstances.set(key, loadingInstance)
  }

  // 隐藏全屏加载
  hideFullscreenLoading(key) {
    const instance = this.loadingInstances.get(key)
    if (instance) {
      instance.close()
      this.loadingInstances.delete(key)
    }
  }

  // 清除所有加载状态
  clearAll() {
    Object.keys(this.loadingStates).forEach(key => {
      this.loadingStates[key] = false
    })
    
    // 关闭所有全屏加载
    this.loadingInstances.forEach(instance => instance.close())
    this.loadingInstances.clear()
    
    this.updateGlobalLoading()
  }

  // 创建加载装饰器
  createLoadingDecorator(key, options = {}) {
    return (target, propertyKey, descriptor) => {
      const originalMethod = descriptor.value
      
      descriptor.value = async function(...args) {
        try {
          this.setLoading(key, true, options)
          const result = await originalMethod.apply(this, args)
          return result
        } finally {
          this.setLoading(key, false)
        }
      }
      
      return descriptor
    }
  }

  // 包装异步函数
  wrapAsync(asyncFn, key, options = {}) {
    return async (...args) => {
      try {
        this.setLoading(key, true, options)
        const result = await asyncFn(...args)
        return result
      } finally {
        this.setLoading(key, false)
      }
    }
  }
}

// 错误处理管理器
class ErrorManager {
  constructor() {
    this.errorHistory = []
    this.maxHistorySize = 50
  }

  // 处理错误
  handleError(error, options = {}) {
    const errorInfo = this.parseError(error)
    
    // 记录错误历史
    this.recordError(errorInfo)
    
    // 显示错误消息
    this.showErrorMessage(errorInfo, options)
    
    // 上报错误（如果需要）
    if (options.report !== false) {
      this.reportError(errorInfo)
    }
    
    return errorInfo
  }

  // 解析错误信息
  parseError(error) {
    let message = '操作失败'
    let code = null
    let type = 'unknown'
    
    if (error?.response) {
      // HTTP错误
      code = error.response.status
      message = error.response.data?.msg || this.getHttpErrorMessage(code)
      type = 'http'
    } else if (error?.code) {
      // 业务错误
      code = error.code
      message = error.msg || this.getBusinessErrorMessage(code)
      type = 'business'
    } else if (typeof error === 'string') {
      message = error
      type = 'string'
    } else if (error?.message) {
      message = error.message
      type = 'exception'
    }
    
    return {
      message,
      code,
      type,
      timestamp: new Date(),
      originalError: error
    }
  }

  // 获取HTTP错误消息
  getHttpErrorMessage(code) {
    const messages = {
      400: '请求参数错误',
      401: '未授权访问',
      403: '权限不足',
      404: '请求的资源不存在',
      405: '请求方法不允许',
      408: '请求超时',
      500: '服务器内部错误',
      502: '网关错误',
      503: '服务不可用',
      504: '网关超时'
    }
    return messages[code] || `HTTP错误 ${code}`
  }

  // 获取业务错误消息
  getBusinessErrorMessage(code) {
    const messages = {
      7001: '部门名称已存在',
      7002: '上级部门不能是自身或子部门',
      7003: '该部门下还有子部门，无法删除',
      7004: '部门不存在',
      7005: '设备不存在',
      7006: '设备已被其他部门关联'
    }
    return messages[code] || `业务错误 ${code}`
  }

  // 显示错误消息
  showErrorMessage(errorInfo, options = {}) {
    const {
      showType = 'message', // message | notification | console
      duration = 3000,
      showClose = true
    } = options
    
    switch (showType) {
      case 'message':
        ElMessage.error({
          message: errorInfo.message,
          duration,
          showClose
        })
        break
      case 'notification':
        ElNotification.error({
          title: '操作失败',
          message: errorInfo.message,
          duration,
          showClose
        })
        break
      case 'console':
        console.error('Error:', errorInfo)
        break
    }
  }

  // 记录错误历史
  recordError(errorInfo) {
    this.errorHistory.unshift(errorInfo)
    
    // 限制历史记录大小
    if (this.errorHistory.length > this.maxHistorySize) {
      this.errorHistory = this.errorHistory.slice(0, this.maxHistorySize)
    }
  }

  // 上报错误
  reportError(errorInfo) {
    // 这里可以实现错误上报逻辑
    // 例如发送到错误监控服务
    if (process.env.NODE_ENV === 'development') {
      console.group('Error Report')
      console.error('Error Info:', errorInfo)
      console.error('Stack Trace:', errorInfo.originalError)
      console.groupEnd()
    }
  }

  // 获取错误历史
  getErrorHistory() {
    return [...this.errorHistory]
  }

  // 清除错误历史
  clearErrorHistory() {
    this.errorHistory = []
  }
}

// 创建全局实例
export const loadingManager = new LoadingManager()
export const errorManager = new ErrorManager()

// 组合式API
export function useLoading() {
  return {
    setLoading: loadingManager.setLoading.bind(loadingManager),
    getLoading: loadingManager.getLoading.bind(loadingManager),
    globalLoading: loadingManager.globalLoading,
    loadingStates: loadingManager.loadingStates,
    wrapAsync: loadingManager.wrapAsync.bind(loadingManager)
  }
}

export function useError() {
  return {
    handleError: errorManager.handleError.bind(errorManager),
    getErrorHistory: errorManager.getErrorHistory.bind(errorManager),
    clearErrorHistory: errorManager.clearErrorHistory.bind(errorManager)
  }
}

// 便捷函数
export function withLoading(asyncFn, key, options = {}) {
  return loadingManager.wrapAsync(asyncFn, key, options)
}

export function handleApiError(error, options = {}) {
  return errorManager.handleError(error, options)
}