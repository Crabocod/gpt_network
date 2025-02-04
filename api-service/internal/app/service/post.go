package service

import (
	"github.com/Crabocod/gpt_network/api-service/internal/app/store"
	"github.com/Crabocod/gpt_network/api-service/internal/models"
)

type PostServiceInterface interface {
	GetCount() (int, error)
	GetList(r GetPostsRequest) ([]models.Post, error)
	Save(post models.Post) error
	Delete(post models.Post) error
}

type GetPostsRequest struct {
	Pagination Pagination `json:"pagination"`
}

type GetPostsResponse struct {
	Posts      []models.Post      `json:"posts"`
	Pagination PaginationResponse `json:"pagination"`
}

type PostService struct {
	store store.Store
}

func NewPostService(s store.Store) *PostService {
	return &PostService{
		store: s,
	}
}

func (s *PostService) Save(post models.Post) error {
	err := s.store.Post().Save(post)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostService) GetCount() (int, error) {
	count, err := s.store.Post().GetCount()
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *PostService) GetList(r GetPostsRequest) ([]models.Post, error) {
	offset := (r.Pagination.PageIndex - 1) * r.Pagination.RecordsPerPage
	posts, err := s.store.Post().GetList(offset, r.Pagination.RecordsPerPage)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *PostService) Delete(post models.Post) error {
	err := s.store.Post().Delete(post)
	if err != nil {
		return err
	}

	return nil
}
