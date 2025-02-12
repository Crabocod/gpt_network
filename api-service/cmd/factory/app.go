package factory

import (
	"context"
	"github.com/Crabocod/gpt_network/api-service/internal/app/apiserver"
	"github.com/Crabocod/gpt_network/api-service/internal/app/grpcserver"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	apiServer  *apiserver.APIServer
	grpcServer *grpcserver.GrpcServer
	logger     *logrus.Logger
}

func New(apiServer *apiserver.APIServer, grpcServer *grpcserver.GrpcServer) *App {
	return &App{
		apiServer:  apiServer,
		grpcServer: grpcServer,
		logger:     logrus.New(),
	}
}

func (a *App) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	errChan := make(chan error, 2)

	go func() {
		if err := a.apiServer.Run(ctx); err != nil {
			errChan <- err
		}
	}()

	go func() {
		if err := a.grpcServer.Run(ctx); err != nil {
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:
		return err
	case sig := <-sigChan:
		a.logger.Info("Received signal: ", sig)
		cancel()
		return nil
	}
}
