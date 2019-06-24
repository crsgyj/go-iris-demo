package redisdb

import "github.com/go-redis/redis"

// UserModel - 用户Model
type UserModel struct {
	Username     string `json:"user_name"`
	Password     string `json:"password"`
	RegisterDate string `json:"register_date"`
}

type Client = redis.Client
type Options = redis.Options
type Z = redis.Z
type ZRangeBy = redis.ZRangeBy
type StringSliceCmd = redis.StringSliceCmd
