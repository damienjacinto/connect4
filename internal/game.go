package game

import (
	"image/color"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	widthBoard, heightBoard = 7, 6
)

// Game implements ebiten.Game interface.
type Game struct {
	width, height int
	title         string
	inited        bool
	board         *Board
}

func NewGame(width int, height int, title string) *Game {
	return &Game{
		width:  width,
		height: height,
		title:  title,
		inited: false,
	}
}

func (g *Game) Update() error {
	if !g.inited {
		g.init()
	}

	input := HandleInput()
	switch input {
	case LEFT:
		slog.Info("left")
	case RIGHT:
		slog.Info("right")
	case DOWN:
		slog.Info("down")
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	g.board.Print()

	vector.DrawFilledRect(screen, 10.0, 10.0, 50, 50, color.White, true)
	vector.DrawFilledCircle(screen, 35, 35, 20, color.Black, true)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.width, g.height
}

func (g *Game) init() {
	defer func() {
		g.inited = true
	}()

	ebiten.SetWindowSize(g.width, g.height)
	ebiten.SetWindowTitle(g.title)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	g.board = NewBoard(widthBoard, heightBoard)
}

func (g *Game) Start() {
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
