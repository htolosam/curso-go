package handlers

import (
	"mini-api-go/server"
	"net/http"
)

func Health(c *server.Context) {
	err := c.JSON(http.StatusOK, "OK")
	if err != nil {
		http.Error(c.RWriter, "Error al codificar json", http.StatusInternalServerError)
	}
}
