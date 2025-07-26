package unit

import (
	"fmt"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_department"
	"github.com/stretchr/testify/suite"
)

// WlDepartmentTreeTestSuite 部门树结构测试套件
type WlDepartmentTreeTestSuite struct {
	suite.Suite
}

func TestWlDepartmentTreeTestSuite(t *testing.T) {
	suite.Run(t, new(WlDepartmentTreeTestSuite))
}

// TestBuildDepartmentTree 测试构建部门树的各种场景
func (suite *WlDepartmentTreeTestSuite) TestBuildDepartmentTree() {
	suite.Run("空数据构建树", func() {
		emptyDepts := []wl_department.WlDepartment{}
		tree := buildDepartmentTree(emptyDepts, nil)
		suite.Empty(tree, "空数据应该返回空树")
	})

	suite.Run("单个顶级部门", func() {
		singleDept := []wl_department.WlDepartment{
			{ID: 1, Name: "技术部", ParentID: nil, Sort: 1},
		}
		tree := buildDepartmentTree(singleDept, nil)
		suite.Len(tree, 1, "应该返回1个部门")
		suite.Equal("技术部", tree[0].Name)
		suite.Empty(tree[0].Children, "单个部门应该没有子节点")
	})

	suite.Run("简单的两层结构", func() {
		depts := []wl_department.WlDepartment{
			{ID: 1, Name: "技术部", ParentID: nil, Sort: 1},
			{ID: 2, Name: "前端组", ParentID: uintPtr(1), Sort: 1},
			{ID: 3, Name: "后端组", ParentID: uintPtr(1), Sort: 2},
		}
		tree := buildDepartmentTree(depts, nil)

		suite.Len(tree, 1, "应该有1个顶级部门")
		suite.Equal("技术部", tree[0].Name)
		suite.Len(tree[0].Children, 2, "技术部应该有2个子部门")
		suite.Equal("前端组", tree[0].Children[0].Name)
		suite.Equal("后端组", tree[0].Children[1].Name)
	})

	suite.Run("多个顶级部门", func() {
		depts := []wl_department.WlDepartment{
			{ID: 1, Name: "技术部", ParentID: nil, Sort: 1},
			{ID: 2, Name: "市场部", ParentID: nil, Sort: 2},
			{ID: 3, Name: "前端组", ParentID: uintPtr(1), Sort: 1},
		}
		tree := buildDepartmentTree(depts, nil)

		suite.Len(tree, 2, "应该有2个顶级部门")
		suite.Equal("技术部", tree[0].Name)
		suite.Equal("市场部", tree[1].Name)
		suite.Len(tree[0].Children, 1, "技术部应该有1个子部门")
		suite.Empty(tree[1].Children, "市场部应该没有子部门")
	})

	suite.Run("深层嵌套结构", func() {
		depts := []wl_department.WlDepartment{
			{ID: 1, Name: "公司", ParentID: nil, Sort: 1},
			{ID: 2, Name: "技术部", ParentID: uintPtr(1), Sort: 1},
			{ID: 3, Name: "前端组", ParentID: uintPtr(2), Sort: 1},
			{ID: 4, Name: "React组", ParentID: uintPtr(3), Sort: 1},
			{ID: 5, Name: "Vue组", ParentID: uintPtr(3), Sort: 2},
		}
		tree := buildDepartmentTree(depts, nil)

		suite.Len(tree, 1, "应该有1个顶级部门")
		suite.Equal("公司", tree[0].Name)

		// 验证技术部
		techDept := tree[0].Children[0]
		suite.Equal("技术部", techDept.Name)

		// 验证前端组
		frontendDept := techDept.Children[0]
		suite.Equal("前端组", frontendDept.Name)
		suite.Len(frontendDept.Children, 2, "前端组应该有2个子部门")

		// 验证React组和Vue组
		suite.Equal("React组", frontendDept.Children[0].Name)
		suite.Equal("Vue组", frontendDept.Children[1].Name)
	})

	suite.Run("构建指定父部门的子树", func() {
		depts := []wl_department.WlDepartment{
			{ID: 1, Name: "技术部", ParentID: nil, Sort: 1},
			{ID: 2, Name: "前端组", ParentID: uintPtr(1), Sort: 1},
			{ID: 3, Name: "后端组", ParentID: uintPtr(1), Sort: 2},
			{ID: 4, Name: "React组", ParentID: uintPtr(2), Sort: 1},
		}

		// 构建技术部的子树
		subTree := buildDepartmentTree(depts, uintPtr(1))
		suite.Len(subTree, 2, "技术部应该有2个直接子部门")
		suite.Equal("前端组", subTree[0].Name)
		suite.Equal("后端组", subTree[1].Name)
		suite.Len(subTree[0].Children, 1, "前端组应该有1个子部门")
		suite.Equal("React组", subTree[0].Children[0].Name)
	})

	suite.Run("排序功能测试", func() {
		depts := []wl_department.WlDepartment{
			{ID: 1, Name: "技术部", ParentID: nil, Sort: 2},
			{ID: 2, Name: "市场部", ParentID: nil, Sort: 1},
			{ID: 3, Name: "后端组", ParentID: uintPtr(1), Sort: 2},
			{ID: 4, Name: "前端组", ParentID: uintPtr(1), Sort: 1},
		}
		tree := buildDepartmentTree(depts, nil)

		// 验证顶级部门按排序排列
		suite.Equal("市场部", tree[0].Name, "市场部排序为1，应该在前面")
		suite.Equal("技术部", tree[1].Name, "技术部排序为2，应该在后面")

		// 验证子部门按排序排列
		techDept := tree[1]
		suite.Equal("前端组", techDept.Children[0].Name, "前端组排序为1，应该在前面")
		suite.Equal("后端组", techDept.Children[1].Name, "后端组排序为2，应该在后面")
	})

	suite.Run("处理孤儿部门", func() {
		depts := []wl_department.WlDepartment{
			{ID: 1, Name: "技术部", ParentID: nil, Sort: 1},
			{ID: 2, Name: "孤儿部门", ParentID: uintPtr(999), Sort: 1}, // 父部门不存在
		}
		tree := buildDepartmentTree(depts, nil)

		suite.Len(tree, 1, "孤儿部门应该被忽略，只返回有效的顶级部门")
		suite.Equal("技术部", tree[0].Name)
	})
}

// TestCircularReferenceDetection 测试循环引用检测
func (suite *WlDepartmentTreeTestSuite) TestCircularReferenceDetection() {
	suite.Run("直接循环引用", func() {
		// A -> A (自己指向自己)
		hasCircular := checkCircularReference(1, 1, []int{})
		suite.True(hasCircular, "自己指向自己应该被检测为循环引用")
	})

	suite.Run("两层循环引用", func() {
		// A -> B -> A
		parentChain := []int{1} // B的上级链包含A
		hasCircular := checkCircularReference(1, 2, parentChain)
		suite.True(hasCircular, "A设置B为上级，但B的上级链包含A，应该被检测为循环引用")
	})

	suite.Run("三层循环引用", func() {
		// A -> B -> C -> A
		parentChain := []int{2, 1} // C的上级链：C -> B -> A
		hasCircular := checkCircularReference(1, 3, parentChain)
		suite.True(hasCircular, "A设置C为上级，但C的上级链包含A，应该被检测为循环引用")
	})

	suite.Run("深层循环引用", func() {
		// A -> B -> C -> D -> E -> A
		parentChain := []int{4, 3, 2, 1} // E的上级链：E -> D -> C -> B -> A
		hasCircular := checkCircularReference(1, 5, parentChain)
		suite.True(hasCircular, "深层循环引用应该被检测出来")
	})

	suite.Run("正常的层级关系", func() {
		// A -> B (B的上级链不包含A)
		parentChain := []int{3, 4} // B的上级链：B -> C -> D
		hasCircular := checkCircularReference(1, 2, parentChain)
		suite.False(hasCircular, "正常的层级关系不应该被检测为循环引用")
	})

	suite.Run("空上级链", func() {
		// A -> B (B是顶级部门)
		parentChain := []int{}
		hasCircular := checkCircularReference(1, 2, parentChain)
		suite.False(hasCircular, "设置顶级部门为上级不应该有循环引用")
	})

	suite.Run("复杂的正常层级", func() {
		// 测试复杂但正常的层级关系
		parentChain := []int{5, 6, 7, 8} // 上级链不包含当前部门
		hasCircular := checkCircularReference(1, 4, parentChain)
		suite.False(hasCircular, "复杂但正常的层级关系不应该有循环引用")
	})
}

// TestGetAllChildrenIDs 测试获取所有子部门ID
func (suite *WlDepartmentTreeTestSuite) TestGetAllChildrenIDs() {
	// 构建测试数据：复杂的部门层级结构
	// 1 (公司)
	//   ├── 2 (技术部)
	//   │   ├── 4 (前端组)
	//   │   │   ├── 7 (React组)
	//   │   │   └── 8 (Vue组)
	//   │   └── 5 (后端组)
	//   │       └── 9 (Go组)
	//   └── 3 (市场部)
	//       └── 6 (销售组)

	departmentHierarchy := map[int][]int{
		1: {2, 3}, // 公司的直接子部门
		2: {4, 5}, // 技术部的直接子部门
		3: {6},    // 市场部的直接子部门
		4: {7, 8}, // 前端组的直接子部门
		5: {9},    // 后端组的直接子部门
		6: {},     // 销售组没有子部门
		7: {},     // React组没有子部门
		8: {},     // Vue组没有子部门
		9: {},     // Go组没有子部门
	}

	suite.Run("获取公司所有子部门", func() {
		allChildren := getAllChildrenIDs(1, departmentHierarchy)
		expected := []int{2, 3, 4, 5, 6, 7, 8, 9}
		suite.ElementsMatch(expected, allChildren, "应该获取所有层级的子部门")
	})

	suite.Run("获取技术部所有子部门", func() {
		allChildren := getAllChildrenIDs(2, departmentHierarchy)
		expected := []int{4, 5, 7, 8, 9}
		suite.ElementsMatch(expected, allChildren, "应该获取技术部下所有层级的子部门")
	})

	suite.Run("获取前端组所有子部门", func() {
		allChildren := getAllChildrenIDs(4, departmentHierarchy)
		expected := []int{7, 8}
		suite.ElementsMatch(expected, allChildren, "应该获取前端组的直接子部门")
	})

	suite.Run("获取叶子部门的子部门", func() {
		allChildren := getAllChildrenIDs(7, departmentHierarchy)
		suite.Empty(allChildren, "叶子部门应该没有子部门")
	})

	suite.Run("获取不存在部门的子部门", func() {
		allChildren := getAllChildrenIDs(999, departmentHierarchy)
		suite.Empty(allChildren, "不存在的部门应该返回空列表")
	})
}

// TestDepartmentNameUniqueness 测试部门名称唯一性
func (suite *WlDepartmentTreeTestSuite) TestDepartmentNameUniqueness() {
	// 模拟现有部门数据
	existingDepartments := []wl_department.WlDepartment{
		{ID: 1, Name: "技术部", ParentID: nil},
		{ID: 2, Name: "市场部", ParentID: nil},
		{ID: 3, Name: "前端组", ParentID: uintPtr(1)},
		{ID: 4, Name: "后端组", ParentID: uintPtr(1)},
		{ID: 5, Name: "前端组", ParentID: uintPtr(2)}, // 不同父部门下可以有同名子部门
	}

	suite.Run("顶级部门名称唯一性检查", func() {
		// 创建新的顶级部门
		isUnique := checkDepartmentNameUnique("新技术部", nil, 0, existingDepartments)
		suite.True(isUnique, "新的顶级部门名称应该唯一")

		// 创建重名的顶级部门
		isUnique = checkDepartmentNameUnique("技术部", nil, 0, existingDepartments)
		suite.False(isUnique, "重名的顶级部门应该不唯一")
	})

	suite.Run("同级子部门名称唯一性检查", func() {
		// 在技术部下创建新的子部门
		isUnique := checkDepartmentNameUnique("测试组", uintPtr(1), 0, existingDepartments)
		suite.True(isUnique, "新的子部门名称应该唯一")

		// 在技术部下创建重名的子部门
		isUnique = checkDepartmentNameUnique("前端组", uintPtr(1), 0, existingDepartments)
		suite.False(isUnique, "同级重名的子部门应该不唯一")
	})

	suite.Run("不同级部门可以同名", func() {
		// 在市场部下创建与技术部下同名的子部门
		isUnique := checkDepartmentNameUnique("后端组", uintPtr(2), 0, existingDepartments)
		suite.True(isUnique, "不同父部门下可以有同名子部门")
	})

	suite.Run("更新时排除自身", func() {
		// 更新部门3（前端组）保持原名称
		isUnique := checkDepartmentNameUnique("前端组", uintPtr(1), 3, existingDepartments)
		suite.True(isUnique, "更新时保持原名称应该唯一")

		// 更新部门3为与同级部门重名
		isUnique = checkDepartmentNameUnique("后端组", uintPtr(1), 3, existingDepartments)
		suite.False(isUnique, "更新时与同级部门重名应该不唯一")
	})

	suite.Run("跨级部门名称检查", func() {
		// 部门5（市场部下的前端组）与部门3（技术部下的前端组）同名但不同级
		isUnique := checkDepartmentNameUnique("前端组", uintPtr(2), 0, existingDepartments)
		suite.False(isUnique, "市场部下已经有前端组了")
	})
}

// TestTreePerformance 测试树操作性能
func (suite *WlDepartmentTreeTestSuite) TestTreePerformance() {
	suite.Run("大量数据构建树性能", func() {
		// 创建大量测试数据
		largeDepartments := generateLargeDepartmentData(1000)

		// 测试构建完整树的性能
		tree := buildDepartmentTree(largeDepartments, nil)
		suite.NotEmpty(tree, "大量数据应该能正常构建树")

		// 验证树的结构正确性
		totalNodes := countTreeNodes(tree)
		suite.Equal(1000, totalNodes, "树中节点总数应该等于输入数据量")
	})

	suite.Run("深层嵌套性能", func() {
		// 创建深层嵌套的部门结构
		deepDepartments := generateDeepDepartmentData(100) // 100层深度

		tree := buildDepartmentTree(deepDepartments, nil)
		suite.NotEmpty(tree, "深层嵌套数据应该能正常构建树")

		// 验证树的深度
		maxDepth := getTreeMaxDepth(tree)
		suite.Equal(100, maxDepth, "树的最大深度应该等于输入深度")
	})
}

// TestTreeValidation 测试树结构验证
func (suite *WlDepartmentTreeTestSuite) TestTreeValidation() {
	suite.Run("验证树结构完整性", func() {
		depts := []wl_department.WlDepartment{
			{ID: 1, Name: "公司", ParentID: nil, Sort: 1},
			{ID: 2, Name: "技术部", ParentID: uintPtr(1), Sort: 1},
			{ID: 3, Name: "前端组", ParentID: uintPtr(2), Sort: 1},
		}

		tree := buildDepartmentTree(depts, nil)

		// 验证树结构的完整性
		isValid := validateTreeStructure(tree, depts)
		suite.True(isValid, "构建的树结构应该是完整和正确的")
	})

	suite.Run("验证父子关系正确性", func() {
		depts := []wl_department.WlDepartment{
			{ID: 1, Name: "技术部", ParentID: nil, Sort: 1},
			{ID: 2, Name: "前端组", ParentID: uintPtr(1), Sort: 1},
			{ID: 3, Name: "后端组", ParentID: uintPtr(1), Sort: 2},
		}

		tree := buildDepartmentTree(depts, nil)

		// 验证父子关系
		isValid := validateParentChildRelationships(tree)
		suite.True(isValid, "父子关系应该正确")
	})
}

// 辅助函数实现

// buildDepartmentTree 构建部门树（模拟实现）
func buildDepartmentTree(departments []wl_department.WlDepartment, parentID *uint) []wl_department.WlDepartment {
	var result []wl_department.WlDepartment

	// 找出指定父部门的直接子部门
	for _, dept := range departments {
		if (parentID == nil && dept.ParentID == nil) || (parentID != nil && dept.ParentID != nil && *dept.ParentID == *parentID) {
			// 递归构建子树
			dept.Children = buildDepartmentTree(departments, &dept.ID)
			result = append(result, dept)
		}
	}

	// 按排序字段排序
	for i := 0; i < len(result)-1; i++ {
		for j := i + 1; j < len(result); j++ {
			if result[i].Sort > result[j].Sort {
				result[i], result[j] = result[j], result[i]
			}
		}
	}

	return result
}

// checkCircularReference 检查循环引用（模拟实现）
func checkCircularReference(deptID, newParentID uint, parentChain []uint) bool {
	// 检查是否选择自身作为上级
	if deptID == newParentID {
		return true
	}

	// 检查新上级的上级链中是否包含当前部门
	for _, parentID := range parentChain {
		if parentID == deptID {
			return true
		}
	}

	return false
}

// getAllChildrenIDs 获取所有子部门ID（模拟实现）
func getAllChildrenIDs(parentID uint, hierarchy map[uint][]uint) []uint {
	var result []uint

	if children, exists := hierarchy[parentID]; exists {
		for _, childID := range children {
			result = append(result, childID)
			// 递归获取子部门的子部门
			grandChildren := getAllChildrenIDs(childID, hierarchy)
			result = append(result, grandChildren...)
		}
	}

	return result
}

// checkDepartmentNameUnique 检查部门名称唯一性（模拟实现）
func checkDepartmentNameUnique(name string, parentID *uint, excludeID uint, existingDepts []wl_department.WlDepartment) bool {
	for _, dept := range existingDepts {
		// 排除自身
		if dept.ID == excludeID {
			continue
		}

		// 检查同级部门名称是否重复
		if dept.Name == name {
			// 检查是否在同一父部门下
			if (parentID == nil && dept.ParentID == nil) ||
				(parentID != nil && dept.ParentID != nil && *parentID == *dept.ParentID) {
				return false
			}
		}
	}
	return true
}

// generateLargeDepartmentData 生成大量测试数据
func generateLargeDepartmentData(count int) []wl_department.WlDepartment {
	departments := make([]wl_department.WlDepartment, count)

	for i := 0; i < count; i++ {
		var parentID *uint
		if i > 0 && i%10 != 0 {
			// 每10个部门中有9个是子部门
			parent := uint((i / 10) * 10)
			parentID = &parent
		}

		departments[i] = wl_department.WlDepartment{
			ID:       uint(i + 1),
			Name:     fmt.Sprintf("部门%d", i+1),
			ParentID: parentID,
			Sort:     i,
		}
	}

	return departments
}

// generateDeepDepartmentData 生成深层嵌套测试数据
func generateDeepDepartmentData(depth int) []wl_department.WlDepartment {
	departments := make([]wl_department.WlDepartment, depth)

	for i := 0; i < depth; i++ {
		var parentID *uint
		if i > 0 {
			parent := uint(i)
			parentID = &parent // 每个部门的父部门是前一个部门
		}

		departments[i] = wl_department.WlDepartment{
			ID:       uint(i + 1),
			Name:     fmt.Sprintf("层级%d", i+1),
			ParentID: parentID,
			Sort:     1,
		}
	}

	return departments
}

// countTreeNodes 计算树中节点总数
func countTreeNodes(tree []wl_department.WlDepartment) int {
	count := len(tree)
	for _, dept := range tree {
		count += countTreeNodes(dept.Children)
	}
	return count
}

// getTreeMaxDepth 获取树的最大深度
func getTreeMaxDepth(tree []wl_department.WlDepartment) int {
	if len(tree) == 0 {
		return 0
	}

	maxDepth := 1
	for _, dept := range tree {
		childDepth := getTreeMaxDepth(dept.Children)
		if childDepth+1 > maxDepth {
			maxDepth = childDepth + 1
		}
	}

	return maxDepth
}

// validateTreeStructure 验证树结构完整性
func validateTreeStructure(tree []wl_department.WlDepartment, originalDepts []wl_department.WlDepartment) bool {
	// 计算原始数据中的部门数量
	originalCount := len(originalDepts)

	// 计算树中的节点数量
	treeCount := countTreeNodes(tree)

	// 验证数量是否一致
	return originalCount == treeCount
}

// validateParentChildRelationships 验证父子关系正确性
func validateParentChildRelationships(tree []wl_department.WlDepartment) bool {
	for _, dept := range tree {
		// 验证子部门的ParentID是否正确指向父部门
		for _, child := range dept.Children {
			if child.ParentID == nil || *child.ParentID != dept.ID {
				return false
			}
		}

		// 递归验证子树
		if !validateParentChildRelationships(dept.Children) {
			return false
		}
	}
	return true
}

// uintPtr 辅助函数：返回uint指针
func uintPtr(i uint) *uint {
	return &i
}
