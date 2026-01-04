package main

import (
	"log"
	"mini-api-go/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handlers.Health)
	mux.HandleFunc("GET /posts", handlers.GetPosts)
	mux.HandleFunc("GET /posts/{id}", handlers.GetPostByID)
	mux.HandleFunc("POST /posts", handlers.CreatePost)
	log.Println("server in 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}
