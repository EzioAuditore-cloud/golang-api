package router

import (
	"net/http"
	_ "project/api"
	"project/config"
	"project/middleWare/jwtMw"
	"project/middleWare/logger"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func InnitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(logger.LoggerHandler())
	router.Use(jwtMw.JwtAuthHandler())
	// router.GET("/users", Users)
	// router.POST("/user", Add)
	// router.PUT("/user/:id", Update)
	// router.DELETE("/user/:id", Del)
	router.GET("/login", func(ctx *gin.Context) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtMw.JwtClaims{
			ID:       123,
			UserName: "Ezio",
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			},
		})
		tokenString, err := token.SignedString([]byte(config.GlobleConf.JwtSecret))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot Create Token"})
			logger.StructLog("Error", "Login Cannot Create Token Err: %v", err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
	})
	router.GET("/hello", func(c *gin.Context) {
		id, _ := c.Get("id")
		userName, _ := c.Get("user_name")
		c.JSON(http.StatusOK, gin.H{
			"msg":       "你好",
			"id":        id,
			"user_name": userName,
		})
	})
	router.Run(":8080")
	return router
}
