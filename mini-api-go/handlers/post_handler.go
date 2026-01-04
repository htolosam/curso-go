package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Post struct {
	ID      int    `json:"id"`
	Tittle  string `json:"tittle"`
	Content string `json:"content"`
}

var posts []Post
var nextID int = 1

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(posts)
	if err != nil {
		http.Error(w, "Error al codificar json", http.StatusInternalServerError)
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Error al decodificar json", http.StatusBadRequest)
	}
	post.ID = nextID
	nextID++
	posts = append(posts, post)
}

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	for _, post := range posts {
		if post.ID == id {
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(post)
		}
	}
}
