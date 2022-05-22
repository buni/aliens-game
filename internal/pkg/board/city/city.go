package city

import (
	"log"
	"sync"

	"github.com/buni/aliens-game/internal/pkg/player"
)

type City struct {
	name      string
	links     map[Direction]*City
	destroyed bool
	visitors  player.Players
	rw        *sync.RWMutex
}

func NewCity(name string) *City {
	return &City{name: name, visitors: make(player.Players, 0, 1), links: map[Direction]*City{}, rw: &sync.RWMutex{}}
}

// AddLink - behaves like a put operation
func (c *City) AddLink(dir Direction, city *City) {
	c.rw.Lock()
	defer c.rw.Unlock()
	c.links[dir] = city
}

// RemoveLink - remove link(s) to a given city by name
func (c *City) RemoveLink(cityName string) {
	c.rw.Lock()
	defer c.rw.Unlock()
	for k, v := range c.links {
		if v.GetName() == cityName {
			delete(c.links, k)
			// break
		}
	}
}

// RemoveCityRefs - remove all links/references to and from the current city
func (c *City) RemoveCityRefs() {
	c.rw.RLock()
	for _, link := range c.links {
		// linkName := link.GetName()
		link.RemoveLink(c.name) // remove each link for the given city
	}
	c.rw.RUnlock()

	c.rw.Lock()
	c.links = map[Direction]*City{}
	c.rw.Unlock()
}

// AddVisitor - appends a visitor to the city visitors list
func (c *City) AddVisitor(visitor player.Player) {
	c.rw.Lock()
	defer c.rw.Unlock()

	c.visitors = append(c.visitors, visitor)
}

// ShouldDestroy - checks where the city can be destroyed and returns true or false
// if true it marks the city as destroyed and kills all the current visitors
func (c *City) ShouldDestroy() bool {
	c.rw.Lock()
	defer c.rw.Unlock()
	if len(c.visitors) >= 2 {
		log.Printf("Aliens fighting in %s %v \n", c.name, c.visitors.GetNames())
		c.destroyed = true
		c.visitors.Destroy()
		return true
	}
	return false
}

// GetName - returns the city name
func (c *City) GetName() string {
	c.rw.RLock()
	defer c.rw.RUnlock()
	return c.name
}

// GetRandomCityLink - returns a random linked city
// returns false if there are no links
func (c *City) GetRandomCityLink() (city string, ok bool) {
	c.rw.RLock()
	defer c.rw.RUnlock()

	for _, link := range c.links {
		return link.name, true
	}

	return
}

// RemoveVisitor - remove a visitor from the city
// if not present in the list its a no-op
func (c *City) RemoveVisitor(movingVisitor player.Player) {
	c.rw.Lock()
	defer c.rw.Unlock()

	for k, visitor := range c.visitors {
		if visitor.Name() == movingVisitor.Name() { // assumes that the generated names can't be the same for two players (this is not validated)
			c.visitors[k] = c.visitors[len(c.visitors)-1]
			c.visitors[len(c.visitors)-1] = nil
			c.visitors = c.visitors[:len(c.visitors)-1]
			break
		}
	}
}

// GetCityLinks - returns a copy of all city links
func (c *City) GetCityLinks() map[Direction]*City {
	result := map[Direction]*City{}
	for dir, link := range c.links { // copy the map to prevent possible race conditions
		result[dir] = link
	}
	return result
}

const (
	North Direction = iota + 1 // make sure iota starts from 1 to ensure 0 is an invalid direction
	East
	South
	West
)
