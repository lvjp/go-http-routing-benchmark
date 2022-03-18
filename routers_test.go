package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lvjp/go-http-routing-benchmark/router"
)

var (
	// load functions of all routers
	routers = []struct {
		name string
		load func(routes []router.Route) http.Handler
	}{
		{"Beego", loadBeego},
		{"Chi", loadChi},
		{"Echo", loadEcho},
		{"Gin", loadGin},
		{"GoRestful", loadGoRestful},
		{"GorillaMux", loadGorillaMux},
		{"GowwwRouter", loadGowwwRouter},
		{"HttpRouter", loadHttpRouter},
		{"HttpTreeMux", loadHttpTreeMux},
		{"Macaron", loadMacaron},
		{"Pat", loadPat},
		// {"Revel", loadRevel},
	}

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

func TestRouters(t *testing.T) {
	loadTestHandler = true

	for name, builder := range router.GetRegistry() {
		name := name
		builder := builder
		routers = append(
			routers,
			struct {
				name string
				load func(routes []router.Route) http.Handler
			}{
				name: name,
				load: func(routes []router.Route) http.Handler {
					return builder.Build(routes, router.WritePathMode)
				},
			})
	}

	for _, router := range routers {
		req, _ := http.NewRequest("GET", "/", nil)
		u := req.URL
		rq := u.RawQuery

		t.Run(router.name, func(t *testing.T) {
			for _, api := range apis {
				t.Run(api.name, func(t *testing.T) {
					r := router.load(api.routes)

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
								router.name, api.name, w.Code, w.Body.String(), route.Method, route.Path,
							)
						}
					}
				})
			}
		})
	}

	loadTestHandler = false
}
