package system

import (
	"github.com/taoti888/ziyan/config"
)

// 配置文件结构体
type System struct {
	Config config.Server `json:"config"`
}
