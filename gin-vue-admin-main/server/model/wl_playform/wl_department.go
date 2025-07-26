// 自动生成模板WlDepartment
package wl_playform

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// wlDepartment表 结构体  WlDepartment
type WlDepartment struct {
	global.GVA_MODEL
	ParentID       *int    `json:"parentId" form:"parentId" gorm:"comment:上级部门ID;column:parent_id;"`                         // 上级部门ID
	DepartmentName *string `json:"departmentName" form:"departmentName" gorm:"comment:部门名称;column:department_name;size:64;"` // 部门名称
	Leader         *string `json:"leader" form:"leader" gorm:"comment:负责人;column:leader;size:32;"`                           // 负责人
	Phone          *string `json:"phone" form:"phone" gorm:"comment:电话;column:phone;size:20;"`                               // 电话
	Email          *string `json:"email" form:"email" gorm:"comment:邮箱;column:email;size:64;"`                               // 邮箱
	Status         *string `json:"status" form:"status" gorm:"comment:状态;column:status;size:8;"`                             // 状态
	Sort           *int    `json:"sort" form:"sort" gorm:"comment:排序;column:sort;"`                                          // 排序
	CreatedBy      *int    `json:"createdBy" form:"createdBy" gorm:"comment:创建者;column:created_by;size:20;"`                 // 创建者
	UpdatedBy      *int    `json:"updatedBy" form:"updatedBy" gorm:"comment:更新者;column:updated_by;size:20;"`                 // 更新者
	DeletedBy      *int    `json:"deletedBy" form:"deletedBy" gorm:"comment:删除者;column:deleted_by;size:20;"`                 // 删除者
}

// TableName wlDepartment表 WlDepartment自定义表名 wl_department
func (WlDepartment) TableName() string {
	return "wl_department"
}

// 设备表结构体
// 对应 device 表
// 可根据实际需求添加更多字段
type Device struct {
	ID          int        `json:"id" gorm:"primaryKey"`
	Name        string     `json:"name" gorm:"column:name;size:100;not null;comment:设备名称"`
	ProductName string     `json:"productName" gorm:"column:product_name;size:100;comment:产品名称"`
	CreatedAt   *time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (Device) TableName() string {
	return "device"
}

// 部门-设备关联表结构体
// 对应 department_device 表
type DepartmentDevice struct {
	ID           int        `json:"id" gorm:"primaryKey"`
	DepartmentID int        `json:"departmentId" gorm:"column:department_id;not null;comment:部门ID"`
	DeviceID     int        `json:"deviceId" gorm:"column:device_id;not null;comment:设备ID"`
	CreatedAt    *time.Time `json:"createdAt" gorm:"column:created_at"`
}

func (DepartmentDevice) TableName() string {
	return "department_device"
}
