
<img src="https://genesem.github.io/logo/xroutes.png" alt="xRoutes" title="xRoutes - go lang http api" align="center"/>
<br/><br/>

### `xRoutes` is *golang* fast http routing API
------------------------------------------------

[![GoDoc](https://godoc.org/github.com/genesem/xroutes?status.png)](https://godoc.org/github.com/genesem/xroutes) [![Build Status](https://travis-ci.org/genesem/xroutes.svg?branch=master)](https://travis-ci.org/genesem/xroutes)

===============


It is forked from <https://github.com/drone/routes> with some mods
and used for internal projects at this time.

### Install:

    go get github.com/genesem/xroutes

### Getting Started

    package main

    import (
        "fmt"
	"github.com/genesem/xroutes"
        "net/http"
    )

    func Whoami(w http.ResponseWriter, r *http.Request) {
        params := r.URL.Query()
        lastName := params.Get(":last")
        firstName := params.Get(":first")
        fmt.Fprintf(w, "you are %s %s", firstName, lastName)
    }

    func main() {
        mux := routes.New()
        mux.Get("/:last/:first", Whoami)

        http.Handle("/", mux)
        http.ListenAndServe(":8088", nil)
    }


### Route Examples

You can create routes for all http methods:

    mux.Get("/:param", handler)
    mux.Put("/:param", handler)
    mux.Post("/:param", handler)
    mux.Patch("/:param", handler)
    mux.Del("/:param", handler)

You can specify custom regular expressions for routes:

    mux.Get("/files/:param(.+)", handler)


### Filters / Middleware
You can apply filters to routes, which is useful for enforcing security,
redirects, etc.

You can, for example, filter all request to enforce some type of security:

    var FilterUser = func(w http.ResponseWriter, r *http.Request) {
    	if r.URL.User == nil || r.URL.User.Username() != "admin" {
    		http.Error(w, "", http.StatusUnauthorized)
    	}
    }

    r.Filter(FilterUser)

You can also apply filters only when certain REST URL Parameters exist:

    r.Get("/:id", handler)
    r.Filter("id", func(rw http.ResponseWriter, r *http.Request) {
		...
	})
