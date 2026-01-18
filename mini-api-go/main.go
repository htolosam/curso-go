package main

import (
	"log"
	"mini-api-go/config"
	"mini-api-go/handlers"
	"mini-api-go/server"
)

func main() {
	configuration := config.LoadConfig()
	app := server.NewApp()
	app.Get("/health", handlers.Health)
	app.Get("/posts", handlers.GetPosts)
	app.Get("/posts/{id}", handlers.GetPostByID)
	app.Post("/posts", handlers.CreatePost)

	err := app.RunServer(configuration.Port)
	if err != nil {
		log.Fatal("Error al iniciar el servidor")
	}
}
