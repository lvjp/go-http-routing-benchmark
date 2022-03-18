package echo

import (
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lvjp/go-http-routing-benchmark/router"
)

func init() {
	router.Register(&echoBuilder{})
}

type echoBuilder struct {
}

func (e *echoBuilder) Name() string {
	return "Echo"
}

func (e *echoBuilder) ParamType() router.ParamType {
	return router.ParamColonType
}

func (e *echoBuilder) Build(routes []router.Route, mode router.Mode) http.Handler {
	h := getHandler(mode)
	app := echo.New()

	for _, r := range routes {
		switch r.Method {
		case "GET":
			app.GET(r.Path, h)
		case "POST":
			app.POST(r.Path, h)
		case "PUT":
			app.PUT(r.Path, h)
		case "PATCH":
			app.PATCH(r.Path, h)
		case "DELETE":
			app.DELETE(r.Path, h)
		default:
			panic("Unknow HTTP method: " + r.Method)
		}
	}

	return app
}

func (e *echoBuilder) BuildSingle(method string, path string, mode router.Mode) http.Handler {
	h := getHandler(mode)
	app := echo.New()

	switch method {
	case "GET":
		app.GET(path, h)
	case "POST":
		app.POST(path, h)
	case "PUT":
		app.PUT(path, h)
	case "PATCH":
		app.PATCH(path, h)
	case "DELETE":
		app.DELETE(path, h)
	default:
		panic("Unknow HTTP method: " + method)
	}

	return app
}

func getHandler(mode router.Mode) echo.HandlerFunc {
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

func skipDataModeHandler(c echo.Context) error { return nil }

func writeParameterModeHandler(c echo.Context) error {
	_, _ = io.WriteString(c.Response(), c.Param("name"))
	return nil
}

func writePathModeHandler(c echo.Context) error {
	_, _ = io.WriteString(c.Response(), c.Request().RequestURI)
	return nil
}
