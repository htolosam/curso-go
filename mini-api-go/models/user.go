package models

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SingUpUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
