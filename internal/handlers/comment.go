package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"web.app/internal/models"

	"github.com/gorilla/mux"
)

type GetCommentsRequest struct {
	PostID     int        `json:"postID"`
	Pagination Pagination `json:"pagination"`
}

type GetCommentsResponse struct {
	Comments   []models.Comment   `json:"comments"`
	Pagination PaginationResponse `json:"pagination"`
}

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	comment.AuthorID = r.Context().Value("user_id").(int)
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, `{"error": "Invalid request format"}`, http.StatusBadRequest)
		return
	}

	if comment.Text == "" || comment.PostID == 0 {
		http.Error(w, `{"error": "Missing requierd fields"}`, http.StatusBadRequest)
		return
	}

	err := comment.Save()
	if err != nil {
		http.Error(w, `{"error": "Failed to create comment"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Comment created successfully"}`))
}

func UpdateCommentHandler(w http.ResponseWriter, r *http.Request) {
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

	err := comment.Save()
	if err != nil {
		http.Error(w, `{"error": "Failed to update comment"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Comment updated successfully"}`))
}

func DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	comment.ID, _ = strconv.Atoi(mux.Vars(r)["id"])

	if comment.ID == 0 {
		http.Error(w, `{"error": "Missing comment id field"}`, http.StatusBadRequest)
		return
	}

	err := comment.Delete()
	if err != nil {
		http.Error(w, `{"error": "Failed to delete comment"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Comment deleted successfully"}`))
}

func GetCommentsHandler(w http.ResponseWriter, r *http.Request) {
	var request GetCommentsRequest
	var comment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, `{"error": "Invalid request format"}`, http.StatusBadRequest)
		return
	}

	if request.PostID == 0 {
		http.Error(w, `{"error": "Missing post id field"}`, http.StatusBadRequest)
		return
	}

	if request.Pagination.PageIndex == 0 {
		request.Pagination.PageIndex = 1
	}
	if request.Pagination.RecordsPerPage == 0 {
		request.Pagination.RecordsPerPage = 10
	}

	totalRecords, err := comment.Count(request.PostID)
	if err != nil {
		http.Error(w, `{"error": "Failed to count comments"}`, http.StatusInternalServerError)
		return
	}

	offset := (request.Pagination.PageIndex - 1) * request.Pagination.RecordsPerPage
	comments, err := comment.GetList(offset, request.Pagination.RecordsPerPage)
	if err != nil {
		http.Error(w, `{"error": "Failed to retrieve comments"}`, http.StatusInternalServerError)
		return
	}

	response := GetCommentsResponse{
		Comments: comments,
		Pagination: PaginationResponse{
			PageIndex:      request.Pagination.PageIndex,
			RecordsPerPage: request.Pagination.RecordsPerPage,
			TotalRecords:   totalRecords,
		},
	}

	json.NewEncoder(w).Encode(response)
}
