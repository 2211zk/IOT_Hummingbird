import service from '@/utils/request'
// @Tags WlUser
// @Summary 创建wlUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.WlUser true "创建wlUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wlUser/createWlUser [post]
export const createWlUser = (data) => {
  return service({
    url: '/wlUser/createWlUser',
    method: 'post',
    data
  })
}

// @Tags WlUser
// @Summary 删除wlUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.WlUser true "删除wlUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wlUser/deleteWlUser [delete]
export const deleteWlUser = (params) => {
  return service({
    url: '/wlUser/deleteWlUser',
    method: 'delete',
    params
  })
}

// @Tags WlUser
// @Summary 批量删除wlUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除wlUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wlUser/deleteWlUser [delete]
export const deleteWlUserByIds = (params) => {
  return service({
    url: '/wlUser/deleteWlUserByIds',
    method: 'delete',
    params
  })
}

// @Tags WlUser
// @Summary 更新wlUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.WlUser true "更新wlUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wlUser/updateWlUser [put]
export const updateWlUser = (data) => {
  return service({
    url: '/wlUser/updateWlUser',
    method: 'put',
    data
  })
}

// @Tags WlUser
// @Summary 用id查询wlUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.WlUser true "用id查询wlUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wlUser/findWlUser [get]
export const findWlUser = (params) => {
  return service({
    url: '/wlUser/findWlUser',
    method: 'get',
    params
  })
}

// @Tags WlUser
// @Summary 分页获取wlUser表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取wlUser表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wlUser/getWlUserList [get]
export const getWlUserList = (params) => {
  return service({
    url: '/wlUser/getWlUserList',
    method: 'get',
    params
  })
}

// @Tags WlUser
// @Summary 不需要鉴权的wlUser表接口
// @Accept application/json
// @Produce application/json
// @Param data query wl_playformReq.WlUserSearch true "分页获取wlUser表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /wlUser/getWlUserPublic [get]
export const getWlUserPublic = () => {
  return service({
    url: '/wlUser/getWlUserPublic',
    method: 'get',
  })
}
