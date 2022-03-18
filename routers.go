// Copyright 2014 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"

	// If you add new routers please:
	// - Keep the benchmark functions etc. alphabetically sorted
	// - Make a pull request (without benchmark results) at
	//   https://github.com/lvjp/go-http-routing-benchmark
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/bmizerany/pat"
	"github.com/lvjp/go-http-routing-benchmark/router"

	"github.com/dimfeld/httptreemux"
	"github.com/emicklei/go-restful"
	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi"
	"github.com/gorilla/mux"
	gowwwrouter "github.com/gowww/router"
	"github.com/julienschmidt/httprouter"
	"github.com/labstack/echo/v4"

	// "github.com/revel/pathtree"
	// "github.com/revel/revel"

	"gopkg.in/macaron.v1"
)

type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() (h http.Header) {
	return http.Header{}
}

func (m *mockResponseWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockResponseWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *mockResponseWriter) WriteHeader(int) {}

var nullLogger *log.Logger

// flag indicating if the normal or the test handler should be loaded
var loadTestHandler = false

func init() {
	// beego sets it to runtime.NumCPU()
	// Currently none of the contesters does concurrent routing
	runtime.GOMAXPROCS(1)

	// makes logging 'webscale' (ignores them)
	log.SetOutput(new(mockResponseWriter))
	nullLogger = log.New(new(mockResponseWriter), "", 0)

	initBeego()
	initGin()
	// initRevel()
}

// Common
func httpHandlerFunc(_ http.ResponseWriter, _ *http.Request) {}

func httpHandlerFuncTest(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.RequestURI)
}

// beego
func beegoHandler(ctx *context.Context) {}

func beegoHandlerWrite(ctx *context.Context) {
	ctx.WriteString(ctx.Input.Param(":name"))
}

func beegoHandlerTest(ctx *context.Context) {
	ctx.WriteString(ctx.Request.RequestURI)
}

func initBeego() {
	beego.BConfig.RunMode = beego.PROD
	beego.BeeLogger.Close()
}

func loadBeego(routes []router.Route) http.Handler {
	h := beegoHandler
	if loadTestHandler {
		h = beegoHandlerTest
	}

	re := regexp.MustCompile(":([^/]*)")
	app := beego.NewControllerRegister()
	for _, route := range routes {
		route.Path = re.ReplaceAllString(route.Path, ":$1")
		switch route.Method {
		case "GET":
			app.Get(route.Path, h)
		case "POST":
			app.Post(route.Path, h)
		case "PUT":
			app.Put(route.Path, h)
		case "PATCH":
			app.Patch(route.Path, h)
		case "DELETE":
			app.Delete(route.Path, h)
		default:
			panic("Unknow HTTP method: " + route.Method)
		}
	}
	return app
}

func loadBeegoSingle(method, path string, handler beego.FilterFunc) http.Handler {
	app := beego.NewControllerRegister()
	switch method {
	case "GET":
		app.Get(path, handler)
	case "POST":
		app.Post(path, handler)
	case "PUT":
		app.Put(path, handler)
	case "PATCH":
		app.Patch(path, handler)
	case "DELETE":
		app.Delete(path, handler)
	default:
		panic("Unknow HTTP method: " + method)
	}
	return app
}

// chi
// chi
func chiHandleWrite(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, chi.URLParam(r, "name"))
}

func loadChi(routes []router.Route) http.Handler {
	h := httpHandlerFunc
	if loadTestHandler {
		h = httpHandlerFuncTest
	}

	re := regexp.MustCompile(":([^/]*)")

	mux := chi.NewRouter()
	for _, route := range routes {
		path := re.ReplaceAllString(route.Path, "{$1}")

		switch route.Method {
		case "GET":
			mux.Get(path, h)
		case "POST":
			mux.Post(path, h)
		case "PUT":
			mux.Put(path, h)
		case "PATCH":
			mux.Patch(path, h)
		case "DELETE":
			mux.Delete(path, h)
		default:
			panic("Unknown HTTP method: " + route.Method)
		}
	}
	return mux
}

func loadChiSingle(method, path string, handler http.HandlerFunc) http.Handler {
	mux := chi.NewRouter()
	switch method {
	case "GET":
		mux.Get(path, handler)
	case "POST":
		mux.Post(path, handler)
	case "PUT":
		mux.Put(path, handler)
	case "PATCH":
		mux.Patch(path, handler)
	case "DELETE":
		mux.Delete(path, handler)
	default:
		panic("Unknown HTTP method: " + method)
	}
	return mux
}

// Echo
func echoHandler(c echo.Context) error {
	return nil
}

func echoHandlerWrite(c echo.Context) error {
	io.WriteString(c.Response(), c.Param("name"))
	return nil
}

func echoHandlerTest(c echo.Context) error {
	io.WriteString(c.Response(), c.Request().RequestURI)
	return nil
}

func loadEcho(routes []router.Route) http.Handler {
	var h echo.HandlerFunc = echoHandler
	if loadTestHandler {
		h = echoHandlerTest
	}

	e := echo.New()
	for _, r := range routes {
		switch r.Method {
		case "GET":
			e.GET(r.Path, h)
		case "POST":
			e.POST(r.Path, h)
		case "PUT":
			e.PUT(r.Path, h)
		case "PATCH":
			e.PATCH(r.Path, h)
		case "DELETE":
			e.DELETE(r.Path, h)
		default:
			panic("Unknow HTTP method: " + r.Method)
		}
	}
	return e
}

func loadEchoSingle(method, path string, h echo.HandlerFunc) http.Handler {
	e := echo.New()
	switch method {
	case "GET":
		e.GET(path, h)
	case "POST":
		e.POST(path, h)
	case "PUT":
		e.PUT(path, h)
	case "PATCH":
		e.PATCH(path, h)
	case "DELETE":
		e.DELETE(path, h)
	default:
		panic("Unknow HTTP method: " + method)
	}
	return e
}

// Gin
func ginHandle(_ *gin.Context) {}

func ginHandleWrite(c *gin.Context) {
	io.WriteString(c.Writer, c.Params.ByName("name"))
}

func ginHandleTest(c *gin.Context) {
	io.WriteString(c.Writer, c.Request.RequestURI)
}

func initGin() {
	gin.SetMode(gin.ReleaseMode)
}

func loadGin(routes []router.Route) http.Handler {
	h := ginHandle
	if loadTestHandler {
		h = ginHandleTest
	}

	router := gin.New()
	for _, route := range routes {
		router.Handle(route.Method, route.Path, h)
	}
	return router
}

func loadGinSingle(method, path string, handle gin.HandlerFunc) http.Handler {
	router := gin.New()
	router.Handle(method, path, handle)
	return router
}

// go-restful
func goRestfulHandler(r *restful.Request, w *restful.Response) {}

func goRestfulHandlerWrite(r *restful.Request, w *restful.Response) {
	io.WriteString(w, r.PathParameter("name"))
}

func goRestfulHandlerTest(r *restful.Request, w *restful.Response) {
	io.WriteString(w, r.Request.RequestURI)
}

func loadGoRestful(routes []router.Route) http.Handler {
	h := goRestfulHandler
	if loadTestHandler {
		h = goRestfulHandlerTest
	}

	re := regexp.MustCompile(":([^/]*)")

	wsContainer := restful.NewContainer()
	ws := new(restful.WebService)

	for _, route := range routes {
		path := re.ReplaceAllString(route.Path, "{$1}")

		switch route.Method {
		case "GET":
			ws.Route(ws.GET(path).To(h))
		case "POST":
			ws.Route(ws.POST(path).To(h))
		case "PUT":
			ws.Route(ws.PUT(path).To(h))
		case "PATCH":
			ws.Route(ws.PATCH(path).To(h))
		case "DELETE":
			ws.Route(ws.DELETE(path).To(h))
		default:
			panic("Unknow HTTP method: " + route.Method)
		}
	}
	wsContainer.Add(ws)
	return wsContainer
}

func loadGoRestfulSingle(method, path string, handler restful.RouteFunction) http.Handler {
	wsContainer := restful.NewContainer()
	ws := new(restful.WebService)
	switch method {
	case "GET":
		ws.Route(ws.GET(path).To(handler))
	case "POST":
		ws.Route(ws.POST(path).To(handler))
	case "PUT":
		ws.Route(ws.PUT(path).To(handler))
	case "PATCH":
		ws.Route(ws.PATCH(path).To(handler))
	case "DELETE":
		ws.Route(ws.DELETE(path).To(handler))
	default:
		panic("Unknow HTTP method: " + method)
	}
	wsContainer.Add(ws)
	return wsContainer
}

// gorilla/mux
func gorillaHandlerWrite(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	io.WriteString(w, params["name"])
}

func loadGorillaMux(routes []router.Route) http.Handler {
	h := httpHandlerFunc
	if loadTestHandler {
		h = httpHandlerFuncTest
	}

	re := regexp.MustCompile(":([^/]*)")
	m := mux.NewRouter()
	for _, route := range routes {
		m.HandleFunc(
			re.ReplaceAllString(route.Path, "{$1}"),
			h,
		).Methods(route.Method)
	}
	return m
}

func loadGorillaMuxSingle(method, path string, handler http.HandlerFunc) http.Handler {
	m := mux.NewRouter()
	m.HandleFunc(path, handler).Methods(method)
	return m
}

// gowww/router
func gowwwRouterHandleWrite(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, gowwwrouter.Parameter(r, "name"))
}

func loadGowwwRouter(routes []router.Route) http.Handler {
	h := httpHandlerFunc
	if loadTestHandler {
		h = httpHandlerFuncTest
	}

	router := gowwwrouter.New()
	for _, route := range routes {
		router.Handle(route.Method, route.Path, http.HandlerFunc(h))
	}
	return router
}

func loadGowwwRouterSingle(method, path string, handler http.Handler) http.Handler {
	router := gowwwrouter.New()
	router.Handle(method, path, handler)
	return router
}

// HttpRouter
func httpRouterHandle(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params) {}

func httpRouterHandleWrite(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	io.WriteString(w, ps.ByName("name"))
}

func httpRouterHandleTest(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	io.WriteString(w, r.RequestURI)
}

func loadHttpRouter(routes []router.Route) http.Handler {
	h := httpRouterHandle
	if loadTestHandler {
		h = httpRouterHandleTest
	}

	router := httprouter.New()
	for _, route := range routes {
		router.Handle(route.Method, route.Path, h)
	}
	return router
}

func loadHttpRouterSingle(method, path string, handle httprouter.Handle) http.Handler {
	router := httprouter.New()
	router.Handle(method, path, handle)
	return router
}

// httpTreeMux
func httpTreeMuxHandler(_ http.ResponseWriter, _ *http.Request, _ map[string]string) {}

func httpTreeMuxHandlerWrite(w http.ResponseWriter, _ *http.Request, vars map[string]string) {
	io.WriteString(w, vars["name"])
}

func httpTreeMuxHandlerTest(w http.ResponseWriter, r *http.Request, _ map[string]string) {
	io.WriteString(w, r.RequestURI)
}

func loadHttpTreeMux(routes []router.Route) http.Handler {
	h := httpTreeMuxHandler
	if loadTestHandler {
		h = httpTreeMuxHandlerTest
	}

	router := httptreemux.New()
	for _, route := range routes {
		router.Handle(route.Method, route.Path, h)
	}
	return router
}

func loadHttpTreeMuxSingle(method, path string, handler httptreemux.HandlerFunc) http.Handler {
	router := httptreemux.New()
	router.Handle(method, path, handler)
	return router
}

// Macaron
func macaronHandler() {}

func macaronHandlerWrite(c *macaron.Context) string {
	return c.Params("name")
}

func macaronHandlerTest(c *macaron.Context) string {
	return c.Req.RequestURI
}

func loadMacaron(routes []router.Route) http.Handler {
	var h = []macaron.Handler{macaronHandler}
	if loadTestHandler {
		h[0] = macaronHandlerTest
	}

	m := macaron.New()
	for _, route := range routes {
		m.Handle(route.Method, route.Path, h)
	}
	return m
}

func loadMacaronSingle(method, path string, handler interface{}) http.Handler {
	m := macaron.New()
	m.Handle(method, path, []macaron.Handler{handler})
	return m
}

// pat
func patHandlerWrite(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Query().Get(":name"))
}

func loadPat(routes []router.Route) http.Handler {
	h := http.HandlerFunc(httpHandlerFunc)
	if loadTestHandler {
		h = http.HandlerFunc(httpHandlerFuncTest)
	}

	m := pat.New()
	for _, route := range routes {
		switch route.Method {
		case "GET":
			m.Get(route.Path, h)
		case "POST":
			m.Post(route.Path, h)
		case "PUT":
			m.Put(route.Path, h)
		case "DELETE":
			m.Del(route.Path, h)
		default:
			panic("Unknow HTTP method: " + route.Method)
		}
	}
	return m
}

func loadPatSingle(method, path string, handler http.Handler) http.Handler {
	m := pat.New()
	switch method {
	case "GET":
		m.Get(path, handler)
	case "POST":
		m.Post(path, handler)
	case "PUT":
		m.Put(path, handler)
	case "DELETE":
		m.Del(path, handler)
	default:
		panic("Unknow HTTP method: " + method)
	}
	return m
}

// Revel (Router only)
// In the following code some Revel internals are modeled.
// The original revel code is copyrighted by Rob Figueiredo.
// See https://github.com/revel/revel/blob/master/LICENSE
// type RevelController struct {
// 	*revel.Controller
// 	router *revel.Router
// }

// func (rc *RevelController) Handle() revel.Result {
// 	return revelResult{}
// }

// func (rc *RevelController) HandleWrite() revel.Result {
// 	return rc.RenderText(rc.Params.Get("name"))
// }

// func (rc *RevelController) HandleTest() revel.Result {
// 	return rc.RenderText(rc.Request.GetRequestURI())
// }

// type revelResult struct{}

// func (rr revelResult) Apply(req *revel.Request, resp *revel.Response) {}

// func (rc *RevelController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	// Dirty hacks, do NOT copy!
// 	revel.MainRouter = rc.router

// 	upgrade := r.Header.Get("Upgrade")
// 	if upgrade == "websocket" || upgrade == "Websocket" {
// 		panic("Not implemented")
// 	} else {
// 		var (
// 			req  = revel.NewRequest(r)
// 			resp = revel.NewResponse(w)
// 			c    = revel.NewController(req, resp)
// 		)
// 		req.Websocket = nil
// 		revel.Filters[0](c, revel.Filters[1:])
// 		if c.Result != nil {
// 			c.Result.Apply(req, resp)
// 		} else if c.Response.Status != 0 {
// 			panic("Not implemented")
// 		}
// 		// Close the Writer if we can
// 		if w, ok := resp.Out.(io.Closer); ok {
// 			w.Close()
// 		}
// 	}
// }

// func initRevel() {
// 	// Only use the Revel filters required for this benchmark
// 	revel.Filters = []revel.Filter{
// 		revel.RouterFilter,
// 		revel.ParamsFilter,
// 		revel.ActionInvoker,
// 	}

// 	revel.RegisterController((*RevelController)(nil),
// 		[]*revel.MethodType{
// 			{
// 				Name: "Handle",
// 			},
// 			{
// 				Name: "HandleWrite",
// 			},
// 			{
// 				Name: "HandleTest",
// 			},
// 		})
// }

// func loadRevel(routes []route) http.Handler {
// 	h := "RevelController.Handle"
// 	if loadTestHandler {
// 		h = "RevelController.HandleTest"
// 	}

// 	router := revel.NewRouter("")

// 	// parseRoutes
// 	var rs []*revel.Route
// 	for _, r := range routes {
// 		rs = append(rs, revel.NewRoute(r.method, r.path, h, "", "", 0))
// 	}
// 	router.Routes = rs

// 	// updateTree
// 	router.Tree = pathtree.New()
// 	for _, r := range router.Routes {
// 		err := router.Tree.Add(r.TreePath, r)
// 		// Allow GETs to respond to HEAD requests.
// 		if err == nil && r.Method == "GET" {
// 			err = router.Tree.Add("/HEAD"+r.Path, r)
// 		}
// 		// Error adding a route to the pathtree.
// 		if err != nil {
// 			panic(err)
// 		}
// 	}

// 	rc := new(RevelController)
// 	rc.router = router
// 	return rc
// }

// func loadRevelSingle(method, path, action string) http.Handler {
// 	router := revel.NewRouter("")

// 	route := revel.NewRoute(method, path, action, "", "", 0)
// 	if err := router.Tree.Add(route.TreePath, route); err != nil {
// 		panic(err)
// 	}

// 	rc := new(RevelController)
// 	rc.router = router
// 	return rc
// }

// Usage notice
func main() {
	fmt.Println("Usage: go test -bench=. -timeout=20m")
	os.Exit(1)
}
