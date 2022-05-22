package game

import (
	"context"
	"testing"

	"github.com/buni/aliens-game/internal/pkg/board/mock"
	"github.com/buni/aliens-game/internal/pkg/player"
)

func TestGame_Next(t *testing.T) {
	tests := []struct {
		name string
		g    *Game
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.Next()
		})
	}
}

func TestGame_Start(t *testing.T) {
	tests := []struct {
		name  string
		g     *Game
		setup func(state *mock.MockState)
		ctx   context.Context
	}{
		{
			name: "",
			g: &Game{
				state:         nil,
				players:       []player.Player{},
				iterations:    0,
				maxIterations: 1,
			},
			ctx: context.Background(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.Start(tt.ctx)
		})
	}
}
