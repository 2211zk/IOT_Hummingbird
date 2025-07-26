/**
 * 快捷入口导航服务单元测试
 */

import { describe, it, expect, vi, beforeEach } from 'vitest'
import { QuickEntryNavigationService } from '@/services/QuickEntryNavigationService'

// Mock dependencies
vi.mock('@/config/quickEntryConfig', () => ({
  getQuickEntryConfig: vi.fn(),
  validateQuickEntryConfig: vi.fn()
}))

vi.mock('@/utils/quickEntryErrorHandler', () => ({
  handleQuickEntryError: vi.fn(),
  checkUserPermission: vi.fn()
}))

vi.mock('element-plus', () => ({
  ElMessage: {
    success: vi.fn(),
    error: vi.fn()
  }
}))

describe('QuickEntryNavigationService', () => {
  let service
  let mockRouter
  let mockRouterStore
  let mockGetQuickEntryConfig
  let mockValidateQuickEntryConfig
  let mockHandleQuickEntryError
  let mockCheckUserPermission

  beforeEach(() => {
    // 重置所有 mock
    vi.clearAllMocks()

    // 创建 mock 对象
    mockRouter = {
      push: vi.fn().mockResolvedValue()
    }

    mockRouterStore = {
      routeMap: {
        'WlProducts': { name: 'WlProducts', meta: {} },
        'WlEquipment': { name: 'WlEquipment', meta: {} },
        'State': { name: 'State', meta: {} }
      },
      setLeftMenu: vi.fn()
    }

    // 获取 mock 函数
    const { getQuickEntryConfig, validateQuickEntryConfig } = require('@/config/quickEntryConfig')
    const { handleQuickEntryError, checkUserPermission } = require('@/utils/quickEntryErrorHandler')
    
    mockGetQuickEntryConfig = getQuickEntryConfig
    mockValidateQuickEntryConfig = validateQuickEntryConfig
    mockHandleQuickEntryError = handleQuickEntryError
    mockCheckUserPermission = checkUserPermission

    // 创建服务实例
    service = new QuickEntryNavigationService(mockRouter, mockRouterStore)
  })

  describe('构造函数', () => {
    it('应该正确初始化服务', () => {
      expect(service.router).toBe(mockRouter)
      expect(service.routerStore).toBe(mockRouterStore)
    })
  })

  describe('validateRoute', () => {
    it('应该验证存在的路由', () => {
      const result = service.validateRoute('WlProducts')
      expect(result).toBe(true)
    })

    it('应该拒绝不存在的路由', () => {
      const result = service.validateRoute('NonExistentRoute')
      expect(result).toBe(false)
    })

    it('应该处理空参数', () => {
      expect(service.validateRoute('')).toBe(false)
      expect(service.validateRoute(null)).toBe(false)
      expect(service.validateRoute(undefined)).toBe(false)
    })
  })

  describe('checkRoutePermission', () => {
    it('应该检查路由权限', () => {
      mockCheckUserPermission.mockReturnValue(true)
      
      const result = service.checkRoutePermission('WlProducts')
      expect(result).toBe(true)
      expect(mockCheckUserPermission).toHaveBeenCalledWith('WlProducts', null)
    })

    it('应该拒绝隐藏的路由', () => {
      mockRouterStore.routeMap['HiddenRoute'] = {
        name: 'HiddenRoute',
        meta: { hidden: true }
      }
      
      const result = service.checkRoutePermission('HiddenRoute')
      expect(result).toBe(false)
    })

    it('应该拒绝不存在的路由', () => {
      const result = service.checkRoutePermission('NonExistentRoute')
      expect(result).toBe(false)
    })
  })

  describe('navigateToQuickEntry', () => {
    const mockRouteConfig = {
      name: 'WlProducts',
      label: '产品管理',
      parentMenu: 'wl_playform'
    }

    beforeEach(() => {
      mockValidateQuickEntryConfig.mockReturnValue(true)
      mockGetQuickEntryConfig.mockReturnValue(mockRouteConfig)
      mockCheckUserPermission.mockReturnValue(true)
    })

    it('应该成功导航到有效的快捷入口', async () => {
      const result = await service.navigateToQuickEntry('addProduct')
      
      expect(mockValidateQuickEntryConfig).toHaveBeenCalledWith('addProduct')
      expect(mockGetQuickEntryConfig).toHaveBeenCalledWith('addProduct')
      expect(mockRouterStore.setLeftMenu).toHaveBeenCalledWith('wl_playform')
      expect(mockRouter.push).toHaveBeenCalledWith({ name: 'WlProducts' })
      expect(result).toEqual(mockRouteConfig)
    })

    it('应该抛出错误当配置无效时', async () => {
      mockValidateQuickEntryConfig.mockReturnValue(false)
      
      await expect(service.navigateToQuickEntry('invalidEntry'))
        .rejects.toThrow('未知的快捷入口类型: invalidEntry')
    })

    it('应该抛出错误当路由不存在时', async () => {
      mockGetQuickEntryConfig.mockReturnValue({
        name: 'NonExistentRoute',
        label: '不存在的路由'
      })
      
      await expect(service.navigateToQuickEntry('addProduct'))
        .rejects.toThrow('路由不存在: NonExistentRoute')
    })

    it('应该抛出错误当权限不足时', async () => {
      mockCheckUserPermission.mockReturnValue(false)
      
      await expect(service.navigateToQuickEntry('addProduct'))
        .rejects.toThrow('权限不足: 无法访问产品管理')
    })

    it('应该处理路由跳转失败', async () => {
      const routerError = new Error('Navigation failed')
      mockRouter.push.mockRejectedValue(routerError)
      
      await expect(service.navigateToQuickEntry('addProduct'))
        .rejects.toThrow('跳转到产品管理失败: Navigation failed')
    })

    it('应该在没有父菜单时跳过菜单设置', async () => {
      mockGetQuickEntryConfig.mockReturnValue({
        name: 'WlProducts',
        label: '产品管理'
        // 没有 parentMenu
      })
      
      await service.navigateToQuickEntry('addProduct')
      
      expect(mockRouterStore.setLeftMenu).not.toHaveBeenCalled()
      expect(mockRouter.push).toHaveBeenCalledWith({ name: 'WlProducts' })
    })
  })

  describe('handleNavigationError', () => {
    it('应该调用错误处理函数', () => {
      const error = new Error('Test error')
      const options = { useNotification: true }
      
      service.handleNavigationError(error, 'addProduct', options)
      
      expect(mockHandleQuickEntryError).toHaveBeenCalledWith(error, 'addProduct', options)
    })
  })

  describe('getAvailableQuickEntries', () => {
    it('应该返回可用的快捷入口列表', () => {
      // Mock getAllQuickEntryConfigs
      const mockConfigs = {
        'addProduct': {
          name: 'WlProducts',
          label: '产品管理'
        },
        'addDevice': {
          name: 'WlEquipment', 
          label: '设备管理'
        },
        'invalidEntry': {
          name: 'NonExistentRoute',
          label: '无效入口'
        }
      }

      // Mock require
      vi.doMock('@/config/quickEntryConfig', () => ({
        getAllQuickEntryConfigs: () => mockConfigs
      }))

      mockCheckUserPermission.mockReturnValue(true)
      
      const availableEntries = service.getAvailableQuickEntries()
      
      expect(availableEntries).toHaveLength(2) // 只有存在的路由
      expect(availableEntries[0]).toMatchObject({
        type: 'addProduct',
        name: 'WlProducts',
        label: '产品管理'
      })
    })
  })

  describe('preloadQuickEntryComponent', () => {
    it('应该预加载组件', async () => {
      const mockComponent = vi.fn().mockResolvedValue({})
      mockRouterStore.routeMap['WlProducts'].component = mockComponent
      
      await service.preloadQuickEntryComponent('addProduct')
      
      expect(mockComponent).toHaveBeenCalled()
    })

    it('应该处理预加载失败', async () => {
      const mockComponent = vi.fn().mockRejectedValue(new Error('Load failed'))
      mockRouterStore.routeMap['WlProducts'].component = mockComponent
      
      // 应该不抛出错误
      await expect(service.preloadQuickEntryComponent('addProduct'))
        .resolves.toBeUndefined()
    })

    it('应该处理无效的配置', async () => {
      mockGetQuickEntryConfig.mockReturnValue(null)
      
      await expect(service.preloadQuickEntryComponent('invalidEntry'))
        .resolves.toBeUndefined()
    })
  })
})