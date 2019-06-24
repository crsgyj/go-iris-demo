package main

import (
	"comm-filter/config"
	"comm-filter/controller"
	"comm-filter/redis"
	"comm-filter/service"
	"comm-filter/utils"

	"github.com/kataras/iris/middleware/logger"
	"github.com/rs/cors"

	"github.com/kataras/iris"
)

func main() {
	app := newApp()
	port := config.Conf.App.Port
	app.Run(iris.Addr(port), iris.WithoutServerError(iris.ErrServerClosed))
}

func newApp() *iris.Application {
	// app
	app := iris.New()
	// config
	app.Use(func(ctx iris.Context) {
		ctx.Values().Set("config", config.Conf)
		ctx.Next()
	})
	// cors
	corsWrapper := cors.New(cors.Options{
		AllowedOrigins:   config.Conf.App.AllowedOrigins,
		AllowedMethods:   []string{"POST", "GET", "PUT", "PATCH", "DELETE", "OPTION", "HEAD"},
		AllowCredentials: true,
	}).ServeHTTP
	app.WrapRouter(corsWrapper)

	// logger
	app.Logger().SetLevel("debug")
	app.Use(logger.New())
	app.Use(userutils.New(userutils.UtilsOptions{}))
	// cache
	// app.Use(cache.Handler(100 * time.Millisecond))
	// redis client
	redis := redisdb.New(&redisdb.Options{
		Addr:     config.Conf.Redis.Addr,
		Password: config.Conf.Redis.Password,
	})
	// 退出时关闭redis连接
	iris.RegisterOnInterrupt(func() {
		redis.Close()
	})
	// 加载service层
	app.Use(service.New(
		service.ServOptions{Redis: redis}))

	// 加载路由
	app.PartyFunc("/", controller.RootRoutes).AllowMethods(iris.MethodOptions)

	return app
}
