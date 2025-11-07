package config

import (
	"fmt"
	"os"
)

const pgEnvPort = "PG_PORT"
const pgEnvHost = "PG_HOST"
const pgEnvName = "PG_DB_NAME"
const pgEnvUser = "PG_USER"
const pgEnvPassword = "PG_PASSWORD"

type pgConfig struct {
	dsn string
}

type PGConfig interface {
	DSN() string
}

func NewPGConfig() (PGConfig, error) {
	host := os.Getenv(pgEnvHost)
	if host == "" {
		return nil, fmt.Errorf("environment variable %s must be set", pgEnvHost)
	}

	port := os.Getenv(pgEnvPort)
	if port == "" {
		return nil, fmt.Errorf("environment variable %s must be set", pgEnvPort)
	}

	name := os.Getenv(pgEnvName)
	if name == "" {
		return nil, fmt.Errorf("environment variable %s must be set", pgEnvName)
	}

	user := os.Getenv(pgEnvUser)
	if user == "" {
		return nil, fmt.Errorf("environment variable %s must be set", pgEnvUser)
	}

	password := os.Getenv(pgEnvPassword)
	if password == "" {
		return nil, fmt.Errorf("environment variable %s must be set", pgEnvPassword)
	}

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", host, port, name, user, password)
	return &pgConfig{dsn: dsn}, nil
}

func (p *pgConfig) DSN() string {
	return p.dsn
}
