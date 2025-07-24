package wl_driver

import (
	"time"
)

type WlProtocols struct {
	ID             uint64     `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	CreatedAt      *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt      *time.Time `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt      *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
	ProtocolId     *string    `json:"protocol_id" gorm:"column:protocol_id;size:20;comment:协议编号"`
	ProtocolName   *string    `json:"protocol_name" gorm:"column:protocol_name;size:100;comment:协议名称"`
	ProtocolType   *string    `json:"protocol_type" gorm:"column:protocol_type;size:50;comment:协议类型"`
	Version        *string    `json:"version" gorm:"column:version;size:20;comment:版本"`
	Status         *string    `json:"status" gorm:"column:status;size:1;comment:状态"`
	Description    *string    `json:"description" gorm:"column:description;size:500;comment:描述"`
	ConfigTemplate *string    `json:"config_template" gorm:"column:config_template;type:text;comment:配置模板"`
	CreatedTime    *time.Time `json:"created_time" gorm:"column:created_time"`
	UpdateTime     *time.Time `json:"update_time" gorm:"column:update_time"`
	CreatedBy      *uint64    `json:"created_by" gorm:"column:created_by"`
	UpdatedBy      *uint64    `json:"updated_by" gorm:"column:updated_by"`
	DeletedBy      *uint64    `json:"deleted_by" gorm:"column:deleted_by"`
}

func (WlProtocols) TableName() string {
	return "wl_protocols"
}
