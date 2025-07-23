import { config } from '@vue/test-utils'
import { ElMessage, ElNotification, ElMessageBox, ElLoading } from 'element-plus'

// 全局配置
config.global.stubs = {
  'el-icon': true,
  'el-loading': true
}

// 模拟Element Plus的全局方法
global.ElMessage = {
  success: jest.fn(),
  error: jest.fn(),
  warning: jest.fn(),
  info: jest.fn()
}

global.ElNotification = {
  success: jest.fn(),
  error: jest.fn(),
  warning: jest.fn(),
  info: jest.fn()
}

global.ElMessageBox = {
  confirm: jest.fn().mockResolvedValue('confirm'),
  alert: jest.fn().mockResolvedValue('confirm'),
  prompt: jest.fn().mockResolvedValue({ value: 'test' })
}

global.ElLoading = {
  service: jest.fn().mockReturnValue({
    close: jest.fn()
  })
}

// 模拟lodash-es
jest.mock('lodash-es', () => ({
  debounce: (fn) => fn
}))

// 模拟window对象
Object.defineProperty(window, 'navigator', {
  value: {
    onLine: true
  },
  writable: true
})

// 模拟console方法以避免测试输出污染
global.console = {
  ...console,
  error: jest.fn(),
  warn: jest.fn(),
  log: jest.fn()
}

// 设置测试超时
jest.setTimeout(10000)

// 清理函数
afterEach(() => {
  jest.clearAllMocks()
  jest.clearAllTimers()
})