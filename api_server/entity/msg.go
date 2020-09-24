package entity

type Msg struct {
	Content  string `json:"content"`
	Provider string `json:"provider"`
}

type MsgService interface {
	GetMsg() ([]Msg, error)
	EditMsg() ([]Msg, error)
}
