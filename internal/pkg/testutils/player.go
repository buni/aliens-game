package testutils

import "sync"

type NoopPlayer struct {
	name        string
	currentCity string
	destroyed   bool
	rw          sync.RWMutex
}

func NewNoopPlayer(name string) *NoopPlayer {
	return &NoopPlayer{
		name:        name,
		currentCity: "",
		destroyed:   false,
		rw:          sync.RWMutex{},
	}
}

func NewNoopPlayerWithCity(name string, cityName string) *NoopPlayer {
	return &NoopPlayer{
		name:        name,
		currentCity: cityName,
		destroyed:   false,
		rw:          sync.RWMutex{},
	}
}

func (c *NoopPlayer) Move(city string) bool {
	return true
}

// SetCity - sets the current player's city name
func (c *NoopPlayer) SetCity(city string) {
	c.rw.Lock()
	defer c.rw.Unlock()
	c.currentCity = city
}

// Name - returns the player name
func (c *NoopPlayer) Name() string {
	c.rw.RLock()
	defer c.rw.RUnlock()
	return c.name
}

// Destroy - marks the player as destroyed
func (c *NoopPlayer) Destroy() {
	c.rw.Lock()
	defer c.rw.Unlock()
	c.destroyed = true
}

// IsDestroyed - returns whether the player is destroyed or not
func (c *NoopPlayer) IsDestroyed() bool {
	c.rw.RLock()
	defer c.rw.RUnlock()
	return c.destroyed
}

// City - returns the city that the player resides in
// returns an empty string if there is none
func (c *NoopPlayer) City() string {
	c.rw.RLock()
	defer c.rw.RUnlock()
	return c.currentCity
}

// String - stringer interface impl
func (c *NoopPlayer) String() string {
	c.rw.RLock()
	defer c.rw.RUnlock()
	return c.Name()
}
