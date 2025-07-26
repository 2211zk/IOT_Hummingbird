import { ElMessage, ElNotification } from 'element-plus'

// API错误处理工具

/**
 * 错误码映射
 */
const ERROR_CODE_MAP = {
  // 通用错误
  400: '请求参数错误',
  401: '未授权访问',
  403: '权限不足',
  404: '请求的资源不存在',
  405: '请求方法不允许',
  500: '服务器内部错误',
  502: '网关错误',
  503: '服务不可用',
  504: '网关超时',
  
  // 业务错误码
  7001: '部门名称已存在',
  7002: '上级部门不能是自身或子部门',
  7003: '该部门下还有子部门，无法删除',
  7004: '部门不存在',
  7005: '设备不存在',
  7006: '设备已被其他部门关联'
}

/**
 * 处理API错误
 * @param {Error|Object} error - 错误对象
 * @param {Object} options - 处理选项
 * @param {boolean} options.showMessage - 是否显示错误消息
 * @param {boolean} options.showNotification - 是否显示通知
 * @param {string} options.defaultMessage - 默认错误消息
 * @param {Function} options.onError - 自定义错误处理函数
 */
export function handleApiError(error, options = {}) {
  const {
    showMessage = true,
    showNotification = false,
    defaultMessage = '操作失败',
    onError
  } = options
  
  let errorMessage = defaultMessage
  let errorCode = null
  
  // 解析错误信息
  if (error.response) {
    // HTTP错误
    errorCode = error.response.status
    errorMessage = ERROR_CODE_MAP[errorCode] || error.response.data?.msg || errorMessage
  } else if (error.code) {
    // 业务错误
    errorCode = error.code
    errorMessage = ERROR_CODE_MAP[errorCode] || error.msg || errorMessage
  } else if (error.msg) {
    // 直接的错误消息
    errorMessage = error.msg
  } else if (typeof error === 'string') {
    // 字符串错误
    errorMessage = error
  }
  
  // 显示错误消息
  if (showMessage) {
    ElMessage.error(errorMessage)
  }
  
  if (showNotification) {
    ElNotification.error({
      title: '操作失败',
      message: errorMessage,
      duration: 5000
    })
  }
  
  // 自定义错误处理
  if (onError && typeof onError === 'function') {
    onError(error, errorMessage, errorCode)
  }
  
  // 返回格式化的错误信息
  return {
    message: errorMessage,
    code: errorCode,
    originalError: error
  }
}

/**
 * 创建API调用包装器
 * @param {Function} apiFunction - API函数
 * @param {Object} errorOptions - 错误处理选项
 * @returns {Function} 包装后的API函数
 */
export function createApiWrapper(apiFunction, errorOptions = {}) {
  return async (...args) => {
    try {
      const result = await apiFunction(...args)
      return result
    } catch (error) {
      const errorInfo = handleApiError(error, errorOptions)
      throw errorInfo
    }
  }
}

/**
 * 批量处理API调用结果
 * @param {Array} promises - Promise数组
 * @param {Object} options - 选项
 * @param {boolean} options.failFast - 是否快速失败
 * @param {Function} options.onSuccess - 成功回调
 * @param {Function} options.onError - 错误回调
 */
export async function batchApiCall(promises, options = {}) {
  const {
    failFast = false,
    onSuccess,
    onError
  } = options
  
  const results = []
  const errors = []
  
  if (failFast) {
    // 快速失败模式
    try {
      const batchResults = await Promise.all(promises)
      batchResults.forEach((result, index) => {
        results.push({ index, result, success: true })
        if (onSuccess) onSuccess(result, index)
      })
    } catch (error) {
      const errorInfo = handleApiError(error)
      if (onError) onError(errorInfo)
      throw errorInfo
    }
  } else {
    // 容错模式
    const settledResults = await Promise.allSettled(promises)
    settledResults.forEach((result, index) => {
      if (result.status === 'fulfilled') {
        results.push({ index, result: result.value, success: true })
        if (onSuccess) onSuccess(result.value, index)
      } else {
        const errorInfo = handleApiError(result.reason, { showMessage: false })
        errors.push({ index, error: errorInfo, success: false })
        if (onError) onError(errorInfo, index)
      }
    })
  }
  
  return {
    results,
    errors,
    success: errors.length === 0
  }
}

/**
 * 重试API调用
 * @param {Function} apiFunction - API函数
 * @param {Array} args - 参数
 * @param {Object} options - 重试选项
 * @param {number} options.maxRetries - 最大重试次数
 * @param {number} options.delay - 重试延迟（毫秒）
 * @param {Function} options.shouldRetry - 是否应该重试的判断函数
 */
export async function retryApiCall(apiFunction, args = [], options = {}) {
  const {
    maxRetries = 3,
    delay = 1000,
    shouldRetry = (error) => error.code >= 500
  } = options
  
  let lastError
  
  for (let attempt = 0; attempt <= maxRetries; attempt++) {
    try {
      const result = await apiFunction(...args)
      return result
    } catch (error) {
      lastError = error
      
      if (attempt === maxRetries || !shouldRetry(error)) {
        break
      }
      
      // 等待后重试
      await new Promise(resolve => setTimeout(resolve, delay * Math.pow(2, attempt)))
    }
  }
  
  throw handleApiError(lastError)
}

/**
 * 防抖API调用
 * @param {Function} apiFunction - API函数
 * @param {number} delay - 防抖延迟
 * @returns {Function} 防抖后的函数
 */
export function debounceApiCall(apiFunction, delay = 300) {
  let timeoutId
  let lastPromise
  
  return (...args) => {
    return new Promise((resolve, reject) => {
      clearTimeout(timeoutId)
      
      timeoutId = setTimeout(async () => {
        try {
          const result = await apiFunction(...args)
          resolve(result)
        } catch (error) {
          reject(handleApiError(error))
        }
      }, delay)
      
      lastPromise = { resolve, reject }
    })
  }
}

/**
 * 节流API调用
 * @param {Function} apiFunction - API函数
 * @param {number} delay - 节流延迟
 * @returns {Function} 节流后的函数
 */
export function throttleApiCall(apiFunction, delay = 1000) {
  let lastCallTime = 0
  let timeoutId
  
  return (...args) => {
    return new Promise((resolve, reject) => {
      const now = Date.now()
      const timeSinceLastCall = now - lastCallTime
      
      if (timeSinceLastCall >= delay) {
        lastCallTime = now
        apiFunction(...args).then(resolve).catch(error => reject(handleApiError(error)))
      } else {
        clearTimeout(timeoutId)
        timeoutId = setTimeout(() => {
          lastCallTime = Date.now()
          apiFunction(...args).then(resolve).catch(error => reject(handleApiError(error)))
        }, delay - timeSinceLastCall)
      }
    })
  }
}