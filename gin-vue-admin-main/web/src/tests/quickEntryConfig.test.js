/**
 * 快捷入口配置单元测试
 */

import { describe, it, expect } from 'vitest'
import {
  QUICK_ENTRY_ROUTE_MAP,
  getQuickEntryConfig,
  getAllQuickEntryConfigs,
  validateQuickEntryConfig,
  getQuickEntryDisplayInfo
} from '@/config/quickEntryConfig'

describe('快捷入口配置测试', () => {
  describe('QUICK_ENTRY_ROUTE_MAP', () => {
    it('应该包含所有必需的快捷入口配置', () => {
      const expectedEntries = [
        'addProduct',
        'addDevice', 
        'serviceMonitor',
        'ruleEngine',
        'alarmCenter',
        'dataCenter'
      ]
      
      expectedEntries.forEach(entry => {
        expect(QUICK_ENTRY_ROUTE_MAP).toHaveProperty(entry)
      })
    })

    it('每个配置项应该包含必需的字段', () => {
      const requiredFields = ['name', 'label', 'menuPath', 'parentMenu', 'icon', 'description']
      
      Object.values(QUICK_ENTRY_ROUTE_MAP).forEach(config => {
        requiredFields.forEach(field => {
          expect(config).toHaveProperty(field)
          expect(config[field]).toBeDefined()
        })
      })
    })

    it('路由名称应该是有效的字符串', () => {
      Object.values(QUICK_ENTRY_ROUTE_MAP).forEach(config => {
        expect(typeof config.name).toBe('string')
        expect(config.name.length).toBeGreaterThan(0)
      })
    })

    it('菜单路径应该是数组', () => {
      Object.values(QUICK_ENTRY_ROUTE_MAP).forEach(config => {
        expect(Array.isArray(config.menuPath)).toBe(true)
        expect(config.menuPath.length).toBeGreaterThan(0)
      })
    })
  })

  describe('getQuickEntryConfig', () => {
    it('应该返回存在的配置', () => {
      const config = getQuickEntryConfig('addProduct')
      expect(config).toBeDefined()
      expect(config.name).toBe('WlProducts')
      expect(config.label).toBe('产品管理')
    })

    it('应该为不存在的配置返回null', () => {
      const config = getQuickEntryConfig('nonExistentEntry')
      expect(config).toBeNull()
    })

    it('应该为空参数返回null', () => {
      expect(getQuickEntryConfig('')).toBeNull()
      expect(getQuickEntryConfig(null)).toBeNull()
      expect(getQuickEntryConfig(undefined)).toBeNull()
    })
  })

  describe('getAllQuickEntryConfigs', () => {
    it('应该返回所有配置', () => {
      const allConfigs = getAllQuickEntryConfigs()
      expect(typeof allConfigs).toBe('object')
      expect(Object.keys(allConfigs).length).toBeGreaterThan(0)
      expect(allConfigs).toEqual(QUICK_ENTRY_ROUTE_MAP)
    })
  })

  describe('validateQuickEntryConfig', () => {
    it('应该验证有效的配置', () => {
      expect(validateQuickEntryConfig('addProduct')).toBe(true)
      expect(validateQuickEntryConfig('addDevice')).toBe(true)
    })

    it('应该拒绝无效的配置', () => {
      expect(validateQuickEntryConfig('nonExistentEntry')).toBe(false)
      expect(validateQuickEntryConfig('')).toBe(false)
      expect(validateQuickEntryConfig(null)).toBe(false)
      expect(validateQuickEntryConfig(undefined)).toBe(false)
    })
  })

  describe('getQuickEntryDisplayInfo', () => {
    it('应该返回正确的显示信息', () => {
      const displayInfo = getQuickEntryDisplayInfo('addProduct')
      expect(displayInfo).toBeDefined()
      expect(displayInfo.label).toBe('产品管理')
      expect(displayInfo.icon).toBe('📦')
      expect(displayInfo.description).toBe('管理物联网产品信息')
    })

    it('应该为不存在的配置返回null', () => {
      const displayInfo = getQuickEntryDisplayInfo('nonExistentEntry')
      expect(displayInfo).toBeNull()
    })

    it('应该只包含显示相关的字段', () => {
      const displayInfo = getQuickEntryDisplayInfo('addProduct')
      const expectedFields = ['label', 'icon', 'description']
      
      expect(Object.keys(displayInfo)).toEqual(expectedFields)
      expect(displayInfo).not.toHaveProperty('name')
      expect(displayInfo).not.toHaveProperty('menuPath')
    })
  })
})