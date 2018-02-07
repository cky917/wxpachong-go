package install

import (
	"net/http"
)

type Route struct {
	Method string
	Path   string
}

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
	Routes() []Route
}
