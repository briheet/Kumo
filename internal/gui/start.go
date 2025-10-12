package gui

import (
	"context"
	"log"
	"runtime"

	"github.com/briheet/kumo/internal/tooling"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"go.uber.org/zap"
)

func init() {

	ebiten.SetWindowSize(900, 720)
	ebiten.SetWindowTitle("Hello from kumo")
}

type Game struct {
	logger *zap.Logger
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello from kumo")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func Start(ctx context.Context) error {

	// NewLogger
	logger := tooling.NewDevelopmentLogger() // TODO: Change to prod via flags for release builds
	defer logger.Sync()

	game := Game{
		logger: logger,
	}

	game.logger.Info(
		"starting gui",
		zap.Int("num_cores", runtime.GOMAXPROCS(0)),
	)

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatalf("error start the gui: %e", err)
	}

	game.logger.Info(
		"exiting, take care!",
	)

	return nil
}
