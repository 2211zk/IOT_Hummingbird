package wl_driver

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_driver"
	wl_driverReq "github.com/flipped-aurora/gin-vue-admin/server/model/wl_driver/request"
)

type WlDriverMarketService struct{}

func (s *WlDriverMarketService) CreateWlDriverMarket(ctx context.Context, data *wl_driver.WlDriverMarket) (err error) {
	err = global.GVA_DB.Create(data).Error
	return err
}

func (s *WlDriverMarketService) DeleteWlDriverMarket(ctx context.Context, marketId string) (err error) {
	err = global.GVA_DB.Delete(&wl_driver.WlDriverMarket{}, "market_id = ?", marketId).Error
	return err
}

func (s *WlDriverMarketService) DeleteWlDriverMarketByIds(ctx context.Context, marketIds []string) (err error) {
	err = global.GVA_DB.Where("market_id in ?", marketIds).Delete(&wl_driver.WlDriverMarket{}).Error
	return err
}

func (s *WlDriverMarketService) UpdateWlDriverMarket(ctx context.Context, data wl_driver.WlDriverMarket) (err error) {
	err = global.GVA_DB.Model(&wl_driver.WlDriverMarket{}).Where("market_id = ?", data.MarketId).Updates(&data).Error
	return err
}

func (s *WlDriverMarketService) GetWlDriverMarket(ctx context.Context, marketId string) (result wl_driver.WlDriverMarket, err error) {
	err = global.GVA_DB.Where("market_id = ?", marketId).First(&result).Error
	return
}

func (s *WlDriverMarketService) GetWlDriverMarketInfoList(ctx context.Context, info wl_driverReq.WlDriverMarketSearch) (list []wl_driver.WlDriverMarket, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&wl_driver.WlDriverMarket{})
	var results []wl_driver.WlDriverMarket
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_time BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
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
