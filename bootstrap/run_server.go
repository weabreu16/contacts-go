package bootstrap

import (
	"contacts-go/api/routes"
	"contacts-go/lib"
)

func RunServer() func(lib.Env, lib.RequestHandler, routes.Routes) {
	return func(
		env lib.Env,
		router lib.RequestHandler,
		routes routes.Routes,
	) {
		routes.Setup()

		if env.PORT == "" {
			_ = router.Gin.Run()
		} else {
			_ = router.Gin.Run(":" + env.PORT)
		}
	}
}
