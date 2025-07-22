package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// StandardCategory 标准品类模型
type StandardCategory struct {
	global.GVA_MODEL
	Name        string `json:"name" gorm:"not null;comment:品类名称"`                    // 品类名称
	Code        string `json:"code" gorm:"uniqueIndex;not null;comment:品类编码"`        // 品类编码
	Category    string `json:"category" gorm:"index;comment:所属类别"`                   // 所属类别
	Description string `json:"description" gorm:"type:text;comment:描述信息"`            // 描述信息
	Status      int    `json:"status" gorm:"default:1;comment:状态 1:启用 0:禁用"`        // 状态 1:启用 0:禁用
}

func (StandardCategory) TableName() string {
	return "standard_categories"
}