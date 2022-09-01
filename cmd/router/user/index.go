package user

import (
	"github.com/gin-gonic/gin"
	"go-api/app/middleware"
)

func Router(engine *gin.Engine) {
	g := engine.Group(
		"/user",
	).Use(
		//middleware.Token,
		middleware.NotifyTimeout,
		middleware.ReformatBody,
	)
	g.POST("/index")
}
