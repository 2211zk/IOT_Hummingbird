package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

// IPLocationInfo IP地理位置信息
type IPLocationInfo struct {
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
	ISP      string `json:"isp"`
}

// GetIPLocation 获取IP地址的地理位置信息（百度地图API实现）
func GetIPLocation(ip string) string {
	// 检查是否为内网IP
	if isPrivateIP(ip) {
		return "内网IP"
	}
	if ip == "127.0.0.1" || ip == "::1" || ip == "localhost" {
		return "本地"
	}
	ak := "fbZPlNxNhl49M4bg6zHwLTP0DpJym7eS" // 百度地图server ak
	url := fmt.Sprintf("https://api.map.baidu.com/location/ip?ip=%s&ak=%s&coor=bd09ll", ip, ak)
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return "未知位置"
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var result struct {
		Content struct {
			AddressDetail struct {
				Province string `json:"province"`
				City     string `json:"city"`
			} `json:"address_detail"`
		} `json:"content"`
		Status int `json:"status"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return "未知位置"
	}
	if result.Status != 0 {
		return "未知位置"
	}
	return result.Content.AddressDetail.Province + "-" + result.Content.AddressDetail.City
}

// isPrivateIP 检查是否为内网IP
func isPrivateIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}

	// 检查IPv4私有地址范围
	if parsedIP.To4() != nil {
		return parsedIP.IsPrivate()
	}

	// 检查IPv6私有地址
	return parsedIP.IsPrivate()
}

// getLocationFromAPI 从API获取位置信息
func getLocationFromAPI(ip string) string {
	// 使用免费的IP地理位置API
	apis := []string{
		fmt.Sprintf("http://ip-api.com/json/%s?lang=zh-CN", ip),
		fmt.Sprintf("https://ipapi.co/%s/json/", ip),
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	for _, apiURL := range apis {
		if location := tryGetLocationFromURL(client, apiURL, ip); location != "" {
			return location
		}
	}

	return ""
}

// tryGetLocationFromURL 尝试从指定URL获取位置信息
func tryGetLocationFromURL(client *http.Client, url, ip string) string {
	resp, err := client.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	// 解析不同API的响应格式
	if strings.Contains(url, "ip-api.com") {
		return parseIPAPIResponse(body)
	} else if strings.Contains(url, "ipapi.co") {
		return parseIPAPICoResponse(body)
	}

	return ""
}

// parseIPAPIResponse 解析ip-api.com的响应
func parseIPAPIResponse(body []byte) string {
	var result struct {
		Status     string `json:"status"`
		Country    string `json:"country"`
		RegionName string `json:"regionName"`
		City       string `json:"city"`
		ISP        string `json:"isp"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return ""
	}

	if result.Status != "success" {
		return ""
	}

	location := result.Country
	if result.RegionName != "" && result.RegionName != result.Country {
		location += " " + result.RegionName
	}
	if result.City != "" && result.City != result.RegionName {
		location += " " + result.City
	}

	return location
}

// parseIPAPICoResponse 解析ipapi.co的响应
func parseIPAPICoResponse(body []byte) string {
	var result struct {
		Country string `json:"country_name"`
		Region  string `json:"region"`
		City    string `json:"city"`
		Org     string `json:"org"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return ""
	}

	location := result.Country
	if result.Region != "" {
		location += " " + result.Region
	}
	if result.City != "" {
		location += " " + result.City
	}

	return location
}

// GetClientIP 获取客户端真实IP地址
func GetClientIP(r *http.Request) string {
	// 检查X-Forwarded-For头
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	if xForwardedFor != "" {
		// X-Forwarded-For可能包含多个IP，取第一个
		ips := strings.Split(xForwardedFor, ",")
		if len(ips) > 0 {
			ip := strings.TrimSpace(ips[0])
			if net.ParseIP(ip) != nil {
				return ip
			}
		}
	}

	// 检查X-Real-IP头
	xRealIP := r.Header.Get("X-Real-IP")
	if xRealIP != "" {
		if net.ParseIP(xRealIP) != nil {
			return xRealIP
		}
	}

	// 检查X-Forwarded-For头的其他格式
	forwarded := r.Header.Get("Forwarded")
	if forwarded != "" {
		// 解析Forwarded头格式: for=192.0.2.60;proto=http;by=203.0.113.43
		parts := strings.Split(forwarded, ";")
		for _, part := range parts {
			if strings.HasPrefix(strings.TrimSpace(part), "for=") {
				ip := strings.TrimPrefix(strings.TrimSpace(part), "for=")
				ip = strings.Trim(ip, "\"")
				if net.ParseIP(ip) != nil {
					return ip
				}
			}
		}
	}

	// 使用RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}

	return ip
}
