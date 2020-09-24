package instance

import (
	"dddframework/api_server/entity"
	"dddframework/api_server/entity/adapter"
)

var MsgInstance entity.MsgService

func init() {
	MsgInstance = new(adapter.MysqlMsgService)
}
