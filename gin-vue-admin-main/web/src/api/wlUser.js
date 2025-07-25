import service from '@/utils/request'

// 获取用户列表
export function getWlUserList(data) {
  return service({
    url: '/sysUser/getWlUserList',
    method: 'post',
    data
  })
} 