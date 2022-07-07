package router

import (
	. "project/api"

	"github.com/gin-gonic/gin"
)

func InnitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/users", Users)
	router.POST("/user", Add)
	router.PUT("/user/:id", Update)
	router.DELETE("/user/:id", Del)
	return router
}
