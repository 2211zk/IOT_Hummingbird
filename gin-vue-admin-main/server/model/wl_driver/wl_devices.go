package wl_driver

import (
	"time"
)

type WlDevices struct {
	ID            uint64     `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	CreatedAt     *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt     *time.Time `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt     *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
	DeviceId      *string    `json:"device_id" gorm:"column:device_id;size:50;comment:设备编号"`
	DeviceName    *string    `json:"device_name" gorm:"column:device_name;size:100;comment:设备名称"`
	DeviceType    *string    `json:"device_type" gorm:"column:device_type;size:50;comment:设备类型"`
	ProtocolType  *string    `json:"protocol_type" gorm:"column:protocol_type;size:50;comment:协议类型"`
	Status        *string    `json:"status" gorm:"column:status;size:1;comment:状态"`
	OnlineStatus  *string    `json:"online_status" gorm:"column:online_status;size:1;comment:在线状态"`
	LastHeartbeat *time.Time `json:"last_heartbeat" gorm:"column:last_heartbeat"`
	CreatedTime   *time.Time `json:"created_time" gorm:"column:created_time"`
	UpdateTime    *time.Time `json:"update_time" gorm:"column:update_time"`
	CreatedBy     *uint64    `json:"created_by" gorm:"column:created_by"`
	UpdatedBy     *uint64    `json:"updated_by" gorm:"column:updated_by"`
	DeletedBy     *uint64    `json:"deleted_by" gorm:"column:deleted_by"`
}

func (WlDevices) TableName() string {
	return "wl_devices"
}
