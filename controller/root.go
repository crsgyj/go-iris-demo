package controller

import (
	"comm-filter/service"
	"comm-filter/utils"

	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
)

// RootRoutes - the routes entry
func RootRoutes(p router.Party) {
	p.Get("/status.ok", statusOK)
	p.PartyFunc("/admin", func(p iris.Party) {
		// 用户模块
		p.PartyFunc("/user", userRoutes)
		// 以下路由需要登录
		p.Use(loginRequired)
		// 商品模块
		p.PartyFunc("/goods", goodsRoutes)
		// 百度翻译
		p.PartyFunc("/baiduFY", baiduFYRoutes)
	})
}

// loginRequired - 登录校验中间件
func loginRequired(ctx iris.Context) {
	service := ctx.Values().Get("service").(service.Service)
	utils := ctx.Values().Get("utils").(userutils.Utils)
	profile, err := service.User.Profile()
	if err.Error != nil {
		utils.HTTPError(err)
		return
	}
	ctx.Values().Set("currUser", profile)
	ctx.Next()
}

/**
 *
 * @api {Get} /status.ok
 * @apiName status.ok
 * @apiGroup base
 * @apiVersion 1.0.0
 *
 * @apiSuccessExample {type} Success-Response: "ok"
 *
 */
func statusOK(ctx iris.Context) {
	ctx.WriteString("ok")
}
