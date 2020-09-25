package instance

import (
	"dddframework/api_server/model"
	"dddframework/api_server/model/adapter"
)

var DBInstance model.DBService

func init() {
	DBInstance = new(adapter.MySqlAdapter)
}
