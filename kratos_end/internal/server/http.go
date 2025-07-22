package server

import (
	v1 "IOT_Hummingbird_back_end/api/helloworld/v1"
	uv1 "IOT_Hummingbird_back_end/api/user/v1"
	"IOT_Hummingbird_back_end/internal/conf"
	"IOT_Hummingbird_back_end/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, user *service.UserService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	uv1.RegisterUserHTTPServer(srv, user)
	// 手动输出所有接口信息
	logger.Log(log.LevelInfo, "HTTP ROUTE", "POST", "/api/user/register")
	logger.Log(log.LevelInfo, "HTTP ROUTE", "POST", "/api/user/login")
	logger.Log(log.LevelInfo, "HTTP ROUTE", "POST", "/helloworld/greeter")
	return srv
}
