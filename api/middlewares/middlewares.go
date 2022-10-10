package middlewares

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewMiddlewares),
)

type Middleware interface {
	Setup()
}

type Middlewares []Middleware

func NewMiddlewares() Middlewares {
	return Middlewares{}
}

func (self Middlewares) Setup() {
	for _, middleware := range self {
		middleware.Setup()
	}
}
