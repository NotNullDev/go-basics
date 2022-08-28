package pubsub

import (
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/notnulldev.com/go-example/redis/connection"
)

func StartSubscriber(channelName string) {
	redisClient := connection.GetRedisConnPool().Get()
	defer func() {
		log.Println("Closing redis subscriber connection.")
		redisClient.Close()
	}()
	// for {
	// 	newMessage, err := redisClient.Receive()
	// 	if err != nil {
	// 		log.Printf("Error! [%v]", err.Error())
	// 	} else {
	// 		log.Printf("Received message: [%s]", newMessage)
	// 	}
	// }

	pubSub := redis.PubSubConn{
		Conn: redisClient,
	}

	pubSub.Subscribe(channelName)

	for {
		switch event := pubSub.Receive().(type) {
		case redis.Message:
			log.Printf("Received message: [%s] from channel: [%s]", event.Data, event.Channel)
		case redis.Subscription:
			log.Printf("Subscription has been created, channel: %s, kind: %s, count?: %d", event.Channel, event.Kind, event.Count)
		case redis.Error:
			log.Printf("ERROR: %s", event.Error())
		default:
			log.Panic("Something went wrong!")
		}
	}

}
