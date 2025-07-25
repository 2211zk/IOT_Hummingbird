package system

import (
	"testing"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/stretchr/testify/assert"
)

func TestLoginLogService_CreateLoginLog(t *testing.T) {
	service := LoginLogService{}

	loginLog := system.SysLoginLog{
		UserName:        "testuser",
		LoginAddress:    "192.168.1.1",
		LoginLocation:   "北京 北京",
		Browser:         "Chrome",
		OperatingSystem: "Windows",
		LoginStatus:     "成功",
		OperationalInfo: "登录成功",
		LoginTime:       time.Now(),
	}

	err := service.CreateLoginLog(loginLog)
	assert.NoError(t, err)
}

func TestLoginLogService_GetLoginLogInfoList(t *testing.T) {
	service := LoginLogService{}

	req := systemReq.LoginLogSearchReq{
		PageInfo: systemReq.PageInfo{
			Page:     1,
			PageSize: 10,
		},
		UserName: "testuser",
	}

	list, total, err := service.GetLoginLogInfoList(req)
	assert.NoError(t, err)
	assert.NotNil(t, list)
	assert.GreaterOrEqual(t, total, int64(0))
}

func TestLoginLogService_GetLoginStatistics(t *testing.T) {
	service := LoginLogService{}

	stats, err := service.GetLoginStatistics(7)
	assert.NoError(t, err)
	assert.NotNil(t, stats)

	// 检查统计数据的键是否存在
	assert.Contains(t, stats, "totalLogins")
	assert.Contains(t, stats, "successLogins")
	assert.Contains(t, stats, "failedLogins")
	assert.Contains(t, stats, "uniqueUsers")
	assert.Contains(t, stats, "successRate")
}

func TestLoginLogService_CleanExpiredLogs(t *testing.T) {
	service := LoginLogService{}

	// 测试无效的天数
	count, err := service.CleanExpiredLogs(0)
	assert.Error(t, err)
	assert.Equal(t, int64(0), count)

	// 测试有效的天数
	count, err = service.CleanExpiredLogs(90)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, count, int64(0))
}

func TestLoginLogService_GetFailedLoginAttempts(t *testing.T) {
	service := LoginLogService{}

	count, err := service.GetFailedLoginAttempts("testuser", 24)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, count, int64(0))
}

func TestLoginLogService_GetTopLoginIPs(t *testing.T) {
	service := LoginLogService{}

	results, err := service.GetTopLoginIPs(10, 7)
	assert.NoError(t, err)
	assert.NotNil(t, results)
}
func TestLoginLogService_CleanExpiredLogsWithBackup(t *testing.T) {
	service := LoginLogService{}

	// 测试带备份的清理
	result, err := service.CleanExpiredLogsWithBackup(90, true)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Contains(t, result, "deletedCount")
	assert.Contains(t, result, "cleanTime")
}

func TestLoginLogService_GetCleanupStatistics(t *testing.T) {
	service := LoginLogService{}

	stats, err := service.GetCleanupStatistics(90)
	assert.NoError(t, err)
	assert.NotNil(t, stats)

	// 检查统计数据的键是否存在
	assert.Contains(t, stats, "totalCount")
	assert.Contains(t, stats, "expiredCount")
	assert.Contains(t, stats, "retainCount")
	assert.Contains(t, stats, "expireTime")
	assert.Contains(t, stats, "retentionDays")
}

func TestLoginLogService_GetLogRetentionPolicy(t *testing.T) {
	service := LoginLogService{}

	days, err := service.GetLogRetentionPolicy()
	assert.NoError(t, err)
	assert.Greater(t, days, 0)
}

func TestLoginLogService_SetLogRetentionPolicy(t *testing.T) {
	service := LoginLogService{}

	// 测试有效的天数
	err := service.SetLogRetentionPolicy(30)
	assert.NoError(t, err)

	// 测试无效的天数
	err = service.SetLogRetentionPolicy(0)
	assert.Error(t, err)

	err = service.SetLogRetentionPolicy(-1)
	assert.Error(t, err)
}
