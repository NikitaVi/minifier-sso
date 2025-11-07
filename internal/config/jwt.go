package config

import (
	"fmt"
	"os"
	"time"
)

const (
	jwtEnvTTL    = "TOKEN_TTL"
	jwtEnvSecret = "JWT_SECRET"
)

type JWTConfig interface {
	TTL() time.Duration
	Secret() string
}

type jwtConfig struct {
	tokenTTL  time.Duration
	jwtSecret string
}

func NewJWTConfig() (JWTConfig, error) {
	tokenTTL := os.Getenv(jwtEnvTTL)
	if tokenTTL == "" {
		return nil, fmt.Errorf(jwtEnvTTL + " not set")
	}
	ttl, err := time.ParseDuration(tokenTTL)
	if err != nil {
		return nil, err
	}

	secret := os.Getenv(jwtEnvSecret)
	if secret == "" {
		return nil, fmt.Errorf(jwtEnvSecret + " not set")
	}

	return &jwtConfig{tokenTTL: ttl, jwtSecret: secret}, nil
}

func (p *jwtConfig) TTL() time.Duration {
	return p.tokenTTL
}

func (p *jwtConfig) Secret() string {
	return p.jwtSecret
}
