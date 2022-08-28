package main

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

const (
	redisProtocol = "tcp"
	redisUrl      = "192.168.67.2:30331"
)

func main() {
	redisClient := getRedisConnPool().Get()
	defer redisClient.Close()

	result, err := redisClient.Do("set", "name", "test name :)")

	if err != nil {
		log.Panicln(err)
	}

	log.Printf("Result for set name: [%s]", result)

	name, err := redisClient.Do("get", "name")

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Name from redis:[%s]", name)
}

func getRedisConnPool() *redis.Pool {
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			channel, err := redis.Dial(redisProtocol, redisUrl)

			if err != nil {
				log.Panic("OMG CAN'T CONNECT TO REDIS: " + err.Error())
			}

			return channel, err
		},
	}

	return pool
}
