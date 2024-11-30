package grpc

import (
	"log"
	"time"

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
		log.Fatalf("Не удалось подключиться: %v", err)
		return err
	}

	for Conn.GetState() != connectivity.Ready {
		log.Println("Ожидание установления соединения...")
		time.Sleep(500 * time.Millisecond)
	}
	log.Println("gRPC соединение установлено")

	return nil
}

func Close() {
	if Conn != nil {
		Conn.Close()
	}
}
