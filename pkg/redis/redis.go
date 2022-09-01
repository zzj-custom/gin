package redis

import (
	"github.com/garyburd/redigo/redis"
	"strconv"
	"sync"
	"time"
)

var (
	redisPool *redis.Pool
	redisOnce sync.Once
)

func (config *DialConfig) getDialOption() []redis.DialOption {
	dialOptions := make([]redis.DialOption, 0, 3)
	dialOptions = append(dialOptions,
		redis.DialConnectTimeout(config.ConnectTimeout),
		redis.DialReadTimeout(config.ReadTimeout),
		redis.DialDatabase(config.Database),
	)
	if config.Password != "" {
		dialOptions = append(dialOptions, redis.DialPassword(config.Password))
	}

	return dialOptions
}

func InitPool(config *DialConfig) *redis.Pool {
	redisOnce.Do(func() {
		redisPool = NewPool(config)
	})
	return redisPool
}

func Pool() *redis.Pool {
	if redisPool == nil {
		panic("redis连接池未初始化")
	}
	return redisPool
}

func NewPool(config *DialConfig) *redis.Pool {
	if config == nil {
		panic("配置文件为空")
	}
	pool := redis.Pool{
		Dial: func() (redis.Conn, error) {
			dial, err := redis.Dial(
				"tcp",
				config.Host+":"+strconv.Itoa(config.Port),
				config.getDialOption()...,
			)
			if err != nil {
				return nil, err
			}
			_, err = dial.Do("SELECT", config.Database)
			if err != nil {
				return nil, err
			}
			return dial, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:         config.MaxIdle,
		MaxActive:       config.MaxActive,
		IdleTimeout:     time.Duration(config.IdleTimeout),
		Wait:            config.Wait,
		MaxConnLifetime: time.Duration(config.MaxConnLifetime),
	}
	return &pool
}
