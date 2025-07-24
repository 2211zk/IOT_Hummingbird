package wl_driver

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_driver"
	wl_driverReq "github.com/flipped-aurora/gin-vue-admin/server/model/wl_driver/request"
	"gorm.io/gorm"
)

type WlDriversService struct{}

// CreateWlDrivers 创建wlDrivers表记录
// Author [yourname](https://github.com/yourname)
func (wlDriversService *WlDriversService) CreateWlDrivers(ctx context.Context, wlDrivers *wl_driver.WlDrivers) (err error) {
	err = global.GVA_DB.Create(wlDrivers).Error
	return err
}

// DeleteWlDrivers 删除wlDrivers表记录
// Author [yourname](https://github.com/yourname)
func (wlDriversService *WlDriversService) DeleteWlDrivers(ctx context.Context, ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&wl_driver.WlDrivers{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&wl_driver.WlDrivers{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteWlDriversByIds 批量删除wlDrivers表记录
// Author [yourname](https://github.com/yourname)
func (wlDriversService *WlDriversService) DeleteWlDriversByIds(ctx context.Context, IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&wl_driver.WlDrivers{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&wl_driver.WlDrivers{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateWlDrivers 更新wlDrivers表记录
// Author [yourname](https://github.com/yourname)
func (wlDriversService *WlDriversService) UpdateWlDrivers(ctx context.Context, wlDrivers wl_driver.WlDrivers) (err error) {
	err = global.GVA_DB.Model(&wl_driver.WlDrivers{}).Where("id = ?", wlDrivers.ID).Updates(&wlDrivers).Error
	return err
}

// GetWlDrivers 根据ID获取wlDrivers表记录
// Author [yourname](https://github.com/yourname)
func (wlDriversService *WlDriversService) GetWlDrivers(ctx context.Context, ID string) (wlDrivers wl_driver.WlDrivers, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wlDrivers).Error
	return
}

// GetWlDriversInfoList 分页获取wlDrivers表记录
// Author [yourname](https://github.com/yourname)
func (wlDriversService *WlDriversService) GetWlDriversInfoList(ctx context.Context, info wl_driverReq.WlDriversSearch) (list []wl_driver.WlDrivers, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&wl_driver.WlDrivers{})
	var wlDriverss []wl_driver.WlDrivers
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	// 添加驱动名称搜索条件
	if info.DriverName != "" {
		db = db.Where("driver_name LIKE ?", "%"+info.DriverName+"%")
	}

	// 添加驱动类型搜索条件
	if info.DriverType != "" {
		db = db.Where("driver_type = ?", info.DriverType)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&wlDriverss).Error
	return wlDriverss, total, err
}
func (wlDriversService *WlDriversService) GetWlDriversPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
