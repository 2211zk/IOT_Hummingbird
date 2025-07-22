import service from '@/utils/request'
// @Tags WlProducts
// @Summary 创建wlProducts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.WlProducts true "创建wlProducts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wlProducts/createWlProducts [post]
export const createWlProducts = (data) => {
  return service({
    url: '/wlProducts/createWlProducts',
    method: 'post',
    data
  })
}

// @Tags WlProducts
// @Summary 删除wlProducts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.WlProducts true "删除wlProducts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wlProducts/deleteWlProducts [delete]
export const deleteWlProducts = (params) => {
  return service({
    url: '/wlProducts/deleteWlProducts',
    method: 'delete',
    params
  })
}

// @Tags WlProducts
// @Summary 批量删除wlProducts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除wlProducts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wlProducts/deleteWlProducts [delete]
export const deleteWlProductsByIds = (params) => {
  return service({
    url: '/wlProducts/deleteWlProductsByIds',
    method: 'delete',
    params
  })
}

// @Tags WlProducts
// @Summary 更新wlProducts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.WlProducts true "更新wlProducts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wlProducts/updateWlProducts [put]
export const updateWlProducts = (data) => {
  return service({
    url: '/wlProducts/updateWlProducts',
    method: 'put',
    data
  })
}

// @Tags WlProducts
// @Summary 用id查询wlProducts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.WlProducts true "用id查询wlProducts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wlProducts/findWlProducts [get]
export const findWlProducts = (params) => {
  return service({
    url: '/wlProducts/findWlProducts',
    method: 'get',
    params
  })
}

// @Tags WlProducts
// @Summary 分页获取wlProducts表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取wlProducts表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wlProducts/getWlProductsList [get]
export const getWlProductsList = (params) => {
  return service({
    url: '/wlProducts/getWlProductsList',
    method: 'get',
    params
  })
}

// @Tags WlProducts
// @Summary 不需要鉴权的wlProducts表接口
// @Accept application/json
// @Produce application/json
// @Param data query wl_playformReq.WlProductsSearch true "分页获取wlProducts表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /wlProducts/getWlProductsPublic [get]
export const getWlProductsPublic = () => {
  return service({
    url: '/wlProducts/getWlProductsPublic',
    method: 'get',
  })
}
