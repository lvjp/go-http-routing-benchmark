package bear

import (
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/lvjp/go-http-routing-benchmark/router"
	"github.com/ursiform/bear"
)

func init() {
	router.Register(&bearBuilder{})
}

type bearBuilder struct {
}

func (b *bearBuilder) Name() string {
	return "Bear"
}

func (b *bearBuilder) ParamType() router.ParamType {
	return router.ParamBraceType
}

func (b *bearBuilder) Build(routes []router.Route, mode router.Mode) http.Handler {
	h := getHandlers(mode)

	router := bear.New()
	re := regexp.MustCompile(":([^/]*)")
	for _, route := range routes {
		switch route.Method {
		case "GET", "POST", "PUT", "PATCH", "DELETE":
			router.On(route.Method, re.ReplaceAllString(route.Path, "{$1}"), h)
		default:
			panic("Unknown HTTP method: " + route.Method)
		}
	}
	return router
}

func (b *bearBuilder) BuildSingle(method string, path string, mode router.Mode) http.Handler {
	h := getHandlers(mode)
	router := bear.New()
	switch method {
	case "GET", "POST", "PUT", "PATCH", "DELETE":
		router.On(method, path, h)
	default:
		panic("Unknown HTTP method: " + method)
	}
	return router
}

func getHandlers(mode router.Mode) interface{} {
	switch mode {
	case router.SkipDataMode:
		return skipDataModeHandler
	case router.WritePathMode:
		return writePathModeHandler
	case router.WriteParameterMode:
		return writeParameterMode
	default:
		panic(fmt.Sprint("unknow mode:", mode))
	}
}

func skipDataModeHandler(_ http.ResponseWriter, _ *http.Request, _ *bear.Context) {}

func writeParameterMode(w http.ResponseWriter, _ *http.Request, ctx *bear.Context) {
	io.WriteString(w, ctx.Params["name"])
}

func writePathModeHandler(w http.ResponseWriter, r *http.Request, _ *bear.Context) {
	io.WriteString(w, r.RequestURI)
}
