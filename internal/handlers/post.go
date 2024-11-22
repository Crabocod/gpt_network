package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"web.app/internal/models"

	"github.com/gorilla/mux"
)

type GetPostsRequest struct {
	Pagination Pagination `json:"pagination"`
}

type GetPostsResponse struct {
	Posts      []models.Post      `json:"posts"`
	Pagination PaginationResponse `json:"pagination"`
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

	if post.ID == 0 {
		http.Error(w, `{"error": "Missing post id field"}`, http.StatusBadRequest)
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

	queryParams := r.URL.Query()
	request.Pagination.PageIndex, _ = strconv.Atoi(queryParams.Get("pageIndex"))
	request.Pagination.RecordsPerPage, _ = strconv.Atoi(queryParams.Get("recordsPerPage"))

	if request.Pagination.PageIndex == 0 || request.Pagination.RecordsPerPage == 0 {
		http.Error(w, `{"error": "Missing requierd fields"}`, http.StatusBadRequest)
		return
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
