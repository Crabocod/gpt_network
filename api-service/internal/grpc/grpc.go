package grpc

import (
	"log"
	"time"

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
		log.Fatalf("Не удалось подключиться: %v", err)
	}

	for Conn.GetState() != connectivity.Ready {
		log.Println("Ожидание установления соединения...")
		time.Sleep(500 * time.Millisecond)
	}
	log.Println("gRPC соединение установлено")
}

func Close() {
	if Conn != nil {
		Conn.Close()
	}
}
