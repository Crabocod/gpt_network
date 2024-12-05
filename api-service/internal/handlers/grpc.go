package handlers

import (
	"context"
	"log"
	"strconv"

	"web.app/internal/models"
	pb "web.app/internal/proto"
)

type SavePostService struct {
	pb.UnimplementedSavePostServiceServer
}

type GetPostService struct {
	pb.UnimplementedGetPostServiceServer
}

type SaveCommentService struct {
	pb.UnimplementedSaveCommentServiceServer
}

func (s *SaveCommentService) SaveComment(ctx context.Context, req *pb.SaveCommentRequest) (*pb.SaveCommentResponse, error) {
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

func (s *GetPostService) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {

	Post, err := models.GetLatestFilteredPost(req.GetAuthorName())
	if err != nil {
		log.Printf("Ошибка при получении поста: %v", err)
		return &pb.GetPostResponse{}, err
	}

	return &pb.GetPostResponse{
		PostId:   strconv.Itoa(Post.ID),
		PostText: Post.Text,
	}, nil
}

func (s *SavePostService) SavePost(ctx context.Context, req *pb.SavePostRequest) (*pb.SavePostResponse, error) {
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
