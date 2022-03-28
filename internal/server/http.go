package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "kratos-layout/api/open/v1"
	"kratos-layout/internal/conf"
	"kratos-layout/internal/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, open *service.OpenService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		//修改api返回的数据格式
		http.ResponseEncoder(ResponseEncoder),
		http.ErrorEncoder(ErrorEncoder),

		http.Middleware(
			recovery.Recovery(),
			//selector.Server(middleware.VerifySign()).Prefix("/api.open.v1.Open/").Build(),
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
	v1.RegisterOpenHTTPServer(srv, open)
	return srv
}
