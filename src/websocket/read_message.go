package websocket

import (
	"log"
	"time"

	"github.com/MuriloFlores/Live-Chat/src/model"
	"nhooyr.io/websocket/wsjson"
)

func ReadMessage(newUser *model.User) {
	for {
		var messageFromClient model.Message

		err := wsjson.Read(newUser.Context, newUser.Conn, &messageFromClient)
		if err != nil {
			log.Println("User " + newUser.Nickname + " was desconnected")
			delete(ConnectedUsers, newUser)
			break
		}

		MessageChannel <- model.Message{From: messageFromClient.From, Content: messageFromClient.Content, SentAt: time.Now().Format("02-01-2006 15:04:05")}
	}
}