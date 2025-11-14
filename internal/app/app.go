package app

import (
	"context"
	"github.com/NikitaVi/minifier-sso/internal/config"
	"github.com/NikitaVi/minifier-sso/internal/logger"
	"github.com/NikitaVi/minifier-sso/pkg/auth_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"sync"
)

type App struct {
	grpcServer      *grpc.Server
	serviceProvider *serviceProvider
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initApp(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) initApp(ctx context.Context) error {
	inits := []func(ctx2 context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPC,
	}

	for _, init := range inits {
		if err := init(ctx); err != nil {
			logger.Error("Failed to init: %v", err)
			return err
		}
	}

	return nil
}

func (a *App) initConfig(ctx context.Context) error {
	err := config.Load(ctx, ".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initGRPC(ctx context.Context) error {
	a.grpcServer = grpc.NewServer()

	reflection.Register(a.grpcServer)

	auth_v1.RegisterAuthV1Server(a.grpcServer, a.serviceProvider.AuthImpl(ctx))

	return nil
}

func (a *App) initServiceProvider(ctx context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) Run(ctx context.Context) error {
	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		defer wg.Done()

		if err := a.runGRPC(ctx); err != nil {
			panic(err)
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) runGRPC(context.Context) error {

	addr := a.serviceProvider.GRPCConfig().Address()

	logger.Info("Running GRPC server on: %s", addr)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Error("Failed to listen: %v", err)
		return err
	}

	logger.Info("Listening on: %s", addr)

	err = a.grpcServer.Serve(lis)
	if err != nil {
		logger.Error("Failed to server: %v", err)
		return err
	}

	return nil
}
