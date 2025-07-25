package system

import (
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"

	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LoginLogApi struct{}

// CreateLoginLog 创建登录日志
// @Tags LoginLog
// @Summary 创建登录日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.CreateLoginLogReq true "创建登录日志"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /loginLog/createLoginLog [post]
func (loginLogApi *LoginLogApi) CreateLoginLog(c *gin.Context) {
	var req systemReq.CreateLoginLogReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	loginLog := system.SysLoginLog{
		AccessNumber:    req.AccessNumber,
		UserName:        req.UserName,
		LoginAddress:    req.LoginAddress,
		LoginLocation:   req.LoginLocation,
		Browser:         req.Browser,
		OperatingSystem: req.OperatingSystem,
		LoginStatus:     req.LoginStatus,
		OperationalInfo: req.OperationalInfo,
	}

	err = loginLogService.CreateLoginLog(loginLog)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteLoginLog 删除登录日志
// @Tags LoginLog
// @Summary 删除登录日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "删除登录日志"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /loginLog/deleteLoginLog [delete]
func (loginLogApi *LoginLogApi) DeleteLoginLog(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindJSON(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(reqId, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	loginLog := system.SysLoginLog{GVA_MODEL: global.GVA_MODEL{ID: uint(reqId.ID)}}
	err = loginLogService.DeleteLoginLog(loginLog)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteLoginLogByIds 批量删除登录日志
// @Tags LoginLog
// @Summary 批量删除登录日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.DeleteLoginLogReq true "批量删除登录日志"
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /loginLog/deleteLoginLogByIds [delete]
func (loginLogApi *LoginLogApi) DeleteLoginLogByIds(c *gin.Context) {
	var req systemReq.DeleteLoginLogReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = loginLogService.DeleteLoginLogByIds(req.IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateLoginLog 更新登录日志
// @Tags LoginLog
// @Summary 更新登录日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysLoginLog true "更新登录日志"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /loginLog/updateLoginLog [put]
func (loginLogApi *LoginLogApi) UpdateLoginLog(c *gin.Context) {
	var loginLog system.SysLoginLog
	err := c.ShouldBindJSON(&loginLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(loginLog, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = loginLogService.UpdateLoginLog(loginLog)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindLoginLog 用id查询登录日志
// @Tags LoginLog
// @Summary 用id查询登录日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.GetById true "用id查询登录日志"
// @Success 200 {object} response.Response{data=systemRes.LoginLogResponse,msg=string} "查询成功"
// @Router /loginLog/findLoginLog [get]
func (loginLogApi *LoginLogApi) FindLoginLog(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindQuery(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(reqId, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	loginLog, err := loginLogService.GetLoginLog(uint(reqId.ID))
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithDetailed(systemRes.LoginLogResponse{LoginLog: loginLog}, "查询成功", c)
}

// GetLoginLogList 分页获取登录日志列表
// @Tags LoginLog
// @Summary 分页获取登录日志列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.LoginLogSearchReq true "分页获取登录日志列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /loginLog/getLoginLogList [post]
func (loginLogApi *LoginLogApi) GetLoginLogList(c *gin.Context) {
	var pageInfo systemReq.LoginLogSearchReq
	var raw map[string]interface{}

	// 兼容空body（EOF）
	err := c.ShouldBindJSON(&raw)
	if err != nil && err.Error() != "EOF" {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 先绑定除时间外的字段
	_ = c.ShouldBindJSON(&pageInfo)

	// 分页参数默认值
	if pageInfo.Page == 0 {
		pageInfo.Page = 1
	}
	if pageInfo.PageSize == 0 {
		pageInfo.PageSize = 10
	}

	// 手动处理时间字段
	if start, ok := raw["startTime"].(string); ok && start != "" {
		t, err := time.Parse(time.RFC3339, start)
		if err == nil {
			pageInfo.StartTime = t
		}
	}
	if end, ok := raw["endTime"].(string); ok && end != "" {
		t, err := time.Parse(time.RFC3339, end)
		if err == nil {
			pageInfo.EndTime = t
		}
	}

	// 移除分页参数校验，避免Page值不能为空报错

	global.GVA_LOG.Info("分页参数", zap.Int("page", pageInfo.Page), zap.Int("pageSize", pageInfo.PageSize))
	list, total, err := loginLogService.GetLoginLogInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetLoginStatistics 获取登录统计信息
// @Tags LoginLog
// @Summary 获取登录统计信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param days query int false "统计天数" default(7)
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取成功"
// @Router /loginLog/getLoginStatistics [get]
func (loginLogApi *LoginLogApi) GetLoginStatistics(c *gin.Context) {
	daysStr := c.DefaultQuery("days", "7")
	days, err := strconv.Atoi(daysStr)
	if err != nil || days <= 0 {
		days = 7
	}

	statistics, err := loginLogService.GetLoginStatistics(days)
	if err != nil {
		global.GVA_LOG.Error("获取统计信息失败!", zap.Error(err))
		response.FailWithMessage("获取统计信息失败", c)
		return
	}
	response.OkWithDetailed(statistics, "获取成功", c)
}

// GetRecentLoginLogs 获取最近登录日志
// @Tags LoginLog
// @Summary 获取最近登录日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param hours query int false "小时数" default(24)
// @Param limit query int false "限制数量" default(50)
// @Success 200 {object} response.Response{data=[]system.SysLoginLog,msg=string} "获取成功"
// @Router /loginLog/getRecentLoginLogs [get]
func (loginLogApi *LoginLogApi) GetRecentLoginLogs(c *gin.Context) {
	hoursStr := c.DefaultQuery("hours", "24")
	limitStr := c.DefaultQuery("limit", "50")

	hours, err := strconv.Atoi(hoursStr)
	if err != nil || hours <= 0 {
		hours = 24
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 50
	}

	logs, err := loginLogService.GetRecentLoginLogs(hours, limit)
	if err != nil {
		global.GVA_LOG.Error("获取最近登录日志失败!", zap.Error(err))
		response.FailWithMessage("获取最近登录日志失败", c)
		return
	}
	response.OkWithDetailed(logs, "获取成功", c)
}

// GetTopLoginIPs 获取热门登录IP
// @Tags LoginLog
// @Summary 获取热门登录IP
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param limit query int false "限制数量" default(10)
// @Param days query int false "统计天数" default(7)
// @Success 200 {object} response.Response{data=[]map[string]interface{},msg=string} "获取成功"
// @Router /loginLog/getTopLoginIPs [get]
func (loginLogApi *LoginLogApi) GetTopLoginIPs(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	daysStr := c.DefaultQuery("days", "7")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	days, err := strconv.Atoi(daysStr)
	if err != nil || days <= 0 {
		days = 7
	}

	ips, err := loginLogService.GetTopLoginIPs(limit, days)
	if err != nil {
		global.GVA_LOG.Error("获取热门登录IP失败!", zap.Error(err))
		response.FailWithMessage("获取热门登录IP失败", c)
		return
	}
	response.OkWithDetailed(ips, "获取成功", c)
}

// ExportLoginLog 导出登录日志
// @Tags LoginLog
// @Summary 导出登录日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/octet-stream
// @Param data body systemReq.LoginLogSearchReq true "导出登录日志"
// @Success 200 {file} file "导出成功"
// @Router /loginLog/exportLoginLog [post]
func (loginLogApi *LoginLogApi) ExportLoginLog(c *gin.Context) {
	var searchReq systemReq.LoginLogSearchReq
	err := c.ShouldBindJSON(&searchReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 导出Excel文件
	excelData, filename, err := loginLogService.ExportLoginLogsToExcel(searchReq)
	if err != nil {
		global.GVA_LOG.Error("导出失败!", zap.Error(err))
		response.FailWithMessage("导出失败", c)
		return
	}

	// 设置响应头
	headers := utils.SetExcelResponseHeaders(filename)
	for key, value := range headers {
		c.Header(key, value)
	}

	// 返回文件数据
	c.Data(200, headers["Content-Type"], excelData)
}

// ExportLoginStatistics 导出登录统计信息
// @Tags LoginLog
// @Summary 导出登录统计信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/octet-stream
// @Param days query int false "统计天数" default(7)
// @Success 200 {file} file "导出成功"
// @Router /loginLog/exportLoginStatistics [get]
func (loginLogApi *LoginLogApi) ExportLoginStatistics(c *gin.Context) {
	daysStr := c.DefaultQuery("days", "7")
	days, err := strconv.Atoi(daysStr)
	if err != nil || days <= 0 {
		days = 7
	}

	// 导出Excel文件
	excelData, filename, err := loginLogService.ExportLoginStatisticsToExcel(days)
	if err != nil {
		global.GVA_LOG.Error("导出统计信息失败!", zap.Error(err))
		response.FailWithMessage("导出统计信息失败", c)
		return
	}

	// 设置响应头
	headers := utils.SetExcelResponseHeaders(filename)
	for key, value := range headers {
		c.Header(key, value)
	}

	// 返回文件数据
	c.Data(200, headers["Content-Type"], excelData)
}

// CleanExpiredLogs 清理过期日志
// @Tags LoginLog
// @Summary 清理过期日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.CleanExpiredLogsReq true "清理过期日志"
// @Success 200 {object} response.Response{data=systemRes.CleanLogsResponse,msg=string} "清理成功"
// @Router /loginLog/cleanExpiredLogs [post]
func (loginLogApi *LoginLogApi) CleanExpiredLogs(c *gin.Context) {
	var req systemReq.CleanExpiredLogsReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if req.Days <= 0 {
		response.FailWithMessage("保留天数必须大于0", c)
		return
	}

	// 执行清理（带备份）
	result, err := loginLogService.CleanExpiredLogsWithBackup(req.Days, true)
	if err != nil {
		global.GVA_LOG.Error("清理过期日志失败!", zap.Error(err))
		response.FailWithMessage("清理失败: "+err.Error(), c)
		return
	}

	cleanResponse := systemRes.CleanLogsResponse{
		DeletedCount: result["deletedCount"].(int64),
	}

	response.OkWithDetailed(cleanResponse, "清理成功", c)
}

// GetCleanupStatistics 获取清理统计信息
// @Tags LoginLog
// @Summary 获取清理统计信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param days query int true "保留天数"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取成功"
// @Router /loginLog/getCleanupStatistics [get]
func (loginLogApi *LoginLogApi) GetCleanupStatistics(c *gin.Context) {
	daysStr := c.Query("days")
	if daysStr == "" {
		response.FailWithMessage("请提供保留天数参数", c)
		return
	}

	days, err := strconv.Atoi(daysStr)
	if err != nil || days <= 0 {
		response.FailWithMessage("保留天数必须是大于0的整数", c)
		return
	}

	statistics, err := loginLogService.GetCleanupStatistics(days)
	if err != nil {
		global.GVA_LOG.Error("获取清理统计信息失败!", zap.Error(err))
		response.FailWithMessage("获取清理统计信息失败", c)
		return
	}

	response.OkWithDetailed(statistics, "获取成功", c)
}

// GetLogRetentionPolicy 获取日志保留策略
// @Tags LoginLog
// @Summary 获取日志保留策略
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取成功"
// @Router /loginLog/getLogRetentionPolicy [get]
func (loginLogApi *LoginLogApi) GetLogRetentionPolicy(c *gin.Context) {
	days, err := loginLogService.GetLogRetentionPolicy()
	if err != nil {
		global.GVA_LOG.Error("获取日志保留策略失败!", zap.Error(err))
		response.FailWithMessage("获取日志保留策略失败", c)
		return
	}

	result := map[string]interface{}{
		"retentionDays": days,
	}

	response.OkWithDetailed(result, "获取成功", c)
}

// SetLogRetentionPolicy 设置日志保留策略
// @Tags LoginLog
// @Summary 设置日志保留策略
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body map[string]int true "保留策略"
// @Success 200 {object} response.Response{msg=string} "设置成功"
// @Router /loginLog/setLogRetentionPolicy [post]
func (loginLogApi *LoginLogApi) SetLogRetentionPolicy(c *gin.Context) {
	var req map[string]int
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	days, exists := req["days"]
	if !exists || days <= 0 {
		response.FailWithMessage("保留天数必须大于0", c)
		return
	}

	err = loginLogService.SetLogRetentionPolicy(days)
	if err != nil {
		global.GVA_LOG.Error("设置日志保留策略失败!", zap.Error(err))
		response.FailWithMessage("设置失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("设置成功", c)
}

// GetFailedLoginAttempts 获取失败登录尝试次数
// @Tags LoginLog
// @Summary 获取失败登录尝试次数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param username query string true "用户名"
// @Param hours query int false "小时数" default(24)
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取成功"
// @Router /loginLog/getFailedLoginAttempts [get]
func (loginLogApi *LoginLogApi) GetFailedLoginAttempts(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		response.FailWithMessage("请提供用户名参数", c)
		return
	}

	hoursStr := c.DefaultQuery("hours", "24")
	hours, err := strconv.Atoi(hoursStr)
	if err != nil || hours <= 0 {
		hours = 24
	}

	count, err := loginLogService.GetFailedLoginAttempts(username, hours)
	if err != nil {
		global.GVA_LOG.Error("获取失败登录尝试次数失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	result := map[string]interface{}{
		"username":    username,
		"hours":       hours,
		"failedCount": count,
	}

	response.OkWithDetailed(result, "获取成功", c)
}
