package main

import (
	"testing"

	"github.com/kataras/iris"
	"github.com/kataras/iris/httptest"
)

var app *iris.Application

func init() {
	app = newApp()
}

func TestStatusOK(t *testing.T) {
	e := httptest.New(t, app) // httptest.URL("http://localhost"+config.Conf.App.Port)
	t1 := e.GET("/status.ok").Expect().Status(iris.StatusOK)
	t1.Body().Equal("ok")
}
