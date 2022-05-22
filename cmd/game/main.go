package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/buni/aliens-game/internal/pkg/game"
)

func main() {
	boardFile := flag.String("board-file", "board.txt", "path to board file")
	computerPlayersCount := flag.Int("computer-players-count", 1, "set the computer players count")
	gameTickRate := flag.Duration("tick-duration", 1, "game loop tick duration")
	maxIterations := flag.Int("max-iterations", 10000, "max game loop iterations")

	sig := make(chan os.Signal, 1)
	// signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-sig
		cancel()
	}()

	g, err := game.NewGame(*boardFile, game.WithComputerPlayers(*computerPlayersCount), game.WithGameTick(*gameTickRate), game.WithMaxIterations(*maxIterations), game.WithHumanPlayers(1))
	if err != nil {
		log.Fatalf("Failed to create new game: %v \n", err)
	}

	g.Start(ctx)
}
