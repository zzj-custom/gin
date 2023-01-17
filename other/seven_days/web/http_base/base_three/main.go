package main

import (
	"fmt"
	"go-api/other/seven_days/web/http_base/base_three/gmn"
	"net/http"
)

func main() {
	engine := gmn.New()
	engine.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Host=%q,Path=%q\n", req.Host, req.URL.Path)
	})
	engine.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Host=%q,Path=%q,Query=%q\n", req.Host, req.URL, req.Form)
	})
	engine.Run("127.0.0.1:8976")
}
