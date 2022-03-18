package httprouter

import (
	"fmt"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lvjp/go-http-routing-benchmark/router"
)

func init() {
	router.Register(&httpRouterBuilder{})
}

type httpRouterBuilder struct {
}

func (hr *httpRouterBuilder) Name() string {
	return "HttpRouter"
}

func (hr *httpRouterBuilder) ParamType() router.ParamType {
	return router.ParamColonType
}

func (hr *httpRouterBuilder) Build(routes []router.Route, mode router.Mode) http.Handler {
	h := getHandler(mode)
	app := httprouter.New()

	for _, route := range routes {
		app.Handle(route.Method, route.Path, h)
	}

	return app
}

func getHandler(mode router.Mode) httprouter.Handle {
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

func skipDataModeHandler(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params) {}

func writeParameterModeHandler(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	_, _ = io.WriteString(w, ps.ByName("name"))
}

func writePathModeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, _ = io.WriteString(w, r.RequestURI)
}
