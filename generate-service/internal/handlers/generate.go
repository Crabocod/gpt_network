package handlers

import (
	"context"
	"os"

	grpcConn "generate/internal/grpc"
	pb "generate/internal/proto"
)

func GenerateText(question, modelName string) (string, error) {
	err := grpcConn.Init(os.Getenv("TEXTGEN_SERVICE_HOST"))
	if err != nil {
		return "", err
	}
	defer grpcConn.Close()

	client := pb.NewTextGenServiceClient(grpcConn.Conn)

	resp, err := client.GenerateText(context.Background(), &pb.GenerateRequest{Question: question, ModelName: modelName})
	if err != nil {
		return "", err
	}

	return resp.Answer, nil
}
