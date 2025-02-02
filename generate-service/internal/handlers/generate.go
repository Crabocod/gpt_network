package handlers

import (
	"context"

	"github.com/Crabocod/gpt_network/generate-service/internal/config"
	grpcConn "github.com/Crabocod/gpt_network/generate-service/internal/grpc"
	pb "github.com/Crabocod/gpt_network/generate-service/internal/proto"
)

func GenerateText(question, modelName string) (string, error) {
	err := grpcConn.Init(config.Data.Hosts.TextgenService)
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
