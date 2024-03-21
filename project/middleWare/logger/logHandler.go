package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LogOut struct {
	// logger   *logrus.Logger
	fileName string
	// osFile   *os.File
}

var logger = logrus.New()

var dayLog LogOut

func CreateLogFolder(path string) {
	if err := os.MkdirAll(path, 0777); err != nil {
		fmt.Println("createLogFolder err")
		panic(err)
	}
}

func StructLog(logLevel, format string, args ...interface{}) {
	tim := time.Now().Format("2006-01-02")
	dayLog.fileName = tim + ".log"
	f, err := os.OpenFile("../logs/"+dayLog.fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	logger.Out = f
	logger.SetFormatter(&logrus.TextFormatter{
		// ForceColors:               true,
		EnvironmentOverrideColors: true,
		TimestampFormat:           "2006-01-02 15:04:05",
	})
	switch logLevel {
	case "Info":
		logger.SetLevel(logrus.InfoLevel)
		logger.Infof(format, args...)
	case "Debug":
		logger.SetLevel(logrus.DebugLevel)
		logger.Debugf(format, args...)
	case "Warn":
		logger.SetLevel(logrus.WarnLevel)
		logger.Warnf(format, args...)
	case "Error":
		logger.SetLevel(logrus.ErrorLevel)
		logger.Errorf(format, args...)
	}
}

func LoggerHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		endTime := time.Now()
		handleTime := endTime.Sub(startTime)
		reqMethod := ctx.Request.Method
		reqUri := ctx.Request.RequestURI
		status := ctx.Writer.Status()
		reqIP := ctx.ClientIP()
		StructLog("Info", "| %d | %v | %s | %s | %s |", status, handleTime, reqIP, reqMethod, reqUri)
	}
}
