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
	// c.HTML(200, "manager.html", result)
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}

func GetUser(c *gin.Context) {

	result, err := models.GetUsers(c.Param("searchUsername"))
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "error",
		})
		return
	}
	// c.HTML(200, "manager.html", result)
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}

func Login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -3,
			"message": "invalid bind",
		})
	}
	// username := c.PostForm("username")
	// password := c.PostForm("password")
	// user := models.User{Password: password, Username: username}
	fmt.Println("user ", user.Username, user.Password)
	u, err1 := user.FindUser(user.Username)
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -4,
			"message": "find user error",
		})
	}
	fmt.Println("u ", u.Username, u.Password)
	if u.Username == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "no found user",
		})
		// c.HTML(200, "index.html", "no found user")
	} else {
		if u.Password != user.Password {
			c.JSON(http.StatusOK, gin.H{
				"code":     -2,
				"message":  "Wrong password",
				"username": u.Username,
			})
			// c.HTML(200, "index.html", "Wrong password")
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":     1,
				"message":  "ok",
				"username": u.Username,
				"password": u.Password,
			})
			// c.HTML(200, "manager.html", nil)
		}
	}

}

func Add(c *gin.Context) {
	// username := c.PostForm("username")
	// password := c.PostForm("password")
	// user := models.User{Password: password, Username: username}

	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -3,
			"message": "invalid bind",
		})
	}
	_, err1 := user.Insert()
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "insert error",
		})
		// c.HTML(200, "register.html", "error")
	}
	c.JSON(http.StatusOK, gin.H{
		"code":     1,
		"message":  "ok",
		"data":     user.ID,
		"username": user.Username,
	})
	// c.HTML(200, "register.html", "ok")

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
	err1 := c.ShouldBindJSON(&user)
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -3,
			"message": "invalid bind",
		})
	}
	fmt.Printf("user: %v\n", user)
	result, err := user.Update(id)
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "update error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "ok",
	})
}
