package gmn

import (
	"fmt"
	"net/http"
	"sync"
)

type Handler func(w http.ResponseWriter, req *http.Request)

type Engine struct {
	router map[string]Handler
	mu     sync.Mutex
}

func New() *Engine {
	return &Engine{router: make(map[string]Handler)}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handle, ok := e.router[key]; ok {
		handle(w, req)
	} else {
		fmt.Printf("Not Found Method=%q Path=%q", req.Method, req.URL.Path)
	}
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) addRoute(method string, pattern string, handler Handler) {
	key := method + "-" + pattern
	if _, ok := e.router[key]; !ok {
		e.mu.Lock()
		defer e.mu.Unlock()
		e.router[key] = handler
	}
}

func (e *Engine) GET(pattern string, handler Handler) {
	e.addRoute(http.MethodGet, pattern, handler)
}

func (e *Engine) POST(pattern string, handler Handler) {
	e.addRoute(http.MethodPost, pattern, handler)
}
