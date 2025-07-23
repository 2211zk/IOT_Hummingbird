import { ref, reactive } from 'vue'
import { ElLoading } from 'element-plus'

/**
 * 加载状态管理器
 */
class LoadingManager {
  constructor() {
    this.loadingStates = reactive({})
    this.loadingInstances = new Map()
  }

  /**
   * 设置加载状态
   * @param {string} key - 加载状态的键
   * @param {boolean} loading - 是否加载中
   */
  setLoading(key, loading) {
    this.loadingStates[key] = loading
  }

  /**
   * 获取加载状态
   * @param {string} key - 加载状态的键
   * @returns {boolean}
   */
  getLoading(key) {
    return this.loadingStates[key] || false
  }

  /**
   * 显示全屏加载
   * @param {string} text - 加载文本
   * @param {string} key - 加载实例的键
   */
  showFullScreenLoading(text = '加载中...', key = 'default') {
    if (this.loadingInstances.has(key)) {
      return
    }

    const loading = ElLoading.service({
      lock: true,
      text,
      background: 'rgba(0, 0, 0, 0.7)'
    })

    this.loadingInstances.set(key, loading)
  }

  /**
   * 隐藏全屏加载
   * @param {string} key - 加载实例的键
   */
  hideFullScreenLoading(key = 'default') {
    const loading = this.loadingInstances.get(key)
    if (loading) {
      loading.close()
      this.loadingInstances.delete(key)
    }
  }

  /**
   * 清除所有加载状态
   */
  clearAll() {
    // 清除所有加载状态
    Object.keys(this.loadingStates).forEach(key => {
      this.loadingStates[key] = false
    })

    // 关闭所有全屏加载
    this.loadingInstances.forEach(loading => {
      loading.close()
    })
    this.loadingInstances.clear()
  }
}

// 创建全局加载管理器实例
export const loadingManager = new LoadingManager()

/**
 * 创建加载状态的组合式函数
 * @param {string} initialKey - 初始键名
 */
export const useLoading = (initialKey = 'default') => {
  const loading = ref(false)
  const loadingText = ref('加载中...')

  const setLoading = (value, text = '加载中...') => {
    loading.value = value
    loadingText.value = text
    loadingManager.setLoading(initialKey, value)
  }

  const withLoading = async (asyncFn, text = '加载中...') => {
    try {
      setLoading(true, text)
      const result = await asyncFn()
      return result
    } finally {
      setLoading(false)
    }
  }

  return {
    loading,
    loadingText,
    setLoading,
    withLoading
  }
}

/**
 * 批量加载状态管理
 */
export const useBatchLoading = () => {
  const loadingStates = reactive({})

  const setLoading = (key, value, text = '加载中...') => {
    loadingStates[key] = {
      loading: value,
      text
    }
    loadingManager.setLoading(key, value)
  }

  const getLoading = (key) => {
    return loadingStates[key]?.loading || false
  }

  const getLoadingText = (key) => {
    return loadingStates[key]?.text || '加载中...'
  }

  const isAnyLoading = () => {
    return Object.values(loadingStates).some(state => state.loading)
  }

  const clearAllLoading = () => {
    Object.keys(loadingStates).forEach(key => {
      loadingStates[key] = {
        loading: false,
        text: '加载中...'
      }
      loadingManager.setLoading(key, false)
    })
  }

  return {
    loadingStates,
    setLoading,
    getLoading,
    getLoadingText,
    isAnyLoading,
    clearAllLoading
  }
}

/**
 * 加载状态装饰器
 * @param {string} loadingKey - 加载状态键
 * @param {string} loadingText - 加载文本
 */
export const withLoadingState = (loadingKey, loadingText = '加载中...') => {
  return (target, propertyKey, descriptor) => {
    const originalMethod = descriptor.value

    descriptor.value = async function(...args) {
      try {
        loadingManager.setLoading(loadingKey, true)
        const result = await originalMethod.apply(this, args)
        return result
      } finally {
        loadingManager.setLoading(loadingKey, false)
      }
    }

    return descriptor
  }
}

/**
 * 延迟加载工具
 * @param {number} delay - 延迟时间（毫秒）
 */
export const delayedLoading = (delay = 300) => {
  let timer = null
  let isLoading = false

  const show = (callback) => {
    if (isLoading) return

    timer = setTimeout(() => {
      isLoading = true
      callback(true)
    }, delay)
  }

  const hide = (callback) => {
    if (timer) {
      clearTimeout(timer)
      timer = null
    }

    if (isLoading) {
      isLoading = false
      callback(false)
    }
  }

  return { show, hide }
}