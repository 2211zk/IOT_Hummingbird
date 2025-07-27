import service from '@/utils/request'

// 设备相关API

/**
 * 获取设备列表
 * @param {Object} params - 查询参数
 * @param {number} params.page - 页码
 * @param {number} params.pageSize - 每页数量
 * @param {string} params.eqName - 设备名称
 * @param {string} params.productsId - 产品ID
 * @param {string} params.status - 状态
 */
export function getDeviceList(params) {
  return service({
    url: '/wlEquipment/getWlEquipmentList',
    method: 'get',
    params
  })
}

/**
 * 获取设备详情
 * @param {number} id - 设备ID
 */
export function getDeviceDetail(id) {
  return service({
    url: `/wlEquipment/findWlEquipment`,
    method: 'get',
    params: { ID: id }
  })
}

/**
 * 创建设备
 * @param {Object} data - 设备数据
 * @param {string} data.eqName - 设备名称
 * @param {string} data.productsId - 产品ID
 * @param {string} data.eqInfo - 设备描述
 */
export function createDevice(data) {
  return service({
    url: '/wlEquipment/createWlEquipment',
    method: 'post',
    data
  })
}

/**
 * 更新设备
 * @param {Object} data - 设备数据
 * @param {number} data.ID - 设备ID
 * @param {string} data.eqName - 设备名称
 * @param {string} data.productsId - 产品ID
 * @param {string} data.eqInfo - 设备描述
 */
export function updateDevice(data) {
  return service({
    url: '/wlEquipment/updateWlEquipment',
    method: 'put',
    data
  })
}

/**
 * 删除设备
 * @param {Object} data - 删除参数
 * @param {number} data.id - 设备ID
 */
export function deleteDevice(data) {
  return service({
    url: '/wlEquipment/deleteWlEquipment',
    method: 'delete',
    params: { ID: data.id }
  })
}

/**
 * 批量删除设备
 * @param {Object} data - 删除参数
 * @param {Array} data.ids - 设备ID列表
 */
export function batchDeleteDevices(data) {
  return service({
    url: '/device/batch-delete',
    method: 'delete',
    data
  })
}

/**
 * 获取设备统计信息
 */
export function getDeviceStats() {
  return service({
    url: '/device/stats',
    method: 'get'
  })
}

// 设备状态管理

/**
 * 启用设备
 * @param {Object} data - 参数
 * @param {number} data.id - 设备ID
 */
export function enableDevice(data) {
  return service({
    url: '/device/enable',
    method: 'put',
    data
  })
}

/**
 * 禁用设备
 * @param {Object} data - 参数
 * @param {number} data.id - 设备ID
 */
export function disableDevice(data) {
  return service({
    url: '/device/disable',
    method: 'put',
    data
  })
}

/**
 * 批量更新设备状态
 * @param {Object} data - 参数
 * @param {Array} data.ids - 设备ID列表
 * @param {string} data.status - 状态
 */
export function batchUpdateDeviceStatus(data) {
  return service({
    url: '/device/batch-status',
    method: 'put',
    data
  })
}

// 设备搜索和过滤

/**
 * 搜索设备
 * @param {Object} params - 搜索参数
 * @param {string} params.keyword - 关键词
 * @param {string} params.type - 搜索类型（name, product, all）
 */
export function searchDevices(params) {
  return service({
    url: '/device/search',
    method: 'get',
    params
  })
}

/**
 * 获取设备选项（用于下拉选择）
 * @param {Object} params - 查询参数
 * @param {string} params.status - 状态过滤
 * @param {boolean} params.available - 是否只返回可用设备
 */
export function getDeviceOptions(params = {}) {
  return service({
    url: '/device/options',
    method: 'get',
    params
  })
}

// 工具函数

/**
 * 格式化设备状态
 * @param {string} status - 状态值
 * @returns {Object} 格式化后的状态信息
 */
export function formatDeviceStatus(status) {
  const statusMap = {
    '启用': { text: '启用', type: 'success', color: '#67C23A' },
    '禁用': { text: '禁用', type: 'danger', color: '#F56C6C' },
    '维护': { text: '维护中', type: 'warning', color: '#E6A23C' },
    '故障': { text: '故障', type: 'danger', color: '#F56C6C' }
  }
  
  return statusMap[status] || { text: status, type: 'info', color: '#909399' }
}

/**
 * 验证设备数据
 * @param {Object} data - 设备数据
 * @returns {Object} 验证结果
 */
export function validateDeviceData(data) {
  const errors = []
  
  if (!data.deviceName || data.deviceName.trim() === '') {
    errors.push('设备名称不能为空')
  }
  
  if (data.deviceName && data.deviceName.length > 100) {
    errors.push('设备名称长度不能超过100个字符')
  }
  
  if (data.productName && data.productName.length > 100) {
    errors.push('产品名称长度不能超过100个字符')
  }
  
  return {
    valid: errors.length === 0,
    errors
  }
}