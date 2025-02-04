package rest

import (
	"github.com/Crabocod/gpt_network/api-service/internal/app/service"
)

type Controller struct {
	UserController    UserControllerInterface
	PostController    PostControllerInterface
	CommentController CommentControllerInterface
}

func NewController(service service.Service) *Controller {
	return &Controller{
		UserController:    NewUserController(service),
		PostController:    NewPostController(service),
		CommentController: NewCommentController(service),
	}
}
