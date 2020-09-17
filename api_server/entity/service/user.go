package service

import (
	"dddframework/api_server/entity"
)

type UserService interface {
	Auth(token string) (entity.User, error)
}
