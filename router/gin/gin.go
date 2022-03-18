package gin

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lvjp/go-http-routing-benchmark/router"
)

func init() {
	gin.SetMode(gin.ReleaseMode)

	router.Register(&ginBuilder{})
}

type ginBuilder struct {
}

func (g *ginBuilder) Name() string {
	return "Gin"
}

func (g *ginBuilder) ParamType() router.ParamType {
	return router.ParamColonType
}

func (g *ginBuilder) Build(routes []router.Route, mode router.Mode) http.Handler {
	h := getHandler(mode)
	app := gin.New()

	for _, route := range routes {
		app.Handle(route.Method, route.Path, h)
	}

	return app
}

func getHandler(mode router.Mode) gin.HandlerFunc {
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

func skipDataModeHandler(_ *gin.Context) {}

func writeParameterModeHandler(c *gin.Context) {
	_, _ = io.WriteString(c.Writer, c.Params.ByName("name"))
}

func writePathModeHandler(c *gin.Context) {
	_, _ = io.WriteString(c.Writer, c.Request.RequestURI)
}
