package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	sizeElement = 60
	padding     = 2
)

var (
	colorBoard = color.RGBA{R: 0, G: 128, B: 255, A: 255}
)

type Game struct {
	width, height int
	title         string
	inited        bool
	board         *GameBoard
	inputs        *Inputs
	result        Result
}

func NewGame(width int, height int, title string) *Game {
	return &Game{
		width:  width,
		height: height,
		title:  title,
		inited: false,
		result: UNKNOWNRESULT,
	}
}

func (g *Game) Update() error {
	if !g.inited {
		g.init()
	}

	input := g.inputs.HandleInput()

	if g.result == UNKNOWNRESULT {
		switch input {
		case LEFT:
			g.board.MoveLeft()
		case RIGHT:
			g.board.MoveRight()
		case DOWN:
			g.board.DropPiece()
			g.searchGameOver()
			if g.result == UNKNOWNRESULT {
				g.board.ChangePlayer()
			}
		}
	} else {
		switch input {
		case DOWN:
			g.Reset()
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.result != UNKNOWNRESULT {
		g.result.Draw(screen)
		return
	}
	wBoard, hBoard := g.board.Size()
	g.board.Draw(screen, (float32(g.width)-wBoard)/2, (float32(g.height)-hBoard)/2)
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

	g.board = NewGameBoard(colorBoard, sizeElement, padding)
	g.inputs = NewInput()
}

func (g *Game) Start() {
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

func (g *Game) searchGameOver() {
	g.result = g.board.GetResult()
}

func (g *Game) Reset() {
	g.board.NewGame()
	g.result = UNKNOWNRESULT
}
