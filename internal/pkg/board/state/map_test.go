package state

import (
	"io"
	"os"
	"sync"
	"testing"

	"github.com/buni/aliens-game/internal/pkg/board/city"
	"github.com/buni/aliens-game/internal/pkg/player"
	"github.com/buni/aliens-game/internal/pkg/testutils"
	"github.com/stretchr/testify/assert"
)

func Test_boardState_GetCities(t *testing.T) {
	tests := []struct {
		name       string
		m          *boardState
		wantCities []*city.City
	}{
		{
			name: "get cities",
			m: &boardState{
				Cities: map[string]*city.City{
					"city1": city.NewCity("city1"),
					"city2": city.NewCity("city2"),
				},
				rw: &sync.RWMutex{},
			},
			wantCities: []*city.City{city.NewCity("city1"), city.NewCity("city2")},
		},
		{
			name: "get cities empty",
			m: &boardState{
				Cities: map[string]*city.City{},
				rw:     &sync.RWMutex{},
			},
			wantCities: []*city.City{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCities := tt.m.GetCities()
			assert.Len(t, gotCities, len(tt.wantCities))
		})
	}
}

func Test_boardState_GetNextDirection(t *testing.T) {
	tests := []struct {
		name         string
		m            *boardState
		currentCity  string
		wantCityName string
		wantOk       bool
	}{
		{
			name: "get initial city",
			m: &boardState{
				Cities: map[string]*city.City{"test": city.NewCity("test")},
				rw:     &sync.RWMutex{},
			},
			currentCity:  "",
			wantCityName: "",
			wantOk:       true,
		},
		{
			name: "get next city",
			m: &boardState{
				Cities: map[string]*city.City{"test": city.NewCity("test"), "test2": city.NewCity("test2")},
				rw:     &sync.RWMutex{},
			},
			currentCity:  "",
			wantCityName: "",
			wantOk:       true,
		},
		{
			name: "no next city player - stuck",
			m: &boardState{
				Cities: map[string]*city.City{"test": city.NewCity("test")},
				rw:     &sync.RWMutex{},
			},
			currentCity:  "test",
			wantCityName: "",
			wantOk:       false,
		},
		{
			name: "current city doesn't exist",
			m: &boardState{
				Cities: map[string]*city.City{"test": city.NewCity("test")},
				rw:     &sync.RWMutex{},
			},
			currentCity:  "test2",
			wantCityName: "",
			wantOk:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCityName, gotOk := tt.m.GetNextDirection(tt.currentCity)
			if tt.wantOk {
				assert.NotEmpty(t, gotCityName)
			}

			if gotOk != tt.wantOk {
				t.Errorf("boardState.GetNextDirection() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_boardState_DeleteCityAndLinks(t *testing.T) {
	type args struct{}
	tests := []struct {
		name       string
		m          *boardState
		wantCities map[string]*city.City
		cityName   string
	}{
		{
			name: "delete a city",
			m: &boardState{
				Cities: map[string]*city.City{"test": city.NewCity("test")},
				rw:     &sync.RWMutex{},
			},
			cityName:   "test",
			wantCities: map[string]*city.City{},
		},
		{
			name: "delete a city one left",
			m: &boardState{
				Cities: map[string]*city.City{"test": city.NewCity("test"), "test2": city.NewCity("test2")},
				rw:     &sync.RWMutex{},
			},
			cityName:   "test2",
			wantCities: map[string]*city.City{"test": city.NewCity("test")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.DeleteCityAndLinks(tt.cityName)
			assert.Len(t, tt.m.Cities, len(tt.wantCities))
		})
	}
}

func Test_boardState_DeleteCity(t *testing.T) {
	tests := []struct {
		name       string
		m          *boardState
		wantCities map[string]*city.City
		cityName   string
	}{
		{
			name: "delete a city",
			m: &boardState{
				Cities: map[string]*city.City{"test": city.NewCity("test")},
				rw:     &sync.RWMutex{},
			},
			cityName:   "test",
			wantCities: map[string]*city.City{},
		},
		{
			name: "delete a city one left",
			m: &boardState{
				Cities: map[string]*city.City{"test": city.NewCity("test"), "test2": city.NewCity("test2")},
				rw:     &sync.RWMutex{},
			},
			cityName:   "test2",
			wantCities: map[string]*city.City{"test": city.NewCity("test")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.DeleteCity(tt.cityName)
		})
	}
}

func Test_boardState_MoveVisitor(t *testing.T) {
	tests := []struct {
		name         string
		m            *boardState
		nextCityName string
		visitor      player.Player
		wantOk       bool
	}{
		{
			name: "successfully move player to next city",
			m: &boardState{
				Cities: map[string]*city.City{"test": city.NewCity("test"), "test2": city.NewCity("test2")},
				rw:     &sync.RWMutex{},
			},
			nextCityName: "test2",
			visitor:      testutils.NewNoopPlayer("test-player"),
			wantOk:       true,
		},
		{
			name: "successfully move player form one to another city",
			m: &boardState{
				Cities: map[string]*city.City{"test": city.NewCity("test"), "test2": city.NewCity("test2")},
				rw:     &sync.RWMutex{},
			},
			nextCityName: "test2",
			visitor:      testutils.NewNoopPlayerWithCity("test-player", "test"),
			wantOk:       true,
		},
		{
			name: "fail move player to non existant city",
			m: &boardState{
				Cities: map[string]*city.City{"test": city.NewCity("test"), "test2": city.NewCity("test2")},
				rw:     &sync.RWMutex{},
			},
			nextCityName: "",
			visitor:      testutils.NewNoopPlayer("test-player"),
			wantOk:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOk := tt.m.MoveVisitor(tt.nextCityName, tt.visitor); gotOk != tt.wantOk {
				t.Errorf("boardState.MoveVisitor() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_boardState_String(t *testing.T) {
	tests := []struct {
		name  string
		m     *boardState
		want  string
		setup func(state *boardState) string
	}{
		{
			name: "",
			m:    &boardState{},

			setup: func(state *boardState) string {
				testCities, err := os.Open("testdata/board-3-cities.txt")
				assert.NoError(t, err)
				testState, err := NewBoardState(testCities)
				assert.NoError(t, err)
				t.Log(testState)
				state.Cities = testState.(*boardState).Cities

				testCities, err = os.Open("testdata/board-3-cities.txt")
				body, err := io.ReadAll(testCities)
				assert.NoError(t, err)
				return string(body)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.want = tt.setup(tt.m)
			got := tt.m.String()
			assert.Len(t, tt.want, len(got))
		})
	}
}
