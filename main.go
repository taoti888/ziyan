package main

import (
	"github.com/gin-gonic/gin"
	"github.com/taoti888/ziyan/core"
	"github.com/taoti888/ziyan/global"
	"github.com/taoti888/ziyan/initialize"
	"go.uber.org/zap"
)

func main() {
	// 生产环境启用ReleaseMode
	gin.SetMode(gin.DebugMode)

	// 初始化Viper
	global.GVA_VP = core.Viper()

	// 初始化zap日志库
	global.GVA_LOG = core.Zap()
	zap.ReplaceGlobals(global.GVA_LOG)

	// gorm连接数据库,注意程序结束前关闭数据库链接
	global.GVA_DB = initialize.Gorm()
	db, _ := global.GVA_DB.DB()
	defer db.Close()

	// 运行程序
	core.RunWindowsServer()
}
