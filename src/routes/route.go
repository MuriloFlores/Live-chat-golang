package routes

import (
	"github.com/MuriloFlores/Live-Chat/src/websocket"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/chat/:nickname", websocket.WsHandler)
	r.GET("/listUsers", websocket.ListUsers)
}