package player

import (
	"fmt"
	"image/color"
	"math"
	"time"

	"github.com/damienjacinto/connect4/internal/gameboard"
)

type MinMaxIAPlayer struct {
	Player
	iatype iatype
}

func NewMinMaxIAPlayer(color color.RGBA, value int) IAPlayer {
	return &MinMaxIAPlayer{
		Player{
			color: color,
			value: value,
			name:  MINMAX.String(),
		},
		MINMAX,
	}
}

func (p *MinMaxIAPlayer) construct(n *gameboard.Node, depth int, player int) *gameboard.Node {
	if depth < maxDepth && !n.Data.IsFinished() && !n.Data.IsFull() {
		depth++
		nextPlayer := (player % 2) + 1

		moves := n.Data.GetAvailableMoves()
		for _, m := range moves {
			newBoard := n.Data.Copy()
			newBoard.Play(m, player)
			child := gameboard.NewNode(newBoard, depth, m)
			n.AddChild(p.construct(child, depth, nextPlayer))
		}
	}
	return n
}

func (p *MinMaxIAPlayer) minmax(n *gameboard.Node, player int) (int, int) {
	var score int = 0
	var bestMove int = n.Move
	if len(n.Childs) == 0 {
		return bestMove, n.Data.GetEvaluation(p.value)
	} else {
		nextPlayer := (player % 2) + 1
		if nextPlayer != p.value {
			score = math.MinInt
			for _, child := range n.Childs {
				_, eval := p.minmax(child, nextPlayer)
				if eval > score {
					score = eval
					bestMove = child.Move
				}
			}
		} else {
			score = math.MaxInt
			for _, child := range n.Childs {
				_, eval := p.minmax(child, nextPlayer)
				if eval < score {
					score = eval
					bestMove = child.Move
				}
			}
		}
		return bestMove, score
	}
}

func (p *MinMaxIAPlayer) Play(b *gameboard.Board) int {
	startTime := time.Now()
	depth := 0
	tree := gameboard.NewNode(b, depth, 0)
	p.construct(tree, depth, p.value)

	move, _ := p.minmax(tree, p.value)
	fmt.Println("Time to play : ", time.Since(startTime))
	return move
}

func (p *MinMaxIAPlayer) GetColor() color.RGBA {
	return p.color
}

func (p *MinMaxIAPlayer) GetValue() int {
	return p.value
}

func (p *MinMaxIAPlayer) GetName() string {
	return p.name
}

func (p *MinMaxIAPlayer) GetType() iatype {
	return p.iatype
}
