package main

import (
	"go-api/other/seven_days/web/context/gmn"
	"net/http"
)

func main() {
	engine := gmn.New()
	engine.GET("/", func(ctx *gmn.Context) {
		ctx.Json(http.StatusOK, gmn.H{
			"data": "success",
		})
	})
	engine.GET("/hello", func(ctx *gmn.Context) {
		ctx.String(http.StatusOK, "hello %s, you are at %s", ctx.Query("name"), ctx.Path)
	})
	engine.Run("127.0.0.1:8975")
}
