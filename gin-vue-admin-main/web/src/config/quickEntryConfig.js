/**
 * 快捷入口路由映射配置
 * 定义首页快捷入口与对应页面路由的映射关系
 */

export const QUICK_ENTRY_ROUTE_MAP = {
  // 产品管理
  'addProduct': {
    name: 'WlProducts',
    label: '产品管理',
    menuPath: ['wl_playform', 'WlProducts'],
    parentMenu: 'wl_playform',
    icon: '📦',
    description: '管理物联网产品信息'
  },
  
  // 设备管理
  'addDevice': {
    name: 'WlEquipment', 
    label: '设备管理',
    menuPath: ['wl_playform', 'WlEquipment'],
    parentMenu: 'wl_playform',
    icon: '📱',
    description: '管理物联网设备信息'
  },
  
  // 系统监控
  'serviceMonitor': {
    name: 'State',
    label: '系统监控',
    menuPath: ['opsMonitor', 'State'],
    parentMenu: 'opsMonitor',
    icon: '🖥️',
    description: '监控系统运行状态'
  },
  
  // 引擎规则
  'ruleEngine': {
    name: 'WlEngineRules',
    label: '引擎规则',
    menuPath: ['advancedCapabilities', 'WlEngineRules'],
    parentMenu: 'advancedCapabilities',
    icon: '⚙️',
    description: '配置业务规则引擎'
  },
  
  // 告警中心
  'alarmCenter': {
    name: 'WlAlarm',
    label: '告警中心',
    menuPath: ['opsMonitor', 'WlAlarm'],
    parentMenu: 'opsMonitor',
    icon: '🔔',
    description: '查看和处理系统告警'
  },
  
  // 服务器状态
  'dataCenter': {
    name: 'State',
    label: '服务器状态',
    menuPath: ['opsMonitor', 'State'],
    parentMenu: 'opsMonitor',
    icon: '💾',
    description: '查看服务器运行状态'
  }
}

/**
 * 获取快捷入口配置
 * @param {string} entryType - 快捷入口类型
 * @returns {Object|null} 配置对象或null
 */
export function getQuickEntryConfig(entryType) {
  return QUICK_ENTRY_ROUTE_MAP[entryType] || null
}

/**
 * 获取所有快捷入口配置
 * @returns {Object} 所有配置对象
 */
export function getAllQuickEntryConfigs() {
  return QUICK_ENTRY_ROUTE_MAP
}

/**
 * 验证快捷入口配置的完整性
 * @param {string} entryType - 快捷入口类型
 * @returns {boolean} 是否有效
 */
export function validateQuickEntryConfig(entryType) {
  const config = getQuickEntryConfig(entryType)
  if (!config) return false
  
  // 检查必需字段
  const requiredFields = ['name', 'label', 'menuPath', 'parentMenu']
  return requiredFields.every(field => config[field] !== undefined && config[field] !== null)
}

/**
 * 获取快捷入口的显示信息
 * @param {string} entryType - 快捷入口类型
 * @returns {Object} 显示信息对象
 */
export function getQuickEntryDisplayInfo(entryType) {
  const config = getQuickEntryConfig(entryType)
  if (!config) return null
  
  return {
    label: config.label,
    icon: config.icon,
    description: config.description
  }
}