package rest

import (
	"encoding/json"
	"github.com/Crabocod/gpt_network/api-service/internal/app/service"
	"github.com/Crabocod/gpt_network/api-service/internal/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type PostControllerInterface interface {
	GetPostsHandler(w http.ResponseWriter, r *http.Request)
	CreatePostHandler(w http.ResponseWriter, r *http.Request)
	UpdatePostHandler(w http.ResponseWriter, r *http.Request)
	DeletePostHandler(w http.ResponseWriter, r *http.Request)
}

type PostController struct {
	service service.Service
}

func NewPostController(s service.Service) *PostController {
	return &PostController{
		service: s,
	}
}

func (c *PostController) GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	var request service.GetPostsRequest

	queryParams := r.URL.Query()
	request.Pagination.PageIndex, _ = strconv.Atoi(queryParams.Get("pageIndex"))
	request.Pagination.RecordsPerPage, _ = strconv.Atoi(queryParams.Get("recordsPerPage"))

	if request.Pagination.PageIndex == 0 || request.Pagination.RecordsPerPage == 0 {
		http.Error(w, `{"error": "Missing requierd fields"}`, http.StatusBadRequest)
		return
	}

	totalRecords, err := c.service.PostService.GetCount()
	if err != nil {
		http.Error(w, `{"error": "Failed to count posts"}`, http.StatusInternalServerError)
		return
	}

	posts, err := c.service.PostService.GetList(request)
	if err != nil {
		http.Error(w, `{"error": "Failed to retrieve posts"}`, http.StatusInternalServerError)
		return
	}

	response := service.GetPostsResponse{
		Posts: posts,
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

func (c *PostController) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	post.AuthorID = r.Context().Value("user_id").(int)
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, `{"error": "Invalid request format"}`, http.StatusBadRequest)
		return
	}

	if post.Text == "" {
		http.Error(w, `{"error": "Missing text field"}`, http.StatusBadRequest)
		return
	}

	err := c.service.PostService.Save(post)
	if err != nil {
		http.Error(w, `{"error": "Failed to create post"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(`{"message": "Post created successfully"}`))
	if err != nil {
		return
	}
}

func (c *PostController) UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	post.ID, _ = strconv.Atoi(mux.Vars(r)["id"])
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, `{"error": "Invalid request format"}`, http.StatusBadRequest)
		return
	}

	if post.Text == "" || post.ID == 0 {
		http.Error(w, `{"error": "Missing requierd fields"}`, http.StatusBadRequest)
		return
	}

	err := c.service.PostService.Save(post)
	if err != nil {
		http.Error(w, `{"error": "Failed to update post"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(`{"message": "Post updated successfully"}`))
	if err != nil {
		return
	}
}

func (c *PostController) DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	post.ID, _ = strconv.Atoi(mux.Vars(r)["id"])

	if post.ID == 0 {
		http.Error(w, `{"error": "Missing post id field"}`, http.StatusBadRequest)
		return
	}

	err := c.service.PostService.Delete(post)
	if err != nil {
		http.Error(w, `{"error": "Failed to delete post"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(`{"message": "Post deleted successfully"}`))
	if err != nil {
		return
	}
}
