
package wl_playform

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform"
    wl_playformReq "github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform/request"
)

type WlUserService struct {}
// CreateWlUser 创建wlUser表记录
// Author [yourname](https://github.com/yourname)
func (wlUserService *WlUserService) CreateWlUser(ctx context.Context, wlUser *wl_playform.WlUser) (err error) {
	err = global.GVA_DB.Create(wlUser).Error
	return err
}

// DeleteWlUser 删除wlUser表记录
// Author [yourname](https://github.com/yourname)
func (wlUserService *WlUserService)DeleteWlUser(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&wl_playform.WlUser{},"id = ?",id).Error
	return err
}

// DeleteWlUserByIds 批量删除wlUser表记录
// Author [yourname](https://github.com/yourname)
func (wlUserService *WlUserService)DeleteWlUserByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]wl_playform.WlUser{},"id in ?",ids).Error
	return err
}

// UpdateWlUser 更新wlUser表记录
// Author [yourname](https://github.com/yourname)
func (wlUserService *WlUserService)UpdateWlUser(ctx context.Context, wlUser wl_playform.WlUser) (err error) {
	err = global.GVA_DB.Model(&wl_playform.WlUser{}).Where("id = ?",wlUser.Id).Updates(&wlUser).Error
	return err
}

// GetWlUser 根据id获取wlUser表记录
// Author [yourname](https://github.com/yourname)
func (wlUserService *WlUserService)GetWlUser(ctx context.Context, id string) (wlUser wl_playform.WlUser, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&wlUser).Error
	return
}
// GetWlUserInfoList 分页获取wlUser表记录
// Author [yourname](https://github.com/yourname)
func (wlUserService *WlUserService)GetWlUserInfoList(ctx context.Context, info wl_playformReq.WlUserSearch) (list []wl_playform.WlUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&wl_playform.WlUser{})
    var wlUsers []wl_playform.WlUser
    // 如果有条件搜索 下方会自动创建搜索语句
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&wlUsers).Error
	return  wlUsers, total, err
}
func (wlUserService *WlUserService)GetWlUserPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
