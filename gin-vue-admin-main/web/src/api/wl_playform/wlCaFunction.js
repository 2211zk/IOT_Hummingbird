import service from '@/utils/request'

// @Tags WlCaFunction
// @Summary 创建wlCaFunction表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WlCaFunction true "创建wlCaFunction表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wlCaFunction/createWlCaFunction [post]
export const createWlCaFunction = (data) => {
  return service({
    url: '/wlCaFunction/createWlCaFunction',
    method: 'post',
    data
  })
}

// @Tags WlCaFunction
// @Summary 删除wlCaFunction表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WlCaFunction true "删除wlCaFunction表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wlCaFunction/deleteWlCaFunction [delete]
export const deleteWlCaFunction = (params) => {
  return service({
    url: '/wlCaFunction/deleteWlCaFunction',
    method: 'delete',
    params
  })
}

// @Tags WlCaFunction
// @Summary 批量删除wlCaFunction表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除wlCaFunction表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wlCaFunction/deleteWlCaFunction [delete]
export const deleteWlCaFunctionByIds = (params) => {
  return service({
    url: '/wlCaFunction/deleteWlCaFunctionByIds',
    method: 'delete',
    params
  })
}

// @Tags WlCaFunction
// @Summary 更新wlCaFunction表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WlCaFunction true "更新wlCaFunction表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wlCaFunction/updateWlCaFunction [put]
export const updateWlCaFunction = (data) => {
  return service({
    url: '/wlCaFunction/updateWlCaFunction',
    method: 'put',
    data
  })
}

// @Tags WlCaFunction
// @Summary 用id查询wlCaFunction表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WlCaFunction true "用id查询wlCaFunction表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wlCaFunction/findWlCaFunction [get]
export const findWlCaFunction = (params) => {
  return service({
    url: '/wlCaFunction/findWlCaFunction',
    method: 'get',
    params
  })
}

// @Tags WlCaFunction
// @Summary 分页获取wlCaFunction表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取wlCaFunction表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wlCaFunction/getWlCaFunctionList [get]
export const getWlCaFunctionList = (params) => {
  return service({
    url: '/wlCaFunction/getWlCaFunctionList',
    method: 'get',
    params
  })
}

// @Tags WlCaFunction
// @Summary 公开获取wlCaFunction表列表（无需权限）
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取wlCaFunction表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wlCaFunction/getWlCaFunctionPublic [get]
export const getWlCaFunctionPublic = (params) => {
  return service({
    url: '/wlCaFunction/getWlCaFunctionPublic',
    method: 'get',
    params
  })
}

// @Tags WlCaFunction
// @Summary 根据品类ID获取功能定义列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param caId query int true "品类ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wlCaFunction/getWlCaFunctionByCategory [get]
export const getWlCaFunctionByCategory = (params) => {
  return service({
    url: '/wlCaFunction/getWlCaFunctionList',
    method: 'get',
    params
  })
} 