package adapter

import (
	"dddframework/api_server/entity"
)

type MysqlMsgService struct {
	Content string `json:"content"`
}

func (service *MysqlMsgService) EditMsg() ([]entity.Msg, error) {
	msgList := mysqlQuery()
	var returnMsgList []entity.Msg
	for _, msg := range msgList {
		returnMsgList = append(returnMsgList, entity.Msg{
			Content:  msg.Content,
			Provider: "Mysql",
		})
	}
	return returnMsgList, nil
}

func (service *MysqlMsgService) GetMsg() ([]entity.Msg, error) {

	msgList := mysqlQuery()
	var returnMsgList []entity.Msg
	for _, msg := range msgList {
		returnMsgList = append(returnMsgList, entity.Msg{
			Content:  msg.Content,
			Provider: "Mysql",
		})
	}
	return returnMsgList, nil
}

func mysqlQuery() []MysqlMsgService {
	return []MysqlMsgService{
		{"MySql1"},
		{"MySql2"},
	}
}
