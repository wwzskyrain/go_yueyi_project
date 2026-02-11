package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"log"
)

// HelloWorld 处理hello请求
func HelloWorld(c context.Context, ctx *app.RequestContext) {
	ctx.JSON(consts.StatusOK, map[string]string{"message": "Hello, Hertz!"})
}

// 定义User结构体
type User struct {
	ID        int    `json:"id"`
	HtmlUrl   string `json:"html_url"`
	AvatarUrl string `json:"avatar_url"`
	Login     string `json:"login"`
}

// SearchUsers2 处理search/users2请求
func SearchUsers2(c context.Context, ctx *app.RequestContext) {
	// 从请求参数中获取q参数
	q := ctx.Query("q")
	log.Printf("q: %s\n", q)
	users := doSearchUsers2(q)
	ctx.JSON(consts.StatusOK, users)
}

func doSearchUsers2(key string) []User {
	log.Printf("key: %s\n", key)
	users := getAllUsers()
	//	先全部返回
	return users
}
func getAllUsers() []User {
	return []User{
		{ID: 1, HtmlUrl: "https://github.com/yueyi", AvatarUrl: "https://github.com/yueyi.png", Login: "yueyi"},
		{ID: 2, HtmlUrl: "https://github.com/wwzskyrain", AvatarUrl: "https://github.com/wwzskyrain.png", Login: "yueyi2"},
		{ID: 3, HtmlUrl: "https://github.com/yueyi", AvatarUrl: "https://github.com/yueyi.png", Login: "yueyi3"},
		{ID: 4, HtmlUrl: "https://github.com/wwzskyrain", AvatarUrl: "https://github.com/wwzskyrain.png", Login: "wwzskyrain"},
		{ID: 5, HtmlUrl: "https://github.com/wwzskyrain", AvatarUrl: "https://github.com/wwzskyrain.png", Login: "wwzskyrain2"},
		{ID: 6, HtmlUrl: "https://github.com/wwzskyrain", AvatarUrl: "https://github.com/wwzskyrain.png", Login: "wwzskyrain3"},
	}
}
