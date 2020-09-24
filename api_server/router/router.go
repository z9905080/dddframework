package router

import (
	Cmd "dddframework/api_server/controller/cmd"
	"dddframework/api_server/entity/adapter"
	Middleware "dddframework/api_server/middleware"
	"github.com/gin-gonic/gin"
	GinDIGroup "github.com/z9905080/gin-di-router"
)

type router struct {
	EntityServer *gin.Engine
}

// SetGroup 設定url prefix (相同prefix 一同為Group)
func (r *router) SetGroup(prefix string) *gin.RouterGroup {
	group := r.EntityServer.Group(prefix)
	return group
}

// SetRoute 設定Route
func SetRoute(server *gin.Engine) {
	// 因為在這個模組會一直使用到Server，直接傳Reference到global var
	// 由於Server的Life Circle與Program是一樣的，所以不會有影響
	// serverRef = server
	routerEntity := router{
		EntityServer: server,
	}

	{
		// Google API設定
		GoogleSettingGroup := routerEntity.EntityServer.Group("", Middleware.Auth(new(adapter.GoogleUserService)))
		// API
		registry := GinDIGroup.New(GoogleSettingGroup)
		registry.RegisterWithGroup(new(Cmd.CmdController), GoogleSettingGroup)
	}

	{
		// Google API設定
		FbSettingGroup := routerEntity.EntityServer.Group("", Middleware.Auth(&adapter.FacebookUserService{FacebookToken: "FBT"}))
		// API
		registry := GinDIGroup.New(FbSettingGroup)
		registry.RegisterWithGroup(new(Cmd.CmdController), FbSettingGroup)

	}

}
