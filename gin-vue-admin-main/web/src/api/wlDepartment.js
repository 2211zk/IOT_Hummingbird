import service from '@/utils/request'
import { withLoading, handleApiError } from '@/utils/loadingManager'
import { withRetry, RETRY_CONFIGS } from '@/utils/retryManager'

// 部门基础CRUD操作

/**
 * 获取部门列表
 * @param {Object} params - 查询参数
 * @param {number} params.page - 页码
 * @param {number} params.pageSize - 每页数量
 * @param {string} params.name - 部门名称
 * @param {string} params.status - 状态
 * @param {boolean} params.treeMode - 是否返回树形结构
 */
export const getDepartmentList = withRetry(
  'DEPARTMENT_LIST',
  withLoading(async (params) => {
    try {
      return await service({
        url: '/department/list',
        method: 'get',
        params
      })
    } catch (error) {
      throw handleApiError(error, { 
        showType: 'message',
        report: true 
      })
    }
  }, 'departmentList'),
  RETRY_CONFIGS.DEPARTMENT_LIST
)

/**
 * 获取部门树（用于选择上级部门）
 * @param {Object} params - 查询参数
 * @param {number} params.excludeId - 排除的部门ID
 */
export const getDepartmentTree = withRetry(
  'DEPARTMENT_TREE',
  withLoading(async (params = {}) => {
    try {
      return await service({
        url: '/department/tree',
        method: 'get',
        params
      })
    } catch (error) {
      throw handleApiError(error, { 
        showType: 'message',
        report: true 
      })
    }
  }, 'departmentTree'),
  RETRY_CONFIGS.DEPARTMENT_TREE
)

/**
 * 获取部门详情
 * @param {number} id - 部门ID
 */
export function getDepartmentDetail(id) {
  return service({
    url: `/department/${id}`,
    method: 'get'
  })
}

/**
 * 创建部门
 * @param {Object} data - 部门数据
 * @param {number} data.parentId - 上级部门ID
 * @param {string} data.name - 部门名称
 * @param {string} data.leader - 负责人
 * @param {string} data.phone - 电话
 * @param {string} data.email - 邮箱
 * @param {string} data.status - 状态
 * @param {number} data.sort - 排序
 * @param {Array} data.deviceIds - 关联设备ID列表
 */
export const createDepartment = withRetry(
  'DEPARTMENT_CREATE',
  withLoading(async (data) => {
    try {
      return await service({
        url: '/department/create',
        method: 'post',
        data
      })
    } catch (error) {
      throw handleApiError(error, { 
        showType: 'message',
        report: true 
      })
    }
  }, 'departmentCreate', { fullscreen: true, text: '正在创建部门...' }),
  RETRY_CONFIGS.DEPARTMENT_CREATE
)

/**
 * 更新部门
 * @param {Object} data - 部门数据
 * @param {number} data.id - 部门ID
 * @param {number} data.parentId - 上级部门ID
 * @param {string} data.name - 部门名称
 * @param {string} data.leader - 负责人
 * @param {string} data.phone - 电话
 * @param {string} data.email - 邮箱
 * @param {string} data.status - 状态
 * @param {number} data.sort - 排序
 * @param {Array} data.deviceIds - 关联设备ID列表
 */
export const updateDepartment = withRetry(
  'DEPARTMENT_UPDATE',
  withLoading(async (data) => {
    try {
      return await service({
        url: '/department/update',
        method: 'put',
        data
      })
    } catch (error) {
      throw handleApiError(error, { 
        showType: 'message',
        report: true 
      })
    }
  }, 'departmentUpdate', { fullscreen: true, text: '正在更新部门...' }),
  RETRY_CONFIGS.DEPARTMENT_UPDATE
)

/**
 * 删除部门
 * @param {Object} data - 删除参数
 * @param {number} data.id - 部门ID
 */
export const deleteDepartment = withRetry(
  'DEPARTMENT_DELETE',
  withLoading(async (data) => {
    try {
      return await service({
        url: '/department/delete',
        method: 'delete',
        data
      })
    } catch (error) {
      throw handleApiError(error, { 
        showType: 'message',
        report: true 
      })
    }
  }, 'departmentDelete', { fullscreen: true, text: '正在删除部门...' }),
  RETRY_CONFIGS.DEPARTMENT_DELETE
)

// 设备关联相关接口

/**
 * 获取可关联的设备列表
 * @param {Object} params - 查询参数
 * @param {number} params.page - 页码
 * @param {number} params.pageSize - 每页数量
 * @param {string} params.deviceName - 设备名称
 * @param {string} params.productName - 产品名称
 * @param {number} params.departmentId - 排除已关联此部门的设备
 */
export function getAvailableDevices(params) {
  return service({
    url: '/department/devices/available',
    method: 'get',
    params
  })
}

/**
 * 获取部门已关联的设备
 * @param {Object} params - 查询参数
 * @param {number} params.departmentId - 部门ID
 * @param {number} params.page - 页码
 * @param {number} params.pageSize - 每页数量
 */
export function getDepartmentDevices(params) {
  return service({
    url: '/department/devices',
    method: 'get',
    params
  })
}

// 兼容旧接口（保持向后兼容）

/**
 * 获取部门列表（兼容旧接口）
 * @param {Object} data - 查询参数
 */
export function getWlDepartmentList(data) {
  return service({
    url: '/department/getWlDepartmentList',
    method: 'post',
    data
  })
}

/**
 * 添加部门（兼容旧接口）
 * @param {Object} data - 部门数据
 */
export function addWlDepartment(data) {
  return service({
    url: '/department/createWlDepartment',
    method: 'post',
    data
  })
}

/**
 * 更新部门（兼容旧接口）
 * @param {Object} data - 部门数据
 */
export function updateWlDepartment(data) {
  return service({
    url: '/department/updateWlDepartment',
    method: 'post',
    data
  })
}

/**
 * 删除部门（兼容旧接口）
 * @param {Object} data - 删除参数
 */
export function deleteWlDepartment(data) {
  return service({
    url: '/department/deleteWlDepartment',
    method: 'post',
    data
  })
}

// 工具函数

/**
 * 构建部门树形结构
 * @param {Array} departments - 部门列表
 * @param {number|null} parentId - 父部门ID
 * @returns {Array} 树形结构数据
 */
export function buildDepartmentTree(departments, parentId = null) {
  const result = []
  
  departments.forEach(dept => {
    if (dept.parentId === parentId) {
      const children = buildDepartmentTree(departments, dept.id)
      if (children.length > 0) {
        dept.children = children
      }
      result.push(dept)
    }
  })
  
  return result
}

/**
 * 扁平化部门树
 * @param {Array} tree - 树形结构数据
 * @returns {Array} 扁平化数据
 */
export function flattenDepartmentTree(tree) {
  const result = []
  
  function traverse(nodes, level = 0) {
    nodes.forEach(node => {
      result.push({
        ...node,
        level
      })
      if (node.children && node.children.length > 0) {
        traverse(node.children, level + 1)
      }
    })
  }
  
  traverse(tree)
  return result
}

/**
 * 查找部门路径
 * @param {Array} departments - 部门列表
 * @param {number} targetId - 目标部门ID
 * @returns {Array} 部门路径
 */
export function findDepartmentPath(departments, targetId) {
  const path = []
  
  function findPath(deptId) {
    const dept = departments.find(d => d.id === deptId)
    if (!dept) return false
    
    path.unshift(dept)
    
    if (dept.parentId) {
      return findPath(dept.parentId)
    }
    
    return true
  }
  
  findPath(targetId)
  return path
} 