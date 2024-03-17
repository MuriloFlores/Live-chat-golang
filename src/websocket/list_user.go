package websocket

import (
	"encoding/json"

	"github.com/MuriloFlores/Live-Chat/src/model"
	"github.com/gin-gonic/gin"
)

func ListUsers(c *gin.Context) {
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var response []*model.User

	for user := range ConnectedUsers {
		response = append(response, user)
	}

	json.NewEncoder(c.Writer).Encode(response)
}
