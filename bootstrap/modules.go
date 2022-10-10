package bootstrap

import (
	"contacts-go/api/controllers"
	"contacts-go/api/routes"
	"contacts-go/lib"
	"contacts-go/repositories"
	"contacts-go/services"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	lib.Module,
	routes.Module,
	controllers.Module,
	repositories.Module,
	services.Module,
)
