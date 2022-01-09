package model

import (
	"github.com/go-redis/redis/v7"
	"gmimo/push/config"
	"time"
)

var (
	RedisPool     *redis.Client
	defaultExpire time.Duration
)

func InitRedisClient(cf config.ConfRedis) error {
	if RedisPool == nil {
		opts := &redis.Options{
			Addr:        cf.HostPort,
			Password:    cf.Password,
			DB:          cf.Db,
			PoolSize:    cf.PoolSize,
			MaxRetries:  cf.MaxRetries,
			IdleTimeout: cf.IdleTimeout,
		}
		RedisPool = redis.NewClient(opts)
		defaultExpire = cf.Ttl
		return RedisPool.Ping().Err()
	}
	return nil
}
