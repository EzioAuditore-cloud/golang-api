package api

import (
	"net/http"
	"project/app/manage"
	"project/config"
	"project/middleWare/jwtMw"
	"project/middleWare/logger"
	general "project/model/General.go"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func Login(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.PostForm("id"))
	if err != nil {
		logger.StructLog("Error", "Login ID String to Int err: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Login ID String to Int Err"})
		return
	}
	userName := ctx.PostForm("userName")
	user := general.UserClient{
		ID:   id,
		Name: userName,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtMw.JwtClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	})
	tokenString, err := token.SignedString([]byte(config.GlobleConf.JwtSecret))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot Create Token"})
		logger.StructLog("Error", "Login Cannot Create Token Err: %v", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func ConnectServer(ctx *gin.Context) {
	// wsUpgrader := &websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}}
	wsConn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Upgrade WebSocket Error"})
		logger.StructLog("Error", "Websocket Upgrade Error: %v", err)
		return
	}
	user, _ := ctx.Get("user")
	userClient := user.(general.UserClient)
	go manage.Srv.Handler(wsConn, userClient)
}

// func SendMsg(ctx *gin.Context) {
// 	wsUpgrader := websocket.Upgrader{}
// 	ws, err := wsUpgrader.Upgrade(ctx.Writer, ctx.Request, nil)
// 	if err != nil {
// 		logger.StructLog("Error", "Websocket Upgrade Error: %v", err)
// 		return
// 	}
// 	defer ws.Close()

// }
