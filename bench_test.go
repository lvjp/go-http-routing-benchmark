// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
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

// Micro Benchmarks
func BenchmarkNewMicro(b *testing.B) {
	benchs := []struct {
		name   string
		method string
		colon  string
		brace  string
		url    string
	}{
		{"param/1", "GET", "/user/:name", "/user/{name}", "/user/gordon"},
		{"param/5", "GET", fiveColon, fiveBrace, fiveRoute},
		{"param/20", "GET", twentyColon, twentyBrace, twentyRoute},
	}

	for name, builder := range router.GetRegistry() {
		b.Run(name, func(b *testing.B) {
			for _, bench := range benchs {
				var matcher string
				switch builder.ParamType() {
				case router.ParamColonType:
					matcher = bench.colon
				case router.ParamBraceType:
					matcher = bench.brace
				default:
					b.Fatal("Unsupported param type:", builder.ParamType())
				}
				r := builder.BuildSingle("GET", matcher, router.SkipDataMode)

				req, _ := http.NewRequest(bench.method, bench.url, nil)
				b.Run(bench.name, func(b *testing.B) {
					benchRequest(b, r, req)
				})
			}

			var matcher string
			switch builder.ParamType() {
			case router.ParamColonType:
				matcher = "/user/:name"
			case router.ParamBraceType:
				matcher = "/user/{name}"
			default:
				b.Fatal("Unsupported param type:", builder.ParamType())
			}
			r := builder.BuildSingle("GET", matcher, router.WriteParameterMode)
			req, _ := http.NewRequest("GET", "/user/gordon", nil)
			b.Run("param/write", func(b *testing.B) {
				benchRequest(b, r, req)
			})
		})
	}
}

func BenchmarkBeego_Param(b *testing.B) {
	router := loadBeegoSingle("GET", "/user/:name", beegoHandler)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkChi_Param(b *testing.B) {
	router := loadChiSingle("GET", "/user/{name}", httpHandlerFunc)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkEcho_Param(b *testing.B) {
	router := loadEchoSingle("GET", "/user/:name", echoHandler)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGin_Param(b *testing.B) {
	router := loadGinSingle("GET", "/user/:name", ginHandle)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGoRestful_Param(b *testing.B) {
	router := loadGoRestfulSingle("GET", "/user/{name}", goRestfulHandler)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGorillaMux_Param(b *testing.B) {
	router := loadGorillaMuxSingle("GET", "/user/{name}", httpHandlerFunc)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGowwwRouter_Param(b *testing.B) {
	router := loadGowwwRouterSingle("GET", "/user/:name", http.HandlerFunc(httpHandlerFunc))

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpRouter_Param(b *testing.B) {
	router := loadHttpRouterSingle("GET", "/user/:name", httpRouterHandle)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpTreeMux_Param(b *testing.B) {
	router := loadHttpTreeMuxSingle("GET", "/user/:name", httpTreeMuxHandler)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkMacaron_Param(b *testing.B) {
	router := loadMacaronSingle("GET", "/user/:name", macaronHandler)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkPat_Param(b *testing.B) {
	router := loadPatSingle("GET", "/user/:name", http.HandlerFunc(httpHandlerFunc))

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}

// func BenchmarkRevel_Param(b *testing.B) {
// 	router := loadRevelSingle("GET", "/user/:name", "RevelController.Handle")

// 	r, _ := http.NewRequest("GET", "/user/gordon", nil)
// 	benchRequest(b, router, r)
// }

// Route with 5 Params (no write)
const fiveColon = "/:a/:b/:c/:d/:e"
const fiveBrace = "/{a}/{b}/{c}/{d}/{e}"
const fiveRoute = "/test/test/test/test/test"

func BenchmarkBeego_Param5(b *testing.B) {
	router := loadBeegoSingle("GET", fiveColon, beegoHandler)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkChi_Param5(b *testing.B) {
	router := loadChiSingle("GET", fiveBrace, httpHandlerFunc)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkEcho_Param5(b *testing.B) {
	router := loadEchoSingle("GET", fiveColon, echoHandler)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGin_Param5(b *testing.B) {
	router := loadGinSingle("GET", fiveColon, ginHandle)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGoRestful_Param5(b *testing.B) {
	router := loadGoRestfulSingle("GET", fiveBrace, goRestfulHandler)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGorillaMux_Param5(b *testing.B) {
	router := loadGorillaMuxSingle("GET", fiveBrace, httpHandlerFunc)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGowwwRouter_Param5(b *testing.B) {
	router := loadGowwwRouterSingle("GET", fiveColon, http.HandlerFunc(httpHandlerFunc))

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpRouter_Param5(b *testing.B) {
	router := loadHttpRouterSingle("GET", fiveColon, httpRouterHandle)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpTreeMux_Param5(b *testing.B) {
	router := loadHttpTreeMuxSingle("GET", fiveColon, httpTreeMuxHandler)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkMacaron_Param5(b *testing.B) {
	router := loadMacaronSingle("GET", fiveColon, macaronHandler)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkPat_Param5(b *testing.B) {
	router := loadPatSingle("GET", fiveColon, http.HandlerFunc(httpHandlerFunc))

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}

// func BenchmarkRevel_Param5(b *testing.B) {
// 	router := loadRevelSingle("GET", fiveColon, "RevelController.Handle")

// 	r, _ := http.NewRequest("GET", fiveRoute, nil)
// 	benchRequest(b, router, r)
// }

// Route with 20 Params (no write)
const twentyColon = "/:a/:b/:c/:d/:e/:f/:g/:h/:i/:j/:k/:l/:m/:n/:o/:p/:q/:r/:s/:t"
const twentyBrace = "/{a}/{b}/{c}/{d}/{e}/{f}/{g}/{h}/{i}/{j}/{k}/{l}/{m}/{n}/{o}/{p}/{q}/{r}/{s}/{t}"
const twentyRoute = "/a/b/c/d/e/f/g/h/i/j/k/l/m/n/o/p/q/r/s/t"

func BenchmarkBeego_Param20(b *testing.B) {
	router := loadBeegoSingle("GET", twentyColon, beegoHandler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkChi_Param20(b *testing.B) {
	router := loadChiSingle("GET", twentyBrace, httpHandlerFunc)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkEcho_Param20(b *testing.B) {
	router := loadEchoSingle("GET", twentyColon, echoHandler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGin_Param20(b *testing.B) {
	router := loadGinSingle("GET", twentyColon, ginHandle)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGoRestful_Param20(b *testing.B) {
	handler := loadGoRestfulSingle("GET", twentyBrace, goRestfulHandler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, handler, r)
}
func BenchmarkGorillaMux_Param20(b *testing.B) {
	router := loadGorillaMuxSingle("GET", twentyBrace, httpHandlerFunc)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGowwwRouter_Param20(b *testing.B) {
	router := loadGowwwRouterSingle("GET", twentyColon, http.HandlerFunc(httpHandlerFunc))

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpRouter_Param20(b *testing.B) {
	router := loadHttpRouterSingle("GET", twentyColon, httpRouterHandle)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpTreeMux_Param20(b *testing.B) {
	router := loadHttpTreeMuxSingle("GET", twentyColon, httpTreeMuxHandler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkMacaron_Param20(b *testing.B) {
	router := loadMacaronSingle("GET", twentyColon, macaronHandler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkPat_Param20(b *testing.B) {
	router := loadPatSingle("GET", twentyColon, http.HandlerFunc(httpHandlerFunc))

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}

// func BenchmarkRevel_Param20(b *testing.B) {
// 	router := loadRevelSingle("GET", twentyColon, "RevelController.Handle")

// 	r, _ := http.NewRequest("GET", twentyRoute, nil)
// 	benchRequest(b, router, r)
// }

// Route with Param and write
func BenchmarkBeego_ParamWrite(b *testing.B) {
	router := loadBeegoSingle("GET", "/user/:name", beegoHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkChi_ParamWrite(b *testing.B) {
	router := loadChiSingle("GET", "/user/{name}", chiHandleWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkEcho_ParamWrite(b *testing.B) {
	router := loadEchoSingle("GET", "/user/:name", echoHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGin_ParamWrite(b *testing.B) {
	router := loadGinSingle("GET", "/user/:name", ginHandleWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGoRestful_ParamWrite(b *testing.B) {
	handler := loadGoRestfulSingle("GET", "/user/{name}", goRestfulHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, handler, r)
}
func BenchmarkGorillaMux_ParamWrite(b *testing.B) {
	router := loadGorillaMuxSingle("GET", "/user/{name}", gorillaHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGowwwRouter_ParamWrite(b *testing.B) {
	router := loadGowwwRouterSingle("GET", "/user/:name", http.HandlerFunc(gowwwRouterHandleWrite))

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpRouter_ParamWrite(b *testing.B) {
	router := loadHttpRouterSingle("GET", "/user/:name", httpRouterHandleWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpTreeMux_ParamWrite(b *testing.B) {
	router := loadHttpTreeMuxSingle("GET", "/user/:name", httpTreeMuxHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkMacaron_ParamWrite(b *testing.B) {
	router := loadMacaronSingle("GET", "/user/:name", macaronHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkPat_ParamWrite(b *testing.B) {
	router := loadPatSingle("GET", "/user/:name", http.HandlerFunc(patHandlerWrite))

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}

// func BenchmarkRevel_ParamWrite(b *testing.B) {
// 	router := loadRevelSingle("GET", "/user/:name", "RevelController.HandleWrite")

// 	r, _ := http.NewRequest("GET", "/user/gordon", nil)
// 	benchRequest(b, router, r)
// }
