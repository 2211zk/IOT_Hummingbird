package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

const (
	test_loginURL = "https://demo.winc-link.com/auth/login"
	test_testURL  = "https://demo.winc-link.com/driver/image"
	test_username = "admin"
	test_password = "123456"
)

func testOfficialProtocolAPI() {
	// 创建cookie jar
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar:     jar,
		Timeout: 30 * time.Second,
	}

	// 登录获取token
	token, err := test_login(client)
	if err != nil {
		fmt.Printf("登录失败: %v\n", err)
		return
	}
	fmt.Printf("登录成功，token: %s\n", token[:50]+"...")

	// 获取官方协议数据
	fmt.Printf("\n获取官方协议数据: %s\n", test_testURL)

	// 构建请求参数，获取官方协议分类的数据
	req, _ := http.NewRequest("GET", test_testURL, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	req.Header.Set("Origin", "https://demo.winc-link.com")
	req.Header.Set("Referer", "https://demo.winc-link.com/")
	req.Header.Set("X-Token", token)

	// 添加查询参数来获取官方协议数据
	q := req.URL.Query()
	q.Add("category", "official") // 官方协议分类
	q.Add("type", "protocol")     // 协议类型
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Printf("响应状态: %d\n", resp.StatusCode)
	fmt.Printf("响应内容: %s\n", string(body))

	// 解析并分析官方协议数据
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err == nil {
		fmt.Println("\n=== 官方协议数据结构分析 ===")
		test_printJSONStructure(result, 0)

		// 提取官方协议列表
		test_extractOfficialProtocols(result)
	} else {
		fmt.Printf("JSON解析失败: %v\n", err)
	}
}

func test_login(client *http.Client) (string, error) {
	loginData := map[string]string{
		"username": test_username,
		"password": test_password,
	}

	jsonData, _ := json.Marshal(loginData)
	req, _ := http.NewRequest("POST", test_loginURL, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	fmt.Printf("正在登录: %s\n", test_loginURL)
	fmt.Printf("登录数据: %s\n", string(jsonData))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("HTTP请求失败: %v\n", err)
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("登录响应状态: %d\n", resp.StatusCode)
	fmt.Printf("登录响应内容: %s\n", string(body))

	var loginResponse struct {
		Success bool `json:"success"`
		Result  struct {
			Token string `json:"token"`
		} `json:"result"`
		Message string `json:"message"`
	}

	if err := json.Unmarshal(body, &loginResponse); err != nil {
		fmt.Printf("JSON解析失败: %v\n", err)
		return "", fmt.Errorf("JSON解析失败: %v", err)
	}

	if !loginResponse.Success {
		return "", fmt.Errorf("登录失败: %s", loginResponse.Message)
	}

	// 设置cookie
	u, _ := url.Parse("https://demo.winc-link.com")
	cookie := &http.Cookie{
		Name:  "token",
		Value: loginResponse.Result.Token,
	}
	client.Jar.SetCookies(u, []*http.Cookie{cookie})

	return loginResponse.Result.Token, nil
}

func test_extractOfficialProtocols(data map[string]interface{}) {
	fmt.Println("\n=== 官方协议数据提取 ===")

	// 尝试从不同可能的字段中提取协议列表
	possibleFields := []string{"data", "result", "list", "items", "protocols", "drivers"}

	for _, field := range possibleFields {
		if value, exists := data[field]; exists {
			fmt.Printf("找到字段 '%s': %T\n", field, value)

			switch v := value.(type) {
			case []interface{}:
				fmt.Printf("发现协议列表，共 %d 个协议:\n", len(v))
				for i, protocol := range v {
					if protocolMap, ok := protocol.(map[string]interface{}); ok {
						fmt.Printf("协议 %d:\n", i+1)
						test_printProtocolInfo(protocolMap)
					}
				}
				return
			case map[string]interface{}:
				fmt.Printf("字段 '%s' 是对象，继续查找...\n", field)
				test_extractOfficialProtocols(v)
				return
			}
		}
	}

	fmt.Println("未找到协议列表数据")
}

func test_printProtocolInfo(protocol map[string]interface{}) {
	// 打印协议的基本信息
	fields := []string{"name", "title", "description", "version", "status", "type", "category"}

	for _, field := range fields {
		if value, exists := protocol[field]; exists {
			fmt.Printf("  %s: %v\n", field, value)
		}
	}

	// 检查是否有下载状态
	if status, exists := protocol["downloadStatus"]; exists {
		fmt.Printf("  下载状态: %v\n", status)
	}

	// 检查是否有付费信息
	if payment, exists := protocol["payment"]; exists {
		fmt.Printf("  付费状态: %v\n", payment)
	}

	fmt.Println()
}

func test_printJSONStructure(data interface{}, indent int) {
	prefix := ""
	for i := 0; i < indent; i++ {
		prefix += "  "
	}

	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			fmt.Printf("%s%s: ", prefix, key)
			switch value.(type) {
			case map[string]interface{}:
				fmt.Println("(object)")
				test_printJSONStructure(value, indent+1)
			case []interface{}:
				fmt.Println("(array)")
				if arr := value.([]interface{}); len(arr) > 0 {
					fmt.Printf("%s  [0]: ", prefix)
					test_printJSONStructure(arr[0], indent+1)
				}
			default:
				fmt.Printf("(%T) %v\n", value, value)
			}
		}
	case []interface{}:
		fmt.Printf("%s(array with %d items)\n", prefix, len(v))
		if len(v) > 0 {
			test_printJSONStructure(v[0], indent)
		}
	default:
		fmt.Printf("%s(%T) %v\n", prefix, v, v)
	}
}
