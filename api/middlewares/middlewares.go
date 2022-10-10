package middlewares

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewCorsMiddleware),
	fx.Provide(NewMiddlewares),
)

type Middleware interface {
	Setup()
}

type Middlewares []Middleware

func NewMiddlewares(
	corsMiddleware CorsMiddleware,
) Middlewares {
	return Middlewares{
		corsMiddleware,
	}
}

func (self Middlewares) Setup() {
	for _, middleware := range self {
		middleware.Setup()
	}
}
