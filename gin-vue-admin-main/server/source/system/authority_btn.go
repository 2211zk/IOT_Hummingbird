package system

import (
	"context"
	"fmt"

	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderAuthorityBtn = initOrderAuthority + 1

type initAuthorityBtn struct{}

// auto run
func init() {
	system.RegisterInit(initOrderAuthorityBtn, &initAuthorityBtn{})
}

func (i *initAuthorityBtn) MigrateTable(ctx context.Context) (context.Context, error) {
	return ctx, nil // do nothing
}

func (i *initAuthorityBtn) TableCreated(ctx context.Context) bool {
	return false // always replace
}

func (i *initAuthorityBtn) InitializerName() string {
	return "sys_authority_btn"
}

func (i *initAuthorityBtn) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	// 获取所有角色
	var authorities []sysModel.SysAuthority
	if err = db.Find(&authorities).Error; err != nil {
		return ctx, errors.Wrap(err, "获取角色列表失败")
	}

	// 获取wlResources菜单
	var wlResourcesMenu sysModel.SysBaseMenu
	if err = db.Where("name = ?", "wlResources").First(&wlResourcesMenu).Error; err != nil {
		return ctx, errors.Wrap(err, "查找wlResources菜单失败")
	}

	// 获取wlResources菜单的所有按钮
	var wlResourcesBtns []sysModel.SysBaseMenuBtn
	if err = db.Where("sys_base_menu_id = ?", wlResourcesMenu.ID).Find(&wlResourcesBtns).Error; err != nil {
		return ctx, errors.Wrap(err, "获取wlResources按钮失败")
	}

	// 为每个角色分配wlResources菜单的按钮权限
	var allAuthorityBtns []sysModel.SysAuthorityBtn

	for _, authority := range authorities {
		var authorityBtns []sysModel.SysAuthorityBtn

		// 为每个按钮创建权限记录
		for _, btn := range wlResourcesBtns {
			authorityBtns = append(authorityBtns, sysModel.SysAuthorityBtn{
				AuthorityId:      authority.AuthorityId,
				SysMenuID:        wlResourcesMenu.ID,
				SysBaseMenuBtnID: btn.ID,
			})
		}

		// 删除旧的按钮权限记录
		if err = db.Where("authority_id = ? AND sys_menu_id = ?", authority.AuthorityId, wlResourcesMenu.ID).Delete(&sysModel.SysAuthorityBtn{}).Error; err != nil {
			return ctx, errors.Wrap(err, "删除旧按钮权限失败")
		}

		// 创建新的按钮权限记录
		if err = db.Create(&authorityBtns).Error; err != nil {
			return ctx, errors.Wrap(err, "创建按钮权限失败")
		}

		// 收集所有按钮权限记录
		allAuthorityBtns = append(allAuthorityBtns, authorityBtns...)

		fmt.Printf("为角色 %s 分配了 %d 个按钮权限\n", authority.AuthorityName, len(authorityBtns))
	}

	next = context.WithValue(ctx, i.InitializerName(), allAuthorityBtns)
	return next, nil
}

func (i *initAuthorityBtn) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	// 检查是否存在按钮权限记录
	var count int64
	if err := db.Model(&sysModel.SysAuthorityBtn{}).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}
