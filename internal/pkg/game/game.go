package game

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/buni/aliens-game/internal/pkg/board"
	"github.com/buni/aliens-game/internal/pkg/board/state"
	"github.com/buni/aliens-game/internal/pkg/player"
	"github.com/buni/aliens-game/internal/pkg/player/computer"
	"github.com/buni/aliens-game/internal/pkg/player/human"

	"github.com/tjarratt/babble"
)

// Game - game
type Game struct {
	state         board.State
	players       player.Players
	iterations    int
	maxIterations int
	ticker        *time.Ticker
}

// Option - game options
type Option func(*Game) error

// WithPlayers set the number of computer players
func WithComputerPlayers(playerCount int) Option {
	return func(g *Game) error {
		for i := 0; i < playerCount; i++ {
			p := computer.NewPlayer(g.state, babble.NewBabbler().Babble())
			g.players = append(g.players, p)
		}
		return nil
	}
}

// WithPlayers set the number of computer players
func WithHumanPlayers(playerCount int) Option {
	return func(g *Game) error {
		for i := 0; i < playerCount; i++ {
			p := human.NewPlayer(g.state, babble.NewBabbler().Babble(), os.Stdin)
			g.players = append(g.players, p)
		}

		return nil
	}
}

// WithGameTick - set a custom game tick rate
func WithGameTick(tickDuration time.Duration) Option {
	return func(g *Game) error {
		g.ticker = time.NewTicker(tickDuration)
		return nil
	}
}

// WithMaxIterations - set a custom max iterations, after which the game ends
func WithMaxIterations(maxIterations int) Option {
	return func(g *Game) error {
		g.maxIterations = maxIterations
		return nil
	}
}

// NewGame -
func NewGame(boardFile string, opts ...Option) (*Game, error) {
	game := &Game{}
	f, err := os.Open(boardFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open board state file %w", err)
	}

	board, err := state.NewBoardState(f)
	if err != nil {
		return nil, fmt.Errorf("failed parse board state file %w", err)
	}

	game.state = board

	for _, opt := range opts {
		err := opt(game)
		if err != nil {
			return nil, err
		}
	}
	return game, nil
}

func (g *Game) Next() {
	wg := &sync.WaitGroup{}
	for _, player := range g.players {
		if !player.IsDestroyed() {
			player := player
			wg.Add(1)
			go func() { // this is not necessarily faster than in most if not all cases due to lock contention and other overheads
				// it was done mostly as an exercise
				player.Move("")
				wg.Done()
			}()
		}
	}
	wg.Wait()

	for _, v := range g.state.GetCities() {
		if v.ShouldDestroy() {
			name := v.GetName()
			g.state.DeleteCityAndLinks(name)
		}
	}
	g.iterations++
}

func (g *Game) Start(ctx context.Context) {
	fmt.Println("Starting game loop")
loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		// case <-g.ticker.C:
		default:
			if g.players.Destroyed() {
				fmt.Printf("Game Over all aliens are destroyed \n")
				break loop
			}

			if g.iterations >= g.maxIterations { // if total iterations are 10k and not all aliens are dead (aliens that are stuck are not counted as "dead"), take that as 10k moves for all aliens that are alive
				fmt.Printf("Game Over aliens made >= %v moves \n", g.iterations)
				for _, v := range g.players {
					fmt.Println(v.Name(), v.IsDestroyed())
				}
				break loop
			}

			g.Next()
		}
	}

	fmt.Printf("Board state:\n%s", g.state)
}
