import service from '@/utils/request'
// @Tags WlEngineRules
// @Summary 创建wlEngineRules表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.WlEngineRules true "创建wlEngineRules表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wlEngineRules/createWlEngineRules [post]
export const createWlEngineRules = (data) => {
  return service({
    url: '/wlEngineRules/createWlEngineRules',
    method: 'post',
    data
  })
}

// @Tags WlEngineRules
// @Summary 删除wlEngineRules表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.WlEngineRules true "删除wlEngineRules表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wlEngineRules/deleteWlEngineRules [delete]
export const deleteWlEngineRules = (params) => {
  return service({
    url: '/wlEngineRules/deleteWlEngineRules',
    method: 'delete',
    params
  })
}

// @Tags WlEngineRules
// @Summary 批量删除wlEngineRules表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除wlEngineRules表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wlEngineRules/deleteWlEngineRules [delete]
export const deleteWlEngineRulesByIds = (params) => {
  return service({
    url: '/wlEngineRules/deleteWlEngineRulesByIds',
    method: 'delete',
    params
  })
}

// @Tags WlEngineRules
// @Summary 更新wlEngineRules表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.WlEngineRules true "更新wlEngineRules表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wlEngineRules/updateWlEngineRules [put]
export const updateWlEngineRules = (data) => {
  return service({
    url: '/wlEngineRules/updateWlEngineRules',
    method: 'put',
    data
  })
}

// @Tags WlEngineRules
// @Summary 用id查询wlEngineRules表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.WlEngineRules true "用id查询wlEngineRules表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wlEngineRules/findWlEngineRules [get]
export const findWlEngineRules = (params) => {
  return service({
    url: '/wlEngineRules/findWlEngineRules',
    method: 'get',
    params
  })
}

// @Tags WlEngineRules
// @Summary 分页获取wlEngineRules表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取wlEngineRules表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wlEngineRules/getWlEngineRulesList [get]
export const getWlEngineRulesList = (params) => {
  return service({
    url: '/wlEngineRules/getWlEngineRulesList',
    method: 'get',
    params
  })
}

// @Tags WlEngineRules
// @Summary 不需要鉴权的wlEngineRules表接口
// @Accept application/json
// @Produce application/json
// @Param data query wl_playformReq.WlEngineRulesSearch true "分页获取wlEngineRules表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /wlEngineRules/getWlEngineRulesPublic [get]
export const getWlEngineRulesPublic = () => {
  return service({
    url: '/wlEngineRules/getWlEngineRulesPublic',
    method: 'get',
  })
}
