package bootstrap

import (
	"contacts-go/api/middlewares"
	"contacts-go/api/routes"
	"contacts-go/lib"
)

func RunServer() func(lib.Env, lib.RequestHandler, routes.Routes, middlewares.Middlewares, lib.Database) {
	return func(
		env lib.Env,
		router lib.RequestHandler,
		routes routes.Routes,
		middlewares middlewares.Middlewares,
		database lib.Database,
	) {
		routes.Setup()
		middlewares.Setup()

		if env.PORT == "" {
			_ = router.Gin.Run()
		} else {
			_ = router.Gin.Run(":" + env.PORT)
		}
	}
}
