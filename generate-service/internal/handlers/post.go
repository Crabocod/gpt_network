package handlers

import (
	"context"
	"os"

	grpcConn "generate/internal/grpc"
	pb "generate/internal/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Post struct {
	ID         string
	Text       string
	AuthorName string
}

func (p *Post) Save() error {
	err := grpcConn.Init(os.Getenv("API_SERVICE_HOST"))
	if err != nil {
		return err
	}
	defer grpcConn.Close()

	saveClient := pb.NewApiServiceClient(grpcConn.Conn)
	resp, err := saveClient.SavePost(context.Background(), &pb.SavePostRequest{Text: p.Text, AuthorName: "test"})
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

	getPostClient := pb.NewApiServiceClient(grpcConn.Conn)
	resp, err := getPostClient.GetPost(context.Background(), &pb.GetPostRequest{AuthorName: authorName})
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return nil, nil
		}
		return nil, err
	}

	post.Text = resp.PostText
	post.AuthorName = authorName
	post.ID = resp.PostId

	return &post, nil
}
