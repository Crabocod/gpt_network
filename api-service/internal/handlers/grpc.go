package handlers

import (
	"context"
	"log"

	"web.app/internal/models"
	pb "web.app/internal/proto"
)

type SaveTextService struct {
	pb.UnimplementedSaveTextServiceServer
}

func (s *SaveTextService) SaveGeneratedText(ctx context.Context, req *pb.SaveRequest) (*pb.SaveResponse, error) {
	var post models.Post
	post.Text = req.GetGeneratedText()
	post.AuthorID = 1
	err := post.Save()
	if err != nil {
		log.Printf("Ошибка при сохранении текста в БД: %v", err)
		return &pb.SaveResponse{Success: false}, err
	}

	return &pb.SaveResponse{Success: true}, nil
}
