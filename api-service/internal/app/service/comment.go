package service

import (
	"github.com/Crabocod/gpt_network/api-service/internal/app/store"
	"github.com/Crabocod/gpt_network/api-service/internal/models"
)

type CommentServiceInterface interface {
	buildTree(comments []models.Comment) []models.Comment
	GetCount(postID int) (int, error)
	GetList(r GetCommentsRequest) ([]models.Comment, error)
	Save(comment models.Comment) error
	Delete(comment models.Comment) error
}

type GetCommentsRequest struct {
	PostID     int        `json:"postID"`
	Pagination Pagination `json:"pagination"`
}

type GetCommentsResponse struct {
	Comments   []models.Comment   `json:"comments"`
	Pagination PaginationResponse `json:"pagination"`
}

type CommentService struct {
	store store.Store
}

func NewCommentService(s store.Store) *CommentService {
	return &CommentService{
		store: s,
	}
}

func (s *CommentService) Delete(comment models.Comment) error {
	err := s.store.Comment().Delete(comment)
	if err != nil {
		return err
	}

	return nil
}

func (s *CommentService) Save(comment models.Comment) error {
	err := s.store.Comment().Save(comment)
	if err != nil {
		return err
	}

	return nil
}

func (s *CommentService) GetList(r GetCommentsRequest) ([]models.Comment, error) {
	offset := (r.Pagination.PageIndex - 1) * r.Pagination.RecordsPerPage
	comments, err := s.store.Comment().GetList(r.PostID, offset, r.Pagination.RecordsPerPage)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (s *CommentService) GetCount(postID int) (int, error) {
	count, err := s.store.Comment().GetCount(postID)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *CommentService) buildTree(comments []models.Comment) []models.Comment {
	commentMap := make(map[int]*models.Comment, len(comments))
	var roots []models.Comment

	for i := range comments {
		comment := &comments[i]
		comment.Children = make([]models.Comment, 0)
		commentMap[comment.ID] = comment
	}

	for i := range comments {
		comment := &comments[i]
		if comment.ParentID == nil {
			roots = append(roots, *comment)
		} else if parent, exists := commentMap[*comment.ParentID]; exists {
			parent.Children = append(parent.Children, *comment)
		}
	}

	return roots
}
