package handlers

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strconv"

	"web.app/internal/models"
	pb "web.app/internal/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ApiService struct {
	pb.UnimplementedApiServiceServer
}

func (s *ApiService) SaveComment(ctx context.Context, req *pb.SaveCommentRequest) (*pb.SaveCommentResponse, error) {
	var comment models.Comment

	User, err := models.GetUserByName(req.GetAuthorName())
	if err != nil {
		log.Printf("Ошибка при получении пользователя: %v", err)
		return &pb.SaveCommentResponse{Success: false}, err
	}

	comment.Text = req.GetText()
	comment.PostID, _ = strconv.Atoi(req.GetPostId())
	comment.AuthorID = User.ID
	err = comment.Save()
	if err != nil {
		log.Printf("Ошибка при сохранении комментария: %v", err)
		return &pb.SaveCommentResponse{Success: false}, err
	}

	return &pb.SaveCommentResponse{Success: true}, nil
}

func (s *ApiService) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	Post, err := models.GetLatestFilteredPost(req.GetAuthorName())
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

func (s *ApiService) SavePost(ctx context.Context, req *pb.SavePostRequest) (*pb.SavePostResponse, error) {
	var post models.Post

	User, err := models.GetUserByName(req.GetAuthorName())
	if err != nil {
		log.Printf("Ошибка при получении пользователя: %v", err)
		return &pb.SavePostResponse{Success: false}, err
	}

	post.Text = req.GetText()
	post.AuthorID = User.ID
	err = post.Save()
	if err != nil {
		log.Printf("Ошибка при сохранении поста: %v", err)
		return &pb.SavePostResponse{Success: false}, err
	}

	return &pb.SavePostResponse{Success: true}, nil
}
