package install

import (
	"github.com/julienschmidt/httprouter"
)

type Router struct {
	*httprouter.Router
	m []Route // key: method, value: path
}

func NewRouter() *Router {
	router := httprouter.New()
	m := []Route{}
	return &Router{router, m}
}

func (r *Router) GET(path string, handle httprouter.Handle) {
	r.handle("GET", path, handle)
}

func (r *Router) POST(path string, handle httprouter.Handle) {
	r.handle("POST", path, handle)
}

func (r *Router) DELETE(path string, handle httprouter.Handle) {
	r.handle("DELETE", path, handle)
}

func (r *Router) PUT(path string, handle httprouter.Handle) {
	r.handle("PUT", path, handle)
}

func (r *Router) PATCH(path string, handle httprouter.Handle) {
	r.handle("PATCH", path, handle)
}

func (r *Router) HEAD(path string, handle httprouter.Handle) {
	r.handle("HEAD", path, handle)
}

func (r *Router) OPTIONS(path string, handle httprouter.Handle) {
	r.handle("OPTIONS", path, handle)
}

func (r *Router) Routes() []Route {
	return r.m
}

func (r *Router) handle(method string, path string, handle httprouter.Handle) {
	r.m = append(r.m, Route{method, path})
	r.Handle(method, path, handle)
}
