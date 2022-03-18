// Copyright 2014 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	// "github.com/revel/pathtree"
	// "github.com/revel/revel"
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

func init() {
	// makes logging 'webscale' (ignores them)
	log.SetOutput(new(mockResponseWriter))

	// initRevel()
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
