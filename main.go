package main

import (
	_ "github.com/lvjp/go-http-routing-benchmark/router/aero"
	_ "github.com/lvjp/go-http-routing-benchmark/router/beego"
	_ "github.com/lvjp/go-http-routing-benchmark/router/chi"
	_ "github.com/lvjp/go-http-routing-benchmark/router/echo"
	_ "github.com/lvjp/go-http-routing-benchmark/router/gin"
	_ "github.com/lvjp/go-http-routing-benchmark/router/gorestful"
	_ "github.com/lvjp/go-http-routing-benchmark/router/gorilla"
	_ "github.com/lvjp/go-http-routing-benchmark/router/gowww"
	_ "github.com/lvjp/go-http-routing-benchmark/router/httprouter"
	_ "github.com/lvjp/go-http-routing-benchmark/router/httptreemux"
	_ "github.com/lvjp/go-http-routing-benchmark/router/macaron"
	_ "github.com/lvjp/go-http-routing-benchmark/router/pat"
)
