package wl_driver

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_driver"
	wl_driverReq "github.com/flipped-aurora/gin-vue-admin/server/model/wl_driver/request"
	"gorm.io/gorm"
)

type WlProtocolsService struct{}

func (s *WlProtocolsService) CreateWlProtocols(ctx context.Context, data *wl_driver.WlProtocols) (err error) {
	err = global.GVA_DB.Create(data).Error
	return err
}

func (s *WlProtocolsService) DeleteWlProtocols(ctx context.Context, ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&wl_driver.WlProtocols{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&wl_driver.WlProtocols{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (s *WlProtocolsService) DeleteWlProtocolsByIds(ctx context.Context, IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&wl_driver.WlProtocols{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&wl_driver.WlProtocols{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (s *WlProtocolsService) UpdateWlProtocols(ctx context.Context, data wl_driver.WlProtocols) (err error) {
	err = global.GVA_DB.Model(&wl_driver.WlProtocols{}).Where("id = ?", data.ID).Updates(&data).Error
	return err
}

func (s *WlProtocolsService) GetWlProtocols(ctx context.Context, ID string) (result wl_driver.WlProtocols, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&result).Error
	return
}

func (s *WlProtocolsService) GetWlProtocolsInfoList(ctx context.Context, info wl_driverReq.WlProtocolsSearch) (list []wl_driver.WlProtocols, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&wl_driver.WlProtocols{})
	var results []wl_driver.WlProtocols
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
