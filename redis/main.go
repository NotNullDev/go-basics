package main

import (
	"fmt"

	"github.com/notnulldev.com/go-example/redis/pubsub"
)

func main() {
	// readwrite.ReadWriteExample(redisClient)
	go pubsub.StartSubscriber("dummy")

	go pubsub.StartPublisher("dummy", 100)

	fmt.Scanln()
}
