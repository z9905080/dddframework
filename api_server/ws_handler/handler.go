package ws_handler

import (
	Response "dddframework/api_server/response"
	Handler "dddframework/api_server/ws_handler/handler"
	OperationCode "dddframework/api_server/ws_handler/operation_code"
	Tool "dddframework/package/tool"
	Logging "github.com/z9905080/gloger"
	"github.com/z9905080/melody"
)

func ClientRequestHandler(server *melody.Melody, session *melody.Session, clientMsgBuffer []byte) {

	var data Response.WSResponse
	_ = Tool.JSONDeSerialize(clientMsgBuffer, &data)

	funcMap := getServerFuncMap() // 百人模組

	if function, isExist := funcMap[data.OperationCode]; isExist {
		function(server, session, data)
	} else {
		Logging.Error("錯誤的OperationCode:", data.OperationCode)
	}

}

// getServerFuncMap WS的OperationCode要註冊進來並對應到相對的Func
// 這裡有參照 C# Reflection的一個核心概念
// 使用函數變數延遲執行的時機
func getServerFuncMap() map[string]func(*melody.Melody, *melody.Session, Response.WSResponse) {

	return map[string]func(*melody.Melody, *melody.Session, Response.WSResponse){
		OperationCode.C2S_Login.String(): Handler.Login,
	}
}
