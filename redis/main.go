package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/notnulldev.com/go-example/redis/pubsub"
)

func main() {
	// readwrite.ReadWriteExample(redisClient)

	logFile, err := os.Create("log.log") // ~6700 req/s

	if err != nil {
		println(err)
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(os.Stdout, logFile)

	log.SetOutput(multiWriter)

	before := time.Now()

	defer func() {
		after := time.Now()
		diff := after.Sub(before)
		log.Printf("Time passed: %v", diff)
	}()

	go pubsub.StartSubscriber("dummy")

	go pubsub.StartPublisher("dummy", 0)

	fmt.Scanln()
}
