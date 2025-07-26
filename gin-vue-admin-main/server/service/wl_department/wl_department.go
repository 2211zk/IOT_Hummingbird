package wl_department

import (
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_department"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_department/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_department/response"
	"gorm.io/gorm"
)

type WlDepartmentService struct{}

var WlDepartmentServiceApp = new(WlDepartmentService)

// GetWlDepartmentList 获取部门列表（支持树形和平铺）
func (s *WlDepartmentService) GetWlDepartmentList(req request.WlDepartmentSearch) (list []wl_department.WlDepartment, total int64, err error) {
	db := global.GVA_DB.Model(&wl_department.WlDepartment{})

	// 搜索条件
	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.DepartmentName != "" {
		db = db.Where("department_name LIKE ? OR name LIKE ?", "%"+req.DepartmentName+"%", "%"+req.DepartmentName+"%")
	}
	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}

	if req.TreeMode {
		// 树形模式：优化查询性能
		var allDepts []wl_department.WlDepartment

		// 使用索引优化的查询，只选择必要字段
		err = db.Select("id, parent_id, name, department_name, leader, phone, email, status, sort, created_at, updated_at").
			Order("sort asc, created_at desc").Find(&allDepts).Error
		if err != nil {
			return
		}

		// 批量预加载设备数量（避免N+1查询）
		if len(allDepts) > 0 {
			var deptIDs []uint
			for _, dept := range allDepts {
				deptIDs = append(deptIDs, dept.ID)
			}

			// 批量查询设备关联数量
			var deviceCounts []struct {
				DepartmentID uint  `json:"department_id"`
				DeviceCount  int64 `json:"device_count"`
			}

			err = global.GVA_DB.Model(&wl_department.WlDepartmentDevice{}).
				Select("department_id, COUNT(*) as device_count").
				Where("department_id IN ?", deptIDs).
				Group("department_id").
				Find(&deviceCounts).Error

			if err == nil {
				// 创建设备数量映射
				deviceCountMap := make(map[uint]int64)
				for _, count := range deviceCounts {
					deviceCountMap[count.DepartmentID] = count.DeviceCount
				}

				// 设置设备数量到部门对象
				for i := range allDepts {
					allDepts[i].DeviceCount = deviceCountMap[allDepts[i].ID]
				}
			}
		}

		list = s.buildDepartmentTreeOptimized(allDepts, nil)
		total = int64(len(allDepts))
	} else {
		// 平铺模式：分页查询
		err = db.Count(&total).Error
		if err != nil {
			return
		}

		// 只预加载必要的设备信息
		err = db.Select("id, parent_id, name, department_name, leader, phone, email, status, sort, created_at, updated_at").
			Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize).
			Order("sort asc, created_at desc").Find(&list).Error

		if err == nil && len(list) > 0 {
			// 批量查询设备数量
			var deptIDs []uint
			for _, dept := range list {
				deptIDs = append(deptIDs, dept.ID)
			}

			var deviceCounts []struct {
				DepartmentID uint  `json:"department_id"`
				DeviceCount  int64 `json:"device_count"`
			}

			countErr := global.GVA_DB.Model(&wl_department.WlDepartmentDevice{}).
				Select("department_id, COUNT(*) as device_count").
				Where("department_id IN ?", deptIDs).
				Group("department_id").
				Find(&deviceCounts).Error

			if countErr == nil {
				deviceCountMap := make(map[uint]int64)
				for _, count := range deviceCounts {
					deviceCountMap[count.DepartmentID] = count.DeviceCount
				}

				for i := range list {
					list[i].DeviceCount = deviceCountMap[list[i].ID]
				}
			}
		}
	}

	return
}

// GetDepartmentTree 获取部门树（用于选择上级部门）
func (s *WlDepartmentService) GetDepartmentTree(req request.DepartmentTreeReq) (tree []*response.DepartmentTreeNode, err error) {
	// 尝试从缓存获取
	tree, err = s.GetDepartmentTreeFromCache(req.ExcludeID)
	if err != nil {
		return nil, err
	}

	return tree, nil
}

// CreateWlDepartment 创建部门
func (s *WlDepartmentService) CreateWlDepartment(req request.CreateWlDepartmentReq) error {
	// 使用事务确保数据一致性
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 验证上级部门是否存在
		if req.ParentID != nil {
			var parent wl_department.WlDepartment
			if err := tx.First(&parent, *req.ParentID).Error; err != nil {
				return errors.New("上级部门不存在")
			}
		}

		// 检查同级部门名称是否重复
		if err := s.checkDepartmentNameUnique(tx, req.ParentID, getNameFromReq(req), 0); err != nil {
			return err
		}

		// 创建部门
		dept := wl_department.WlDepartment{
			ParentID:       req.ParentID,
			Name:           getNameFromReq(req),
			DepartmentName: getNameFromReq(req), // 兼容字段
			Leader:         req.Leader,
			Phone:          req.Phone,
			Email:          req.Email,
			Status:         getStatusOrDefault(req.Status),
			Sort:           req.Sort,
		}

		if err := tx.Create(&dept).Error; err != nil {
			return err
		}

		// 关联设备
		if len(req.DeviceIDs) > 0 {
			if err := s.associateDevices(tx, dept.ID, req.DeviceIDs); err != nil {
				return err
			}
		}

		// 清除缓存
		s.InvalidateCache()

		return nil
	})
}

// UpdateWlDepartment 更新部门
func (s *WlDepartmentService) UpdateWlDepartment(req request.UpdateWlDepartmentReq) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var dept wl_department.WlDepartment
		if err := tx.First(&dept, req.ID).Error; err != nil {
			return errors.New("部门不存在")
		}

		// 验证上级部门关系
		if req.ParentID != nil {
			if *req.ParentID == req.ID {
				return errors.New("不能选择自身作为上级部门")
			}

			// 检查是否会形成循环引用
			if err := s.checkCircularReference(tx, req.ID, *req.ParentID); err != nil {
				return err
			}

			// 验证上级部门是否存在
			var parent wl_department.WlDepartment
			if err := tx.First(&parent, *req.ParentID).Error; err != nil {
				return errors.New("上级部门不存在")
			}
		}

		// 检查同级部门名称是否重复
		if err := s.checkDepartmentNameUnique(tx, req.ParentID, getNameFromUpdateReq(req), req.ID); err != nil {
			return err
		}

		// 更新部门信息
		dept.ParentID = req.ParentID
		dept.Name = getNameFromUpdateReq(req)
		dept.DepartmentName = getNameFromUpdateReq(req) // 兼容字段
		dept.Leader = req.Leader
		dept.Phone = req.Phone
		dept.Email = req.Email
		dept.Status = getStatusOrDefault(req.Status)
		dept.Sort = req.Sort

		if err := tx.Save(&dept).Error; err != nil {
			return err
		}

		// 更新设备关联
		if err := s.updateDeviceAssociations(tx, dept.ID, req.DeviceIDs); err != nil {
			return err
		}

		// 如果禁用部门，同时禁用所有子部门
		if req.Status == "禁用" {
			if err := s.disableChildrenDepartments(tx, req.ID); err != nil {
				return err
			}
		}

		// 清除缓存
		s.InvalidateCache()

		return nil
	})
}

// DeleteWlDepartment 删除部门
func (s *WlDepartmentService) DeleteWlDepartment(req request.DeleteWlDepartmentReq) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 检查是否有子部门
		var childCount int64
		if err := tx.Model(&wl_department.WlDepartment{}).Where("parent_id = ?", req.ID).Count(&childCount).Error; err != nil {
			return err
		}
		if childCount > 0 {
			return errors.New("该部门下还有子部门，无法删除")
		}

		// 删除设备关联
		if err := tx.Where("department_id = ?", req.ID).Delete(&wl_department.WlDepartmentDevice{}).Error; err != nil {
			return err
		}

		// 删除部门
		err := tx.Delete(&wl_department.WlDepartment{}, req.ID).Error
		if err != nil {
			return err
		}

		// 清除缓存
		s.InvalidateCache()

		return nil
	})
}

// GetDepartmentDetail 获取部门详情
func (s *WlDepartmentService) GetDepartmentDetail(id uint) (dept wl_department.WlDepartment, err error) {
	err = global.GVA_DB.Preload("Devices").Preload("Parent").First(&dept, id).Error
	if err != nil {
		return
	}

	// 获取子部门
	var children []wl_department.WlDepartment
	err = global.GVA_DB.Where("parent_id = ?", id).Find(&children).Error
	if err != nil {
		return
	}
	dept.Children = children

	return
}

// GetAvailableDevices 获取可关联的设备列表
func (s *WlDepartmentService) GetAvailableDevices(req request.AvailableDevicesReq) (list []wl_department.WlDevice, total int64, err error) {
	db := global.GVA_DB.Model(&wl_department.WlDevice{})

	// 搜索条件
	if req.DeviceName != "" {
		db = db.Where("device_name LIKE ?", "%"+req.DeviceName+"%")
	}
	if req.ProductName != "" {
		db = db.Where("product_name LIKE ?", "%"+req.ProductName+"%")
	}

	// 排除已关联指定部门的设备
	if req.DepartmentID != nil {
		subQuery := global.GVA_DB.Model(&wl_department.WlDepartmentDevice{}).
			Select("device_id").
			Where("department_id = ?", *req.DepartmentID)
		db = db.Where("id NOT IN (?)", subQuery)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize).
		Order("created_at desc").Find(&list).Error

	return
}

// GetDepartmentDevices 获取部门已关联的设备
func (s *WlDepartmentService) GetDepartmentDevices(req request.DepartmentDevicesReq) (list []wl_department.WlDevice, total int64, err error) {
	db := global.GVA_DB.Model(&wl_department.WlDevice{}).
		Joins("JOIN wl_department_device ON wl_device.id = wl_department_device.device_id").
		Where("wl_department_device.department_id = ?", req.DepartmentID)

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize).
		Order("wl_department_device.created_at desc").Find(&list).Error

	return
}

// buildDepartmentTree 构建部门树形结构
func (s *WlDepartmentService) buildDepartmentTree(departments []wl_department.WlDepartment, parentID *uint) []wl_department.WlDepartment {
	var result []wl_department.WlDepartment

	for _, dept := range departments {
		if (parentID == nil && dept.ParentID == nil) || (parentID != nil && dept.ParentID != nil && *dept.ParentID == *parentID) {
			dept.Children = s.buildDepartmentTree(departments, &dept.ID)
			result = append(result, dept)
		}
	}

	return result
}

// buildDepartmentTreeOptimized 优化的部门树构建（使用映射提高性能）
func (s *WlDepartmentService) buildDepartmentTreeOptimized(departments []wl_department.WlDepartment, parentID *uint) []wl_department.WlDepartment {
	// 创建父子关系映射，避免重复遍历
	childrenMap := make(map[uint][]wl_department.WlDepartment)
	var rootDepts []wl_department.WlDepartment

	// 第一次遍历：建立映射关系
	for _, dept := range departments {
		if dept.ParentID == nil {
			rootDepts = append(rootDepts, dept)
		} else {
			childrenMap[*dept.ParentID] = append(childrenMap[*dept.ParentID], dept)
		}
	}

	// 递归构建树形结构
	var buildTree func(depts []wl_department.WlDepartment) []wl_department.WlDepartment
	buildTree = func(depts []wl_department.WlDepartment) []wl_department.WlDepartment {
		var result []wl_department.WlDepartment
		for _, dept := range depts {
			if children, exists := childrenMap[dept.ID]; exists {
				dept.Children = buildTree(children)
			}
			result = append(result, dept)
		}
		return result
	}

	if parentID == nil {
		return buildTree(rootDepts)
	}

	// 如果指定了parentID，返回该父节点的子树
	if children, exists := childrenMap[*parentID]; exists {
		return buildTree(children)
	}

	return []wl_department.WlDepartment{}
}

// checkCircularReference 检查循环引用
func (s *WlDepartmentService) checkCircularReference(tx *gorm.DB, deptID uint, newParentID uint) error {
	// 获取新上级部门的所有上级部门ID
	parentIDs, err := s.getAllParentIDs(tx, newParentID)
	if err != nil {
		return err
	}

	// 检查当前部门是否在新上级部门的上级链中
	for _, parentID := range parentIDs {
		if parentID == deptID {
			return errors.New("不能选择子部门作为上级部门")
		}
	}

	return nil
}

// getAllParentIDs 获取部门的所有上级部门ID
func (s *WlDepartmentService) getAllParentIDs(tx *gorm.DB, deptID uint) ([]uint, error) {
	var parentIDs []uint
	currentID := deptID

	for {
		var dept wl_department.WlDepartment
		if err := tx.Select("parent_id").First(&dept, currentID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				break
			}
			return nil, err
		}

		if dept.ParentID == nil {
			break
		}

		parentIDs = append(parentIDs, *dept.ParentID)
		currentID = *dept.ParentID
	}

	return parentIDs, nil
}

// getAllChildrenIDs 获取部门的所有子部门ID
func (s *WlDepartmentService) getAllChildrenIDs(deptID uint) ([]uint, error) {
	var childrenIDs []uint

	var directChildren []wl_department.WlDepartment
	if err := global.GVA_DB.Select("id").Where("parent_id = ?", deptID).Find(&directChildren).Error; err != nil {
		return nil, err
	}

	for _, child := range directChildren {
		childrenIDs = append(childrenIDs, child.ID)
		subChildren, err := s.getAllChildrenIDs(child.ID)
		if err != nil {
			return nil, err
		}
		childrenIDs = append(childrenIDs, subChildren...)
	}

	return childrenIDs, nil
}

// checkDepartmentNameUnique 检查同级部门名称唯一性
func (s *WlDepartmentService) checkDepartmentNameUnique(tx *gorm.DB, parentID *uint, name string, excludeID uint) error {
	db := tx.Model(&wl_department.WlDepartment{}).Where("name = ?", name)

	if parentID == nil {
		db = db.Where("parent_id IS NULL")
	} else {
		db = db.Where("parent_id = ?", *parentID)
	}

	if excludeID > 0 {
		db = db.Where("id != ?", excludeID)
	}

	var count int64
	if err := db.Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return errors.New("同级部门名称已存在")
	}

	return nil
}

// associateDevices 关联设备
func (s *WlDepartmentService) associateDevices(tx *gorm.DB, deptID uint, deviceIDs []uint) error {
	for _, deviceID := range deviceIDs {
		// 检查设备是否存在
		var device wl_department.WlDevice
		if err := tx.First(&device, deviceID).Error; err != nil {
			return fmt.Errorf("设备ID %d 不存在", deviceID)
		}

		// 创建关联
		association := wl_department.WlDepartmentDevice{
			DepartmentID: deptID,
			DeviceID:     deviceID,
		}
		if err := tx.Create(&association).Error; err != nil {
			// 忽略重复关联错误
			if !isDuplicateError(err) {
				return err
			}
		}
	}
	return nil
}

// updateDeviceAssociations 更新设备关联
func (s *WlDepartmentService) updateDeviceAssociations(tx *gorm.DB, deptID uint, deviceIDs []uint) error {
	// 删除现有关联
	if err := tx.Where("department_id = ?", deptID).Delete(&wl_department.WlDepartmentDevice{}).Error; err != nil {
		return err
	}

	// 创建新关联
	if len(deviceIDs) > 0 {
		return s.associateDevices(tx, deptID, deviceIDs)
	}

	return nil
}

// disableChildrenDepartments 禁用所有子部门
func (s *WlDepartmentService) disableChildrenDepartments(tx *gorm.DB, parentID uint) error {
	return tx.Model(&wl_department.WlDepartment{}).
		Where("parent_id = ?", parentID).
		Update("status", "禁用").Error
}

// 辅助函数
func getNameFromReq(req request.CreateWlDepartmentReq) string {
	if req.Name != "" {
		return req.Name
	}
	return req.DepartmentName
}

func getNameFromUpdateReq(req request.UpdateWlDepartmentReq) string {
	if req.Name != "" {
		return req.Name
	}
	return req.DepartmentName
}

func getStatusOrDefault(status string) string {
	if status == "" {
		return "启用"
	}
	return status
}

func isDuplicateError(err error) bool {
	// 检查是否为重复键错误，具体实现根据数据库类型而定
	return false // 简化处理，实际应该检查具体的错误类型
}
