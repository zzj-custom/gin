package router

import (
	"github.com/gin-gonic/gin"
	"go-api/cmd/router/login"
	"go-api/cmd/router/user"
)

type Router func(engine *gin.Engine)

var routers []Router

func Register(router ...Router) {
	routers = append(routers, router...)
}

func Routers() []Router {
	return routers
}

func init() {
	Register(login.Router)
	Register(user.Router)
}
