package models

type Post struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Tittle    string `json:"tittle"`
	Content   string `json:"content"`
	Author    string `json:"author"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
