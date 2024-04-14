package game

import (
	"fmt"
	"image/color"

	"github.com/damienjacinto/connect4/internal/player"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	paddingHeader = 40.0
)

var (
	red    = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	yellow = color.RGBA{R: 255, G: 255, B: 0, A: 255}
)

type GameBoard struct {
	width, height int
	board         *Board
	colorBoard    color.RGBA
	sizeElement   float32
	padding       float32
	header        *Header
	player1       player.IPlayer
	player2       player.IPlayer
	currentPlayer player.IPlayer
}

func NewGameBoard(colorBoard color.RGBA, sizeElement float32, padding float32) *GameBoard {
	board := NewBoard()
	header := NewHeader(sizeElement, padding, board.GetWidth())
	player1 := player.NewHumanPlayer(red, 1, "Player 1")
	//player2 := NewHumanPlayer(yellow, 2, "Player 2")
	player2 := player.NewAIPlayer(yellow, 2, player.RANDOM)

	return &GameBoard{
		width:         width,
		height:        height,
		board:         NewBoard(),
		colorBoard:    colorBoard,
		sizeElement:   sizeElement,
		padding:       padding,
		header:        header,
		player1:       player1,
		player2:       player2,
		currentPlayer: player1,
	}
}

func (b *GameBoard) Print() {
	fmt.Println("-----------------")
	for i := range b.height {
		for j := range b.width {
			fmt.Printf("%d ", b.board.GetBoard()[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Println("-----------------")
}

func (b *GameBoard) Size() (float32, float32) {
	return float32(b.width)*b.sizeElement + float32((b.width-1))*b.padding, float32(b.height)*b.sizeElement + float32((b.height-1))*padding
}

func (b *GameBoard) Draw(screen *ebiten.Image, initx float32, inity float32, endscreen string, currentError string) {
	x := initx
	y := inity + 40

	if currentError != "" && endscreen == "" {
		ebitenutil.DebugPrint(screen, currentError)
	}

	if endscreen != "" {
		ebitenutil.DebugPrint(screen, endscreen)
	} else {
		middleScreen := initx + (float32(b.width)*b.sizeElement+float32((b.width-1))*b.padding)/2
		b.header.Draw(screen, middleScreen, paddingHeader, b.currentPlayer)
	}

	for i := 0; i < b.height; i++ {
		for j := 0; j < b.width; j++ {
			b.drawSlot(screen, x, y, b.sizeElement, b.colorBoard, b.board.GetBoard()[i][j])
			x += b.sizeElement + b.padding
		}
		x = initx
		y += b.sizeElement + b.padding
	}
}

func (b *GameBoard) drawSlot(screen *ebiten.Image, x, y, size float32, clr color.Color, player int) {
	vector.DrawFilledRect(screen, x, y, size, size, clr, true)
	switch player {
	case 1:
		vector.DrawFilledCircle(screen, x+size/2, y+size/2, size/2.5, b.player1.GetColor(), true)
	case 2:
		vector.DrawFilledCircle(screen, x+size/2, y+size/2, size/2.5, b.player2.GetColor(), true)
	default:
		vector.DrawFilledCircle(screen, x+size/2, y+size/2, size/2.5, color.Black, true)
	}
}

func (b *GameBoard) MoveLeft() {
	b.header.MoveLeft()
}

func (b *GameBoard) MoveRight() {
	b.header.MoveRight()
}

func (b *GameBoard) Play(col int) error {
	i := b.height - 1
	for ; i > 0; i-- {
		if b.board.GetBoard()[i][col] == 0 {
			break
		}
	}
	if b.board.GetBoard()[i][col] != 0 {
		return fmt.Errorf("column is full")
	}
	b.board.GetBoard()[i][col] = b.currentPlayer.GetValue()
	return nil
}

func (b *GameBoard) ChangePlayer() {
	if b.currentPlayer.GetValue() == 1 {
		b.currentPlayer = b.player2
	} else {
		b.currentPlayer = b.player1
	}
}

func (b *GameBoard) NewGame() {
	b.board.Reset()
	b.header.Reset()
	b.currentPlayer = b.player1
}

func (b *GameBoard) GetResult() Result {
	if b.board.IsFinished() {
		return Result(b.currentPlayer.GetValue())
	}
	if b.board.IsFull() {
		return DRAW
	}
	return UNKNOWNRESULT
}

func (b *GameBoard) IsCurrentPlayerHuman() bool {
	_, ok := b.currentPlayer.(*player.HumanPlayer)
	return ok
}

func (b *GameBoard) IsCurrentPlayerAI() bool {
	_, ok := b.currentPlayer.(player.IAPlayer)
	return ok
}
