package repositories

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewContactRepository),
	fx.Provide(NewUserRepository),
)
