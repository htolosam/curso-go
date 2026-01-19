package main

import (
	"log"
	"mini-api-go/config"
	"mini-api-go/database"
	"mini-api-go/handlers"
	"mini-api-go/server"
)

func main() {
	configuration := config.LoadConfig()

	if err := database.Connect(configuration.DatabaseURL); err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	defer func() {
		err := database.Close()
		if err != nil {
			log.Fatal("Error al cerrar la base de datos:", err)
		}
	}()
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
