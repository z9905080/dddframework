package response

//RespCode Enum
type RespCode struct {
	Num int
	Msg string
}

// RespCodeEnum編號
var (
	SUCCESS = RespCode{0, "Y"} // value --> 0
	FAIL    = RespCode{1, "N"} // value --> 1
)

// String fmt
func (respCodeEnum RespCode) String() string {
	return respCodeEnum.Msg
}
