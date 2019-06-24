package userutils

import (
	"math/rand"
	"time"

	"github.com/kataras/iris"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

type Utils struct {
	ctx iris.Context
}

type UtilsOptions struct {
}

// HTTPErr - httpError struct with Status Error, Code
type HTTPErr struct {
	Status int
	Error  error
	Code   int64
}

// New - add utils to ctx
func New(options UtilsOptions) iris.Handler {
	return func(ctx iris.Context) {
		u := Utils{ctx: ctx}
		ctx.Values().Set("utils", u)
		ctx.Next()
	}
}

// HTTPError - handler HTTPError callback
func (u *Utils) HTTPError(httpErr HTTPErr) {
	ctx := u.ctx
	ctx.StatusCode(httpErr.Status)
	ctx.JSON(iris.Map{
		"message": httpErr.Error.Error(),
		"code":    httpErr.Code,
	})
}

// SuccBody - Success body of http
type SuccBody struct {
	Data    interface{}
	Message string
	Code    int64
	Restful bool
}

// HTTPSuccess - handler HTTPSuccess callback
func (u *Utils) HTTPSuccess(succBody SuccBody) {
	ctx := u.ctx
	if ctx.GetStatusCode() != 200 {
		return
	}
	ctx.StatusCode(iris.StatusOK)
	if succBody.Restful == true {
		ctx.JSON(succBody.Data)
	} else {
		ctx.JSON(iris.Map{
			"code":    succBody.Code,
			"data":    succBody.Data,
			"message": succBody.Message,
		})
	}

}

// RandString - return a random String
func (u *Utils) RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func (u *Utils) RandInt(limit int) int {
	return r.Intn(limit)
}
