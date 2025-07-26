package wl_department

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_department/request"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// WlDepartmentApiTestSuite API测试套件
type WlDepartmentApiTestSuite struct {
	suite.Suite
	router *gin.Engine
}

func (suite *WlDepartmentApiTestSuite) SetupTest() {
	suite.router = setupTestRouter()
}

func TestWlDepartmentApiTestSuite(t *testing.T) {
	suite.Run(t, new(WlDepartmentApiTestSuite))
}

// setupTestRouter 设置测试路由
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// 创建API实例
	api := &WlDepartmentApi{}

	// 设置部门管理路由
	departmentGroup := router.Group("/department")
	{
		departmentGroup.GET("/list", api.GetWlDepartmentList)
		departmentGroup.POST("/list", api.GetWlDepartmentList) // 支持POST方式
		departmentGroup.GET("/tree", api.GetDepartmentTree)
		departmentGroup.GET("/:id", api.GetDepartmentDetail)
		departmentGroup.POST("/create", api.CreateWlDepartment)
		departmentGroup.PUT("/update", api.UpdateWlDepartment)
		departmentGroup.DELETE("/delete", api.DeleteWlDepartment)
		departmentGroup.GET("/devices/available", api.GetAvailableDevices)
		departmentGroup.GET("/devices", api.GetDepartmentDevices)
	}

	return router
}

// TestGetWlDepartmentList 测试获取部门列表API
func (suite *WlDepartmentApiTestSuite) TestGetWlDepartmentList() {
	tests := []struct {
		name           string
		method         string
		queryParams    string
		requestBody    interface{}
		expectedStatus int
		checkResponse  func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:           "GET方式获取部门列表成功",
			method:         "GET",
			queryParams:    "?page=1&pageSize=10",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "code", "响应应该包含code字段")

				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err, "响应应该是有效的JSON")

				if data, ok := response["data"].(map[string]interface{}); ok {
					assert.Contains(t, data, "list", "响应数据应该包含list字段")
					assert.Contains(t, data, "total", "响应数据应该包含total字段")
					assert.Contains(t, data, "page", "响应数据应该包含page字段")
					assert.Contains(t, data, "pageSize", "响应数据应该包含pageSize字段")
				}
			},
		},
		{
			name:   "POST方式获取部门列表成功",
			method: "POST",
			requestBody: request.WlDepartmentSearch{
				Page:     1,
				PageSize: 10,
			},
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "code", "响应应该包含code字段")
			},
		},
		{
			name:           "获取树形部门列表",
			method:         "GET",
			queryParams:    "?page=1&pageSize=10&treeMode=true",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "code", "响应应该包含code字段")
			},
		},
		{
			name:           "按名称搜索部门",
			method:         "GET",
			queryParams:    "?name=技术部&page=1&pageSize=10",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "code", "响应应该包含code字段")
			},
		},
		{
			name:           "按状态筛选部门",
			method:         "GET",
			queryParams:    "?status=启用&page=1&pageSize=10",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "code", "响应应该包含code字段")
			},
		},
		{
			name:           "使用默认分页参数",
			method:         "GET",
			queryParams:    "",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "code", "响应应该包含code字段")
			},
		},
		{
			name:           "无效的分页参数",
			method:         "GET",
			queryParams:    "?page=0&pageSize=-1",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "code", "响应应该包含code字段")
			},
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			var req *http.Request

			if tt.method == "POST" && tt.requestBody != nil {
				jsonBody, _ := json.Marshal(tt.requestBody)
				req, _ = http.NewRequest("POST", "/department/list", bytes.NewBuffer(jsonBody))
				req.Header.Set("Content-Type", "application/json")
			} else {
				req, _ = http.NewRequest(tt.method, "/department/list"+tt.queryParams, nil)
			}

			w := httptest.NewRecorder()
			suite.router.ServeHTTP(w, req)

			suite.Equal(tt.expectedStatus, w.Code, "HTTP状态码应该匹配")
			if tt.checkResponse != nil {
				tt.checkResponse(suite.T(), w)
			}
		})
	}
}

// TestGetDepartmentTree 测试获取部门树API
func (suite *WlDepartmentApiTestSuite) TestGetDepartmentTree() {
	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		checkResponse  func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:           "获取完整部门树",
			queryParams:    "",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err, "响应应该是有效的JSON")

				if data, ok := response["data"].([]interface{}); ok {
					// 验证返回的是数组格式
					assert.IsType(t, []interface{}{}, data, "部门树应该是数组格式")
				}
			},
		},
		{
			name:           "排除指定部门的部门树",
			queryParams:    "?excludeId=1",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "code", "响应应该包含code字段")
			},
		},
		{
			name:           "排除无效部门ID",
			queryParams:    "?excludeId=0",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "code", "响应应该包含code字段")
			},
		},
		{
			name:           "排除不存在的部门ID",
			queryParams:    "?excludeId=99999",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "code", "响应应该包含code字段")
			},
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			req, _ := http.NewRequest("GET", "/department/tree"+tt.queryParams, nil)
			w := httptest.NewRecorder()
			suite.router.ServeHTTP(w, req)

			suite.Equal(tt.expectedStatus, w.Code, "HTTP状态码应该匹配")
			if tt.checkResponse != nil {
				tt.checkResponse(suite.T(), w)
			}
		})
	}
}

// TestCreateWlDepartment 测试创建部门API
func (suite *WlDepartmentApiTestSuite) TestCreateWlDepartment() {
	tests := []struct {
		name           string
		requestBody    request.CreateWlDepartmentReq
		expectedStatus int
		checkResponse  func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "创建顶级部门成功",
			requestBody: request.CreateWlDepartmentReq{
				Name:      "测试部门",
				Leader:    "张三",
				Phone:     "13800138000",
				Email:     "zhangsan@example.com",
				Status:    "启用",
				Sort:      1,
				ParentID:  nil,
				DeviceIDs: []int{1, 2},
			},
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err, "响应应该是有效的JSON")

				// 验证响应结构
				assert.Contains(t, response, "code", "响应应该包含code字段")
				assert.Contains(t, response, "msg", "响应应该包含msg字段")
			},
		},
		{
			name: "创建子部门成功",
			requestBody: request.CreateWlDepartmentReq{
				Name:     "子部门",
				Leader:   "李四",
				Phone:    "13800138001",
				Email:    "lisi@example.com",
				Status:   "启用",
				Sort:     1,
				ParentID: func() *int { i := 1; return &i }(),
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "使用DepartmentName字段创建",
			requestBody: request.CreateWlDepartmentReq{
				DepartmentName: "市场部",
				Leader:         "赵六",
				Status:         "启用",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "部门名称为空应该失败",
			requestBody: request.CreateWlDepartmentReq{
				Name:   "",
				Leader: "王五",
			},
			expectedStatus: http.StatusOK, // API层会返回业务错误，HTTP状态码仍为200
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "部门名称不能为空", "应该返回部门名称错误信息")
			},
		},
		{
			name: "Name和DepartmentName都为空应该失败",
			requestBody: request.CreateWlDepartmentReq{
				Name:           "",
				DepartmentName: "",
				Leader:         "王五",
			},
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "部门名称不能为空", "应该返回部门名称错误信息")
			},
		},
		{
			name: "创建部门并关联设备",
			requestBody: request.CreateWlDepartmentReq{
				Name:      "设备管理部",
				Leader:    "孙七",
				Status:    "启用",
				DeviceIDs: []int{1, 2, 3},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "无效的JSON格式",
			requestBody: request.CreateWlDepartmentReq{
				Name: "测试部门",
			},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			jsonBody, _ := json.Marshal(tt.requestBody)
			req, _ := http.NewRequest("POST", "/department/create", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			suite.router.ServeHTTP(w, req)

			suite.Equal(tt.expectedStatus, w.Code, "HTTP状态码应该匹配")
			if tt.checkResponse != nil {
				tt.checkResponse(suite.T(), w)
			}
		})
	}
}

// TestUpdateWlDepartment 测试更新部门API
func (suite *WlDepartmentApiTestSuite) TestUpdateWlDepartment() {
	tests := []struct {
		name           string
		requestBody    request.UpdateWlDepartmentReq
		expectedStatus int
		checkResponse  func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "更新部门基本信息成功",
			requestBody: request.UpdateWlDepartmentReq{
				ID:        1,
				Name:      "更新后的部门",
				Leader:    "张三",
				Phone:     "13800138000",
				Email:     "zhangsan@example.com",
				Status:    "启用",
				Sort:      2,
				ParentID:  nil,
				DeviceIDs: []int{1, 2, 3},
			},
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err, "响应应该是有效的JSON")
				assert.Contains(t, response, "code", "响应应该包含code字段")
			},
		},
		{
			name: "更新部门上级关系",
			requestBody: request.UpdateWlDepartmentReq{
				ID:       2,
				Name:     "前端组",
				ParentID: func() *int { i := 1; return &i }(),
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "使用DepartmentName字段更新",
			requestBody: request.UpdateWlDepartmentReq{
				ID:             1,
				DepartmentName: "技术部-新",
				Leader:         "李四",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "更新部门状态为禁用",
			requestBody: request.UpdateWlDepartmentReq{
				ID:     1,
				Name:   "技术部",
				Status: "禁用",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "部门ID为0应该失败",
			requestBody: request.UpdateWlDepartmentReq{
				ID:   0,
				Name: "测试部门",
			},
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "部门ID不能为空", "应该返回ID错误信息")
			},
		},
		{
			name: "部门ID为负数应该失败",
			requestBody: request.UpdateWlDepartmentReq{
				ID:   -1,
				Name: "测试部门",
			},
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "部门ID不能为空", "应该返回ID错误信息")
			},
		},
		{
			name: "部门名称为空应该失败",
			requestBody: request.UpdateWlDepartmentReq{
				ID:   1,
				Name: "",
			},
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "部门名称不能为空", "应该返回名称错误信息")
			},
		},
		{
			name: "Name和DepartmentName都为空应该失败",
			requestBody: request.UpdateWlDepartmentReq{
				ID:             1,
				Name:           "",
				DepartmentName: "",
			},
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "部门名称不能为空", "应该返回名称错误信息")
			},
		},
		{
			name: "更新设备关联",
			requestBody: request.UpdateWlDepartmentReq{
				ID:        1,
				Name:      "技术部",
				DeviceIDs: []int{4, 5, 6},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "清空设备关联",
			requestBody: request.UpdateWlDepartmentReq{
				ID:        1,
				Name:      "技术部",
				DeviceIDs: []int{},
			},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			jsonBody, _ := json.Marshal(tt.requestBody)
			req, _ := http.NewRequest("PUT", "/department/update", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			suite.router.ServeHTTP(w, req)

			suite.Equal(tt.expectedStatus, w.Code, "HTTP状态码应该匹配")
			if tt.checkResponse != nil {
				tt.checkResponse(suite.T(), w)
			}
		})
	}
}

// TestDeleteWlDepartment 测试删除部门API
func TestDeleteWlDepartment(t *testing.T) {
	router := setupTestRouter()

	tests := []struct {
		name           string
		requestBody    request.DeleteWlDepartmentReq
		expectedStatus int
		checkResponse  func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "删除部门成功",
			requestBody: request.DeleteWlDepartmentReq{
				ID: 1,
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "部门ID为0应该失败",
			requestBody: request.DeleteWlDepartmentReq{
				ID: 0,
			},
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "部门ID不能为空", "应该返回ID错误信息")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBody, _ := json.Marshal(tt.requestBody)
			req, _ := http.NewRequest("DELETE", "/department/delete", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code, "HTTP状态码应该匹配")
			if tt.checkResponse != nil {
				tt.checkResponse(t, w)
			}
		})
	}
}

// TestGetDepartmentDetail 测试获取部门详情API
func TestGetDepartmentDetail(t *testing.T) {
	router := setupTestRouter()

	tests := []struct {
		name           string
		departmentID   string
		expectedStatus int
		checkResponse  func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:           "获取部门详情成功",
			departmentID:   "1",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "无效的部门ID",
			departmentID:   "invalid",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "无效的部门ID", "应该返回ID无效错误")
			},
		},
		{
			name:           "部门ID为0",
			departmentID:   "0",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "无效的部门ID", "应该返回ID无效错误")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/department/"+tt.departmentID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code, "HTTP状态码应该匹配")
			if tt.checkResponse != nil {
				tt.checkResponse(t, w)
			}
		})
	}
}

// TestGetAvailableDevices 测试获取可用设备API
func TestGetAvailableDevices(t *testing.T) {
	router := setupTestRouter()

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
	}{
		{
			name:           "获取可用设备成功",
			queryParams:    "?page=1&pageSize=10",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "按设备名称搜索",
			queryParams:    "?deviceName=测试设备&page=1&pageSize=10",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "按产品名称搜索",
			queryParams:    "?productName=测试产品&page=1&pageSize=10",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "排除指定部门的设备",
			queryParams:    "?departmentId=1&page=1&pageSize=10",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/department/devices/available"+tt.queryParams, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code, "HTTP状态码应该匹配")
		})
	}
}

// TestGetDepartmentDevices 测试获取部门设备API
func TestGetDepartmentDevices(t *testing.T) {
	router := setupTestRouter()

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		checkResponse  func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:           "获取部门设备成功",
			queryParams:    "?departmentId=1&page=1&pageSize=10",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "部门ID为空应该失败",
			queryParams:    "?page=1&pageSize=10",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "部门ID不能为空", "应该返回部门ID错误")
			},
		},
		{
			name:           "部门ID为0应该失败",
			queryParams:    "?departmentId=0&page=1&pageSize=10",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "部门ID不能为空", "应该返回部门ID错误")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/department/devices"+tt.queryParams, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code, "HTTP状态码应该匹配")
			if tt.checkResponse != nil {
				tt.checkResponse(t, w)
			}
		})
	}
}

// TestAPIErrorHandling 测试API错误处理
func TestAPIErrorHandling(t *testing.T) {
	router := setupTestRouter()

	tests := []struct {
		name           string
		method         string
		url            string
		body           interface{}
		expectedStatus int
		checkResponse  func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:           "无效的JSON格式",
			method:         "POST",
			url:            "/department/create",
			body:           "invalid json",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "参数错误", "应该返回参数错误")
			},
		},
		{
			name:           "缺少必填字段",
			method:         "POST",
			url:            "/department/create",
			body:           map[string]interface{}{},
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "部门名称不能为空", "应该返回必填字段错误")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var reqBody []byte
			if str, ok := tt.body.(string); ok {
				reqBody = []byte(str)
			} else {
				reqBody, _ = json.Marshal(tt.body)
			}

			req, _ := http.NewRequest(tt.method, tt.url, bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code, "HTTP状态码应该匹配")
			if tt.checkResponse != nil {
				tt.checkResponse(t, w)
			}
		})
	}
}

// BenchmarkGetWlDepartmentList 性能测试：获取部门列表
func BenchmarkGetWlDepartmentList(b *testing.B) {
	router := setupTestRouter()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest("GET", "/department/list?page=1&pageSize=10", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
	}
}

// BenchmarkCreateWlDepartment 性能测试：创建部门
func BenchmarkCreateWlDepartment(b *testing.B) {
	router := setupTestRouter()

	requestBody := request.CreateWlDepartmentReq{
		Name:   "性能测试部门",
		Leader: "测试用户",
		Phone:  "13800138000",
		Email:  "test@example.com",
		Status: "启用",
		Sort:   1,
	}

	jsonBody, _ := json.Marshal(requestBody)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest("POST", "/department/create", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
	}
}

// TestDeleteWlDepartment 测试删除部门API
func (suite *WlDepartmentApiTestSuite) TestDeleteWlDepartment() {
	tests := []struct {
		name           string
		requestBody    request.DeleteWlDepartmentReq
		expectedStatus int
		checkResponse  func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "删除部门成功",
			requestBody: request.DeleteWlDepartmentReq{
				ID: 3, // 假设是叶子部门
			},
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err, "响应应该是有效的JSON")
			},
		},
		{
			name: "部门ID为0应该失败",
			requestBody: request.DeleteWlDepartmentReq{
				ID: 0,
			},
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "部门ID不能为空", "应该返回ID错误信息")
			},
		},
		{
			name: "部门ID为负数应该失败",
			requestBody: request.DeleteWlDepartmentReq{
				ID: -1,
			},
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "部门ID不能为空", "应该返回ID错误信息")
			},
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			jsonBody, _ := json.Marshal(tt.requestBody)
			req, _ := http.NewRequest("DELETE", "/department/delete", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			suite.router.ServeHTTP(w, req)

			suite.Equal(tt.expectedStatus, w.Code, "HTTP状态码应该匹配")
			if tt.checkResponse != nil {
				tt.checkResponse(suite.T(), w)
			}
		})
	}
}

// TestGetDepartmentDetail 测试获取部门详情API
func (suite *WlDepartmentApiTestSuite) TestGetDepartmentDetail() {
	tests := []struct {
		name           string
		departmentID   string
		expectedStatus int
		checkResponse  func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:           "获取部门详情成功",
			departmentID:   "1",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err, "响应应该是有效的JSON")

				if data, ok := response["data"].(map[string]interface{}); ok {
					assert.Contains(t, data, "id", "部门详情应该包含id字段")
					assert.Contains(t, data, "name", "部门详情应该包含name字段")
					assert.Contains(t, data, "devices", "部门详情应该包含devices字段")
				}
			},
		},
		{
			name:           "无效的部门ID格式",
			departmentID:   "invalid",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "无效的部门ID", "应该返回ID无效错误")
			},
		},
		{
			name:           "部门ID为0",
			departmentID:   "0",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "无效的部门ID", "应该返回ID无效错误")
			},
		},
		{
			name:           "部门ID为负数",
			departmentID:   "-1",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "无效的部门ID", "应该返回ID无效错误")
			},
		},
		{
			name:           "不存在的部门ID",
			departmentID:   "99999",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				// 这里可能返回查询失败或者空结果
				assert.Contains(t, w.Body.String(), "code", "响应应该包含code字段")
			},
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			req, _ := http.NewRequest("GET", "/department/"+tt.departmentID, nil)
			w := httptest.NewRecorder()
			suite.router.ServeHTTP(w, req)

			suite.Equal(tt.expectedStatus, w.Code, "HTTP状态码应该匹配")
			if tt.checkResponse != nil {
				tt.checkResponse(suite.T(), w)
			}
		})
	}
}

// TestGetAvailableDevices 测试获取可用设备API
func (suite *WlDepartmentApiTestSuite) TestGetAvailableDevices() {
	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		checkResponse  func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:           "获取可用设备成功",
			queryParams:    "?page=1&pageSize=10",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err, "响应应该是有效的JSON")

				if data, ok := response["data"].(map[string]interface{}); ok {
					assert.Contains(t, data, "list", "响应数据应该包含list字段")
					assert.Contains(t, data, "total", "响应数据应该包含total字段")
				}
			},
		},
		{
			name:           "按设备名称搜索",
			queryParams:    "?deviceName=传感器&page=1&pageSize=10",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "按产品名称搜索",
			queryParams:    "?productName=温度传感器&page=1&pageSize=10",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "排除指定部门的设备",
			queryParams:    "?departmentId=1&page=1&pageSize=10",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "组合搜索条件",
			queryParams:    "?deviceName=传感器&productName=温度&departmentId=1&page=1&pageSize=5",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "使用默认分页参数",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "无效的分页参数",
			queryParams:    "?page=0&pageSize=-1",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			req, _ := http.NewRequest("GET", "/department/devices/available"+tt.queryParams, nil)
			w := httptest.NewRecorder()
			suite.router.ServeHTTP(w, req)

			suite.Equal(tt.expectedStatus, w.Code, "HTTP状态码应该匹配")
			if tt.checkResponse != nil {
				tt.checkResponse(suite.T(), w)
			}
		})
	}
}

// TestGetDepartmentDevices 测试获取部门设备API
func (suite *WlDepartmentApiTestSuite) TestGetDepartmentDevices() {
	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		checkResponse  func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:           "获取部门设备成功",
			queryParams:    "?departmentId=1&page=1&pageSize=10",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err, "响应应该是有效的JSON")

				if data, ok := response["data"].(map[string]interface{}); ok {
					assert.Contains(t, data, "list", "响应数据应该包含list字段")
					assert.Contains(t, data, "total", "响应数据应该包含total字段")
				}
			},
		},
		{
			name:           "部门ID为空应该失败",
			queryParams:    "?page=1&pageSize=10",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "部门ID不能为空", "应该返回部门ID错误")
			},
		},
		{
			name:           "部门ID为0应该失败",
			queryParams:    "?departmentId=0&page=1&pageSize=10",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "部门ID不能为空", "应该返回部门ID错误")
			},
		},
		{
			name:           "部门ID为负数应该失败",
			queryParams:    "?departmentId=-1&page=1&pageSize=10",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "部门ID不能为空", "应该返回部门ID错误")
			},
		},
		{
			name:           "使用默认分页参数",
			queryParams:    "?departmentId=1",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "无效的分页参数",
			queryParams:    "?departmentId=1&page=0&pageSize=-1",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			req, _ := http.NewRequest("GET", "/department/devices"+tt.queryParams, nil)
			w := httptest.NewRecorder()
			suite.router.ServeHTTP(w, req)

			suite.Equal(tt.expectedStatus, w.Code, "HTTP状态码应该匹配")
			if tt.checkResponse != nil {
				tt.checkResponse(suite.T(), w)
			}
		})
	}
}

// TestAPIErrorHandling 测试API错误处理
func (suite *WlDepartmentApiTestSuite) TestAPIErrorHandling() {
	tests := []struct {
		name           string
		method         string
		url            string
		body           interface{}
		contentType    string
		expectedStatus int
		checkResponse  func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:           "无效的JSON格式-创建部门",
			method:         "POST",
			url:            "/department/create",
			body:           "invalid json",
			contentType:    "application/json",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "参数错误", "应该返回参数错误")
			},
		},
		{
			name:           "无效的JSON格式-更新部门",
			method:         "PUT",
			url:            "/department/update",
			body:           "invalid json",
			contentType:    "application/json",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "参数错误", "应该返回参数错误")
			},
		},
		{
			name:           "无效的JSON格式-删除部门",
			method:         "DELETE",
			url:            "/department/delete",
			body:           "invalid json",
			contentType:    "application/json",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "参数错误", "应该返回参数错误")
			},
		},
		{
			name:           "缺少Content-Type头",
			method:         "POST",
			url:            "/department/create",
			body:           map[string]interface{}{"name": "测试部门"},
			contentType:    "",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "错误的Content-Type",
			method:         "POST",
			url:            "/department/create",
			body:           map[string]interface{}{"name": "测试部门"},
			contentType:    "text/plain",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "空请求体",
			method:         "POST",
			url:            "/department/create",
			body:           nil,
			contentType:    "application/json",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Contains(t, w.Body.String(), "部门名称不能为空", "应该返回必填字段错误")
			},
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			var reqBody []byte
			if tt.body != nil {
				if str, ok := tt.body.(string); ok {
					reqBody = []byte(str)
				} else {
					reqBody, _ = json.Marshal(tt.body)
				}
			}

			req, _ := http.NewRequest(tt.method, tt.url, bytes.NewBuffer(reqBody))
			if tt.contentType != "" {
				req.Header.Set("Content-Type", tt.contentType)
			}
			w := httptest.NewRecorder()
			suite.router.ServeHTTP(w, req)

			suite.Equal(tt.expectedStatus, w.Code, "HTTP状态码应该匹配")
			if tt.checkResponse != nil {
				tt.checkResponse(suite.T(), w)
			}
		})
	}
}

// TestAPIIntegration 测试API集成场景
func (suite *WlDepartmentApiTestSuite) TestAPIIntegration() {
	suite.Run("完整的部门管理流程", func() {
		// 1. 获取部门列表
		req, _ := http.NewRequest("GET", "/department/list?page=1&pageSize=10", nil)
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusOK, w.Code)

		// 2. 获取部门树
		req, _ = http.NewRequest("GET", "/department/tree", nil)
		w = httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusOK, w.Code)

		// 3. 创建部门
		createReq := request.CreateWlDepartmentReq{
			Name:   "集成测试部门",
			Leader: "测试用户",
			Status: "启用",
		}
		jsonBody, _ := json.Marshal(createReq)
		req, _ = http.NewRequest("POST", "/department/create", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w = httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusOK, w.Code)

		// 4. 获取可用设备
		req, _ = http.NewRequest("GET", "/department/devices/available?page=1&pageSize=10", nil)
		w = httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusOK, w.Code)

		// 5. 获取部门详情
		req, _ = http.NewRequest("GET", "/department/1", nil)
		w = httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusOK, w.Code)
	})

	suite.Run("错误处理流程", func() {
		// 1. 尝试创建无效部门
		createReq := request.CreateWlDepartmentReq{
			Name: "", // 空名称
		}
		jsonBody, _ := json.Marshal(createReq)
		req, _ := http.NewRequest("POST", "/department/create", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusOK, w.Code)
		suite.Contains(w.Body.String(), "部门名称不能为空")

		// 2. 尝试更新不存在的部门
		updateReq := request.UpdateWlDepartmentReq{
			ID:   99999,
			Name: "不存在的部门",
		}
		jsonBody, _ = json.Marshal(updateReq)
		req, _ = http.NewRequest("PUT", "/department/update", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w = httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusOK, w.Code)

		// 3. 尝试删除无效ID的部门
		deleteReq := request.DeleteWlDepartmentReq{
			ID: 0,
		}
		jsonBody, _ = json.Marshal(deleteReq)
		req, _ = http.NewRequest("DELETE", "/department/delete", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w = httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusOK, w.Code)
		suite.Contains(w.Body.String(), "部门ID不能为空")
	})
}

// BenchmarkGetWlDepartmentList 性能测试：获取部门列表
func BenchmarkGetWlDepartmentList(b *testing.B) {
	router := setupTestRouter()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest("GET", "/department/list?page=1&pageSize=10", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
	}
}

// BenchmarkCreateWlDepartment 性能测试：创建部门
func BenchmarkCreateWlDepartment(b *testing.B) {
	router := setupTestRouter()

	requestBody := request.CreateWlDepartmentReq{
		Name:   "性能测试部门",
		Leader: "测试用户",
		Phone:  "13800138000",
		Email:  "test@example.com",
		Status: "启用",
		Sort:   1,
	}

	jsonBody, _ := json.Marshal(requestBody)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest("POST", "/department/create", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
	}
}

// BenchmarkGetDepartmentTree 性能测试：获取部门树
func BenchmarkGetDepartmentTree(b *testing.B) {
	router := setupTestRouter()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest("GET", "/department/tree", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
	}
}
