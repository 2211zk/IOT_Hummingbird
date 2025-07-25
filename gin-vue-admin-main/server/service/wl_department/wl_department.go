package wl_department

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_department"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_department/request"
)

type WlDepartmentService struct{}

var WlDepartmentServiceApp = new(WlDepartmentService)

func (s *WlDepartmentService) GetWlDepartmentList(req request.WlDepartmentSearch) (list []wl_department.WlDepartment, total int64, err error) {
	db := global.GVA_DB.Model(&wl_department.WlDepartment{})
	if req.DepartmentName != "" {
		db = db.Where("department_name LIKE ?", "%"+req.DepartmentName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize).Order("sort asc, created_at desc").Find(&list).Error
	return
}

func (s *WlDepartmentService) CreateWlDepartment(req request.CreateWlDepartmentReq) error {
	dep := wl_department.WlDepartment{
		DepartmentName: req.DepartmentName,
		Leader:         req.Leader,
		Phone:          req.Phone,
		Email:          req.Email,
		Status:         req.Status,
		Sort:           req.Sort,
	}
	return global.GVA_DB.Create(&dep).Error
}

func (s *WlDepartmentService) UpdateWlDepartment(req request.UpdateWlDepartmentReq) error {
	var dep wl_department.WlDepartment
	if err := global.GVA_DB.First(&dep, req.ID).Error; err != nil {
		return err
	}
	dep.DepartmentName = req.DepartmentName
	dep.Leader = req.Leader
	dep.Phone = req.Phone
	dep.Email = req.Email
	dep.Status = req.Status
	dep.Sort = req.Sort
	return global.GVA_DB.Save(&dep).Error
}

func (s *WlDepartmentService) DeleteWlDepartment(req request.DeleteWlDepartmentReq) error {
	return global.GVA_DB.Delete(&wl_department.WlDepartment{}, req.ID).Error
}
