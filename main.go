package main

import (
	"github.com/ChatRoom1234/ChatChatChat/chatroom"
	"github.com/ChatRoom1234/ChatChatChat/server"
)

func main() {
	room := chatroom.NewRoom()
	go room.Run()

	err := server.NewServer(room)
	if err != nil {
		panic(err)
	}
}
