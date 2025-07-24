package wl_driver

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_driver"
	wl_driverReq "github.com/flipped-aurora/gin-vue-admin/server/model/wl_driver/request"
)

type WlDevicesService struct{}

func (s *WlDevicesService) CreateWlDevices(ctx context.Context, data *wl_driver.WlDevices) (err error) {
	err = global.GVA_DB.Create(data).Error
	return err
}

func (s *WlDevicesService) DeleteWlDevices(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&wl_driver.WlDevices{}, "id = ?", ID).Error
	return err
}

func (s *WlDevicesService) DeleteWlDevicesByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Where("id in ?", IDs).Delete(&wl_driver.WlDevices{}).Error
	return err
}

func (s *WlDevicesService) UpdateWlDevices(ctx context.Context, data wl_driver.WlDevices) (err error) {
	err = global.GVA_DB.Model(&wl_driver.WlDevices{}).Where("id = ?", data.ID).Updates(&data).Error
	return err
}

func (s *WlDevicesService) GetWlDevices(ctx context.Context, ID string) (result wl_driver.WlDevices, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&result).Error
	return
}

func (s *WlDevicesService) GetWlDevicesInfoList(ctx context.Context, info wl_driverReq.WlDevicesSearch) (list []wl_driver.WlDevices, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&wl_driver.WlDevices{})
	var results []wl_driver.WlDevices
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Find(&results).Error
	return results, total, err
}
