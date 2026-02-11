package support_web

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/yueyi/projects/support_web/controller"
)

// RegisterRouter 统一注册路由，以及中间件
func RegisterRouter(h *server.Hertz) {
	// 注册路由
	h.GET("/hello", controller.HelloWorld)
	// 注册search/users2路由
	h.GET("/search/users2", controller.SearchUsers2)
}
