package services

import (
	"github.com/arifluthfi16/game-matchmaking/services/match"
	"github.com/arifluthfi16/game-matchmaking/services/player"
	"github.com/arifluthfi16/game-matchmaking/services/room"
	"gorm.io/gorm"
)

var (
	PlayerService 	player.PlayerService
	RoomService		room.RoomService
	MatchService	match.MatchService
)

func InjectDBIntoServices (db *gorm.DB) {
	PlayerService.DB 	= db
	RoomService.DB		= db
	MatchService.DB		= db
}