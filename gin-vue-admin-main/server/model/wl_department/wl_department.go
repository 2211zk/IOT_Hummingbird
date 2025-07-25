package wl_department

import "time"

type WlDepartment struct {
	ID             int       `json:"id" gorm:"primaryKey"`
	DepartmentName string    `json:"departmentName" gorm:"column:department_name"`
	Leader         string    `json:"leader"`
	Phone          string    `json:"phone"`
	Email          string    `json:"email"`
	Status         string    `json:"status"`
	Sort           int       `json:"sort"`
	CreatedAt      time.Time `json:"createdAt" gorm:"column:created_at"`
}

func (WlDepartment) TableName() string {
	return "wl_department"
}
