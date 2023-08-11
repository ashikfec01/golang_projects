package models

type RegisterUserCommand struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserCommand struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
