package app

import (
	"context"
	impl "github.com/NikitaVi/minifier-sso/internal/api/auth"
	"github.com/NikitaVi/minifier-sso/internal/config"
	"github.com/NikitaVi/minifier-sso/internal/repo"
	repository "github.com/NikitaVi/minifier-sso/internal/repo/auth"
	"github.com/NikitaVi/minifier-sso/internal/service"
	serv "github.com/NikitaVi/minifier-sso/internal/service/auth"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type serviceProvider struct {
	grpcConfig config.GRPCConfig
	pgConfig   config.PGConfig

	jwtConfig config.JWTConfig

	db *pgxpool.Pool

	authImpl    *impl.Implementation
	authService service.AuthService
	authRepo    repo.AuthRepo
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatal(err)
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatal(err)
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) JWTConfig() config.JWTConfig {
	if s.jwtConfig == nil {
		cfg, err := config.NewJWTConfig()
		if err != nil {
			log.Fatal(err)
		}

		s.jwtConfig = cfg
	}

	return s.jwtConfig
}

func (s *serviceProvider) AuthRepo(ctx context.Context) repo.AuthRepo {
	if s.authRepo == nil {
		s.authRepo = repository.NewRepo(s.DB(ctx, s.PGConfig().DSN()))
	}

	return s.authRepo
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = serv.NewService(s.AuthRepo(ctx), s.JWTConfig())
	}

	return s.authService
}

func (s *serviceProvider) DB(ctx context.Context, dsn string) *pgxpool.Pool {
	if s.db == nil {
		pool, err := pgxpool.Connect(ctx, dsn)
		if err != nil {
			log.Fatal(err)
		}

		s.db = pool
	}

	return s.db
}

func (s *serviceProvider) AuthImpl(ctx context.Context) *impl.Implementation {
	if s.authImpl == nil {
		s.authImpl = impl.NewImplementation(s.AuthService(ctx))
	}

	return s.authImpl
}
