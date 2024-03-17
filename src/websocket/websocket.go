package websocket

import (
	"log"

	"github.com/MuriloFlores/Live-Chat/src/model"
	"github.com/gin-gonic/gin"
	"nhooyr.io/websocket"
)

var (
	ConnectedUsers map[*model.User]bool = make(map[*model.User]bool)
	MessageChannel chan model.Message   = make(chan model.Message)
	JoinChannel    chan *model.User     = make(chan *model.User)
)

func WsHandler(c *gin.Context) {
	nickname, _ := c.Params.Get("nickname")

	connection, err := websocket.Accept(c.Writer, c.Request, &websocket.AcceptOptions{
		OriginPatterns: []string{"*"},
	})

	if err != nil {
		log.Fatal("Error to trying accept websocket connection: ", err.Error())
	}

	go SendMessage()
	go NewJoiner()

	NewUser := model.User{Nickname: nickname, Conn: connection, Context: c.Request.Context()}
	JoinChannel <- &NewUser

	ReadMessage(&NewUser)
}
