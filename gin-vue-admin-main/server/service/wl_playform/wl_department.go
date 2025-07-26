package wl_playform

import (
	"context"
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform"
	wl_playformReq "github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform/request"
	"gorm.io/gorm"
)

type WlDepartmentService struct{}

// CreateWlDepartment 创建wlDepartment表记录
// Author [yourname](https://github.com/yourname)
func (wlDepartmentService *WlDepartmentService) CreateWlDepartment(ctx context.Context, wlDepartment *wl_playform.WlDepartment) (err error) {
	err = global.GVA_DB.Create(wlDepartment).Error
	return err
}

// DeleteWlDepartment 删除wlDepartment表记录
// Author [yourname](https://github.com/yourname)
func (wlDepartmentService *WlDepartmentService) DeleteWlDepartment(ctx context.Context, ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&wl_playform.WlDepartment{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&wl_playform.WlDepartment{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteWlDepartmentByIds 批量删除wlDepartment表记录
// Author [yourname](https://github.com/yourname)
func (wlDepartmentService *WlDepartmentService) DeleteWlDepartmentByIds(ctx context.Context, IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&wl_playform.WlDepartment{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&wl_playform.WlDepartment{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateWlDepartment 更新wlDepartment表记录
// Author [yourname](https://github.com/yourname)
func (wlDepartmentService *WlDepartmentService) UpdateWlDepartment(ctx context.Context, wlDepartment wl_playform.WlDepartment) (err error) {
	err = global.GVA_DB.Model(&wl_playform.WlDepartment{}).Where("id = ?", wlDepartment.ID).Updates(&wlDepartment).Error
	return err
}

// GetWlDepartment 根据ID获取wlDepartment表记录
// Author [yourname](https://github.com/yourname)
func (wlDepartmentService *WlDepartmentService) GetWlDepartment(ctx context.Context, ID string) (wlDepartment wl_playform.WlDepartment, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wlDepartment).Error
	return
}

// GetWlDepartmentInfoList 分页获取wlDepartment表记录
// Author [yourname](https://github.com/yourname)
func (wlDepartmentService *WlDepartmentService) GetWlDepartmentInfoList(ctx context.Context, info wl_playformReq.WlDepartmentSearch) (list []wl_playform.WlDepartment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&wl_playform.WlDepartment{})
	var wlDepartments []wl_playform.WlDepartment
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&wlDepartments).Error
	return wlDepartments, total, err
}
func (wlDepartmentService *WlDepartmentService) GetWlDepartmentPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// 分配设备到部门（覆盖式分配，先清空再插入）
func (wlDepartmentService *WlDepartmentService) AssignDevicesToDepartment(ctx context.Context, departmentId int, deviceIds []int) error {
	tx := global.GVA_DB.Begin()
	// 先删除原有分配
	if err := tx.Where("department_id = ?", departmentId).Delete(&wl_playform.DepartmentDevice{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 批量插入新分配
	for _, deviceId := range deviceIds {
		dd := wl_playform.DepartmentDevice{
			DepartmentID: departmentId,
			DeviceID:     deviceId,
		}
		if err := tx.Create(&dd).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

// 查询部门下所有设备
func (wlDepartmentService *WlDepartmentService) GetDevicesByDepartment(ctx context.Context, departmentId int) ([]wl_playform.Device, error) {
	var devices []wl_playform.Device
	err := global.GVA_DB.Table("device").
		Select("device.*").
		Joins("JOIN department_device ON device.id = department_device.device_id").
		Where("department_device.department_id = ?", departmentId).
		Find(&devices).Error
	return devices, err
}

// 移除部门下某个设备
func (wlDepartmentService *WlDepartmentService) RemoveDeviceFromDepartment(ctx context.Context, departmentId int, deviceId int) error {
	result := global.GVA_DB.Where("department_id = ? AND device_id = ?", departmentId, deviceId).Delete(&wl_playform.DepartmentDevice{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("未找到对应的部门设备关系")
	}
	return nil
}
