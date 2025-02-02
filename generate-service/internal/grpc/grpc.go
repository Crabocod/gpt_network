package grpc

import (
	"time"

	"github.com/Crabocod/gpt_network/generate-service/internal/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

var Conn *grpc.ClientConn

func Init(host string) error {
	var err error
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	Conn, err = grpc.Dial(host, opts...)
	if err != nil {
		logger.Logrus.Fatalf("Failed to connect to %s: %v", host, err)
		return err
	}

	for Conn.GetState() != connectivity.Ready {
		logger.Logrus.Infof("Waiting for connection to %s...", host)
		time.Sleep(500 * time.Millisecond)
	}
	logger.Logrus.Infof("gRPC connection with %s established", host)

	return nil
}

func Close() {
	if Conn != nil {
		Conn.Close()
	}
}
