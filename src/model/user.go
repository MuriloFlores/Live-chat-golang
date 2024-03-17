package model

import (
	"context"

	"nhooyr.io/websocket"
)

type User struct {
	Nickname string          `json:"nickname"`
	Conn     *websocket.Conn 
	Context  context.Context 
}
