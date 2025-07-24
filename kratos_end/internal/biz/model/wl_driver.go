package model

import "time"

type WlDrivers struct {
	Id             uint64    `gorm:"column:id;type:bigint UNSIGNED;comment:驱动编号;primaryKey;not null;" json:"id"`                // 驱动编号
	DriverName     string    `gorm:"column:driver_name;type:varchar(100);comment:驱动名称;default:NULL;" json:"driver_name"`        // 驱动名称
	Version        string    `gorm:"column:version;type:varchar(20);comment:版本;default:NULL;" json:"version"`                     // 版本
	DriverType     string    `gorm:"column:driver_type;type:varchar(1);comment:当前状态;default:NULL;" json:"driver_type"`          // 当前状态
	Status         string    `gorm:"column:status;type:varchar(1);comment:状态;default:NULL;" json:"status"`                        // 状态
	CreatedTime    time.Time `gorm:"column:created_time;type:datetime(3);comment:创建时间;default:NULL;" json:"created_time"`       // 创建时间
	UpdateTime     time.Time `gorm:"column:update_time;type:datetime(3);comment:更新时间;default:NULL;" json:"update_time"`         // 更新时间
	ProtocolType   string    `gorm:"column:protocol_type;type:varchar(50);comment:协议类型;default:NULL;" json:"protocol_type"`     // 协议类型
	DeviceCategory string    `gorm:"column:device_category;type:varchar(50);comment:设备类型;default:NULL;" json:"device_category"` // 设备类型
	CreatedAt      time.Time `gorm:"column:created_at;type:datetime(3);default:NULL;" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:datetime(3);default:NULL;" json:"updated_at"`
	DeletedAt      time.Time `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
	DriverId       string    `gorm:"column:driver_id;type:varchar(20);comment:驱动编号;default:NULL;" json:"driver_id"`     // 驱动编号
	CreatedBy      uint64    `gorm:"column:created_by;type:bigint UNSIGNED;comment:创建者;default:NULL;" json:"created_by"` // 创建者
	UpdatedBy      uint64    `gorm:"column:updated_by;type:bigint UNSIGNED;comment:更新者;default:NULL;" json:"updated_by"` // 更新者
	DeletedBy      uint64    `gorm:"column:deleted_by;type:bigint UNSIGNED;comment:删除者;default:NULL;" json:"deleted_by"` // 删除者
}

func (WlDrivers) TableName() string {
	return "wl_drivers"
}
