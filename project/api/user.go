package api

import (
	"fmt"
	"net/http"
	models "project/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Users(c *gin.Context) {

	var user models.User

	result, err := user.Users()

	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}

func Add(c *gin.Context) {
	var user models.User
	c.ShouldBind(&user)
	fmt.Println("user", user.Username, user.Password)
	id, err := user.Insert()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "ok",
		"data":    id,
	})
}

func Del(c *gin.Context) {
	var user models.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	result, err := user.Destroy(id)
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "ok",
	})
}

func Update(c *gin.Context) {
	var user models.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	c.ShouldBind(&user)
	result, err := user.Update(id)
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "ok",
	})
}
