# 菜单重构说明

## 更改概述

根据用户需求，将"场景联动"、"引擎规则"、"资源管理"这三个页面从主页移动到"高级能力"菜单下，作为其子菜单项。

## 具体更改

### 1. 添加新的父级菜单
在 `server/source/system/menu.go` 中添加了新的父级菜单：
```go
{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "advancedCapabilities", Name: "advancedCapabilities", Component: "view/routerHolder.vue", Sort: 10, Meta: Meta{Title: "高级能力", Icon: "cloud"}}
```

### 2. 添加子菜单
在同一个文件中添加了三个子菜单：
```go
// 高级能力子菜单
{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["advancedCapabilities"], Path: "wlScenes", Name: "wlScenes", Component: "view/wl_playform/wlScenes/wlScenes.vue", Sort: 1, Meta: Meta{Title: "场景联动", Icon: "connection"}},
{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["advancedCapabilities"], Path: "wlEngineRules", Name: "wlEngineRules", Component: "view/wl_playform/wlEngineRules/wlEngineRules.vue", Sort: 2, Meta: Meta{Title: "引擎规则", Icon: "document"}},
{MenuLevel: 1, Hidden: false, ParentId: menuNameMap["advancedCapabilities"], Path: "wlResources", Name: "wlResources", Component: "view/wl_playform/wlResources/wlResources.vue", Sort: 3, Meta: Meta{Title: "资源管理", Icon: "link"}},
```

## 菜单结构

### 修改前
- 场景联动 (独立菜单)
- 引擎规则 (独立菜单)  
- 资源管理 (独立菜单)

### 修改后
- 高级能力
  - 场景联动
  - 引擎规则
  - 资源管理

## 应用更改

要应用这些更改，需要：

1. 重新启动后端服务
2. 清除浏览器缓存
3. 重新登录系统

## 图标说明

- 高级能力: `cloud` (云朵图标)
- 场景联动: `connection` (连接图标)
- 引擎规则: `document` (文档图标)
- 资源管理: `link` (链接图标)

## 注意事项

1. 这些更改只影响菜单结构，不影响页面功能
2. 所有原有的API路由和页面组件保持不变
3. 用户权限需要重新分配以访问新的菜单结构 