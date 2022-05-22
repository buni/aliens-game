package computer

import (
	"sync"
	"testing"

	"github.com/buni/aliens-game/internal/pkg/board/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestComputer_Move(t *testing.T) {
	tests := []struct {
		name  string
		setup func(state *mock.MockState, c *Computer)
		c     *Computer
		want  bool
	}{
		{
			name: "successfully move computer player",
			setup: func(state *mock.MockState, c *Computer) {
				state.EXPECT().GetNextDirection(gomock.Any()).Return("next_city", true)
				state.EXPECT().MoveVisitor("next_city", c).Return(true)
			},
			c: &Computer{
				moves:       0,
				name:        "",
				currentCity: "",
				board:       nil,
				destroyed:   false,
				rw:          &sync.RWMutex{},
			},
			want: true,
		},
		{
			name: "move noop when player is destroyed",
			setup: func(state *mock.MockState, c *Computer) {
			},
			c: &Computer{
				moves:       0,
				name:        "",
				currentCity: "",
				board:       nil,
				destroyed:   true,
				rw:          &sync.RWMutex{},
			},
			want: false,
		},
		{
			name: "no next direction",
			setup: func(state *mock.MockState, c *Computer) {
				state.EXPECT().GetNextDirection(gomock.Any()).Return("", false)
			},
			c: &Computer{
				moves:       0,
				name:        "",
				currentCity: "",
				board:       nil,
				destroyed:   false,
				rw:          &sync.RWMutex{},
			},
			want: false,
		},
		{
			name: "fail to move visitor",
			setup: func(state *mock.MockState, c *Computer) {
				state.EXPECT().GetNextDirection(gomock.Any()).Return("next_city", true)
				state.EXPECT().MoveVisitor("next_city", c).Return(false)
			},
			c: &Computer{
				moves:       0,
				name:        "",
				currentCity: "",
				board:       nil,
				destroyed:   false,
				rw:          &sync.RWMutex{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			stateMock := mock.NewMockState(ctrl)
			tt.setup(stateMock, tt.c)
			tt.c.board = stateMock

			if got := tt.c.Move(""); got != tt.want {
				t.Errorf("Computer.Move() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputer_SetCity(t *testing.T) {
	tests := []struct {
		name     string
		c        *Computer
		city     string
		wantCity string
	}{
		{
			name:     "successfully set city",
			c:        &Computer{moves: 0, name: "", currentCity: "", board: nil, destroyed: false, rw: &sync.RWMutex{}},
			city:     "test",
			wantCity: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.SetCity(tt.city)
			assert.Equal(t, tt.wantCity, tt.city)
		})
	}
}

func TestComputer_Name(t *testing.T) {
	tests := []struct {
		name string
		c    *Computer
		want string
	}{
		{
			name: "successfully get player name",
			c: &Computer{
				moves:       0,
				name:        "test-name",
				currentCity: "",
				board:       nil,
				destroyed:   false,
				rw:          &sync.RWMutex{},
			},
			want: "test-name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Name(); got != tt.want {
				t.Errorf("Computer.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputer_Destroy(t *testing.T) {
	tests := []struct {
		name string
		c    *Computer
		want bool
	}{
		{
			name: "successfully destroy player",
			c: &Computer{
				moves:       0,
				name:        "",
				currentCity: "",
				board:       nil,
				destroyed:   false,
				rw:          &sync.RWMutex{},
			},
			want: true,
		},
		{
			name: "destroy noop player",
			c: &Computer{
				moves:       0,
				name:        "",
				currentCity: "",
				board:       nil,
				destroyed:   true,
				rw:          &sync.RWMutex{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Destroy()
			assert.Equal(t, tt.want, tt.c.destroyed)
		})
	}
}

func TestComputer_IsDestroyed(t *testing.T) {
	tests := []struct {
		name string
		c    *Computer
		want bool
	}{
		{
			name: "player is not destroyed",
			c: &Computer{
				moves:       0,
				name:        "",
				currentCity: "",
				board:       nil,
				destroyed:   false,
				rw:          &sync.RWMutex{},
			},
			want: false,
		},
		{
			name: "player is  destroyed",
			c: &Computer{
				moves:       0,
				name:        "",
				currentCity: "",
				board:       nil,
				destroyed:   true,
				rw:          &sync.RWMutex{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.IsDestroyed(); got != tt.want {
				t.Errorf("Computer.IsDestroyed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputer_City(t *testing.T) {
	tests := []struct {
		name string
		c    *Computer
		want string
	}{
		{
			name: "get city",
			c: &Computer{
				moves:       0,
				name:        "",
				currentCity: "test",
				board:       nil,
				destroyed:   false,
				rw:          &sync.RWMutex{},
			},
			want: "test",
		},
		{
			name: "get city empty",
			c: &Computer{
				moves:       0,
				name:        "",
				currentCity: "",
				board:       nil,
				destroyed:   false,
				rw:          &sync.RWMutex{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.City(); got != tt.want {
				t.Errorf("Computer.City() = %v, want %v", got, tt.want)
			}
		})
	}
}
