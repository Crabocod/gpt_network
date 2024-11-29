package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	pb "web.app/internal/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

func GenerateText(w http.ResponseWriter, r *http.Request) {

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial("python_service:50051", opts...)
	if err != nil {
		log.Fatalf("Не удалось подключиться: %v", err)
	}
	defer conn.Close()

	for conn.GetState() != connectivity.Ready {
		log.Println("Ожидание установления соединения...")
		time.Sleep(500 * time.Millisecond)
	}

	client := pb.NewTextGenServiceClient(conn)

	question := "Привет, как дела?"
	resp, err := client.GenerateText(context.Background(), &pb.GenerateRequest{Question: question})
	if err != nil {
		log.Fatalf("Ошибка при вызове GenerateText: %v", err)
	}

	json.NewEncoder(w).Encode(map[string]string{
		"resp": resp.Answer,
	})
	log.Printf("Ответ от Python: %s", resp.Answer)
}
