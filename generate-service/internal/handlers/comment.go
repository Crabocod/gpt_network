package handlers

import (
	"context"
	"os"

	grpcConn "generate/internal/grpc"
	pb "generate/internal/proto"
)

type Comment struct {
	ID        int
	Question  string
	Answer    string
	ModelName string
}

func (p *Comment) Save(post *Post) error {
	err := grpcConn.Init(os.Getenv("API_SERVICE_HOST"))
	if err != nil {
		return err
	}
	defer grpcConn.Close()

	saveClient := pb.NewSaveCommentServiceClient(grpcConn.Conn)

	resp, err := saveClient.SaveComment(context.Background(),
		&pb.SaveCommentRequest{
			Text:       post.Answer,
			AuthorName: post.ModelName,
			PostId:     post.ID,
		})

	if err != nil || !resp.Success {
		return err
	}

	return nil
}
