// 自动生成模板WlResources
package wl_playform

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// wlResources表 结构体  WlResources
type WlResources struct {
	global.GVA_MODEL
	InstanceName       *string    `json:"instanceName" form:"instanceName" gorm:"comment:实例名称;column:instance_name;size:30;" binding:"required"` //实例名称
	TimeoutMs          *time.Time `json:"timeoutMs" form:"timeoutMs" gorm:"comment:超市时间;column:timeout_ms;"`                                     //超时时间
	VerificationStatus *string    `json:"verificationStatus" form:"verificationStatus" gorm:"comment:验证状态;column:verification_status;size:20;"`  //验证状态
	ResourcesKey       *string    `json:"resourcesKey" form:"resourcesKey" gorm:"comment:MongoDB资源key;column:resources_key;size:100;"`           //MongoDB资源key
	CreatedBy          uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy          uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy          uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName wlResources表 WlResources自定义表名 wl_resources
func (WlResources) TableName() string {
	return "wl_resources"
}
