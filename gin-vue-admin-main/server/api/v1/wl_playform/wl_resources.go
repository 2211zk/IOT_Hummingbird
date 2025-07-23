package wl_playform

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WlResourcesApi struct{}

// CreateWlResourcesWithTransaction 创建资源（事务处理）
// @Tags WlResources
// @Summary 创建资源（事务处理）
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CreateWlResourcesRequest true "创建资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wlResources/createWithTransaction [post]
func (wlResourcesApi *WlResourcesApi) CreateWlResourcesWithTransaction(c *gin.Context) {
	var requestData struct {
		InstanceName string                 `json:"instanceName"`
		ResourceType string                 `json:"resourceType"`
		ResourceData map[string]interface{} `json:"resourceData"`
		RequestID    string                 `json:"requestId"` // 添加请求ID
	}

	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 检查请求ID是否已处理过（简单的防重复机制）
	if requestData.RequestID != "" {
		// 这里可以添加Redis缓存检查，暂时用日志记录
		global.GVA_LOG.Info("处理请求", zap.String("requestId", requestData.RequestID))
	}

	// 创建WlResources对象
	wlResources := &wl_playform.WlResources{
		InstanceName: &requestData.InstanceName,
	}

	err = wlResourcesService.CreateWlResourcesWithTransaction(c.Request.Context(), wlResources, requestData.ResourceData)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// CreateWlResources 创建资源
// @Tags WlResources
// @Summary 创建资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wl_playform.WlResources true "创建资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wlResources/createWlResources [post]
func (wlResourcesApi *WlResourcesApi) CreateWlResources(c *gin.Context) {
	var wlResources wl_playform.WlResources
	err := c.ShouldBindJSON(&wlResources)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = wlResourcesService.CreateWlResources(&wlResources)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWlResources 删除资源
// @Tags WlResources
// @Summary 删除资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wl_playform.WlResources true "删除资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wlResources/deleteWlResources [delete]
func (wlResourcesApi *WlResourcesApi) DeleteWlResources(c *gin.Context) {
	ID := c.Query("ID")
	id, err := strconv.ParseUint(ID, 10, 32)
	if err != nil {
		response.FailWithMessage("ID格式错误", c)
		return
	}
	err = wlResourcesService.DeleteWlResourcesByIds([]uint{uint(id)})
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWlResourcesByIds 批量删除资源
// @Tags WlResources
// @Summary 批量删除资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wlResources/deleteWlResources [delete]
func (wlResourcesApi *WlResourcesApi) DeleteWlResourcesByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	var uintIDs []uint
	for _, idStr := range IDs {
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			response.FailWithMessage("ID格式错误", c)
			return
		}
		uintIDs = append(uintIDs, uint(id))
	}
	err := wlResourcesService.DeleteWlResourcesByIds(uintIDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWlResources 更新资源
// @Tags WlResources
// @Summary 更新资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wl_playform.WlResources true "更新资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wlResources/updateWlResources [put]
func (wlResourcesApi *WlResourcesApi) UpdateWlResources(c *gin.Context) {
	var wlResources wl_playform.WlResources
	err := c.ShouldBindJSON(&wlResources)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = wlResourcesService.UpdateWlResources(&wlResources)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWlResources 用id查询资源
// @Tags WlResources
// @Summary 用id查询资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wl_playform.WlResources true "用id查询资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wlResources/findWlResources [get]
func (wlResourcesApi *WlResourcesApi) FindWlResources(c *gin.Context) {
	ID := c.Query("ID")
	id, err := strconv.ParseUint(ID, 10, 32)
	if err != nil {
		response.FailWithMessage("ID格式错误", c)
		return
	}
	rewlResources, err := wlResourcesService.GetWlResources(uint(id))
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewlResources": rewlResources}, c)
	}
}

// GetWlResourcesList 分页获取资源列表
// @Tags WlResources
// @Summary 分页获取资源列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.WlResourcesSearch true "分页获取资源列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wlResources/getWlResourcesList [get]
func (wlResourcesApi *WlResourcesApi) GetWlResourcesList(c *gin.Context) {
	var pageInfo request.WlResourcesSearch

	// 手动绑定查询参数，避免验证错误
	if pageStr := c.Query("page"); pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil {
			pageInfo.Page = page
		} else {
			pageInfo.Page = 1
		}
	} else {
		pageInfo.Page = 1
	}

	if pageSizeStr := c.Query("pageSize"); pageSizeStr != "" {
		if pageSize, err := strconv.Atoi(pageSizeStr); err == nil {
			pageInfo.PageSize = pageSize
		} else {
			pageInfo.PageSize = 10
		}
	} else {
		pageInfo.PageSize = 10
	}

	// 处理搜索参数
	if instanceName := c.Query("instanceName"); instanceName != "" {
		pageInfo.InstanceName = &instanceName
	}
	if timeoutMs := c.Query("timeoutMs"); timeoutMs != "" {
		pageInfo.TimeoutMs = &timeoutMs
	}
	if verificationStatus := c.Query("verificationStatus"); verificationStatus != "" {
		pageInfo.VerificationStatus = &verificationStatus
	}
	if resourcesKey := c.Query("resourcesKey"); resourcesKey != "" {
		pageInfo.ResourcesKey = &resourcesKey
	}
	if startCreatedAt := c.Query("startCreatedAt"); startCreatedAt != "" {
		pageInfo.StartCreatedAt = &startCreatedAt
	}
	if endCreatedAt := c.Query("endCreatedAt"); endCreatedAt != "" {
		pageInfo.EndCreatedAt = &endCreatedAt
	}

	list, total, err := wlResourcesService.GetWlResourcesInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// VerifyWlResources 验证资源
// @Tags WlResources
// @Summary 验证资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.VerifyWlResourcesRequest true "验证资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证成功"}"
// @Router /wlResources/verifyWlResources [post]
func (wlResourcesApi *WlResourcesApi) VerifyWlResources(c *gin.Context) {
	var req request.VerifyWlResourcesRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 从数据库获取完整记录
	fullRecord, err := wlResourcesService.GetWlResources(req.ID)
	if err != nil {
		global.GVA_LOG.Error("获取资源记录失败!", zap.Error(err))
		response.FailWithMessage("获取资源记录失败", c)
		return
	}

	// 执行验证逻辑
	err = wlResourcesService.VerifyWlResources(fullRecord)
	if err != nil {
		global.GVA_LOG.Error("验证失败!", zap.Error(err))
		response.FailWithMessage("验证失败", c)
	} else {
		response.OkWithMessage("验证成功", c)
	}
}
