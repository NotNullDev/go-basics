package pubsub

import (
	"fmt"
	"log"
	"time"

	"github.com/notnulldev.com/go-example/redis/connection"
)

func StartPublisher(channelName string, sleepTimeInMs int64) {
	redisClient := connection.GetRedisConnPool().Get()
	defer func() {
		log.Println("Closing redis subscriber connection.")
		redisClient.Close()
	}()

	var i = 0
	for {
		msg := fmt.Sprintf("Message %d", i)
		i++

		err := redisClient.Send("publish", channelName, msg)

		if err != nil {
			log.Printf("ERROR: %s", err.Error())
		}

		redisClient.Flush()
		time.Sleep(time.Duration(sleepTimeInMs) * time.Millisecond)
	}
}
