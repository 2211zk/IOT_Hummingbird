import service from '@/utils/request'
// @Tags WlEquipment
// @Summary 创建wlEquipment表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.WlEquipment true "创建wlEquipment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wlEquipment/createWlEquipment [post]
export const createWlEquipment = (data) => {
  return service({
    url: '/wlEquipment/createWlEquipment',
    method: 'post',
    data
  })
}

// @Tags WlEquipment
// @Summary 删除wlEquipment表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.WlEquipment true "删除wlEquipment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wlEquipment/deleteWlEquipment [delete]
export const deleteWlEquipment = (params) => {
  return service({
    url: '/wlEquipment/deleteWlEquipment',
    method: 'delete',
    params
  })
}

// @Tags WlEquipment
// @Summary 批量删除wlEquipment表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除wlEquipment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wlEquipment/deleteWlEquipment [delete]
export const deleteWlEquipmentByIds = (params) => {
  return service({
    url: '/wlEquipment/deleteWlEquipmentByIds',
    method: 'delete',
    params
  })
}

// @Tags WlEquipment
// @Summary 更新wlEquipment表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.WlEquipment true "更新wlEquipment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wlEquipment/updateWlEquipment [put]
export const updateWlEquipment = (data) => {
  return service({
    url: '/wlEquipment/updateWlEquipment',
    method: 'put',
    data
  })
}

// @Tags WlEquipment
// @Summary 用id查询wlEquipment表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.WlEquipment true "用id查询wlEquipment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wlEquipment/findWlEquipment [get]
export const findWlEquipment = (params) => {
  return service({
    url: '/wlEquipment/findWlEquipment',
    method: 'get',
    params
  })
}

// @Tags WlEquipment
// @Summary 分页获取wlEquipment表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取wlEquipment表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wlEquipment/getWlEquipmentList [get]
export const getWlEquipmentList = (params) => {
  return service({
    url: '/wlEquipment/getWlEquipmentList',
    method: 'get',
    params
  })
}

// @Tags WlEquipment
// @Summary 不需要鉴权的wlEquipment表接口
// @Accept application/json
// @Produce application/json
// @Param data query wl_playformReq.WlEquipmentSearch true "分页获取wlEquipment表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /wlEquipment/getWlEquipmentPublic [get]
export const getWlEquipmentPublic = () => {
  return service({
    url: '/wlEquipment/getWlEquipmentPublic',
    method: 'get',
  })
}
