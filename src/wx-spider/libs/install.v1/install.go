package install

import (
	"fmt"
	"net/http"

	"bilibili.com/libs/alog.v1"
	"github.com/urfave/negroni"
)

type Install struct {
	*negroni.Negroni
	handler Handler
}

func New() *Install {
	n := negroni.New(negroni.NewRecovery(), alog.NewLogger())
	n.UseFunc(commonHandleFunc)
	return &Install{n, nil}
}

func (p *Install) WithRouter(handler Handler) {
	p.handler = handler
	p.UseHandler(handler)
}

func (p *Install) UseStatic(dir string, prefix string) {
	static := negroni.NewStatic(http.Dir(dir))
	static.Prefix = prefix
	p.Use(static)
}

func (p *Install) ListenAndServe(port string) {
	for _, route := range p.handler.Routes() {
		fmt.Printf("%s:\t\t%s\n", route.Method, route.Path)
	}
	p.Run(port)
}

// define your middleware below by negroni.Handler and http.Handler
func commonHandleFunc(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(w, r)
}
