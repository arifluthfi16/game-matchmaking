package room

import (
	"database/sql"
	"github.com/arifluthfi16/game-matchmaking/model"
)

func (s RoomService) CreateOne (ownerUsername string, title string, description string) (model.Room, error) {
	room := model.Room{
		Title:         title,
		Description:   sql.NullString{
			String: description,
		},
		Game:          "Warframe",
		MaxPlayer:     4,
		IsActive:      true,
		GameMode:      "Normal",
		MinLevel:      1,
		OwnerUsername: ownerUsername,
	}

	if err := s.DB.Create(&room).Error; err != nil {
		return room, err
	}
	return room, nil
}