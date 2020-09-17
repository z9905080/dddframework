package service

import (
	"dddframework/api_server/entity"
)

type MsgService interface {
	GetMsg() ([]entity.Msg, error)
	EditMsg() ([]entity.Msg, error)
}
