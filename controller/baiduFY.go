package controller

import (
	"comm-filter/config"
	"comm-filter/utils"
	"crypto/md5"
	"encoding/hex"
	"log"
	"strconv"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
)

func baiduFYRoutes(p router.Party) {
	p.Get("/access", baiduFYAccess)
}

// baiduFYAccess - controller of baiduFY - access
func baiduFYAccess(ctx iris.Context) {
	utils := ctx.Values().Get("utils").(userutils.Utils)
	conf := ctx.Values().Get("config").(*config.Config)
	q := ctx.URLParam("q")

	appid := conf.BaiduFY.Appid
	key := conf.BaiduFY.Key
	salt := strconv.FormatInt(time.Now().Unix()*1000, 10)
	from := "zh"
	to := "en"
	md5ctx := md5.New()
	md5ctx.Write([]byte(appid + q + salt + key))
	sign := hex.EncodeToString(md5ctx.Sum(nil))
	log.Printf("Controller.BaiduFY.baiduFYAccess]sign: %v, salt: %v", sign, salt)

	// 返回
	utils.HTTPSuccess(userutils.SuccBody{
		Data: iris.Map{
			"q":     q,
			"appid": appid,
			"from":  from,
			"to":    to,
			"salt":  salt,
			"sign":  sign,
		},
		Restful: true,
	})
}
