import service from '@/utils/request'

export const getWlSystemMoniorList = (params) => {
  return service({
    url: '/wlSystemMonior/getWlSystemMoniorList',
    method: 'get',
    params
  })
}

export const createWlSystemMonior = (data) => {
  return service({
    url: '/wlSystemMonior/createWlSystemMonior',
    method: 'post',
    data
  })
}

export const updateWlSystemMonior = (data) => {
  return service({
    url: '/wlSystemMonior/updateWlSystemMonior',
    method: 'put',
    data
  })
}

export const deleteWlSystemMonior = (params) => {
  return service({
    url: '/wlSystemMonior/deleteWlSystemMonior',
    method: 'delete',
    params
  })
}

export const deleteWlSystemMoniorByIds = (params) => {
  return service({
    url: '/wlSystemMonior/deleteWlSystemMoniorByIds',
    method: 'delete',
    params
  })
}

export const findWlSystemMonior = (params) => {
  return service({
    url: '/wlSystemMonior/findWlSystemMonior',
    method: 'get',
    params
  })
} 