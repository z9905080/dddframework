package adapter

import (
	"dddframework/api_server/entity"
	"dddframework/api_server/model/instance"
)

type MsgAdapter struct {
	Content string `json:"content"`
}

func (service *MsgAdapter) GetMsg() ([]entity.Msg, error) {

	msgList, err := instance.DBInstance.GetMsg()
	if err != nil {
		return nil, err
	}

	var returnMsgList []entity.Msg
	for _, msg := range msgList {
		returnMsgList = append(returnMsgList, entity.Msg{
			Content:  msg.Content,
			Provider: msg.Provider,
		})
	}

	return returnMsgList, nil
}
