package pat

import (
	"fmt"
	"io"
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/lvjp/go-http-routing-benchmark/router"
)

func init() {
	router.Register(&patBuilder{})
}

type patBuilder struct {
}

func (p *patBuilder) Name() string {
	return "Pat"
}

func (p *patBuilder) ParamType() router.ParamType {
	return router.ParamColonType
}

func (p *patBuilder) Build(routes []router.Route, mode router.Mode) http.Handler {
	h := getHandler(mode)
	app := pat.New()

	for _, route := range routes {
		switch route.Method {
		case "GET":
			app.Get(route.Path, h)
		case "POST":
			app.Post(route.Path, h)
		case "PUT":
			app.Put(route.Path, h)
		case "DELETE":
			app.Del(route.Path, h)
		default:
			panic("Unknow HTTP method: " + route.Method)
		}
	}

	return app
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
	_, _ = io.WriteString(w, r.URL.Query().Get(":name"))
}

func writePathModeHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, r.RequestURI)
}
