package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/arifluthfi16/game-matchmaking/controller"
)

type Room struct {}

func (p *Room) Route (route *gin.Engine){
	router := route.Group("/room")
	Controller := controller.RoomController{}

	router.POST("/", Controller.CreateRoom)
}

