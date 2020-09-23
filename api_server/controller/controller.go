package controller

import (
	"dddframework/api_server/entity/adapter"
	"dddframework/api_server/entity/service"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

var hash = uuid.NewV4()

func Get(c *gin.Context) {
	var model service.MsgService
	model = new(adapter.MsSqlMsgService)

	_, _ = model.GetMsg()

	c.JSON(http.StatusOK, hash)
	return
}
