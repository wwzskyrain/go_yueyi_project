package middlewares

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
)

// RequestLogMiddleware 记录请求日志
func RequestLogMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		// 用log组件来打印请求日志
		log.Printf("[Request_Log]: %s %s\n", ctx.Request.Method(), ctx.Request.Path())
		ctx.Next(c)
	}
}
