package httptreemux

import (
	"fmt"
	"io"
	"net/http"

	"github.com/dimfeld/httptreemux"
	"github.com/lvjp/go-http-routing-benchmark/router"
)

func init() {
	router.Register(&gorillaBuilder{})
}

type gorillaBuilder struct {
}

func (g *gorillaBuilder) Name() string {
	return "HttpTreeMux"
}

func (g *gorillaBuilder) ParamType() router.ParamType {
	return router.ParamColonType
}

func (g *gorillaBuilder) Build(routes []router.Route, mode router.Mode) http.Handler {
	h := getHandler(mode)
	app := httptreemux.New()
	for _, route := range routes {
		app.Handle(route.Method, route.Path, h)
	}
	return app
}

func (g *gorillaBuilder) BuildSingle(method string, path string, mode router.Mode) http.Handler {
	h := getHandler(mode)
	app := httptreemux.New()
	app.Handle(method, path, h)
	return app
}

func getHandler(mode router.Mode) httptreemux.HandlerFunc {
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

func skipDataModeHandler(_ http.ResponseWriter, _ *http.Request, _ map[string]string) {}

func writeParameterModeHandler(w http.ResponseWriter, _ *http.Request, vars map[string]string) {
	_, _ = io.WriteString(w, vars["name"])
}

func writePathModeHandler(w http.ResponseWriter, r *http.Request, _ map[string]string) {
	_, _ = io.WriteString(w, r.RequestURI)
}
