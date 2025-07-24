package biz

import (
	"context"

	"gorm.io/gorm"
)

type DashboardUsecase struct {
	db *gorm.DB
}

func NewDashboardUsecase(db *gorm.DB) *DashboardUsecase {
	return &DashboardUsecase{db: db}
}

func (uc *DashboardUsecase) GetOverview(ctx context.Context) (map[string]int64, error) {
	var productTotal, productPublished, productUnpublished int64
	var deviceTotal, deviceOnline, deviceOffline int64
	var driverTotal, driverRunning, driverStopped int64
	var alarmTotal int64

	uc.db.Table("wl_products").Count(&productTotal)
	uc.db.Table("wl_products").Where("standard_quality=1").Count(&productPublished)
	uc.db.Table("wl_products").Where("standard_quality!=1").Count(&productUnpublished)
	uc.db.Table("wl_equipment").Count(&deviceTotal)
	uc.db.Table("wl_equipment").Where("eq_status='在线'").Count(&deviceOnline)
	uc.db.Table("wl_equipment").Where("eq_status='离线'").Count(&deviceOffline)
	uc.db.Table("wl_drivers").Count(&driverTotal)
	uc.db.Table("wl_drivers").Where("status='1'").Count(&driverRunning)
	uc.db.Table("wl_drivers").Where("status!='1'").Count(&driverStopped)
	uc.db.Table("wl_alarm").Count(&alarmTotal)

	return map[string]int64{
		"product_total":       productTotal,
		"product_published":   productPublished,
		"product_unpublished": productUnpublished,
		"device_total":        deviceTotal,
		"device_online":       deviceOnline,
		"device_offline":      deviceOffline,
		"driver_total":        driverTotal,
		"driver_running":      driverRunning,
		"driver_stopped":      driverStopped,
		"alarm_total":         alarmTotal,
	}, nil
}

func (uc *DashboardUsecase) GetDB() *gorm.DB {
	return uc.db
}

// 其它接口同理实现...
