package router

import "fmt"

type Registry map[string]Builder

var routerRegistry = make(Registry)

func Register(builder Builder) {
	if builder == nil {
		panic("Given builder is nil")
	}

	if _, exist := routerRegistry[builder.Name()]; exist {
		panic(fmt.Sprint("Already registered builder: ", builder.Name()))
	}

	routerRegistry[builder.Name()] = builder
}

func GetRegistry() Registry {
	registry := make(Registry, len(routerRegistry))

	for k, v := range routerRegistry {
		registry[k] = v
	}

	return registry
}
