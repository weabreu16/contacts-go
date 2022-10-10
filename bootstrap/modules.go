package bootstrap

import (
	"contacts-go/lib"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	lib.Module,
)
