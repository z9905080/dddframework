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
