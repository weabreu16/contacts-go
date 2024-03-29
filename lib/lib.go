package lib

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewEnv),
	fx.Provide(NewRequestHandler),
	fx.Provide(NewDatabase),
	fx.Provide(NewRepository),
	fx.Provide(NewFilesBucket),
)
