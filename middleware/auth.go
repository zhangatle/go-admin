package middleware

import (
	"go-admin/handler"
	"go-admin/pkg/jwtauth"
	"go-admin/tools/config"
	"time"
)

func AuthInit() (*jwtauth.GinJWTMiddleware, error) {
	return jwtauth.New(&jwtauth.GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte(config.ApplicationConfig.JwtSecret),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		PayloadFunc:     handler.PayloadFunc,
		IdentityHandler: handler.IdentityHandler,
		Authenticator:   handler.Authenticator,
		Authorizator:    handler.Authorizator,
		Unauthorized:    handler.Unauthorized,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})
}
