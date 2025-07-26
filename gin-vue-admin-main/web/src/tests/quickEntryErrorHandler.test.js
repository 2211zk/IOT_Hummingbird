/**
 * 快捷入口错误处理单元测试
 */

import { describe, it, expect, vi, beforeEach } from 'vitest'
import {
  ERROR_TYPES,
  classifyError,
  showErrorMessage,
  handleQuickEntryError,
  checkUserPermission,
  createErrorHandler
} from '@/utils/quickEntryErrorHandler'

// Mock Element Plus
vi.mock('element-plus', () => ({
  ElMessage: vi.fn(),
  ElNotification: vi.fn()
}))

// Mock config
vi.mock('@/config/quickEntryConfig', () => ({
  getQuickEntryDisplayInfo: vi.fn()
}))

describe('快捷入口错误处理测试', () => {
  let mockElMessage
  let mockElNotification
  let mockGetQuickEntryDisplayInfo

  beforeEach(() => {
    vi.clearAllMocks()
    
    const { ElMessage, ElNotification } = require('element-plus')
    const { getQuickEntryDisplayInfo } = require('@/config/quickEntryConfig')
    
    mockElMessage = ElMessage
    mockElNotification = ElNotification
    mockGetQuickEntryDisplayInfo = getQuickEntryDisplayInfo
    
    mockGetQuickEntryDisplayInfo.mockReturnValue({
      label: '产品管理',
      icon: '📦',
      description: '管理物联网产品信息'
    })
  })

  describe('ERROR_TYPES', () => {
    it('应该定义所有错误类型', () => {
      const expectedTypes = [
        'CONFIG_NOT_FOUND',
        'ROUTE_NOT_FOUND', 
        'PERMISSION_DENIED',
        'NAVIGATION_FAILED',
        'NETWORK_ERROR',
        'UNKNOWN_ERROR'
      ]
      
      expectedTypes.forEach(type => {
        expect(ERROR_TYPES).toHaveProperty(type)
      })
    })
  })

  describe('classifyError', () => {
    it('应该正确分类配置不存在错误', () => {
      const error = new Error('未知的快捷入口类型: test')
      expect(classifyError(error)).toBe(ERROR_TYPES.CONFIG_NOT_FOUND)
      
      const error2 = new Error('配置不存在')
      expect(classifyError(error2)).toBe(ERROR_TYPES.CONFIG_NOT_FOUND)
    })

    it('应该正确分类路由不存在错误', () => {
      const error = new Error('路由不存在: TestRoute')
      expect(classifyError(error)).toBe(ERROR_TYPES.ROUTE_NOT_FOUND)
      
      const error2 = new Error('页面不存在')
      expect(classifyError(error2)).toBe(ERROR_TYPES.ROUTE_NOT_FOUND)
    })

    it('应该正确分类权限错误', () => {
      const error = new Error('权限不足')
      expect(classifyError(error)).toBe(ERROR_TYPES.PERMISSION_DENIED)
      
      const error2 = new Error('Permission denied')
      expect(classifyError(error2)).toBe(ERROR_TYPES.PERMISSION_DENIED)
    })

    it('应该正确分类导航错误', () => {
      const error = new Error('跳转失败')
      expect(classifyError(error)).toBe(ERROR_TYPES.NAVIGATION_FAILED)
      
      const error2 = new Error('Navigation failed')
      expect(classifyError(error2)).toBe(ERROR_TYPES.NAVIGATION_FAILED)
    })

    it('应该正确分类网络错误', () => {
      const error = new Error('网络连接失败')
      expect(classifyError(error)).toBe(ERROR_TYPES.NETWORK_ERROR)
      
      const error2 = new Error('Network error')
      expect(classifyError(error2)).toBe(ERROR_TYPES.NETWORK_ERROR)
    })

    it('应该将未知错误分类为UNKNOWN_ERROR', () => {
      const error = new Error('Some random error')
      expect(classifyError(error)).toBe(ERROR_TYPES.UNKNOWN_ERROR)
    })
  })

  describe('showErrorMessage', () => {
    it('应该使用ElMessage显示错误', () => {
      showErrorMessage(ERROR_TYPES.CONFIG_NOT_FOUND)
      
      expect(mockElMessage).toHaveBeenCalledWith({
        message: '该功能正在开发中，敬请期待',
        type: 'warning',
        duration: 3000,
        showClose: true
      })
    })

    it('应该使用ElNotification显示错误', () => {
      showErrorMessage(ERROR_TYPES.ROUTE_NOT_FOUND, { useNotification: true })
      
      expect(mockElNotification).toHaveBeenCalledWith({
        title: '页面不存在',
        message: '目标页面不存在或暂未配置，请联系管理员',
        type: 'error',
        duration: 4000,
        position: 'top-right'
      })
    })

    it('应该使用自定义消息', () => {
      const customMessage = '自定义错误消息'
      showErrorMessage(ERROR_TYPES.CONFIG_NOT_FOUND, { customMessage })
      
      expect(mockElMessage).toHaveBeenCalledWith({
        message: customMessage,
        type: 'warning',
        duration: 3000,
        showClose: true
      })
    })

    it('应该处理未知错误类型', () => {
      showErrorMessage('INVALID_ERROR_TYPE')
      
      expect(mockElMessage).toHaveBeenCalledWith({
        message: '发生未知错误，请稍后重试',
        type: 'error',
        duration: 3000,
        showClose: true
      })
    })
  })

  describe('handleQuickEntryError', () => {
    it('应该处理配置不存在错误', () => {
      const error = new Error('未知的快捷入口类型: test')
      
      handleQuickEntryError(error, 'addProduct')
      
      expect(mockElMessage).toHaveBeenCalledWith({
        message: '产品管理功能暂未开放，敬请期待',
        type: 'warning',
        duration: 3000,
        showClose: true
      })
    })

    it('应该处理路由不存在错误', () => {
      const error = new Error('路由不存在: TestRoute')
      
      handleQuickEntryError(error, 'addProduct')
      
      expect(mockElMessage).toHaveBeenCalledWith({
        message: '产品管理页面暂未配置，请联系管理员',
        type: 'error',
        duration: 3000,
        showClose: true
      })
    })

    it('应该处理权限不足错误', () => {
      const error = new Error('权限不足')
      
      handleQuickEntryError(error, 'addProduct')
      
      expect(mockElMessage).toHaveBeenCalledWith({
        message: '您没有访问产品管理的权限',
        type: 'warning',
        duration: 3000,
        showClose: true
      })
    })

    it('应该处理导航失败错误', () => {
      const error = new Error('跳转失败')
      
      handleQuickEntryError(error, 'addProduct')
      
      expect(mockElMessage).toHaveBeenCalledWith({
        message: '跳转到产品管理失败，请稍后重试',
        type: 'error',
        duration: 3000,
        showClose: true
      })
    })

    it('应该处理未知快捷入口类型', () => {
      mockGetQuickEntryDisplayInfo.mockReturnValue(null)
      const error = new Error('Some error')
      
      handleQuickEntryError(error, 'unknownEntry')
      
      expect(mockElMessage).toHaveBeenCalled()
    })
  })

  describe('checkUserPermission', () => {
    it('应该允许有权限的用户访问', () => {
      const userStore = {
        userInfo: {
          permissions: ['product:view', 'device:view']
        }
      }
      
      expect(checkUserPermission('WlProducts', userStore)).toBe(true)
    })

    it('应该拒绝无权限的用户访问', () => {
      const userStore = {
        userInfo: {
          permissions: ['device:view'] // 没有 product:view
        }
      }
      
      expect(checkUserPermission('WlProducts', userStore)).toBe(false)
    })

    it('应该处理无用户信息的情况', () => {
      expect(checkUserPermission('WlProducts', null)).toBe(false)
      expect(checkUserPermission('WlProducts', {})).toBe(false)
      expect(checkUserPermission('WlProducts', { userInfo: null })).toBe(false)
    })

    it('应该允许访问没有权限要求的路由', () => {
      const userStore = {
        userInfo: {
          permissions: []
        }
      }
      
      // 未在权限映射中的路由应该默认允许访问
      expect(checkUserPermission('UnmappedRoute', userStore)).toBe(true)
    })
  })

  describe('createErrorHandler', () => {
    it('应该创建错误处理中间件', () => {
      const errorHandler = createErrorHandler()
      expect(typeof errorHandler).toBe('function')
    })

    it('应该处理错误并记录日志', () => {
      const consoleSpy = vi.spyOn(console, 'group').mockImplementation(() => {})
      const consoleErrorSpy = vi.spyOn(console, 'error').mockImplementation(() => {})
      const consoleLogSpy = vi.spyOn(console, 'log').mockImplementation(() => {})
      const consoleGroupEndSpy = vi.spyOn(console, 'groupEnd').mockImplementation(() => {})
      
      const errorHandler = createErrorHandler({ enableLogging: true })
      const error = new Error('Test error')
      
      errorHandler(error, 'addProduct', { test: 'context' })
      
      expect(consoleSpy).toHaveBeenCalledWith('🚨 快捷入口错误 [addProduct]')
      expect(consoleErrorSpy).toHaveBeenCalledWith('错误对象:', error)
      expect(consoleLogSpy).toHaveBeenCalledWith('上下文:', { test: 'context' })
      expect(consoleGroupEndSpy).toHaveBeenCalled()
      
      consoleSpy.mockRestore()
      consoleErrorSpy.mockRestore()
      consoleLogSpy.mockRestore()
      consoleGroupEndSpy.mockRestore()
    })

    it('应该支持禁用日志记录', () => {
      const consoleSpy = vi.spyOn(console, 'group').mockImplementation(() => {})
      
      const errorHandler = createErrorHandler({ enableLogging: false })
      const error = new Error('Test error')
      
      errorHandler(error, 'addProduct')
      
      expect(consoleSpy).not.toHaveBeenCalled()
      
      consoleSpy.mockRestore()
    })
  })
})