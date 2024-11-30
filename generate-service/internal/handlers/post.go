package handlers

import (
	"context"

	grpcConn "generate/internal/grpc"
	pb "generate/internal/proto"
	"generate/internal/services"
)

var hosts = map[string]string{
	"textgenService": "textgen:50051",
	"apiService":     "api:50052",
}

var questions = []string{
	"Как дела?",
	"Что делаешь?",
	"Что хочешь?",
	"Хаха",
	"Привет",
}

var models = []string{
	"МихаилGPT",
	"ЕваGPT",
	"АртурGPT",
	"РомаGPT",
	"РусланGPT",
	"СеняGPT",
}

func Save(text string) error {
	err := grpcConn.Init(hosts["apiService"])
	if err != nil {
		return err
	}
	defer grpcConn.Close()

	saveClient := pb.NewSaveTextServiceClient(grpcConn.Conn)
	resp, err := saveClient.SaveGeneratedText(context.Background(), &pb.SaveRequest{GeneratedText: text})
	if err != nil || !resp.Success {
		return err
	}

	return nil
}

func Generate() (string, error) {
	err := grpcConn.Init(hosts["textgenService"])
	if err != nil {
		return "", err
	}
	defer grpcConn.Close()

	client := pb.NewTextGenServiceClient(grpcConn.Conn)

	question := services.RandomChoice(questions)
	model := services.RandomChoice(models)
	resp, err := client.GenerateText(context.Background(), &pb.GenerateRequest{Question: question, ModelName: model})
	if err != nil {
		return "", err
	}

	return resp.Answer, nil
}
