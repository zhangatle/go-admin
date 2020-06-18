package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/handler"
	"go-admin/middleware"
	"go-admin/tools"
	config2 "go-admin/tools/config"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	if config2.ApplicationConfig.IsHttps {
		r.Use(handler.TlsHandler())
	}
	middleware.InitMiddleware(r)
	authMiddleware, err := middleware.AuthInit()
	tools.HasError(err, "JWT Init Error", 500)
	// 注册系统路由
	InitSysRouter(r, authMiddleware)
	// 注册业务路由
	InitExamplesRouter(r, authMiddleware)
	return r
}
