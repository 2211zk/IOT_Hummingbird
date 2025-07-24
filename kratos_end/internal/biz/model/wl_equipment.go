package model

import "time"

type WlEquipment struct {
	Id           int32     `gorm:"column:id;type:int;comment:主键id;primaryKey;not null;" json:"id"`                        // 主键id
	EqName       string    `gorm:"column:eq_name;type:varchar(20);comment:设备名称;default:NULL;" json:"eq_name"`             // 设备名称
	EqLogotype   string    `gorm:"column:eq_logotype;type:varchar(50);comment:设备唯一标识;default:NULL;" json:"eq_logotype"`   // 设备唯一标识
	ProductsId   int16     `gorm:"column:products_id;type:smallint;comment:所属产品;default:NULL;" json:"products_id"`        // 所属产品
	DriveId      int16     `gorm:"column:drive_id;type:smallint;comment:驱动id;default:NULL;" json:"drive_id"`              // 驱动id
	EqCoordinate string    `gorm:"column:eq_coordinate;type:varchar(50);comment:设备坐标;default:NULL;" json:"eq_coordinate"` // 设备坐标
	EqAddress    string    `gorm:"column:eq_address;type:varchar(100);comment:设备详细地址;default:NULL;" json:"eq_address"`    // 设备详细地址
	EqInfo       string    `gorm:"column:eq_info;type:varchar(150);comment:设备描述;default:NULL;" json:"eq_info"`            // 设备描述
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime(3);default:NULL;" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime(3);default:NULL;" json:"updated_at"`
	DeletedAt    time.Time `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
	CreatedBy    uint64    `gorm:"column:created_by;type:bigint UNSIGNED;comment:创建者;default:NULL;" json:"created_by"` // 创建者
	UpdatedBy    uint64    `gorm:"column:updated_by;type:bigint UNSIGNED;comment:更新者;default:NULL;" json:"updated_by"` // 更新者
	DeletedBy    uint64    `gorm:"column:deleted_by;type:bigint UNSIGNED;comment:删除者;default:NULL;" json:"deleted_by"` // 删除者
}

func (WlEquipment) TableName() string {
	return "wl_equipment"
}
