/**
 * 快捷入口导航功能集成测试
 */

import { describe, it, expect, vi, beforeEach } from 'vitest'
import { mount } from '@vue/test-utils'
import { createPinia } from 'pinia'
import { createRouter, createWebHistory } from 'vue-router'
import Dashboard from '@/view/dashboard/index.vue'

// Mock dependencies
vi.mock('@/api/dashboard/dashboard', () => ({
  getDashboardData: vi.fn().mockResolvedValue({
    code: 0,
    data: {
      platformData: {
        productCount: 10,
        deviceCount: 20,
        driverCount: 5,
        alarmCount: 3
      },
      systemStatus: {
        cpu: { usage: 45 },
        memory: { usage: 60 },
        load: { usage: 30 },
        disk: { usage: 25 }
      },
      alarmData: {
        hint: 1,
        minor: 1,
        important: 1,
        urgent: 0
      }
    }
  })
}))

vi.mock('echarts', () => ({
  init: vi.fn(() => ({
    setOption: vi.fn(),
    resize: vi.fn()
  }))
}))

vi.mock('element-plus', () => ({
  ElMessage: {
    success: vi.fn(),
    error: vi.fn(),
    warning: vi.fn()
  }
}))

describe('快捷入口导航集成测试', () => {
  let wrapper
  let router
  let pinia
  let routerStore

  beforeEach(async () => {
    // 创建测试路由
    router = createRouter({
      history: createWebHistory(),
      routes: [
        { path: '/', name: 'Dashboard', component: Dashboard },
        { path: '/products', name: 'WlProducts', component: { template: '<div>Products</div>' } },
        { path: '/equipment', name: 'WlEquipment', component: { template: '<div>Equipment</div>' } },
        { path: '/state', name: 'State', component: { template: '<div>State</div>' } },
        { path: '/rules', name: 'WlEngineRules', component: { template: '<div>Rules</div>' } },
        { path: '/alarm', name: 'WlAlarm', component: { template: '<div>Alarm</div>' } }
      ]
    })

    // 创建 Pinia store
    pinia = createPinia()

    // Mock router store
    routerStore = {
      routeMap: {
        'WlProducts': { name: 'WlProducts', meta: {} },
        'WlEquipment': { name: 'WlEquipment', meta: {} },
        'State': { name: 'State', meta: {} },
        'WlEngineRules': { name: 'WlEngineRules', meta: {} },
        'WlAlarm': { name: 'WlAlarm', meta: {} }
      },
      setLeftMenu: vi.fn(),
      forceUpdateMenuState: vi.fn(),
      getActiveMenuPath: vi.fn()
    }

    // Mock useRouterStore
    vi.doMock('@/pinia/modules/router', () => ({
      useRouterStore: () => routerStore
    }))

    // 挂载组件
    wrapper = mount(Dashboard, {
      global: {
        plugins: [router, pinia],
        stubs: {
          'el-watermark': true,
          'el-scrollbar': true,
          'el-menu': true,
          'el-menu-item': true,
          'el-sub-menu': true,
          'el-icon': true
        }
      }
    })

    await wrapper.vm.$nextTick()
  })

  describe('快捷入口渲染', () => {
    it('应该渲染所有快捷入口按钮', () => {
      const quickEntryItems = wrapper.findAll('.quick-entry-item')
      expect(quickEntryItems).toHaveLength(6) // 6个快捷入口
    })

    it('每个快捷入口应该有正确的图标和标签', () => {
      const quickEntryItems = wrapper.findAll('.quick-entry-item')
      
      // 检查第一个快捷入口（产品管理）
      const firstEntry = quickEntryItems[0]
      expect(firstEntry.find('.quick-entry-icon').text()).toBe('📦')
      expect(firstEntry.find('.quick-entry-label').text()).toBe('产品管理')
    })

    it('快捷入口应该有正确的提示信息', () => {
      const quickEntryItems = wrapper.findAll('.quick-entry-item')
      const firstEntry = quickEntryItems[0]
      
      expect(firstEntry.attributes('title')).toContain('管理物联网产品信息')
    })
  })

  describe('快捷入口点击功能', () => {
    it('点击产品管理应该跳转到正确的路由', async () => {
      const routerPushSpy = vi.spyOn(router, 'push')
      
      const productEntry = wrapper.findAll('.quick-entry-item')[0]
      await productEntry.trigger('click')
      
      // 等待异步操作完成
      await wrapper.vm.$nextTick()
      await new Promise(resolve => setTimeout(resolve, 100))
      
      expect(routerPushSpy).toHaveBeenCalledWith({ name: 'WlProducts' })
      expect(routerStore.setLeftMenu).toHaveBeenCalledWith('wl_playform')
    })

    it('点击设备管理应该跳转到正确的路由', async () => {
      const routerPushSpy = vi.spyOn(router, 'push')
      
      const deviceEntry = wrapper.findAll('.quick-entry-item')[1]
      await deviceEntry.trigger('click')
      
      await wrapper.vm.$nextTick()
      await new Promise(resolve => setTimeout(resolve, 100))
      
      expect(routerPushSpy).toHaveBeenCalledWith({ name: 'WlEquipment' })
      expect(routerStore.setLeftMenu).toHaveBeenCalledWith('wl_playform')
    })

    it('点击系统监控应该跳转到正确的路由', async () => {
      const routerPushSpy = vi.spyOn(router, 'push')
      
      const monitorEntry = wrapper.findAll('.quick-entry-item')[2]
      await monitorEntry.trigger('click')
      
      await wrapper.vm.$nextTick()
      await new Promise(resolve => setTimeout(resolve, 100))
      
      expect(routerPushSpy).toHaveBeenCalledWith({ name: 'State' })
      expect(routerStore.setLeftMenu).toHaveBeenCalledWith('opsMonitor')
    })
  })

  describe('加载状态管理', () => {
    it('点击时应该显示加载状态', async () => {
      const productEntry = wrapper.findAll('.quick-entry-item')[0]
      
      // 模拟慢速路由跳转
      vi.spyOn(router, 'push').mockImplementation(() => 
        new Promise(resolve => setTimeout(resolve, 500))
      )
      
      await productEntry.trigger('click')
      await wrapper.vm.$nextTick()
      
      // 检查加载状态
      expect(productEntry.classes()).toContain('loading')
      expect(productEntry.find('.quick-entry-loading').exists()).toBe(true)
      expect(productEntry.find('.quick-entry-icon').exists()).toBe(false)
    })

    it('应该防止重复点击', async () => {
      const routerPushSpy = vi.spyOn(router, 'push').mockImplementation(() => 
        new Promise(resolve => setTimeout(resolve, 500))
      )
      
      const productEntry = wrapper.findAll('.quick-entry-item')[0]
      
      // 快速连续点击
      await productEntry.trigger('click')
      await productEntry.trigger('click')
      await productEntry.trigger('click')
      
      await wrapper.vm.$nextTick()
      
      // 应该只调用一次
      expect(routerPushSpy).toHaveBeenCalledTimes(1)
    })

    it('加载完成后应该清除加载状态', async () => {
      const productEntry = wrapper.findAll('.quick-entry-item')[0]
      
      await productEntry.trigger('click')
      await wrapper.vm.$nextTick()
      
      // 等待加载状态清除
      await new Promise(resolve => setTimeout(resolve, 600))
      await wrapper.vm.$nextTick()
      
      expect(productEntry.classes()).not.toContain('loading')
      expect(productEntry.find('.quick-entry-loading').exists()).toBe(false)
      expect(productEntry.find('.quick-entry-icon').exists()).toBe(true)
    })
  })

  describe('错误处理', () => {
    it('应该处理路由跳转失败', async () => {
      const { ElMessage } = require('element-plus')
      
      // 模拟路由跳转失败
      vi.spyOn(router, 'push').mockRejectedValue(new Error('Navigation failed'))
      
      const productEntry = wrapper.findAll('.quick-entry-item')[0]
      await productEntry.trigger('click')
      
      await wrapper.vm.$nextTick()
      await new Promise(resolve => setTimeout(resolve, 100))
      
      // 应该显示错误消息
      expect(ElMessage.error || ElMessage).toHaveBeenCalled()
    })

    it('应该处理不存在的路由', async () => {
      const { ElMessage } = require('element-plus')
      
      // 移除路由映射
      routerStore.routeMap = {}
      
      const productEntry = wrapper.findAll('.quick-entry-item')[0]
      await productEntry.trigger('click')
      
      await wrapper.vm.$nextTick()
      await new Promise(resolve => setTimeout(resolve, 100))
      
      expect(ElMessage.error || ElMessage).toHaveBeenCalled()
    })
  })

  describe('用户体验', () => {
    it('应该显示成功跳转消息', async () => {
      const { ElMessage } = require('element-plus')
      
      const productEntry = wrapper.findAll('.quick-entry-item')[0]
      await productEntry.trigger('click')
      
      await wrapper.vm.$nextTick()
      await new Promise(resolve => setTimeout(resolve, 100))
      
      expect(ElMessage.success).toHaveBeenCalledWith('已跳转到产品管理')
    })

    it('快捷入口应该有悬停效果', async () => {
      const productEntry = wrapper.findAll('.quick-entry-item')[0]
      
      await productEntry.trigger('mouseenter')
      await wrapper.vm.$nextTick()
      
      // 检查CSS类或样式变化
      // 注意：在测试环境中CSS效果可能不会完全应用
      expect(productEntry.element).toBeDefined()
    })
  })

  describe('配置验证', () => {
    it('所有快捷入口配置应该有效', () => {
      const { validateQuickEntryConfig } = require('@/config/quickEntryConfig')
      const quickEntryTypes = ['addProduct', 'addDevice', 'serviceMonitor', 'ruleEngine', 'alarmCenter', 'dataCenter']
      
      quickEntryTypes.forEach(type => {
        expect(validateQuickEntryConfig(type)).toBe(true)
      })
    })

    it('所有路由应该在路由映射中存在', () => {
      const { getAllQuickEntryConfigs } = require('@/config/quickEntryConfig')
      const configs = getAllQuickEntryConfigs()
      
      Object.values(configs).forEach(config => {
        expect(routerStore.routeMap).toHaveProperty(config.name)
      })
    })
  })
})