package bootstrap

import (
	"context"

	"log"

	"go.uber.org/fx"
)

func Bootstrap() {
	opts := fx.Options(
		fx.Invoke(RunServer()),
	)

	ctx := context.Background()

	app := fx.New(CommonModules, opts)

	err := app.Start(ctx)

	defer app.Stop(ctx)

	if err != nil {
		log.Fatal(err)
	}
}
