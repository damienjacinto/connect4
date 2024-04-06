package game

import (
	"fmt"
	"image/color"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	width, height = 7, 6
	paddingHeader = 40.0
)

type Board struct {
	width, height int
	board         [][]int
	colorBoard    color.RGBA
	sizeElement   float32
	padding       float32
	header        *Header
	currentPlayer int
}

func NewBoard(colorBoard color.RGBA, sizeElement float32, padding float32) *Board {
	board := make([][]int, height)
	for i := range board {
		board[i] = make([]int, width)
	}

	header := NewHeader(sizeElement, padding, width)

	return &Board{
		width:         width,
		height:        height,
		board:         board,
		colorBoard:    colorBoard,
		sizeElement:   sizeElement,
		padding:       padding,
		header:        header,
		currentPlayer: 1,
	}
}

func (b *Board) Random() {
	for i := range b.height {
		for j := range b.width {
			b.board[i][j] = rand.IntN(3)
		}
	}
}

func (b *Board) Print() {
	fmt.Println("-----------------")
	for i := range b.height {
		for j := range b.width {
			fmt.Printf("%d ", b.board[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Println("-----------------")
}

func (b *Board) Size() (float32, float32) {
	return float32(b.width)*b.sizeElement + float32((b.width-1))*b.padding, float32(b.height)*b.sizeElement + float32((b.height-1))*padding
}

func (b *Board) Draw(screen *ebiten.Image, initx float32, inity float32) {
	x := initx
	y := inity + 40

	middleScreen := initx + (float32(b.width)*b.sizeElement+float32((b.width-1))*b.padding)/2
	b.header.Draw(screen, middleScreen, paddingHeader, b.currentPlayer)

	for i := 0; i < b.height; i++ {
		for j := 0; j < b.width; j++ {
			drawSlot(screen, x, y, b.sizeElement, b.colorBoard, b.board[i][j])
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

func (b *Board) GetWidth() int {
	return b.width
}

func (b *Board) MoveLeft() {
	b.header.MoveLeft()
}

func (b *Board) MoveRight() {
	b.header.MoveRight()
}

func (b *Board) DropPiece() error {
	i := b.height - 1
	for ; i > 0; i-- {
		if b.board[i][b.header.positionPiece] == 0 {
			break
		}
	}
	if b.board[i][b.header.positionPiece] != 0 {
		return fmt.Errorf("column is full")
	}
	b.board[i][b.header.positionPiece] = b.currentPlayer
	return nil
}

func (b *Board) IsFinished() bool {
	// Check horizontal
	for i := 0; i < b.height; i++ {
		for j := 0; j < b.width-3; j++ {
			if b.board[i][j] != 0 && b.board[i][j] == b.board[i][j+1] && b.board[i][j] == b.board[i][j+2] && b.board[i][j] == b.board[i][j+3] {
				return true
			}
		}
	}
	// Check vertical
	for i := 0; i < b.height-3; i++ {
		for j := 0; j < b.width; j++ {
			if b.board[i][j] != 0 && b.board[i][j] == b.board[i+1][j] && b.board[i][j] == b.board[i+2][j] && b.board[i][j] == b.board[i+3][j] {
				return true
			}
		}
	}
	// Check diagonal
	for i := 0; i < b.height-3; i++ {
		for j := 0; j < b.width-3; j++ {
			if b.board[i][j] != 0 && b.board[i][j] == b.board[i+1][j+1] && b.board[i][j] == b.board[i+2][j+2] && b.board[i][j] == b.board[i+3][j+3] {
				return true
			}
		}
	}
	for i := 0; i < b.height-3; i++ {
		for j := 3; j < b.width; j++ {
			if b.board[i][j] != 0 && b.board[i][j] == b.board[i+1][j-1] && b.board[i][j] == b.board[i+2][j-2] && b.board[i][j] == b.board[i+3][j-3] {
				return true
			}
		}
	}
	return false
}

func (b *Board) ChangePlayer() {
	b.currentPlayer = b.currentPlayer + 1
	if b.currentPlayer > 2 {
		b.currentPlayer = 1
	}
}

func (b *Board) Reset() {
	for i := range b.height {
		for j := range b.width {
			b.board[i][j] = 0
		}
	}
	b.header.Reset()
	b.currentPlayer = 1
}
