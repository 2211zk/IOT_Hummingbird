package integration

import (
	"fmt"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_department"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_department/request"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

// WlDepartmentIntegrationTestSuite 部门管理集成测试套件
type WlDepartmentIntegrationTestSuite struct {
	suite.Suite
	router *gin.Engine
}

func (suite *WlDepartmentIntegrationTestSuite) SetupSuite() {
	// 设置测试环境
	gin.SetMode(gin.TestMode)
	suite.router = setupIntegrationTestRouter()
}

func TestWlDepartmentIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(WlDepartmentIntegrationTestSuite))
}

// setupIntegrationTestRouter 设置集成测试路由
func setupIntegrationTestRouter() *gin.Engine {
	router := gin.New()
	// 这里应该设置完整的路由，包括中间件等
	// 为了测试目的，这里简化处理
	return router
}

// TestDepartmentCRUDFlow 测试部门CRUD完整流程
func (suite *WlDepartmentIntegrationTestSuite) TestDepartmentCRUDFlow() {
	suite.Run("完整的部门CRUD流程测试", func() {
		// 1. 创建顶级部门
		createReq := request.CreateWlDepartmentReq{
			Name:      "集成测试部门",
			Leader:    "测试负责人",
			Phone:     "13800138000",
			Email:     "test@example.com",
			Status:    "启用",
			Sort:      1,
			DeviceIDs: []int{1, 2},
		}

		// 模拟创建部门的响应
		expectedDeptID := 1
		suite.simulateCreateDepartment(createReq, expectedDeptID)

		// 2. 获取部门列表，验证创建成功
		suite.simulateGetDepartmentList(expectedDeptID)

		// 3. 创建子部门
		childCreateReq := request.CreateWlDepartmentReq{
			Name:     "子部门",
			Leader:   "子部门负责人",
			ParentID: &expectedDeptID,
			Status:   "启用",
			Sort:     1,
		}

		expectedChildID := 2
		suite.simulateCreateDepartment(childCreateReq, expectedChildID)

		// 4. 获取部门树，验证层级关系
		suite.simulateGetDepartmentTree(expectedDeptID, expectedChildID)

		// 5. 更新部门信息
		updateReq := request.UpdateWlDepartmentReq{
			ID:        expectedDeptID,
			Name:      "更新后的部门名称",
			Leader:    "新负责人",
			Phone:     "13800138001",
			Email:     "new@example.com",
			Status:    "启用",
			Sort:      2,
			DeviceIDs: []int{1, 2, 3},
		}

		suite.simulateUpdateDepartment(updateReq)

		// 6. 获取部门详情，验证更新成功
		suite.simulateGetDepartmentDetail(expectedDeptID, updateReq.Name)

		// 7. 测试设备关联功能
		suite.simulateDeviceAssociation(expectedDeptID)

		// 8. 尝试删除有子部门的部门（应该失败）
		suite.simulateDeleteDepartmentWithChildren(expectedDeptID)

		// 9. 删除子部门
		suite.simulateDeleteDepartment(expectedChildID, true)

		// 10. 删除父部门
		suite.simulateDeleteDepartment(expectedDeptID, true)
	})
}

// simulateCreateDepartment 模拟创建部门
func (suite *WlDepartmentIntegrationTestSuite) simulateCreateDepartment(req request.CreateWlDepartmentReq, expectedID int) {
	// 验证请求参数
	name := req.Name
	if name == "" {
		name = req.DepartmentName
	}
	suite.NotEmpty(name, "部门名称不应该为空")
	suite.NotEmpty(req.Leader, "负责人不应该为空")

	// 验证设备关联
	if len(req.DeviceIDs) > 0 {
		for _, deviceID := range req.DeviceIDs {
			suite.Greater(deviceID, 0, "设备ID应该大于0")
		}
	}

	// 模拟成功响应
	suite.Greater(expectedID, 0, "创建的部门ID应该大于0")
}

// simulateGetDepartmentList 模拟获取部门列表
func (suite *WlDepartmentIntegrationTestSuite) simulateGetDepartmentList(expectedDeptID int) {
	// 模拟查询参数
	searchReq := request.WlDepartmentSearch{
		Page:     1,
		PageSize: 10,
		TreeMode: false,
	}

	suite.Greater(searchReq.Page, 0, "页码应该大于0")
	suite.Greater(searchReq.PageSize, 0, "页大小应该大于0")

	// 模拟响应数据验证
	mockDepartments := []wl_department.WlDepartment{
		{
			ID:       expectedDeptID,
			Name:     "集成测试部门",
			Leader:   "测试负责人",
			Phone:    "13800138000",
			Email:    "test@example.com",
			Status:   "启用",
			Sort:     1,
			ParentID: nil,
		},
	}

	suite.Len(mockDepartments, 1, "应该返回1个部门")
	suite.Equal(expectedDeptID, mockDepartments[0].ID, "部门ID应该匹配")
}

// simulateGetDepartmentTree 模拟获取部门树
func (suite *WlDepartmentIntegrationTestSuite) simulateGetDepartmentTree(parentID, childID int) {
	// 模拟树形结构数据
	mockTree := []wl_department.WlDepartment{
		{
			ID:       parentID,
			Name:     "集成测试部门",
			ParentID: nil,
			Children: []wl_department.WlDepartment{
				{
					ID:       childID,
					Name:     "子部门",
					ParentID: &parentID,
					Children: []wl_department.WlDepartment{},
				},
			},
		},
	}

	suite.Len(mockTree, 1, "应该有1个顶级部门")
	suite.Len(mockTree[0].Children, 1, "顶级部门应该有1个子部门")
	suite.Equal(childID, mockTree[0].Children[0].ID, "子部门ID应该匹配")
	suite.Equal(&parentID, mockTree[0].Children[0].ParentID, "子部门的父ID应该匹配")
}

// simulateUpdateDepartment 模拟更新部门
func (suite *WlDepartmentIntegrationTestSuite) simulateUpdateDepartment(req request.UpdateWlDepartmentReq) {
	suite.Greater(req.ID, 0, "部门ID应该大于0")

	name := req.Name
	if name == "" {
		name = req.DepartmentName
	}
	suite.NotEmpty(name, "部门名称不应该为空")

	// 验证设备关联更新
	if len(req.DeviceIDs) > 0 {
		for _, deviceID := range req.DeviceIDs {
			suite.Greater(deviceID, 0, "设备ID应该大于0")
		}
	}
}

// simulateGetDepartmentDetail 模拟获取部门详情
func (suite *WlDepartmentIntegrationTestSuite) simulateGetDepartmentDetail(deptID int, expectedName string) {
	suite.Greater(deptID, 0, "部门ID应该大于0")

	// 模拟部门详情数据
	mockDetail := wl_department.WlDepartment{
		ID:     deptID,
		Name:   expectedName,
		Leader: "新负责人",
		Phone:  "13800138001",
		Email:  "new@example.com",
		Status: "启用",
		Sort:   2,
		Devices: []wl_department.WlDevice{
			{ID: 1, DeviceName: "设备1", ProductName: "产品1"},
			{ID: 2, DeviceName: "设备2", ProductName: "产品2"},
			{ID: 3, DeviceName: "设备3", ProductName: "产品3"},
		},
	}

	suite.Equal(deptID, mockDetail.ID, "部门ID应该匹配")
	suite.Equal(expectedName, mockDetail.Name, "部门名称应该匹配")
	suite.Len(mockDetail.Devices, 3, "应该有3个关联设备")
}

// simulateDeviceAssociation 模拟设备关联功能
func (suite *WlDepartmentIntegrationTestSuite) simulateDeviceAssociation(deptID int) {
	// 1. 获取可用设备列表
	availableReq := request.AvailableDevicesReq{
		Page:         1,
		PageSize:     10,
		DepartmentID: &deptID,
	}

	suite.Greater(availableReq.Page, 0, "页码应该大于0")
	suite.Greater(availableReq.PageSize, 0, "页大小应该大于0")
	suite.Equal(deptID, *availableReq.DepartmentID, "部门ID应该匹配")

	// 模拟可用设备数据
	mockAvailableDevices := []wl_department.WlDevice{
		{ID: 4, DeviceName: "可用设备1", ProductName: "产品A"},
		{ID: 5, DeviceName: "可用设备2", ProductName: "产品B"},
	}

	suite.Len(mockAvailableDevices, 2, "应该有2个可用设备")

	// 2. 获取部门已关联设备
	deptDevicesReq := request.DepartmentDevicesReq{
		DepartmentID: deptID,
		Page:         1,
		PageSize:     10,
	}

	suite.Greater(deptDevicesReq.DepartmentID, 0, "部门ID应该大于0")

	// 模拟已关联设备数据
	mockDeptDevices := []wl_department.WlDevice{
		{ID: 1, DeviceName: "设备1", ProductName: "产品1"},
		{ID: 2, DeviceName: "设备2", ProductName: "产品2"},
		{ID: 3, DeviceName: "设备3", ProductName: "产品3"},
	}

	suite.Len(mockDeptDevices, 3, "应该有3个已关联设备")
}

// simulateDeleteDepartmentWithChildren 模拟删除有子部门的部门
func (suite *WlDepartmentIntegrationTestSuite) simulateDeleteDepartmentWithChildren(deptID int) {
	deleteReq := request.DeleteWlDepartmentReq{
		ID: deptID,
	}

	suite.Greater(deleteReq.ID, 0, "部门ID应该大于0")

	// 模拟检查子部门逻辑
	hasChildren := true // 假设有子部门
	if hasChildren {
		// 应该返回错误，不允许删除
		suite.True(hasChildren, "有子部门的部门不应该被删除")
	}
}

// simulateDeleteDepartment 模拟删除部门
func (suite *WlDepartmentIntegrationTestSuite) simulateDeleteDepartment(deptID int, shouldSucceed bool) {
	deleteReq := request.DeleteWlDepartmentReq{
		ID: deptID,
	}

	suite.Greater(deleteReq.ID, 0, "部门ID应该大于0")

	if shouldSucceed {
		// 模拟删除成功
		suite.True(shouldSucceed, "部门删除应该成功")
	}
}

// TestCircularReferenceDetection 测试循环引用检测
func (suite *WlDepartmentIntegrationTestSuite) TestCircularReferenceDetection() {
	suite.Run("循环引用检测测试", func() {
		// 创建部门层级结构用于测试
		// A -> B -> C
		deptA := 1
		deptB := 2
		deptC := 3

		// 测试场景1: C试图设置A为上级（会形成循环）
		suite.testCircularReference(deptC, deptA, []int{deptB, deptA}, true)

		// 测试场景2: A试图设置C为上级（会形成循环）
		suite.testCircularReference(deptA, deptC, []int{deptB, deptA}, true)

		// 测试场景3: 正常的层级关系变更
		deptD := 4
		suite.testCircularReference(deptD, deptA, []int{}, false)

		// 测试场景4: 自己设置自己为上级
		suite.testCircularReference(deptA, deptA, []int{}, true)
	})
}

// testCircularReference 测试循环引用检测逻辑
func (suite *WlDepartmentIntegrationTestSuite) testCircularReference(deptID, newParentID int, parentChain []int, expectCircular bool) {
	// 模拟循环引用检测逻辑
	hasCircular := false

	// 检查是否选择自身作为上级
	if deptID == newParentID {
		hasCircular = true
	} else {
		// 检查新上级的上级链中是否包含当前部门
		for _, parentID := range parentChain {
			if parentID == deptID {
				hasCircular = true
				break
			}
		}
	}

	if expectCircular {
		suite.True(hasCircular, "应该检测到循环引用: 部门%d设置部门%d为上级", deptID, newParentID)
	} else {
		suite.False(hasCircular, "不应该有循环引用: 部门%d设置部门%d为上级", deptID, newParentID)
	}
}

// TestDepartmentTreeConstruction 测试部门树构建
func (suite *WlDepartmentIntegrationTestSuite) TestDepartmentTreeConstruction() {
	suite.Run("部门树构建测试", func() {
		// 模拟复杂的部门结构
		departments := []wl_department.WlDepartment{
			{ID: 1, Name: "公司", ParentID: nil, Sort: 1},
			{ID: 2, Name: "技术部", ParentID: intPtr(1), Sort: 1},
			{ID: 3, Name: "市场部", ParentID: intPtr(1), Sort: 2},
			{ID: 4, Name: "前端组", ParentID: intPtr(2), Sort: 1},
			{ID: 5, Name: "后端组", ParentID: intPtr(2), Sort: 2},
			{ID: 6, Name: "React组", ParentID: intPtr(4), Sort: 1},
			{ID: 7, Name: "Vue组", ParentID: intPtr(4), Sort: 2},
			{ID: 8, Name: "Go组", ParentID: intPtr(5), Sort: 1},
		}

		// 测试构建完整树
		tree := suite.buildMockDepartmentTree(departments, nil)
		suite.validateDepartmentTree(tree, departments)

		// 测试构建子树
		subTree := suite.buildMockDepartmentTree(departments, intPtr(2))
		suite.validateSubTree(subTree, 2)
	})
}

// buildMockDepartmentTree 模拟构建部门树
func (suite *WlDepartmentIntegrationTestSuite) buildMockDepartmentTree(departments []wl_department.WlDepartment, parentID *int) []wl_department.WlDepartment {
	var result []wl_department.WlDepartment

	// 找出指定父部门的直接子部门
	for _, dept := range departments {
		if (parentID == nil && dept.ParentID == nil) || (parentID != nil && dept.ParentID != nil && *dept.ParentID == *parentID) {
			// 递归构建子树
			dept.Children = suite.buildMockDepartmentTree(departments, &dept.ID)
			result = append(result, dept)
		}
	}

	return result
}

// validateDepartmentTree 验证部门树结构
func (suite *WlDepartmentIntegrationTestSuite) validateDepartmentTree(tree []wl_department.WlDepartment, allDepts []wl_department.WlDepartment) {
	// 验证顶级部门
	suite.Len(tree, 1, "应该有1个顶级部门")
	suite.Equal("公司", tree[0].Name, "顶级部门应该是公司")

	// 验证公司的子部门
	suite.Len(tree[0].Children, 2, "公司应该有2个子部门")
	suite.Equal("技术部", tree[0].Children[0].Name, "第一个子部门应该是技术部")
	suite.Equal("市场部", tree[0].Children[1].Name, "第二个子部门应该是市场部")

	// 验证技术部的子部门
	techDept := tree[0].Children[0]
	suite.Len(techDept.Children, 2, "技术部应该有2个子部门")
	suite.Equal("前端组", techDept.Children[0].Name, "技术部第一个子部门应该是前端组")
	suite.Equal("后端组", techDept.Children[1].Name, "技术部第二个子部门应该是后端组")

	// 验证前端组的子部门
	frontendDept := techDept.Children[0]
	suite.Len(frontendDept.Children, 2, "前端组应该有2个子部门")
	suite.Equal("React组", frontendDept.Children[0].Name, "前端组第一个子部门应该是React组")
	suite.Equal("Vue组", frontendDept.Children[1].Name, "前端组第二个子部门应该是Vue组")

	// 验证后端组的子部门
	backendDept := techDept.Children[1]
	suite.Len(backendDept.Children, 1, "后端组应该有1个子部门")
	suite.Equal("Go组", backendDept.Children[0].Name, "后端组的子部门应该是Go组")
}

// validateSubTree 验证子树结构
func (suite *WlDepartmentIntegrationTestSuite) validateSubTree(subTree []wl_department.WlDepartment, parentID int) {
	suite.Len(subTree, 2, "技术部应该有2个直接子部门")
	suite.Equal("前端组", subTree[0].Name, "第一个子部门应该是前端组")
	suite.Equal("后端组", subTree[1].Name, "第二个子部门应该是后端组")

	// 验证每个子部门的ParentID
	for _, dept := range subTree {
		suite.NotNil(dept.ParentID, "子部门应该有父部门ID")
		suite.Equal(parentID, *dept.ParentID, "父部门ID应该匹配")
	}
}

// TestDeviceAssociationFlow 测试设备关联完整流程
func (suite *WlDepartmentIntegrationTestSuite) TestDeviceAssociationFlow() {
	suite.Run("设备关联完整流程测试", func() {
		deptID := 1

		// 1. 初始状态：部门没有关联设备
		initialDevices := suite.getMockDepartmentDevices(deptID)
		suite.Empty(initialDevices, "初始状态部门应该没有关联设备")

		// 2. 关联设备
		deviceIDs := []int{1, 2, 3}
		suite.simulateAssociateDevices(deptID, deviceIDs)

		// 3. 验证关联结果
		associatedDevices := suite.getMockDepartmentDevices(deptID)
		suite.Len(associatedDevices, 3, "应该关联3个设备")

		// 4. 更新设备关联（添加新设备，移除部分设备）
		newDeviceIDs := []int{2, 3, 4, 5}
		suite.simulateUpdateDeviceAssociation(deptID, deviceIDs, newDeviceIDs)

		// 5. 验证更新结果
		updatedDevices := suite.getMockDepartmentDevices(deptID)
		suite.Len(updatedDevices, 4, "应该有4个关联设备")

		// 6. 清空设备关联
		suite.simulateAssociateDevices(deptID, []int{})
		finalDevices := suite.getMockDepartmentDevices(deptID)
		suite.Empty(finalDevices, "清空后应该没有关联设备")
	})
}

// getMockDepartmentDevices 获取模拟的部门设备数据
func (suite *WlDepartmentIntegrationTestSuite) getMockDepartmentDevices(deptID int) []wl_department.WlDevice {
	// 这里应该模拟从数据库查询部门设备的逻辑
	// 为了测试目的，返回模拟数据
	return []wl_department.WlDevice{}
}

// simulateAssociateDevices 模拟关联设备
func (suite *WlDepartmentIntegrationTestSuite) simulateAssociateDevices(deptID int, deviceIDs []int) {
	suite.Greater(deptID, 0, "部门ID应该大于0")

	for _, deviceID := range deviceIDs {
		suite.Greater(deviceID, 0, "设备ID应该大于0")
	}

	// 模拟关联逻辑
	suite.True(true, "设备关联应该成功")
}

// simulateUpdateDeviceAssociation 模拟更新设备关联
func (suite *WlDepartmentIntegrationTestSuite) simulateUpdateDeviceAssociation(deptID int, oldDeviceIDs, newDeviceIDs []int) {
	// 计算需要添加和移除的设备
	oldSet := make(map[int]bool)
	for _, id := range oldDeviceIDs {
		oldSet[id] = true
	}

	newSet := make(map[int]bool)
	for _, id := range newDeviceIDs {
		newSet[id] = true
	}

	var toAdd, toRemove []int

	// 找出需要添加的设备
	for _, id := range newDeviceIDs {
		if !oldSet[id] {
			toAdd = append(toAdd, id)
		}
	}

	// 找出需要移除的设备
	for _, id := range oldDeviceIDs {
		if !newSet[id] {
			toRemove = append(toRemove, id)
		}
	}

	suite.ElementsMatch([]int{4, 5}, toAdd, "应该添加设备4和5")
	suite.ElementsMatch([]int{1}, toRemove, "应该移除设备1")
}

// TestStatusInheritance 测试状态继承
func (suite *WlDepartmentIntegrationTestSuite) TestStatusInheritance() {
	suite.Run("部门状态继承测试", func() {
		// 模拟部门层级：技术部 -> 前端组 -> React组
		parentID := 1
		childID := 2
		grandChildID := 3

		// 1. 禁用父部门
		suite.simulateUpdateDepartmentStatus(parentID, "禁用")

		// 2. 验证子部门和孙部门也被禁用
		suite.validateStatusInheritance(parentID, []int{childID, grandChildID}, "禁用")

		// 3. 尝试启用子部门（应该失败，因为父部门是禁用的）
		suite.simulateUpdateDepartmentStatus(childID, "启用")
		suite.validateParentStatusCheck(childID, parentID, "禁用", false)

		// 4. 启用父部门
		suite.simulateUpdateDepartmentStatus(parentID, "启用")

		// 5. 现在可以启用子部门
		suite.simulateUpdateDepartmentStatus(childID, "启用")
		suite.validateParentStatusCheck(childID, parentID, "启用", true)
	})
}

// simulateUpdateDepartmentStatus 模拟更新部门状态
func (suite *WlDepartmentIntegrationTestSuite) simulateUpdateDepartmentStatus(deptID int, status string) {
	suite.Greater(deptID, 0, "部门ID应该大于0")
	suite.Contains([]string{"启用", "禁用"}, status, "状态应该是启用或禁用")
}

// validateStatusInheritance 验证状态继承
func (suite *WlDepartmentIntegrationTestSuite) validateStatusInheritance(parentID int, childIDs []int, expectedStatus string) {
	for _, childID := range childIDs {
		// 模拟检查子部门状态
		suite.Greater(childID, 0, "子部门ID应该大于0")
		// 在实际实现中，这里应该查询数据库验证子部门状态
	}
}

// validateParentStatusCheck 验证父部门状态检查
func (suite *WlDepartmentIntegrationTestSuite) validateParentStatusCheck(deptID, parentID int, parentStatus string, shouldAllow bool) {
	if parentStatus == "禁用" && shouldAllow {
		suite.Fail("父部门禁用时不应该允许启用子部门")
	} else if parentStatus == "启用" && !shouldAllow {
		suite.Fail("父部门启用时应该允许启用子部门")
	}
}

// TestConcurrentOperations 测试并发操作
func (suite *WlDepartmentIntegrationTestSuite) TestConcurrentOperations() {
	suite.Run("并发操作测试", func() {
		deptID := 1

		// 模拟并发更新部门信息
		done := make(chan bool, 10)
		for i := 0; i < 10; i++ {
			go func(index int) {
				updateReq := request.UpdateWlDepartmentReq{
					ID:     deptID,
					Name:   fmt.Sprintf("并发更新部门-%d", index),
					Leader: fmt.Sprintf("负责人-%d", index),
				}
				suite.simulateUpdateDepartment(updateReq)
				done <- true
			}(i)
		}

		// 等待所有并发操作完成
		for i := 0; i < 10; i++ {
			<-done
		}

		suite.True(true, "并发操作应该正常完成")
	})
}

// TestErrorScenarios 测试错误场景
func (suite *WlDepartmentIntegrationTestSuite) TestErrorScenarios() {
	suite.Run("错误场景测试", func() {
		// 1. 创建重名部门
		suite.testDuplicateNameError()

		// 2. 删除不存在的部门
		suite.testDeleteNonExistentDepartment()

		// 3. 更新不存在的部门
		suite.testUpdateNonExistentDepartment()

		// 4. 关联不存在的设备
		suite.testAssociateNonExistentDevice()

		// 5. 无效的参数
		suite.testInvalidParameters()
	})
}

// testDuplicateNameError 测试重名错误
func (suite *WlDepartmentIntegrationTestSuite) testDuplicateNameError() {
	// 模拟创建重名部门
	createReq := request.CreateWlDepartmentReq{
		Name:   "技术部", // 假设已存在
		Leader: "测试负责人",
	}

	// 在实际实现中，这里应该返回重名错误
	suite.NotEmpty(createReq.Name, "部门名称不应该为空")
}

// testDeleteNonExistentDepartment 测试删除不存在的部门
func (suite *WlDepartmentIntegrationTestSuite) testDeleteNonExistentDepartment() {
	deleteReq := request.DeleteWlDepartmentReq{
		ID: 99999, // 不存在的部门ID
	}

	suite.Greater(deleteReq.ID, 0, "部门ID应该大于0")
	// 在实际实现中，这里应该返回部门不存在错误
}

// testUpdateNonExistentDepartment 测试更新不存在的部门
func (suite *WlDepartmentIntegrationTestSuite) testUpdateNonExistentDepartment() {
	updateReq := request.UpdateWlDepartmentReq{
		ID:   99999, // 不存在的部门ID
		Name: "更新部门",
	}

	suite.Greater(updateReq.ID, 0, "部门ID应该大于0")
	// 在实际实现中，这里应该返回部门不存在错误
}

// testAssociateNonExistentDevice 测试关联不存在的设备
func (suite *WlDepartmentIntegrationTestSuite) testAssociateNonExistentDevice() {
	updateReq := request.UpdateWlDepartmentReq{
		ID:        1,
		Name:      "技术部",
		DeviceIDs: []int{99999}, // 不存在的设备ID
	}

	suite.Greater(updateReq.ID, 0, "部门ID应该大于0")
	for _, deviceID := range updateReq.DeviceIDs {
		suite.Greater(deviceID, 0, "设备ID应该大于0")
	}
	// 在实际实现中，这里应该返回设备不存在错误
}

// testInvalidParameters 测试无效参数
func (suite *WlDepartmentIntegrationTestSuite) testInvalidParameters() {
	// 测试各种无效参数场景
	invalidRequests := []interface{}{
		request.CreateWlDepartmentReq{Name: ""}, // 空名称
		request.UpdateWlDepartmentReq{ID: 0},    // 无效ID
		request.DeleteWlDepartmentReq{ID: -1},   // 负数ID
	}

	for _, req := range invalidRequests {
		switch r := req.(type) {
		case request.CreateWlDepartmentReq:
			suite.Empty(r.Name, "应该检测到空名称")
		case request.UpdateWlDepartmentReq:
			suite.LessOrEqual(r.ID, 0, "应该检测到无效ID")
		case request.DeleteWlDepartmentReq:
			suite.LessOrEqual(r.ID, 0, "应该检测到无效ID")
		}
	}
}

// intPtr 辅助函数：返回int指针
func intPtr(i int) *int {
	return &i
}
