package controller

import "github.com/gin-gonic/gin"

func GoIndex(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func GoRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func GoManager(c *gin.Context) {
	c.HTML(200, "manager.html", nil)
}

func GoGet(c *gin.Context) {
	c.HTML(200, "getuser.html", nil)
}

func GoDel(c *gin.Context) {
	c.HTML(200, "deluser.html", nil)
}
