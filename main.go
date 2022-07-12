package main

import (
	router "project/routes"
	"project/utils"
)

func main() {
	r := router.InnitRouter()
	r.Run(utils.HttpPort)
}
