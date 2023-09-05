package core

import (
	"fmt"
	"github.com/taoti888/ziyan/service/system"
	"go.uber.org/zap"
	"strings"
	"time"

	"github.com/taoti888/ziyan/global"
	"github.com/taoti888/ziyan/initialize"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()

	port := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Port)
	s := initServer(port, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("Server run success on ", zap.Any("port", strings.Split(port, ":")[1]))
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
