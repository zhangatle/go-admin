package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"go-admin/tools/config"
)

func TlsHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     config.ApplicationConfig.Domain,
		})
		err := secureMiddleware.Process(context.Writer, context.Request)
		if err != nil {
			return
		}
		context.Next()
	}
}
