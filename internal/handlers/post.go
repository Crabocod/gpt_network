package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"web.app/internal/models"

	"github.com/gorilla/mux"
)

type Pagination struct {
	PageIndex      int `json:"pageIndex"`
	RecordsPerPage int `json:"recordsPerPage"`
}

type GetPostsRequest struct {
	Pagination Pagination `json:"pagination"`
}

type GetPostsResponse struct {
	Posts      []models.Post      `json:"posts"`
	Pagination PaginationResponse `json:"pagination"`
}

type PaginationResponse struct {
	PageIndex      int `json:"pageIndex"`
	RecordsPerPage int `json:"recordsPerPage"`
	TotalRecords   int `json:"totalRecords"`
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
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

	err := post.Save()
	if err != nil {
		http.Error(w, `{"error": "Failed to create post"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Post created successfully"}`))
}

func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	post.ID, _ = strconv.Atoi(mux.Vars(r)["id"])
	post.AuthorID = r.Context().Value("user_id").(int)
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, `{"error": "Invalid request format"}`, http.StatusBadRequest)
		return
	}

	if post.Text == "" || post.ID == 0 {
		http.Error(w, `{"error": "Missing requierd fields"}`, http.StatusBadRequest)
		return
	}

	err := post.Save()
	if err != nil {
		http.Error(w, `{"error": "Failed to update post"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Post updated successfully"}`))
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	post.ID, _ = strconv.Atoi(mux.Vars(r)["id"])
	post.AuthorID = r.Context().Value("user_id").(int)

	if post.ID == 0 {
		http.Error(w, `{"error": "Missing post_id field"}`, http.StatusBadRequest)
		return
	}

	err := post.Delete()
	if err != nil {
		http.Error(w, `{"error": "Failed to delete post"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Post deleted successfully"}`))
}

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	var request GetPostsRequest
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, `{"error": "Invalid request format"}`, http.StatusBadRequest)
		return
	}

	if request.Pagination.PageIndex == 0 {
		request.Pagination.PageIndex = 1
	}
	if request.Pagination.RecordsPerPage == 0 {
		request.Pagination.RecordsPerPage = 10
	}

	totalRecords, err := post.Count()
	if err != nil {
		http.Error(w, `{"error": "Failed to count posts"}`, http.StatusInternalServerError)
		return
	}

	offset := (request.Pagination.PageIndex - 1) * request.Pagination.RecordsPerPage
	posts, err := post.GetList(offset, request.Pagination.RecordsPerPage)
	if err != nil {
		http.Error(w, `{"error": "Failed to retrieve posts"}`, http.StatusInternalServerError)
		return
	}

	response := GetPostsResponse{
		Posts: posts,
		Pagination: PaginationResponse{
			PageIndex:      request.Pagination.PageIndex,
			RecordsPerPage: request.Pagination.RecordsPerPage,
			TotalRecords:   totalRecords,
		},
	}

	json.NewEncoder(w).Encode(response)
}
