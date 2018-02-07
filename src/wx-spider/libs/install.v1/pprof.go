package install

import (
	"net/http"
	"net/http/pprof"
	"github.com/julienschmidt/httprouter"
)

func (install *Install) DoPProf(r *Router) {
	r.GET("/debug/pprof", wrap(pprof.Index))
	r.GET("/debug/pprof/cmdline", wrap(pprof.Cmdline))
	r.GET("/debug/pprof/symbol", wrap(pprof.Symbol))
	r.GET("/debug/pprof/profile", wrap(pprof.Profile))
	r.GET("/debug/pprof/heap", wrap(pprof.Handler("heap").ServeHTTP))
	r.GET("/debug/pprof/goroutine", wrap(pprof.Handler("goroutine").ServeHTTP))
	r.GET("/debug/pprof/block", wrap(pprof.Handler("block").ServeHTTP))
	r.GET("/debug/pprof/threadcreate", wrap(pprof.Handler("threadcreate").ServeHTTP))
}

func wrap(handler http.HandlerFunc) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		handler(w, r)
	})
}
