package service

import "github.com/Crabocod/gpt_network/api-service/internal/app/store"

type Service struct {
	UserService    UserServiceInterface
	TokenService   TokenServiceInterface
	PostService    PostServiceInterface
	CommentService CommentServiceInterface
}

func NewService(s store.Store) *Service {
	return &Service{
		UserService:    NewUserService(s),
		TokenService:   NewTokenService(s),
		PostService:    NewPostService(s),
		CommentService: NewCommentService(s),
	}
}
