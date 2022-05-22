package computer

import (
	"log"
	"sync"

	"github.com/buni/aliens-game/internal/pkg/board"
)

type Computer struct {
	moves       int
	name        string
	currentCity string
	board       board.State
	destroyed   bool
	rw          *sync.RWMutex
}

// NewPlayer - new computer player instance
func NewPlayer(board board.State, name string) *Computer {
	return &Computer{board: board, name: name, rw: &sync.RWMutex{}} // TODO: replace board
}

// Move - the player on the map
// in this implementation a random move is made if possible
// the function input is ignored
func (c *Computer) Move(city string) bool {
	if c.IsDestroyed() {
		return false
	}

	nextCityName, ok := c.board.GetNextDirection(c.City())
	if !ok {
		c.rw.RLock() // TODO: remove
		defer c.rw.RUnlock()
		log.Println("couldn't get next dir", c.currentCity, c.name)
		log.Println(c.board)
		return false
	}

	if c.board.MoveVisitor(nextCityName, c) {
		c.SetCity(nextCityName)
		return true
	}

	return false
}

// SetCity - sets the current player's city name
func (c *Computer) SetCity(city string) {
	c.rw.Lock()
	defer c.rw.Unlock()
	c.currentCity = city
}

// Name - returns the player name
func (c *Computer) Name() string {
	c.rw.RLock()
	defer c.rw.RUnlock()
	return c.name
}

// Destroy - marks the player as destroyed
func (c *Computer) Destroy() {
	c.rw.Lock()
	defer c.rw.Unlock()
	c.destroyed = true
}

// IsDestroyed - returns whether the player is destroyed or not
func (c *Computer) IsDestroyed() bool {
	c.rw.RLock()
	defer c.rw.RUnlock()
	return c.destroyed
}

// City - returns the city that the player resides in
// returns an empty string if there is none
func (c *Computer) City() string {
	c.rw.RLock()
	defer c.rw.RUnlock()
	return c.currentCity
}

// String - stringer interface impl
func (c *Computer) String() string {
	c.rw.RLock()
	defer c.rw.RUnlock()
	return c.Name()
}
