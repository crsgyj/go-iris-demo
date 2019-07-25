package service

import (
	"comm-filter/redis"

	"github.com/kataras/iris"
)

// Service - struct
type Service struct {
	User  User
	Goods Goods
}

// ServOptions options to init Service
type ServOptions struct {
	Redis *redisdb.Client
}

var serviceHandler iris.Handler

// New add service to iris.Context
func New(options ServOptions) iris.Handler {
	if serviceHandler != nil {
		return serviceHandler
	}

	serviceHandler = func(ctx iris.Context) {
		s := Service{
			User:  User{ctx: ctx, redis: options.Redis},
			Goods: Goods{ctx: ctx, redis: options.Redis},
		}
		ctx.Values().Set("service", s)
		ctx.Next()
	}
	return serviceHandler
}
