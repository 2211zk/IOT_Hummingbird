import { ElMessage } from 'element-plus'

// 重试管理器
class RetryManager {
  constructor() {
    this.retryConfigs = new Map()
    this.activeRetries = new Map()
  }

  // 设置重试配置
  setRetryConfig(key, config = {}) {
    const defaultConfig = {
      maxRetries: 3,
      delay: 1000,
      backoff: 'exponential', // linear | exponential | fixed
      shouldRetry: (error) => {
        // 默认只重试网络错误和5xx错误
        return !error.response || error.response.status >= 500
      },
      onRetry: (attempt, error) => {
        console.log(`Retry attempt ${attempt} for ${key}:`, error.message)
      },
      onMaxRetriesReached: (error) => {
        ElMessage.error(`操作失败，已重试${config.maxRetries || 3}次`)
      }
    }
    
    this.retryConfigs.set(key, { ...defaultConfig, ...config })
  }

  // 执行带重试的异步操作
  async executeWithRetry(key, asyncFn, ...args) {
    const config = this.retryConfigs.get(key) || {}
    const {
      maxRetries = 3,
      delay = 1000,
      backoff = 'exponential',
      shouldRetry = () => true,
      onRetry = () => {},
      onMaxRetriesReached = () => {}
    } = config

    let lastError
    let attempt = 0

    // 检查是否已有相同的重试在进行
    if (this.activeRetries.has(key)) {
      return this.activeRetries.get(key)
    }

    const retryPromise = (async () => {
      while (attempt <= maxRetries) {
        try {
          const result = await asyncFn(...args)
          return result
        } catch (error) {
          lastError = error
          attempt++

          if (attempt > maxRetries || !shouldRetry(error)) {
            break
          }

          // 调用重试回调
          onRetry(attempt, error)

          // 计算延迟时间
          const retryDelay = this.calculateDelay(delay, attempt, backoff)
          await this.sleep(retryDelay)
        }
      }

      // 达到最大重试次数
      onMaxRetriesReached(lastError)
      throw lastError
    })()

    // 记录活跃的重试
    this.activeRetries.set(key, retryPromise)

    try {
      const result = await retryPromise
      return result
    } finally {
      this.activeRetries.delete(key)
    }
  }

  // 计算延迟时间
  calculateDelay(baseDelay, attempt, backoff) {
    switch (backoff) {
      case 'linear':
        return baseDelay * attempt
      case 'exponential':
        return baseDelay * Math.pow(2, attempt - 1)
      case 'fixed':
      default:
        return baseDelay
    }
  }

  // 睡眠函数
  sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms))
  }

  // 取消重试
  cancelRetry(key) {
    this.activeRetries.delete(key)
  }

  // 清除所有重试
  clearAllRetries() {
    this.activeRetries.clear()
  }
}

// 断路器模式
class CircuitBreaker {
  constructor(options = {}) {
    this.failureThreshold = options.failureThreshold || 5
    this.resetTimeout = options.resetTimeout || 60000
    this.monitoringPeriod = options.monitoringPeriod || 10000
    
    this.state = 'CLOSED' // CLOSED | OPEN | HALF_OPEN
    this.failureCount = 0
    this.lastFailureTime = null
    this.successCount = 0
  }

  // 执行操作
  async execute(asyncFn, ...args) {
    if (this.state === 'OPEN') {
      if (Date.now() - this.lastFailureTime > this.resetTimeout) {
        this.state = 'HALF_OPEN'
        this.successCount = 0
      } else {
        throw new Error('Circuit breaker is OPEN')
      }
    }

    try {
      const result = await asyncFn(...args)
      this.onSuccess()
      return result
    } catch (error) {
      this.onFailure()
      throw error
    }
  }

  // 成功处理
  onSuccess() {
    this.failureCount = 0
    
    if (this.state === 'HALF_OPEN') {
      this.successCount++
      if (this.successCount >= 3) {
        this.state = 'CLOSED'
      }
    }
  }

  // 失败处理
  onFailure() {
    this.failureCount++
    this.lastFailureTime = Date.now()
    
    if (this.failureCount >= this.failureThreshold) {
      this.state = 'OPEN'
    }
  }

  // 获取状态
  getState() {
    return this.state
  }

  // 重置断路器
  reset() {
    this.state = 'CLOSED'
    this.failureCount = 0
    this.lastFailureTime = null
    this.successCount = 0
  }
}

// 创建全局实例
export const retryManager = new RetryManager()

// 便捷函数
export function withRetry(key, asyncFn, config = {}) {
  retryManager.setRetryConfig(key, config)
  return (...args) => retryManager.executeWithRetry(key, asyncFn, ...args)
}

export function createCircuitBreaker(options = {}) {
  return new CircuitBreaker(options)
}

// 组合式API
export function useRetry() {
  return {
    setRetryConfig: retryManager.setRetryConfig.bind(retryManager),
    executeWithRetry: retryManager.executeWithRetry.bind(retryManager),
    cancelRetry: retryManager.cancelRetry.bind(retryManager)
  }
}

// 预设的重试配置
export const RETRY_CONFIGS = {
  // 部门相关操作
  DEPARTMENT_LIST: {
    maxRetries: 2,
    delay: 500,
    backoff: 'exponential'
  },
  DEPARTMENT_TREE: {
    maxRetries: 3,
    delay: 300,
    backoff: 'linear'
  },
  DEPARTMENT_CREATE: {
    maxRetries: 1,
    delay: 1000,
    shouldRetry: (error) => error.response?.status >= 500
  },
  DEPARTMENT_UPDATE: {
    maxRetries: 1,
    delay: 1000,
    shouldRetry: (error) => error.response?.status >= 500
  },
  DEPARTMENT_DELETE: {
    maxRetries: 1,
    delay: 1000,
    shouldRetry: (error) => error.response?.status >= 500
  },
  DEVICE_LIST: {
    maxRetries: 2,
    delay: 500,
    backoff: 'exponential'
  }
}