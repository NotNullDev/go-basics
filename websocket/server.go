package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type ChatUser struct {
	id string
}

type ChatMessage struct {
	Id      string
	Author  ChatUser
	Content string
	Channel string
}

type Server struct {
	id               string
	newMessage       chan ChatMessage
	connectClient    chan ChatUser
	disconnectClient chan ChatUser
}

func (server *Server) Run() {
	for {
		select {
		case user := <-server.connectClient:
			println("Connected client with id " + user.id)
		case user := <-server.disconnectClient:
			println("Disconnected client with id " + user.id)
		case message := <-server.newMessage:
			println("Received message from " + message.Author.id + " with content: " + message.Content)
		}

	}
}

func createServer() *Server {
	return &Server{
		id:               uuid.NewString(),
		newMessage:       make(chan ChatMessage),
		connectClient:    make(chan ChatUser),
		disconnectClient: make(chan ChatUser),
	}
}

var addr = flag.String("a", "localhost:7777", "address to server websocket")
var endpoint = flag.String("e", "/server", "endpoint ot serve websocket connection")

func main() {
	logFile, err := os.Create("log.log")

	if err != nil {
		println(err.Error() + "\nCan't create log file!")
	} else {
		log.SetOutput(logFile)
	}

	server := createServer()

	log.Println("addr:", *addr)
	log.Println("endpoint:", *endpoint)

	go server.Run()

	url := url.URL{
		Scheme: "ws",
		Host:   *addr,
		Path:   *endpoint,
	}

	log.Println("URL STRING: ", url.String())

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	http.HandleFunc(*endpoint, func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Starting server on %v", *addr)
		wsConnection, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			log.Println(err.Error())
			return
		}

		defer wsConnection.Close()

		for {
			msgType, msgContent, err := wsConnection.ReadMessage()

			log.Printf("New connection!")

			if err != nil {
				log.Fatalln(err.Error())
				return
			}

			log.Printf("Received new message with type: [%v] and content: [%v]", msgType, string(msgContent))
		}
	})

	err = http.ListenAndServe(*addr, nil)

	if err != nil {
		log.Fatalln(err.Error())
	}
}
