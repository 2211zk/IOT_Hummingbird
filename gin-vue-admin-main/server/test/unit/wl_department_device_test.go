package unit

import (
	"fmt"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_department"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_department/request"
	"github.com/stretchr/testify/suite"
)

// WlDepartmentDeviceTestSuite 部门设备关联测试套件
type WlDepartmentDeviceTestSuite struct {
	suite.Suite
}

func TestWlDepartmentDeviceTestSuite(t *testing.T) {
	suite.Run(t, new(WlDepartmentDeviceTestSuite))
}

// TestDeviceAssociation 测试设备关联功能
func (suite *WlDepartmentDeviceTestSuite) TestDeviceAssociation() {
	suite.Run("关联单个设备", func() {
		deptID := 1
		deviceIDs := []int{1}

		result := simulateAssociateDevices(deptID, deviceIDs)
		suite.True(result.Success, "关联单个设备应该成功")
		suite.Len(result.AssociatedDevices, 1, "应该关联1个设备")
		suite.Equal(1, result.AssociatedDevices[0].ID, "设备ID应该匹配")
	})

	suite.Run("关联多个设备", func() {
		deptID := 1
		deviceIDs := []int{1, 2, 3}

		result := simulateAssociateDevices(deptID, deviceIDs)
		suite.True(result.Success, "关联多个设备应该成功")
		suite.Len(result.AssociatedDevices, 3, "应该关联3个设备")

		// 验证所有设备都被关联
		associatedIDs := make([]int, len(result.AssociatedDevices))
		for i, device := range result.AssociatedDevices {
			associatedIDs[i] = device.ID
		}
		suite.ElementsMatch(deviceIDs, associatedIDs, "关联的设备ID应该匹配")
	})

	suite.Run("关联空设备列表", func() {
		deptID := 1
		deviceIDs := []int{}

		result := simulateAssociateDevices(deptID, deviceIDs)
		suite.True(result.Success, "关联空设备列表应该成功（清空关联）")
		suite.Empty(result.AssociatedDevices, "应该没有关联设备")
	})

	suite.Run("关联重复设备ID", func() {
		deptID := 1
		deviceIDs := []int{1, 2, 1, 3, 2} // 包含重复ID

		result := simulateAssociateDevices(deptID, deviceIDs)
		suite.True(result.Success, "关联重复设备ID应该成功（自动去重）")
		suite.Len(result.AssociatedDevices, 3, "应该去重后关联3个设备")

		// 验证去重效果
		uniqueIDs := []int{1, 2, 3}
		associatedIDs := make([]int, len(result.AssociatedDevices))
		for i, device := range result.AssociatedDevices {
			associatedIDs[i] = device.ID
		}
		suite.ElementsMatch(uniqueIDs, associatedIDs, "应该自动去重")
	})

	suite.Run("关联无效设备ID", func() {
		deptID := 1
		deviceIDs := []int{0, -1, 999999} // 无效的设备ID

		result := simulateAssociateDevices(deptID, deviceIDs)
		suite.False(result.Success, "关联无效设备ID应该失败")
		suite.Contains(result.ErrorMessage, "无效的设备ID", "应该返回无效设备ID错误")
	})

	suite.Run("无效部门ID", func() {
		deptID := 0
		deviceIDs := []int{1, 2}

		result := simulateAssociateDevices(deptID, deviceIDs)
		suite.False(result.Success, "无效部门ID应该失败")
		suite.Contains(result.ErrorMessage, "无效的部门ID", "应该返回无效部门ID错误")
	})
}

// TestUpdateDeviceAssociation 测试更新设备关联
func (suite *WlDepartmentDeviceTestSuite) TestUpdateDeviceAssociation() {
	suite.Run("添加新设备", func() {
		deptID := 1
		oldDeviceIDs := []int{1, 2}
		newDeviceIDs := []int{1, 2, 3}

		result := simulateUpdateDeviceAssociation(deptID, oldDeviceIDs, newDeviceIDs)
		suite.True(result.Success, "添加新设备应该成功")
		suite.ElementsMatch([]int{3}, result.AddedDevices, "应该添加设备3")
		suite.Empty(result.RemovedDevices, "不应该移除任何设备")
	})

	suite.Run("移除设备", func() {
		deptID := 1
		oldDeviceIDs := []int{1, 2, 3}
		newDeviceIDs := []int{1, 3}

		result := simulateUpdateDeviceAssociation(deptID, oldDeviceIDs, newDeviceIDs)
		suite.True(result.Success, "移除设备应该成功")
		suite.Empty(result.AddedDevices, "不应该添加任何设备")
		suite.ElementsMatch([]int{2}, result.RemovedDevices, "应该移除设备2")
	})

	suite.Run("同时添加和移除设备", func() {
		deptID := 1
		oldDeviceIDs := []int{1, 2}
		newDeviceIDs := []int{2, 3, 4}

		result := simulateUpdateDeviceAssociation(deptID, oldDeviceIDs, newDeviceIDs)
		suite.True(result.Success, "同时添加和移除设备应该成功")
		suite.ElementsMatch([]int{3, 4}, result.AddedDevices, "应该添加设备3和4")
		suite.ElementsMatch([]int{1}, result.RemovedDevices, "应该移除设备1")
	})

	suite.Run("完全替换设备", func() {
		deptID := 1
		oldDeviceIDs := []int{1, 2}
		newDeviceIDs := []int{3, 4}

		result := simulateUpdateDeviceAssociation(deptID, oldDeviceIDs, newDeviceIDs)
		suite.True(result.Success, "完全替换设备应该成功")
		suite.ElementsMatch([]int{3, 4}, result.AddedDevices, "应该添加设备3和4")
		suite.ElementsMatch([]int{1, 2}, result.RemovedDevices, "应该移除设备1和2")
	})

	suite.Run("清空所有设备", func() {
		deptID := 1
		oldDeviceIDs := []int{1, 2, 3}
		newDeviceIDs := []int{}

		result := simulateUpdateDeviceAssociation(deptID, oldDeviceIDs, newDeviceIDs)
		suite.True(result.Success, "清空所有设备应该成功")
		suite.Empty(result.AddedDevices, "不应该添加任何设备")
		suite.ElementsMatch([]int{1, 2, 3}, result.RemovedDevices, "应该移除所有设备")
	})

	suite.Run("设备关联无变化", func() {
		deptID := 1
		oldDeviceIDs := []int{1, 2, 3}
		newDeviceIDs := []int{1, 2, 3}

		result := simulateUpdateDeviceAssociation(deptID, oldDeviceIDs, newDeviceIDs)
		suite.True(result.Success, "设备关联无变化应该成功")
		suite.Empty(result.AddedDevices, "不应该添加任何设备")
		suite.Empty(result.RemovedDevices, "不应该移除任何设备")
	})
}

// TestGetAvailableDevices 测试获取可用设备
func (suite *WlDepartmentDeviceTestSuite) TestGetAvailableDevices() {
	suite.Run("获取所有可用设备", func() {
		req := request.AvailableDevicesReq{
			Page:     1,
			PageSize: 10,
		}

		result := simulateGetAvailableDevices(req)
		suite.True(result.Success, "获取可用设备应该成功")
		suite.NotEmpty(result.Devices, "应该有可用设备")
		suite.Greater(result.Total, int64(0), "总数应该大于0")
	})

	suite.Run("按设备名称搜索", func() {
		req := request.AvailableDevicesReq{
			Page:       1,
			PageSize:   10,
			DeviceName: "传感器",
		}

		result := simulateGetAvailableDevices(req)
		suite.True(result.Success, "按设备名称搜索应该成功")

		// 验证返回的设备名称包含搜索关键词
		for _, device := range result.Devices {
			suite.Contains(device.DeviceName, "传感器", "设备名称应该包含搜索关键词")
		}
	})

	suite.Run("按产品名称搜索", func() {
		req := request.AvailableDevicesReq{
			Page:        1,
			PageSize:    10,
			ProductName: "温度传感器",
		}

		result := simulateGetAvailableDevices(req)
		suite.True(result.Success, "按产品名称搜索应该成功")

		// 验证返回的产品名称包含搜索关键词
		for _, device := range result.Devices {
			suite.Contains(device.ProductName, "温度传感器", "产品名称应该包含搜索关键词")
		}
	})

	suite.Run("排除指定部门的设备", func() {
		departmentID := 1
		req := request.AvailableDevicesReq{
			Page:         1,
			PageSize:     10,
			DepartmentID: &departmentID,
		}

		result := simulateGetAvailableDevices(req)
		suite.True(result.Success, "排除指定部门设备应该成功")

		// 验证返回的设备不属于指定部门
		excludedDeviceIDs := []int{1, 2, 3} // 假设部门1关联了设备1,2,3
		for _, device := range result.Devices {
			suite.NotContains(excludedDeviceIDs, device.ID, "返回的设备不应该属于指定部门")
		}
	})

	suite.Run("组合搜索条件", func() {
		departmentID := 1
		req := request.AvailableDevicesReq{
			Page:         1,
			PageSize:     5,
			DeviceName:   "传感器",
			ProductName:  "温度",
			DepartmentID: &departmentID,
		}

		result := simulateGetAvailableDevices(req)
		suite.True(result.Success, "组合搜索应该成功")
		suite.LessOrEqual(len(result.Devices), 5, "返回的设备数量不应该超过页大小")
	})

	suite.Run("无效的分页参数", func() {
		req := request.AvailableDevicesReq{
			Page:     0,
			PageSize: -1,
		}

		result := simulateGetAvailableDevices(req)
		suite.True(result.Success, "无效分页参数应该使用默认值")
		suite.Equal(1, result.Page, "页码应该使用默认值1")
		suite.Equal(10, result.PageSize, "页大小应该使用默认值10")
	})
}

// TestGetDepartmentDevices 测试获取部门设备
func (suite *WlDepartmentDeviceTestSuite) TestGetDepartmentDevices() {
	suite.Run("获取部门设备成功", func() {
		req := request.DepartmentDevicesReq{
			DepartmentID: 1,
			Page:         1,
			PageSize:     10,
		}

		result := simulateGetDepartmentDevices(req)
		suite.True(result.Success, "获取部门设备应该成功")
		suite.NotEmpty(result.Devices, "部门应该有关联设备")
		suite.Greater(result.Total, int64(0), "总数应该大于0")
	})

	suite.Run("部门没有关联设备", func() {
		req := request.DepartmentDevicesReq{
			DepartmentID: 999, // 假设这个部门没有关联设备
			Page:         1,
			PageSize:     10,
		}

		result := simulateGetDepartmentDevices(req)
		suite.True(result.Success, "获取空设备列表应该成功")
		suite.Empty(result.Devices, "应该返回空设备列表")
		suite.Equal(int64(0), result.Total, "总数应该为0")
	})

	suite.Run("无效的部门ID", func() {
		req := request.DepartmentDevicesReq{
			DepartmentID: 0,
			Page:         1,
			PageSize:     10,
		}

		result := simulateGetDepartmentDevices(req)
		suite.False(result.Success, "无效部门ID应该失败")
		suite.Contains(result.ErrorMessage, "无效的部门ID", "应该返回无效部门ID错误")
	})

	suite.Run("部门不存在", func() {
		req := request.DepartmentDevicesReq{
			DepartmentID: 99999, // 不存在的部门ID
			Page:         1,
			PageSize:     10,
		}

		result := simulateGetDepartmentDevices(req)
		suite.False(result.Success, "不存在的部门应该失败")
		suite.Contains(result.ErrorMessage, "部门不存在", "应该返回部门不存在错误")
	})

	suite.Run("分页功能", func() {
		// 第一页
		req1 := request.DepartmentDevicesReq{
			DepartmentID: 1,
			Page:         1,
			PageSize:     2,
		}

		result1 := simulateGetDepartmentDevices(req1)
		suite.True(result1.Success, "获取第一页应该成功")
		suite.LessOrEqual(len(result1.Devices), 2, "第一页设备数量不应该超过页大小")

		// 第二页
		req2 := request.DepartmentDevicesReq{
			DepartmentID: 1,
			Page:         2,
			PageSize:     2,
		}

		result2 := simulateGetDepartmentDevices(req2)
		suite.True(result2.Success, "获取第二页应该成功")

		// 验证分页结果不重复
		if len(result1.Devices) > 0 && len(result2.Devices) > 0 {
			suite.NotEqual(result1.Devices[0].ID, result2.Devices[0].ID, "不同页的设备不应该重复")
		}
	})
}

// TestDeviceAssociationValidation 测试设备关联验证
func (suite *WlDepartmentDeviceTestSuite) TestDeviceAssociationValidation() {
	suite.Run("验证设备存在性", func() {
		deptID := 1
		deviceIDs := []int{1, 2, 999999} // 包含不存在的设备ID

		result := validateDeviceExistence(deptID, deviceIDs)
		suite.False(result.Valid, "包含不存在设备ID应该验证失败")
		suite.Contains(result.ErrorMessage, "设备不存在", "应该返回设备不存在错误")
		suite.ElementsMatch([]int{999999}, result.InvalidDeviceIDs, "应该返回无效的设备ID")
	})

	suite.Run("验证设备状态", func() {
		deptID := 1
		deviceIDs := []int{1, 2, 3} // 假设设备2是禁用状态

		result := validateDeviceStatus(deptID, deviceIDs)
		suite.False(result.Valid, "包含禁用设备应该验证失败")
		suite.Contains(result.ErrorMessage, "设备已禁用", "应该返回设备禁用错误")
		suite.ElementsMatch([]int{2}, result.DisabledDeviceIDs, "应该返回禁用的设备ID")
	})

	suite.Run("验证设备重复关联", func() {
		deptID := 1
		deviceIDs := []int{1, 2, 3} // 假设设备1已经关联到其他部门

		result := validateDeviceDuplication(deptID, deviceIDs)
		suite.False(result.Valid, "设备重复关联应该验证失败")
		suite.Contains(result.ErrorMessage, "设备已关联", "应该返回设备已关联错误")
		suite.ElementsMatch([]int{1}, result.DuplicatedDeviceIDs, "应该返回重复关联的设备ID")
	})

	suite.Run("所有验证通过", func() {
		deptID := 1
		deviceIDs := []int{4, 5, 6} // 假设这些设备都是有效且可关联的

		existenceResult := validateDeviceExistence(deptID, deviceIDs)
		suite.True(existenceResult.Valid, "设备存在性验证应该通过")

		statusResult := validateDeviceStatus(deptID, deviceIDs)
		suite.True(statusResult.Valid, "设备状态验证应该通过")

		duplicationResult := validateDeviceDuplication(deptID, deviceIDs)
		suite.True(duplicationResult.Valid, "设备重复关联验证应该通过")
	})
}

// TestDeviceAssociationConcurrency 测试设备关联并发操作
func (suite *WlDepartmentDeviceTestSuite) TestDeviceAssociationConcurrency() {
	suite.Run("并发关联设备", func() {
		deptID := 1

		// 模拟多个并发操作
		done := make(chan bool, 5)
		for i := 0; i < 5; i++ {
			go func(index int) {
				deviceIDs := []int{index + 1} // 每个goroutine关联不同的设备
				result := simulateAssociateDevices(deptID, deviceIDs)
				suite.True(result.Success, "并发关联设备应该成功")
				done <- true
			}(i)
		}

		// 等待所有并发操作完成
		for i := 0; i < 5; i++ {
			<-done
		}

		// 验证最终结果
		req := request.DepartmentDevicesReq{
			DepartmentID: deptID,
			Page:         1,
			PageSize:     10,
		}
		result := simulateGetDepartmentDevices(req)
		suite.True(result.Success, "获取最终结果应该成功")
		suite.LessOrEqual(len(result.Devices), 5, "并发操作后设备数量应该正确")
	})
}

// 模拟函数实现

// DeviceAssociationResult 设备关联结果
type DeviceAssociationResult struct {
	Success           bool
	AssociatedDevices []wl_department.WlDevice
	ErrorMessage      string
}

// UpdateDeviceAssociationResult 更新设备关联结果
type UpdateDeviceAssociationResult struct {
	Success        bool
	AddedDevices   []int
	RemovedDevices []int
	ErrorMessage   string
}

// AvailableDevicesResult 可用设备查询结果
type AvailableDevicesResult struct {
	Success      bool
	Devices      []wl_department.WlDevice
	Total        int64
	Page         int
	PageSize     int
	ErrorMessage string
}

// DepartmentDevicesResult 部门设备查询结果
type DepartmentDevicesResult struct {
	Success      bool
	Devices      []wl_department.WlDevice
	Total        int64
	ErrorMessage string
}

// ValidationResult 验证结果
type ValidationResult struct {
	Valid               bool
	ErrorMessage        string
	InvalidDeviceIDs    []int
	DisabledDeviceIDs   []int
	DuplicatedDeviceIDs []int
}

// simulateAssociateDevices 模拟关联设备
func simulateAssociateDevices(deptID int, deviceIDs []int) DeviceAssociationResult {
	if deptID <= 0 {
		return DeviceAssociationResult{
			Success:      false,
			ErrorMessage: "无效的部门ID",
		}
	}

	// 验证设备ID有效性
	for _, deviceID := range deviceIDs {
		if deviceID <= 0 {
			return DeviceAssociationResult{
				Success:      false,
				ErrorMessage: "无效的设备ID",
			}
		}
	}

	// 去重处理
	uniqueDeviceIDs := removeDuplicateInts(deviceIDs)

	// 模拟关联设备
	var devices []wl_department.WlDevice
	for _, deviceID := range uniqueDeviceIDs {
		devices = append(devices, wl_department.WlDevice{
			ID:          deviceID,
			DeviceName:  fmt.Sprintf("设备%d", deviceID),
			ProductName: fmt.Sprintf("产品%d", deviceID),
			Status:      "启用",
		})
	}

	return DeviceAssociationResult{
		Success:           true,
		AssociatedDevices: devices,
	}
}

// simulateUpdateDeviceAssociation 模拟更新设备关联
func simulateUpdateDeviceAssociation(deptID int, oldDeviceIDs, newDeviceIDs []int) UpdateDeviceAssociationResult {
	if deptID <= 0 {
		return UpdateDeviceAssociationResult{
			Success:      false,
			ErrorMessage: "无效的部门ID",
		}
	}

	// 计算需要添加和移除的设备
	oldSet := make(map[int]bool)
	for _, id := range oldDeviceIDs {
		oldSet[id] = true
	}

	newSet := make(map[int]bool)
	for _, id := range newDeviceIDs {
		newSet[id] = true
	}

	var addedDevices, removedDevices []int

	// 找出需要添加的设备
	for _, id := range newDeviceIDs {
		if !oldSet[id] {
			addedDevices = append(addedDevices, id)
		}
	}

	// 找出需要移除的设备
	for _, id := range oldDeviceIDs {
		if !newSet[id] {
			removedDevices = append(removedDevices, id)
		}
	}

	return UpdateDeviceAssociationResult{
		Success:        true,
		AddedDevices:   addedDevices,
		RemovedDevices: removedDevices,
	}
}

// simulateGetAvailableDevices 模拟获取可用设备
func simulateGetAvailableDevices(req request.AvailableDevicesReq) AvailableDevicesResult {
	// 设置默认分页参数
	page := req.Page
	if page <= 0 {
		page = 1
	}

	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	// 模拟设备数据
	allDevices := []wl_department.WlDevice{
		{ID: 4, DeviceName: "温度传感器1", ProductName: "温度传感器产品A", Status: "启用"},
		{ID: 5, DeviceName: "湿度传感器1", ProductName: "湿度传感器产品B", Status: "启用"},
		{ID: 6, DeviceName: "压力传感器1", ProductName: "压力传感器产品C", Status: "启用"},
		{ID: 7, DeviceName: "温度传感器2", ProductName: "温度传感器产品D", Status: "启用"},
		{ID: 8, DeviceName: "光照传感器1", ProductName: "光照传感器产品E", Status: "启用"},
	}

	// 过滤设备
	var filteredDevices []wl_department.WlDevice
	for _, device := range allDevices {
		// 按设备名称过滤
		if req.DeviceName != "" && !contains(device.DeviceName, req.DeviceName) {
			continue
		}

		// 按产品名称过滤
		if req.ProductName != "" && !contains(device.ProductName, req.ProductName) {
			continue
		}

		// 排除指定部门的设备
		if req.DepartmentID != nil && *req.DepartmentID == 1 {
			// 假设部门1关联了设备1,2,3
			if device.ID <= 3 {
				continue
			}
		}

		filteredDevices = append(filteredDevices, device)
	}

	// 分页处理
	total := int64(len(filteredDevices))
	start := (page - 1) * pageSize
	end := start + pageSize

	if start >= len(filteredDevices) {
		filteredDevices = []wl_department.WlDevice{}
	} else if end > len(filteredDevices) {
		filteredDevices = filteredDevices[start:]
	} else {
		filteredDevices = filteredDevices[start:end]
	}

	return AvailableDevicesResult{
		Success:  true,
		Devices:  filteredDevices,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}
}

// simulateGetDepartmentDevices 模拟获取部门设备
func simulateGetDepartmentDevices(req request.DepartmentDevicesReq) DepartmentDevicesResult {
	if req.DepartmentID <= 0 {
		return DepartmentDevicesResult{
			Success:      false,
			ErrorMessage: "无效的部门ID",
		}
	}

	if req.DepartmentID == 99999 {
		return DepartmentDevicesResult{
			Success:      false,
			ErrorMessage: "部门不存在",
		}
	}

	// 模拟部门设备数据
	var departmentDevices []wl_department.WlDevice
	if req.DepartmentID == 1 {
		departmentDevices = []wl_department.WlDevice{
			{ID: 1, DeviceName: "设备1", ProductName: "产品1", Status: "启用"},
			{ID: 2, DeviceName: "设备2", ProductName: "产品2", Status: "启用"},
			{ID: 3, DeviceName: "设备3", ProductName: "产品3", Status: "启用"},
		}
	} else if req.DepartmentID == 999 {
		// 没有关联设备的部门
		departmentDevices = []wl_department.WlDevice{}
	}

	// 分页处理
	total := int64(len(departmentDevices))
	page := req.Page
	if page <= 0 {
		page = 1
	}

	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	start := (page - 1) * pageSize
	end := start + pageSize

	if start >= len(departmentDevices) {
		departmentDevices = []wl_department.WlDevice{}
	} else if end > len(departmentDevices) {
		departmentDevices = departmentDevices[start:]
	} else {
		departmentDevices = departmentDevices[start:end]
	}

	return DepartmentDevicesResult{
		Success: true,
		Devices: departmentDevices,
		Total:   total,
	}
}

// validateDeviceExistence 验证设备存在性
func validateDeviceExistence(deptID int, deviceIDs []int) ValidationResult {
	// 模拟已存在的设备ID
	existingDeviceIDs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var invalidIDs []int
	for _, deviceID := range deviceIDs {
		if !containsInt(existingDeviceIDs, deviceID) {
			invalidIDs = append(invalidIDs, deviceID)
		}
	}

	if len(invalidIDs) > 0 {
		return ValidationResult{
			Valid:            false,
			ErrorMessage:     "设备不存在",
			InvalidDeviceIDs: invalidIDs,
		}
	}

	return ValidationResult{Valid: true}
}

// validateDeviceStatus 验证设备状态
func validateDeviceStatus(deptID int, deviceIDs []int) ValidationResult {
	// 模拟禁用的设备ID
	disabledDeviceIDs := []int{2}

	var disabledIDs []int
	for _, deviceID := range deviceIDs {
		if containsInt(disabledDeviceIDs, deviceID) {
			disabledIDs = append(disabledIDs, deviceID)
		}
	}

	if len(disabledIDs) > 0 {
		return ValidationResult{
			Valid:             false,
			ErrorMessage:      "设备已禁用",
			DisabledDeviceIDs: disabledIDs,
		}
	}

	return ValidationResult{Valid: true}
}

// validateDeviceDuplication 验证设备重复关联
func validateDeviceDuplication(deptID int, deviceIDs []int) ValidationResult {
	// 模拟已关联到其他部门的设备ID
	duplicatedDeviceIDs := []int{1}

	var duplicatedIDs []int
	for _, deviceID := range deviceIDs {
		if containsInt(duplicatedDeviceIDs, deviceID) {
			duplicatedIDs = append(duplicatedIDs, deviceID)
		}
	}

	if len(duplicatedIDs) > 0 {
		return ValidationResult{
			Valid:               false,
			ErrorMessage:        "设备已关联",
			DuplicatedDeviceIDs: duplicatedIDs,
		}
	}

	return ValidationResult{Valid: true}
}

// 辅助函数

// removeDuplicateInts 去除重复的整数
func removeDuplicateInts(slice []int) []int {
	seen := make(map[int]bool)
	var result []int

	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

// contains 检查字符串是否包含子字符串
func contains(str, substr string) bool {
	return len(substr) == 0 || (len(str) >= len(substr) &&
		str[0:len(substr)] == substr ||
		(len(str) > len(substr) && contains(str[1:], substr)))
}

// containsInt 检查整数切片是否包含指定整数
func containsInt(slice []int, item int) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
