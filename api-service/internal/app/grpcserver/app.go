package grpcserver

import (
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
	logger *logrus.Logger
	config *config.Config
	db     *sqlx.DB
}

func New(db *sqlx.DB, config *config.Config) *GrpcServer {
	return &GrpcServer{
		logger: logrus.New(),
		config: config,
		db:     db,
	}
}

func (s *GrpcServer) Run() error {
	store := postgresql.New(s.db)
	business := service.NewService(store)
	controller := grpcController.New(*business)

	server := grpc.NewServer()
	pb.RegisterApiServiceServer(server, controller)

	listener, err := net.Listen("tcp", s.config.GrpcServer.BindAddr)
	if err != nil {
		return err
	}

	s.logger.Info("Starting gRPC server on ", s.config.GrpcServer.BindAddr)
	if err := server.Serve(listener); err != nil {
		return err
	}

	return nil
}
