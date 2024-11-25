package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	pb "proto/go/textgen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

func GenerateText(w http.ResponseWriter, r *http.Request) {
	// Создаем контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Создаем соединение с gRPC сервером
	conn, err := grpc.NewClient(
		"localhost:50051", // адрес сервера
		grpc.WithTransportCredentials(insecure.NewCredentials()), // Указываем, что соединение небезопасное
	)
	if err != nil {
		log.Fatalf("Не удалось подключиться: %v", err)
	}
	defer conn.Close()

	// Проверяем состояние соединения
	if conn.GetState() != connectivity.Ready {
		log.Fatalf("Соединение не готово: %v", conn.GetState())
	}

	// Создаем клиента с использованием нового соединения
	client := pb.NewTextGenServiceClient(conn)

	// Вызов gRPC метода
	question := "Привет, как дела?"
	resp, err := client.GenerateText(ctx, &pb.GenerateRequest{Question: question})
	if err != nil {
		log.Fatalf("Ошибка при вызове GenerateText: %v", err)
	}

	json.NewEncoder(w).Encode(map[string]string{
		"resp": resp.Answer,
	})
	log.Printf("Ответ от Python: %s", resp.Answer)
}
