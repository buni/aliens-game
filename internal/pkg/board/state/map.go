package state

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"log"
	"strings"
	"sync"

	"github.com/buni/aliens-game/internal/pkg/board"
	"github.com/buni/aliens-game/internal/pkg/board/city"

	"github.com/buni/aliens-game/internal/pkg/player"
)

// State - game state
type boardState struct {
	Cities map[string]*city.City
	rw     *sync.RWMutex
}

func NewBoardState(state io.Reader) (board.State, error) {
	mapState := &boardState{
		Cities: map[string]*city.City{},
		rw:     &sync.RWMutex{},
	}

	s := bufio.NewScanner(state)
	for s.Scan() {
		lineFields := strings.Fields(s.Text()) // assumes that all cities don't have space in their name
		if len(lineFields) < 1 {
			return nil, errors.New("no city in line")
		}

		cityName := lineFields[0]

		stateCity, ok := mapState.Cities[cityName]
		if !ok {
			stateCity = city.NewCity(cityName)
			mapState.Cities[cityName] = stateCity
		}

		for _, field := range lineFields[1:] {
			kv := strings.Split(field, "=")
			if len(kv) != 2 {
				return nil, errors.New("bad link")
			}
			key := kv[0]
			value := kv[1]
			dir, err := city.ParseDirection(key)
			if err != nil {
				return nil, err
			}

			linkCity, ok := mapState.Cities[value]
			if !ok {
				linkCity = city.NewCity(value)
				mapState.Cities[value] = linkCity
			}

			stateCity.AddLink(dir, linkCity) // TODO: make a func ?h
		}
	}
	return mapState, nil
}

func (m *boardState) GetCities() (cities []*city.City) {
	m.rw.RLock()
	defer m.rw.RUnlock()
	for _, city := range m.Cities {
		cities = append(cities, city)
	}
	return
}

func (m *boardState) GetNextDirection(currentCity string) (cityName string, ok bool) {
	m.rw.RLock()
	defer m.rw.RUnlock()

	if currentCity == "" {
		for _, city := range m.Cities {
			// maps don't have a specific iteration order, we can utilize that to get the first city (the same principal is used for GetRandomCityLink)
			// the key random distribution is far from evenly distributed, as there is bias towards the first few keys
			// but its fine for this case (as alternative the keys can be extracted or kept in a slice an picked with an alternative algorithm form crypto/rand or math/rand)
			return city.GetName(), true
		}
	}

	city, ok := m.GetCity(currentCity)
	if !ok {
		return
	}

	return city.GetRandomCityLink()
}

func (m *boardState) DeleteCityAndLinks(cityName string) {
	city, ok := m.GetCity(cityName)
	if !ok {
		return
	}

	m.DeleteCity(cityName)
	city.RemoveCityRefs()
}

func (m *boardState) DeleteCity(cityName string) {
	m.rw.Lock()
	defer m.rw.Unlock()
	delete(m.Cities, cityName)
}

func (m *boardState) GetCity(cityName string) (stateCity *city.City, ok bool) {
	m.rw.RLock()
	defer m.rw.RUnlock()
	stateCity, ok = m.Cities[cityName]
	return
}

func (m *boardState) MoveVisitor(nextCityName string, visitor player.Player) (ok bool) {
	nextCity, ok := m.GetCity(nextCityName)
	if !ok {
		log.Println("no next city", m.String())
		return false
	}

	currentCity, ok := m.GetCity(visitor.City())
	if ok {
		currentCity.RemoveVisitor(visitor)
	}

	nextCity.AddVisitor(visitor)

	// if ok = nextCity.ShouldDestroy(); ok { // TODO: move from here probably not the right place for game logic
	// 	// destroy is called after a move, not after the round ends, so a player could potentially 'die' before making a move in the current round (it is also racy in nature)
	// 	m.DeleteCityAndLinks(nextCityName)
	// 	return true
	// }

	return true
}

// String - stringer impl
// because maps are used the string representation of the printed cities
// won't match the insertion/parse order
func (m *boardState) String() string {
	buff := &bytes.Buffer{}
	for _, v := range m.Cities {

		line := []string{v.GetName()}
		for dir, city := range v.GetCityLinks() {
			line = append(line, strings.Join([]string{dir.String(), city.GetName()}, "="))
		}
		buff.WriteString(strings.Join(line, " "))
		buff.WriteString("\n")
	}
	return buff.String()
}

// Foo north=Bar west=Baz south=Qu-ux
// Bar south=Foo west=Bee
