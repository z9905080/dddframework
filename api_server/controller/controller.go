package controller

import (
	"dddframework/api_server/entity/adapter"
	"dddframework/api_server/entity/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Get(c *gin.Context) {
	var model service.MsgService
	model = new(adapter.MsSqlMsgService)

	msgList, _ := model.GetMsg()

	c.JSON(http.StatusOK, msgList)
	return
}
