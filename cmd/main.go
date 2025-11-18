package main

import (
	"context"
	"github.com/NikitaVi/minifier-sso/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		panic(err)
	}

	err = a.Run(ctx)
	if err != nil {
		panic(err)
	}
}
