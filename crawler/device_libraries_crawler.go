package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 数据库配置
const (
	dbHost     = "14.103.143.229"
	dbPort     = "3306"
	dbUser     = "root"
	dbPassword = "82AF916059F13E331E633AC0B8A8191B"
	dbName     = "wl_playform"
	tableName  = "driver_cards"
)

// API配置
const (
	loginURL = "https://demo.winc-link.com/api/v1/auth/login"
	apiURL   = "https://demo.winc-link.com/api/v1/device-libraries?isAll=true&is_internal=true&classify_id=9001"
	username = "admin"
	password = "123456"
)

// 数据结构定义
type DeviceLibrariesResponse struct {
	Success   bool                    `json:"success"`
	ErrorCode int                     `json:"errorCode"`
	ErrorMsg  string                  `json:"errorMsg"`
	Result    DeviceLibrariesResult   `json:"result"`
}

type DeviceLibrariesResult struct {
	List     []DeviceLibrary `json:"list"`
	Total    int             `json:"total"`
	Page     int             `json:"page"`
	PageSize int             `json:"pageSize"`
}

type DeviceLibrary struct {
	ID               string           `json:"id"`
	Name             string           `json:"name"`
	Description      string           `json:"description"`
	Protocol         string           `json:"protocol"`
	Version          string           `json:"version"`
	ContainerName    string           `json:"container_name"`
	DockerConfigID   string           `json:"docker_config_id"`
	DockerRepoName   string           `json:"docker_repo_name"`
	OperateStatus    string           `json:"operate_status"`
	IsInternal       bool             `json:"is_internal"`
	Manual           string           `json:"manual"`
	Icon             string           `json:"icon"`
	ClassifyID       string           `json:"classify_id"`
	Created          int64            `json:"created"`
	Language         string           `json:"language"`
	IsFree           bool             `json:"is_free"`
	SupportVersions  []SupportVersion `json:"support_versions"`
}

type SupportVersion struct {
	Version    string `json:"version"`
	IsDefault  bool   `json:"is_default"`
	ConfigFile string `json:"config_file"`
}

func main() {
	// 连接数据库
	db, err := connectDB()
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	defer db.Close()

	// 清空现有数据
	fmt.Println("清空现有驱动卡片数据...")
	_, err = db.Exec(fmt.Sprintf("DELETE FROM %s", tableName))
	if err != nil {
		log.Fatalf("清空数据失败: %v", err)
	}

	// 重置自增ID
	_, err = db.Exec(fmt.Sprintf("ALTER TABLE %s AUTO_INCREMENT = 1", tableName))
	if err != nil {
		log.Printf("重置自增ID失败: %v", err)
	}

	// 创建HTTP客户端并登录
	client, token, err := createAuthenticatedClient()
	if err != nil {
		log.Fatalf("登录失败: %v", err)
	}
	fmt.Printf("登录成功，开始爬取设备库数据...\n")

	// 获取设备库数据
	deviceLibraries, err := fetchDeviceLibraries(client, token)
	if err != nil {
		log.Fatalf("获取设备库数据失败: %v", err)
	}

	fmt.Printf("获取到 %d 条设备库数据\n", len(deviceLibraries))

	// 存储到数据库
	err = saveToDatabase(db, deviceLibraries)
	if err != nil {
		log.Fatalf("保存数据失败: %v", err)
	}

	fmt.Println("数据导入完成！")
	
	// 显示统计信息
	showStatistics(db)
}

// 连接数据库
func connectDB() (*sql.DB, error) {
	var dsn string
	if dbPassword == "" {
		dsn = fmt.Sprintf("%s@tcp(%s:%s)/%s", dbUser, dbHost, dbPort, dbName)
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	}
	
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("打开数据库连接失败: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("数据库连接测试失败: %v", err)
	}

	fmt.Println("成功连接到数据库")
	return db, nil
}

// 创建认证的HTTP客户端
func createAuthenticatedClient() (*http.Client, string, error) {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar:     jar,
		Timeout: 30 * time.Second,
	}

	// 登录获取token
	loginData := map[string]string{
		"username": username,
		"password": password,
	}

	jsonData, _ := json.Marshal(loginData)
	req, _ := http.NewRequest("POST", loginURL, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	
	var loginResponse struct {
		Success bool `json:"success"`
		Result  struct {
			Token string `json:"token"`
		} `json:"result"`
	}

	json.Unmarshal(body, &loginResponse)
	
	if !loginResponse.Success {
		return nil, "", fmt.Errorf("登录失败")
	}

	// 设置cookie
	u, _ := url.Parse("https://demo.winc-link.com")
	cookie := &http.Cookie{
		Name:  "token",
		Value: loginResponse.Result.Token,
	}
	client.Jar.SetCookies(u, []*http.Cookie{cookie})

	return client, loginResponse.Result.Token, nil
}

// 获取设备库数据
func fetchDeviceLibraries(client *http.Client, token string) ([]DeviceLibrary, error) {
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	req.Header.Set("Origin", "https://demo.winc-link.com")
	req.Header.Set("Referer", "https://demo.winc-link.com/")
	req.Header.Set("X-Token", token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("API响应状态: %d\n", resp.StatusCode)
	fmt.Printf("API响应内容: %s\n", string(body))

	var response DeviceLibrariesResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("解析JSON失败: %v", err)
	}

	if !response.Success {
		return nil, fmt.Errorf("API返回错误: %s", response.ErrorMsg)
	}

	return response.Result.List, nil
}

// 保存数据到数据库
func saveToDatabase(db *sql.DB, deviceLibraries []DeviceLibrary) error {
	// 开始事务
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// 准备插入语句
	insertSQL := fmt.Sprintf(`INSERT INTO %s (name, img, description, tags) VALUES (?, ?, ?, ?)`, tableName)
	stmt, err := tx.Prepare(insertSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	successCount := 0

	for _, library := range deviceLibraries {
		// 处理字段长度限制
		name := truncateString(library.Name, 128)
		img := truncateString(library.Icon, 255)
		description := truncateString(library.Description, 255)
		
		// 构建tags字符串，包含协议、版本、状态等信息
		tags := []string{
			library.Protocol,
			"v" + library.Version,
			library.OperateStatus,
		}
		if library.IsFree {
			tags = append(tags, "免费")
		}
		if library.IsInternal {
			tags = append(tags, "内置")
		}
		
		tagsStr := truncateString(strings.Join(tags, ","), 255)

		_, err = stmt.Exec(name, img, description, tagsStr)
		if err != nil {
			log.Printf("插入数据失败: %v, 数据: %s", err, library.Name)
			continue
		}
		successCount++
	}

	// 提交事务
	err = tx.Commit()
	if err != nil {
		return err
	}

	fmt.Printf("成功插入 %d 条记录\n", successCount)
	return nil
}

// 截断字符串以适应数据库字段长度
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen]
}

// 显示统计信息
func showStatistics(db *sql.DB) {
	var totalCount int
	err := db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", tableName)).Scan(&totalCount)
	if err != nil {
		log.Printf("查询总记录数失败: %v", err)
		return
	}

	fmt.Printf("\n=== 数据库统计信息 ===\n")
	fmt.Printf("总记录数: %d\n", totalCount)

	// 显示前10条记录作为示例
	fmt.Println("\n前10条记录示例:")
	rows, err := db.Query(fmt.Sprintf("SELECT id, name, img, description, tags FROM %s LIMIT 10", tableName))
	if err != nil {
		log.Printf("查询示例数据失败: %v", err)
		return
	}
	defer rows.Close()

	fmt.Printf("%-5s %-20s %-30s %-30s %-20s\n", "ID", "名称", "图标", "描述", "标签")
	fmt.Println(strings.Repeat("-", 110))
	
	for rows.Next() {
		var id int
		var name, img, description, tags string
		err = rows.Scan(&id, &name, &img, &description, &tags)
		if err != nil {
			continue
		}
		
		// 截断过长的字符串用于显示
		if len(name) > 18 {
			name = name[:15] + "..."
		}
		if len(img) > 28 {
			img = img[:25] + "..."
		}
		if len(description) > 28 {
			description = description[:25] + "..."
		}
		if len(tags) > 18 {
			tags = tags[:15] + "..."
		}
		
		fmt.Printf("%-5d %-20s %-30s %-30s %-20s\n", id, name, img, description, tags)
	}
}