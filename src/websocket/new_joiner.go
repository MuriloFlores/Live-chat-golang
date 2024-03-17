package websocket

import (
	"time"

	"github.com/MuriloFlores/Live-Chat/src/model"
)

func NewJoiner() {

	for user := range JoinChannel {
		ConnectedUsers[user] = true

		MessageChannel <- model.Message{From: user.Nickname, Content: "the user " + user.Nickname + " has joined in the channel", SentAt: time.Now().Format("02-01-2006 15:04:05")}
	}
}
