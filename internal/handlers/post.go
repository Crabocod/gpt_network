package handlers

import (
	"encoding/json"
	"net/http"

	"web.app/internal/models"
)

type Pagination struct {
	PageIndex      int `json:"pageIndex"`
	RecordsPerPage int `json:"recordsPerPage"`
}

type GetPostsRequest struct {
	Pagination Pagination `json:"pagination"`
}

type Post struct {
	ID       string `json:"id"`
	AuthorID string `json:"authorId"`
	Text     string `json:"text"`
	Date     string `json:"date"`
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
	var post Post
	user_id := r.Context().Value("user_id").(int)
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, `{"error": "Invalid request format"}`, http.StatusBadRequest)
		return
	}

	if post.Text == "" {
		http.Error(w, `{"error": "Missing text field"}`, http.StatusBadRequest)
		return
	}

	err := models.CreatePost(user_id, post.Text)
	if err != nil {
		http.Error(w, `{"error": "Failed to create post"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Post created successfully"}`))
}

func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	var post Post
	user_id := r.Context().Value("user_id").(int)
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, `{"error": "Invalid request format"}`, http.StatusBadRequest)
		return
	}

	if post.Text == "" || post.ID == "" {
		http.Error(w, `{"error": "Missing requierd fields"}`, http.StatusBadRequest)
		return
	}

	err := models.UpdatePost(user_id, post.ID, post.Text)
	if err != nil {
		http.Error(w, `{"error": "Failed to update post"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Post updated successfully"}`))
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	var post Post
	user_id := r.Context().Value("user_id").(int)
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, `{"error": "Invalid request format"}`, http.StatusBadRequest)
		return
	}

	if post.ID == "" {
		http.Error(w, `{"error": "Missing post_id field"}`, http.StatusBadRequest)
		return
	}

	err := models.DeletePost(user_id, post.ID)
	if err != nil {
		http.Error(w, `{"error": "Failed to delete post"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Post deleted successfully"}`))
}

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	var request GetPostsRequest
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

	totalRecords, err := models.CountPosts()
	if err != nil {
		http.Error(w, `{"error": "Failed to count posts"}`, http.StatusInternalServerError)
		return
	}

	offset := (request.Pagination.PageIndex - 1) * request.Pagination.RecordsPerPage
	posts, err := models.GetPosts(offset, request.Pagination.RecordsPerPage)
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
