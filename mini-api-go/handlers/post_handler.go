package handlers

import (
	"mini-api-go/server"
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

func GetPosts(c *server.Context) {
	err := c.JSON(http.StatusOK, posts)
	if err != nil {
		http.Error(c.RWriter, "Error al codificar json", http.StatusInternalServerError)
	}
}

func CreatePost(c *server.Context) {
	var post Post
	err := c.BindJSON(&post)
	if err != nil {
		http.Error(c.RWriter, "Error al decodificar json", http.StatusBadRequest)
	}
	post.ID = nextID
	nextID++
	posts = append(posts, post)

	err = c.JSON(http.StatusCreated, post)
	if err != nil {
		http.Error(c.RWriter, "Error al codificar json", http.StatusInternalServerError)
	}
}

func GetPostByID(c *server.Context) {
	idStr := c.Request.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	for _, post := range posts {
		if post.ID == id {
			_ = c.JSON(http.StatusOK, post)
			return
		}
	}
	http.Error(c.RWriter, "Post not found", http.StatusNotFound)
}
