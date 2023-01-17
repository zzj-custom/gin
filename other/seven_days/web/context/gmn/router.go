package gmn

import (
	"fmt"
	"sync"
)

type HandlerFun func(ctx *Context)

type router struct {
	handlers map[string]HandlerFun
	mu       sync.Mutex
}

var (
	routerRepo *router
	routerOnce sync.Once
)

func newRouter() *router {
	routerOnce.Do(func() {
		router := new(router)
		router.handlers = make(map[string]HandlerFun)
		routerRepo = router
	})
	return routerRepo
}

func (r *router) addRoute(method string, pattern string, handler HandlerFun) {
	key := method + "-" + pattern
	if _, ok := r.handlers[key]; !ok {
		r.handlers[key] = handler
	}
}

func (r *router) handle(ctx *Context) {
	key := ctx.Method + "-" + ctx.Path
	if handler, ok := r.handlers[key]; ok {
		handler(ctx)
	} else {
		fmt.Printf("Page Not Found Method=%q Path=%q", ctx.Method, ctx.Path)
	}
}
