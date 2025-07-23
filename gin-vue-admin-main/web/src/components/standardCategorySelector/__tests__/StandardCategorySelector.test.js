import { describe, it, expect, beforeEach, afterEach, jest } from '@jest/globals'
import { mount } from '@vue/test-utils'
import StandardCategorySelector from '../index.vue'
import CategoryModal from '../CategoryModal.vue'
import { mockCategories, mockApiResponse } from './mockData'
import { 
  createTestConfig, 
  waitForAsync, 
  mockClick, 
  expectElementExists,
  expectElementText,
  expectEventEmitted,
  cleanupTest
} from './testUtils'

// 模拟API
jest.mock('@/api/standardCategory', () => ({
  getStandardCategoryList: jest.fn(),
  getStandardCategoryCategories: jest.fn()
}))

describe('StandardCategorySelector', () => {
  let wrapper
  const testConfig = createTestConfig()

  beforeEach(() => {
    jest.clearAllMocks()
  })

  afterEach(() => {
    if (wrapper) {
      wrapper.unmount()
    }
    cleanupTest()
  })

  describe('基础渲染', () => {
    it('应该正确渲染触发按钮', () => {
      wrapper = mount(StandardCategorySelector, testConfig)
      
      expectElementExists(wrapper, '.selector-trigger')
      expectElementExists(wrapper, 'el-button')
      expectElementText(wrapper, 'el-button', '选择标准品类')
    })

    it('应该支持自定义按钮文本', () => {
      wrapper = mount(StandardCategorySelector, {
        ...testConfig,
        props: {
          placeholder: '自定义按钮文本'
        }
      })
      
      expectElementText(wrapper, 'el-button', '自定义按钮文本')
    })

    it('应该支持禁用状态', () => {
      wrapper = mount(StandardCategorySelector, {
        ...testConfig,
        props: {
          disabled: true
        }
      })
      
      const button = wrapper.find('el-button')
      expect(button.attributes('disabled')).toBeDefined()
    })
  })

  describe('弹框交互', () => {
    it('点击按钮应该打开弹框', async () => {
      wrapper = mount(StandardCategorySelector, testConfig)
      
      await mockClick(wrapper, 'el-button')
      
      const modal = wrapper.findComponent(CategoryModal)
      expect(modal.exists()).toBe(true)
      expect(modal.props('modelValue')).toBe(true)
    })

    it('禁用状态下点击按钮不应该打开弹框', async () => {
      wrapper = mount(StandardCategorySelector, {
        ...testConfig,
        props: {
          disabled: true
        }
      })
      
      await mockClick(wrapper, 'el-button')
      
      const modal = wrapper.findComponent(CategoryModal)
      expect(modal.props('modelValue')).toBe(false)
    })
  })

  describe('选择功能', () => {
    it('应该正确显示已选择的品类', () => {
      const selectedCategories = [mockCategories[0], mockCategories[1]]
      
      wrapper = mount(StandardCategorySelector, {
        ...testConfig,
        props: {
          modelValue: selectedCategories
        }
      })
      
      expectElementExists(wrapper, '.selected-display')
      
      const tags = wrapper.findAll('el-tag')
      expect(tags).toHaveLength(2)
      expect(tags[0].text()).toContain(mockCategories[0].name)
      expect(tags[1].text()).toContain(mockCategories[1].name)
    })

    it('应该支持移除单个已选择的品类', async () => {
      const selectedCategories = [mockCategories[0]]
      
      wrapper = mount(StandardCategorySelector, {
        ...testConfig,
        props: {
          modelValue: selectedCategories
        }
      })
      
      const tag = wrapper.find('el-tag')
      await tag.find('.el-tag__close').trigger('click')
      
      expectEventEmitted(wrapper, 'update:modelValue', [])
      expectEventEmitted(wrapper, 'change', [])
    })

    it('应该限制显示的品类数量', () => {
      const selectedCategories = mockCategories.slice(0, 5)
      
      wrapper = mount(StandardCategorySelector, {
        ...testConfig,
        props: {
          modelValue: selectedCategories,
          maxDisplay: 3
        }
      })
      
      const tags = wrapper.findAll('el-tag')
      expect(tags).toHaveLength(4) // 3个品类 + 1个"+2"标签
      
      const moreTag = tags[3]
      expect(moreTag.text()).toContain('+2')
    })
  })

  describe('Props传递', () => {
    it('应该正确传递props给CategoryModal', () => {
      wrapper = mount(StandardCategorySelector, {
        ...testConfig,
        props: {
          multiple: false,
          maxSelections: 5
        }
      })
      
      const modal = wrapper.findComponent(CategoryModal)
      expect(modal.props('multiple')).toBe(false)
      expect(modal.props('maxSelections')).toBe(5)
    })
  })

  describe('事件处理', () => {
    it('应该正确处理确认事件', async () => {
      wrapper = mount(StandardCategorySelector, testConfig)
      
      const modal = wrapper.findComponent(CategoryModal)
      await modal.vm.$emit('confirm', mockCategories.slice(0, 2))
      
      expectEventEmitted(wrapper, 'update:modelValue', mockCategories.slice(0, 2))
      expectEventEmitted(wrapper, 'change', mockCategories.slice(0, 2))
      expectEventEmitted(wrapper, 'confirm', mockCategories.slice(0, 2))
    })

    it('应该正确处理取消事件', async () => {
      wrapper = mount(StandardCategorySelector, testConfig)
      
      const modal = wrapper.findComponent(CategoryModal)
      await modal.vm.$emit('cancel')
      
      // 取消时不应该触发任何更新事件
      expect(wrapper.emitted('update:modelValue')).toBeFalsy()
      expect(wrapper.emitted('change')).toBeFalsy()
    })
  })

  describe('暴露的方法', () => {
    it('应该暴露openModal方法', () => {
      wrapper = mount(StandardCategorySelector, testConfig)
      
      expect(typeof wrapper.vm.openModal).toBe('function')
      
      wrapper.vm.openModal()
      expect(wrapper.vm.modalVisible).toBe(true)
    })

    it('应该暴露clearSelection方法', () => {
      wrapper = mount(StandardCategorySelector, {
        ...testConfig,
        props: {
          modelValue: mockCategories.slice(0, 2)
        }
      })
      
      expect(typeof wrapper.vm.clearSelection).toBe('function')
      
      wrapper.vm.clearSelection()
      
      expectEventEmitted(wrapper, 'update:modelValue', [])
      expectEventEmitted(wrapper, 'change', [])
    })

    it('应该暴露getSelectedCategories方法', () => {
      const selectedCategories = mockCategories.slice(0, 2)
      
      wrapper = mount(StandardCategorySelector, {
        ...testConfig,
        props: {
          modelValue: selectedCategories
        }
      })
      
      expect(typeof wrapper.vm.getSelectedCategories).toBe('function')
      expect(wrapper.vm.getSelectedCategories()).toEqual(selectedCategories)
    })
  })

  describe('响应式更新', () => {
    it('应该响应外部modelValue的变化', async () => {
      wrapper = mount(StandardCategorySelector, {
        ...testConfig,
        props: {
          modelValue: []
        }
      })
      
      await wrapper.setProps({
        modelValue: mockCategories.slice(0, 1)
      })
      
      expectElementExists(wrapper, '.selected-display')
      const tags = wrapper.findAll('el-tag')
      expect(tags).toHaveLength(1)
      expect(tags[0].text()).toContain(mockCategories[0].name)
    })
  })
})