package human

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"

	"github.com/buni/aliens-game/internal/pkg/board"
	"github.com/buni/aliens-game/internal/pkg/board/city"
)

type Human struct {
	moves       int
	name        string
	currentCity string
	board       board.State
	destroyed   bool
	inPipe      io.Reader
	rw          *sync.RWMutex
}

// NewPlayer - new computer player instance
func NewPlayer(board board.State, name string, inPipe io.Reader) *Human {
	return &Human{board: board, name: name, inPipe: inPipe, rw: &sync.RWMutex{}} // TODO: replace board
}

// Move - the player on the map
func (c *Human) Move(nextCityName string) bool {
	if c.IsDestroyed() {
		return false
	}

	currentCityName := c.City()
	ok := false

	switch currentCityName {
	case "":
		log.Println(currentCityName, "dadas")
		nextCityName, ok = c.board.GetNextDirection(currentCityName)
		if !ok {
			return false
		}

		if c.board.MoveVisitor(nextCityName, c) {
			c.SetCity(nextCityName)
			return true
		}

	default:
		scanner := bufio.NewScanner(c.inPipe)

		cc, _ := c.board.GetCity(currentCityName)
		fmt.Printf("Chose your next direction %s \n", c.Name())
		for dir, city := range cc.GetCityLinks() {
			log.Printf("Direction:%s City:%s \n", dir, city.GetName())
		}

		for scanner.Scan() { // TODO: break logic a bit
			currentCity, _ := c.board.GetCity(c.City())
			links := currentCity.GetCityLinks()
			if len(links) == 0 {
				fmt.Println("player is stuck")
				break
			}
			for dir, city := range links {
				fmt.Printf("Direction:%s City:%s \n", dir, city.GetName())
			}
			if c.IsDestroyed() {
				return false
			}

			move := strings.TrimSpace(strings.ToLower(scanner.Text()))
			dir, err := city.ParseDirection(move)
			if err != nil {
				log.Printf("invalid direction chosen: %v \n", err)
				continue
			}

			nextCity, ok := links[dir]
			if !ok {
				log.Println("direction not found")
				continue
			}

			ok = c.board.MoveVisitor(nextCity.GetName(), c)
			if !ok {
				log.Println("failed to move to city, try again")
				continue
			}
			c.SetCity(nextCity.GetName())

			break
		}
	}

	return false
}

// SetCity - sets the current player's city name
func (c *Human) SetCity(city string) {
	c.rw.Lock()
	defer c.rw.Unlock()
	c.currentCity = city
}

// Name - returns the player name
func (c *Human) Name() string {
	c.rw.RLock()
	defer c.rw.RUnlock()
	return c.name
}

// Destroy - marks the player as destroyed
func (c *Human) Destroy() {
	c.rw.Lock()
	defer c.rw.Unlock()
	c.destroyed = true
}

// IsDestroyed - returns whether the player is destroyed or not
func (c *Human) IsDestroyed() bool {
	c.rw.RLock()
	defer c.rw.RUnlock()
	return c.destroyed
}

// City - returns the city that the player resides in
// returns an empty string if there is none
func (c *Human) City() string {
	c.rw.RLock()
	defer c.rw.RUnlock()
	return c.currentCity
}

// String - stringer interface impl
func (c *Human) String() string {
	c.rw.RLock()
	defer c.rw.RUnlock()
	return c.Name()
}
