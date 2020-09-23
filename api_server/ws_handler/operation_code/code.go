package operation_code

// OperationCode Enum
type OperationCode string

const (
	C2S_Login OperationCode = "C2S_Login"
	S2C_Login OperationCode = "S2C_Login"
)

// String fmt
func (funcCodeEnum OperationCode) String() string {
	return string(funcCodeEnum)
}
