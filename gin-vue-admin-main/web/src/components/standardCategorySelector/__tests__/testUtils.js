import { mount, createLocalVue } from '@vue/test-utils'
import { ElButton, ElDialog, ElTable, ElInput, ElSelect, ElTag, ElPagination } from 'element-plus'
import { nextTick } from 'vue'

// 创建测试用的Vue实例
export const createTestVue = () => {
  const localVue = createLocalVue()
  
  // 注册Element Plus组件
  localVue.component('ElButton', ElButton)
  localVue.component('ElDialog', ElDialog)
  localVue.component('ElTable', ElTable)
  localVue.component('ElInput', ElInput)
  localVue.component('ElSelect', ElSelect)
  localVue.component('ElTag', ElTag)
  localVue.component('ElPagination', ElPagination)
  
  return localVue
}

// 创建组件挂载器
export const createWrapper = (component, options = {}) => {
  const localVue = createTestVue()
  
  return mount(component, {
    localVue,
    ...options
  })
}

// 等待异步操作完成
export const waitForAsync = async (wrapper, timeout = 1000) => {
  await nextTick()
  await new Promise(resolve => setTimeout(resolve, 100))
  await wrapper.vm.$nextTick()
}

// 模拟用户点击
export const mockClick = async (wrapper, selector) => {
  const element = wrapper.find(selector)
  if (!element.exists()) {
    throw new Error(`Element with selector "${selector}" not found`)
  }
  
  await element.trigger('click')
  await waitForAsync(wrapper)
}

// 模拟用户输入
export const mockInput = async (wrapper, selector, value) => {
  const input = wrapper.find(selector)
  if (!input.exists()) {
    throw new Error(`Input with selector "${selector}" not found`)
  }
  
  await input.setValue(value)
  await input.trigger('input')
  await waitForAsync(wrapper)
}

// 模拟用户选择
export const mockSelect = async (wrapper, selector, value) => {
  const select = wrapper.find(selector)
  if (!select.exists()) {
    throw new Error(`Select with selector "${selector}" not found`)
  }
  
  await select.setValue(value)
  await select.trigger('change')
  await waitForAsync(wrapper)
}

// 检查元素是否存在
export const expectElementExists = (wrapper, selector) => {
  const element = wrapper.find(selector)
  expect(element.exists()).toBe(true)
  return element
}

// 检查元素是否不存在
export const expectElementNotExists = (wrapper, selector) => {
  const element = wrapper.find(selector)
  expect(element.exists()).toBe(false)
}

// 检查元素文本内容
export const expectElementText = (wrapper, selector, expectedText) => {
  const element = expectElementExists(wrapper, selector)
  expect(element.text()).toContain(expectedText)
}

// 检查元素是否禁用
export const expectElementDisabled = (wrapper, selector, disabled = true) => {
  const element = expectElementExists(wrapper, selector)
  expect(element.attributes('disabled')).toBe(disabled ? '' : undefined)
}

// 检查元素是否可见
export const expectElementVisible = (wrapper, selector, visible = true) => {
  const element = expectElementExists(wrapper, selector)
  expect(element.isVisible()).toBe(visible)
}

// 模拟API调用
export const mockApiCall = (mockFn, response, delay = 100) => {
  return mockFn.mockImplementation(() => 
    new Promise(resolve => 
      setTimeout(() => resolve(response), delay)
    )
  )
}

// 模拟API错误
export const mockApiError = (mockFn, error, delay = 100) => {
  return mockFn.mockImplementation(() => 
    new Promise((resolve, reject) => 
      setTimeout(() => reject(error), delay)
    )
  )
}

// 检查事件是否被触发
export const expectEventEmitted = (wrapper, eventName, expectedData = undefined) => {
  const emittedEvents = wrapper.emitted(eventName)
  expect(emittedEvents).toBeTruthy()
  
  if (expectedData !== undefined) {
    const lastEvent = emittedEvents[emittedEvents.length - 1]
    expect(lastEvent[0]).toEqual(expectedData)
  }
}

// 检查组件props
export const expectProps = (wrapper, expectedProps) => {
  Object.keys(expectedProps).forEach(key => {
    expect(wrapper.props(key)).toEqual(expectedProps[key])
  })
}

// 创建测试用的全局配置
export const createTestConfig = () => ({
  global: {
    stubs: {
      'el-icon': true,
      'el-loading': true,
      'el-message': true,
      'el-notification': true
    },
    mocks: {
      $t: (key) => key,
      $message: {
        success: jest.fn(),
        error: jest.fn(),
        warning: jest.fn(),
        info: jest.fn()
      },
      $notification: {
        success: jest.fn(),
        error: jest.fn(),
        warning: jest.fn(),
        info: jest.fn()
      }
    }
  }
})

// 清理测试环境
export const cleanupTest = () => {
  jest.clearAllMocks()
  jest.clearAllTimers()
}