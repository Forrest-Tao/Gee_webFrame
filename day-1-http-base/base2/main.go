package main

import (
	"fmt"
	"log"
	"net/http"
)

// ServeHttp interface
// handler  func (w http.ResponseWriter, r http.Request)
// 自定义 ServeHttp

type Engine struct{}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path=%q", r.URL.Path)
	case "/hello":
		for k, v := range r.Header {
			fmt.Fprintf(w, "header[%q]=%q", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND %s", r.URL.Path)
	}
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9091", engine))
}
