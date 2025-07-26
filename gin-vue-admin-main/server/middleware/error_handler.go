package middleware

import (
	"context"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ErrorHandler 错误处理中间件
func ErrorHandler() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			handlePanicError(c, err)
		} else if err, ok := recovered.(error); ok {
			handlePanicError(c, err.Error())
		} else {
			handlePanicError(c, fmt.Sprintf("Unknown error: %v", recovered))
		}
	})
}

// handlePanicError 处理panic错误
func handlePanicError(c *gin.Context, err string) {
	// 记录错误日志
	global.GVA_LOG.Error("Panic recovered",
		zap.String("error", err),
		zap.String("path", c.Request.URL.Path),
		zap.String("method", c.Request.Method),
		zap.String("ip", c.ClientIP()),
		zap.String("user_agent", c.Request.UserAgent()),
		zap.String("stack", string(debug.Stack())),
		zap.Time("time", time.Now()),
	)

	// 返回错误响应
	response.FailWithMessage("服务器内部错误，请稍后重试", c)
	c.Abort()
}

// BusinessErrorHandler 业务错误处理中间件
func BusinessErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 检查是否有错误
		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			// 记录业务错误
			global.GVA_LOG.Warn("Business error",
				zap.String("error", err.Error()),
				zap.String("path", c.Request.URL.Path),
				zap.String("method", c.Request.Method),
				zap.String("ip", c.ClientIP()),
				zap.Time("time", time.Now()),
			)

			// 根据错误类型返回不同的响应
			handleBusinessError(c, err.Err)
		}
	}
}

// handleBusinessError 处理业务错误
func handleBusinessError(c *gin.Context, err error) {
	switch err.Error() {
	case "部门名称已存在":
		response.FailWithDetailed(nil, "部门名称已存在", c)
	case "上级部门不能是自身或子部门":
		response.FailWithDetailed(nil, "上级部门不能是自身或子部门", c)
	case "该部门下还有子部门，无法删除":
		response.FailWithDetailed(nil, "该部门下还有子部门，无法删除", c)
	case "部门不存在":
		response.FailWithDetailed(nil, "部门不存在", c)
	case "设备不存在":
		response.FailWithDetailed(nil, "设备不存在", c)
	case "设备已被其他部门关联":
		response.FailWithDetailed(nil, "设备已被其他部门关联", c)
	default:
		response.FailWithMessage(err.Error(), c)
	}
}

// RequestTimeoutHandler 请求超时处理中间件
func RequestTimeoutHandler(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置请求超时
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)

		// 创建一个channel来接收处理完成的信号
		done := make(chan bool, 1)

		go func() {
			c.Next()
			done <- true
		}()

		select {
		case <-done:
			// 请求正常完成
			return
		case <-ctx.Done():
			// 请求超时
			response.FailWithMessage("请求超时，请稍后重试", c)
			c.Abort()
			return
		}
	}
}

// RateLimitErrorHandler 限流错误处理
func RateLimitErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 检查是否触发限流
		if c.GetHeader("X-RateLimit-Remaining") == "0" {
			global.GVA_LOG.Warn("Rate limit exceeded",
				zap.String("ip", c.ClientIP()),
				zap.String("path", c.Request.URL.Path),
				zap.Time("time", time.Now()),
			)
		}
	}
}

// DatabaseErrorHandler 数据库错误处理
func DatabaseErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if recovered := recover(); recovered != nil {
				if err, ok := recovered.(error); ok {
					handleDatabaseError(c, err)
				}
			}
		}()

		c.Next()
	}
}

// handleDatabaseError 处理数据库错误
func handleDatabaseError(c *gin.Context, err error) {
	global.GVA_LOG.Error("Database error",
		zap.String("error", err.Error()),
		zap.String("path", c.Request.URL.Path),
		zap.String("method", c.Request.Method),
		zap.Time("time", time.Now()),
	)

	// 根据错误类型返回不同响应
	errorMsg := err.Error()
	if contains(errorMsg, "connection") {
		response.FailWithMessage("数据库连接失败，请稍后重试", c)
	} else if contains(errorMsg, "timeout") {
		response.FailWithMessage("数据库操作超时，请稍后重试", c)
	} else if contains(errorMsg, "duplicate") {
		response.FailWithMessage("数据已存在，请检查后重试", c)
	} else {
		response.FailWithMessage("数据操作失败，请稍后重试", c)
	}

	c.Abort()
}

// ValidationErrorHandler 参数验证错误处理
func ValidationErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 检查绑定错误
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				if err.Type == gin.ErrorTypeBind {
					global.GVA_LOG.Warn("Validation error",
						zap.String("error", err.Error()),
						zap.String("path", c.Request.URL.Path),
						zap.Time("time", time.Now()),
					)

					response.FailWithMessage("参数验证失败: "+err.Error(), c)
					c.Abort()
					return
				}
			}
		}
	}
}

// contains 检查字符串是否包含子字符串
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr ||
		(len(s) > len(substr) &&
			(s[:len(substr)] == substr ||
				s[len(s)-len(substr):] == substr ||
				containsInMiddle(s, substr))))
}

func containsInMiddle(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// HealthCheckHandler 健康检查处理
func HealthCheckHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/health" {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"time":   time.Now().Unix(),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
