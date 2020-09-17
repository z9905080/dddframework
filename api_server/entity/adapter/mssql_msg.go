package adapter

import (
	"dddframework/api_server/entity"
)

type MsSqlMsgService struct {
	MsContent string `json:"ms_content"`
}

func (service *MsSqlMsgService) EditMsg() ([]entity.Msg, error) {
	msgList := mssqlQuery()
	var returnMsgList []entity.Msg
	for _, msg := range msgList {
		returnMsgList = append(returnMsgList, entity.Msg{
			Content:  msg.MsContent,
			Provider: "Mssql",
		})
	}
	return returnMsgList, nil
}

func (service *MsSqlMsgService) GetMsg() ([]entity.Msg, error) {
	msgList := mssqlQuery()
	var returnMsgList []entity.Msg
	for _, msg := range msgList {
		returnMsgList = append(returnMsgList, entity.Msg{
			Content:  msg.MsContent,
			Provider: "Mssql",
		})
	}
	return returnMsgList, nil
}

func mssqlQuery() []MsSqlMsgService {
	return []MsSqlMsgService{
		{"MsSql3"},
		{"MsSql4"},
	}
}
