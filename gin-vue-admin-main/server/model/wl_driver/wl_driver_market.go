package wl_driver

import (
	"time"
)

type WlDriverMarket struct {
	MarketId      string     `json:"market_id" gorm:"primaryKey;column:market_id;size:20;comment:市场ID"`
	DriverName    *string    `json:"driver_name" gorm:"column:driver_name;size:100;comment:驱动名称"`
	Version       *string    `json:"version" gorm:"column:version;size:20;comment:版本"`
	DriverType    *string    `json:"driver_type" gorm:"column:driver_type;type:enum('官方','自定义');comment:驱动类型"`
	DownloadCount *int       `json:"download_count" gorm:"column:download_count;default:0;comment:下载次数"`
	CreatedTime   *time.Time `json:"created_time" gorm:"column:created_time;comment:上架时间"`
	UpdatedTime   *time.Time `json:"updated_time" gorm:"column:updated_time;comment:更新时间"`
}

func (WlDriverMarket) TableName() string {
	return "wl_driver_market"
}
