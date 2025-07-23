
package wl_playform

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform"
    wl_playformReq "github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform/request"
    "gorm.io/gorm"
)

type WlEngineRulesService struct {}
// CreateWlEngineRules 创建wlEngineRules表记录
// Author [yourname](https://github.com/yourname)
func (wlEngineRulesService *WlEngineRulesService) CreateWlEngineRules(ctx context.Context, wlEngineRules *wl_playform.WlEngineRules) (err error) {
	err = global.GVA_DB.Create(wlEngineRules).Error
	return err
}

// DeleteWlEngineRules 删除wlEngineRules表记录
// Author [yourname](https://github.com/yourname)
func (wlEngineRulesService *WlEngineRulesService)DeleteWlEngineRules(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&wl_playform.WlEngineRules{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&wl_playform.WlEngineRules{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteWlEngineRulesByIds 批量删除wlEngineRules表记录
// Author [yourname](https://github.com/yourname)
func (wlEngineRulesService *WlEngineRulesService)DeleteWlEngineRulesByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&wl_playform.WlEngineRules{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&wl_playform.WlEngineRules{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateWlEngineRules 更新wlEngineRules表记录
// Author [yourname](https://github.com/yourname)
func (wlEngineRulesService *WlEngineRulesService)UpdateWlEngineRules(ctx context.Context, wlEngineRules wl_playform.WlEngineRules) (err error) {
	err = global.GVA_DB.Model(&wl_playform.WlEngineRules{}).Where("id = ?",wlEngineRules.ID).Updates(&wlEngineRules).Error
	return err
}

// GetWlEngineRules 根据ID获取wlEngineRules表记录
// Author [yourname](https://github.com/yourname)
func (wlEngineRulesService *WlEngineRulesService)GetWlEngineRules(ctx context.Context, ID string) (wlEngineRules wl_playform.WlEngineRules, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wlEngineRules).Error
	return
}
// GetWlEngineRulesInfoList 分页获取wlEngineRules表记录
// Author [yourname](https://github.com/yourname)
func (wlEngineRulesService *WlEngineRulesService)GetWlEngineRulesInfoList(ctx context.Context, info wl_playformReq.WlEngineRulesSearch) (list []wl_playform.WlEngineRules, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&wl_playform.WlEngineRules{})
    var wlEngineRuless []wl_playform.WlEngineRules
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.RuleName != nil && *info.RuleName != "" {
        db = db.Where("rule_name LIKE ?", "%"+ *info.RuleName+"%")
    }
    if info.MessageSource != nil && *info.MessageSource != "" {
        db = db.Where("message_source = ?", *info.MessageSource)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&wlEngineRuless).Error
	return  wlEngineRuless, total, err
}
func (wlEngineRulesService *WlEngineRulesService)GetWlEngineRulesPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
