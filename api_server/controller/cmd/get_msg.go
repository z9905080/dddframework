package cmd

import (
	"dddframework/api_server/entity/instance"
	ErrorCode "dddframework/api_server/error_code"
	Response "dddframework/api_server/response"
	"github.com/gin-gonic/gin"
	GinDIGroup "github.com/z9905080/gin-di-router"
	"net/http"
)

// GetMsg B0001 取得訊息
func (*CmdController) GetMsg() (GinDIGroup.APIType, []gin.HandlerFunc) {
	apiType := GinDIGroup.Get
	var middlewareFuncList []gin.HandlerFunc
	controlFunc := GetMsg
	return apiType, append(middlewareFuncList, controlFunc)
}

func GetMsg(c *gin.Context) {

	resp, _ := instance.MsgInstance.GetMsg()

	c.JSON(http.StatusOK, Response.APIResponse{
		Status: Response.SUCCESS.String(),
		Code:   ErrorCode.OK.NumString(),
		Data:   resp,
	})
	return
}
