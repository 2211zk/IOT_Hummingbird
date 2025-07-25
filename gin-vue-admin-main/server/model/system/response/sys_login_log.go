package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/system"

// LoginLogResponse 登录日志响应结构体
type LoginLogResponse struct {
	LoginLog system.SysLoginLog `json:"loginLog"`
}

// LoginLogListResponse 登录日志列表响应结构体
type LoginLogListResponse struct {
	List     []system.SysLoginLog `json:"list"`
	Total    int64                `json:"total"`
	Page     int                  `json:"page"`
	PageSize int                  `json:"pageSize"`
}

// CleanLogsResponse 清理日志响应结构体
type CleanLogsResponse struct {
	DeletedCount int64 `json:"deletedCount"` // 删除的记录数
}
