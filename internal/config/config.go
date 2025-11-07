package config

import (
	"context"
	"github.com/joho/godotenv"
)

func Load(ctx context.Context, filename string) error {
	err := godotenv.Load(filename)
	if err != nil {
		return err
	}

	return nil
}
