import { ElMessage, ElNotification } from 'element-plus'

/**
 * 统一错误处理函数
 * @param {Error} error - 错误对象
 * @param {string} defaultMessage - 默认错误消息
 * @param {string} type - 错误类型 ('message' | 'notification')
 */
export const handleError = (error, defaultMessage = '操作失败', type = 'message') => {
  console.error('StandardCategorySelector Error:', error)
  
  let message = defaultMessage
  let title = '错误'
  
  // 解析错误信息
  if (error?.response?.data?.msg) {
    message = error.response.data.msg
  } else if (error?.message) {
    message = error.message
  }
  
  // 根据错误类型显示不同的提示
  if (error?.response?.status) {
    switch (error.response.status) {
      case 401:
        message = '登录已过期，请重新登录'
        title = '认证失败'
        break
      case 403:
        message = '没有权限执行此操作'
        title = '权限不足'
        break
      case 404:
        message = '请求的资源不存在'
        title = '资源未找到'
        break
      case 500:
        message = '服务器内部错误，请稍后重试'
        title = '服务器错误'
        break
      case 502:
      case 503:
      case 504:
        message = '服务暂时不可用，请稍后重试'
        title = '服务不可用'
        break
    }
  }
  
  // 显示错误提示
  if (type === 'notification') {
    ElNotification.error({
      title,
      message,
      duration: 5000
    })
  } else {
    ElMessage.error({
      message,
      duration: 3000
    })
  }
}

/**
 * 网络错误处理
 * @param {Error} error - 网络错误
 */
export const handleNetworkError = (error) => {
  if (!navigator.onLine) {
    ElMessage.error('网络连接已断开，请检查网络设置')
    return
  }
  
  if (error.code === 'ECONNABORTED') {
    ElMessage.error('请求超时，请稍后重试')
    return
  }
  
  handleError(error, '网络请求失败，请检查网络连接')
}

/**
 * API调用包装器，统一处理错误
 * @param {Function} apiCall - API调用函数
 * @param {string} errorMessage - 错误消息
 * @param {Object} options - 选项
 */
export const withErrorHandling = async (apiCall, errorMessage = '操作失败', options = {}) => {
  const { 
    showLoading = false, 
    loadingText = '加载中...',
    successMessage = '',
    notificationType = 'message'
  } = options
  
  let loading = null
  
  try {
    if (showLoading) {
      loading = ElMessage({
        message: loadingText,
        type: 'info',
        duration: 0
      })
    }
    
    const result = await apiCall()
    
    if (loading) {
      loading.close()
    }
    
    if (successMessage) {
      if (notificationType === 'notification') {
        ElNotification.success({
          title: '成功',
          message: successMessage
        })
      } else {
        ElMessage.success(successMessage)
      }
    }
    
    return result
  } catch (error) {
    if (loading) {
      loading.close()
    }
    
    if (error.name === 'NetworkError' || !error.response) {
      handleNetworkError(error)
    } else {
      handleError(error, errorMessage, notificationType)
    }
    
    throw error
  }
}

/**
 * 重试机制
 * @param {Function} fn - 要重试的函数
 * @param {number} maxRetries - 最大重试次数
 * @param {number} delay - 重试延迟（毫秒）
 */
export const withRetry = async (fn, maxRetries = 3, delay = 1000) => {
  let lastError
  
  for (let i = 0; i <= maxRetries; i++) {
    try {
      return await fn()
    } catch (error) {
      lastError = error
      
      if (i === maxRetries) {
        break
      }
      
      // 如果是网络错误或服务器错误，进行重试
      if (
        !error.response || 
        error.response.status >= 500 ||
        error.code === 'ECONNABORTED'
      ) {
        await new Promise(resolve => setTimeout(resolve, delay * (i + 1)))
        continue
      }
      
      // 其他错误不重试
      break
    }
  }
  
  throw lastError
}

/**
 * 防抖错误处理
 */
let errorTimer = null
export const debouncedError = (error, message, delay = 1000) => {
  if (errorTimer) {
    clearTimeout(errorTimer)
  }
  
  errorTimer = setTimeout(() => {
    handleError(error, message)
    errorTimer = null
  }, delay)
}