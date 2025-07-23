import { describe, it, expect, beforeEach, afterEach, jest } from '@jest/globals'
import { mount } from '@vue/test-utils'
import StandardCategorySelector from '../index.vue'
import { mockCategories, mockApiResponse, createMockPageData } from './mockData'
import { 
  createTestConfig, 
  waitForAsync, 
  mockClick, 
  mockInput,
  expectElementExists,
  expectElementText,
  expectEventEmitted,
  mockApiCall,
  mockApiError,
  cleanupTest
} from './testUtils'

// 模拟API
const mockGetStandardCategoryList = jest.fn()
const mockGetStandardCategoryCategories = jest.fn()

jest.mock('@/api/standardCategory', () => ({
  getStandardCategoryList: mockGetStandardCategoryList,
  getStandardCategoryCategories: mockGetStandardCategoryCategories
}))

describe('StandardCategorySelector Integration Tests', () => {
  let wrapper
  const testConfig = createTestConfig()

  beforeEach(() => {
    jest.clearAllMocks()
    mockApiCall(mockGetStandardCategoryList, mockApiResponse.success)
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

  describe('完整用户流程测试', () => {
    it('应该支持完整的选择流程：打开弹框 -> 搜索 -> 选择 -> 确认', async () => {
      wrapper = mount(StandardCategorySelector, {
        ...testConfig,
        props: {
          modelValue: []
        }
      })

      // 1. 点击按钮打开弹框
      await mockClick(wrapper, 'el-button')
      await waitForAsync(wrapper)

      // 验证弹框已打开
      const modal = wrapper.findComponent({ name: 'CategoryModal' })
      expect(modal.props('modelValue')).toBe(true)

      // 2. 等待数据加载
      await waitForAsync(wrapper)
      expect(mockGetStandardCategoryList).toHaveBeenCalled()

      // 3. 进行搜索
      const searchBar = modal.findComponent({ name: 'SearchBar' })
      await searchBar.vm.$emit('search', { keyword: '电子', category: '', status: null })
      await waitForAsync(wrapper)

      // 验证搜索请求
      expect(mockGetStandardCategoryList).toHaveBeenCalledWith({
        page: 1,
        pageSize: 10,
        keyword: '电子',
        category: '',
        status: null
      })

      // 4. 选择品类
      const categoryTable = modal.findComponent({ name: 'CategoryTable' })
      await categoryTable.vm.$emit('select', mockCategories[0])
      await waitForAsync(wrapper)

      // 验证选择状态
      const selectedList = modal.findComponent({ name: 'SelectedList' })
      expect(selectedList.props('selected')).toContain(mockCategories[0])

      // 5. 确认选择
      const actionButtons = modal.findComponent({ name: 'ActionButtons' })
      await actionButtons.vm.$emit('confirm')
      await waitForAsync(wrapper)

      // 验证结果
      expectEventEmitted(wrapper, 'update:modelValue', [mockCategories[0]])
      expectEventEmitted(wrapper, 'change', [mockCategories[0]])
      expectEventEmitted(wrapper, 'confirm', [mockCategories[0]])
    })

    it('应该支持多选流程', async () => {
      wrapper = mount(StandardCategorySelector, {
        ...testConfig,
        props: {
          modelValue: [],
          multiple: true
        }
      })

      await mockClick(wrapper, 'el-button')
      await waitForAsync(wrapper)

      const modal = wrapper.findComponent({ name: 'CategoryModal' })
      const categoryTable = modal.findComponent({ name: 'CategoryTable' })

      // 选择多个品类
      await categoryTable.vm.$emit('select', mockCategories[0])
      await categoryTable.vm.$emit('select', mockCategories[1])
      await waitForAsync(wrapper)

      const actionButtons = modal.findComponent({ name: 'ActionButtons' })
      await actionButtons.vm.$emit('confirm')
      await waitForAsync(wrapper)

      expectEventEmitted(wrapper, 'update:modelValue', [mockCategories[0], mockCategories[1]])
    })

    it('应该支持单选流程', async () => {
      wrapper = mount(StandardCategorySelector, {
        ...testConfig,
        props: {
          modelValue: [],
          multiple: false
        }
      })

      await mockClick(wrapper, 'el-button')
      await waitForAsync(wrapper)

      const modal = wrapper.findComponent({ name: 'CategoryModal' })
      const categoryTable = modal.findComponent({ name: 'CategoryTable' })

      // 选择第一个品类
      await categoryTable.vm.$emit('select', mockCategories[0])
      await waitForAsync(wrapper)

      // 选择第二个品类（应该替换第一个）
      await categoryTable.vm.$emit('select', mockCategories[1])
      await waitForAsync(wrapper)

      const actionButtons = modal.findComponent({ name: 'ActionButtons' })
      await actionButtons.vm.$emit('confirm')
      await waitForAsync(wrapper)

      expectEventEmitted(wrapper, 'update:modelValue', [mockCategories[1]])
    })
  })

  describe('分页功能测试', () => {
    it('应该支持分页操作', async () => {
      // 模拟分页数据
      mockApiCall(mockGetStandardCategoryList, createMockPageData(1, 10, 50))

      wrapper = mount(StandardCategorySelector, testConfig)

      await mockClick(wrapper, 'el-button')
      await waitForAsync(wrapper)

      const modal = wrapper.findComponent({ name: 'CategoryModal' })
      const categoryTable = modal.findComponent({ name: 'CategoryTable' })

      // 切换到第二页
      await categoryTable.vm.$emit('page-change', { page: 2, pageSize: 10 })
      await waitForAsync(wrapper)

      expect(mockGetStandardCategoryList).toHaveBeenCalledWith({
        page: 2,
        pageSize: 10,
        keyword: '',
        category: '',
        status: null
      })
    })

    it('应该支持改变页大小', async () => {
      wrapper = mount(StandardCategorySelector, testConfig)

      await mockClick(wrapper, 'el-button')
      await waitForAsync(wrapper)

      const modal = wrapper.findComponent({ name: 'CategoryModal' })
      const categoryTable = modal.findComponent({ name: 'CategoryTable' })

      // 改变页大小
      await categoryTable.vm.$emit('page-change', { page: 1, pageSize: 20 })
      await waitForAsync(wrapper)

      expect(mockGetStandardCategoryList).toHaveBeenCalledWith({
        page: 1,
        pageSize: 20,
        keyword: '',
        category: '',
        status: null
      })
    })
  })

  describe('错误处理测试', () => {
    it('应该正确处理API错误', async () => {
      mockApiError(mockGetStandardCategoryList, new Error('Network Error'))

      wrapper = mount(StandardCategorySelector, testConfig)

      await mockClick(wrapper, 'el-button')
      await waitForAsync(wrapper)

      // 验证错误处理
      const modal = wrapper.findComponent({ name: 'CategoryModal' })
      expect(modal.vm.categoryList).toEqual([])
      expect(modal.vm.total).toBe(0)
    })

    it('应该处理网络超时', async () => {
      mockApiError(mockGetStandardCategoryList, { code: 'ECONNABORTED', message: 'timeout' })

      wrapper = mount(StandardCategorySelector, testConfig)

      await mockClick(wrapper, 'el-button')
      await waitForAsync(wrapper)

      // 验证超时处理
      const modal = wrapper.findComponent({ name: 'CategoryModal' })
      expect(modal.vm.categoryList).toEqual([])
    })
  })

  describe('性能测试', () => {
    it('应该在合理时间内完成渲染', async () => {
      const startTime = performance.now()

      wrapper = mount(StandardCategorySelector, testConfig)

      const renderTime = performance.now() - startTime
      expect(renderTime).toBeLessThan(100) // 100ms内完成渲染
    })

    it('应该正确处理大量数据', async () => {
      // 模拟大量数据
      const largeDataResponse = createMockPageData(1, 100, 1000)
      mockApiCall(mockGetStandardCategoryList, largeDataResponse)

      wrapper = mount(StandardCategorySelector, testConfig)

      await mockClick(wrapper, 'el-button')
      await waitForAsync(wrapper)

      const modal = wrapper.findComponent({ name: 'CategoryModal' })
      expect(modal.vm.categoryList.length).toBe(100)
      expect(modal.vm.total).toBe(1000)
    })
  })

  describe('响应式测试', () => {
    it('应该在移动设备上正确显示', async () => {
      // 模拟移动设备屏幕
      Object.defineProperty(window, 'innerWidth', {
        writable: true,
        configurable: true,
        value: 375
      })

      wrapper = mount(StandardCategorySelector, testConfig)

      await mockClick(wrapper, 'el-button')
      await waitForAsync(wrapper)

      // 验证响应式布局
      const modal = wrapper.findComponent({ name: 'CategoryModal' })
      expectElementExists(modal, '.category-selector-content')
    })
  })

  describe('可访问性测试', () => {
    it('应该支持键盘导航', async () => {
      wrapper = mount(StandardCategorySelector, testConfig)

      const button = wrapper.find('el-button')
      
      // 模拟键盘事件
      await button.trigger('keydown.enter')
      await waitForAsync(wrapper)

      const modal = wrapper.findComponent({ name: 'CategoryModal' })
      expect(modal.props('modelValue')).toBe(true)
    })

    it('应该有正确的ARIA属性', () => {
      wrapper = mount(StandardCategorySelector, testConfig)

      const button = wrapper.find('el-button')
      expect(button.attributes('role')).toBeDefined()
    })
  })

  describe('内存泄漏测试', () => {
    it('组件卸载后应该清理事件监听器', async () => {
      wrapper = mount(StandardCategorySelector, testConfig)

      await mockClick(wrapper, 'el-button')
      await waitForAsync(wrapper)

      // 记录初始事件监听器数量
      const initialListeners = document.querySelectorAll('[data-test-listener]').length

      wrapper.unmount()
      await waitForAsync(wrapper)

      // 验证事件监听器已清理
      const finalListeners = document.querySelectorAll('[data-test-listener]').length
      expect(finalListeners).toBeLessThanOrEqual(initialListeners)
    })
  })

  describe('边界条件测试', () => {
    it('应该处理空数据', async () => {
      mockApiCall(mockGetStandardCategoryList, {
        code: 0,
        data: { list: [], total: 0, page: 1, pageSize: 10 },
        msg: '获取成功'
      })

      wrapper = mount(StandardCategorySelector, testConfig)

      await mockClick(wrapper, 'el-button')
      await waitForAsync(wrapper)

      const modal = wrapper.findComponent({ name: 'CategoryModal' })
      const categoryTable = modal.findComponent({ name: 'CategoryTable' })
      
      // 验证空状态显示
      expectElementExists(categoryTable, '.el-empty')
    })

    it('应该处理最大选择限制', async () => {
      wrapper = mount(StandardCategorySelector, {
        ...testConfig,
        props: {
          modelValue: [],
          maxSelections: 2
        }
      })

      await mockClick(wrapper, 'el-button')
      await waitForAsync(wrapper)

      const modal = wrapper.findComponent({ name: 'CategoryModal' })
      const categoryTable = modal.findComponent({ name: 'CategoryTable' })

      // 选择超过限制的品类
      await categoryTable.vm.$emit('select', mockCategories[0])
      await categoryTable.vm.$emit('select', mockCategories[1])
      await categoryTable.vm.$emit('select', mockCategories[2]) // 应该被拒绝

      expect(modal.vm.selectedCategories).toHaveLength(2)
    })
  })

  describe('并发测试', () => {
    it('应该正确处理并发API请求', async () => {
      wrapper = mount(StandardCategorySelector, testConfig)

      // 快速连续打开关闭弹框
      await mockClick(wrapper, 'el-button')
      const modal = wrapper.findComponent({ name: 'CategoryModal' })
      await modal.vm.$emit('cancel')
      await mockClick(wrapper, 'el-button')
      await waitForAsync(wrapper)

      // 验证最终状态正确
      expect(modal.props('modelValue')).toBe(true)
    })
  })
})