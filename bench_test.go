// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/lvjp/go-http-routing-benchmark/router"
)

func benchRequest(b *testing.B, router http.Handler, r *http.Request) {
	w := new(mockResponseWriter)
	u := r.URL
	rq := u.RawQuery
	r.RequestURI = u.RequestURI()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		u.RawQuery = rq
		router.ServeHTTP(w, r)
	}
}

func benchRoutes(b *testing.B, router http.Handler, routes []router.Route) {
	w := new(mockResponseWriter)
	r, _ := http.NewRequest("GET", "/", nil)
	u := r.URL
	rq := u.RawQuery

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, route := range routes {
			r.Method = route.Method
			r.RequestURI = route.Path
			u.Path = route.Path
			u.RawQuery = rq
			router.ServeHTTP(w, r)
		}
	}
}

type microTestCase struct {
	name   string
	method string
	colon  string
	brace  string
	url    string
}

func (mtc *microTestCase) matcher(pt router.ParamType) string {
	switch pt {
	case router.ParamColonType:
		return mtc.colon
	case router.ParamBraceType:
		return mtc.brace
	default:
		panic(fmt.Sprint("Unsupported param type: ", pt))
	}
}

// Micro Benchmarks
func BenchmarkNewMicro(b *testing.B) {
	benchs := []microTestCase{
		{"param/1", "GET", "/user/:name", "/user/{name}", "/user/gordon"},
		{"param/5", "GET", fiveColon, fiveBrace, fiveRoute},
		{"param/20", "GET", twentyColon, twentyBrace, twentyRoute},
	}

	writeBench := microTestCase{
		colon: "/user/:name",
		brace: "/user/{name}",
		url:   "/user/gordon",
	}

	for name, builder := range router.GetRegistry() {
		b.Run(name, func(b *testing.B) {
			for _, bench := range benchs {
				r := builder.Build(
					[]router.Route{{
						Method: "GET",
						Path:   bench.matcher(builder.ParamType()),
					}},
					router.SkipDataMode,
				)

				req, _ := http.NewRequest(bench.method, bench.url, nil)
				b.Run(bench.name, func(b *testing.B) {
					benchRequest(b, r, req)
				})
			}

			r := builder.Build(
				[]router.Route{{
					Method: "GET",
					Path:   writeBench.matcher(builder.ParamType()),
				}},
				router.WriteParameterMode,
			)
			req, _ := http.NewRequest("GET", writeBench.url, nil)
			b.Run("param/write", func(b *testing.B) {
				benchRequest(b, r, req)
			})
		})
	}
}

// func BenchmarkRevel_Param(b *testing.B) {
// 	router := loadRevelSingle("GET", "/user/:name", "RevelController.Handle")

// 	r, _ := http.NewRequest("GET", "/user/gordon", nil)
// 	benchRequest(b, router, r)
// }

// Route with 5 Params (no write)
const fiveColon = "/:a/:b/:c/:d/:e"
const fiveBrace = "/{a}/{b}/{c}/{d}/{e}"
const fiveRoute = "/a/b/c/d/e"

// func BenchmarkRevel_Param5(b *testing.B) {
// 	router := loadRevelSingle("GET", fiveColon, "RevelController.Handle")

// 	r, _ := http.NewRequest("GET", fiveRoute, nil)
// 	benchRequest(b, router, r)
// }

// Route with 20 Params (no write)
const twentyColon = "/:a/:b/:c/:d/:e/:f/:g/:h/:i/:j/:k/:l/:m/:n/:o/:p/:q/:r/:s/:t"
const twentyBrace = "/{a}/{b}/{c}/{d}/{e}/{f}/{g}/{h}/{i}/{j}/{k}/{l}/{m}/{n}/{o}/{p}/{q}/{r}/{s}/{t}"
const twentyRoute = "/a/b/c/d/e/f/g/h/i/j/k/l/m/n/o/p/q/r/s/t"

// func BenchmarkRevel_Param20(b *testing.B) {
// 	router := loadRevelSingle("GET", twentyColon, "RevelController.Handle")

// 	r, _ := http.NewRequest("GET", twentyRoute, nil)
// 	benchRequest(b, router, r)
// }

// func BenchmarkRevel_ParamWrite(b *testing.B) {
// 	router := loadRevelSingle("GET", "/user/:name", "RevelController.HandleWrite")

// 	r, _ := http.NewRequest("GET", "/user/gordon", nil)
// 	benchRequest(b, router, r)
// }
