package router

import (
	"fmt"
	"net/http"
	"project/api"
	_ "project/api"
	"project/app/kafkaMQ"
	"project/app/manage"
	"project/config"
	"project/middleWare/jwtMw"
	"project/middleWare/logger"
	general "project/model/General.go"

	"github.com/gin-contrib/pprof"

	"github.com/gin-gonic/gin"
)

func InnitRouter() *gin.Engine {
	go manage.Srv.ListenMessage()
	go kafkaMQ.Consumer()
	logger.StructLog("Info", "服务，启动！")
	router := gin.Default()
	pprof.Register(router, "debug/pprof")
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
	fmt.Println("config.GlobleConf.ListenIPAndPort.Port", config.GlobleConf.ListenIPAndPort.Port)
	router.Run(config.GlobleConf.ListenIPAndPort.Port)
	return router
}
