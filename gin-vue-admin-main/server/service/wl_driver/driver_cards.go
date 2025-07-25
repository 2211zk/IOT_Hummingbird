package wl_driver

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_driver"
	wl_driverReq "github.com/flipped-aurora/gin-vue-admin/server/model/wl_driver/request"
	"gorm.io/gorm"
)

type DriverCardsService struct{}

func (driverCardsService *DriverCardsService) CreateDriverCards(ctx context.Context, driverCards *wl_driver.DriverCards) (err error) {
	err = global.GVA_DB.Create(driverCards).Error
	return err
}

func (driverCardsService *DriverCardsService) DeleteDriverCards(ctx context.Context, ID uint, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&wl_driver.DriverCards{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&wl_driver.DriverCards{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (driverCardsService *DriverCardsService) DeleteDriverCardsByIds(ctx context.Context, IDs []uint, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&wl_driver.DriverCards{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&wl_driver.DriverCards{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (driverCardsService *DriverCardsService) UpdateDriverCards(ctx context.Context, driverCards wl_driver.DriverCards) (err error) {
	err = global.GVA_DB.Model(&wl_driver.DriverCards{}).Where("id = ?", driverCards.ID).Updates(&driverCards).Error
	return err
}

func (driverCardsService *DriverCardsService) GetDriverCards(ctx context.Context, ID uint) (driverCards wl_driver.DriverCards, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&driverCards).Error
	return
}

func (driverCardsService *DriverCardsService) GetDriverCardsInfoList(ctx context.Context, info wl_driverReq.DriverCardsSearch) (list []wl_driver.DriverCards, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&wl_driver.DriverCards{})
	var driverCardss []wl_driver.DriverCards
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
	err = db.Find(&driverCardss).Error
	return driverCardss, total, err
}

func (driverCardsService *DriverCardsService) GetDriverCardsPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
