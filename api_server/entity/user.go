package entity

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Provider string `json:"provider"`
	Token    string `json:"token"`
}

type UserService interface {
	Auth(token string) (User, error)
}
