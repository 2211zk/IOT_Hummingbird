package utils

import (
	"strings"
)

// UserAgentInfo 用户代理信息结构体
type UserAgentInfo struct {
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browserVersion"`
	OS             string `json:"os"`
	OSVersion      string `json:"osVersion"`
	Device         string `json:"device"`
	IsMobile       bool   `json:"isMobile"`
	IsTablet       bool   `json:"isTablet"`
	IsDesktop      bool   `json:"isDesktop"`
}

// ParseUserAgent 解析用户代理字符串
func ParseUserAgent(userAgent string) UserAgentInfo {
	info := UserAgentInfo{
		Browser:   "Unknown",
		OS:        "Unknown",
		Device:    "Desktop",
		IsDesktop: true,
	}

	if userAgent == "" {
		return info
	}

	userAgent = strings.ToLower(userAgent)

	// 检测移动设备
	mobilePatterns := []string{
		"mobile", "android", "iphone", "ipad", "ipod", "blackberry", "windows phone",
		"opera mini", "opera mobi", "palm", "symbian", "nokia", "samsung",
	}

	for _, pattern := range mobilePatterns {
		if strings.Contains(userAgent, pattern) {
			info.IsMobile = true
			info.IsDesktop = false
			info.Device = "Mobile"
			break
		}
	}

	// 检测平板设备
	if strings.Contains(userAgent, "ipad") ||
		(strings.Contains(userAgent, "android") && !strings.Contains(userAgent, "mobile")) {
		info.IsTablet = true
		info.IsMobile = false
		info.IsDesktop = false
		info.Device = "Tablet"
	}

	// 解析浏览器
	info.Browser = parseBrowser(userAgent)

	// 解析操作系统
	info.OS = parseOS(userAgent)

	return info
}

// parseBrowser 解析浏览器信息
func parseBrowser(userAgent string) string {
	browsers := map[string][]string{
		"Chrome":            {"chrome/", "chromium/"},
		"Firefox":           {"firefox/"},
		"Safari":            {"safari/"},
		"Edge":              {"edge/", "edg/"},
		"Opera":             {"opera/", "opr/"},
		"Internet Explorer": {"msie", "trident/"},
	}

	for browser, patterns := range browsers {
		for _, pattern := range patterns {
			if strings.Contains(userAgent, pattern) {
				return browser
			}
		}
	}

	return "Unknown"
}

// parseOS 解析操作系统信息
func parseOS(userAgent string) string {
	osPatterns := map[string][]string{
		"Windows": {"windows nt 10.0", "windows nt 6.3", "windows nt 6.2", "windows nt 6.1", "windows nt 6.0", "windows nt 5.1", "windows"},
		"macOS":   {"mac os x", "macos", "darwin"},
		"Linux":   {"linux", "ubuntu", "debian", "fedora", "centos"},
		"Android": {"android"},
		"iOS":     {"iphone os", "ios", "ipad"},
		"Unix":    {"unix", "bsd"},
	}

	for os, patterns := range osPatterns {
		for _, pattern := range patterns {
			if strings.Contains(userAgent, pattern) {
				return os
			}
		}
	}

	return "Unknown"
}

// GetSimpleBrowserName 获取简化的浏览器名称
func GetSimpleBrowserName(userAgent string) string {
	info := ParseUserAgent(userAgent)
	return info.Browser
}

// GetSimpleOSName 获取简化的操作系统名称
func GetSimpleOSName(userAgent string) string {
	info := ParseUserAgent(userAgent)
	return info.OS
}

// IsValidUserAgent 验证用户代理字符串是否有效
func IsValidUserAgent(userAgent string) bool {
	if userAgent == "" {
		return false
	}

	// 基本的用户代理字符串应该包含一些常见的标识符
	commonPatterns := []string{
		"mozilla", "webkit", "chrome", "safari", "firefox", "edge", "opera", "msie",
	}

	userAgentLower := strings.ToLower(userAgent)
	for _, pattern := range commonPatterns {
		if strings.Contains(userAgentLower, pattern) {
			return true
		}
	}

	return false
}
