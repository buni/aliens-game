package player

//go:generate mockgen -source=entity.go -destination=mock/player_mocks.go -package mock

// Player interface
type Player interface {
	Move(city string) bool
	Name() string
	City() string
	Destroy()
	SetCity(city string)
	IsDestroyed() bool
}

// Players - player slice type
type Players []Player

// GetNames - get the names of all players in the slice
func (p Players) GetNames() (names []string) {
	for _, player := range p {
		names = append(names, player.Name())
	}
	return
}

// Destroyed - check if all players in the slice are destroyed
func (p Players) Destroyed() bool {
	for _, player := range p {
		if !player.IsDestroyed() {
			return false
		}
	}
	return true
}

// Destroy - destroyed all players in the slice
func (p Players) Destroy() {
	for _, player := range p {
		player.Destroy()
	}
}
