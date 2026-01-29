package handlers

import (
	"log"
	"mini-api-go/models"
	"mini-api-go/server"
	"mini-api-go/services"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type PostHandler struct {
	postService *services.PostService
}

func NewPostHandler(postService *services.PostService) *PostHandler {
	return &PostHandler{postService: postService}
}

func (h *PostHandler) CreatePostHandler(c *server.Context) {
	var post models.Post
	if err := c.BindJSON(&post); err != nil {
		models.ResponseError(c, models.NewAppError("Error al decodificar json", http.StatusBadRequest))
		return
	}
	if err := validate.Struct(post); err != nil {
		models.ResponseError(c, models.NewAppError(err.Error(), http.StatusBadRequest))
		return
	}
	err := h.postService.CreatePost(c.Request.Context(), &post)
	if err != nil {
		models.ResponseError(c, models.NewAppError(err.Error(), http.StatusInternalServerError))
		return
	}
	err = c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Post creado con exito",
		"post":    post})
	if err != nil {
		log.Println("Error al codificar json en la respuesta del handler")
		return
	}
}

func (h *PostHandler) GetAllPostHandler(c *server.Context) {
	posts, err := h.postService.GetAllPost(c.Request.Context())
	if err != nil {
		return
	}
	_ = c.JSON(http.StatusOK, map[string]interface{}{
		"message": "posts revisados",
		"posts":   posts,
	})
}

func (h *PostHandler) GetPostHandler(c *server.Context) {
	postIDStr, err := strconv.Atoi(c.Request.PathValue("id"))
	if err != nil {
		models.ResponseError(c, models.NewAppError("Error al obtener el id del post", http.StatusBadRequest))
		return
	}
	postID := uint(postIDStr)
	post, err := h.postService.GetPostByID(c.Request.Context(), postID)
	if err != nil {
		models.ResponseError(c, models.NewAppError(err.Error(), http.StatusInternalServerError))
		return
	}
	_ = c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Post encontrado",
		"post":    post,
	})
}

func (h *PostHandler) GetPostsByUserIDHandler(c *server.Context) {
	userIDStr, err := strconv.Atoi(c.Request.PathValue("id"))
	if err != nil {
		models.ResponseError(c, models.NewAppError("Error al obtener el id del usuario", http.StatusBadRequest))
		return
	}
	userID := uint(userIDStr)
	posts, err := h.postService.GetPostsByUserID(c.Request.Context(), userID)
	if err != nil {
		models.ResponseError(c, models.NewAppError(err.Error(), http.StatusInternalServerError))
	}
	_ = c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Posts encontrados",
		"user_id": userID,
		"posts":   posts,
	})
}

func (h *PostHandler) UpdatePostHandler(c *server.Context) {
	var postUpd models.Post
	postIDStr, err := strconv.Atoi(c.Request.PathValue("id"))
	if err != nil {
		models.ResponseError(c, models.NewAppError("Error al obtener el id del post", http.StatusBadRequest))
		return
	}
	if err := c.BindJSON(&postUpd); err != nil {
		models.ResponseError(c, models.NewAppError("Error al decodificar json", http.StatusBadRequest))
		return
	}
	if err := validate.Struct(postUpd); err != nil {
		models.ResponseError(c, models.NewAppError(err.Error(), http.StatusBadRequest))
	}
	err = h.postService.UpdatePost(c.Request.Context(), &postUpd, uint(postIDStr))
	if err != nil {
		models.ResponseError(c, models.NewAppError(err.Error(), http.StatusInternalServerError))
		return
	}
	_ = c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Post actualizado con exito",
		"post":    postUpd,
	})
}

func (h *PostHandler) DeletePostHandler(c *server.Context) {
	postIDStr, err := strconv.Atoi(c.Request.PathValue("id"))
	if err != nil {
		models.ResponseError(c, models.NewAppError("Error al obtener el id del post", http.StatusBadRequest))
		return
	}
	postID := uint(postIDStr)
	err = h.postService.DeletePost(c.Request.Context(), postID)
	if err != nil {
		models.ResponseError(c, models.NewAppError(err.Error(), http.StatusInternalServerError))
		return
	}
	_ = c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Post eliminado con exito",
	})
}

func (h *PostHandler) GetAllPostsPaginatedHandler(c *server.Context) {
	page, err := strconv.Atoi(c.Request.PathValue("page"))
	if err != nil {
		page = 1
	}
	size, err := strconv.Atoi(c.Request.PathValue("size"))
	if err != nil {
		size = 10
	}
	log.Println(c.Request.URL.Query().Get("page"))
	log.Println(c.Request.URL.Query().Get("size"))
	posts, err := h.postService.GetAllPostsPaginated(c.Request.Context(), page, size)
	if err != nil {
		models.ResponseError(c, models.NewAppError(err.Error(), http.StatusInternalServerError))
		return
	}
	_ = c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Posts encontrados",
		"page":    page,
		"size":    size,
		"total":   len(posts),
		"posts":   posts,
	})
}
