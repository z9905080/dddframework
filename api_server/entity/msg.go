package entity

type Msg struct {
	Content  string `json:"content"`
	Provider string `json:"provider"`
}

//func (super *Msg) Auth(token string) (Msg, error) {
//	return Msg{
//		Token:    token,
//		Name:     "Default",
//		Password: "Default",
//		Provider: "Default",
//	}, nil
//}
