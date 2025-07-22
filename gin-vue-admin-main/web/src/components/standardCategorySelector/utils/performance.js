/**
 * 性能优化工具集
 */

import { ref, computed, nextTick } from 'vue'

/**
 * 虚拟滚动实现
 * @param {Array} items - 数据项
 * @param {number} itemHeight - 每项高度
 * @param {number} containerHeight - 容器高度
 */
export const useVirtualScroll = (items, itemHeight = 50, containerHeight = 400) => {
  const scrollTop = ref(0)
  const visibleCount = Math.ceil(containerHeight / itemHeight)
  const bufferSize = 5 // 缓冲区大小

  const visibleItems = computed(() => {
    const startIndex = Math.max(0, Math.floor(scrollTop.value / itemHeight) - bufferSize)
    const endIndex = Math.min(items.value.length, startIndex + visibleCount + bufferSize * 2)
    
    return {
      items: items.value.slice(startIndex, endIndex),
      startIndex,
      endIndex,
      offsetY: startIndex * itemHeight
    }
  })

  const totalHeight = computed(() => items.value.length * itemHeight)

  const handleScroll = (event) => {
    scrollTop.value = event.target.scrollTop
  }

  return {
    visibleItems,
    totalHeight,
    handleScroll,
    scrollTop
  }
}

/**
 * 防抖 Hook
 * @param {Function} fn - 要防抖的函数
 * @param {number} delay - 延迟时间
 */
export const useDebounce = (fn, delay = 300) => {
  let timer = null

  const debouncedFn = (...args) => {
    if (timer) {
      clearTimeout(timer)
    }
    timer = setTimeout(() => {
      fn.apply(this, args)
      timer = null
    }, delay)
  }

  const cancel = () => {
    if (timer) {
      clearTimeout(timer)
      timer = null
    }
  }

  return { debouncedFn, cancel }
}

/**
 * 节流 Hook
 * @param {Function} fn - 要节流的函数
 * @param {number} delay - 延迟时间
 */
export const useThrottle = (fn, delay = 300) => {
  let timer = null
  let lastExecTime = 0

  const throttledFn = (...args) => {
    const currentTime = Date.now()
    
    if (currentTime - lastExecTime > delay) {
      fn.apply(this, args)
      lastExecTime = currentTime
    } else if (!timer) {
      timer = setTimeout(() => {
        fn.apply(this, args)
        lastExecTime = Date.now()
        timer = null
      }, delay - (currentTime - lastExecTime))
    }
  }

  const cancel = () => {
    if (timer) {
      clearTimeout(timer)
      timer = null
    }
  }

  return { throttledFn, cancel }
}

/**
 * 缓存管理
 */
class CacheManager {
  constructor(maxSize = 100, ttl = 5 * 60 * 1000) { // 默认5分钟过期
    this.cache = new Map()
    this.maxSize = maxSize
    this.ttl = ttl
  }

  set(key, value) {
    // 如果缓存已满，删除最旧的项
    if (this.cache.size >= this.maxSize) {
      const firstKey = this.cache.keys().next().value
      this.cache.delete(firstKey)
    }

    this.cache.set(key, {
      value,
      timestamp: Date.now()
    })
  }

  get(key) {
    const item = this.cache.get(key)
    
    if (!item) {
      return null
    }

    // 检查是否过期
    if (Date.now() - item.timestamp > this.ttl) {
      this.cache.delete(key)
      return null
    }

    return item.value
  }

  has(key) {
    return this.get(key) !== null
  }

  clear() {
    this.cache.clear()
  }

  size() {
    return this.cache.size
  }
}

// 创建全局缓存实例
export const apiCache = new CacheManager()

/**
 * 带缓存的 API 调用
 * @param {string} key - 缓存键
 * @param {Function} apiCall - API 调用函数
 * @param {boolean} useCache - 是否使用缓存
 */
export const cachedApiCall = async (key, apiCall, useCache = true) => {
  if (useCache && apiCache.has(key)) {
    return apiCache.get(key)
  }

  const result = await apiCall()
  
  if (useCache && result) {
    apiCache.set(key, result)
  }

  return result
}

/**
 * 图片懒加载
 */
export const useLazyLoad = () => {
  const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        const img = entry.target
        const src = img.dataset.src
        
        if (src) {
          img.src = src
          img.removeAttribute('data-src')
          observer.unobserve(img)
        }
      }
    })
  })

  const observe = (element) => {
    observer.observe(element)
  }

  const unobserve = (element) => {
    observer.unobserve(element)
  }

  const disconnect = () => {
    observer.disconnect()
  }

  return { observe, unobserve, disconnect }
}

/**
 * 性能监控
 */
export class PerformanceMonitor {
  constructor() {
    this.metrics = new Map()
  }

  start(name) {
    this.metrics.set(name, {
      startTime: performance.now(),
      endTime: null,
      duration: null
    })
  }

  end(name) {
    const metric = this.metrics.get(name)
    if (metric) {
      metric.endTime = performance.now()
      metric.duration = metric.endTime - metric.startTime
    }
  }

  getMetric(name) {
    return this.metrics.get(name)
  }

  getAllMetrics() {
    const result = {}
    this.metrics.forEach((value, key) => {
      result[key] = value
    })
    return result
  }

  clear() {
    this.metrics.clear()
  }

  // 记录组件渲染时间
  measureRender(componentName, renderFn) {
    return async (...args) => {
      this.start(`${componentName}_render`)
      const result = await renderFn(...args)
      await nextTick()
      this.end(`${componentName}_render`)
      return result
    }
  }

  // 记录 API 调用时间
  measureApi(apiName, apiFn) {
    return async (...args) => {
      this.start(`${apiName}_api`)
      try {
        const result = await apiFn(...args)
        this.end(`${apiName}_api`)
        return result
      } catch (error) {
        this.end(`${apiName}_api`)
        throw error
      }
    }
  }
}

// 创建全局性能监控实例
export const performanceMonitor = new PerformanceMonitor()

/**
 * 内存使用监控
 */
export const useMemoryMonitor = () => {
  const getMemoryUsage = () => {
    if (performance.memory) {
      return {
        used: Math.round(performance.memory.usedJSHeapSize / 1048576), // MB
        total: Math.round(performance.memory.totalJSHeapSize / 1048576), // MB
        limit: Math.round(performance.memory.jsHeapSizeLimit / 1048576) // MB
      }
    }
    return null
  }

  const logMemoryUsage = (label = 'Memory Usage') => {
    const usage = getMemoryUsage()
    if (usage) {
      console.log(`${label}:`, usage)
    }
  }

  return { getMemoryUsage, logMemoryUsage }
}

/**
 * 组件卸载时清理资源
 */
export const useCleanup = () => {
  const cleanupTasks = []

  const addCleanupTask = (task) => {
    cleanupTasks.push(task)
  }

  const cleanup = () => {
    cleanupTasks.forEach(task => {
      try {
        task()
      } catch (error) {
        console.error('Cleanup task failed:', error)
      }
    })
    cleanupTasks.length = 0
  }

  return { addCleanupTask, cleanup }
}

/**
 * 批量更新优化
 */
export const useBatchUpdate = () => {
  let updateQueue = []
  let isUpdating = false

  const batchUpdate = (updateFn) => {
    updateQueue.push(updateFn)
    
    if (!isUpdating) {
      isUpdating = true
      nextTick(() => {
        const updates = [...updateQueue]
        updateQueue = []
        isUpdating = false
        
        updates.forEach(update => {
          try {
            update()
          } catch (error) {
            console.error('Batch update failed:', error)
          }
        })
      })
    }
  }

  return { batchUpdate }
}

/**
 * 组件预加载
 */
export const preloadComponent = (componentLoader) => {
  return componentLoader().catch(error => {
    console.warn('Component preload failed:', error)
  })
}

/**
 * 资源预加载
 */
export const preloadResource = (url, type = 'fetch') => {
  return new Promise((resolve, reject) => {
    if (type === 'image') {
      const img = new Image()
      img.onload = () => resolve(img)
      img.onerror = reject
      img.src = url
    } else {
      fetch(url)
        .then(response => response.json())
        .then(resolve)
        .catch(reject)
    }
  })
}