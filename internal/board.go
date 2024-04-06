package game

import "fmt"

type Board struct {
	width, height int
	board         [][]int
}

func NewBoard(width, height int) *Board {
	board := make([][]int, height)
	for i := range board {
		board[i] = make([]int, width)
	}
	return &Board{
		width:  width,
		height: height,
		board:  board,
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
