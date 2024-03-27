package logger

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type LogConfig struct {
	// logger   *logrus.Logger
	Path            string `yaml:"path"`
	TimestampFormat string `yaml:"timestampFormat"`
	// osFile   *os.File
}

var logger = logrus.New()

var logConfig LogConfig

func init() {
	configFile := "../middleWare/logger/config/log.yaml"
	dataBytes, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Errorf("log config init ReadFile err: %v", err)
	}
	err = yaml.Unmarshal(dataBytes, &logConfig)
	if err != nil {
		fmt.Errorf("log config init Unmarshal err: %v", err)
	}
	StructLog("Info", "log 配置成功！")
}

func CreateLogFolder(path string) {
	if err := os.MkdirAll(path, 0777); err != nil {
		fmt.Println("createLogFolder err")
		panic(err)
	}
}

func StructLog(logLevel, format string, args ...interface{}) {
	tim := time.Now().Format("2006-01-02")
	fileName := tim + ".log"
	f, err := os.OpenFile(logConfig.Path+"/"+fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	logger.SetFormatter(&logrus.TextFormatter{
		// ForceColors:               true,
		// EnvironmentOverrideColors: true,
		TimestampFormat: logConfig.TimestampFormat,
	})
	writers := []io.Writer{f, os.Stdout}
	outs := io.MultiWriter(writers...)
	logger.SetOutput(outs)
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
