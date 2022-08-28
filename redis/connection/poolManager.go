package connection

import (
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/notnulldev.com/go-example/redis/globals"
)

func GetRedisConnPool() *redis.Pool {
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			channel, err := redis.Dial(globals.RedisProtocol, globals.RedisUrl)

			if err != nil {
				log.Panic("OMG CAN'T CONNECT TO REDIS: " + err.Error())
			}

			return channel, err
		},
	}

	return pool
}
