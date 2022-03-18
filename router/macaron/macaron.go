package macaron

import (
	"fmt"
	"net/http"

	"github.com/lvjp/go-http-routing-benchmark/router"
	"gopkg.in/macaron.v1"
)

func init() {
	router.Register(&macaronBuilder{})
}

type macaronBuilder struct {
}

func (m *macaronBuilder) Name() string {
	return "Macaron"
}

func (m *macaronBuilder) ParamType() router.ParamType {
	return router.ParamColonType
}

func (m *macaronBuilder) Build(routes []router.Route, mode router.Mode) http.Handler {
	h := getHandler(mode)
	app := macaron.New()

	for _, route := range routes {
		app.Handle(route.Method, route.Path, h)
	}

	return app
}

func (m *macaronBuilder) BuildSingle(method string, path string, mode router.Mode) http.Handler {
	h := getHandler(mode)
	app := macaron.New()
	app.Handle(method, path, h)
	return app
}

func getHandler(mode router.Mode) []macaron.Handler {
	switch mode {
	case router.SkipDataMode:
		return []macaron.Handler{skipDataModeHandler}
	case router.WritePathMode:
		return []macaron.Handler{writePathModeHandler}
	case router.WriteParameterMode:
		return []macaron.Handler{writeParameterModeHandler}
	default:
		panic(fmt.Sprint("unknow mode:", mode))
	}
}

func skipDataModeHandler() {}

func writeParameterModeHandler(c *macaron.Context) string {
	return c.Params("name")
}

func writePathModeHandler(c *macaron.Context) string {
	return c.Req.RequestURI
}
