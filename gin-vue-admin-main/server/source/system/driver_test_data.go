package system

import (
	"context"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_driver"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderDriverTestData = initOrderDictDetail + 1

type initDriverTestData struct{}

// auto run
func init() {
	system.RegisterInit(initOrderDriverTestData, &initDriverTestData{})
}

func (i *initDriverTestData) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&wl_driver.WlDrivers{})
}

func (i *initDriverTestData) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&wl_driver.WlDrivers{})
}

func (i *initDriverTestData) InitializerName() string {
	return "driver_test_data"
}

func (i *initDriverTestData) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	// 检查是否已有测试数据
	var count int64
	db.Model(&wl_driver.WlDrivers{}).Count(&count)
	if count > 0 {
		return ctx, nil // 已有数据，跳过初始化
	}

	now := time.Now()
	driverName1 := "MODBUS RTU协议驱动"
	driverName2 := "TCP协议驱动"
	driverId1 := "DRV001"
	driverId2 := "DRV002"
	version := "2.7"
	driverType1 := "official"
	driverType2 := "custom"
	status1 := "运行中"
	status2 := "已停止"
	protocolType1 := "MODBUS"
	protocolType2 := "TCP"
	deviceCategory1 := "网关"
	deviceCategory2 := "传感器"

	entities := []wl_driver.WlDrivers{
		{
			DriverName:     &driverName1,
			DriverId:       &driverId1,
			Version:        &version,
			DriverType:     &driverType1,
			Status:         &status1,
			ProtocolType:   &protocolType1,
			DeviceCategory: &deviceCategory1,
			CreatedTime:    &now,
			UpdateTime:     &now,
		},
		{
			DriverName:     &driverName2,
			DriverId:       &driverId2,
			Version:        &version,
			DriverType:     &driverType2,
			Status:         &status2,
			ProtocolType:   &protocolType2,
			DeviceCategory: &deviceCategory2,
			CreatedTime:    &now,
			UpdateTime:     &now,
		},
	}

	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, "驱动测试数据初始化失败!")
	}

	next = context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initDriverTestData) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	var count int64
	db.Model(&wl_driver.WlDrivers{}).Count(&count)
	return count > 0
}
