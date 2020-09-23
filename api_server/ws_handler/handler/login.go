package handler

import (
	Response "dddframework/api_server/response"
	Logging "github.com/z9905080/gloger"
	"github.com/z9905080/melody"
)

// Login 登入功能
func Login(server *melody.Melody, session *melody.Session, data Response.WSResponse) {
	Logging.Debug("it's login.")
}
