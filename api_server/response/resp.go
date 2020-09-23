package response

// APIResponse api的回傳格式
type APIResponse struct {
	// API狀態
	Status string `json:"status" example:"Y"`
	// API錯誤碼
	Code string `json:"code" example:"0"`
	// API訊息
	Message string `json:"message" example:"回傳成功"`
	// API回傳資料
	Data interface{} `json:"data"`
}

type WSResponse struct {
	// WS狀態
	Status string `json:"status" example:"Y"`
	// WS錯誤碼
	ErrorCode int `json:"error_code" example:"0"`
	// WS訊息
	Message string `json:"message" example:"回傳成功"`
	// WS回傳資料
	Data interface{} `json:"data"`
	// WS FuncCode
	OperationCode string `json:"operation_code"`
}
