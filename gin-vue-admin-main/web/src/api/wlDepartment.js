import service from '@/utils/request'

export function getWlDepartmentList(data) {
  return service({
    url: '/department/getWlDepartmentList',
    method: 'post',
    data
  })
}

export function addWlDepartment(data) {
  return service({
    url: '/department/createWlDepartment',
    method: 'post',
    data
  })
}

export function updateWlDepartment(data) {
  return service({
    url: '/department/updateWlDepartment',
    method: 'post',
    data
  })
}

export function deleteWlDepartment(data) {
  return service({
    url: '/department/deleteWlDepartment',
    method: 'post',
    data
  })
} 