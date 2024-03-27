package server

import (
	"project/app/manage"
	"project/middleWare/logger"
)

func ServerStart() {
	logger.StructLog("Info", "服务，启动！")
	manage.Srv.ListenMessage()
}
