package xroutes

import (
	"net/http"
	"regexp"
)

const (
	CONNECT = "CONNECT"
	DELETE  = "DELETE"
	GET     = "GET"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"
	PATCH   = "PATCH"
	POST    = "POST"
	PUT     = "PUT"
	TRACE   = "TRACE"
)

type route struct {
	method  string
	regex   *regexp.Regexp
	params  map[int]string
	handler http.HandlerFunc
}

type RouteMux struct {
	routes  []*route
	filters []http.HandlerFunc
}
