package login

import (
	"github.com/gin-gonic/gin"
	"go-api/app/controller/login"
	"go-api/app/middleware"
)

func Router(engine *gin.Engine) {
	g := engine.Group(
		"/login",
	).Use(
		middleware.ReformatBody,
	)
	g.POST("/register", login.Register)
	g.POST("/index", login.UserLogin)
}
