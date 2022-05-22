package board

import (
	"github.com/buni/aliens-game/internal/pkg/board/city"
	"github.com/buni/aliens-game/internal/pkg/player"
)

//go:generate mockgen -source=entity.go -destination=mock/board_state_mocks.go -package mock

// State board state interface
type State interface {
	DeleteCity(cityName string)
	DeleteCityAndLinks(cityName string)
	MoveVisitor(cityName string, visitor player.Player) bool
	GetCity(cityName string) (city *city.City, ok bool)
	GetNextDirection(currentCity string) (string, bool)
	GetCities() (cities []*city.City)
}
