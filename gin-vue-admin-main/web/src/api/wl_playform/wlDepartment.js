import service from '@/utils/request'
// @Tags WlDepartment
// @Summary 创建wlDepartment表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.WlDepartment true "创建wlDepartment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wlDepartment/createWlDepartment [post]
export const createWlDepartment = (data) => {
  return service({
    url: '/wlDepartment/createWlDepartment',
    method: 'post',
    data
  })
}

// @Tags WlDepartment
// @Summary 删除wlDepartment表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.WlDepartment true "删除wlDepartment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wlDepartment/deleteWlDepartment [delete]
export const deleteWlDepartment = (params) => {
  return service({
    url: '/wlDepartment/deleteWlDepartment',
    method: 'delete',
    params
  })
}

// @Tags WlDepartment
// @Summary 批量删除wlDepartment表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除wlDepartment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wlDepartment/deleteWlDepartment [delete]
export const deleteWlDepartmentByIds = (params) => {
  return service({
    url: '/wlDepartment/deleteWlDepartmentByIds',
    method: 'delete',
    params
  })
}

// @Tags WlDepartment
// @Summary 更新wlDepartment表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.WlDepartment true "更新wlDepartment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wlDepartment/updateWlDepartment [put]
export const updateWlDepartment = (data) => {
  return service({
    url: '/wlDepartment/updateWlDepartment',
    method: 'put',
    data
  })
}

// @Tags WlDepartment
// @Summary 用id查询wlDepartment表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.WlDepartment true "用id查询wlDepartment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wlDepartment/findWlDepartment [get]
export const findWlDepartment = (params) => {
  return service({
    url: '/wlDepartment/findWlDepartment',
    method: 'get',
    params
  })
}

// @Tags WlDepartment
// @Summary 分页获取wlDepartment表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取wlDepartment表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wlDepartment/getWlDepartmentList [get]
export const getWlDepartmentList = (params) => {
  return service({
    url: '/wlDepartment/getWlDepartmentList',
    method: 'get',
    params
  })
}

// @Tags WlDepartment
// @Summary 不需要鉴权的wlDepartment表接口
// @Accept application/json
// @Produce application/json
// @Param data query wl_playformReq.WlDepartmentSearch true "分页获取wlDepartment表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /wlDepartment/getWlDepartmentPublic [get]
export const getWlDepartmentPublic = () => {
  return service({
    url: '/wlDepartment/getWlDepartmentPublic',
    method: 'get',
  })
}

// 分配设备到部门
export const assignDevicesToDepartment = (data) => {
  return service({
    url: '/wlDepartment/assignDevicesToDepartment',
    method: 'post',
    data
  })
}

// 查询部门下所有设备
export const getDevicesByDepartment = (data) => {
  return service({
    url: '/wlDepartment/getDevicesByDepartment',
    method: 'post',
    data
  })
}

// 移除部门下某个设备
export const removeDeviceFromDepartment = (data) => {
  return service({
    url: '/wlDepartment/removeDeviceFromDepartment',
    method: 'post',
    data
  })
}
