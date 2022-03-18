package beego

import (
	"fmt"
	"net/http"
	"regexp"
	"runtime"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/lvjp/go-http-routing-benchmark/router"
)

func init() {
	// beego sets it to runtime.NumCPU()
	// Currently none of the contesters does concurrent routing
	runtime.GOMAXPROCS(1)

	beego.BConfig.RunMode = beego.PROD
	beego.BeeLogger.Close()

	router.Register(&beegoBuilder{})
}

type beegoBuilder struct {
}

func (b *beegoBuilder) Name() string {
	return "Beego"
}

func (b *beegoBuilder) ParamType() router.ParamType {
	return router.ParamColonType
}

func (b *beegoBuilder) Build(routes []router.Route, mode router.Mode) http.Handler {
	h := getHandler(mode)
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

func (b *beegoBuilder) BuildSingle(method string, path string, mode router.Mode) http.Handler {
	h := getHandler(mode)
	app := beego.NewControllerRegister()

	switch method {
	case "GET":
		app.Get(path, h)
	case "POST":
		app.Post(path, h)
	case "PUT":
		app.Put(path, h)
	case "PATCH":
		app.Patch(path, h)
	case "DELETE":
		app.Delete(path, h)
	default:
		panic("Unknow HTTP method: " + method)
	}
	return app
}

func getHandler(mode router.Mode) beego.FilterFunc {
	switch mode {
	case router.SkipDataMode:
		return skipDataModeHandler
	case router.WritePathMode:
		return writePathModeHandler
	case router.WriteParameterMode:
		return writeParameterModeHandler
	default:
		panic(fmt.Sprint("unknow mode:", mode))
	}
}

func skipDataModeHandler(ctx *context.Context) {}

func writeParameterModeHandler(ctx *context.Context) {
	ctx.WriteString(ctx.Input.Param(":name"))
}

func writePathModeHandler(ctx *context.Context) {
	ctx.WriteString(ctx.Request.RequestURI)
}
