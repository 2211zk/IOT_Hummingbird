package system

import (
	"context"

	. "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderMenu = initOrderAuthority + 1

type initMenu struct{}

// auto run
func init() {
	system.RegisterInit(initOrderMenu, &initMenu{})
}

func (i *initMenu) InitializerName() string {
	return SysBaseMenu{}.TableName()
}

func (i *initMenu) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(
		&SysBaseMenu{},
		&SysBaseMenuParameter{},
		&SysBaseMenuBtn{},
	)
}

func (i *initMenu) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	m := db.Migrator()
	return m.HasTable(&SysBaseMenu{}) &&
		m.HasTable(&SysBaseMenuParameter{}) &&
		m.HasTable(&SysBaseMenuBtn{})
}

// InitializeData 初始化菜单数据
// 功能：创建系统所需的所有菜单项，包括父级菜单和子菜单
// 特别处理：为设备接入模块创建层级菜单结构（设备接入 -> 产品管理、设备管理）
func (i *initMenu) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	// 清理现有的菜单数据，重新创建
	if err = db.Exec("DELETE FROM sys_base_menus").Error; err != nil {
		return ctx, errors.Wrap(err, "清理现有菜单数据失败!")
	}

	// 定义所有菜单
	allMenus := []SysBaseMenu{
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "dashboard", Name: "dashboard", Component: "view/dashboard/index.vue", Sort: 1, Meta: Meta{Title: "仪表盘", Icon: "odometer"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "about", Name: "about", Component: "view/about/index.vue", Sort: 9, Meta: Meta{Title: "关于我们", Icon: "info-filled"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 3, Meta: Meta{Title: "超级管理员", Icon: "user"}},
		{MenuLevel: 0, Hidden: true, ParentId: 0, Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 4, Meta: Meta{Title: "个人信息", Icon: "message"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "example", Name: "example", Component: "view/example/index.vue", Sort: 7, Meta: Meta{Title: "示例文件", Icon: "management"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "systemTools", Name: "systemTools", Component: "view/systemTools/index.vue", Sort: 5, Meta: Meta{Title: "系统工具", Icon: "tools"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "https://www.gin-vue-admin.com", Name: "https://www.gin-vue-admin.com", Component: "/", Sort: 0, Meta: Meta{Title: "官方网站", Icon: "customer-gva"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "state", Name: "state", Component: "view/system/state.vue", Sort: 8, Meta: Meta{Title: "服务器状态", Icon: "cloudy"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "plugin", Name: "plugin", Component: "view/routerHolder.vue", Sort: 6, Meta: Meta{Title: "插件系统", Icon: "cherry"}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "advancedCapabilities", Name: "advancedCapabilities", Component: "view/routerHolder.vue", Sort: 10, Meta: Meta{Title: "高级能力", Icon: "cloud"}},
		// 设备接入父菜单 - 作为产品管理和设备管理的父级菜单
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "wl_playform", Name: "wl_playform", Component: "view/wl_playform/deviceAccess/index.vue", Sort: 2, Meta: Meta{Title: "设备接入", Icon: "connection"}},
	}

	// 先创建父级菜单（ParentId = 0 的菜单）
	if err = db.Create(&allMenus).Error; err != nil {
		return ctx, errors.Wrap(err, SysBaseMenu{}.TableName()+"父级菜单初始化失败!")
	}

	// 建立菜单映射 - 通过Name查找已创建的菜单及其ID
	// 目的：为后续创建子菜单时提供父菜单的ID引用
	// 使用场景：子菜单需要设置ParentId字段，指向对应的父菜单ID
	menuNameMap := make(map[string]uint)
	for _, menu := range allMenus {
		menuNameMap[menu.Name] = menu.ID
	}

	// 定义子菜单，并设置正确的ParentId
	// 注意：子菜单的ParentId必须指向已存在的父菜单ID，确保层级关系正确建立
	childMenus := []SysBaseMenu{
		// superAdmin子菜单
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["superAdmin"], Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1, Meta: Meta{Title: "角色管理", Icon: "avatar"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["superAdmin"], Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 2, Meta: Meta{Title: "菜单管理", Icon: "tickets", KeepAlive: true}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["superAdmin"], Path: "api", Name: "api", Component: "view/superAdmin/api/api.vue", Sort: 3, Meta: Meta{Title: "api管理", Icon: "platform", KeepAlive: true}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["superAdmin"], Path: "user", Name: "user", Component: "view/superAdmin/user/user.vue", Sort: 4, Meta: Meta{Title: "用户管理", Icon: "coordinate"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["superAdmin"], Path: "dictionary", Name: "dictionary", Component: "view/superAdmin/dictionary/sysDictionary.vue", Sort: 5, Meta: Meta{Title: "字典管理", Icon: "notebook"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["superAdmin"], Path: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 6, Meta: Meta{Title: "操作历史", Icon: "pie-chart"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["superAdmin"], Path: "sysParams", Name: "sysParams", Component: "view/superAdmin/params/sysParams.vue", Sort: 7, Meta: Meta{Title: "参数管理", Icon: "compass"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["superAdmin"], Path: "departmentManage", Name: "departmentManage", Component: "view/system/departmentManage/index.vue", Sort: 8, Meta: Meta{Title: "部门管理", Icon: "office-building"}},

		// example子菜单
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["example"], Path: "upload", Name: "upload", Component: "view/example/upload/upload.vue", Sort: 5, Meta: Meta{Title: "媒体库（上传下载）", Icon: "upload"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["example"], Path: "breakpoint", Name: "breakpoint", Component: "view/example/breakpoint/breakpoint.vue", Sort: 6, Meta: Meta{Title: "断点续传", Icon: "upload-filled"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["example"], Path: "customer", Name: "customer", Component: "view/example/customer/customer.vue", Sort: 7, Meta: Meta{Title: "客户列表（资源示例）", Icon: "avatar"}},

		// systemTools子菜单
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["systemTools"], Path: "autoCode", Name: "autoCode", Component: "view/systemTools/autoCode/index.vue", Sort: 1, Meta: Meta{Title: "代码生成器", Icon: "cpu", KeepAlive: true}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["systemTools"], Path: "formCreate", Name: "formCreate", Component: "view/systemTools/formCreate/index.vue", Sort: 3, Meta: Meta{Title: "表单生成器", Icon: "magic-stick", KeepAlive: true}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["systemTools"], Path: "system", Name: "system", Component: "view/systemTools/system/system.vue", Sort: 4, Meta: Meta{Title: "系统配置", Icon: "operation"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["systemTools"], Path: "autoCodeAdmin", Name: "autoCodeAdmin", Component: "view/systemTools/autoCodeAdmin/index.vue", Sort: 2, Meta: Meta{Title: "自动化代码管理", Icon: "magic-stick"}},
		{MenuLevel: 1, Hidden: true, ParentId: menuNameMap["systemTools"], Path: "autoCodeEdit/:id", Name: "autoCodeEdit", Component: "view/systemTools/autoCode/index.vue", Sort: 0, Meta: Meta{Title: "自动化代码-${id}", Icon: "magic-stick"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["systemTools"], Path: "autoPkg", Name: "autoPkg", Component: "view/systemTools/autoPkg/autoPkg.vue", Sort: 0, Meta: Meta{Title: "模板配置", Icon: "folder"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["systemTools"], Path: "exportTemplate", Name: "exportTemplate", Component: "view/systemTools/exportTemplate/exportTemplate.vue", Sort: 5, Meta: Meta{Title: "导出模板", Icon: "reading"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["systemTools"], Path: "picture", Name: "picture", Component: "view/systemTools/autoCode/picture.vue", Sort: 6, Meta: Meta{Title: "AI页面绘制", Icon: "picture-filled"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["systemTools"], Path: "mcpTool", Name: "mcpTool", Component: "view/systemTools/autoCode/mcp.vue", Sort: 7, Meta: Meta{Title: "Mcp Tools模板", Icon: "magnet"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["systemTools"], Path: "mcpTest", Name: "mcpTest", Component: "view/systemTools/autoCode/mcpTest.vue", Sort: 7, Meta: Meta{Title: "Mcp Tools测试", Icon: "partly-cloudy"}},

		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["plugin"], Path: "https://plugin.gin-vue-admin.com/", Name: "https://plugin.gin-vue-admin.com/", Component: "https://plugin.gin-vue-admin.com/", Sort: 0, Meta: Meta{Title: "插件市场", Icon: "shop"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["plugin"], Path: "installPlugin", Name: "installPlugin", Component: "view/systemTools/installPlugin/index.vue", Sort: 1, Meta: Meta{Title: "插件安装", Icon: "box"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["plugin"], Path: "pubPlug", Name: "pubPlug", Component: "view/systemTools/pubPlug/pubPlug.vue", Sort: 3, Meta: Meta{Title: "打包插件", Icon: "files"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["plugin"], Path: "plugin-email", Name: "plugin-email", Component: "plugin/email/view/index.vue", Sort: 4, Meta: Meta{Title: "邮件插件", Icon: "message"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["plugin"], Path: "anInfo", Name: "anInfo", Component: "plugin/announcement/view/info.vue", Sort: 5, Meta: Meta{Title: "公告管理[示例]", Icon: "scaleToOriginal"}},

		// 高级能力子菜单
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["advancedCapabilities"], Path: "wlScenes", Name: "wlScenes", Component: "view/wl_playform/wlScenes/wlScenes.vue", Sort: 1, Meta: Meta{Title: "场景联动", Icon: "connection"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["advancedCapabilities"], Path: "wlEngineRules", Name: "wlEngineRules", Component: "view/wl_playform/wlEngineRules/wlEngineRules.vue", Sort: 2, Meta: Meta{Title: "引擎规则", Icon: "document"}},
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["advancedCapabilities"], Path: "wlResources", Name: "wlResources", Component: "view/wl_playform/wlResources/wlResources.vue", Sort: 3, Meta: Meta{Title: "资源管理", Icon: "link"}},

		// wl_playform子菜单 - 设备接入的子菜单项
		// 产品管理：用于管理物联网产品信息
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["wl_playform"], Path: "wl_playform/wlProducts", Name: "wlProducts", Component: "view/wl_playform/wlProducts/wlProducts.vue", Sort: 1, Meta: Meta{Title: "产品管理", Icon: "box"}},
		// 设备管理：用于管理物联网设备信息
		{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["wl_playform"], Path: "wl_playform/wlEquipment", Name: "wlEquipment", Component: "view/wl_playform/wlEquipment/wlEquipment.vue", Sort: 2, Meta: Meta{Title: "设备管理", Icon: "monitor"}},
	}

	// 创建子菜单到数据库
	// 注意：必须先创建父菜单，再创建子菜单，确保ParentId引用有效
	if err = db.Create(&childMenus).Error; err != nil {
		return ctx, errors.Wrap(err, SysBaseMenu{}.TableName()+"子菜单初始化失败!")
	}

	// 组合所有菜单作为返回结果
	// 包含：父级菜单 + 子菜单，形成完整的菜单树结构
	allEntities := append(allMenus, childMenus...)

	// 为wlResources菜单添加按钮权限
	var wlResourcesMenu SysBaseMenu
	if err = db.Where("name = ?", "wlResources").First(&wlResourcesMenu).Error; err != nil {
		return ctx, errors.Wrap(err, "查找wlResources菜单失败!")
	}

	// 定义wlResources菜单的按钮权限
	wlResourcesBtns := []SysBaseMenuBtn{
		{Name: "add", Desc: "新增", SysBaseMenuID: wlResourcesMenu.ID},
		{Name: "edit", Desc: "编辑", SysBaseMenuID: wlResourcesMenu.ID},
		{Name: "delete", Desc: "删除", SysBaseMenuID: wlResourcesMenu.ID},
		{Name: "info", Desc: "查看", SysBaseMenuID: wlResourcesMenu.ID},
		{Name: "batchDelete", Desc: "批量删除", SysBaseMenuID: wlResourcesMenu.ID},
		{Name: "exportTemplate", Desc: "导出模板", SysBaseMenuID: wlResourcesMenu.ID},
		{Name: "exportExcel", Desc: "导出Excel", SysBaseMenuID: wlResourcesMenu.ID},
		{Name: "importExcel", Desc: "导入Excel", SysBaseMenuID: wlResourcesMenu.ID},
		{Name: "verify", Desc: "验证", SysBaseMenuID: wlResourcesMenu.ID},
	}

	// 创建wlResources菜单按钮权限
	if err = db.Create(&wlResourcesBtns).Error; err != nil {
		return ctx, errors.Wrap(err, "wlResources菜单按钮权限初始化失败!")
	}

	// 为wlEngineRules菜单添加按钮权限
	var wlEngineRulesMenu SysBaseMenu
	if err = db.Where("name = ?", "wlEngineRules").First(&wlEngineRulesMenu).Error; err != nil {
		return ctx, errors.Wrap(err, "查找wlEngineRules菜单失败!")
	}

	// 定义wlEngineRules菜单的按钮权限
	wlEngineRulesBtns := []SysBaseMenuBtn{
		{Name: "add", Desc: "新增", SysBaseMenuID: wlEngineRulesMenu.ID},
		{Name: "edit", Desc: "编辑", SysBaseMenuID: wlEngineRulesMenu.ID},
		{Name: "delete", Desc: "删除", SysBaseMenuID: wlEngineRulesMenu.ID},
		{Name: "info", Desc: "查看", SysBaseMenuID: wlEngineRulesMenu.ID},
		{Name: "batchDelete", Desc: "批量删除", SysBaseMenuID: wlEngineRulesMenu.ID},
		{Name: "exportTemplate", Desc: "导出模板", SysBaseMenuID: wlEngineRulesMenu.ID},
		{Name: "exportExcel", Desc: "导出Excel", SysBaseMenuID: wlEngineRulesMenu.ID},
		{Name: "importExcel", Desc: "导入Excel", SysBaseMenuID: wlEngineRulesMenu.ID},
	}

	// 创建wlEngineRules菜单按钮权限
	if err = db.Create(&wlEngineRulesBtns).Error; err != nil {
		return ctx, errors.Wrap(err, "wlEngineRules菜单按钮权限初始化失败!")
	}

	// 为wlScenes菜单添加按钮权限
	var wlScenesMenu SysBaseMenu
	if err = db.Where("name = ?", "wlScenes").First(&wlScenesMenu).Error; err != nil {
		return ctx, errors.Wrap(err, "查找wlScenes菜单失败!")
	}

	// 定义wlScenes菜单的按钮权限
	wlScenesBtns := []SysBaseMenuBtn{
		{Name: "add", Desc: "新增", SysBaseMenuID: wlScenesMenu.ID},
		{Name: "edit", Desc: "编辑", SysBaseMenuID: wlScenesMenu.ID},
		{Name: "delete", Desc: "删除", SysBaseMenuID: wlScenesMenu.ID},
		{Name: "info", Desc: "查看", SysBaseMenuID: wlScenesMenu.ID},
		{Name: "batchDelete", Desc: "批量删除", SysBaseMenuID: wlScenesMenu.ID},
		{Name: "exportTemplate", Desc: "导出模板", SysBaseMenuID: wlScenesMenu.ID},
		{Name: "exportExcel", Desc: "导出Excel", SysBaseMenuID: wlScenesMenu.ID},
		{Name: "importExcel", Desc: "导入Excel", SysBaseMenuID: wlScenesMenu.ID},
	}

	// 创建wlScenes菜单按钮权限
	if err = db.Create(&wlScenesBtns).Error; err != nil {
		return ctx, errors.Wrap(err, "wlScenes菜单按钮权限初始化失败!")
	}

	next = context.WithValue(ctx, i.InitializerName(), allEntities)
	return next, nil
}

// DataInserted 检查菜单数据是否已插入
// 功能：判断是否需要执行菜单初始化
// 检查策略：通过查询设备接入菜单是否存在来判断整个菜单系统是否需要初始化
func (i *initMenu) DataInserted(ctx context.Context) bool {
	// 总是返回false，强制重新初始化菜单数据
	return false
}
