package chi

import (
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/go-chi/chi"
	"github.com/lvjp/go-http-routing-benchmark/router"
)

func init() {
	router.Register(&chiBuilder{})
}

type chiBuilder struct {
}

func (c *chiBuilder) Name() string {
	return "Chi"
}

func (c *chiBuilder) ParamType() router.ParamType {
	return router.ParamBraceType
}

func (c *chiBuilder) Build(routes []router.Route, mode router.Mode) http.Handler {
	h := getHandler(mode)
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

func (c *chiBuilder) BuildSingle(method string, path string, mode router.Mode) http.Handler {
	h := getHandler(mode)
	mux := chi.NewRouter()

	switch method {
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
		panic("Unknown HTTP method: " + method)
	}

	return mux
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
	_, _ = io.WriteString(w, chi.URLParam(r, "name"))
}

func writePathModeHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, r.RequestURI)
}
