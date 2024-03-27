package router

import (
	"net/http"
	"project/api"
	_ "project/api"
	"project/app/server"
	"project/middleWare/jwtMw"
	"project/middleWare/logger"
	general "project/model/General.go"

	"github.com/gin-gonic/gin"
)

func InnitRouter() *gin.Engine {
	go server.ServerStart()
	router := gin.Default()
	router.Use(logger.LoggerHandler())
	router.Use(jwtMw.JwtAuthHandler())
	// router.GET("/users", Users)
	// router.POST("/user", Add)
	// router.PUT("/user/:id", Update)
	// router.DELETE("/user/:id", Del)
	router.POST("/login", api.Login)
	router.GET("/conn", api.ConnectServer)
	router.GET("/hello", func(c *gin.Context) {
		user, _ := c.Get("user")
		userClient := user.(general.UserClient)
		c.JSON(http.StatusOK, gin.H{
			"msg":  "你好",
			"user": userClient,
		})
	})
	router.Run(":8080")
	return router
}
