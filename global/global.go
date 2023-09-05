package global

import (
	"github.com/go-redis/redis/v8"
	"sync"

	"github.com/taoti888/ziyan/config"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB                  *gorm.DB
	GVA_REDIS               *redis.ClusterClient
	GVA_CONFIG              config.Server
	GVA_VP                  *viper.Viper
	GVA_LOG                 *zap.Logger
	GVA_Concurrency_Control = &singleflight.Group{}

	lock sync.RWMutex
)
