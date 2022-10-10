package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewContactRoutes),
	fx.Provide(NewAuthRoutes),
	fx.Provide(NewRoutes),
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(
	contactRoutes ContactRoutes,
	authRoutes AuthRoutes,
) Routes {
	return Routes{
		contactRoutes,
		authRoutes,
	}
}

func (self Routes) Setup() {
	for _, route := range self {
		route.Setup()
	}
}
