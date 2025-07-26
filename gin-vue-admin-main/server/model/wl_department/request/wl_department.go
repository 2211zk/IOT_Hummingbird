package request

// WlDepartmentSearch 部门搜索请求
type WlDepartmentSearch struct {
	Page           int    `json:"page" form:"page"`
	PageSize       int    `json:"pageSize" form:"pageSize"`
	Name           string `json:"name" form:"name"`
	DepartmentName string `json:"departmentName" form:"departmentName"` // 兼容字段
	Status         string `json:"status" form:"status"`
	TreeMode       bool   `json:"treeMode" form:"treeMode"` // 是否返回树形结构
}

// CreateWlDepartmentReq 创建部门请求
type CreateWlDepartmentReq struct {
	ParentID       *uint  `json:"parentId"`
	Name           string `json:"name" binding:"required"`
	DepartmentName string `json:"departmentName"` // 兼容字段，如果提供则使用此值作为name
	Leader         string `json:"leader"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	Status         string `json:"status"`
	Sort           int    `json:"sort"`
	DeviceIDs      []uint `json:"deviceIds"` // 关联的设备ID列表
}

// UpdateWlDepartmentReq 更新部门请求
type UpdateWlDepartmentReq struct {
	ID             uint   `json:"id" binding:"required"`
	ParentID       *uint  `json:"parentId"`
	Name           string `json:"name" binding:"required"`
	DepartmentName string `json:"departmentName"` // 兼容字段
	Leader         string `json:"leader"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	Status         string `json:"status"`
	Sort           int    `json:"sort"`
	DeviceIDs      []uint `json:"deviceIds"` // 关联的设备ID列表
}

// DeleteWlDepartmentReq 删除部门请求
type DeleteWlDepartmentReq struct {
	ID uint `json:"id" binding:"required"`
}

// DepartmentTreeReq 获取部门树请求
type DepartmentTreeReq struct {
	ExcludeID *uint `json:"excludeId" form:"excludeId"` // 排除的部门ID（用于编辑时排除自身和子部门）
}

// AvailableDevicesReq 获取可用设备请求
type AvailableDevicesReq struct {
	Page         int    `json:"page" form:"page"`
	PageSize     int    `json:"pageSize" form:"pageSize"`
	DeviceName   string `json:"deviceName" form:"deviceName"`
	ProductName  string `json:"productName" form:"productName"`
	DepartmentID *uint  `json:"departmentId" form:"departmentId"` // 排除已关联此部门的设备
}

// DepartmentDevicesReq 获取部门设备请求
type DepartmentDevicesReq struct {
	DepartmentID uint `json:"departmentId" form:"departmentId" binding:"required"`
	Page         int  `json:"page" form:"page"`
	PageSize     int  `json:"pageSize" form:"pageSize"`
}
