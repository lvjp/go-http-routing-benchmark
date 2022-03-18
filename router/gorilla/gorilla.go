package gorilla

import (
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
	"github.com/lvjp/go-http-routing-benchmark/router"
)

func init() {
	router.Register(&gorillaBuilder{})
}

type gorillaBuilder struct {
}

func (g *gorillaBuilder) Name() string {
	return "GorillaMux"
}

func (g *gorillaBuilder) ParamType() router.ParamType {
	return router.ParamBraceType
}

func (g *gorillaBuilder) Build(routes []router.Route, mode router.Mode) http.Handler {
	h := getHandler(mode)
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

func (g *gorillaBuilder) BuildSingle(method string, path string, mode router.Mode) http.Handler {
	h := getHandler(mode)
	m := mux.NewRouter()
	m.HandleFunc(path, h).Methods(method)
	return m
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
	params := mux.Vars(r)
	_, _ = io.WriteString(w, params["name"])
}

func writePathModeHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, r.RequestURI)
}
