package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform"
)

type WlResourcesSearch struct {
	wl_playform.WlResources
	request.PageInfo
	StartCreatedAt *string `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *string `json:"endCreatedAt" form:"endCreatedAt"`
}

// VerifyWlResourcesRequest 验证资源请求
type VerifyWlResourcesRequest struct {
	ID uint `json:"ID" form:"ID" binding:"required"`
}
