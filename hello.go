package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	printHello("HI!\n")
	fmt.Print(getRandomNumber())
	startWebServer()

	// router := gin.Default()
	router := gin.New()

	router.GET("/", func(ctx *gin.Context) {
		body, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			fmt.Println(err)
		}
		ctx.JSON(200, gin.H{
			"id":  0,
			"msg": body,
		})
	})

	router.Run(":8081")

	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	userInput, _, _ := reader.ReadLine()

	fmt.Print(userInput)
}

func printHello(input string) {
	fmt.Printf("%s", input)
}

func getRandomNumber() int32 {
	rand.Seed(time.Now().UnixMilli())
	return rand.Int31()
}

func startWebServer() {

}
