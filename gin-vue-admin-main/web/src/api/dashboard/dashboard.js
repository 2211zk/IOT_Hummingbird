import service from '@/utils/request'

// @Tags Dashboard
// @Summary 获取仪表盘数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body object true "获取仪表盘数据"
// @Success 200 {object} response.Response{data=object} "获取成功"
// @Router /dashboard/getDashboardData [get]
export const getDashboardData = () => {
  return service({
    url: '/dashboard/getDashboardData',
    method: 'get',
    donNotShowLoading: true
  })
} 