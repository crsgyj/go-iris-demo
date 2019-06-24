package redisdb

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"time"
)

var (
	Redisdb   *Client
	adminUser = &UserModel{"admin", "123456", time.Now().UTC().String()}
)

func New(options *Options) *Client {
	if Redisdb != nil {
		return Redisdb
	}

	Redisdb = redis.NewClient(options)
	// 初始化唯一admin用户
	initOnlyUser(Redisdb)

	return Redisdb
}

func initOnlyUser(c *Client) {
	// 查询admin用户是否存在
	user := c.Get("user:" + adminUser.Username + "12").Val()
	if user != "" {
		return
	}
	// 创建admin用户
	s, _ := json.Marshal(adminUser)
	c.Set("user:"+adminUser.Username, string(s), 0)
}
