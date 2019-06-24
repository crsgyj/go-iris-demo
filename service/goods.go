package service

import (
	"comm-filter/redis"
	"comm-filter/utils"
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/kataras/iris"
)

// Goods - serv struct
type Goods struct {
	ctx   iris.Context
	redis *redisdb.Client
}

// Item - item of goods
type Item struct {
	GoodsID     string `json:"goods_id" validate:"required,gt=1"`
	Name        string `json:"name" validate:"required,gt=1"`
	NameE       string `json:"name_e"`
	Addr        string `json:"addr" validate:"required"`
	AddrE       string `json:"addr_e"`
	Tag         string `json:"tag"`
	CreatedDate string `json:"created_date"`
}

// ItemList - list of item
type ItemList = []Item

// AddGoods - add item to list
func (g *Goods) AddGoods(source []Item) []redisdb.Z {
	utils := g.ctx.Values().Get("utils").(userutils.Utils)
	score := dateScore(time.Now())

	members := make([]redisdb.Z, len(source))
	for i, item := range source {
		item.CreatedDate = strconv.FormatFloat(score, 'f', -1, 64)
		members[i].Score = score
		members[i].Member = item.GoodsID
		m, _ := json.Marshal(item)
		t := time.Second*3600*24*31 + /* 为了不在同一时间解锁 */ time.Duration(utils.RandInt(600))*time.Second
		g.redis.Set("goods_item:"+item.GoodsID, m, t)
	}
	// reids zadd goods
	cmd := g.redis.ZAdd("goods_zset", members...)

	log.Println(cmd)
	return members
}

// ListOptions - 查询参数
type ListOptions struct {
	Page    int64  `json:"page" validate:"gte=0"`
	PerPage int64  `json:"per_page" validate:"gte=1,lte=50"`
	GoodsID string `json:"goods_id"`
}

// Suggest - 建议
func (g *Goods) Suggest(goodsID string) []string {
	var (
		offset int64 = 0
		limit  int64 = 15
		result *redisdb.StringSliceCmd
	)
	if goodsID != "" {
		result = g.redis.ZRangeByLex("goods_zset", redisdb.ZRangeBy{
			Min:    "[" + goodsID,
			Max:    "+",
			Offset: offset,
			Count:  limit,
		})
	} else {
		result = g.redis.ZRange("goods_zset", offset, offset+limit)
	}
	return result.Val()
}

// List - 获取分页列表
func (g *Goods) List(listOptions *ListOptions) iris.Map {
	offset := (listOptions.Page - 1) * listOptions.PerPage
	var list = make([]*Item, 0)
	var result *redisdb.StringSliceCmd
	// 总数
	count := g.redis.ZCount("goods_zset", "0", "30000101").Val()
	// Get StringSlice
	if listOptions.GoodsID != "" {
		count = 0
		result = g.redis.ZRangeByLex("goods_zset", redisdb.ZRangeBy{
			Min:    "[" + listOptions.GoodsID,
			Max:    "+",
			Offset: offset,
			Count:  listOptions.PerPage,
		})
	} else {
		result = g.redis.ZRange("goods_zset", offset, offset+listOptions.PerPage)
	}
	for _, gid := range result.Val() {
		item := &Item{}
		data := g.redis.Get("goods_item:" + gid).Val()
		// 已过期
		if data == "" {
			list = append(list, &Item{GoodsID: gid, Tag: "已过期"})
			continue
		}
		// jsonParse
		err := json.Unmarshal([]byte(data), item)
		if err != nil {
			list = append(list, &Item{GoodsID: gid, Tag: "数据错误：" + data})
			continue
		}

		// goodsId存在 则最多只查出一条
		if listOptions.GoodsID != "" {
			if gid == listOptions.GoodsID {
				list = []*Item{item}
				count = 1
				break
			}
		} else {
			list = append(list, item)
		}
	}

	return iris.Map{
		"count": count,
		"list":  list,
	}
}

// UpdateItem - 更新某一项
func (g *Goods) UpdateItem(item *Item) userutils.HTTPErr {
	utils := g.ctx.Values().Get("utils").(userutils.Utils)
	// ZRANK key member
	score := g.redis.ZScore("goods_zset", item.GoodsID)

	if score.Err() != nil {
		return userutils.HTTPErr{
			Status: iris.StatusNotFound,
			Error:  errors.New("该货品不存在或已删除"),
			Code:   40004,
		}
	}

	rest := restExpires(int(score.Val()), utils.RandInt(600))
	m, _ := json.Marshal(item)
	g.redis.Set("goods_item:"+item.GoodsID, m, rest)

	return userutils.HTTPErr{}
}

// DelItem - 删除某一项
func (g *Goods) DelItem(goodsID string) userutils.HTTPErr {
	// ZRANK key member
	score := g.redis.ZScore("goods_zset", goodsID)

	if score.Err() != nil {
		return userutils.HTTPErr{
			Status: iris.StatusNotFound,
			Error:  errors.New("该货品不存在或已删除"),
			Code:   40004,
		}
	}

	result1 := g.redis.ZRem("goods_zset", goodsID)
	result2 := g.redis.Del("goods_item:" + goodsID)
	log.Println(result1)
	log.Println(result2)

	return userutils.HTTPErr{}
}

// DateScore - score by date
func dateScore(date time.Time) float64 {
	return float64(date.Year()*10000 + int(date.Month())*100 + date.Day())
}

func restExpires(score int, randT int) time.Duration {
	s := strconv.Itoa(score)
	year, _ := strconv.Atoi(s[:4])
	month, _ := strconv.Atoi(s[4:6])
	day, _ := strconv.Atoi(s[6:])
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	dur := time.Since(t)
	rest := time.Second*3600*24*31 - dur + time.Duration(randT)*time.Second
	return rest
}
