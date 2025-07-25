
package wl_playform

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform"
    wl_playformReq "github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform/request"
    "gorm.io/gorm"
)

type WlAlarmService struct {}
// CreateWlAlarm 创建wlAlarm表记录
// Author [yourname](https://github.com/yourname)
func (wlAlarmService *WlAlarmService) CreateWlAlarm(ctx context.Context, wlAlarm *wl_playform.WlAlarm) (err error) {
	err = global.GVA_DB.Create(wlAlarm).Error
	return err
}

// DeleteWlAlarm 删除wlAlarm表记录
// Author [yourname](https://github.com/yourname)
func (wlAlarmService *WlAlarmService)DeleteWlAlarm(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&wl_playform.WlAlarm{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&wl_playform.WlAlarm{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteWlAlarmByIds 批量删除wlAlarm表记录
// Author [yourname](https://github.com/yourname)
func (wlAlarmService *WlAlarmService)DeleteWlAlarmByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&wl_playform.WlAlarm{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&wl_playform.WlAlarm{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateWlAlarm 更新wlAlarm表记录
// Author [yourname](https://github.com/yourname)
func (wlAlarmService *WlAlarmService)UpdateWlAlarm(ctx context.Context, wlAlarm wl_playform.WlAlarm) (err error) {
	err = global.GVA_DB.Model(&wl_playform.WlAlarm{}).Where("id = ?",wlAlarm.ID).Updates(&wlAlarm).Error
	return err
}

// GetWlAlarm 根据ID获取wlAlarm表记录
// Author [yourname](https://github.com/yourname)
func (wlAlarmService *WlAlarmService)GetWlAlarm(ctx context.Context, ID string) (wlAlarm wl_playform.WlAlarm, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wlAlarm).Error
	return
}
// GetWlAlarmInfoList 分页获取wlAlarm表记录
// Author [yourname](https://github.com/yourname)
func (wlAlarmService *WlAlarmService)GetWlAlarmInfoList(ctx context.Context, info wl_playformReq.WlAlarmSearch) (list []wl_playform.WlAlarm, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&wl_playform.WlAlarm{})
    var wlAlarms []wl_playform.WlAlarm
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
           orderMap["ID"] = true
           orderMap["CreatedAt"] = true
         	orderMap["create_time"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&wlAlarms).Error
	return  wlAlarms, total, err
}
func (wlAlarmService *WlAlarmService)GetWlAlarmPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
