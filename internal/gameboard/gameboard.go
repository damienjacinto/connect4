package gameboard

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type GameBoard struct {
	board       *Board
	colorBoard  color.RGBA
	sizeElement float32
	padding     float32
}

func NewGameBoard(width int, height int, colorBoard color.RGBA, sizeElement float32, padding float32) *GameBoard {
	board := NewBoard(width, height)
	return &GameBoard{
		board:       board,
		colorBoard:  colorBoard,
		sizeElement: sizeElement,
		padding:     padding,
	}
}

func (b *GameBoard) Print() {
	fmt.Println("-----------------")
	for i := range b.board.GetHeigth() {
		for j := range b.board.GetWidth() {
			fmt.Printf("%d ", b.board.GetBoard()[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Println("-----------------")
}

func (b *GameBoard) Size() (float32, float32) {
	return float32(b.board.GetWidth())*b.sizeElement + float32((b.board.GetWidth()-1))*b.padding, float32(b.board.GetHeigth())*b.sizeElement + float32((b.board.GetHeigth()-1))*b.padding
}

func (b *GameBoard) Draw(screen *ebiten.Image, wScreen int, hScreen int, player1Color color.RGBA, player2Color color.RGBA) {
	wBoard, hBoard := b.Size()
	initx := (float32(wScreen) - wBoard) / 2
	inity := (float32(hScreen) - hBoard) / 2

	x := initx
	y := inity + 40

	for i := 0; i < b.board.GetHeigth(); i++ {
		for j := 0; j < b.board.GetWidth(); j++ {
			b.drawSlot(screen, x, y, b.sizeElement, b.colorBoard, b.board.GetBoard()[i][j], player1Color, player2Color)
			x += b.sizeElement + b.padding
		}
		x = initx
		y += b.sizeElement + b.padding
	}
}

func (b *GameBoard) drawSlot(screen *ebiten.Image, x, y, size float32, clr color.Color, player int, player1Color color.RGBA, player2Color color.RGBA) {
	vector.DrawFilledRect(screen, x, y, size, size, clr, true)
	switch player {
	case 1:
		vector.DrawFilledCircle(screen, x+size/2, y+size/2, size/2.5, player1Color, true)
	case 2:
		vector.DrawFilledCircle(screen, x+size/2, y+size/2, size/2.5, player2Color, true)
	default:
		vector.DrawFilledCircle(screen, x+size/2, y+size/2, size/2.5, color.Black, true)
	}
}

func (b *GameBoard) Play(col int, playerValue int) error {
	return b.board.Play(col, playerValue)
}

func (b *GameBoard) NewGame() {
	b.board.Reset()
}

func (b *GameBoard) GetBoard() *Board {
	return b.board
}
