package readwrite

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

func ReadWriteExample(redisClient redis.Conn) {
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
