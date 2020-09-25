package model

import Structure "dddframework/api_server/model/structure"

type DB struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Provider string `json:"provider"`
	Token    string `json:"token"`
}

type DBService interface {
	GetMsg() ([]Structure.MsgST, error)
}
