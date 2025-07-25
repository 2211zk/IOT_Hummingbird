package wl_user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_user"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_user/request"
)

type WlUserService struct{}

var WlUserServiceApp = new(WlUserService)

func (s *WlUserService) GetWlUserList(req request.WlUserSearch) (list []wl_user.WlUser, total int64, err error) {
	db := global.GVA_DB.Model(&wl_user.WlUser{})
	if req.UserName != "" {
		db = db.Where("user_name LIKE ?", "%"+req.UserName+"%")
	}
	if req.Mobile != "" {
		db = db.Where("mobile = ?", req.Mobile)
	}
	if req.Status != "" {
		db = db.Where("user_status = ?", req.Status)
	}
	if req.Department != 0 {
		db = db.Where("department = ?", req.Department)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize).Order("creation_time desc").Find(&list).Error
	return
}
