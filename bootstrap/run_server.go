package bootstrap

import (
	"contacts-go/lib"
)

func RunServer() func(lib.Env, lib.RequestHandler) {
	return func(
		env lib.Env,
		router lib.RequestHandler,
	) {
		if env.PORT == "" {
			_ = router.Gin.Run()
		} else {
			_ = router.Gin.Run(":" + env.PORT)
		}
	}
}
