package error_code

import "strconv"

type ErrCode struct {
	CodeNum int
	//MainCode   string
	//ModuleCode string
	//FuncCode   string
	Msg string
}

func (c *ErrCode) String() string {
	return c.Msg
}

func (c *ErrCode) Int() int {
	return c.CodeNum
}

// NumString 回傳錯誤代碼
func (c *ErrCode) NumString() string {
	return strconv.Itoa(c.Int())
}

var (
	OK = ErrCode{0, "OK"}
)

var (
	AuthErr = ErrCode{1110010001, "驗證錯誤"}
)
