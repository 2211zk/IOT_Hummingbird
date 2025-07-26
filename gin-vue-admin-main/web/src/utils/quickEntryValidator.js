/**
 * 快捷入口功能验证工具
 * 用于验证快捷入口配置和功能是否正常工作
 */

import { getAllQuickEntryConfigs, validateQuickEntryConfig } from '@/config/quickEntryConfig'

/**
 * 验证所有快捷入口配置
 * @returns {Object} 验证结果
 */
export function validateAllQuickEntryConfigs() {
  const configs = getAllQuickEntryConfigs()
  const results = {
    total: 0,
    valid: 0,
    invalid: 0,
    details: []
  }

  Object.keys(configs).forEach(entryType => {
    results.total++
    const isValid = validateQuickEntryConfig(entryType)
    
    if (isValid) {
      results.valid++
    } else {
      results.invalid++
    }

    results.details.push({
      entryType,
      isValid,
      config: configs[entryType]
    })
  })

  return results
}

/**
 * 验证路由映射是否存在
 * @param {Object} routerStore - 路由存储
 * @returns {Object} 验证结果
 */
export function validateRouteMapping(routerStore) {
  const configs = getAllQuickEntryConfigs()
  const results = {
    total: 0,
    found: 0,
    missing: 0,
    details: []
  }

  Object.values(configs).forEach(config => {
    results.total++
    const routeExists = !!routerStore.routeMap[config.name]
    
    if (routeExists) {
      results.found++
    } else {
      results.missing++
    }

    results.details.push({
      routeName: config.name,
      label: config.label,
      exists: routeExists
    })
  })

  return results
}

/**
 * 打印验证报告
 * @param {Object} configResults - 配置验证结果
 * @param {Object} routeResults - 路由验证结果
 */
export function printValidationReport(configResults, routeResults) {
  console.group('🔍 快捷入口功能验证报告')
  
  console.group('📋 配置验证')
  console.log(`总计: ${configResults.total}`)
  console.log(`有效: ${configResults.valid}`)
  console.log(`无效: ${configResults.invalid}`)
  
  if (configResults.invalid > 0) {
    console.warn('无效配置详情:')
    configResults.details
      .filter(item => !item.isValid)
      .forEach(item => console.warn(`- ${item.entryType}:`, item.config))
  }
  console.groupEnd()

  console.group('🛣️ 路由验证')
  console.log(`总计: ${routeResults.total}`)
  console.log(`存在: ${routeResults.found}`)
  console.log(`缺失: ${routeResults.missing}`)
  
  if (routeResults.missing > 0) {
    console.warn('缺失路由详情:')
    routeResults.details
      .filter(item => !item.exists)
      .forEach(item => console.warn(`- ${item.routeName} (${item.label})`))
  }
  console.groupEnd()

  const overallSuccess = configResults.invalid === 0 && routeResults.missing === 0
  if (overallSuccess) {
    console.log('✅ 所有验证通过，快捷入口功能可以正常使用')
  } else {
    console.error('❌ 验证失败，请检查上述问题')
  }
  
  console.groupEnd()
  
  return overallSuccess
}

/**
 * 运行完整验证
 * @param {Object} routerStore - 路由存储
 * @returns {boolean} 验证是否通过
 */
export function runFullValidation(routerStore) {
  const configResults = validateAllQuickEntryConfigs()
  const routeResults = validateRouteMapping(routerStore)
  
  return printValidationReport(configResults, routeResults)
}