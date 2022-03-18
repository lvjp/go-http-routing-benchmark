package router

import (
	"net/http"
)

type Route struct {
	Method string
	Path   string
}

type Mode int

const (
	SkipDataMode Mode = iota
	WriteParameterMode
	WritePathMode
)

type ParamType int

const (
	ParamColonType ParamType = iota
	ParamBraceType
)

type Builder interface {
	Name() string
	ParamType() ParamType
	Build(routes []Route, mode Mode) http.Handler
}
