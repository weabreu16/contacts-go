package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewContactRoutes),
	fx.Provide(NewRoutes),
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(
	contactRoutes ContactRoutes,
) Routes {
	return Routes{
		contactRoutes,
	}
}

func (self Routes) Setup() {
	for _, route := range self {
		route.Setup()
	}
}
