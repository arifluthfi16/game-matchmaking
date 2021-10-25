package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/arifluthfi16/game-matchmaking/controller"
)

type Match struct {}

func (p *Match) Route (route *gin.Engine){
	router := route.Group("/match")
	Controller := controller.MatchController{}

	router.POST("/join", Controller.JoinMatch)
}

