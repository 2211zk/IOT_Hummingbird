import service from '@/utils/request'
// @Tags DriverCards
// @Summary 创建driverCards表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DriverCards true "创建driverCards表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /driverCards/createDriverCards [post]
export const createDriverCards = (data) => {
  return service({
    url: '/driverCards/createDriverCards',
    method: 'post',
    data
  })
}

// @Tags DriverCards
// @Summary 删除driverCards表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DriverCards true "删除driverCards表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /driverCards/deleteDriverCards [delete]
export const deleteDriverCards = (params) => {
  return service({
    url: '/driverCards/deleteDriverCards',
    method: 'delete',
    params
  })
}

// @Tags DriverCards
// @Summary 批量删除driverCards表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除driverCards表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /driverCards/deleteDriverCards [delete]
export const deleteDriverCardsByIds = (params) => {
  return service({
    url: '/driverCards/deleteDriverCardsByIds',
    method: 'delete',
    params
  })
}

// @Tags DriverCards
// @Summary 更新driverCards表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DriverCards true "更新driverCards表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /driverCards/updateDriverCards [put]
export const updateDriverCards = (data) => {
  return service({
    url: '/driverCards/updateDriverCards',
    method: 'put',
    data
  })
}

// @Tags DriverCards
// @Summary 用id查询driverCards表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.DriverCards true "用id查询driverCards表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /driverCards/findDriverCards [get]
export const findDriverCards = (params) => {
  return service({
    url: '/driverCards/findDriverCards',
    method: 'get',
    params
  })
}

// @Tags DriverCards
// @Summary 分页获取driverCards表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取driverCards表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /driverCards/getDriverCardsList [get]
export const getDriverCardsList = (params) => {
  return service({
    url: '/driverCards/getDriverCardsList',
    method: 'get',
    params
  })
}

// @Tags DriverCards
// @Summary 不需要鉴权的driverCards表接口
// @Accept application/json
// @Produce application/json
// @Param data query wl_driverReq.DriverCardsSearch true "分页获取driverCards表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /driverCards/getDriverCardsPublic [get]
export const getDriverCardsPublic = () => {
  return service({
    url: '/driverCards/getDriverCardsPublic',
    method: 'get',
  })
}
