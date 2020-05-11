package service

import (
	"bytes"
	redisdb "comm-filter/redis"
	userutils "comm-filter/utils"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/kataras/iris"
)

// User UserRoutes
type User struct {
	ctx    iris.Context
	redis  *redisdb.Client
	models []interface{}
}

// Login 登录
func (u *User) Login() (data iris.Map, httpErr *userutils.HTTPErr) {
	// Context
	ctx := u.ctx
	// utils
	utils := ctx.Values().Get("utils").(userutils.Utils)
	formUser := redisdb.UserModel{}
	if err := ctx.ReadJSON(&formUser); err != nil {
		httpErr = &userutils.HTTPErr{
			Status:  iris.StatusNotAcceptable,
			Error:   err,
			Code:    40001,
			Restful: true,
		}
		return
	}
	// 查询用户信息
	userInfoStr := u.redis.Get("user:" + formUser.Username).Val()
	log.Println(formUser.Username, userInfoStr, u.redis.Get("user:"+formUser.Username))
	if userInfoStr == "" {
		httpErr = &userutils.HTTPErr{
			Status:  iris.StatusNotAcceptable,
			Error:   errors.New("用户不存在"),
			Code:    40001,
			Restful: true,
		}
		return
	}
	// jsonParse
	userInfo := &redisdb.UserModel{}
	json.Unmarshal([]byte(userInfoStr), userInfo)
	if formUser.Password != userInfo.Password {
		httpErr = &userutils.HTTPErr{
			Status:  iris.StatusNotAcceptable,
			Error:   errors.New("密码错误"),
			Code:    40001,
			Restful: true,
		}
		return
	}
	var buffer bytes.Buffer
	userJSONbyte, _ := json.Marshal(userInfo)
	buffer.Write(userJSONbyte)
	buffer.WriteString(utils.RandString(8))
	md5byte := md5.Sum(buffer.Bytes())
	md5str := hex.EncodeToString(md5byte[:])
	u.redis.Set("access_token:"+md5str, string(userJSONbyte), 24*60*60*time.Second)
	// cookie
	u.ctx.SetCookie(&http.Cookie{
		Name:   "token",
		Value:  md5str,
		MaxAge: 24 * 60 * 60,
		Path:   "/",
	})

	data = iris.Map{
		"user":  userInfo.Username,
		"token": md5str,
	}
	return
}

// Logout - 登出
func (u *User) Logout() {
	token := u.ctx.GetCookie("token")
	log.Println(token)
	u.redis.Del("access_token:" + token)
	u.ctx.RemoveCookie("token")
}

// Profile - 获取用户数据
func (u *User) Profile() (userProfile *redisdb.UserModel, httpErr *userutils.HTTPErr) {
	token := u.ctx.GetCookie("token")
	if token == "" {
		httpErr = &userutils.HTTPErr{
			Status:  iris.StatusUnauthorized,
			Error:   errors.New("用户未登录"),
			Code:    40001,
			Restful: true,
		}
		return
	}
	currUserStr := u.redis.Get("access_token:" + token).Val()
	if currUserStr == "" {
		httpErr = &userutils.HTTPErr{
			Status:  iris.StatusUnauthorized,
			Error:   errors.New("用户未登录"),
			Code:    40001,
			Restful: true,
		}
		return
	}

	userProfile = &redisdb.UserModel{}
	if err := json.Unmarshal([]byte(currUserStr), userProfile); err != nil {
		u.redis.Del("access_token:" + token)
		httpErr = &userutils.HTTPErr{
			Status:  iris.StatusNotAcceptable,
			Error:   errors.New("未知错误，请重新登录"),
			Code:    40001,
			Restful: true,
		}
		return
	}
	return
}
