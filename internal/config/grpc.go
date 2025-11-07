package config

import (
	"fmt"
	"os"
)

const (
	grpcEnvHost = "GRPC_HOST"
	grpcEnvPort = "GRPC_PORT"
)

type GRPCConfig interface {
	Address() string
}

type grpcConfig struct {
	host string
	port string
}

func NewGRPCConfig() (GRPCConfig, error) {
	host := os.Getenv(grpcEnvHost)
	if host == "" {
		return nil, fmt.Errorf("environment variable %s must be set", grpcEnvHost)
	}

	port := os.Getenv(grpcEnvPort)
	if port == "" {
		return nil, fmt.Errorf("environment variable %s must be set", grpcEnvPort)
	}

	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}

func (g *grpcConfig) Address() string {
	return g.host + ":" + g.port
}
