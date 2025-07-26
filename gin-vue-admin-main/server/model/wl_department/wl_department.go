package wl_department

import (
	"time"

	"gorm.io/gorm"
)

type WlDepartment struct {
	ID             uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	ParentID       *uint          `json:"parentId" gorm:"column:parent_id;index"`
	Name           string         `json:"name" gorm:"column:name;not null;size:100"`
	Leader         string         `json:"leader" gorm:"column:leader;size:32"`
	Phone          string         `json:"phone" gorm:"column:phone;size:20"`
	Email          string         `json:"email" gorm:"column:email;size:64"`
	Status         string         `json:"status" gorm:"column:status;size:8;default:启用"`
	Sort           int            `json:"sort" gorm:"column:sort;default:0"`
	CreatedAt      time.Time      `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt      time.Time      `json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at;index"`
	DepartmentName string         `json:"departmentName" gorm:"column:department_name;size:64"` // 兼容字段
	CreatedBy      uint           `json:"createdBy" gorm:"column:created_by"`
	UpdatedBy      uint           `json:"updatedBy" gorm:"column:updated_by"`
	DeletedBy      uint           `json:"deletedBy" gorm:"column:deleted_by"`

	// 关联字段
	Children    []WlDepartment `json:"children" gorm:"-"`
	Parent      *WlDepartment  `json:"parent" gorm:"foreignKey:ParentID"`
	Devices     []WlDevice     `json:"devices" gorm:"many2many:wl_department_device;"`
	DeviceCount int64          `json:"deviceCount" gorm:"-"` // 设备数量（不存储到数据库）
}

// TableName 指定表名
func (WlDepartment) TableName() string {
	return "wl_department"
}

// WlDevice 设备模型
type WlDevice struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	DeviceName  string    `json:"deviceName" gorm:"column:device_name"`
	ProductName string    `json:"productName" gorm:"column:product_name"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

// TableName 指定表名
func (WlDevice) TableName() string {
	return "wl_device"
}

// WlDepartmentDevice 部门设备关联模型
type WlDepartmentDevice struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	DepartmentID uint      `json:"departmentId" gorm:"column:department_id;not null"`
	DeviceID     uint      `json:"deviceId" gorm:"column:device_id;not null"`
	CreatedAt    time.Time `json:"createdAt" gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
}

// TableName 指定表名
func (WlDepartmentDevice) TableName() string {
	return "wl_department_device"
}
