
package wl_playform

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform"
    wl_playformReq "github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform/request"
    "gorm.io/gorm"
)

type WlCategoryService struct {}
// CreateWlCategory 创建wlCategory表记录
// Author [yourname](https://github.com/yourname)
func (wlCategoryService *WlCategoryService) CreateWlCategory(ctx context.Context, wlCategory *wl_playform.WlCategory) (err error) {
	err = global.GVA_DB.Create(wlCategory).Error
	return err
}

// DeleteWlCategory 删除wlCategory表记录
// Author [yourname](https://github.com/yourname)
func (wlCategoryService *WlCategoryService)DeleteWlCategory(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&wl_playform.WlCategory{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&wl_playform.WlCategory{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteWlCategoryByIds 批量删除wlCategory表记录
// Author [yourname](https://github.com/yourname)
func (wlCategoryService *WlCategoryService)DeleteWlCategoryByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&wl_playform.WlCategory{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&wl_playform.WlCategory{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateWlCategory 更新wlCategory表记录
// Author [yourname](https://github.com/yourname)
func (wlCategoryService *WlCategoryService)UpdateWlCategory(ctx context.Context, wlCategory wl_playform.WlCategory) (err error) {
	err = global.GVA_DB.Model(&wl_playform.WlCategory{}).Where("id = ?",wlCategory.ID).Updates(&wlCategory).Error
	return err
}

// GetWlCategory 根据ID获取wlCategory表记录
// Author [yourname](https://github.com/yourname)
func (wlCategoryService *WlCategoryService)GetWlCategory(ctx context.Context, ID string) (wlCategory wl_playform.WlCategory, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wlCategory).Error
	return
}
// GetWlCategoryInfoList 分页获取wlCategory表记录
// Author [yourname](https://github.com/yourname)
func (wlCategoryService *WlCategoryService)GetWlCategoryInfoList(ctx context.Context, info wl_playformReq.WlCategorySearch) (list []wl_playform.WlCategory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&wl_playform.WlCategory{})
    var wlCategorys []wl_playform.WlCategory
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    // 添加品类名称搜索
    if info.Name != nil && *info.Name != "" {
        db = db.Where("ca_name LIKE ?", "%"+*info.Name+"%")
    }
    if info.CaName != nil && *info.CaName != "" {
        db = db.Where("ca_name LIKE ?", "%"+*info.CaName+"%")
    }
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&wlCategorys).Error
	return  wlCategorys, total, err
}
func (wlCategoryService *WlCategoryService)GetWlCategoryPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
