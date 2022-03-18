package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lvjp/go-http-routing-benchmark/router"
)

var (
	// all APIs
	apis = []struct {
		name   string
		routes []router.Route
	}{
		{"GitHub", githubAPI},
		{"GPlus", gplusAPI},
		{"Parse", parseAPI},
		{"Static", staticRoutes},
	}
)

func TestRouters_pathMode(t *testing.T) {
	for _, builder := range router.GetRegistry() {
		req, _ := http.NewRequest("GET", "/", nil)
		u := req.URL
		rq := u.RawQuery

		t.Run(builder.Name(), func(t *testing.T) {
			for _, api := range apis {
				t.Run(api.name, func(t *testing.T) {
					r := builder.Build(api.routes, router.WritePathMode)

					for _, route := range api.routes {
						w := httptest.NewRecorder()
						req.Method = route.Method
						req.RequestURI = route.Path
						u.Path = route.Path
						u.RawQuery = rq
						r.ServeHTTP(w, req)
						if w.Code != 200 || w.Body.String() != route.Path {
							t.Errorf(
								"%s in API %s: %d - %s; expected %s %s\n",
								builder.Name(), api.name, w.Code, w.Body.String(), route.Method, route.Path,
							)
						}
					}
				})
			}
		})
	}
}

func TestRouters_parameterMode(t *testing.T) {
	name := "gordon"
	method := "GET"
	path := "/user/" + name

	for _, builder := range router.GetRegistry() {
		req, _ := http.NewRequest(method, path, nil)
		u := req.URL
		rq := u.RawQuery

		t.Run(builder.Name(), func(t *testing.T) {
			var matcher string
			switch builder.ParamType() {
			case router.ParamColonType:
				matcher = "/user/:name"
			case router.ParamBraceType:
				matcher = "/user/{name}"
			default:
				panic(fmt.Sprint("Unsupported param type: ", builder.ParamType()))
			}

			r := builder.Build(
				[]router.Route{{
					Method: method,
					Path:   matcher,
				}},
				router.WriteParameterMode,
			)
			w := httptest.NewRecorder()
			req.Method = method
			req.RequestURI = path
			u.Path = path
			u.RawQuery = rq
			r.ServeHTTP(w, req)
			if w.Code != 200 || w.Body.String() != name {
				t.Errorf(
					"%s: %d - %s; expected '%#v' '%#v'\n",
					builder.Name(), w.Code, w.Body.String(), method, name,
				)
			}
		})
	}
}
