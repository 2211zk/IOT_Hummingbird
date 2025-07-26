import { ElMessage, ElNotification, ElMessageBox } from 'element-plus'
import { h } from 'vue'

// 用户通知管理器
class UserNotificationManager {
  constructor() {
    this.notificationQueue = []
    this.isProcessing = false
    this.maxQueueSize = 10
  }

  // 显示成功消息
  showSuccess(message, options = {}) {
    const config = {
      type: 'success',
      message,
      duration: 3000,
      showClose: true,
      ...options
    }

    if (options.notification) {
      ElNotification.success({
        title: '操作成功',
        ...config
      })
    } else {
      ElMessage.success(config)
    }
  }

  // 显示错误消息
  showError(message, options = {}) {
    const config = {
      type: 'error',
      message,
      duration: 5000,
      showClose: true,
      ...options
    }

    if (options.notification) {
      ElNotification.error({
        title: '操作失败',
        ...config
      })
    } else {
      ElMessage.error(config)
    }
  }

  // 显示警告消息
  showWarning(message, options = {}) {
    const config = {
      type: 'warning',
      message,
      duration: 4000,
      showClose: true,
      ...options
    }

    if (options.notification) {
      ElNotification.warning({
        title: '注意',
        ...config
      })
    } else {
      ElMessage.warning(config)
    }
  }

  // 显示信息消息
  showInfo(message, options = {}) {
    const config = {
      type: 'info',
      message,
      duration: 3000,
      showClose: true,
      ...options
    }

    if (options.notification) {
      ElNotification.info({
        title: '提示',
        ...config
      })
    } else {
      ElMessage.info(config)
    }
  }

  // 显示加载消息
  showLoading(message = '加载中...', options = {}) {
    return ElMessage({
      type: 'info',
      message,
      duration: 0,
      showClose: false,
      iconClass: 'el-icon-loading',
      ...options
    })
  }

  // 显示确认对话框
  async showConfirm(message, title = '确认', options = {}) {
    const config = {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
      ...options
    }

    try {
      await ElMessageBox.confirm(message, title, config)
      return true
    } catch {
      return false
    }
  }

  // 显示输入对话框
  async showPrompt(message, title = '输入', options = {}) {
    const config = {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      inputPattern: options.pattern,
      inputErrorMessage: options.errorMessage,
      inputPlaceholder: options.placeholder,
      inputValue: options.defaultValue,
      ...options
    }

    try {
      const { value } = await ElMessageBox.prompt(message, title, config)
      return value
    } catch {
      return null
    }
  }

  // 显示自定义对话框
  async showCustomDialog(options = {}) {
    const config = {
      title: '提示',
      message: '',
      showCancelButton: true,
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      ...options
    }

    try {
      await ElMessageBox(config)
      return true
    } catch {
      return false
    }
  }

  // 显示进度提示
  showProgress(message, progress = 0, options = {}) {
    const progressMessage = h('div', [
      h('p', message),
      h('el-progress', {
        percentage: progress,
        strokeWidth: 6,
        ...options.progressProps
      })
    ])

    return ElNotification({
      title: '进度',
      message: progressMessage,
      duration: 0,
      showClose: false,
      ...options
    })
  }

  // 显示操作结果
  showOperationResult(success, successMessage, errorMessage, options = {}) {
    if (success) {
      this.showSuccess(successMessage, options)
    } else {
      this.showError(errorMessage, options)
    }
  }

  // 显示网络错误
  showNetworkError(error, options = {}) {
    let message = '网络连接失败，请检查网络设置'
    
    if (error.code === 'NETWORK_ERROR') {
      message = '网络连接中断，请检查网络连接'
    } else if (error.code === 'TIMEOUT') {
      message = '请求超时，请稍后重试'
    } else if (error.response?.status === 0) {
      message = '无法连接到服务器，请检查网络连接'
    }

    this.showError(message, {
      duration: 6000,
      notification: true,
      ...options
    })
  }

  // 显示权限错误
  showPermissionError(action = '执行此操作', options = {}) {
    this.showError(`您没有权限${action}，请联系管理员`, {
      duration: 5000,
      notification: true,
      ...options
    })
  }

  // 显示验证错误
  showValidationError(errors, options = {}) {
    if (Array.isArray(errors)) {
      const errorList = errors.map(error => `• ${error}`).join('\n')
      this.showError(`请修正以下错误：\n${errorList}`, {
        duration: 6000,
        dangerouslyUseHTMLString: true,
        ...options
      })
    } else {
      this.showError(errors, options)
    }
  }

  // 显示批量操作结果
  showBatchResult(results, options = {}) {
    const { success = [], failed = [] } = results
    
    if (failed.length === 0) {
      this.showSuccess(`批量操作完成，成功处理 ${success.length} 项`, options)
    } else if (success.length === 0) {
      this.showError(`批量操作失败，${failed.length} 项处理失败`, options)
    } else {
      this.showWarning(
        `批量操作部分成功：成功 ${success.length} 项，失败 ${failed.length} 项`,
        {
          duration: 6000,
          notification: true,
          ...options
        }
      )
    }
  }

  // 队列化通知（避免通知过多）
  queueNotification(notification) {
    if (this.notificationQueue.length >= this.maxQueueSize) {
      this.notificationQueue.shift() // 移除最旧的通知
    }
    
    this.notificationQueue.push(notification)
    
    if (!this.isProcessing) {
      this.processQueue()
    }
  }

  // 处理通知队列
  async processQueue() {
    this.isProcessing = true
    
    while (this.notificationQueue.length > 0) {
      const notification = this.notificationQueue.shift()
      
      try {
        await notification()
        await this.delay(500) // 通知间隔
      } catch (error) {
        console.error('Notification error:', error)
      }
    }
    
    this.isProcessing = false
  }

  // 延迟函数
  delay(ms) {
    return new Promise(resolve => setTimeout(resolve, ms))
  }

  // 清除所有消息
  clearAll() {
    ElMessage.closeAll()
    ElNotification.closeAll()
    this.notificationQueue = []
  }
}

// 部门管理专用通知
export class DepartmentNotifications {
  constructor(notificationManager) {
    this.nm = notificationManager
  }

  // 部门操作成功
  departmentCreated(name) {
    this.nm.showSuccess(`部门"${name}"创建成功`)
  }

  departmentUpdated(name) {
    this.nm.showSuccess(`部门"${name}"更新成功`)
  }

  departmentDeleted(name) {
    this.nm.showSuccess(`部门"${name}"删除成功`)
  }

  // 部门操作错误
  departmentNameExists(name) {
    this.nm.showError(`部门名称"${name}"已存在，请使用其他名称`)
  }

  cannotDeleteWithChildren(name) {
    this.nm.showError(`部门"${name}"下还有子部门，请先删除子部门`)
  }

  cannotSelectSelfAsParent() {
    this.nm.showError('不能选择自身或子部门作为上级部门')
  }

  departmentNotFound() {
    this.nm.showError('部门不存在或已被删除')
  }

  // 设备关联通知
  devicesAssociated(count) {
    this.nm.showSuccess(`成功关联 ${count} 个设备`)
  }

  deviceAssociationFailed(deviceName) {
    this.nm.showError(`设备"${deviceName}"关联失败`)
  }

  // 状态变更通知
  departmentEnabled(name) {
    this.nm.showSuccess(`部门"${name}"已启用`)
  }

  departmentDisabled(name) {
    this.nm.showWarning(`部门"${name}"已禁用，其子部门也将被禁用`)
  }

  // 权限相关通知
  noPermissionToDelete() {
    this.nm.showPermissionError('删除部门')
  }

  noPermissionToEdit() {
    this.nm.showPermissionError('编辑部门')
  }

  // 数据加载通知
  loadingDepartments() {
    return this.nm.showLoading('正在加载部门数据...')
  }

  loadingDevices() {
    return this.nm.showLoading('正在加载设备数据...')
  }

  // 确认对话框
  async confirmDelete(name, hasChildren = false, hasDevices = false) {
    let message = `确定要删除部门"${name}"吗？`
    
    if (hasChildren) {
      message += '\n注意：该部门下还有子部门，请先删除子部门。'
      return false
    }
    
    if (hasDevices) {
      message += '\n删除后将解除所有设备关联。'
    }
    
    return await this.nm.showConfirm(message, '删除确认', {
      type: 'warning',
      dangerouslyUseHTMLString: true
    })
  }

  async confirmDisable(name) {
    return await this.nm.showConfirm(
      `确定要禁用部门"${name}"吗？\n禁用后其子部门也将被禁用。`,
      '禁用确认',
      {
        type: 'warning',
        dangerouslyUseHTMLString: true
      }
    )
  }
}

// 创建全局实例
export const userNotificationManager = new UserNotificationManager()
export const departmentNotifications = new DepartmentNotifications(userNotificationManager)

// 便捷函数
export const notify = {
  success: userNotificationManager.showSuccess.bind(userNotificationManager),
  error: userNotificationManager.showError.bind(userNotificationManager),
  warning: userNotificationManager.showWarning.bind(userNotificationManager),
  info: userNotificationManager.showInfo.bind(userNotificationManager),
  loading: userNotificationManager.showLoading.bind(userNotificationManager),
  confirm: userNotificationManager.showConfirm.bind(userNotificationManager),
  prompt: userNotificationManager.showPrompt.bind(userNotificationManager)
}

// 组合式API
export function useNotification() {
  return {
    notify,
    departmentNotifications,
    clearAll: userNotificationManager.clearAll.bind(userNotificationManager)
  }
}