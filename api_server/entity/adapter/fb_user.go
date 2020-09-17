package adapter

import (
	"dddframework/api_server/entity"
	"log"
)

type FacebookUserService struct {
	FacebookAccount string `json:"facebook_account"`
	FacebookToken   string `json:"facebook_token"`
}

func (super *FacebookUserService) Auth(token string) (entity.User, error) {
	log.Println("it's FB auth.")
	return entity.User{
		Token:    token,
		Name:     super.FacebookAccount,
		Password: super.FacebookToken,
		Provider: "FB",
	}, nil
}
