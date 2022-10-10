package middlewares

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewCorsMiddleware),
	fx.Provide(NewJWTMiddleware),
	fx.Provide(NewMiddlewares),
)

type Middleware interface {
	Setup()
}

type Middlewares []Middleware

func NewMiddlewares(
	corsMiddleware CorsMiddleware,
	jwtMiddleware JWTMiddleware,
) Middlewares {
	return Middlewares{
		corsMiddleware,
		jwtMiddleware,
	}
}

func (self Middlewares) Setup() {
	for _, middleware := range self {
		middleware.Setup()
	}
}
