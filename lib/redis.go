package lib

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var redisPool *redis.Pool

func InitRedis() *redis.Pool {
	address := config.DatabaseConfig.RedisConfig.Host + ":" + config.DatabaseConfig.RedisConfig.Port
	password := redis.DialPassword(config.DatabaseConfig.RedisConfig.Password)
	conTimeout := redis.DialConnectTimeout(240 * time.Second)
	writeTimeout := redis.DialWriteTimeout(240 * time.Second)
	readTimeout := redis.DialReadTimeout(240 * time.Second)
	redisPool = &redis.Pool{
		MaxIdle:     config.DatabaseConfig.RedisConfig.MaxIdleConnections,
		MaxActive:   config.DatabaseConfig.RedisConfig.MaxOpenConnections,
		IdleTimeout: 240 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", address, password,
				readTimeout, writeTimeout, conTimeout)
			if err != nil {
				SendSlackMessage(Slack{
					Text: "REDIS: " + err.Error(),
				})
				return nil, err
			}
			return c, nil
		},
	}
	return redisPool
}

func GetRedisPool() *redis.Pool {
	return redisPool
}

func GetCache(key string) ([]byte, error) {
	conn := GetRedisPool().Get()
	prefix := "API:"

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", prefix+key))
	if err != nil {
		return data, err
	}
	return data, err
}

func SetCache(key string, value []byte, seconds float64) error {
	conn := GetRedisPool().Get()

	prefix := "API:"
	_, err := conn.Do("SET", prefix+key, value)
	if err != nil {
		v := string(value)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		return err
	}
	if seconds != 0 {
		conn.Do("EXPIRE", prefix+key, seconds)
	} else {
		conn.Do("DEL", prefix+key)
	}
	return err
}

func DeleteRedis(key string) error {
	conn := GetRedisPool().Get()
	prefix := "API:"
	_, err := conn.Do("DEL", prefix+key)
	return err
}

func RedisPing() error {
	conn := GetRedisPool().Get()
	_, err := conn.Do("PING")
	return err
}
