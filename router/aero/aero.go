package aero

import (
	"fmt"
	"io"
	"net/http"

	"github.com/aerogo/aero"
	"github.com/lvjp/go-http-routing-benchmark/router"
)

func init() {
	router.Register(&aeroBuilder{})
}

type aeroBuilder struct {
}

func (a *aeroBuilder) Name() string {
	return "Aero"
}

func (a *aeroBuilder) ParamType() router.ParamType {
	return router.ParamBraceType
}

func (a *aeroBuilder) Build(routes []router.Route, mode router.Mode) http.Handler {
	h := getHandler(mode)
	app := aero.New()

	for _, r := range routes {
		switch r.Method {
		case "GET":
			app.Get(r.Path, h)
		case "POST":
			app.Post(r.Path, h)
		case "PUT":
			app.Put(r.Path, h)
		case "PATCH":
			app.Router().Add(http.MethodPatch, r.Path, h)
		case "DELETE":
			app.Delete(r.Path, h)
		default:
			panic("Unknow HTTP method: " + r.Method)
		}
	}

	return app
}

func (a *aeroBuilder) BuildSingle(method string, path string, mode router.Mode) http.Handler {
	h := getHandler(mode)
	app := aero.New()
	switch method {
	case "GET":
		app.Get(path, h)
	case "POST":
		app.Post(path, h)
	case "PUT":
		app.Put(path, h)
	case "PATCH":
		app.Router().Add(http.MethodPatch, path, h)
	case "DELETE":
		app.Delete(path, h)
	default:
		panic("Unknow HTTP method: " + method)
	}

	return app
}

func getHandler(mode router.Mode) aero.Handler {
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

func skipDataModeHandler(c aero.Context) error {
	return nil
}

func writeParameterMode(ctx aero.Context) error {
	io.WriteString(ctx.Response().Internal(), ctx.Get("name"))
	return nil
}
func writePathModeHandler(ctx aero.Context) error {
	io.WriteString(ctx.Response().Internal(), ctx.Request().Path())
	return nil
}
