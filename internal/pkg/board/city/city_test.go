package city

import (
	"sync"
	"testing"

	"github.com/buni/aliens-game/internal/pkg/player"
	"github.com/buni/aliens-game/internal/pkg/testutils"
	"github.com/stretchr/testify/assert"
)

func TestCity_RemoveLink(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		c        *City
		wantCity *City
		cityName string
	}{
		{
			name: "successfully remove links",
			c: &City{
				name: "test",
				links: map[Direction]*City{
					East:  {name: "remove-1", rw: &sync.RWMutex{}},
					North: {name: "city2", rw: &sync.RWMutex{}},
					West:  {name: "city3", rw: &sync.RWMutex{}},
					South: {name: "city4", rw: &sync.RWMutex{}},
				},
				rw: &sync.RWMutex{},
			},
			wantCity: &City{
				name: "test",
				links: map[Direction]*City{
					North: {name: "city2", rw: &sync.RWMutex{}},
					West:  {name: "city3", rw: &sync.RWMutex{}},
					South: {name: "city4", rw: &sync.RWMutex{}},
				},
				rw: &sync.RWMutex{},
			},
			cityName: "remove-1",
		},
		{
			name: "successfully remove multiple links",
			c: &City{
				name: "test",
				links: map[Direction]*City{
					East:  {name: "remove-1", rw: &sync.RWMutex{}},
					North: {name: "remove-1", rw: &sync.RWMutex{}},
					West:  {name: "city3", rw: &sync.RWMutex{}},
					South: {name: "city4", rw: &sync.RWMutex{}},
				},
				rw: &sync.RWMutex{},
			},
			wantCity: &City{
				name: "test",
				links: map[Direction]*City{
					West:  {name: "city3", rw: &sync.RWMutex{}},
					South: {name: "city4", rw: &sync.RWMutex{}},
				},
				rw: &sync.RWMutex{},
			},
			cityName: "remove-1",
		},
		{
			name: "non existent city noop",
			c: &City{
				name: "test",
				links: map[Direction]*City{
					North: {name: "city2", rw: &sync.RWMutex{}},
					West:  {name: "city4", rw: &sync.RWMutex{}},
					South: {name: "city3", rw: &sync.RWMutex{}},
				},
				rw: &sync.RWMutex{},
			},
			wantCity: &City{
				name: "test",
				links: map[Direction]*City{
					North: {name: "city2", rw: &sync.RWMutex{}},
					West:  {name: "city4", rw: &sync.RWMutex{}},
					South: {name: "city3", rw: &sync.RWMutex{}},
				},
				rw: &sync.RWMutex{},
			},
			cityName: "remove-1",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.c.RemoveLink(tt.cityName)
			assert.EqualValues(t, tt.c.links, tt.wantCity.links)
		})
	}
}

func TestCity_RemoveCityRefs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		c         *City
		wantLinks map[Direction]*City
	}{
		{
			name: "remove multiple refs",
			c: &City{
				name: "test",
				links: map[Direction]*City{
					North: {
						name: "city2",
						links: map[Direction]*City{
							North: {name: "test", rw: &sync.RWMutex{}},
						},
						rw: &sync.RWMutex{},
					},
					West: {
						name: "city3", links: map[Direction]*City{
							North: {name: "test", rw: &sync.RWMutex{}},
						},
						rw: &sync.RWMutex{},
					},
				},
				rw: &sync.RWMutex{},
			},
			wantLinks: map[Direction]*City{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tt.c.RemoveCityRefs()
			assert.EqualValues(t, tt.c.links, tt.wantLinks)
		})
	}
}

func TestCity_AddVisitor(t *testing.T) {
	t.Parallel()
	tests := []struct {
		visitor        player.Player
		name           string
		c              *City
		wantPlayersLen int
	}{
		{
			name:           "successfully add a visitor and keep existing one",
			wantPlayersLen: 2,
			visitor:        testutils.NewNoopPlayer("test-player-2"),
			c: &City{
				name:      "",
				links:     map[Direction]*City{},
				destroyed: false,
				visitors: []player.Player{
					testutils.NewNoopPlayer("test-player-1"),
				},
				rw: &sync.RWMutex{},
			},
		},
		{
			name:           "successfully add first visitor",
			wantPlayersLen: 1,
			visitor:        testutils.NewNoopPlayer("test-player-2"),
			c: &City{
				name:      "",
				links:     map[Direction]*City{},
				destroyed: false,
				visitors:  nil,
				rw:        &sync.RWMutex{},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.c.AddVisitor(tt.visitor)
			assert.Len(t, tt.c.visitors, tt.wantPlayersLen)
		})
	}
}

func TestCity_ShouldDestroy(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		c    *City
		want bool
	}{
		{
			name: "successfully destroy city",
			c: &City{
				name:      "",
				links:     map[Direction]*City{},
				destroyed: false,
				visitors:  []player.Player{testutils.NewNoopPlayer("test-player1"), testutils.NewNoopPlayer("test-player2")},
				rw:        &sync.RWMutex{},
			},
			want: true,
		},
		{
			name: "single player destroy check",
			c: &City{
				name:      "",
				links:     map[Direction]*City{},
				destroyed: false,
				visitors:  []player.Player{testutils.NewNoopPlayer("test-player1")},
				rw:        &sync.RWMutex{},
			},
			want: false,
		},
		{
			name: "no players destroy check",
			c: &City{
				name:      "",
				links:     map[Direction]*City{},
				destroyed: false,
				rw:        &sync.RWMutex{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.c.ShouldDestroy(); got != tt.want {
				t.Errorf("City.ShouldDestroy() = %v, want %v", got, tt.want)
			}
		})
	}
}

 func TestCity_GetName(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		c    *City
		want string
	}{
		{
			name: "test get name",
			c: &City{
				name:      "test1",
				links:     map[Direction]*City{},
				destroyed: false,
				visitors:  []player.Player{},
				rw:        &sync.RWMutex{},
			},
			want: "test1",
		},
		{
			name: "test get name",
			c: &City{
				name:      "test2",
				links:     map[Direction]*City{},
				destroyed: false,
				visitors:  []player.Player{},
				rw:        &sync.RWMutex{},
			},
			want: "test2",
		},
		{
			name: "test get name",
			c: &City{
				name:      "test3",
				links:     map[Direction]*City{},
				destroyed: false,
				visitors:  []player.Player{},
				rw:        &sync.RWMutex{},
			},
			want: "test3",
		},
		{
			name: "test get name",
			c: &City{
				name:      "test4",
				links:     map[Direction]*City{},
				destroyed: false,
				visitors:  []player.Player{},
				rw:        &sync.RWMutex{},
			},
			want: "test4",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.c.GetName(); got != tt.want {
				t.Errorf("City.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCity_GetRandomCityLink(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		c        *City
		wantCity string
		wantOk   bool
	}{
		{
			name: "successfully get random link",
			c: &City{
				name: "test-city",
				links: map[Direction]*City{
					South: {
						name:      "test-city2",
						links:     map[Direction]*City{},
						destroyed: false,
						visitors:  []player.Player{},
						rw:        &sync.RWMutex{},
					},
					North: {
						name:      "test-city3",
						links:     map[Direction]*City{},
						destroyed: false,
						visitors:  []player.Player{},
						rw:        &sync.RWMutex{},
					},
					West: {
						name:      "test-city4",
						links:     map[Direction]*City{},
						destroyed: false,
						visitors:  []player.Player{},
						rw:        &sync.RWMutex{},
					},
					East: {
						name:      "test-city5",
						links:     map[Direction]*City{},
						destroyed: false,
						visitors:  []player.Player{},
						rw:        &sync.RWMutex{},
					},
				},
				destroyed: false,
				visitors:  []player.Player{},
				rw:        &sync.RWMutex{},
			},
			wantCity: "",
			wantOk:   true,
		},
		{
			name: "fail to get random link",
			c: &City{
				name:      "test-city",
				links:     map[Direction]*City{},
				destroyed: false,
				visitors:  []player.Player{},
				rw:        &sync.RWMutex{},
			},
			wantCity: "",
			wantOk:   false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotCity, gotOk := tt.c.GetRandomCityLink()
			if gotOk {
				assert.NotEmpty(t, gotCity)
			} else {
				assert.Empty(t, gotCity)
			}

			if gotOk != tt.wantOk {
				t.Errorf("City.GetRandomCityLink() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestCity_RemoveVisitor(t *testing.T) {
	t.Parallel()
	type args struct{}
	tests := []struct {
		name           string
		c              *City
		visitor        player.Player
		wantPlayersLen int
	}{
		{
			name: "successfully remove one out of two visitors",
			c: &City{
				name:      "",
				links:     map[Direction]*City{},
				destroyed: false,
				visitors:  []player.Player{testutils.NewNoopPlayer("test1"), testutils.NewNoopPlayer("test2")},
				rw:        &sync.RWMutex{},
			},
			wantPlayersLen: 1,
			visitor:        testutils.NewNoopPlayer("test1"),
		},
		{
			name: "successfully remove one out of one visitors",
			c: &City{
				name:      "",
				links:     map[Direction]*City{},
				destroyed: false,
				visitors:  []player.Player{testutils.NewNoopPlayer("test1"), testutils.NewNoopPlayer("test2")},
				rw:        &sync.RWMutex{},
			},
			wantPlayersLen: 1,
			visitor:        testutils.NewNoopPlayer("test1"),
		},
		{
			name: " remove visitor no-op",
			c: &City{
				name:      "",
				links:     map[Direction]*City{},
				destroyed: false,
				rw:        &sync.RWMutex{},
			},
			wantPlayersLen: 0,
			visitor:        testutils.NewNoopPlayer("test1"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.c.RemoveVisitor(tt.visitor)
			assert.Len(t, tt.c.visitors, tt.wantPlayersLen)
		})
	}
}
