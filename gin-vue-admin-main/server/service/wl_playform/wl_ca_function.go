
package wl_playform

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform"
    wl_playformReq "github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform/request"
    "gorm.io/gorm"
)

type WlCaFunctionService struct {}
// CreateWlCaFunction 创建wlCaFunction表记录
// Author [yourname](https://github.com/yourname)
func (wlCaFunctionService *WlCaFunctionService) CreateWlCaFunction(ctx context.Context, wlCaFunction *wl_playform.WlCaFunction) (err error) {
	err = global.GVA_DB.Create(wlCaFunction).Error
	return err
}

// DeleteWlCaFunction 删除wlCaFunction表记录
// Author [yourname](https://github.com/yourname)
func (wlCaFunctionService *WlCaFunctionService)DeleteWlCaFunction(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&wl_playform.WlCaFunction{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&wl_playform.WlCaFunction{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteWlCaFunctionByIds 批量删除wlCaFunction表记录
// Author [yourname](https://github.com/yourname)
func (wlCaFunctionService *WlCaFunctionService)DeleteWlCaFunctionByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&wl_playform.WlCaFunction{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&wl_playform.WlCaFunction{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateWlCaFunction 更新wlCaFunction表记录
// Author [yourname](https://github.com/yourname)
func (wlCaFunctionService *WlCaFunctionService)UpdateWlCaFunction(ctx context.Context, wlCaFunction wl_playform.WlCaFunction) (err error) {
	err = global.GVA_DB.Model(&wl_playform.WlCaFunction{}).Where("id = ?",wlCaFunction.ID).Updates(&wlCaFunction).Error
	return err
}

// GetWlCaFunction 根据ID获取wlCaFunction表记录
// Author [yourname](https://github.com/yourname)
func (wlCaFunctionService *WlCaFunctionService)GetWlCaFunction(ctx context.Context, ID string) (wlCaFunction wl_playform.WlCaFunction, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wlCaFunction).Error
	return
}
// GetWlCaFunctionInfoList 分页获取wlCaFunction表记录
// Author [yourname](https://github.com/yourname)
func (wlCaFunctionService *WlCaFunctionService)GetWlCaFunctionInfoList(ctx context.Context, info wl_playformReq.WlCaFunctionSearch) (list []wl_playform.WlCaFunction, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&wl_playform.WlCaFunction{})
    var wlCaFunctions []wl_playform.WlCaFunction
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    // 添加品类ID过滤
    if info.CaId != nil {
        db = db.Where("ca_id = ?", *info.CaId)
    }
    
    // 添加功能名称过滤
    if info.FunctionName != nil && *info.FunctionName != "" {
        db = db.Where("function_name LIKE ?", "%"+*info.FunctionName+"%")
    }
    
    // 添加功能类型过滤
    if info.FunctionType != nil && *info.FunctionType != "" {
        db = db.Where("function_type = ?", *info.FunctionType)
    }
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&wlCaFunctions).Error
	return  wlCaFunctions, total, err
}
func (wlCaFunctionService *WlCaFunctionService)GetWlCaFunctionPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
