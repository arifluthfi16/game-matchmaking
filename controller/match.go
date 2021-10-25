package controller

import (
	"fmt"
	"github.com/arifluthfi16/game-matchmaking/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MatchController struct {}

type createMatchInput struct {
	Username 	string 	`json:"username" binding:"required"`
	RoomID		int		`json:"room_id" binding:"required"`
}

func (controller MatchController) JoinMatch (c *gin.Context) {
	// * Input Validation
	var input createMatchInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var response = ErrorResponse{
			Msg: "Validation Error",
			Err: err,
		}
		c.AbortWithStatusJSON(http.StatusOK, response)
		return
	}

	// * PLayer Join Room
	isPlayerJoined, err := services.MatchService.FindUserInRoom(input.Username, uint(input.RoomID))
	if  err != nil {
		var response = ErrorResponse{
			Msg: "Failed to check player",
			Err: err,
		}
		c.AbortWithStatusJSON(http.StatusOK, response)
		return
	}

	fmt.Println(isPlayerJoined)
	match, err := services.MatchService.PlayerJoinRoom(input.Username, uint(input.RoomID))
	if  err != nil {
		var response = ErrorResponse{
			Msg: "Failed to Create Player",
			Err: err,
		}
		c.AbortWithStatusJSON(http.StatusOK, response)
		return
	}

	var responseString = fmt.Sprintf("Successfully joined room %s", string(match.RoomID))
	var response = SuccessResponse{
		Status: true,
		Msg:    responseString,
	}
	c.AbortWithStatusJSON(http.StatusOK, response)
}