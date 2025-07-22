import service from '@/utils/request'

// 获取设备地图数据
export const getDeviceMapData = (params) => {
  return service({
    url: '/wlEquipment/getWlEquipmentList',
    method: 'get',
    params
  })
}

// 获取设备分布统计
export const getDeviceDistribution = () => {
  return service({
    url: '/wlEquipment/getDeviceDistribution',
    method: 'get'
  })
} 