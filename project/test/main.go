package main

import (
	"project/middleWare/logger"
	router "project/routes"
)

func init() {
	logger.CreateLogFolder("../logs")
}
func main() {
	r := router.InnitRouter()
	r.Run(":8081")
}
