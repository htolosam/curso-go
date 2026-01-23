package main

import (
	"log"
	"mini-api-go/config"
	"mini-api-go/database"
	"mini-api-go/handlers"
	"mini-api-go/middleware"
	"mini-api-go/repositories"
	"mini-api-go/server"
	"mini-api-go/services"
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
	userRepo := repositories.NewUserRepository(database.DB)
	userService := services.NewUserService(userRepo, configuration)
	userHandler := handlers.NewUserHandler(userService)

	app := server.NewApp()
	app.Get("/health", handlers.Health)
	app.Post("/singup", userHandler.SignUp)
	app.Get("/posts/{id}", handlers.GetPostByID)
	app.Post("/posts", handlers.CreatePost)
	app.Post("/login", userHandler.Login)

	app.Get("/me", middleware.AuthMiddleware(userHandler.MeHandler))

	err := app.RunServer(configuration.Port)
	if err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
