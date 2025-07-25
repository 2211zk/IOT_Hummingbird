package utils

import (
	"fmt"
	"time"

	"github.com/xuri/excelize/v2"
)

// LoginLogData 登录日志数据结构（避免循环导入）
type LoginLogData struct {
	ID              uint      `json:"id"`
	UserName        string    `json:"userName"`
	LoginAddress    string    `json:"loginAddress"`
	LoginLocation   string    `json:"loginLocation"`
	Browser         string    `json:"browser"`
	OperatingSystem string    `json:"operatingSystem"`
	LoginStatus     string    `json:"loginStatus"`
	OperationalInfo string    `json:"operationalInfo"`
	LoginTime       time.Time `json:"loginTime"`
}

// ExportLoginLogsToExcel 导出登录日志到Excel
func ExportLoginLogsToExcel(logs []LoginLogData) ([]byte, error) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// 创建工作表
	sheetName := "登录日志"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return nil, err
	}

	// 设置表头
	headers := []string{
		"ID", "用户名", "登录IP", "登录地点", "浏览器",
		"操作系统", "登录状态", "操作信息", "登录时间",
	}

	// 写入表头
	for i, header := range headers {
		cell := fmt.Sprintf("%s1", string(rune('A'+i)))
		f.SetCellValue(sheetName, cell, header)
	}

	// 设置表头样式
	headerStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
			Size: 12,
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#E0E0E0"},
			Pattern: 1,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
	})
	if err != nil {
		return nil, err
	}

	// 应用表头样式
	f.SetRowStyle(sheetName, 1, 1, headerStyle)

	// 创建数据样式
	dataStyle, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
		Alignment: &excelize.Alignment{
			Vertical: "center",
		},
	})
	if err != nil {
		return nil, err
	}

	// 写入数据
	for i, log := range logs {
		row := i + 2 // 从第2行开始（第1行是表头）

		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), log.ID)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), log.UserName)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), log.LoginAddress)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), log.LoginLocation)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), log.Browser)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), log.OperatingSystem)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), log.LoginStatus)
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), log.OperationalInfo)
		f.SetCellValue(sheetName, fmt.Sprintf("I%d", row), log.LoginTime.Format("2006-01-02 15:04:05"))

		// 应用数据样式
		f.SetRowStyle(sheetName, row, row, dataStyle)
	}

	// 设置列宽
	columnWidths := map[string]float64{
		"A": 8,  // ID
		"B": 15, // 用户名
		"C": 18, // 登录IP
		"D": 20, // 登录地点
		"E": 15, // 浏览器
		"F": 15, // 操作系统
		"G": 10, // 登录状态
		"H": 20, // 操作信息
		"I": 20, // 登录时间
	}

	for col, width := range columnWidths {
		f.SetColWidth(sheetName, col, col, width)
	}

	// 设置活动工作表
	f.SetActiveSheet(index)

	// 删除默认的Sheet1
	f.DeleteSheet("Sheet1")

	// 保存到字节数组
	buffer, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// ExportLoginStatisticsToExcel 导出登录统计信息到Excel
func ExportLoginStatisticsToExcel(statistics map[string]interface{}, topIPs []map[string]interface{}, recentLogs []LoginLogData) ([]byte, error) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// 创建统计信息工作表
	statsSheet := "登录统计"
	index, err := f.NewSheet(statsSheet)
	if err != nil {
		return nil, err
	}

	// 写入统计信息
	f.SetCellValue(statsSheet, "A1", "统计项目")
	f.SetCellValue(statsSheet, "B1", "数值")

	f.SetCellValue(statsSheet, "A2", "总登录次数")
	f.SetCellValue(statsSheet, "B2", statistics["totalLogins"])

	f.SetCellValue(statsSheet, "A3", "成功登录次数")
	f.SetCellValue(statsSheet, "B3", statistics["successLogins"])

	f.SetCellValue(statsSheet, "A4", "失败登录次数")
	f.SetCellValue(statsSheet, "B4", statistics["failedLogins"])

	f.SetCellValue(statsSheet, "A5", "独立用户数")
	f.SetCellValue(statsSheet, "B5", statistics["uniqueUsers"])

	f.SetCellValue(statsSheet, "A6", "成功率(%)")
	if successRate, ok := statistics["successRate"].(float64); ok {
		f.SetCellValue(statsSheet, "B6", fmt.Sprintf("%.2f", successRate))
	}

	// 创建热门IP工作表
	if len(topIPs) > 0 {
		ipSheet := "热门登录IP"
		f.NewSheet(ipSheet)

		f.SetCellValue(ipSheet, "A1", "IP地址")
		f.SetCellValue(ipSheet, "B1", "地理位置")
		f.SetCellValue(ipSheet, "C1", "登录次数")

		for i, ip := range topIPs {
			row := i + 2
			f.SetCellValue(ipSheet, fmt.Sprintf("A%d", row), ip["login_address"])
			f.SetCellValue(ipSheet, fmt.Sprintf("B%d", row), ip["login_location"])
			f.SetCellValue(ipSheet, fmt.Sprintf("C%d", row), ip["login_count"])
		}
	}

	// 创建最近登录工作表
	if len(recentLogs) > 0 {
		recentSheet := "最近登录"
		f.NewSheet(recentSheet)

		headers := []string{"用户名", "登录IP", "登录地点", "浏览器", "登录状态", "登录时间"}
		for i, header := range headers {
			cell := fmt.Sprintf("%s1", string(rune('A'+i)))
			f.SetCellValue(recentSheet, cell, header)
		}

		for i, log := range recentLogs {
			row := i + 2
			f.SetCellValue(recentSheet, fmt.Sprintf("A%d", row), log.UserName)
			f.SetCellValue(recentSheet, fmt.Sprintf("B%d", row), log.LoginAddress)
			f.SetCellValue(recentSheet, fmt.Sprintf("C%d", row), log.LoginLocation)
			f.SetCellValue(recentSheet, fmt.Sprintf("D%d", row), log.Browser)
			f.SetCellValue(recentSheet, fmt.Sprintf("E%d", row), log.LoginStatus)
			f.SetCellValue(recentSheet, fmt.Sprintf("F%d", row), log.LoginTime.Format("2006-01-02 15:04:05"))
		}
	}

	// 设置活动工作表
	f.SetActiveSheet(index)
	f.DeleteSheet("Sheet1")

	// 保存到字节数组
	buffer, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// GenerateExcelFileName 生成Excel文件名
func GenerateExcelFileName(prefix string) string {
	timestamp := time.Now().Format("20060102_150405")
	return fmt.Sprintf("%s_%s.xlsx", prefix, timestamp)
}

// SetExcelResponseHeaders 设置Excel下载的响应头
func SetExcelResponseHeaders(filename string) map[string]string {
	return map[string]string{
		"Content-Type":        "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		"Content-Disposition": fmt.Sprintf("attachment; filename=\"%s\"", filename),
		"Cache-Control":       "no-cache",
	}
}
