package wl_driver

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type DriverCards struct {
	global.GVA_MODEL
	Name        *string `json:"name" form:"name" gorm:"comment:驱动名称;column:name;size:128;"`
	Img         *string `json:"img" form:"img" gorm:"comment:图片链接;column:img;size:255;"`
	Description *string `json:"description" form:"description" gorm:"comment:描述;column:description;size:255;"`
	Tags        *string `json:"tags" form:"tags" gorm:"comment:标签;column:tags;size:255;"`
	CreatedBy   uint    `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint    `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint    `gorm:"column:deleted_by;comment:删除者"`
}

func (DriverCards) TableName() string {
	return "driver_cards"
}
