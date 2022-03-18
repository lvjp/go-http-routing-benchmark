package gowww

import (
	"fmt"
	"io"
	"net/http"

	gowww "github.com/gowww/router"
	"github.com/lvjp/go-http-routing-benchmark/router"
)

func init() {
	router.Register(&gowwwBuilder{})
}

type gowwwBuilder struct {
}

func (g *gowwwBuilder) Name() string {
	return "Gowww"
}

func (g *gowwwBuilder) ParamType() router.ParamType {
	return router.ParamColonType
}

func (g *gowwwBuilder) Build(routes []router.Route, mode router.Mode) http.Handler {
	h := getHandler(mode)
	router := gowww.New()

	for _, route := range routes {
		router.Handle(route.Method, route.Path, http.HandlerFunc(h))
	}

	return router
}

func (g *gowwwBuilder) BuildSingle(method string, path string, mode router.Mode) http.Handler {
	h := getHandler(mode)
	router := gowww.New()
	router.Handle(method, path, h)
	return router
}

func getHandler(mode router.Mode) http.HandlerFunc {
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

func skipDataModeHandler(_ http.ResponseWriter, _ *http.Request) {}

func writeParameterModeHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, gowww.Parameter(r, "name"))
}

func writePathModeHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, r.RequestURI)
}
