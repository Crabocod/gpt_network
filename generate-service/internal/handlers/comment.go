package handlers

import (
	"context"

	"github.com/Crabocod/gpt_network/generate-service/internal/config"
	grpcConn "github.com/Crabocod/gpt_network/generate-service/internal/grpc"
	pb "github.com/Crabocod/gpt_network/generate-service/internal/proto"
)

type Comment struct {
	ID         string
	PostID     string
	Text       string
	AuthorName string
}

func (c *Comment) Save() error {
	err := grpcConn.Init(config.Data.Hosts.ApiService)
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
