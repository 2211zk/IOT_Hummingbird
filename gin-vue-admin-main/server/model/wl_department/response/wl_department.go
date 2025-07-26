package response

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_department"
)

// DepartmentTreeNode 部门树节点响应
type DepartmentTreeNode struct {
	ID       uint                  `json:"id"`
	Name     string                `json:"name"`
	ParentID *uint                 `json:"parentId"`
	Children []*DepartmentTreeNode `json:"children"`
}

// DepartmentListResponse 部门列表响应
type DepartmentListResponse struct {
	ID             uint                      `json:"id"`
	ParentID       *uint                     `json:"parentId"`
	Name           string                    `json:"name"`
	DepartmentName string                    `json:"departmentName"` // 兼容字段
	Leader         string                    `json:"leader"`
	Phone          string                    `json:"phone"`
	Email          string                    `json:"email"`
	Status         string                    `json:"status"`
	Sort           int                       `json:"sort"`
	CreatedAt      time.Time                 `json:"createdAt"`
	UpdatedAt      time.Time                 `json:"updatedAt"`
	Children       []*DepartmentListResponse `json:"children,omitempty"`
	DeviceCount    int                       `json:"deviceCount"` // 关联设备数量
}

// DepartmentDetailResponse 部门详情响应
type DepartmentDetailResponse struct {
	ID             uint                `json:"id"`
	ParentID       *uint               `json:"parentId"`
	Name           string              `json:"name"`
	DepartmentName string              `json:"departmentName"`
	Leader         string              `json:"leader"`
	Phone          string              `json:"phone"`
	Email          string              `json:"email"`
	Status         string              `json:"status"`
	Sort           int                 `json:"sort"`
	CreatedAt      time.Time           `json:"createdAt"`
	UpdatedAt      time.Time           `json:"updatedAt"`
	Devices        []DeviceResponse    `json:"devices"`
	Parent         *DepartmentTreeNode `json:"parent,omitempty"`
}

// DeviceResponse 设备响应
type DeviceResponse struct {
	ID          uint   `json:"id"`
	DeviceName  string `json:"deviceName"`
	ProductName string `json:"productName"`
	Status      string `json:"status"`
}

// AvailableDevicesResponse 可用设备响应
type AvailableDevicesResponse struct {
	List  []DeviceResponse `json:"list"`
	Total int64            `json:"total"`
}

// ConvertToTreeNode 将部门模型转换为树节点
func ConvertToTreeNode(dept *wl_department.WlDepartment) *DepartmentTreeNode {
	node := &DepartmentTreeNode{
		ID:       dept.ID,
		Name:     dept.Name,
		ParentID: dept.ParentID,
		Children: make([]*DepartmentTreeNode, 0),
	}

	for _, child := range dept.Children {
		childNode := ConvertToTreeNode(&child)
		node.Children = append(node.Children, childNode)
	}

	return node
}

// ConvertToListResponse 将部门模型转换为列表响应
func ConvertToListResponse(dept *wl_department.WlDepartment) *DepartmentListResponse {
	resp := &DepartmentListResponse{
		ID:             dept.ID,
		ParentID:       dept.ParentID,
		Name:           dept.Name,
		DepartmentName: dept.DepartmentName,
		Leader:         dept.Leader,
		Phone:          dept.Phone,
		Email:          dept.Email,
		Status:         dept.Status,
		Sort:           dept.Sort,
		CreatedAt:      dept.CreatedAt,
		UpdatedAt:      dept.UpdatedAt,
		DeviceCount:    len(dept.Devices),
		Children:       make([]*DepartmentListResponse, 0),
	}

	for _, child := range dept.Children {
		childResp := ConvertToListResponse(&child)
		resp.Children = append(resp.Children, childResp)
	}

	return resp
}

// ConvertToDetailResponse 将部门模型转换为详情响应
func ConvertToDetailResponse(dept *wl_department.WlDepartment) *DepartmentDetailResponse {
	resp := &DepartmentDetailResponse{
		ID:             dept.ID,
		ParentID:       dept.ParentID,
		Name:           dept.Name,
		DepartmentName: dept.DepartmentName,
		Leader:         dept.Leader,
		Phone:          dept.Phone,
		Email:          dept.Email,
		Status:         dept.Status,
		Sort:           dept.Sort,
		CreatedAt:      dept.CreatedAt,
		UpdatedAt:      dept.UpdatedAt,
		Devices:        make([]DeviceResponse, 0),
	}

	// 转换设备信息
	for _, device := range dept.Devices {
		deviceResp := DeviceResponse{
			ID:          device.ID,
			DeviceName:  device.DeviceName,
			ProductName: device.ProductName,
			Status:      device.Status,
		}
		resp.Devices = append(resp.Devices, deviceResp)
	}

	// 转换父部门信息
	if dept.Parent != nil {
		resp.Parent = &DepartmentTreeNode{
			ID:       dept.Parent.ID,
			Name:     dept.Parent.Name,
			ParentID: dept.Parent.ParentID,
		}
	}

	return resp
}
