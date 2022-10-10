package bootstrap

import (
	"contacts-go/api/routes"
	"contacts-go/lib"
)

func RunServer() func(lib.Env, lib.RequestHandler, routes.Routes, lib.Database) {
	return func(
		env lib.Env,
		router lib.RequestHandler,
		routes routes.Routes,
		database lib.Database,
	) {
		routes.Setup()

		if env.PORT == "" {
			_ = router.Gin.Run()
		} else {
			_ = router.Gin.Run(":" + env.PORT)
		}
	}
}
