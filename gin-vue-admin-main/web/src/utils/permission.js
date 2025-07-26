/**
 * 权限管理工具
 * 用于检查用户是否有访问特定页面或执行特定操作的权限
 */

// 权限常量定义
export const PERMISSIONS = {
  // 部门管理权限
  DEPARTMENT_VIEW: 'department:view',
  DEPARTMENT_CREATE: 'department:create',
  DEPARTMENT_EDIT: 'department:edit',
  DEPARTMENT_DELETE: 'department:delete',
  DEPARTMENT_STATUS: 'department:status',
  DEPARTMENT_SORT: 'department:sort',
  DEPARTMENT_DEVICE: 'department:device',
  DEPARTMENT_BATCH: 'department:batch',
  
  // 设备管理权限
  DEVICE_VIEW: 'device:view',
  DEVICE_CREATE: 'device:create',
  DEVICE_EDIT: 'device:edit',
  DEVICE_DELETE: 'device:delete',
  DEVICE_STATUS: 'device:status',
  DEVICE_BATCH: 'device:batch',
  
  // 系统管理权限
  SYSTEM_MANAGE: 'system:manage',
  USER_MANAGE: 'user:manage',
  ROLE_MANAGE: 'role:manage'
}

// 角色权限映射
export const ROLE_PERMISSIONS = {
  super_admin: [
    // 拥有所有权限
    ...Object.values(PERMISSIONS)
  ],
  admin: [
    // 管理员权限（除了一些高级系统权限）
    PERMISSIONS.DEPARTMENT_VIEW,
    PERMISSIONS.DEPARTMENT_CREATE,
    PERMISSIONS.DEPARTMENT_EDIT,
    PERMISSIONS.DEPARTMENT_DELETE,
    PERMISSIONS.DEPARTMENT_STATUS,
    PERMISSIONS.DEPARTMENT_SORT,
    PERMISSIONS.DEPARTMENT_DEVICE,
    PERMISSIONS.DEVICE_VIEW,
    PERMISSIONS.DEVICE_CREATE,
    PERMISSIONS.DEVICE_EDIT,
    PERMISSIONS.DEVICE_DELETE,
    PERMISSIONS.DEVICE_STATUS,
    PERMISSIONS.USER_MANAGE
  ],
  user: [
    // 普通用户权限（只读）
    PERMISSIONS.DEPARTMENT_VIEW,
    PERMISSIONS.DEVICE_VIEW
  ]
}

/**
 * 获取当前用户角色
 * @returns {string} 用户角色
 */
export function getCurrentUserRole() {
  // 这里应该从用户状态管理或localStorage获取
  // 暂时返回默认角色用于演示
  return localStorage.getItem('userRole') || 'admin'
}

/**
 * 获取当前用户权限列表
 * @returns {Array} 权限列表
 */
export function getCurrentUserPermissions() {
  const role = getCurrentUserRole()
  return ROLE_PERMISSIONS[role] || []
}

/**
 * 检查用户是否有指定权限
 * @param {string} permission 权限标识
 * @returns {boolean} 是否有权限
 */
export function hasPermission(permission) {
  const userPermissions = getCurrentUserPermissions()
  return userPermissions.includes(permission)
}

/**
 * 检查用户是否有任意一个权限
 * @param {Array} permissions 权限列表
 * @returns {boolean} 是否有权限
 */
export function hasAnyPermission(permissions) {
  return permissions.some(permission => hasPermission(permission))
}

/**
 * 检查用户是否有所有权限
 * @param {Array} permissions 权限列表
 * @returns {boolean} 是否有权限
 */
export function hasAllPermissions(permissions) {
  return permissions.every(permission => hasPermission(permission))
}

/**
 * 权限检查装饰器
 * @param {string} permission 权限标识
 * @param {Function} callback 有权限时执行的回调
 * @param {Function} fallback 无权限时执行的回调
 */
export function withPermission(permission, callback, fallback) {
  if (hasPermission(permission)) {
    return callback()
  } else {
    return fallback ? fallback() : null
  }
}

/**
 * 页面权限检查
 * @param {string} routeName 路由名称
 * @returns {boolean} 是否可以访问
 */
export function canAccessRoute(routeName) {
  const routePermissions = {
    'DepartmentManagement': PERMISSIONS.DEPARTMENT_VIEW,
    'DeviceManagement': PERMISSIONS.DEVICE_VIEW,
    'SystemManagement': PERMISSIONS.SYSTEM_MANAGE
  }
  
  const requiredPermission = routePermissions[routeName]
  return requiredPermission ? hasPermission(requiredPermission) : true
}

/**
 * 操作权限检查
 * @param {string} action 操作类型
 * @param {string} resource 资源类型
 * @returns {boolean} 是否可以执行操作
 */
export function canPerformAction(action, resource) {
  const permission = `${resource}:${action}`
  return hasPermission(permission)
}

/**
 * 批量权限检查
 * @param {Array} checks 检查项数组，每项包含 {permission, required}
 * @returns {Object} 权限检查结果
 */
export function batchPermissionCheck(checks) {
  const results = {}
  
  checks.forEach(({ key, permission, required = true }) => {
    results[key] = {
      hasPermission: hasPermission(permission),
      required,
      permission
    }
  })
  
  return results
}

// Vue 3 组合式API权限检查
export function usePermission() {
  return {
    hasPermission,
    hasAnyPermission,
    hasAllPermissions,
    canAccessRoute,
    canPerformAction,
    getCurrentUserRole,
    getCurrentUserPermissions,
    PERMISSIONS
  }
}

// 权限指令（用于模板中的权限控制）
export const permissionDirective = {
  mounted(el, binding) {
    const { value } = binding
    if (value && !hasPermission(value)) {
      el.style.display = 'none'
    }
  },
  updated(el, binding) {
    const { value } = binding
    if (value && !hasPermission(value)) {
      el.style.display = 'none'
    } else {
      el.style.display = ''
    }
  }
}

export default {
  hasPermission,
  hasAnyPermission,
  hasAllPermissions,
  canAccessRoute,
  canPerformAction,
  getCurrentUserRole,
  getCurrentUserPermissions,
  PERMISSIONS,
  ROLE_PERMISSIONS
}