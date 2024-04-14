package game

import "math/rand/v2"

const width, height = 7, 6

type Board struct {
	state [][]int
}

func NewBoard() *Board {
	board := make([][]int, height)
	for i := range board {
		board[i] = make([]int, width)
	}

	return &Board{
		state: board,
	}
}

func (b *Board) GetBoard() [][]int {
	return b.state
}

func (b *Board) Random() {
	for i := range height {
		for j := range width {
			b.state[i][j] = rand.IntN(3)
		}
	}
}

func (b *Board) Reset() {
	for i := range height {
		for j := range width {
			b.state[i][j] = 0
		}
	}
}

func (b *Board) GetWidth() int {
	return width
}

func (b *Board) GetHeigth() int {
	return height
}

func (b *Board) IsFinished() bool {
	// Check horizontal
	for i := 0; i < height; i++ {
		for j := 0; j < width-3; j++ {
			if b.state[i][j] != 0 && b.state[i][j] == b.state[i][j+1] && b.state[i][j] == b.state[i][j+2] && b.state[i][j] == b.state[i][j+3] {
				return true
			}
		}
	}
	// Check vertical
	for i := 0; i < height-3; i++ {
		for j := 0; j < width; j++ {
			if b.state[i][j] != 0 && b.state[i][j] == b.state[i+1][j] && b.state[i][j] == b.state[i+2][j] && b.state[i][j] == b.state[i+3][j] {
				return true
			}
		}
	}
	// Check diagonal
	for i := 0; i < height-3; i++ {
		for j := 0; j < width-3; j++ {
			if b.state[i][j] != 0 && b.state[i][j] == b.state[i+1][j+1] && b.state[i][j] == b.state[i+2][j+2] && b.state[i][j] == b.state[i+3][j+3] {
				return true
			}
		}
	}
	for i := 0; i < height-3; i++ {
		for j := 3; j < width; j++ {
			if b.state[i][j] != 0 && b.state[i][j] == b.state[i+1][j-1] && b.state[i][j] == b.state[i+2][j-2] && b.state[i][j] == b.state[i+3][j-3] {
				return true
			}
		}
	}
	return false
}

func (b *Board) IsFull() bool {
	for i := range height {
		for j := range width {
			if b.state[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func (b *Board) Copy() *Board {
	board := make([][]int, height)
	for i := range height {
		board[i] = make([]int, width)
		copy(board[i], b.state[i])
	}
	return &Board{
		state: board,
	}
}
