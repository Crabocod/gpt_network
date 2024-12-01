package handlers

import (
	"context"

	grpcConn "generate/internal/grpc"
	pb "generate/internal/proto"
)

var hosts = map[string]string{
	"textgenService": "textgen:50051",
	"apiService":     "api:50052",
}

type Post struct {
	Question  string
	Answer    string
	ModelName string
}

func (p *Post) Save() error {
	err := grpcConn.Init(hosts["apiService"])
	if err != nil {
		return err
	}
	defer grpcConn.Close()

	saveClient := pb.NewSaveTextServiceClient(grpcConn.Conn)
	resp, err := saveClient.SaveGeneratedText(context.Background(), &pb.SaveRequest{GeneratedText: p.Answer, AuthorName: p.ModelName})
	if err != nil || !resp.Success {
		return err
	}

	return nil
}

func (p *Post) Generate() (string, error) {
	err := grpcConn.Init(hosts["textgenService"])
	if err != nil {
		return "", err
	}
	defer grpcConn.Close()

	client := pb.NewTextGenServiceClient(grpcConn.Conn)

	resp, err := client.GenerateText(context.Background(), &pb.GenerateRequest{Question: p.Question, ModelName: p.ModelName})
	if err != nil {
		return "", err
	}

	return resp.Answer, nil
}
