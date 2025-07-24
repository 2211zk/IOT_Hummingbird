package wl_driver

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_driver"
	wl_driverReq "github.com/flipped-aurora/gin-vue-admin/server/model/wl_driver/request"
)

type WlSystemMoniorService struct{}

func (s *WlSystemMoniorService) CreateWlSystemMonior(ctx context.Context, data *wl_driver.WlSystemMonior) (err error) {
	err = global.GVA_DB.Create(data).Error
	return err
}

func (s *WlSystemMoniorService) DeleteWlSystemMonior(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&wl_driver.WlSystemMonior{}, "id = ?", ID).Error
	return err
}

func (s *WlSystemMoniorService) DeleteWlSystemMoniorByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Where("id in ?", IDs).Delete(&wl_driver.WlSystemMonior{}).Error
	return err
}

func (s *WlSystemMoniorService) UpdateWlSystemMonior(ctx context.Context, data wl_driver.WlSystemMonior) (err error) {
	err = global.GVA_DB.Model(&wl_driver.WlSystemMonior{}).Where("id = ?", data.ID).Updates(&data).Error
	return err
}

func (s *WlSystemMoniorService) GetWlSystemMonior(ctx context.Context, ID string) (result wl_driver.WlSystemMonior, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&result).Error
	return
}

func (s *WlSystemMoniorService) GetWlSystemMoniorInfoList(ctx context.Context, info wl_driverReq.WlSystemMoniorSearch) (list []wl_driver.WlSystemMonior, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&wl_driver.WlSystemMonior{})
	var results []wl_driver.WlSystemMonior
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("timestamp BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
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
