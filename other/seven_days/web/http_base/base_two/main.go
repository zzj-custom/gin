package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Engine struct{}

func New() *Engine {
	return new(Engine)
}

func (e *Engine) ServeHTTP(write http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	switch path {
	case "/":
		fmt.Fprintf(write, "Host=%q,path=%q\n", req.Host, path)
	case "/hello":
		headers := req.Header
		for key, header := range headers {
			fmt.Fprintf(write, "Header[%q]=%q\n", key, header)
		}
	default:
		fmt.Fprintf(write, "Page Not Found Path = %q\n", req.URL)
	}
}

func main() {
	engine := New()
	log.Fatal(http.ListenAndServe("localhost:9999", engine))
}
