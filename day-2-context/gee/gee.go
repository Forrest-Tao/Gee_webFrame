package gee

import "net/http"

type HandlerFunc func(*Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method, patten string, handler HandlerFunc) {
	engine.router.addRoute(method, patten, handler)
}

func (engine *Engine) GET(patten string, handler HandlerFunc) {
	engine.addRoute("GET", patten, handler)
}

func (engine *Engine) POST(patten string, handler HandlerFunc) {
	engine.addRoute("POST", patten, handler)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	engine.router.handle(c)
}
