import service from '@/utils/request'
// @Tags WlAlarm
// @Summary 创建wlAlarm表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.WlAlarm true "创建wlAlarm表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wlAlarm/createWlAlarm [post]
export const createWlAlarm = (data) => {
  return service({
    url: '/wlAlarm/createWlAlarm',
    method: 'post',
    data
  })
}

// @Tags WlAlarm
// @Summary 删除wlAlarm表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.WlAlarm true "删除wlAlarm表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wlAlarm/deleteWlAlarm [delete]
export const deleteWlAlarm = (params) => {
  return service({
    url: '/wlAlarm/deleteWlAlarm',
    method: 'delete',
    params
  })
}

// @Tags WlAlarm
// @Summary 批量删除wlAlarm表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除wlAlarm表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wlAlarm/deleteWlAlarm [delete]
export const deleteWlAlarmByIds = (params) => {
  return service({
    url: '/wlAlarm/deleteWlAlarmByIds',
    method: 'delete',
    params
  })
}

// @Tags WlAlarm
// @Summary 更新wlAlarm表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.WlAlarm true "更新wlAlarm表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wlAlarm/updateWlAlarm [put]
export const updateWlAlarm = (data) => {
  return service({
    url: '/wlAlarm/updateWlAlarm',
    method: 'put',
    data
  })
}

// @Tags WlAlarm
// @Summary 用id查询wlAlarm表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.WlAlarm true "用id查询wlAlarm表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wlAlarm/findWlAlarm [get]
export const findWlAlarm = (params) => {
  return service({
    url: '/wlAlarm/findWlAlarm',
    method: 'get',
    params
  })
}

// @Tags WlAlarm
// @Summary 分页获取wlAlarm表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取wlAlarm表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wlAlarm/getWlAlarmList [get]
export const getWlAlarmList = (params) => {
  return service({
    url: '/wlAlarm/getWlAlarmList',
    method: 'get',
    params
  })
}

// @Tags WlAlarm
// @Summary 不需要鉴权的wlAlarm表接口
// @Accept application/json
// @Produce application/json
// @Param data query wl_playformReq.WlAlarmSearch true "分页获取wlAlarm表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /wlAlarm/getWlAlarmPublic [get]
export const getWlAlarmPublic = () => {
  return service({
    url: '/wlAlarm/getWlAlarmPublic',
    method: 'get',
  })
}
