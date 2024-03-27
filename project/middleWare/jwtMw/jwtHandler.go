package jwtMw

import (
	"net/http"
	"os"
	"project/middleWare/logger"
	general "project/model/General.go"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type JwtConfig struct {
	Secret string `yaml: "Secret"`
}

type JwtClaims struct {
	User general.UserClient
	jwt.StandardClaims
}

var skipPathMap = map[string]bool{
	"/login":    true,
	"/register": true,
}

var jwtConfig JwtConfig

func init() {
	dataBytes, err := os.ReadFile("../middleWare/jwtMw/config/Jwt.yaml")
	if err != nil {
		logger.StructLog("Error", "db config init ReadFile err: %v", err)
	}
	err = yaml.Unmarshal(dataBytes, &jwtConfig)
	if err != nil {
		logger.StructLog("Error", "Jwt config init Unmarshal err: %v", err)
	}
	logger.StructLog("Info", "log 配置成功！")
}

func JwtAuthHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if _, ok := skipPathMap[ctx.FullPath()]; ok {
			ctx.Next()
			return
		}

		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"GetAuthorization Error": "Unauthorized"})
			return
		}
		// logger.StructLog("Info", "tokenString: %s", tokenString)
		token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtConfig.Secret), nil
		})
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"ParseWithClaims Error": "Unauthorized"})
			return
		}

		if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
			ctx.Set("token", token)
			ctx.Set("user", claims.User)
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"tokenValid Error": "Unauthorized"})
			return
		}
	}
}
