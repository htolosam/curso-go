package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mini-api-go/models"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(cxt context.Context, post *models.Post) error {
	query := "INSERT INTO posts (title, content, user_id) VALUES (?, ?, ?)"
	result, err := r.db.ExecContext(cxt, query, post.Tittle, post.Content, post.UserID)
	if err != nil {
		return fmt.Errorf("Error al insertar post: %s", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("Error al obtener el id del post insertado: %s", err)
	}
	post.ID = uint(id)
	return nil
}

func (r *PostRepository) FindAll(cxt context.Context) ([]models.Post, error) {
	query := "SELECT id, title, content, user_id FROM posts ORDER BY created_at DESC"
	rows, err := r.db.QueryContext(cxt, query)
	if err != nil {
		return nil, fmt.Errorf("Error al obtener los posts: %s", err)
	}
	defer rows.Close()
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err = rows.Scan(&post.ID, &post.Tittle, &post.Content, &post.UserID,
			&post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, fmt.Errorf("Error al decodificar los posts: %s", err)
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *PostRepository) FindByID(cxt context.Context, id uint) (*models.Post, error) {
	query := "SELECT * FROM posts WHERE id = ?"
	post := &models.Post{}
	err := r.db.QueryRowContext(cxt, query, id).Scan(&post.ID, &post.Tittle, &post.Content, &post.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("No existe un post con id %d", id)
		}
		return nil, fmt.Errorf("Error al obtener el post con id %d: %s", id, err)
	}
	return post, nil
}

func (r *PostRepository) FindByUserID(cxt context.Context, userID uint) ([]models.Post, error) {
	query := "SELECT * FROM posts WHERE user_id = ? ORDER BY created_at DESC"
	rows, err := r.db.QueryContext(cxt, query, userID)
	if err != nil {
		return nil, fmt.Errorf("Error al obtener los posts de un usuario: %s", err)
	}
	defer rows.Close()
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err = rows.Scan(&post.ID, &post.Tittle, &post.Content, &post.UserID,
			&post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, fmt.Errorf("Error al decodificar los posts de un usuario: %s", err)
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *PostRepository) Update(cxt context.Context, post *models.Post) error {
	query := "UPDATE posts SET title = ?, content = ? WHERE id = ?"
	result, err := r.db.ExecContext(cxt, query, post.Tittle, post.Content, post.ID)
	if err != nil {
		return fmt.Errorf("Error al actualizar el post: %s", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error al verificar la actualización %s", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("post no encontrado")
	}
	return nil
}

func (r *PostRepository) Delete(ctx context.Context, id uint) error {
	query := "DELETE FROM posts WHERE id = ?"
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("Error al eliminar el post: %s", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error al verificar la eliminación %s", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("post no encontrado")
	}
	return nil
}
