package board

import (
	"fmt"
)

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
	for i := height - 1; i >= 0; i-- {
		for j := width - 1; j >= 0; j-- {
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

func (b *Board) GetAvailableMoves() []int {
	var moves []int
	for i := range width {
		if b.state[0][i] == 0 {
			moves = append(moves, i)
		}
	}
	return moves
}

func (b *Board) Play(col int, player int) error {
	i := b.GetHeigth() - 1
	for ; i > 0; i-- {
		if b.state[i][col] == 0 {
			break
		}
	}
	if b.state[i][col] != 0 {
		return fmt.Errorf("column is full")
	}
	b.state[i][col] = player
	return nil
}

func (b *Board) String() string {
	str := ""
	for i := range height {
		if i > 3 {
			for j := range width {
				str += fmt.Sprintf("%d ", b.state[i][j])
			}
			str += "\n"
		}
	}
	return str
}

func (b *Board) evaluateLine(player, val1, val2, val3, val4 int) int {
	countPlayer := 0
	countOpposite := 0

	// Count the pieces in the line
	for _, cell := range []int{val1, val2, val3, val4} {
		if cell == player {
			countPlayer++
		} else if cell != 0 {
			countOpposite++
		}
	}

	// Score the line based on the counts
	if countPlayer == 4 {
		return 40000
	} else if countPlayer == 3 && countOpposite == 0 {
		return 1000
	} else if countPlayer == 2 && countOpposite < 2 {
		return 200
	} else if countOpposite == 2 && countPlayer < 2 {
		return -500
	} else if countOpposite == 3 && countPlayer == 0 {
		return -2000
	} else if countOpposite == 4 {
		return -20000
	}
	return 0
}

func (b *Board) GetEvaluation(player int) int {

	var score int = 0
	// Check horizontal
	for i := 0; i < height; i++ {
		for j := 0; j < width-3; j++ {
			score += b.evaluateLine(player, b.state[i][j], b.state[i][j+1], b.state[i][j+2], b.state[i][j+3])
		}
	}
	// Check vertical
	for i := 0; i < height-3; i++ {
		for j := 0; j < width; j++ {
			score += b.evaluateLine(player, b.state[i][j], b.state[i+1][j], b.state[i+2][j], b.state[i+3][j])
		}
	}
	// Check diagonal
	for i := 0; i < height-3; i++ {
		for j := 0; j < width-3; j++ {
			score += b.evaluateLine(player, b.state[i][j], b.state[i+1][j+1], b.state[i+2][j+2], b.state[i+3][j+3])
		}
	}

	for i := 0; i < height-3; i++ {
		for j := 3; j < width; j++ {
			score += b.evaluateLine(player, b.state[i][j], b.state[i+1][j-1], b.state[i+2][j-2], b.state[i+3][j-3])
		}
	}

	// Better valuation for middle columns
	for i := 0; i < height; i++ {
		for j := 3; j < width-2; j++ {
			if b.state[i][j] == player {
				score += 1
				if j == 3 {
					score += 1
				}
			} else if b.state[i][j] != 0 {
				score -= 2
				if j == 3 {
					score -= 2
				}
			}
		}
	}
	return score
}
