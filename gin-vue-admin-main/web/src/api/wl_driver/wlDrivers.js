import service from '@/utils/request'
// @Tags WlDrivers
// @Summary 创建wlDrivers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.WlDrivers true "创建wlDrivers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wlDrivers/createWlDrivers [post]
export const createWlDrivers = (data) => {
  return service({
    url: '/wlDrivers/createWlDrivers',
    method: 'post',
    data
  })
}

// @Tags WlDrivers
// @Summary 删除wlDrivers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.WlDrivers true "删除wlDrivers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wlDrivers/deleteWlDrivers [delete]
export const deleteWlDrivers = (params) => {
  return service({
    url: '/wlDrivers/deleteWlDrivers',
    method: 'delete',
    params
  })
}

// @Tags WlDrivers
// @Summary 批量删除wlDrivers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除wlDrivers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wlDrivers/deleteWlDrivers [delete]
export const deleteWlDriversByIds = (params) => {
  return service({
    url: '/wlDrivers/deleteWlDriversByIds',
    method: 'delete',
    params
  })
}

// @Tags WlDrivers
// @Summary 更新wlDrivers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.WlDrivers true "更新wlDrivers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wlDrivers/updateWlDrivers [put]
export const updateWlDrivers = (data) => {
  return service({
    url: '/wlDrivers/updateWlDrivers',
    method: 'put',
    data
  })
}

// @Tags WlDrivers
// @Summary 用id查询wlDrivers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.WlDrivers true "用id查询wlDrivers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wlDrivers/findWlDrivers [get]
export const findWlDrivers = (params) => {
  return service({
    url: '/wlDrivers/findWlDrivers',
    method: 'get',
    params
  })
}

// @Tags WlDrivers
// @Summary 分页获取wlDrivers表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取wlDrivers表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wlDrivers/getWlDriversList [get]
export const getWlDriversList = (params) => {
  return service({
    url: '/wlDrivers/getWlDriversList',
    method: 'get',
    params
  })
}

// @Tags WlDrivers
// @Summary 不需要鉴权的wlDrivers表接口
// @Accept application/json
// @Produce application/json
// @Param data query wl_driverReq.WlDriversSearch true "分页获取wlDrivers表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /wlDrivers/getWlDriversPublic [get]
export const getWlDriversPublic = () => {
  return service({
    url: '/wlDrivers/getWlDriversPublic',
    method: 'get',
  })
}
