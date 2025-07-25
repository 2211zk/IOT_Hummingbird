package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExportLoginLogsToExcel(t *testing.T) {
	// 创建测试数据
	logs := []LoginLogData{
		{
			ID:              1,
			UserName:        "testuser1",
			LoginAddress:    "192.168.1.1",
			LoginLocation:   "北京 北京",
			Browser:         "Chrome",
			OperatingSystem: "Windows",
			LoginStatus:     "成功",
			OperationalInfo: "登录成功",
			LoginTime:       time.Now(),
		},
		{
			ID:              2,
			UserName:        "testuser2",
			LoginAddress:    "192.168.1.2",
			LoginLocation:   "上海 上海",
			Browser:         "Firefox",
			OperatingSystem: "macOS",
			LoginStatus:     "失败",
			OperationalInfo: "密码错误",
			LoginTime:       time.Now().Add(-time.Hour),
		},
	}

	// 测试导出功能
	data, err := ExportLoginLogsToExcel(logs)
	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Greater(t, len(data), 0)
}

func TestExportLoginStatisticsToExcel(t *testing.T) {
	// 创建测试统计数据
	statistics := map[string]interface{}{
		"totalLogins":   100,
		"successLogins": 85,
		"failedLogins":  15,
		"uniqueUsers":   20,
		"successRate":   85.0,
	}

	// 创建测试IP数据
	topIPs := []map[string]interface{}{
		{
			"login_address":  "192.168.1.1",
			"login_location": "北京 北京",
			"login_count":    25,
		},
		{
			"login_address":  "192.168.1.2",
			"login_location": "上海 上海",
			"login_count":    20,
		},
	}

	// 创建测试最近登录数据
	recentLogs := []LoginLogData{
		{
			ID:            1,
			UserName:      "testuser1",
			LoginAddress:  "192.168.1.1",
			LoginLocation: "北京 北京",
			Browser:       "Chrome",
			LoginStatus:   "成功",
			LoginTime:     time.Now(),
		},
	}

	// 测试导出功能
	data, err := ExportLoginStatisticsToExcel(statistics, topIPs, recentLogs)
	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Greater(t, len(data), 0)
}

func TestGenerateExcelFileName(t *testing.T) {
	filename := GenerateExcelFileName("test")
	assert.Contains(t, filename, "test_")
	assert.Contains(t, filename, ".xlsx")
}

func TestSetExcelResponseHeaders(t *testing.T) {
	headers := SetExcelResponseHeaders("test.xlsx")
	assert.Equal(t, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", headers["Content-Type"])
	assert.Contains(t, headers["Content-Disposition"], "test.xlsx")
	assert.Equal(t, "no-cache", headers["Cache-Control"])
}
