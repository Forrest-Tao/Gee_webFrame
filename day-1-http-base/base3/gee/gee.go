package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addPatten(method string, patten string, handler HandlerFunc) {
	key := method + patten
	engine.router[key] = handler
}

func (engine *Engine) GET(patten string, handler HandlerFunc) {
	engine.addPatten("GET", patten, handler)
}

func (engine *Engine) POST(patten string, handler HandlerFunc) {
	engine.addPatten("POST", patten, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + r.URL.Path
	// 匹配路由
	if handler, ok := engine.router[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND %q", r.URL.Path)
	}
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
