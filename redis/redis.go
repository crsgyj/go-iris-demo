package redisdb

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

var (
	Redisdb   *Client
	adminUser = &UserModel{"admin", "123456", time.Now().UTC().String()}
)

func New(options *Options) *Client {
	var (
		// 连接错误
		connErr error
	)
	if Redisdb != nil {
		return Redisdb
	}

	Redisdb = redis.NewClient(options)

	if connErr = Redisdb.ClientGetName().Err(); connErr != redis.Nil {
		panic("数据库连接失败。" + connErr.Error())
	}
	// 初始化唯一admin用户
	initOnlyUser(Redisdb)

	return Redisdb
}

func initOnlyUser(c *Client) {
	// 查询admin用户是否存在
	user := c.Get("user:" + adminUser.Username).Val()
	if user != "" {
		return
	}
	// 创建admin用户
	s, _ := json.Marshal(adminUser)
	c.Set("user:"+adminUser.Username, string(s), 0)
}
