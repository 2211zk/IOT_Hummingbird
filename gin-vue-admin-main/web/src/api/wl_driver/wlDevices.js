import service from '@/utils/request'

// 获取设备列表（分页）
export const getWlDevicesList = (params) => {
  return service({
    url: '/wlDevices/getWlDevicesList',
    method: 'get',
    params
  })
}

// 新增设备
export const createWlDevices = (data) => {
  return service({
    url: '/wlDevices/createWlDevices',
    method: 'post',
    data
  })
}

// 编辑设备
export const updateWlDevices = (data) => {
  return service({
    url: '/wlDevices/updateWlDevices',
    method: 'put',
    data
  })
}

// 删除设备
export const deleteWlDevices = (params) => {
  return service({
    url: '/wlDevices/deleteWlDevices',
    method: 'delete',
    params
  })
}

// 批量删除设备
export const deleteWlDevicesByIds = (params) => {
  return service({
    url: '/wlDevices/deleteWlDevicesByIds',
    method: 'delete',
    params
  })
}

// 查询单个设备详情
export const findWlDevices = (params) => {
  return service({
    url: '/wlDevices/findWlDevices',
    method: 'get',
    params
  })
} 