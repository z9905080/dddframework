package entity

import "log"

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Provider string `json:"provider"`
	Token    string `json:"token"`
}

func (u *User) Auth(token string) (User, error) {
	log.Println("it's default auth.")
	return User{
		Token:    token,
		Name:     "Default",
		Password: "Default",
		Provider: "Default",
	}, nil
}
