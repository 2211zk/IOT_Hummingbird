package server

import (
	v1 "kratos/api/helloworld/v1"
	"kratos/internal/conf"
	"kratos/internal/handler"
	"kratos/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	httpx "github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, userService *service.UserService, logger log.Logger) *httpx.Server {
	var opts = []httpx.ServerOption{
		httpx.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, httpx.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, httpx.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, httpx.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := httpx.NewServer(opts...)

	handler.RegisterUserRoutes(srv, userService)

	v1.RegisterGreeterHTTPServer(srv, greeter)
	return srv
}
