package handlers

import (
	"context"

	"generate/internal/config"
	grpcConn "generate/internal/grpc"
	pb "generate/internal/proto"
)

type Comment struct {
	ID         string
	PostID     string
	Text       string
	AuthorName string
}

func (c *Comment) Save() error {
	err := grpcConn.Init(config.Data.ApiServiceHost)
	if err != nil {
		return err
	}
	defer grpcConn.Close()

	saveClient := pb.NewApiServiceClient(grpcConn.Conn)

	resp, err := saveClient.SaveComment(context.Background(),
		&pb.SaveCommentRequest{
			Text:       c.Text,
			AuthorName: c.AuthorName,
			PostId:     c.PostID,
		})

	if err != nil || !resp.Success {
		return err
	}

	return nil
}
