package api_server

import (
	"dddframework/api_server/controller"
	"dddframework/api_server/entity/adapter"
	"dddframework/api_server/entity/service"
	"dddframework/api_server/middleware"
	WsHandler "dddframework/api_server/ws_handler"
	EnvVar "dddframework/environment/env_variable"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	Logging "github.com/z9905080/gloger"
	"github.com/z9905080/melody"
	"log"
	"net/http"
	"strconv"
)

func Run() {

	wsEngine := melody.New(melody.DialChannelBufferSize(500))
	wsEngine.Config.MaxMessageSize = 0

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:    []string{"*"},
	}))

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

	r.GET("ws", func(c *gin.Context) {
		sessionDataMap := make(map[string]interface{}, 0)
		sessionDataMap["client_ip"] = c.ClientIP()
		wsEngine.HandleRequestWithKeys(c.Writer, c.Request, sessionDataMap)
	})

	wsEngine.HandleMessage(func(s *melody.Session, msg []byte) {
		WsHandler.ClientRequestHandler(wsEngine, s, msg)
	})

	// 設定server 參數
	s := &http.Server{
		Addr:           ":" + strconv.Itoa(EnvVar.Setting.Server.Port),
		Handler:        r,
		ReadTimeout:    15,
		WriteTimeout:   15,
		MaxHeaderBytes: 1 << 20,
	}

	Logging.Force("目前啟動的Port:", EnvVar.Setting.Server.Port)

	log.Println(s.ListenAndServe())
}
