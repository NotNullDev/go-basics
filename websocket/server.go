package main

type ChatUser struct {
	id string
}

type ChatMessage struct {
	Id      string
	Author  ChatUser
	Content string
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
		id:               uuid.,
		newMessage:       make(chan ChatMessage),
		connectClient:    make(chan ChatUser),
		disconnectClient: make(chan ChatUser),
	}
}

func main() {
	createServer()
}
