package router

import (
	"project/api"
	"project/utils"

	"github.com/gin-gonic/gin"
)

func InnitRouter() *gin.Engine {
	gin.SetMode(utils.AppMode)
	router := gin.Default()
	// router.Use(middlewares.Cors())
	// router.LoadHTMLGlob("templates/*")
	// router.Static("/assets", "./assets")
	// router.GET("/", controller.GoIndex)
	// router.GET("/register", controller.GoRegister)
	// router.GET("/manager", controller.GoManager)
	// router.GET("/getuser", controller.GoGet)
	// router.GET("/deluser", controller.GoDel)
	// router.POST("/deluser", controller.GoDel)
	router.POST("/login", api.Login)
	router.GET("/users", api.Users)
	router.GET("/user/:searchUsername", api.GetUser)
	router.POST("/register", api.Add)
	router.POST("/user/:id", api.Update)
	router.DELETE("/user/:id", api.Del)
	return router
}
