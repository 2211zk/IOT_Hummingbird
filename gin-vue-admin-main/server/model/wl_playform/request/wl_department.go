package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type WlDepartmentSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	request.PageInfo
}

// 部门创建请求
type CreateDepartmentRequest struct {
	ParentID       *int   `json:"parentId" form:"parentId"`
	DepartmentName string `json:"departmentName" form:"departmentName" binding:"required"`
	Leader         string `json:"leader" form:"leader"`
	Phone          string `json:"phone" form:"phone"`
	Email          string `json:"email" form:"email"`
	Status         string `json:"status" form:"status"`
	Sort           int    `json:"sort" form:"sort"`
	DeviceIDs      []int  `json:"deviceIds" form:"deviceIds[]"` // 设备ID列表
}

// 部门编辑请求
type UpdateDepartmentRequest struct {
	ID             int    `json:"id" form:"id" binding:"required"`
	ParentID       *int   `json:"parentId" form:"parentId"`
	DepartmentName string `json:"departmentName" form:"departmentName" binding:"required"`
	Leader         string `json:"leader" form:"leader"`
	Phone          string `json:"phone" form:"phone"`
	Email          string `json:"email" form:"email"`
	Status         string `json:"status" form:"status"`
	Sort           int    `json:"sort" form:"sort"`
	DeviceIDs      []int  `json:"deviceIds" form:"deviceIds[]"` // 设备ID列表
}

// 部门分配设备请求
type AssignDevicesRequest struct {
	DepartmentID int   `json:"departmentId" form:"departmentId" binding:"required"`
	DeviceIDs    []int `json:"deviceIds" form:"deviceIds[]" binding:"required"`
}

// 查询部门下设备请求
type DepartmentDevicesRequest struct {
	DepartmentID int `json:"departmentId" form:"departmentId" binding:"required"`
}
