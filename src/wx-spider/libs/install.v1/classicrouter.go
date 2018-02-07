package install

import (
	"net/http"
)

type ClassicRouter struct {
	*http.ServeMux
	m []Route
}

type handle func(http.ResponseWriter, *http.Request)

func NewClassicRouter() *ClassicRouter {
	mux := http.NewServeMux()
	m := []Route{}
	return &ClassicRouter{mux, m}
}

func (p *ClassicRouter) GET(pattern string, h handle) {
	p.handle("GET", pattern, h)
}

func (p *ClassicRouter) POST(pattern string, h handle) {
	p.handle("POST", pattern, h)
}

func (p *ClassicRouter) DELETE(pattern string, h handle) {
	p.handle("DELETE", pattern, h)
}

func (p *ClassicRouter) PUT(pattern string, h handle) {
	p.handle("PUT", pattern, h)
}

func (p *ClassicRouter) handle(method string, pattern string, h handle) {
	p.m = append(p.m, Route{method, pattern})
	p.HandleFunc(pattern, h)
}

func (p *ClassicRouter) Routes() []Route {
	return p.m
}
