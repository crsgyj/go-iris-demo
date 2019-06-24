package service

import (
	"bytes"
	"comm-filter/redis"
	"comm-filter/utils"
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
func (u *User) Login() (data iris.Map, err error) {
	// Context
	ctx := u.ctx
	// utils
	utils := ctx.Values().Get("utils").(userutils.Utils)
	formUser := redisdb.UserModel{}
	if err := ctx.ReadJSON(&formUser); err != nil {
		utils.HTTPError(userutils.HTTPErr{
			Status: iris.StatusNotAcceptable,
			Error:  err,
			Code:   40001,
		})
		return nil, err
	}
	// 查询用户信息
	userInfoStr := u.redis.Get("user:" + formUser.Username).Val()
	log.Println(formUser.Username, userInfoStr, u.redis.Get("user:"+formUser.Username))
	if userInfoStr == "" {
		err := errors.New("用户不存在")
		utils.HTTPError(userutils.HTTPErr{
			Status: iris.StatusNotAcceptable,
			Error:  err,
			Code:   40001,
		})
		return nil, err
	}
	// jsonParse
	userInfo := &redisdb.UserModel{}
	json.Unmarshal([]byte(userInfoStr), userInfo)
	if formUser.Password != userInfo.Password {
		err := errors.New("密码错误")
		utils.HTTPError(userutils.HTTPErr{
			Status: iris.StatusNotAcceptable,
			Error:  err,
			Code:   40001,
		})
		return nil, err
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

	// ctx.JSON(iris.Map{
	// 	"success": true,
	// 	"token":   md5str,
	// })
	return iris.Map{
		"user":  userInfo.Username,
		"token": md5str,
	}, nil
}

// Logout - 登出
func (u *User) Logout() {
	token := u.ctx.GetCookie("token")
	log.Println(token)
	u.redis.Del("access_token:" + token)
	u.ctx.RemoveCookie("token")
}

// Profile - 获取用户数据
func (u *User) Profile() (userProfile redisdb.UserModel, err userutils.HTTPErr) {
	token := u.ctx.GetCookie("token")
	if token == "" {
		return redisdb.UserModel{}, userutils.HTTPErr{
			Status: iris.StatusUnauthorized,
			Error:  errors.New("用户未登录"),
			Code:   40001,
		}
	}
	currUserStr := u.redis.Get("access_token:" + token).Val()
	if currUserStr == "" {
		return redisdb.UserModel{}, userutils.HTTPErr{
			Status: iris.StatusUnauthorized,
			Error:  errors.New("用户未登录"),
			Code:   40001,
		}
	}

	currUser := redisdb.UserModel{}
	if err := json.Unmarshal([]byte(currUserStr), &currUser); err != nil {
		u.redis.Del("access_token:" + token)
		return redisdb.UserModel{}, userutils.HTTPErr{
			Status: iris.StatusNotAcceptable,
			Error:  errors.New("未知错误，请重新登录"),
			Code:   40001,
		}
	}
	return currUser, userutils.HTTPErr{}
}
