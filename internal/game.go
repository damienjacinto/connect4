package game

import (
	"image/color"

	"github.com/damienjacinto/connect4/internal/player"
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
	currentError  string
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

func (g *Game) humanMove(input Direction) (bool, error) {
	human := g.board.currentPlayer.(*player.HumanPlayer)
	switch input {
	case LEFT:
		g.board.MoveLeft()
	case RIGHT:
		g.board.MoveRight()
	case DOWN:
		human.SetMove(g.board.header.positionPiece)
		col := human.Play()
		if err := g.board.Play(col); err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

func (g *Game) computerMove() (bool, error) {
	ai := g.board.currentPlayer.(player.IAPlayer)
	col := ai.Play()
	if err := g.board.Play(col); err != nil {
		return false, err
	}
	return true, nil
}

func (g *Game) Update() error {
	if !g.inited {
		g.init()
	}

	moved := false
	var err error
	input := g.inputs.HandleInput()

	if g.result == UNKNOWNRESULT {
		if g.board.IsCurrentPlayerHuman() {
			moved, err = g.humanMove(input)
			if err != nil {
				g.currentError = err.Error()
			}
		} else if g.board.IsCurrentPlayerAI() {
			moved, err = g.computerMove()
			if err != nil {
				g.currentError = err.Error()
			}
		}

		if moved {
			g.currentError = ""
			g.updateResult()
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
	resultValue := g.result.FormatResult()
	wBoard, hBoard := g.board.Size()
	g.board.Draw(screen, (float32(g.width)-wBoard)/2, (float32(g.height)-hBoard)/2, resultValue, g.currentError)
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

func (g *Game) updateResult() {
	g.result = g.board.GetResult()
}

func (g *Game) Reset() {
	g.board.NewGame()
	g.result = UNKNOWNRESULT
}
