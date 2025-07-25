package request

type WlDepartmentSearch struct {
	Page           int    `json:"page"`
	PageSize       int    `json:"pageSize"`
	DepartmentName string `json:"departmentName"`
}

type CreateWlDepartmentReq struct {
	DepartmentName string `json:"departmentName" binding:"required"`
	Leader         string `json:"leader"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	Status         string `json:"status"`
	Sort           int    `json:"sort"`
}

type UpdateWlDepartmentReq struct {
	ID             int    `json:"id" binding:"required"`
	DepartmentName string `json:"departmentName" binding:"required"`
	Leader         string `json:"leader"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	Status         string `json:"status"`
	Sort           int    `json:"sort"`
}

type DeleteWlDepartmentReq struct {
	ID int `json:"id" binding:"required"`
}
