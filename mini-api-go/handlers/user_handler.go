package handlers

import (
	"log"
	"mini-api-go/models"
	"mini-api-go/server"
	"mini-api-go/services"
	"net/http"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) SignUp(c *server.Context) {
	var req models.SingUpUser
	if err := c.BindJSON(&req); err != nil {
		models.ResponseError(c, models.NewAppError("Error al decodificar json", http.StatusBadRequest))
		return
	}
	if req.Name == "" || req.Email == "" || req.Password == "" {
		models.ResponseError(c, models.NewAppError("Datos invalidos", http.StatusBadRequest))
		return
	}
	user, err := h.userService.SingUp(c.Request.Context(), req.Name, req.LastName, req.Email, req.Password)
	if err != nil {
		models.ResponseError(c, models.NewAppError(err.Error(), http.StatusInternalServerError))
		return
	}
	err = c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Usuario creado con exito",
		"user":    user})
	if err != nil {
		log.Println("Error al codificar json en la respuesta del handler")
		return
	}
}
