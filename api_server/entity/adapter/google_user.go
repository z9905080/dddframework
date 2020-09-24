package adapter

import (
	"dddframework/api_server/entity"
	"log"
)

type GoogleUserService struct {
}

func (super *GoogleUserService) Auth(token string) (entity.User, error) {
	log.Println("it's google auth.")
	return entity.User{
		Provider: "Google",
		Token:    token,
	}, nil
}
