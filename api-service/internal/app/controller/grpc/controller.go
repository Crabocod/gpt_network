package grpc

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Crabocod/gpt_network/api-service/internal/app/service"
	"github.com/Crabocod/gpt_network/api-service/internal/models"
	pb "github.com/Crabocod/gpt_network/api-service/internal/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

type Controller struct {
	service service.Service
	pb.UnimplementedApiServiceServer
}

func NewController(s service.Service) *Controller {
	return &Controller{
		service: s,
	}
}

func (c *Controller) SaveComment(ctx context.Context, req *pb.SaveCommentRequest) (*pb.SaveCommentResponse, error) {
	var comment models.Comment

	User, err := c.service.UserService.GetByName(req.GetAuthorName())
	if err != nil {
		return &pb.SaveCommentResponse{Success: false}, err
	}

	comment.Text = req.GetText()
	comment.PostID, _ = strconv.Atoi(req.GetPostId())
	comment.AuthorID = User.ID
	err = c.service.CommentService.Save(comment)
	if err != nil {
		return &pb.SaveCommentResponse{Success: false}, err
	}

	return &pb.SaveCommentResponse{Success: true}, nil
}


func (c *Controller) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	Post, err := c.service.PostService.GetLatestFilteredPost(req.GetAuthorName())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "Post not found for author: %s", req.GetAuthorName())
		}
		return nil, status.Errorf(codes.Internal, "Error retrieving post: %v", err)
	}

	return &pb.GetPostResponse{
		PostId:   strconv.Itoa(Post.ID),
		PostText: Post.Text,
	}, nil
}

func (c *Controller) SavePost(ctx context.Context, req *pb.SavePostRequest) (*pb.SavePostResponse, error) {
	var post models.Post

	User, err := c.service.UserService.GetByName(req.GetAuthorName())
	if err != nil {
		return &pb.SavePostResponse{Success: false}, err
	}

	post.Text = req.GetText()
	post.AuthorID = User.ID
	err = c.service.PostService.Save(post)
	if err != nil {
		return &pb.SavePostResponse{Success: false}, err
	}

	return &pb.SavePostResponse{Success: true}, nil
}
