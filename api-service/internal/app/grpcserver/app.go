package grpcserver

import (
	"context"
	grpcController "github.com/Crabocod/gpt_network/api-service/internal/app/controller/grpc"
	"github.com/Crabocod/gpt_network/api-service/internal/app/service"
	"github.com/Crabocod/gpt_network/api-service/internal/app/store/postgresql"
	"github.com/Crabocod/gpt_network/api-service/internal/config"
	pb "github.com/Crabocod/gpt_network/api-service/internal/proto"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type GrpcServer struct {
	logger     *logrus.Logger
	config     *config.Config
	db         *sqlx.DB
	grpcServer *grpc.Server
	listener   net.Listener
}

func New(db *sqlx.DB, config *config.Config) *GrpcServer {
	s := &GrpcServer{
		logger: logrus.New(),
		config: config,
		db:     db,
	}

	store := postgresql.New(s.db)
	business := service.NewService(store)
	controller := grpcController.NewController(*business)

	s.grpcServer = grpc.NewServer()
	pb.RegisterApiServiceServer(s.grpcServer, controller)

	return s
}

func (s *GrpcServer) Run(ctx context.Context) error {
	var err error
	s.listener, err = net.Listen("tcp", s.config.GrpcServer.BindAddr)
	if err != nil {
		return err
	}

	s.logger.Info("Starting gRPC server on ", s.config.GrpcServer.BindAddr)

	go func() {
		<-ctx.Done()
		s.logger.Info("Shutting down gRPC server...")
		s.grpcServer.GracefulStop()
	}()

	if err := s.grpcServer.Serve(s.listener); err != nil {
		return err
	}

	return nil
}

func (s *GrpcServer) Shutdown() error {
	s.grpcServer.GracefulStop()
	return nil
}
