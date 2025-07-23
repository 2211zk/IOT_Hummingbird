import service from '@/utils/request'

// @Tags WlResources
// @Summary 创建资源（事务处理）
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CreateWlResourcesRequest true "创建资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wlResources/createWithTransaction [post]
export const createWlResourcesWithTransaction = (data) => {
  return service({
    url: '/wlResources/createWithTransaction',
    method: 'post',
    data
  })
}

// @Tags WlResources
// @Summary 创建资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wl_playform.WlResources true "创建资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wlResources/createWlResources [post]
export const createWlResources = (data) => {
  return service({
    url: '/wlResources/createWlResources',
    method: 'post',
    data
  })
}

// @Tags WlResources
// @Summary 删除资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wl_playform.WlResources true "删除资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wlResources/deleteWlResources [delete]
export const deleteWlResources = (params) => {
  return service({
    url: '/wlResources/deleteWlResources',
    method: 'delete',
    params
  })
}

// @Tags WlResources
// @Summary 批量删除资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wlResources/deleteWlResources [delete]
export const deleteWlResourcesByIds = (params) => {
  return service({
    url: '/wlResources/deleteWlResourcesByIds',
    method: 'delete',
    params
  })
}

// @Tags WlResources
// @Summary 更新资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wl_playform.WlResources true "更新资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wlResources/updateWlResources [put]
export const updateWlResources = (data) => {
  return service({
    url: '/wlResources/updateWlResources',
    method: 'put',
    data
  })
}

// @Tags WlResources
// @Summary 用id查询资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wl_playform.WlResources true "用id查询资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wlResources/findWlResources [get]
export const findWlResources = (params) => {
  return service({
    url: '/wlResources/findWlResources',
    method: 'get',
    params
  })
}

// @Tags WlResources
// @Summary 分页获取资源列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.WlResourcesSearch true "分页获取资源列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wlResources/getWlResourcesList [get]
export const getWlResourcesList = (params) => {
  return service({
    url: '/wlResources/getWlResourcesList',
    method: 'get',
    params
  })
}

// @Tags WlResources
// @Summary 验证资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wl_playform.WlResources true "验证资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证成功"}"
// @Router /wlResources/verifyWlResources [post]
export const verifyWlResources = (data) => {
  return service({
    url: '/wlResources/verifyWlResources',
    method: 'post',
    data
  })
}

// @Tags WlResources
// @Summary 不需要鉴权的wlResources表接口
// @Accept application/json
// @Produce application/json
// @Param data query wl_playformReq.WlResourcesSearch true "分页获取wlResources表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /wlResources/getWlResourcesPublic [get]
export const getWlResourcesPublic = () => {
  return service({
    url: '/wlResources/getWlResourcesPublic',
    method: 'get',
  })
}
