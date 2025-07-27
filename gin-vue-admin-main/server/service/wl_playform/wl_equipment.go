package wl_playform

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform"
	wl_playformReq "github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform/request"
	"gorm.io/gorm"
)

type WlEquipmentService struct{}

// CreateWlEquipment 创建wlEquipment表记录
// Author [yourname](https://github.com/yourname)
func (wlEquipmentService *WlEquipmentService) CreateWlEquipment(ctx context.Context, wlEquipment *wl_playform.WlEquipment) (err error) {
	err = global.GVA_DB.Create(wlEquipment).Error
	return err
}

// DeleteWlEquipment 删除wlEquipment表记录
// Author [yourname](https://github.com/yourname)
func (wlEquipmentService *WlEquipmentService) DeleteWlEquipment(ctx context.Context, ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&wl_playform.WlEquipment{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&wl_playform.WlEquipment{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteWlEquipmentByIds 批量删除wlEquipment表记录
// Author [yourname](https://github.com/yourname)
func (wlEquipmentService *WlEquipmentService) DeleteWlEquipmentByIds(ctx context.Context, IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&wl_playform.WlEquipment{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&wl_playform.WlEquipment{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateWlEquipment 更新wlEquipment表记录
// Author [yourname](https://github.com/yourname)
func (wlEquipmentService *WlEquipmentService) UpdateWlEquipment(ctx context.Context, wlEquipment wl_playform.WlEquipment) (err error) {
	err = global.GVA_DB.Model(&wl_playform.WlEquipment{}).Where("id = ?", wlEquipment.ID).Updates(&wlEquipment).Error
	return err
}

// GetWlEquipment 根据ID获取wlEquipment表记录
// Author [yourname](https://github.com/yourname)
func (wlEquipmentService *WlEquipmentService) GetWlEquipment(ctx context.Context, ID string) (wlEquipment wl_playform.WlEquipment, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wlEquipment).Error
	return
}

// GetWlEquipmentInfoList 分页获取wlEquipment表记录
// Author [yourname](https://github.com/yourname)
func (wlEquipmentService *WlEquipmentService) GetWlEquipmentInfoList(ctx context.Context, info wl_playformReq.WlEquipmentSearch) (list []wl_playform.WlEquipment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&wl_playform.WlEquipment{})
	var wlEquipments []wl_playform.WlEquipment
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.EqName != nil && *info.EqName != "" {
		db = db.Where("eq_name LIKE ?", "%"+*info.EqName+"%")
	}
	if info.ProductsId != 0 {
		db = db.Where("products_id = ?", info.ProductsId)
	}
	if info.Status != nil && *info.Status != "" {
		db = db.Where("status = ?", *info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&wlEquipments).Error
	return wlEquipments, total, err
}
func (wlEquipmentService *WlEquipmentService) GetWlEquipmentPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
