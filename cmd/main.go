package main

import (
	"context"
	"github.com/NikitaVi/minifier-sso/internal/app"
	"github.com/NikitaVi/minifier-sso/internal/logger"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		panic(err)
	}
	logger.Init()

	err = a.Run(ctx)
	if err != nil {
		panic(err)
	}
}
