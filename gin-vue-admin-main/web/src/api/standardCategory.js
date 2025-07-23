import service from '@/utils/request'

// @Summary 获取标准品类列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.StandardCategorySearch true "分页获取标准品类列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /standardCategory/getStandardCategoryList [get]
export const getStandardCategoryList = (params) => {
  return service({
    url: '/standardCategory/getStandardCategoryList',
    method: 'get',
    params: params
  })
}

// @Summary 根据ID获取标准品类详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id path int true "标准品类ID"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /standardCategory/getStandardCategoryById/{id} [get]
export const getStandardCategoryById = (id) => {
  return service({
    url: `/standardCategory/getStandardCategoryById/${id}`,
    method: 'get'
  })
}

// @Summary 创建标准品类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.StandardCategory true "创建标准品类"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /standardCategory/createStandardCategory [post]
export const createStandardCategory = (data) => {
  return service({
    url: '/standardCategory/createStandardCategory',
    method: 'post',
    data: data
  })
}

// @Summary 更新标准品类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.StandardCategory true "更新标准品类"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /standardCategory/updateStandardCategory [put]
export const updateStandardCategory = (data) => {
  return service({
    url: '/standardCategory/updateStandardCategory',
    method: 'put',
    data: data
  })
}

// @Summary 删除标准品类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "删除标准品类"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /standardCategory/deleteStandardCategory [delete]
export const deleteStandardCategory = (data) => {
  return service({
    url: '/standardCategory/deleteStandardCategory',
    method: 'delete',
    data: data
  })
}

// @Summary 获取标准品类的所有类别
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /standardCategory/getStandardCategoryCategories [get]
export const getStandardCategoryCategories = () => {
  return service({
    url: '/standardCategory/getStandardCategoryCategories',
    method: 'get'
  })
}