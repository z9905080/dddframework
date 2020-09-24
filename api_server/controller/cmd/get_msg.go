package cmd

import (
	"dddframework/api_server/entity/instance"
	"github.com/gin-gonic/gin"
	GinDIGroup "github.com/z9905080/gin-di-router"
	"net/http"
)

// GetMsg B0001 取得訊息
func (*CmdController) GetMsg() (GinDIGroup.APIType, []gin.HandlerFunc) {
	apiType := GinDIGroup.Post
	var middlewareFuncList []gin.HandlerFunc
	controlFunc := GetMsg
	return apiType, append(middlewareFuncList, controlFunc)
}

func GetMsg(c *gin.Context) {

	resp, _ := instance.MsgInstance.GetMsg()

	c.JSON(http.StatusOK, resp)
	return
}
