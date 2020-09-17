package adapter

import (
	"dddframework/api_server/entity"
	"log"
)

type GoogleUserService struct {
	GoogleAccount string `json:"google_account"`
	GoogleToken   string `json:"google_token"`
}

func (super *GoogleUserService) Auth(token string) (entity.User, error) {
	log.Println("it's google auth.")
	return entity.User{
		Name:     super.GoogleAccount,
		Password: super.GoogleToken,
		Provider: "Google",
		Token:    token,
	}, nil
}
