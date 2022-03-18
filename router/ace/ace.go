package ace

import (
	"fmt"
	"io"
	"net/http"

	"github.com/lvjp/go-http-routing-benchmark/router"
	"github.com/plimble/ace"
)

func init() {
	router.Register(&aceBuilder{})
}

type aceBuilder struct {
}

func (a *aceBuilder) Name() string {
	return "Ace"
}

func (a *aceBuilder) ParamType() router.ParamType {
	return router.ParamColonType
}

func (a *aceBuilder) Build(routes []router.Route, mode router.Mode) http.Handler {
	handlers := getHandlers(mode)
	router := ace.New()

	for _, route := range routes {
		router.Handle(route.Method, route.Path, handlers)
	}

	return router
}

func (a *aceBuilder) BuildSingle(method string, path string, mode router.Mode) http.Handler {
	router := ace.New()
	router.Handle(method, path, getHandlers(mode))
	return router
}

func getHandlers(mode router.Mode) []ace.HandlerFunc {
	switch mode {
	case router.SkipDataMode:
		return []ace.HandlerFunc{skipDataModeHandler}
	case router.WritePathMode:
		return []ace.HandlerFunc{writePathModeHandler}
	case router.WriteParameterMode:
		return []ace.HandlerFunc{writeParameterMode}
	default:
		panic(fmt.Sprint("unknow mode:", mode))
	}
}

func skipDataModeHandler(_ *ace.C) {}

func writeParameterMode(c *ace.C) {
	io.WriteString(c.Writer, c.Param("name"))
}

func writePathModeHandler(c *ace.C) {
	io.WriteString(c.Writer, c.Request.RequestURI)
}
