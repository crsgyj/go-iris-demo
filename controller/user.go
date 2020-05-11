package controller

import (
	redisdb "comm-filter/redis"
	"comm-filter/service"
	userutils "comm-filter/utils"

	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
)

func userRoutes(p router.Party) {
	// 登录
	p.Post("/login", userLogin)
	// 以下路由需要登录
	p.Use(loginRequired)
	// 登出
	p.Post("/logout", userLogout)
	// 获取用户数据
	p.Get("/profile", profile)
}

// userLogin - route Handler of user login
func userLogin(ctx iris.Context) {
	utils := ctx.Values().Get("utils").(userutils.Utils)
	service := ctx.Values().Get("service").(service.Service)
	data, err := service.User.Login()
	if err != nil {
		utils.HTTPError(err)
		return
	}

	utils.HTTPSuccess(userutils.SuccBody{
		Data:    data,
		Restful: true,
	})
}

// userLogout - route Handler of user logout
func userLogout(ctx iris.Context) {
	utils := ctx.Values().Get("utils").(userutils.Utils)
	service := ctx.Values().Get("service").(service.Service)

	service.User.Logout()
	utils.HTTPSuccess(userutils.SuccBody{
		Data:    iris.Map{"success": true},
		Restful: true,
	})
}

// profile - route handler of get user profile
func profile(ctx iris.Context) {
	utils := ctx.Values().Get("utils").(userutils.Utils)
	currUser := ctx.Values().Get("currUser").(*redisdb.UserModel)

	utils.HTTPSuccess(userutils.SuccBody{
		Data: iris.Map{
			"user_name":     currUser.Username,
			"register_date": currUser.RegisterDate,
		},
		Restful: true,
	})
}
