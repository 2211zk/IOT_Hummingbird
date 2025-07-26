/**
 * 快捷入口导航服务
 * 处理首页快捷入口的路由跳转和菜单状态管理
 */

import { getQuickEntryConfig, validateQuickEntryConfig } from '@/config/quickEntryConfig'
import { handleQuickEntryError, checkUserPermission } from '@/utils/quickEntryErrorHandler'

export class QuickEntryNavigationService {
  constructor(router, routerStore) {
    this.router = router
    this.routerStore = routerStore
  }

  /**
   * 导航到快捷入口对应的页面
   * @param {string} entryType - 快捷入口类型
   * @param {Object} userStore - 用户状态存储（可选）
   * @returns {Promise<Object>} 路由配置对象
   * @throws {Error} 当配置不存在或路由无效时抛出错误
   */
  async navigateToQuickEntry(entryType, userStore = null) {
    // 验证配置是否存在
    if (!validateQuickEntryConfig(entryType)) {
      throw new Error(`未知的快捷入口类型: ${entryType}`)
    }

    const routeConfig = getQuickEntryConfig(entryType)
    
    // 检查路由是否存在
    if (!this.validateRoute(routeConfig.name)) {
      throw new Error(`路由不存在: ${routeConfig.name}`)
    }

    // 检查用户权限
    if (!this.checkRoutePermission(routeConfig.name, userStore)) {
      throw new Error(`权限不足: 无法访问${routeConfig.label}`)
    }

    try {
      // 设置顶级菜单激活状态
      if (routeConfig.parentMenu) {
        this.routerStore.setLeftMenu(routeConfig.parentMenu)
      }

      // 执行路由跳转
      await this.router.push({ name: routeConfig.name })
      
      console.log(`成功跳转到${routeConfig.label}页面`)
      return routeConfig
      
    } catch (error) {
      console.error(`路由跳转失败:`, error)
      throw new Error(`跳转到${routeConfig.label}失败: ${error.message}`)
    }
  }

  /**
   * 验证路由是否存在
   * @param {string} routeName - 路由名称
   * @returns {boolean} 路由是否存在
   */
  validateRoute(routeName) {
    return !!this.routerStore.routeMap[routeName]
  }

  /**
   * 检查用户是否有访问指定路由的权限
   * @param {string} routeName - 路由名称
   * @param {Object} userStore - 用户状态存储
   * @returns {boolean} 是否有权限
   */
  checkRoutePermission(routeName, userStore = null) {
    const route = this.routerStore.routeMap[routeName]
    if (!route) return false
    
    // 如果路由被隐藏，则认为用户无权限访问
    if (route.meta && route.meta.hidden) {
      return false
    }
    
    // 使用统一的权限检查函数
    return checkUserPermission(routeName, userStore)
  }

  /**
   * 获取所有可用的快捷入口
   * @returns {Array} 可用的快捷入口列表
   */
  getAvailableQuickEntries() {
    const { getAllQuickEntryConfigs } = require('@/config/quickEntryConfig')
    const allConfigs = getAllQuickEntryConfigs()
    
    return Object.keys(allConfigs).filter(entryType => {
      const config = allConfigs[entryType]
      return this.validateRoute(config.name) && this.checkRoutePermission(config.name)
    }).map(entryType => ({
      type: entryType,
      ...allConfigs[entryType]
    }))
  }

  /**
   * 处理导航错误
   * @param {Error} error - 错误对象
   * @param {string} entryType - 快捷入口类型
   * @param {Object} options - 处理选项
   */
  handleNavigationError(error, entryType, options = {}) {
    handleQuickEntryError(error, entryType, options)
  }

  /**
   * 预加载快捷入口对应的页面组件
   * @param {string} entryType - 快捷入口类型
   * @returns {Promise<void>}
   */
  async preloadQuickEntryComponent(entryType) {
    const routeConfig = getQuickEntryConfig(entryType)
    if (!routeConfig) return
    
    try {
      const route = this.routerStore.routeMap[routeConfig.name]
      if (route && route.component && typeof route.component === 'function') {
        await route.component()
        console.log(`预加载组件成功: ${routeConfig.label}`)
      }
    } catch (error) {
      console.warn(`预加载组件失败: ${routeConfig.label}`, error)
    }
  }
}

/**
 * 创建导航服务实例的工厂函数
 * @param {Object} router - Vue Router 实例
 * @param {Object} routerStore - Pinia Router Store 实例
 * @returns {QuickEntryNavigationService} 导航服务实例
 */
export function createQuickEntryNavigationService(router, routerStore) {
  return new QuickEntryNavigationService(router, routerStore)
}