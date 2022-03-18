package gorestful

import (
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/emicklei/go-restful/v3"
	"github.com/lvjp/go-http-routing-benchmark/router"
)

func init() {
	router.Register(&goRestfulBuilder{})
}

type goRestfulBuilder struct {
}

func (g *goRestfulBuilder) Name() string {
	return "GoRestful"
}

func (g *goRestfulBuilder) ParamType() router.ParamType {
	return router.ParamBraceType
}

func (g *goRestfulBuilder) Build(routes []router.Route, mode router.Mode) http.Handler {
	h := getHandler(mode)
	re := regexp.MustCompile(":([^/]*)")
	wsContainer := restful.NewContainer()
	ws := new(restful.WebService)

	for _, route := range routes {
		path := re.ReplaceAllString(route.Path, "{$1}")

		switch route.Method {
		case "GET":
			ws.Route(ws.GET(path).To(h))
		case "POST":
			ws.Route(ws.POST(path).To(h))
		case "PUT":
			ws.Route(ws.PUT(path).To(h))
		case "PATCH":
			ws.Route(ws.PATCH(path).To(h))
		case "DELETE":
			ws.Route(ws.DELETE(path).To(h))
		default:
			panic("Unknow HTTP method: " + route.Method)
		}
	}

	wsContainer.Add(ws)

	return wsContainer
}

func getHandler(mode router.Mode) restful.RouteFunction {
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

func skipDataModeHandler(r *restful.Request, w *restful.Response) {}

func writeParameterModeHandler(r *restful.Request, w *restful.Response) {
	_, _ = io.WriteString(w, r.PathParameter("name"))
}

func writePathModeHandler(r *restful.Request, w *restful.Response) {
	_, _ = io.WriteString(w, r.Request.RequestURI)
}
