package rest

import (
	"encoding/json"
	"github.com/Crabocod/gpt_network/api-service/internal/app/service"
	"github.com/Crabocod/gpt_network/api-service/internal/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type CommentControllerInterface interface {
	GetCommentsHandler(w http.ResponseWriter, r *http.Request)
	CreateCommentHandler(w http.ResponseWriter, r *http.Request)
	UpdateCommentHandler(w http.ResponseWriter, r *http.Request)
	DeleteCommentHandler(w http.ResponseWriter, r *http.Request)
}

type CommentController struct {
	service service.Service
}

func NewCommentController(s service.Service) *CommentController {
	return &CommentController{
		service: s,
	}
}

func (c *CommentController) GetCommentsHandler(w http.ResponseWriter, r *http.Request) {
	var request service.GetCommentsRequest

	queryParams := r.URL.Query()
	request.Pagination.PageIndex, _ = strconv.Atoi(queryParams.Get("pageIndex"))
	request.Pagination.RecordsPerPage, _ = strconv.Atoi(queryParams.Get("recordsPerPage"))
	request.PostID, _ = strconv.Atoi(mux.Vars(r)["post_id"])

	if request.PostID == 0 || request.Pagination.RecordsPerPage == 0 || request.Pagination.PageIndex == 0 {
		http.Error(w, `{"error": "Missing requierd fields"}`, http.StatusBadRequest)
		return
	}

	totalRecords, err := c.service.CommentService.GetCount(request.PostID)
	if err != nil {
		http.Error(w, `{"error": "Failed to count comments"}`, http.StatusInternalServerError)
		return
	}

	comments, err := c.service.CommentService.GetList(request)
	if err != nil {
		http.Error(w, `{"error": "Failed to retrieve comments"}`, http.StatusInternalServerError)
		return
	}

	response := service.GetCommentsResponse{
		Comments: comments,
		Pagination: service.PaginationResponse{
			PageIndex:      request.Pagination.PageIndex,
			RecordsPerPage: request.Pagination.RecordsPerPage,
			TotalRecords:   totalRecords,
		},
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

func (c *CommentController) CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	comment.AuthorID = r.Context().Value("user_id").(int)
	comment.PostID, _ = strconv.Atoi(mux.Vars(r)["post_id"])

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, `{"error": "Invalid request format"}`, http.StatusBadRequest)
		return
	}

	if comment.Text == "" || comment.PostID == 0 {
		http.Error(w, `{"error": "Missing requierd fields"}`, http.StatusBadRequest)
		return
	}

	err := c.service.CommentService.Save(comment)
	if err != nil {
		http.Error(w, `{"error": "Failed to create comment"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(`{"message": "Comment created successfully"}`))
	if err != nil {
		return
	}
}

func (c *CommentController) UpdateCommentHandler(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	comment.ID, _ = strconv.Atoi(mux.Vars(r)["id"])
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, `{"error": "Invalid request format"}`, http.StatusBadRequest)
		return
	}

	if comment.Text == "" || comment.ID == 0 {
		http.Error(w, `{"error": "Missing requierd fields"}`, http.StatusBadRequest)
		return
	}

	err := c.service.CommentService.Save(comment)
	if err != nil {
		http.Error(w, `{"error": "Failed to update comment"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(`{"message": "Comment updated successfully"}`))
	if err != nil {
		return
	}
}

func (c *CommentController) DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	comment.ID, _ = strconv.Atoi(mux.Vars(r)["id"])

	if comment.ID == 0 {
		http.Error(w, `{"error": "Missing comment id field"}`, http.StatusBadRequest)
		return
	}

	err := c.service.CommentService.Delete(comment)
	if err != nil {
		http.Error(w, `{"error": "Failed to delete comment"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(`{"message": "Comment deleted successfully"}`))
	if err != nil {
		return
	}
}
