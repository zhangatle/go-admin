package middleware

import "github.com/gin-gonic/gin"

func InitMiddleware(r *gin.Engine) {
	// 日志处理
	r.Use(LoggerToFile())
	// 自定义错误处理
	r.Use(CustomError)
	r.Use(NoCache)
	// 跨域处理
	r.Use(Options)
	r.Use(Secure)
	r.Use(RequestId())
}
