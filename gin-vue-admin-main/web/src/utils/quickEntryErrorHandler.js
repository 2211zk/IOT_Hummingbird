/**
 * 快捷入口错误处理工具
 * 提供统一的错误分类和用户友好的错误提示
 */

import { ElMessage, ElNotification } from 'element-plus'

// 错误类型枚举
export const ERROR_TYPES = {
  CONFIG_NOT_FOUND: 'CONFIG_NOT_FOUND',
  ROUTE_NOT_FOUND: 'ROUTE_NOT_FOUND',
  PERMISSION_DENIED: 'PERMISSION_DENIED',
  NAVIGATION_FAILED: 'NAVIGATION_FAILED',
  NETWORK_ERROR: 'NETWORK_ERROR',
  UNKNOWN_ERROR: 'UNKNOWN_ERROR'
}

// 错误消息映射
const ERROR_MESSAGES = {
  [ERROR_TYPES.CONFIG_NOT_FOUND]: {
    title: '功能暂未开放',
    message: '该功能正在开发中，敬请期待',
    type: 'warning'
  },
  [ERROR_TYPES.ROUTE_NOT_FOUND]: {
    title: '页面不存在',
    message: '目标页面不存在或暂未配置，请联系管理员',
    type: 'error'
  },
  [ERROR_TYPES.PERMISSION_DENIED]: {
    title: '权限不足',
    message: '您没有访问该功能的权限，请联系管理员',
    type: 'warning'
  },
  [ERROR_TYPES.NAVIGATION_FAILED]: {
    title: '跳转失败',
    message: '页面跳转失败，请稍后重试',
    type: 'error'
  },
  [ERROR_TYPES.NETWORK_ERROR]: {
    title: '网络错误',
    message: '网络连接异常，请检查网络后重试',
    type: 'error'
  },
  [ERROR_TYPES.UNKNOWN_ERROR]: {
    title: '未知错误',
    message: '发生未知错误，请稍后重试',
    type: 'error'
  }
}

/**
 * 根据错误信息分类错误类型
 * @param {Error} error - 错误对象
 * @returns {string} 错误类型
 */
export function classifyError(error) {
  const message = error.message.toLowerCase()
  
  if (message.includes('未知的快捷入口类型') || message.includes('配置不存在')) {
    return ERROR_TYPES.CONFIG_NOT_FOUND
  }
  
  if (message.includes('路由不存在') || message.includes('页面不存在')) {
    return ERROR_TYPES.ROUTE_NOT_FOUND
  }
  
  if (message.includes('权限') || message.includes('permission')) {
    return ERROR_TYPES.PERMISSION_DENIED
  }
  
  if (message.includes('跳转') || message.includes('navigation')) {
    return ERROR_TYPES.NAVIGATION_FAILED
  }
  
  if (message.includes('网络') || message.includes('network')) {
    return ERROR_TYPES.NETWORK_ERROR
  }
  
  return ERROR_TYPES.UNKNOWN_ERROR
}

/**
 * 显示错误提示
 * @param {string} errorType - 错误类型
 * @param {Object} options - 选项
 * @param {boolean} options.useNotification - 是否使用通知而不是消息
 * @param {string} options.customMessage - 自定义错误消息
 */
export function showErrorMessage(errorType, options = {}) {
  const { useNotification = false, customMessage } = options
  const errorConfig = ERROR_MESSAGES[errorType] || ERROR_MESSAGES[ERROR_TYPES.UNKNOWN_ERROR]
  
  const message = customMessage || errorConfig.message
  
  if (useNotification) {
    ElNotification({
      title: errorConfig.title,
      message: message,
      type: errorConfig.type,
      duration: 4000,
      position: 'top-right'
    })
  } else {
    ElMessage({
      message: message,
      type: errorConfig.type,
      duration: 3000,
      showClose: true
    })
  }
}

/**
 * 统一的快捷入口错误处理函数
 * @param {Error} error - 错误对象
 * @param {string} entryType - 快捷入口类型
 * @param {Object} options - 处理选项
 */
export function handleQuickEntryError(error, entryType, options = {}) {
  console.error(`快捷入口错误 [${entryType}]:`, error)
  
  const errorType = classifyError(error)
  const { getQuickEntryDisplayInfo } = require('@/config/quickEntryConfig')
  const displayInfo = getQuickEntryDisplayInfo(entryType)
  const entryLabel = displayInfo?.label || '未知功能'
  
  // 记录错误到监控系统（如果有的话）
  if (typeof window !== 'undefined' && window.errorTracker) {
    window.errorTracker.captureException(error, {
      tags: {
        component: 'quick-entry',
        entryType: entryType,
        errorType: errorType
      },
      extra: {
        entryLabel: entryLabel,
        timestamp: new Date().toISOString()
      }
    })
  }
  
  // 根据错误类型显示相应的用户提示
  let customMessage = null
  
  switch (errorType) {
    case ERROR_TYPES.CONFIG_NOT_FOUND:
      customMessage = `${entryLabel}功能暂未开放，敬请期待`
      break
    case ERROR_TYPES.ROUTE_NOT_FOUND:
      customMessage = `${entryLabel}页面暂未配置，请联系管理员`
      break
    case ERROR_TYPES.PERMISSION_DENIED:
      customMessage = `您没有访问${entryLabel}的权限`
      break
    case ERROR_TYPES.NAVIGATION_FAILED:
      customMessage = `跳转到${entryLabel}失败，请稍后重试`
      break
  }
  
  showErrorMessage(errorType, {
    ...options,
    customMessage
  })
}

/**
 * 检查用户权限
 * @param {string} routeName - 路由名称
 * @param {Object} userStore - 用户状态存储
 * @returns {boolean} 是否有权限
 */
export function checkUserPermission(routeName, userStore) {
  // 这里可以实现具体的权限检查逻辑
  // 例如检查用户角色、权限列表等
  
  if (!userStore || !userStore.userInfo) {
    return false
  }
  
  // 示例：检查用户是否有特定权限
  const userPermissions = userStore.userInfo.permissions || []
  const routePermissionMap = {
    'WlProducts': 'product:view',
    'WlEquipment': 'device:view',
    'WlAlarm': 'alarm:view',
    'WlEngineRules': 'rule:view',
    'State': 'system:view'
  }
  
  const requiredPermission = routePermissionMap[routeName]
  if (requiredPermission) {
    return userPermissions.includes(requiredPermission)
  }
  
  // 默认允许访问
  return true
}

/**
 * 创建错误处理中间件
 * @param {Object} options - 配置选项
 * @returns {Function} 错误处理中间件
 */
export function createErrorHandler(options = {}) {
  const { enableLogging = true, enableTracking = true } = options
  
  return (error, entryType, context = {}) => {
    if (enableLogging) {
      console.group(`🚨 快捷入口错误 [${entryType}]`)
      console.error('错误对象:', error)
      console.log('上下文:', context)
      console.groupEnd()
    }
    
    handleQuickEntryError(error, entryType, {
      useNotification: context.useNotification,
      enableTracking
    })
  }
}