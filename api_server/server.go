package api_server

import (
	"dddframework/api_server/controller"
	"dddframework/api_server/entity/adapter"
	"dddframework/api_server/entity/service"
	"dddframework/api_server/middleware"
	EnvVar "dddframework/environment/env_variable"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func Run() {
	r := gin.Default()

	fbAuthMW := service.UserService(
		&adapter.FacebookUserService{
			FacebookAccount: "FBA",
			FacebookToken:   "FBT",
		})

	googleAuthMW := service.UserService(
		&adapter.GoogleUserService{
			GoogleAccount: "googleA",
			GoogleToken:   "Ggg",
		})

	fbGroup := r.Group("fb", middleware.Auth(fbAuthMW))
	googleGroup := r.Group("google", middleware.Auth(googleAuthMW))

	fbGroup.GET("get", controller.Get)
	googleGroup.GET("get", controller.Get)

	// 設定server 參數
	s := &http.Server{
		Addr:           ":" + strconv.Itoa(EnvVar.Setting.Server.Port),
		Handler:        r,
		ReadTimeout:    15,
		WriteTimeout:   15,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println(s.ListenAndServe())
}
