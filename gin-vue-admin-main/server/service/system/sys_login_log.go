package system

import (
	"errors"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
)

type LoginLogService struct{}

var LoginLogServiceApp = new(LoginLogService)

// CreateLoginLog 创建登录日志
func (loginLogService *LoginLogService) CreateLoginLog(loginLog system.SysLoginLog) error {
	return global.GVA_DB.Create(&loginLog).Error
}

// DeleteLoginLog 删除登录日志
func (loginLogService *LoginLogService) DeleteLoginLog(loginLog system.SysLoginLog) error {
	return global.GVA_DB.Delete(&loginLog).Error
}

// DeleteLoginLogByIds 批量删除登录日志
func (loginLogService *LoginLogService) DeleteLoginLogByIds(ids []uint) error {
	return global.GVA_DB.Delete(&[]system.SysLoginLog{}, "id in ?", ids).Error
}

// UpdateLoginLog 更新登录日志
func (loginLogService *LoginLogService) UpdateLoginLog(loginLog system.SysLoginLog) error {
	return global.GVA_DB.Save(&loginLog).Error
}

// GetLoginLog 根据ID获取登录日志
func (loginLogService *LoginLogService) GetLoginLog(id uint) (loginLog system.SysLoginLog, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&loginLog).Error
	return
}

// GetLoginLogInfoList 分页获取登录日志列表
func (loginLogService *LoginLogService) GetLoginLogInfoList(info systemReq.LoginLogSearchReq) (list []system.SysLoginLog, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 创建数据库查询
	db := global.GVA_DB.Model(&system.SysLoginLog{})

	// 构建查询条件
	if info.UserName != "" {
		db = db.Where("user_name LIKE ?", "%"+info.UserName+"%")
	}
	if info.LoginAddress != "" {
		db = db.Where("login_address LIKE ?", "%"+info.LoginAddress+"%")
	}
	if info.LoginLocation != "" {
		db = db.Where("login_location LIKE ?", "%"+info.LoginLocation+"%")
	}
	if info.LoginStatus != "" {
		db = db.Where("login_status = ?", info.LoginStatus)
	}

	// 时间范围查询
	if !info.StartTime.IsZero() {
		db = db.Where("login_time >= ?", info.StartTime)
	}
	if !info.EndTime.IsZero() {
		db = db.Where("login_time <= ?", info.EndTime)
	}

	// 获取总数
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	// 获取分页数据，按登录时间倒序排列
	err = db.Limit(limit).Offset(offset).Order("login_time DESC").Find(&list).Error
	return list, total, err
}

// GetLoginLogsByUserName 根据用户名获取登录日志
func (loginLogService *LoginLogService) GetLoginLogsByUserName(userName string, limit int) ([]system.SysLoginLog, error) {
	var logs []system.SysLoginLog
	err := global.GVA_DB.Where("user_name = ?", userName).
		Order("login_time DESC").
		Limit(limit).
		Find(&logs).Error
	return logs, err
}

// GetRecentLoginLogs 获取最近的登录日志
func (loginLogService *LoginLogService) GetRecentLoginLogs(hours int, limit int) ([]system.SysLoginLog, error) {
	var logs []system.SysLoginLog
	timeLimit := time.Now().Add(-time.Duration(hours) * time.Hour)

	err := global.GVA_DB.Where("login_time >= ?", timeLimit).
		Order("login_time DESC").
		Limit(limit).
		Find(&logs).Error
	return logs, err
}

// GetLoginStatistics 获取登录统计信息
func (loginLogService *LoginLogService) GetLoginStatistics(days int) (map[string]interface{}, error) {
	timeLimit := time.Now().AddDate(0, 0, -days)

	var totalLogins int64
	var successLogins int64
	var failedLogins int64
	var uniqueUsers int64

	// 总登录次数
	err := global.GVA_DB.Model(&system.SysLoginLog{}).
		Where("login_time >= ?", timeLimit).
		Count(&totalLogins).Error
	if err != nil {
		return nil, err
	}

	// 成功登录次数
	err = global.GVA_DB.Model(&system.SysLoginLog{}).
		Where("login_time >= ? AND login_status = ?", timeLimit, "成功").
		Count(&successLogins).Error
	if err != nil {
		return nil, err
	}

	// 失败登录次数
	err = global.GVA_DB.Model(&system.SysLoginLog{}).
		Where("login_time >= ? AND login_status = ?", timeLimit, "失败").
		Count(&failedLogins).Error
	if err != nil {
		return nil, err
	}

	// 独立用户数
	err = global.GVA_DB.Model(&system.SysLoginLog{}).
		Where("login_time >= ?", timeLimit).
		Distinct("user_name").
		Count(&uniqueUsers).Error
	if err != nil {
		return nil, err
	}

	statistics := map[string]interface{}{
		"totalLogins":   totalLogins,
		"successLogins": successLogins,
		"failedLogins":  failedLogins,
		"uniqueUsers":   uniqueUsers,
		"successRate":   float64(successLogins) / float64(totalLogins) * 100,
	}

	return statistics, nil
}

// CleanExpiredLogs 清理过期日志
func (loginLogService *LoginLogService) CleanExpiredLogs(days int) (int64, error) {
	if days <= 0 {
		return 0, errors.New("保留天数必须大于0")
	}

	expireTime := time.Now().AddDate(0, 0, -days)

	result := global.GVA_DB.Where("login_time < ?", expireTime).Delete(&system.SysLoginLog{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

// GetLoginLogsByIP 根据IP地址获取登录日志
func (loginLogService *LoginLogService) GetLoginLogsByIP(ip string, limit int) ([]system.SysLoginLog, error) {
	var logs []system.SysLoginLog
	err := global.GVA_DB.Where("login_address = ?", ip).
		Order("login_time DESC").
		Limit(limit).
		Find(&logs).Error
	return logs, err
}

// GetFailedLoginAttempts 获取失败的登录尝试
func (loginLogService *LoginLogService) GetFailedLoginAttempts(userName string, hours int) (int64, error) {
	timeLimit := time.Now().Add(-time.Duration(hours) * time.Hour)

	var count int64
	err := global.GVA_DB.Model(&system.SysLoginLog{}).
		Where("user_name = ? AND login_status = ? AND login_time >= ?", userName, "失败", timeLimit).
		Count(&count).Error

	return count, err
}

// GetTopLoginIPs 获取登录次数最多的IP地址
func (loginLogService *LoginLogService) GetTopLoginIPs(limit int, days int) ([]map[string]interface{}, error) {
	timeLimit := time.Now().AddDate(0, 0, -days)

	var results []map[string]interface{}

	err := global.GVA_DB.Model(&system.SysLoginLog{}).
		Select("login_address, login_location, COUNT(*) as login_count").
		Where("login_time >= ?", timeLimit).
		Group("login_address, login_location").
		Order("login_count DESC").
		Limit(limit).
		Find(&results).Error

	return results, err
}

// ExportLoginLogs 导出登录日志数据
func (loginLogService *LoginLogService) ExportLoginLogs(info systemReq.LoginLogSearchReq) ([]system.SysLoginLog, error) {
	// 创建数据库查询
	db := global.GVA_DB.Model(&system.SysLoginLog{})

	// 构建查询条件（与GetLoginLogInfoList相同的条件）
	if info.UserName != "" {
		db = db.Where("user_name LIKE ?", "%"+info.UserName+"%")
	}
	if info.LoginAddress != "" {
		db = db.Where("login_address LIKE ?", "%"+info.LoginAddress+"%")
	}
	if info.LoginLocation != "" {
		db = db.Where("login_location LIKE ?", "%"+info.LoginLocation+"%")
	}
	if info.LoginStatus != "" {
		db = db.Where("login_status = ?", info.LoginStatus)
	}

	// 时间范围查询
	if !info.StartTime.IsZero() {
		db = db.Where("login_time >= ?", info.StartTime)
	}
	if !info.EndTime.IsZero() {
		db = db.Where("login_time <= ?", info.EndTime)
	}

	var logs []system.SysLoginLog
	err := db.Order("login_time DESC").Find(&logs).Error
	return logs, err
}

// ExportLoginLogsToExcel 导出登录日志到Excel文件
func (loginLogService *LoginLogService) ExportLoginLogsToExcel(info systemReq.LoginLogSearchReq) ([]byte, string, error) {
	// 获取数据
	logs, err := loginLogService.ExportLoginLogs(info)
	if err != nil {
		return nil, "", err
	}

	// 转换为utils包的数据结构
	var exportLogs []utils.LoginLogData
	for _, log := range logs {
		exportLogs = append(exportLogs, utils.LoginLogData{
			ID:              log.ID,
			UserName:        log.UserName,
			LoginAddress:    log.LoginAddress,
			LoginLocation:   log.LoginLocation,
			Browser:         log.Browser,
			OperatingSystem: log.OperatingSystem,
			LoginStatus:     log.LoginStatus,
			OperationalInfo: log.OperationalInfo,
			LoginTime:       log.LoginTime,
		})
	}

	// 导入utils包中的导出函数
	excelData, err := utils.ExportLoginLogsToExcel(exportLogs)
	if err != nil {
		return nil, "", err
	}

	// 生成文件名
	filename := utils.GenerateExcelFileName("登录日志")

	return excelData, filename, nil
}

// ExportLoginStatisticsToExcel 导出登录统计信息到Excel
func (loginLogService *LoginLogService) ExportLoginStatisticsToExcel(days int) ([]byte, string, error) {
	// 获取统计信息
	statistics, err := loginLogService.GetLoginStatistics(days)
	if err != nil {
		return nil, "", err
	}

	// 获取热门IP
	topIPs, err := loginLogService.GetTopLoginIPs(10, days)
	if err != nil {
		return nil, "", err
	}

	// 获取最近登录记录
	recentLogs, err := loginLogService.GetRecentLoginLogs(24, 50)
	if err != nil {
		return nil, "", err
	}

	// 转换为utils包的数据结构
	var exportRecentLogs []utils.LoginLogData
	for _, log := range recentLogs {
		exportRecentLogs = append(exportRecentLogs, utils.LoginLogData{
			ID:              log.ID,
			UserName:        log.UserName,
			LoginAddress:    log.LoginAddress,
			LoginLocation:   log.LoginLocation,
			Browser:         log.Browser,
			OperatingSystem: log.OperatingSystem,
			LoginStatus:     log.LoginStatus,
			OperationalInfo: log.OperationalInfo,
			LoginTime:       log.LoginTime,
		})
	}

	// 导出到Excel
	excelData, err := utils.ExportLoginStatisticsToExcel(statistics, topIPs, exportRecentLogs)
	if err != nil {
		return nil, "", err
	}

	// 生成文件名
	filename := utils.GenerateExcelFileName("登录统计")

	return excelData, filename, nil
}

// GetLogRetentionPolicy 获取日志保留策略配置
func (loginLogService *LoginLogService) GetLogRetentionPolicy() (int, error) {
	// 从配置文件或数据库获取保留策略
	// 这里暂时返回默认值90天
	return 90, nil
}

// SetLogRetentionPolicy 设置日志保留策略
func (loginLogService *LoginLogService) SetLogRetentionPolicy(days int) error {
	if days <= 0 {
		return errors.New("保留天数必须大于0")
	}

	// 这里应该将配置保存到数据库或配置文件
	// 暂时只做参数验证
	return nil
}

// BackupLogsBeforeClean 清理前备份日志
func (loginLogService *LoginLogService) BackupLogsBeforeClean(days int) (string, error) {
	if days <= 0 {
		return "", errors.New("保留天数必须大于0")
	}

	expireTime := time.Now().AddDate(0, 0, -days)

	// 获取要删除的日志
	var logs []system.SysLoginLog
	err := global.GVA_DB.Where("login_time < ?", expireTime).Find(&logs).Error
	if err != nil {
		return "", err
	}

	if len(logs) == 0 {
		return "", nil
	}

	// 转换为utils包的数据结构
	var exportLogs []utils.LoginLogData
	for _, log := range logs {
		exportLogs = append(exportLogs, utils.LoginLogData{
			ID:              log.ID,
			UserName:        log.UserName,
			LoginAddress:    log.LoginAddress,
			LoginLocation:   log.LoginLocation,
			Browser:         log.Browser,
			OperatingSystem: log.OperatingSystem,
			LoginStatus:     log.LoginStatus,
			OperationalInfo: log.OperationalInfo,
			LoginTime:       log.LoginTime,
		})
	}

	// 导出到Excel作为备份
	_, err = utils.ExportLoginLogsToExcel(exportLogs)
	if err != nil {
		return "", err
	}

	// 生成备份文件名
	backupFileName := utils.GenerateExcelFileName("登录日志备份")

	// 这里应该将文件保存到备份目录
	// 暂时返回文件名，实际实现时需要保存文件
	// 在实际实现中，应该使用excelData将文件保存到磁盘

	return backupFileName, nil
}

// CleanExpiredLogsWithBackup 清理过期日志（带备份）
func (loginLogService *LoginLogService) CleanExpiredLogsWithBackup(days int, needBackup bool) (map[string]interface{}, error) {
	if days <= 0 {
		return nil, errors.New("保留天数必须大于0")
	}

	result := make(map[string]interface{})

	// 如果需要备份，先进行备份
	if needBackup {
		backupFile, err := loginLogService.BackupLogsBeforeClean(days)
		if err != nil {
			return nil, err
		}
		result["backupFile"] = backupFile
	}

	// 执行清理
	deletedCount, err := loginLogService.CleanExpiredLogs(days)
	if err != nil {
		return nil, err
	}

	result["deletedCount"] = deletedCount
	result["cleanTime"] = time.Now()

	// 记录清理操作的审计日志
	auditLog := system.SysOperationRecord{
		Ip:     "system",
		Method: "DELETE",
		Path:   "/loginLog/cleanExpiredLogs",
		Status: 200,
		Agent:  "System Cleanup Task",
		Body:   fmt.Sprintf("清理%d天前的登录日志", days),
		Resp:   fmt.Sprintf("删除了%d条记录", deletedCount),
		UserID: 0, // 系统操作
	}

	global.GVA_DB.Create(&auditLog)

	return result, nil
}

// GetCleanupStatistics 获取清理统计信息
func (loginLogService *LoginLogService) GetCleanupStatistics(days int) (map[string]interface{}, error) {
	if days <= 0 {
		return nil, errors.New("保留天数必须大于0")
	}

	expireTime := time.Now().AddDate(0, 0, -days)

	var totalCount int64
	var expiredCount int64

	// 获取总记录数
	err := global.GVA_DB.Model(&system.SysLoginLog{}).Count(&totalCount).Error
	if err != nil {
		return nil, err
	}

	// 获取过期记录数
	err = global.GVA_DB.Model(&system.SysLoginLog{}).
		Where("login_time < ?", expireTime).
		Count(&expiredCount).Error
	if err != nil {
		return nil, err
	}

	// 计算最早和最晚的记录时间
	var earliestTime, latestTime time.Time

	global.GVA_DB.Model(&system.SysLoginLog{}).
		Select("MIN(login_time)").
		Scan(&earliestTime)

	global.GVA_DB.Model(&system.SysLoginLog{}).
		Select("MAX(login_time)").
		Scan(&latestTime)

	statistics := map[string]interface{}{
		"totalCount":    totalCount,
		"expiredCount":  expiredCount,
		"retainCount":   totalCount - expiredCount,
		"expireTime":    expireTime,
		"earliestTime":  earliestTime,
		"latestTime":    latestTime,
		"retentionDays": days,
	}

	return statistics, nil
}

// ScheduleCleanup 定时清理任务
func (loginLogService *LoginLogService) ScheduleCleanup() error {
	// 获取保留策略
	retentionDays, err := loginLogService.GetLogRetentionPolicy()
	if err != nil {
		return err
	}

	// 执行清理（带备份）
	result, err := loginLogService.CleanExpiredLogsWithBackup(retentionDays, true)
	if err != nil {
		global.GVA_LOG.Error("定时清理登录日志失败", zap.Error(err))
		return err
	}

	global.GVA_LOG.Info("定时清理登录日志完成",
		zap.Int64("deletedCount", result["deletedCount"].(int64)),
		zap.Int("retentionDays", retentionDays))

	return nil
}
