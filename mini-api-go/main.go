package main

import (
	"fmt"
	"log"
	"mini-api-go/config"
	"mini-api-go/handlers"
	"mini-api-go/server"
	"net/http"
)

func main() {
	configuration := config.LoadConfig()
	fmt.Println("port is:: ", configuration.Port)
	fmt.Println("jwt secret is:: ", configuration.JwtSecret)
	fmt.Println("database url is:: ", configuration.DatabaseURL)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handlers.Health)
	mux.HandleFunc("GET /posts", handlers.GetPosts)
	mux.HandleFunc("GET /posts/{id}", handlers.GetPostByID)
	mux.HandleFunc("POST /posts", handlers.CreatePost)
	log.Println("server in 8080")
	app := server.NewApp()
	err := app.RunServer(configuration.Port)
	if err != nil {
		return
	}
}
