// 性能监控工具

class PerformanceMonitor {
  constructor() {
    this.metrics = new Map()
    this.observers = []
    this.isEnabled = process.env.NODE_ENV === 'development'
  }

  // 开始性能测量
  startMeasure(name) {
    if (!this.isEnabled) return

    const startTime = performance.now()
    this.metrics.set(name, {
      startTime,
      endTime: null,
      duration: null,
      memory: this.getMemoryUsage()
    })
  }

  // 结束性能测量
  endMeasure(name) {
    if (!this.isEnabled) return

    const metric = this.metrics.get(name)
    if (!metric) {
      console.warn(`Performance measure "${name}" not found`)
      return
    }

    const endTime = performance.now()
    const duration = endTime - metric.startTime

    metric.endTime = endTime
    metric.duration = duration
    metric.memoryAfter = this.getMemoryUsage()

    // 记录性能日志
    this.logPerformance(name, metric)

    return metric
  }

  // 获取内存使用情况
  getMemoryUsage() {
    if (performance.memory) {
      return {
        used: performance.memory.usedJSHeapSize,
        total: performance.memory.totalJSHeapSize,
        limit: performance.memory.jsHeapSizeLimit
      }
    }
    return null
  }

  // 记录性能日志
  logPerformance(name, metric) {
    const { duration, memory, memoryAfter } = metric
    
    console.group(`🚀 Performance: ${name}`)
    console.log(`Duration: ${duration.toFixed(2)}ms`)
    
    if (memory && memoryAfter) {
      const memoryDiff = memoryAfter.used - memory.used
      console.log(`Memory: ${this.formatBytes(memoryAfter.used)} (${memoryDiff > 0 ? '+' : ''}${this.formatBytes(memoryDiff)})`)
    }
    
    // 性能警告
    if (duration > 1000) {
      console.warn(`⚠️ Slow operation detected: ${duration.toFixed(2)}ms`)
    }
    
    console.groupEnd()
  }

  // 格式化字节数
  formatBytes(bytes) {
    if (bytes === 0) return '0 Bytes'
    
    const k = 1024
    const sizes = ['Bytes', 'KB', 'MB', 'GB']
    const i = Math.floor(Math.log(bytes) / Math.log(k))
    
    return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
  }

  // 监控API调用性能
  monitorApiCall(apiName, apiFunction) {
    return async (...args) => {
      this.startMeasure(`API: ${apiName}`)
      
      try {
        const result = await apiFunction(...args)
        this.endMeasure(`API: ${apiName}`)
        return result
      } catch (error) {
        this.endMeasure(`API: ${apiName}`)
        throw error
      }
    }
  }

  // 监控组件渲染性能
  monitorComponentRender(componentName, renderFunction) {
    return (...args) => {
      this.startMeasure(`Render: ${componentName}`)
      
      try {
        const result = renderFunction(...args)
        this.endMeasure(`Render: ${componentName}`)
        return result
      } catch (error) {
        this.endMeasure(`Render: ${componentName}`)
        throw error
      }
    }
  }

  // 监控长任务
  observeLongTasks() {
    if (!this.isEnabled || !window.PerformanceObserver) return

    try {
      const observer = new PerformanceObserver((list) => {
        for (const entry of list.getEntries()) {
          if (entry.duration > 50) {
            console.warn(`🐌 Long task detected: ${entry.duration.toFixed(2)}ms`, entry)
          }
        }
      })

      observer.observe({ entryTypes: ['longtask'] })
      this.observers.push(observer)
    } catch (error) {
      console.warn('Long task observation not supported')
    }
  }

  // 监控资源加载
  observeResourceTiming() {
    if (!this.isEnabled || !window.PerformanceObserver) return

    try {
      const observer = new PerformanceObserver((list) => {
        for (const entry of list.getEntries()) {
          if (entry.duration > 1000) {
            console.warn(`🐌 Slow resource: ${entry.name} (${entry.duration.toFixed(2)}ms)`)
          }
        }
      })

      observer.observe({ entryTypes: ['resource'] })
      this.observers.push(observer)
    } catch (error) {
      console.warn('Resource timing observation not supported')
    }
  }

  // 获取页面性能指标
  getPageMetrics() {
    if (!performance.timing) return null

    const timing = performance.timing
    const navigation = performance.navigation

    return {
      // 页面加载时间
      pageLoadTime: timing.loadEventEnd - timing.navigationStart,
      // DNS查询时间
      dnsTime: timing.domainLookupEnd - timing.domainLookupStart,
      // TCP连接时间
      tcpTime: timing.connectEnd - timing.connectStart,
      // 请求时间
      requestTime: timing.responseEnd - timing.requestStart,
      // 解析DOM时间
      domParseTime: timing.domContentLoadedEventEnd - timing.domLoading,
      // 白屏时间
      whiteScreenTime: timing.responseStart - timing.navigationStart,
      // 首屏时间
      firstScreenTime: timing.domContentLoadedEventEnd - timing.navigationStart,
      // 导航类型
      navigationType: navigation.type,
      // 重定向次数
      redirectCount: navigation.redirectCount
    }
  }

  // 监控内存泄漏
  monitorMemoryLeaks() {
    if (!this.isEnabled) return

    let lastMemoryUsage = this.getMemoryUsage()
    
    setInterval(() => {
      const currentMemory = this.getMemoryUsage()
      if (currentMemory && lastMemoryUsage) {
        const memoryIncrease = currentMemory.used - lastMemoryUsage.used
        const increasePercent = (memoryIncrease / lastMemoryUsage.used) * 100
        
        if (increasePercent > 10) {
          console.warn(`🚨 Potential memory leak detected: ${this.formatBytes(memoryIncrease)} increase (${increasePercent.toFixed(2)}%)`)
        }
      }
      lastMemoryUsage = currentMemory
    }, 30000) // 每30秒检查一次
  }

  // 获取所有性能指标
  getAllMetrics() {
    const metrics = {}
    
    for (const [name, metric] of this.metrics) {
      metrics[name] = {
        duration: metric.duration,
        memory: metric.memoryAfter ? metric.memoryAfter.used : null
      }
    }
    
    return {
      customMetrics: metrics,
      pageMetrics: this.getPageMetrics(),
      currentMemory: this.getMemoryUsage()
    }
  }

  // 清理观察者
  cleanup() {
    this.observers.forEach(observer => observer.disconnect())
    this.observers = []
    this.metrics.clear()
  }

  // 启用/禁用监控
  setEnabled(enabled) {
    this.isEnabled = enabled
  }
}

// 创建全局实例
export const performanceMonitor = new PerformanceMonitor()

// 装饰器函数
export function measurePerformance(name) {
  return function(target, propertyKey, descriptor) {
    const originalMethod = descriptor.value
    
    descriptor.value = async function(...args) {
      performanceMonitor.startMeasure(`${target.constructor.name}.${propertyKey}`)
      
      try {
        const result = await originalMethod.apply(this, args)
        performanceMonitor.endMeasure(`${target.constructor.name}.${propertyKey}`)
        return result
      } catch (error) {
        performanceMonitor.endMeasure(`${target.constructor.name}.${propertyKey}`)
        throw error
      }
    }
    
    return descriptor
  }
}

// 便捷函数
export function withPerformanceMonitoring(name, fn) {
  return performanceMonitor.monitorApiCall(name, fn)
}

// 初始化性能监控
export function initPerformanceMonitoring() {
  performanceMonitor.observeLongTasks()
  performanceMonitor.observeResourceTiming()
  performanceMonitor.monitorMemoryLeaks()
  
  // 页面卸载时输出性能报告
  window.addEventListener('beforeunload', () => {
    const metrics = performanceMonitor.getAllMetrics()
    console.log('📊 Performance Report:', metrics)
  })
}

// 组合式API
export function usePerformanceMonitor() {
  return {
    startMeasure: performanceMonitor.startMeasure.bind(performanceMonitor),
    endMeasure: performanceMonitor.endMeasure.bind(performanceMonitor),
    monitorApiCall: performanceMonitor.monitorApiCall.bind(performanceMonitor),
    getAllMetrics: performanceMonitor.getAllMetrics.bind(performanceMonitor)
  }
}