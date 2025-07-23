import service from '@/utils/request'
// @Tags WlScenes
// @Summary 创建wlScenes表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.WlScenes true "创建wlScenes表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wlScenes/createWlScenes [post]
export const createWlScenes = (data) => {
  return service({
    url: '/wlScenes/createWlScenes',
    method: 'post',
    data
  })
}

// @Tags WlScenes
// @Summary 删除wlScenes表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.WlScenes true "删除wlScenes表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wlScenes/deleteWlScenes [delete]
export const deleteWlScenes = (params) => {
  return service({
    url: '/wlScenes/deleteWlScenes',
    method: 'delete',
    params
  })
}

// @Tags WlScenes
// @Summary 批量删除wlScenes表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除wlScenes表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wlScenes/deleteWlScenes [delete]
export const deleteWlScenesByIds = (params) => {
  return service({
    url: '/wlScenes/deleteWlScenesByIds',
    method: 'delete',
    params
  })
}

// @Tags WlScenes
// @Summary 更新wlScenes表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.WlScenes true "更新wlScenes表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wlScenes/updateWlScenes [put]
export const updateWlScenes = (data) => {
  return service({
    url: '/wlScenes/updateWlScenes',
    method: 'put',
    data
  })
}

// @Tags WlScenes
// @Summary 用id查询wlScenes表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.WlScenes true "用id查询wlScenes表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wlScenes/findWlScenes [get]
export const findWlScenes = (params) => {
  return service({
    url: '/wlScenes/findWlScenes',
    method: 'get',
    params
  })
}

// @Tags WlScenes
// @Summary 分页获取wlScenes表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取wlScenes表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wlScenes/getWlScenesList [get]
export const getWlScenesList = (params) => {
  return service({
    url: '/wlScenes/getWlScenesList',
    method: 'get',
    params
  })
}

// @Tags WlScenes
// @Summary 不需要鉴权的wlScenes表接口
// @Accept application/json
// @Produce application/json
// @Param data query wl_playformReq.WlScenesSearch true "分页获取wlScenes表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /wlScenes/getWlScenesPublic [get]
export const getWlScenesPublic = () => {
  return service({
    url: '/wlScenes/getWlScenesPublic',
    method: 'get',
  })
}
