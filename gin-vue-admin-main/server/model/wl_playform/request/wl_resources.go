package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type WlResourcesSearch struct {
	InstanceName       *string `json:"instanceName" form:"instanceName"`             // 实例名称
	TimeoutMs          *string `json:"timeoutMs" form:"timeoutMs"`                   // 超时时间
	VerificationStatus *string `json:"verificationStatus" form:"verificationStatus"` // 验证状态
	ResourcesKey       *string `json:"resourcesKey" form:"resourcesKey"`             // MongoDB资源key
	request.PageInfo
	StartCreatedAt *string `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *string `json:"endCreatedAt" form:"endCreatedAt"`
}

// VerifyWlResourcesRequest 验证资源请求
type VerifyWlResourcesRequest struct {
	ID uint `json:"ID" form:"ID" binding:"required"`
}
