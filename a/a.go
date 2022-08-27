package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	fmt.Print("Started at:", start)
	for i := 1; i < 50_000; i++ {
		http.Get("http://localhost:8081/")
	}
	finish := time.Now()
	fmt.Print("\nFinished at:", finish)
	testTime := finish.Sub(start)
	fmt.Print("\nTotal time:", testTime)
}
