package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/arifluthfi16/game-matchmaking/controller"
)

type Player struct {}

func (p *Player) Route (route *gin.Engine){
	router := route.Group("/player")
	Controller := controller.PlayerController{}

	router.POST("/", Controller.CreatePlayer)
}

