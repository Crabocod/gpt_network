package grpc

import (
	"time"

	"github.com/Crabocod/gpt_network/api-service/internal/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

var Conn *grpc.ClientConn

func Init() {
	var err error
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	Conn, err = grpc.Dial("textgen:50051", opts...)
	if err != nil {
		logger.Logrus.Fatalf("Failed to connect to textgen service: %v", err)
	}

	for Conn.GetState() != connectivity.Ready {
		logger.Logrus.Info("Waiting for connection to textgen service...")
		time.Sleep(500 * time.Millisecond)
	}
	logger.Logrus.Info("gRPC connection with textgen service established")
}

func Close() {
	if Conn != nil {
		Conn.Close()
	}
}
