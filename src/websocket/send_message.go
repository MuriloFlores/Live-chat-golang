package websocket

import "nhooyr.io/websocket/wsjson"

func SendMessage() {
	for newMessage := range MessageChannel {
		for user := range ConnectedUsers {
			wsjson.Write(user.Context, user.Conn, newMessage)
		}
	}
}
