package middleware

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LoginLogMiddleware 登录日志记录中间件
func LoginLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 只处理登录相关的请求
		if c.Request.URL.Path != "/base/login" {
			c.Next()
			return
		}

		// 获取请求信息
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()

		// 解析用户代理信息
		userAgentInfo := utils.ParseUserAgent(userAgent)

		// 获取地理位置信息
		location := utils.GetIPLocation(clientIP)

		// 处理请求
		c.Next()

		// 请求处理完成后记录日志
		go func() {
			recordLoginLog(c, clientIP, userAgent, userAgentInfo, location)
		}()
	}
}

// recordLoginLog 记录登录日志
func recordLoginLog(c *gin.Context, clientIP, userAgent string, userAgentInfo utils.UserAgentInfo, location string) {
	// 从上下文中获取用户信息（如果登录成功）
	username := ""
	loginStatus := "失败"
	operationalInfo := "登录失败"

	// 检查响应状态码来判断登录是否成功
	if c.Writer.Status() == 200 {
		// 尝试从上下文获取用户信息
		if userInfo, exists := c.Get("user"); exists {
			if user, ok := userInfo.(*system.SysUser); ok {
				username = user.Username
				loginStatus = "成功"
				operationalInfo = "登录成功"
			}
		} else {
			// 如果没有用户信息但状态码是200，可能是其他原因
			loginStatus = "失败"
			operationalInfo = "登录验证失败"
		}
	} else {
		// 根据不同的状态码设置不同的操作信息
		switch c.Writer.Status() {
		case 400:
			operationalInfo = "请求参数错误"
		case 401:
			operationalInfo = "用户名或密码错误"
		case 403:
			operationalInfo = "账户被禁用"
		case 429:
			operationalInfo = "登录尝试过于频繁"
		default:
			operationalInfo = "登录失败"
		}
	}

	// 如果没有获取到用户名，尝试从请求体中解析
	if username == "" {
		if loginReq, exists := c.Get("loginRequest"); exists {
			if req, ok := loginReq.(map[string]interface{}); ok {
				if user, ok := req["username"].(string); ok {
					username = user
				}
			}
		}
	}

	// 创建登录日志记录
	loginLog := system.SysLoginLog{
		UserName:        username,
		LoginAddress:    clientIP,
		LoginLocation:   location,
		Browser:         userAgentInfo.Browser,
		OperatingSystem: userAgentInfo.OS,
		LoginStatus:     loginStatus,
		OperationalInfo: operationalInfo,
		LoginTime:       time.Now(),
	}

	// 保存到数据库
	if err := global.GVA_DB.Create(&loginLog).Error; err != nil {
		global.GVA_LOG.Error("Failed to record login log", zap.Error(err))
	}
}

// RecordLoginSuccess 记录登录成功日志（用于在登录成功后调用）
func RecordLoginSuccess(username, clientIP, userAgent string) {
	userAgentInfo := utils.ParseUserAgent(userAgent)
	location := utils.GetIPLocation(clientIP)

	loginLog := system.SysLoginLog{
		UserName:        username,
		LoginAddress:    clientIP,
		LoginLocation:   location,
		Browser:         userAgentInfo.Browser,
		OperatingSystem: userAgentInfo.OS,
		LoginStatus:     "成功",
		OperationalInfo: "登录成功",
		LoginTime:       time.Now(),
	}

	if err := global.GVA_DB.Create(&loginLog).Error; err != nil {
		global.GVA_LOG.Error("Failed to record login success log", zap.Error(err))
	}
}

// RecordLoginFailure 记录登录失败日志
func RecordLoginFailure(username, clientIP, userAgent, reason string) {
	userAgentInfo := utils.ParseUserAgent(userAgent)
	location := utils.GetIPLocation(clientIP)

	loginLog := system.SysLoginLog{
		UserName:        username,
		LoginAddress:    clientIP,
		LoginLocation:   location,
		Browser:         userAgentInfo.Browser,
		OperatingSystem: userAgentInfo.OS,
		LoginStatus:     "失败",
		OperationalInfo: reason,
		LoginTime:       time.Now(),
	}

	if err := global.GVA_DB.Create(&loginLog).Error; err != nil {
		global.GVA_LOG.Error("Failed to record login failure log", zap.Error(err))
	}
}
