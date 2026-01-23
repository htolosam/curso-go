package models

type Post struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id" validate:"required"`
	Tittle    string `json:"tittle" validate:"required"`
	Content   string `json:"content" validate:"required"`
	Author    string `json:"author"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
