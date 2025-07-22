import { describe, it, expect, beforeEach, afterEach, jest } from '@jest/globals'
import { mount } from '@vue/test-utils'
import CategoryModal from '../CategoryModal.vue'
import SearchBar from '../SearchBar.vue'
import CategoryTable from '../CategoryTable.vue'
import SelectedList from '../SelectedList.vue'
import ActionButtons from '../ActionButtons.vue'
import { mockCategories, mockApiResponse, createMockPageData } from './mockData'
import { 
  createTestConfig, 
  waitForAsync, 
  mockClick, 
  expectElementExists,
  expectEventEmitted,
  mockApiCall,
  mockApiError,
  cleanupTest
} from './testUtils'

// 模拟API
const mockGetStandardCategoryList = jest.fn()
jest.mock('@/api/standardCategory', () => ({
  getStandardCategoryList: mockGetStandardCategoryList
}))

describe('CategoryModal', () => {
  let wrapper
  const testConfig = createTestConfig()

  beforeEach(() => {
    jest.clearAllMocks()
    mockApiCall(mockGetStandardCategoryList, mockApiResponse.success)
  })

  afterEach(() => {
    if (wrapper) {
      wrapper.unmount()
    }
    cleanupTest()
  })

  describe('基础渲染', () => {
    it('应该正确渲染弹框结构', () => {
      wrapper = mount(CategoryModal, {
        ...testConfig,
        props: {
          modelValue: true
        }
      })
      
      expectElementExists(wrapper, '.category-modal')
      expectElementExists(wrapper, '.category-selector-content')
      expectElementExists(wrapper, '.left-panel')
      expectElementExists(wrapper, '.right-panel')
      
      expect(wrapper.findComponent(SearchBar).exists()).toBe(true)
      expect(wrapper.findComponent(CategoryTable).exists()).toBe(true)
      expect(wrapper.findComponent(SelectedList).exists()).toBe(true)
      expect(wrapper.findComponent(ActionButtons).exists()).toBe(true)
    })

    it('弹框关闭时不应该渲染内容', () => {
      wrapper = mount(CategoryModal, {
        ...testConfig,
        props: {
          modelValue: false
        }
      })
      
      const dialog = wrapper.find('.el-dialog')
      expect(dialog.isVisible()).toBe(false)
    })
  })

  describe('数据加载', () => {
    it('弹框打开时应该加载品类数据', async () => {
      wrapper = mount(CategoryModal, {
        ...testConfig,
        props: {
          modelValue: false
        }
      })
      
      await wrapper.setProps({ modelValue: true })
      await waitForAsync(wrapper)
      
      expect(mockGetStandardCategoryList).toHaveBeenCalledWith({
        page: 1,
        pageSize: 10,
        keyword: '',
        category: '',
        status: null
      })
    })

    it('应该正确处理API加载错误', async () => {
      mockApiError(mockGetStandardCategoryList, new Error('Network Error'))
      
      wrapper = mount(CategoryModal, {
        ...testConfig,
        props: {
          modelValue: true
        }
      })
      
      await waitForAsync(wrapper)
      
      // 验证错误处理
      expect(wrapper.vm.categoryList).toEqual([])
      expect(wrapper.vm.total).toBe(0)
    })
  })

  describe('搜索功能', () => {
    it('应该响应搜索事件', async () => {
      wrapper = mount(CategoryModal, {
        ...testConfig,
        props: {
          modelValue: true
        }
      })
      
      const searchBar = wrapper.findComponent(SearchBar)
      const searchData = {
        keyword: '电子',
        category: '电子设备',
        status: 1
      }
      
      await searchBar.vm.$emit('search', searchData)
      await waitForAsync(wrapper)
      
      expect(mockGetStandardCategoryList).toHaveBeenCalledWith({
        page: 1,
        pageSize: 10,
        ...searchData
      })
    })
  })

  describe('分页功能', () => {
    it('应该响应分页变化事件', async () => {
      wrapper = mount(CategoryModal, {
        ...testConfig,
        props: {
          modelValue: true
        }
      })
      
      const categoryTable = wrapper.findComponent(CategoryTable)
      
      await categoryTable.vm.$emit('page-change', { page: 2, pageSize: 20 })
      await waitForAsync(wrapper)
      
      expect(mockGetStandardCategoryList).toHaveBeenCalledWith({
        page: 2,
        pageSize: 20,
        keyword: '',
        category: '',
        status: null
      })
    })
  })

  describe('选择功能', () => {
    beforeEach(() => {
      wrapper = mount(CategoryModal, {
        ...testConfig,
        props: {
          modelValue: true,
          multiple: true
        }
      })
    })

    it('应该支持单个品类选择', async () => {
      const categoryTable = wrapper.findComponent(CategoryTable)
      const selectedCategory = mockCategories[0]
      
      await categoryTable.vm.$emit('select', selectedCategory)
      
      expect(wrapper.vm.selectedCategories).toContain(selectedCategory)
    })

    it('应该防止重复选择同一品类', async () => {
      const categoryTable = wrapper.findComponent(CategoryTable)
      const selectedCategory = mockCategories[0]
      
      // 第一次选择
      await categoryTable.vm.$emit('select', selectedCategory)
      expect(wrapper.vm.selectedCategories).toHaveLength(1)
      
      // 第二次选择同一品类
      await categoryTable.vm.$emit('select', selectedCategory)
      expect(wrapper.vm.selectedCategories).toHaveLength(1)
    })

    it('应该支持最大选择数量限制', async () => {
      await wrapper.setProps({ maxSelections: 2 })
      
      const categoryTable = wrapper.findComponent(CategoryTable)
      
      // 选择两个品类
      await categoryTable.vm.$emit('select', mockCategories[0])
      await categoryTable.vm.$emit('select', mockCategories[1])
      expect(wrapper.vm.selectedCategories).toHaveLength(2)
      
      // 尝试选择第三个品类
      await categoryTable.vm.$emit('select', mockCategories[2])
      expect(wrapper.vm.selectedCategories).toHaveLength(2)
    })

    it('单选模式下应该只能选择一个品类', async () => {
      await wrapper.setProps({ multiple: false })
      
      const categoryTable = wrapper.findComponent(CategoryTable)
      
      await categoryTable.vm.$emit('select', mockCategories[0])
      expect(wrapper.vm.selectedCategories).toHaveLength(1)
      
      await categoryTable.vm.$emit('select', mockCategories[1])
      expect(wrapper.vm.selectedCategories).toHaveLength(1)
      expect(wrapper.vm.selectedCategories[0]).toEqual(mockCategories[1])
    })
  })

  describe('移除功能', () => {
    it('应该支持移除单个品类', async () => {
      wrapper = mount(CategoryModal, {
        ...testConfig,
        props: {
          modelValue: true,
          initialSelected: [mockCategories[0], mockCategories[1]]
        }
      })
      
      const selectedList = wrapper.findComponent(SelectedList)
      
      await selectedList.vm.$emit('remove', mockCategories[0])
      
      expect(wrapper.vm.selectedCategories).toHaveLength(1)
      expect(wrapper.vm.selectedCategories[0]).toEqual(mockCategories[1])
    })

    it('应该支持清空所有选择', async () => {
      wrapper = mount(CategoryModal, {
        ...testConfig,
        props: {
          modelValue: true,
          initialSelected: mockCategories.slice(0, 3)
        }
      })
      
      const selectedList = wrapper.findComponent(SelectedList)
      
      await selectedList.vm.$emit('clear-all')
      
      expect(wrapper.vm.selectedCategories).toHaveLength(0)
    })
  })

  describe('确认和取消', () => {
    it('应该正确处理确认操作', async () => {
      wrapper = mount(CategoryModal, {
        ...testConfig,
        props: {
          modelValue: true,
          initialSelected: [mockCategories[0]]
        }
      })
      
      const actionButtons = wrapper.findComponent(ActionButtons)
      
      await actionButtons.vm.$emit('confirm')
      
      expectEventEmitted(wrapper, 'confirm', [mockCategories[0]])
      expect(wrapper.vm.visible).toBe(false)
    })

    it('没有选择品类时不应该允许确认', async () => {
      wrapper = mount(CategoryModal, {
        ...testConfig,
        props: {
          modelValue: true,
          initialSelected: []
        }
      })
      
      const actionButtons = wrapper.findComponent(ActionButtons)
      
      await actionButtons.vm.$emit('confirm')
      
      // 应该显示警告，不触发确认事件
      expect(wrapper.emitted('confirm')).toBeFalsy()
      expect(wrapper.vm.visible).toBe(true)
    })

    it('应该正确处理取消操作', async () => {
      wrapper = mount(CategoryModal, {
        ...testConfig,
        props: {
          modelValue: true
        }
      })
      
      const actionButtons = wrapper.findComponent(ActionButtons)
      
      await actionButtons.vm.$emit('cancel')
      
      expectEventEmitted(wrapper, 'cancel')
      expect(wrapper.vm.visible).toBe(false)
    })
  })

  describe('初始化数据', () => {
    it('应该正确初始化已选择的品类', () => {
      const initialSelected = mockCategories.slice(0, 2)
      
      wrapper = mount(CategoryModal, {
        ...testConfig,
        props: {
          modelValue: true,
          initialSelected
        }
      })
      
      expect(wrapper.vm.selectedCategories).toEqual(initialSelected)
    })
  })

  describe('响应式设计', () => {
    it('应该在小屏幕上调整布局', () => {
      // 模拟小屏幕
      Object.defineProperty(window, 'innerWidth', {
        writable: true,
        configurable: true,
        value: 768
      })
      
      wrapper = mount(CategoryModal, {
        ...testConfig,
        props: {
          modelValue: true
        }
      })
      
      // 验证响应式样式类是否正确应用
      expectElementExists(wrapper, '.category-selector-content')
    })
  })
})