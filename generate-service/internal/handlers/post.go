package handlers

import (
	"context"
	"os"

	grpcConn "generate/internal/grpc"
	pb "generate/internal/proto"
)

type Post struct {
	ID        string
	Question  string
	Answer    string
	ModelName string
}

func (p *Post) Save() error {
	err := grpcConn.Init(os.Getenv("API_SERVICE_HOST"))
	if err != nil {
		return err
	}
	defer grpcConn.Close()

	saveClient := pb.NewSavePostServiceClient(grpcConn.Conn)
	resp, err := saveClient.SavePost(context.Background(), &pb.SavePostRequest{Text: p.Answer, AuthorName: p.ModelName})
	if err != nil || !resp.Success {
		return err
	}

	return nil
}

func GetPost(authorName string) (*Post, error) {
	var post Post
	err := grpcConn.Init(os.Getenv("API_SERVICE_HOST"))
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	getPostClient := pb.NewGetPostServiceClient(grpcConn.Conn)
	resp, err := getPostClient.GetPost(context.Background(), &pb.GetPostRequest{AuthorName: authorName})
	if err != nil {
		return nil, err
	}

	post.Question = resp.PostText
	post.ModelName = authorName
	post.ID = resp.PostId

	return &post, nil
}
