package routers

import "github.com/gin-gonic/gin"

type RouterInterface interface {
	Route(*gin.Engine)
}

type RouteLoader struct{}

func (loader RouteLoader) LoadRoutes() []RouterInterface {
	player 	:= new (Player)
	room 	:= new (Room)
	match 	:= new (Match)
	return []RouterInterface{
		player,
		room,
		match,
	}
}