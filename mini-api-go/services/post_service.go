package services

import (
	"context"
	"fmt"
	"mini-api-go/config"
	"mini-api-go/models"
	"mini-api-go/repositories"
	"time"
)

type PostService struct {
	postRepository *repositories.PostRepository
	userService    *UserService
	config         *config.Config
}

func NewPostService(postRepository *repositories.PostRepository, userService *UserService, cfg *config.Config) *PostService {
	return &PostService{postRepository: postRepository, userService: userService, config: cfg}
}

// CreatePost crear nuevo post
func (s *PostService) CreatePost(ctx context.Context, post *models.Post) error {
	user, err := s.userService.GetUserByID(ctx, post.UserID)
	if err != nil {
		return err
	}
	if user != nil {
		post.Author = fmt.Sprintf("%s %s", user.Name, user.LastName)
		if err = s.postRepository.Create(ctx, post); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("No se pudo crear el post, el usuario no existe")
}

// GetAllPost get all post
func (s *PostService) GetAllPost(ctx context.Context) ([]models.Post, error) {
	return s.postRepository.FindAll(ctx)
}

// GetPostByID get post by id
func (s *PostService) GetPostByID(ctx context.Context, id uint) (*models.Post, error) {
	return s.postRepository.FindByID(ctx, id)
}

// GetPostsByUserID get post by user id
func (s *PostService) GetPostsByUserID(ctx context.Context, userID uint) ([]models.Post, error) {
	user, err := s.userService.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	if user != nil {
		if posts, err = s.postRepository.FindByUserID(ctx, userID); err != nil {
			return nil, err
		}
		return posts, nil
	}
	return nil, fmt.Errorf("No se encontraron posts para el usuario")
}

// UpdatePost update post
func (s *PostService) UpdatePost(ctx context.Context, postUpd *models.Post, postID uint) error {
	postDB, err := s.postRepository.FindByID(ctx, postID)
	if err != nil {
		return err
	}
	if postDB.Tittle != postUpd.Tittle {
		postDB.Tittle = postUpd.Tittle
	}
	if postDB.Content == postUpd.Content {
		postDB.Content = postUpd.Content
	}
	postDB.UpdatedAt = time.Now().String()
	return s.postRepository.Update(ctx, postDB)
}

// DeletePost delete post by ID
func (s *PostService) DeletePost(ctx context.Context, postID uint) error {
	_, err := s.postRepository.FindByID(ctx, postID)
	if err != nil {
		return err
	}
	return s.postRepository.Delete(ctx, postID)
}

// GetAllPostsPaginated get all post paginated
func (s *PostService) GetAllPostsPaginated(ctx context.Context, page, pageSize int) ([]models.Post, error) {
	return s.postRepository.FindAllPaginated(ctx, page, pageSize)
}
