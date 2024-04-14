package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	paddingHeader = 40.0
)

type GameBoard struct {
	width, height int
	board         *Board
	colorBoard    color.RGBA
	sizeElement   float32
	padding       float32
	header        *Header
	currentPlayer int
}

func NewGameBoard(colorBoard color.RGBA, sizeElement float32, padding float32) *GameBoard {
	header := NewHeader(sizeElement, padding, width)

	return &GameBoard{
		width:         width,
		height:        height,
		board:         NewBoard(),
		colorBoard:    colorBoard,
		sizeElement:   sizeElement,
		padding:       padding,
		header:        header,
		currentPlayer: 1,
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

func (b *GameBoard) Draw(screen *ebiten.Image, initx float32, inity float32) {
	x := initx
	y := inity + 40

	middleScreen := initx + (float32(b.width)*b.sizeElement+float32((b.width-1))*b.padding)/2
	b.header.Draw(screen, middleScreen, paddingHeader, b.currentPlayer)

	for i := 0; i < b.height; i++ {
		for j := 0; j < b.width; j++ {
			drawSlot(screen, x, y, b.sizeElement, b.colorBoard, b.board.GetBoard()[i][j])
			x += b.sizeElement + b.padding
		}
		x = initx
		y += b.sizeElement + b.padding
	}
}

func drawSlot(screen *ebiten.Image, x, y, size float32, clr color.Color, player int) {
	vector.DrawFilledRect(screen, x, y, size, size, clr, true)
	switch player {
	case 1:
		vector.DrawFilledCircle(screen, x+size/2, y+size/2, size/2.5, player1, true)
	case 2:
		vector.DrawFilledCircle(screen, x+size/2, y+size/2, size/2.5, player2, true)
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

func (b *GameBoard) DropPiece() error {
	i := b.height - 1
	for ; i > 0; i-- {
		if b.board.GetBoard()[i][b.header.positionPiece] == 0 {
			break
		}
	}
	if b.board.GetBoard()[i][b.header.positionPiece] != 0 {
		return fmt.Errorf("column is full")
	}
	b.board.GetBoard()[i][b.header.positionPiece] = b.currentPlayer
	return nil
}

func (b *GameBoard) ChangePlayer() {
	b.currentPlayer = b.currentPlayer + 1
	if b.currentPlayer > 2 {
		b.currentPlayer = 1
	}
}

func (b *GameBoard) NewGame() {
	b.board.Reset()
	b.header.Reset()
	b.currentPlayer = 1
}

func (b *GameBoard) GetResult() Result {
	if b.board.IsFinished() {
		return Result(b.currentPlayer)
	}
	if b.board.IsFull() {
		return DRAW
	}
	return UNKNOWNRESULT
}
