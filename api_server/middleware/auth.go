package middleware

import (
	"dddframework/api_server/entity/service"
	ErrorCode "dddframework/api_server/error_code"
	Response "dddframework/api_server/response"
	"github.com/gin-gonic/gin"
	Logging "github.com/z9905080/gloger"
	"net/http"
)

// Auth 檢查驗證
func Auth(user service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {

		if userToken, getCookieErr := c.Cookie("user_token"); getCookieErr == nil {
			authResult, authErr := user.Auth(userToken)
			if authErr != nil {
				Logging.Error("auth Err:", authErr)
				c.AbortWithStatusJSON(http.StatusOK, Response.APIResponse{
					Code:    ErrorCode.AuthErr.NumString(),
					Status:  Response.FAIL.String(),
					Message: ErrorCode.AuthErr.String(),
				})
				return
			}
			c.Set("user_data", authResult)
			c.Next()
		} else {
			Logging.Error("Get Cookie Err:", getCookieErr)

			return
		}
	}
}
