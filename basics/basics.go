package main

type AppUser struct {
	name string
	age  string
}

type Message struct {
	id    int64
	data  string
	owner AppUser
}

func (msg *Message) wipeData() {
	msg.data = ""
}

func (msg Message) wipeDataNotWorking() {
	msg.data = ""
}

func main() {
	i := 0

	pi := &i

	println(pi)
	*pi = 1
	println(i)

	me := AppUser{
		name: "Jacek",
		age:  "23",
	}

	messages := []Message{}

	// create 10 more messages
	for i := 0; i < 10; i++ {
		message := Message{
			id:    int64(i),
			data:  "Hello World " + string(i) + "!",
			owner: me,
		}
		messages = append(messages, message)
	}

	println("\n\n\nMessages:")
	for _, msg := range messages {
		println(msg.data)
	}

	println("\nNot working wipe:")
	for _, msg := range messages {
		msg.wipeDataNotWorking()
		println(msg.data)
	}

	println("Wiped messages:")
	for _, msg := range messages {
		msg.wipeData()
		println(msg.data)
	}
}
