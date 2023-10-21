package redis

import (
	"github.com/megadata-dev/go-snmp-olt-zte-c320/config"
	"github.com/redis/go-redis/v9"
	"time"
)

func NewRedisClient(cfg *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:         cfg.RedisCfg.Host + ":" + cfg.RedisCfg.Port,
		Password:     cfg.RedisCfg.Password,
		DB:           cfg.RedisCfg.DB,
		MinIdleConns: cfg.RedisCfg.MinIdleConnections,
		PoolSize:     cfg.RedisCfg.PoolSize,
		PoolTimeout:  time.Duration(cfg.RedisCfg.PoolTimeout) * time.Second,
	})
}
