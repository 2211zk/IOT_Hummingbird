
package wl_playform

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform"
    wl_playformReq "github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform/request"
    "gorm.io/gorm"
)

type WlScenesService struct {}
// CreateWlScenes 创建wlScenes表记录
// Author [yourname](https://github.com/yourname)
func (wlScenesService *WlScenesService) CreateWlScenes(ctx context.Context, wlScenes *wl_playform.WlScenes) (err error) {
	err = global.GVA_DB.Create(wlScenes).Error
	return err
}

// DeleteWlScenes 删除wlScenes表记录
// Author [yourname](https://github.com/yourname)
func (wlScenesService *WlScenesService)DeleteWlScenes(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&wl_playform.WlScenes{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&wl_playform.WlScenes{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteWlScenesByIds 批量删除wlScenes表记录
// Author [yourname](https://github.com/yourname)
func (wlScenesService *WlScenesService)DeleteWlScenesByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&wl_playform.WlScenes{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&wl_playform.WlScenes{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateWlScenes 更新wlScenes表记录
// Author [yourname](https://github.com/yourname)
func (wlScenesService *WlScenesService)UpdateWlScenes(ctx context.Context, wlScenes wl_playform.WlScenes) (err error) {
	err = global.GVA_DB.Model(&wl_playform.WlScenes{}).Where("id = ?",wlScenes.ID).Updates(&wlScenes).Error
	return err
}

// GetWlScenes 根据ID获取wlScenes表记录
// Author [yourname](https://github.com/yourname)
func (wlScenesService *WlScenesService)GetWlScenes(ctx context.Context, ID string) (wlScenes wl_playform.WlScenes, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wlScenes).Error
	return
}
// GetWlScenesInfoList 分页获取wlScenes表记录
// Author [yourname](https://github.com/yourname)
func (wlScenesService *WlScenesService)GetWlScenesInfoList(ctx context.Context, info wl_playformReq.WlScenesSearch) (list []wl_playform.WlScenes, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&wl_playform.WlScenes{})
    var wlSceness []wl_playform.WlScenes
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.SceneName != nil && *info.SceneName != "" {
        db = db.Where("scene_name LIKE ?", "%"+ *info.SceneName+"%")
    }
    if info.ScenesStatus != nil && *info.ScenesStatus != "" {
        db = db.Where("scenes_status = ?", *info.ScenesStatus)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&wlSceness).Error
	return  wlSceness, total, err
}
func (wlScenesService *WlScenesService)GetWlScenesPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
