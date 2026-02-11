package support_web

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/yueyi/projects/support_web/middlewares"
	"log"
)

// InitWebServer 使用hertz初始化web服务器
//
//	1.并加载中间件，
//	2.加载router
//	3.并启动web服务器
//	4.并在结束时关闭web服务器
func InitWebServer() {
	serverOpts := []config.Option{
		server.WithHostPorts(":8088"),
	}
	h := server.Default(serverOpts...)
	// 注册请求日志中间件
	h.Use(middlewares.RequestLogMiddleware())
	// 注册路由
	RegisterRouter(h)
	// 启动web服务器
	h.Spin()
	//	优雅的关闭web服务器
	err := h.Shutdown(context.Background())

	if err != nil {
		log.Printf("Shutdown web server failed: %v\n", err)
		return
	}
	// 关闭web服务器成功
	log.Printf("Shutdown web server success\n")
}
