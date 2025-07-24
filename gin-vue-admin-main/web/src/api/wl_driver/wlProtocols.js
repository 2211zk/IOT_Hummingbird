import service from '@/utils/request'

export const getWlProtocolsList = (params) => {
  return service({
    url: '/wlProtocols/getWlProtocolsList',
    method: 'get',
    params
  })
}

export const createWlProtocols = (data) => {
  return service({
    url: '/wlProtocols/createWlProtocols',
    method: 'post',
    data
  })
}

export const updateWlProtocols = (data) => {
  return service({
    url: '/wlProtocols/updateWlProtocols',
    method: 'put',
    data
  })
}

export const deleteWlProtocols = (params) => {
  return service({
    url: '/wlProtocols/deleteWlProtocols',
    method: 'delete',
    params
  })
}

export const deleteWlProtocolsByIds = (params) => {
  return service({
    url: '/wlProtocols/deleteWlProtocolsByIds',
    method: 'delete',
    params
  })
}

export const findWlProtocols = (params) => {
  return service({
    url: '/wlProtocols/findWlProtocols',
    method: 'get',
    params
  })
} 