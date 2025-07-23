import service from '@/utils/request'

// @Tags WlCategory
// @Summary 创建wlCategory表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WlCategory true "创建wlCategory表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wlCategory/createWlCategory [post]
export const createWlCategory = (data) => {
  return service({
    url: '/wlCategory/createWlCategory',
    method: 'post',
    data
  })
}

// @Tags WlCategory
// @Summary 删除wlCategory表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WlCategory true "删除wlCategory表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wlCategory/deleteWlCategory [delete]
export const deleteWlCategory = (params) => {
  return service({
    url: '/wlCategory/deleteWlCategory',
    method: 'delete',
    params
  })
}

// @Tags WlCategory
// @Summary 批量删除wlCategory表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除wlCategory表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wlCategory/deleteWlCategory [delete]
export const deleteWlCategoryByIds = (params) => {
  return service({
    url: '/wlCategory/deleteWlCategoryByIds',
    method: 'delete',
    params
  })
}

// @Tags WlCategory
// @Summary 更新wlCategory表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WlCategory true "更新wlCategory表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wlCategory/updateWlCategory [put]
export const updateWlCategory = (data) => {
  return service({
    url: '/wlCategory/updateWlCategory',
    method: 'put',
    data
  })
}

// @Tags WlCategory
// @Summary 用id查询wlCategory表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WlCategory true "用id查询wlCategory表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wlCategory/findWlCategory [get]
export const findWlCategory = (params) => {
  return service({
    url: '/wlCategory/findWlCategory',
    method: 'get',
    params
  })
}

// @Tags WlCategory
// @Summary 分页获取wlCategory表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取wlCategory表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wlCategory/getWlCategoryList [get]
export const getWlCategoryList = (params) => {
  return service({
    url: '/wlCategory/getWlCategoryList',
    method: 'get',
    params
  })
}

// @Tags WlCategory
// @Summary 公开获取wlCategory表列表（无需权限）
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取wlCategory表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wlCategory/getWlCategoryPublic [get]
export const getWlCategoryPublic = (params) => {
  return service({
    url: '/wlCategory/getWlCategoryPublic',
    method: 'get',
    params
  })
} 