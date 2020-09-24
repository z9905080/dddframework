package adapter

import (
	"dddframework/api_server/entity"
	"log"
)

type FacebookUserService struct {
	FacebookToken string `json:"facebook_token"`
}

func (super *FacebookUserService) Auth(token string) (entity.User, error) {
	log.Println("it's FB auth.")
	return entity.User{
		Token:    token,
		Password: super.FacebookToken,
		Provider: "FB",
	}, nil
}
