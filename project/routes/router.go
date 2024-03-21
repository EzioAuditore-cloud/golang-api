package router

import (
	"net/http"
	_ "project/api"
	"project/middleWare/logger"

	"github.com/gin-gonic/gin"
)

func InnitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(logger.LoggerHandler())
	// router.GET("/users", Users)
	// router.POST("/user", Add)
	// router.PUT("/user/:id", Update)
	// router.DELETE("/user/:id", Del)
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "你好",
		})
	})
	router.Run(":8080")
	return router
}
