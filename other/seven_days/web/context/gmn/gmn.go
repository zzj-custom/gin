package gmn

import (
	"net/http"
	"sync"
)

type Engine struct {
	router *router
}

var (
	engineRepo *Engine
	engineOnce sync.Once
)

func New() *Engine {
	engineOnce.Do(func() {
		engine := new(Engine)
		engine.router = newRouter()
		engineRepo = engine
	})
	return engineRepo
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := newContext(w, req)
	e.router.handle(ctx)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) addRoute(method string, pattern string, handler HandlerFun) {
	e.router.addRoute(method, pattern, handler)
}

func (e *Engine) GET(pattern string, handler HandlerFun) {
	e.addRoute(http.MethodGet, pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFun) {
	e.addRoute(http.MethodPost, pattern, handler)
}
