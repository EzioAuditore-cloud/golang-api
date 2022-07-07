package main

import (
	router "project/routes"
)

func main() {
	r := router.InnitRouter()
	r.Run(":8081")
}
