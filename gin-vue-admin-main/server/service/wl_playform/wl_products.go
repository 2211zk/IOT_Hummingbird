
package wl_playform

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform"
    wl_playformReq "github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform/request"
    "gorm.io/gorm"
)

type WlProductsService struct {}
// CreateWlProducts 创建wlProducts表记录
// Author [yourname](https://github.com/yourname)
func (wlProductsService *WlProductsService) CreateWlProducts(ctx context.Context, wlProducts *wl_playform.WlProducts) (err error) {
	err = global.GVA_DB.Create(wlProducts).Error
	return err
}

// DeleteWlProducts 删除wlProducts表记录
// Author [yourname](https://github.com/yourname)
func (wlProductsService *WlProductsService)DeleteWlProducts(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&wl_playform.WlProducts{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&wl_playform.WlProducts{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteWlProductsByIds 批量删除wlProducts表记录
// Author [yourname](https://github.com/yourname)
func (wlProductsService *WlProductsService)DeleteWlProductsByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&wl_playform.WlProducts{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&wl_playform.WlProducts{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateWlProducts 更新wlProducts表记录
// Author [yourname](https://github.com/yourname)
func (wlProductsService *WlProductsService)UpdateWlProducts(ctx context.Context, wlProducts wl_playform.WlProducts) (err error) {
	err = global.GVA_DB.Model(&wl_playform.WlProducts{}).Where("id = ?",wlProducts.ID).Updates(&wlProducts).Error
	return err
}

// GetWlProducts 根据ID获取wlProducts表记录
// Author [yourname](https://github.com/yourname)
func (wlProductsService *WlProductsService)GetWlProducts(ctx context.Context, ID string) (wlProducts wl_playform.WlProducts, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wlProducts).Error
	return
}
// GetWlProductsInfoList 分页获取wlProducts表记录
// Author [yourname](https://github.com/yourname)
func (wlProductsService *WlProductsService)GetWlProductsInfoList(ctx context.Context, info wl_playformReq.WlProductsSearch) (list []wl_playform.WlProducts, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&wl_playform.WlProducts{})
    var wlProductss []wl_playform.WlProducts
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.PrName != nil && *info.PrName != "" {
        db = db.Where("pr_name LIKE ?", "%"+ *info.PrName+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&wlProductss).Error
	return  wlProductss, total, err
}
func (wlProductsService *WlProductsService)GetWlProductsPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
