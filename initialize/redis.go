package initialize

import (
	"context"
	"github.com/taoti888/ziyan/global"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.GVA_CONFIG.Redis
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        redisCfg.Addrs,
		Username:     redisCfg.Username,
		Password:     redisCfg.Password,
		PoolSize:     redisCfg.PoolSize,
		MaxRedirects: redisCfg.MaxRedirects,
		MinIdleConns: redisCfg.MinIdleConns,
		IdleTimeout:  redisCfg.IdleTimeout,
		DialTimeout:  redisCfg.DialTimeout,
		ReadTimeout:  redisCfg.ReadTimeout,
		WriteTimeout: redisCfg.WriteTimeout,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.GVA_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.GVA_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.GVA_REDIS = client
	}
}
