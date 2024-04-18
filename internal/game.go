package game

import (
	"image/color"

	"github.com/damienjacinto/connect4/internal/gameboard"
	"github.com/damienjacinto/connect4/internal/header"
	"github.com/damienjacinto/connect4/internal/player"
	"github.com/damienjacinto/connect4/internal/result"
	"github.com/damienjacinto/connect4/internal/selectboard"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	sizeElement = 60
	padding     = 2
	width       = 7
	height      = 6
)

var (
	colorBoard = color.RGBA{R: 0, G: 128, B: 255, A: 255}
)

type Game struct {
	widthScreen, heightScreen int
	title                     string
	inited                    bool
	selected                  bool
	gameBoard                 *gameboard.GameBoard
	selectBoard               *selectboard.SelectBoard
	inputs                    *Inputs
	result                    result.Result
	player1                   player.IPlayer
	player2                   player.IPlayer
	currentPlayer             player.IPlayer
	header                    *header.Header
}

func NewGame(widthScreen int, heightScreen int, title string) *Game {
	header := header.NewHeader(sizeElement, padding, width)
	player1 := player.NewAIPlayer(player.Yellow, 1, player.ALPHABETA)
	player2 := player.NewAIPlayer(player.Red, 2, player.ALPHABETA)
	//player2 := player.NewHumanPlayer(red, 2, "Player 2")

	return &Game{
		widthScreen:   widthScreen,
		heightScreen:  heightScreen,
		title:         title,
		selected:      true,
		inited:        false,
		player1:       player1,
		player2:       player2,
		currentPlayer: player1,
		header:        header,
		result:        result.UNKNOWNRESULT,
	}
}

func (g *Game) humanMove(input Direction) (bool, error) {
	human := g.GetCurrentPlayer().(*player.HumanPlayer)
	switch input {
	case LEFT:
		g.header.MoveLeft()
	case RIGHT:
		g.header.MoveRight()
	case QUIT:
		return false, ebiten.Termination
	case DOWN:
		human.SetMove(g.header.GetPositionPiece())
		col := human.Play(g.gameBoard.GetBoard())
		if err := g.gameBoard.Play(col, human.GetValue()); err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

func (g *Game) computerMove() (bool, error) {
	col := g.GetCurrentPlayer().Play(g.gameBoard.GetBoard())
	if err := g.gameBoard.Play(col, g.currentPlayer.GetValue()); err != nil {
		return false, err
	}
	return true, nil
}

func (g *Game) updateSelectPlayer(input Direction) error {

	return nil
}

func (g *Game) updateSelectMove(input Direction) error {
	moved := false
	var err error
	if g.result == result.UNKNOWNRESULT {
		if g.IsCurrentPlayerHuman() {
			moved, err = g.humanMove(input)
			if err == ebiten.Termination {
				return err
			} else if err != nil {
				g.header.SetDisplay(err.Error())
			}
		} else if g.IsCurrentPlayerAI() {
			moved, err = g.computerMove()
			if err != nil {
				g.header.SetDisplay(err.Error())
			}
		}
		if moved {
			g.header.SetDisplay("")
			g.result = g.getResult()
			if g.result == result.UNKNOWNRESULT {
				g.ChangePlayer()
			}
		}
	} else {
		switch input {
		case DOWN:
			g.Reset()
		case QUIT:
			return ebiten.Termination
		}
	}
	return nil
}

func (g *Game) Update() error {
	if !g.inited {
		g.init()
	}

	input := g.inputs.HandleInput()
	if !g.selected {
		return g.updateSelectPlayer(input)
	}
	return g.updateSelectMove(input)
}

func (g *Game) Draw(screen *ebiten.Image) {
	if !g.selected {
		g.selectBoard.Draw(screen, g.widthScreen, g.heightScreen)
	} else {
		g.gameBoard.Draw(screen, g.widthScreen, g.heightScreen, g.player1.GetColor(), g.player2.GetColor())
		wBoard, _ := g.gameBoard.Size()
		initx := (float32(g.widthScreen) - wBoard) / 2
		middleScreen := initx + (float32(width)*sizeElement+float32((width-1))*padding)/2
		if g.result == result.UNKNOWNRESULT {
			g.header.Draw(screen, middleScreen, g.currentPlayer)
		} else {
			ebitenutil.DebugPrint(screen, g.result.FormatResult())
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.widthScreen, g.heightScreen
}

func (g *Game) init() {
	defer func() {
		g.inited = true
	}()

	ebiten.SetWindowSize(g.widthScreen, g.heightScreen)
	ebiten.SetWindowTitle(g.title)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	g.inputs = NewInput()
	g.gameBoard = gameboard.NewGameBoard(width, height, colorBoard, sizeElement, padding)
	g.selectBoard = selectboard.NewSelectBoard()
	g.Reset()
}

func (g *Game) Start() {
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

func (g *Game) getResult() result.Result {
	res := result.UNKNOWNRESULT
	if g.gameBoard.GetBoard().IsFinished() {
		res = result.Result(g.currentPlayer.GetValue())
	}
	if g.gameBoard.GetBoard().IsFull() {
		res = result.DRAW
	}
	return res
}

func (g *Game) Reset() {
	g.gameBoard.NewGame()
	g.header.Reset()
	g.result = result.UNKNOWNRESULT
	g.currentPlayer = g.player1
}

func (g *Game) GetCurrentPlayer() player.IPlayer {
	return g.currentPlayer
}

func (g *Game) IsCurrentPlayerHuman() bool {
	_, ok := g.currentPlayer.(*player.HumanPlayer)
	return ok
}

func (g *Game) IsCurrentPlayerAI() bool {
	_, ok := g.currentPlayer.(player.IAPlayer)
	return ok
}

func (g *Game) ChangePlayer() {
	if g.currentPlayer.GetValue() == 1 {
		g.currentPlayer = g.player2
	} else {
		g.currentPlayer = g.player1
	}
}
