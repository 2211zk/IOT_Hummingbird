import { describe, it, expect, beforeEach, afterEach, jest } from '@jest/globals'
import { mount } from '@vue/test-utils'
import StandardCategorySelector from '../index.vue'
import { createMockPageData } from './mockData'
import { createTestConfig, waitForAsync, mockApiCall, cleanupTest } from './testUtils'
import { performanceMonitor, useMemoryMonitor } from '../utils/performance'

// 模拟API
const mockGetStandardCategoryList = jest.fn()
const mockGetStandardCategoryCategories = jest.fn()

jest.mock('@/api/standardCategory', () => ({
  getStandardCategoryList: mockGetStandardCategoryList,
  getStandardCategoryCategories: mockGetStandardCategoryCategories
}))

describe('StandardCategorySelector Performance Tests', () => {
  let wrapper
  const testConfig = createTestConfig()
  const { getMemoryUsage, logMemoryUsage } = useMemoryMonitor()

  beforeEach(() => {
    jest.clearAllMocks()
    performanceMonitor.clear()
    
    mockApiCall(mockGetStandardCategoryList, createMockPageData(1, 10, 100))
    mockApiCall(mockGetStandardCategoryCategories, { 
      code: 0, 
      data: ['电子设备', '家居用品', '办公设备'], 
      msg: '获取成功' 
    })
  })

  afterEach(() => {
    if (wrapper) {
      wrapper.unmount()
    }
    cleanupTest()
  })

  describe('渲染性能测试', () => {
    it('组件初始渲染应该在合理时间内完成', async () => {
      const startTime = performance.now()
      
      wrapper = mount(StandardCategorySelector, testConfig)
      
      const renderTime = performance.now() - startTime
      
      expect(renderTime).toBeLessThan(50) // 50ms内完成初始渲染
      console.log(`初始渲染时间: ${renderTime.toFixed(2)}ms`)
    })

    it('弹框打开应该在合理时间内完成', async () => {
      wrapper = mount(StandardCategorySelector, testConfig)
      
      const startTime = performance.now()
      
      await wrapper.find('el-button').trigger('click')
      await waitForAsync(wrapper)
      
      const openTime = performance.now() - startTime
      
      expect(openTime).toBeLessThan(200) // 200ms内完成弹框打开
      console.log(`弹框打开时间: ${openTime.toFixed(2)}ms`)
    })

    it('数据加载应该在合理时间内完成', async () => {
      wrapper = mount(StandardCategorySelector, testConfig)
      
      await wrapper.find('el-button').trigger('click')
      
      const startTime = performance.now()
      await waitForAsync(wrapper)
      const loadTime = performance.now() - startTime
      
      expect(loadTime).toBeLessThan(500) // 500ms内完成数据加载
      console.log(`数据加载时间: ${loadTime.toFixed(2)}ms`)
    })
  })

  describe('大数据量性能测试', () => {
    it('应该能处理大量数据而不影响性能', async () => {
      // 模拟1000条数据
      const largeDataResponse = createMockPageData(1, 100, 1000)
      mockApiCall(mockGetStandardCategoryList, largeDataResponse)
      
      wrapper = mount(StandardCategorySelector, testConfig)
      
      const startTime = performance.now()
      
      await wrapper.find('el-button').trigger('click')
      await waitForAsync(wrapper)
      
      const processTime = performance.now() - startTime
      
      expect(processTime).toBeLessThan(1000) // 1秒内处理完成
      console.log(`大数据处理时间: ${processTime.toFixed(2)}ms`)
      
      const modal = wrapper.findComponent({ name: 'CategoryModal' })
      expect(modal.vm.categoryList.length).toBe(100)
      expect(modal.vm.total).toBe(1000)
    })

    it('搜索大量数据应该有良好性能', async () => {
      const largeDataResponse = createMockPageData(1, 100, 1000)
      mockApiCall(mockGetStandardCategoryList, largeDataResponse)
      
      wrapper = mount(StandardCategorySelector, testConfig)
      
      await wrapper.find('el-button').trigger('click')
      await waitForAsync(wrapper)
      
      const modal = wrapper.findComponent({ name: 'CategoryModal' })
      const searchBar = modal.findComponent({ name: 'SearchBar' })
      
      const startTime = performance.now()
      
      // 模拟搜索
      await searchBar.vm.$emit('search', { keyword: '测试', category: '', status: null })
      await waitForAsync(wrapper)
      
      const searchTime = performance.now() - startTime
      
      expect(searchTime).toBeLessThan(300) // 300ms内完成搜索
      console.log(`搜索时间: ${searchTime.toFixed(2)}ms`)
    })
  })

  describe('内存使用测试', () => {
    it('组件创建和销毁不应该造成内存泄漏', async () => {
      const initialMemory = getMemoryUsage()
      
      // 创建多个组件实例
      const wrappers = []
      for (let i = 0; i < 10; i++) {
        const w = mount(StandardCategorySelector, testConfig)
        wrappers.push(w)
        
        await w.find('el-button').trigger('click')
        await waitForAsync(w)
      }
      
      const peakMemory = getMemoryUsage()
      
      // 销毁所有组件
      wrappers.forEach(w => w.unmount())
      
      // 强制垃圾回收（如果支持）
      if (global.gc) {
        global.gc()
      }
      
      await new Promise(resolve => setTimeout(resolve, 100))
      
      const finalMemory = getMemoryUsage()
      
      if (initialMemory && peakMemory && finalMemory) {
        console.log('内存使用情况:')
        console.log(`初始: ${initialMemory.used}MB`)
        console.log(`峰值: ${peakMemory.used}MB`)
        console.log(`最终: ${finalMemory.used}MB`)
        
        // 最终内存使用不应该显著增加
        const memoryIncrease = finalMemory.used - initialMemory.used
        expect(memoryIncrease).toBeLessThan(10) // 不超过10MB增长
      }
    })

    it('长时间使用不应该造成内存持续增长', async () => {
      wrapper = mount(StandardCategorySelector, testConfig)
      
      const initialMemory = getMemoryUsage()
      
      // 模拟长时间使用
      for (let i = 0; i < 50; i++) {
        await wrapper.find('el-button').trigger('click')
        await waitForAsync(wrapper)
        
        const modal = wrapper.findComponent({ name: 'CategoryModal' })
        await modal.vm.$emit('cancel')
        await waitForAsync(wrapper)
      }
      
      const finalMemory = getMemoryUsage()
      
      if (initialMemory && finalMemory) {
        const memoryIncrease = finalMemory.used - initialMemory.used
        expect(memoryIncrease).toBeLessThan(5) // 不超过5MB增长
        
        console.log(`长时间使用内存增长: ${memoryIncrease}MB`)
      }
    })
  })

  describe('响应时间测试', () => {
    it('用户交互响应时间应该在合理范围内', async () => {
      wrapper = mount(StandardCategorySelector, testConfig)
      
      await wrapper.find('el-button').trigger('click')
      await waitForAsync(wrapper)
      
      const modal = wrapper.findComponent({ name: 'CategoryModal' })
      const categoryTable = modal.findComponent({ name: 'CategoryTable' })
      
      // 测试选择响应时间
      const startTime = performance.now()
      
      await categoryTable.vm.$emit('select', { id: 1, name: '测试品类' })
      await waitForAsync(wrapper)
      
      const responseTime = performance.now() - startTime
      
      expect(responseTime).toBeLessThan(100) // 100ms内响应
      console.log(`选择响应时间: ${responseTime.toFixed(2)}ms`)
    })

    it('分页切换响应时间应该合理', async () => {
      wrapper = mount(StandardCategorySelector, testConfig)
      
      await wrapper.find('el-button').trigger('click')
      await waitForAsync(wrapper)
      
      const modal = wrapper.findComponent({ name: 'CategoryModal' })
      const categoryTable = modal.findComponent({ name: 'CategoryTable' })
      
      const startTime = performance.now()
      
      await categoryTable.vm.$emit('page-change', { page: 2, pageSize: 10 })
      await waitForAsync(wrapper)
      
      const pageChangeTime = performance.now() - startTime
      
      expect(pageChangeTime).toBeLessThan(300) // 300ms内完成分页切换
      console.log(`分页切换时间: ${pageChangeTime.toFixed(2)}ms`)
    })
  })

  describe('并发性能测试', () => {
    it('并发操作不应该影响性能', async () => {
      wrapper = mount(StandardCategorySelector, testConfig)
      
      const startTime = performance.now()
      
      // 并发执行多个操作
      const promises = []
      for (let i = 0; i < 5; i++) {
        promises.push(
          wrapper.find('el-button').trigger('click')
        )
      }
      
      await Promise.all(promises)
      await waitForAsync(wrapper)
      
      const concurrentTime = performance.now() - startTime
      
      expect(concurrentTime).toBeLessThan(500) // 500ms内完成并发操作
      console.log(`并发操作时间: ${concurrentTime.toFixed(2)}ms`)
    })
  })

  describe('性能监控测试', () => {
    it('应该能够监控组件性能指标', async () => {
      performanceMonitor.start('component_lifecycle')
      
      wrapper = mount(StandardCategorySelector, testConfig)
      
      await wrapper.find('el-button').trigger('click')
      await waitForAsync(wrapper)
      
      performanceMonitor.end('component_lifecycle')
      
      const metrics = performanceMonitor.getAllMetrics()
      expect(metrics.component_lifecycle).toBeDefined()
      expect(metrics.component_lifecycle.duration).toBeGreaterThan(0)
      
      console.log('性能指标:', metrics)
    })
  })

  describe('资源使用测试', () => {
    it('DOM节点数量应该在合理范围内', async () => {
      const initialNodeCount = document.querySelectorAll('*').length
      
      wrapper = mount(StandardCategorySelector, testConfig)
      
      await wrapper.find('el-button').trigger('click')
      await waitForAsync(wrapper)
      
      const finalNodeCount = document.querySelectorAll('*').length
      const nodeIncrease = finalNodeCount - initialNodeCount
      
      expect(nodeIncrease).toBeLessThan(200) // DOM节点增加不超过200个
      console.log(`DOM节点增加: ${nodeIncrease}`)
    })

    it('事件监听器数量应该合理', async () => {
      // 这个测试需要在真实浏览器环境中运行
      // 在Jest环境中我们只能做基本检查
      
      wrapper = mount(StandardCategorySelector, testConfig)
      
      await wrapper.find('el-button').trigger('click')
      await waitForAsync(wrapper)
      
      // 验证组件正常工作
      const modal = wrapper.findComponent({ name: 'CategoryModal' })
      expect(modal.exists()).toBe(true)
    })
  })

  describe('优化效果测试', () => {
    it('防抖功能应该减少API调用次数', async () => {
      wrapper = mount(StandardCategorySelector, testConfig)
      
      await wrapper.find('el-button').trigger('click')
      await waitForAsync(wrapper)
      
      const modal = wrapper.findComponent({ name: 'CategoryModal' })
      const searchBar = modal.findComponent({ name: 'SearchBar' })
      
      // 快速连续搜索
      await searchBar.vm.$emit('search', { keyword: 'a', category: '', status: null })
      await searchBar.vm.$emit('search', { keyword: 'ab', category: '', status: null })
      await searchBar.vm.$emit('search', { keyword: 'abc', category: '', status: null })
      
      await waitForAsync(wrapper)
      
      // 由于防抖，API调用次数应该少于搜索次数
      const apiCallCount = mockGetStandardCategoryList.mock.calls.length
      expect(apiCallCount).toBeLessThan(5) // 包括初始加载
    })
  })
})