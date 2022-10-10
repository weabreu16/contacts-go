package services

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewContactService),
	fx.Provide(NewJWTService),
	fx.Provide(NewAuthService),
)
