package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", greetHandler)
	log.Fatal(http.ListenAndServe("localhost:9999", nil))
}

func indexHandler(write http.ResponseWriter, req *http.Request) {
	fmt.Println(fmt.Fprintf(write, "host:%s,path=%q\n", req.Host, req.URL.Path))
}

func greetHandler(write http.ResponseWriter, req *http.Request) {
	q := req.URL.Query().Get("q")
	if q != "" {
		fmt.Fprintf(write, "请求参数:%s", reverse(q))
	}
	headers := req.Header
	for key, header := range headers {
		fmt.Fprintf(write, "Header[%q]=%q\n", key, header)
	}
}

func reverse(str string) string {
	s := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}
