package api_server

import (
	"dddframework/api_server/router"

	WsHandler "dddframework/api_server/ws_handler"
	EnvVar "dddframework/environment/env_variable"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	Logging "github.com/z9905080/gloger"
	"github.com/z9905080/melody"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Run() {

	wsEngine := melody.New(melody.DialChannelBufferSize(500))
	wsEngine.Config.MaxMessageSize = 0

	r := gin.Default()

	setRootMiddleWare(r)

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:    []string{"*"},
	}))

	router.SetRoute(r)

	r.GET("ws", func(c *gin.Context) {
		sessionDataMap := make(map[string]interface{}, 0)
		sessionDataMap["client_ip"] = c.ClientIP()
		wsEngine.HandleRequestWithKeys(c.Writer, c.Request, sessionDataMap)
	})

	wsEngine.HandleMessage(func(s *melody.Session, msg []byte) {
		WsHandler.ClientRequestHandler(wsEngine, s, msg)
	})

	wsEngine.HandleConnect(func(s *melody.Session) {
		s.AddSub("top")
	})

	//go func() {
	//	hash := uuid.NewV4()
	//	for {
	//		wsEngine.PubTextMsg([]byte(hash.String()), true, "top")
	//		time.Sleep(5 * time.Second)
	//	}
	//}()

	// 設定server 參數
	s := &http.Server{
		Addr:           ":" + strconv.Itoa(EnvVar.Setting.Server.Port),
		Handler:        r,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	Logging.Force("目前啟動的Port:", EnvVar.Setting.Server.Port)

	log.Println(s.ListenAndServe())
}

// setRootMiddleWare 共用的middleware
func setRootMiddleWare(server *gin.Engine) {
	// gzip 壓縮response
	server.Use(gzip.Gzip(gzip.BestCompression))
}
