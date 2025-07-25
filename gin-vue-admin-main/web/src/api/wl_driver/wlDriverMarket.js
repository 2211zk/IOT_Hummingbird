import service from '@/utils/request'

export const getWlDriverMarketList = (params) => {
  return service({
    url: '/wlDriverMarket/getWlDriverMarketList',
    method: 'get',
    params
  })
}

export const createWlDriverMarket = (data) => {
  return service({
    url: '/wlDriverMarket/createWlDriverMarket',
    method: 'post',
    data
  })
}

export const updateWlDriverMarket = (data) => {
  return service({
    url: '/wlDriverMarket/updateWlDriverMarket',
    method: 'put',
    data
  })
}

export const deleteWlDriverMarket = (params) => {
  return service({
    url: '/wlDriverMarket/deleteWlDriverMarket',
    method: 'delete',
    params
  })
}

export const deleteWlDriverMarketByIds = (params) => {
  return service({
    url: '/wlDriverMarket/deleteWlDriverMarketByIds',
    method: 'delete',
    params
  })
}

export const findWlDriverMarket = (params) => {
  return service({
    url: '/wlDriverMarket/findWlDriverMarket',
    method: 'get',
    params
  })
} 