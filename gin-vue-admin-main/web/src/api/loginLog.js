import service from '@/utils/request'

// 获取登录日志列表
export const getLoginLogList = (data) => {
  return service({
    url: '/loginLog/getLoginLogList',
    method: 'post',
    data
  })
}

// 获取登录日志详情
export const getLoginLogDetail = (id) => {
  return service({
    url: `/loginLog/findLoginLog?ID=${id}`,
    method: 'get'
  })
}

// 删除登录日志
export const deleteLoginLog = (data) => {
  return service({
    url: '/loginLog/deleteLoginLog',
    method: 'delete',
    data
  })
}

// 批量删除登录日志
export const deleteLoginLogByIds = (data) => {
  return service({
    url: '/loginLog/deleteLoginLogByIds',
    method: 'delete',
    data
  })
}

// 导出登录日志
export const exportLoginLog = (data) => {
  return service({
    url: '/loginLog/exportLoginLog',
    method: 'post',
    data,
    responseType: 'blob'
  })
}

// 获取登录统计信息
export const getLoginStatistics = (days = 7) => {
  return service({
    url: `/loginLog/getLoginStatistics?days=${days}`,
    method: 'get'
  })
}

// 获取最近登录日志
export const getRecentLoginLogs = (hours = 24, limit = 50) => {
  return service({
    url: `/loginLog/getRecentLoginLogs?hours=${hours}&limit=${limit}`,
    method: 'get'
  })
}

// 获取热门登录IP
export const getTopLoginIPs = (limit = 10, days = 7) => {
  return service({
    url: `/loginLog/getTopLoginIPs?limit=${limit}&days=${days}`,
    method: 'get'
  })
}

// 清理过期日志
export const cleanExpiredLogs = (data) => {
  return service({
    url: '/loginLog/cleanExpiredLogs',
    method: 'post',
    data
  })
}

// 获取清理统计信息
export const getCleanupStatistics = (days) => {
  return service({
    url: `/loginLog/getCleanupStatistics?days=${days}`,
    method: 'get'
  })
}

// 获取日志保留策略
export const getLogRetentionPolicy = () => {
  return service({
    url: '/loginLog/getLogRetentionPolicy',
    method: 'get'
  })
}

// 设置日志保留策略
export const setLogRetentionPolicy = (data) => {
  return service({
    url: '/loginLog/setLogRetentionPolicy',
    method: 'post',
    data
  })
}

// 获取失败登录尝试次数
export const getFailedLoginAttempts = (username, hours = 24) => {
  return service({
    url: `/loginLog/getFailedLoginAttempts?username=${username}&hours=${hours}`,
    method: 'get'
  })
}

// 导出登录统计信息
export const exportLoginStatistics = (days = 7) => {
  return service({
    url: `/loginLog/exportLoginStatistics?days=${days}`,
    method: 'get',
    responseType: 'blob'
  })
}