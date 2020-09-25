package adapter

import Structure "dddframework/api_server/model/structure"

type MySqlAdapter struct {
}

func (m MySqlAdapter) GetMsg() ([]Structure.MsgST, error) {

	// DB query
	return []Structure.MsgST{
		{"a", "Mysql"},
	}, nil
}
