package player

import (
	"fmt"
	"image/color"
	"math"
	"time"

	"github.com/damienjacinto/connect4/internal/gameboard"
)

type AlphaBetaIAPlayer struct {
	Player
	iatype iatype
}

func NewAlphaBetaIAPlayer(color color.RGBA, value int) IAPlayer {
	return &AlphaBetaIAPlayer{
		Player{
			color: color,
			value: value,
			name:  ALPHABETA.String(),
		},
		ALPHABETA,
	}
}

func (p *AlphaBetaIAPlayer) construct(n *gameboard.Node, depth int, player int) *gameboard.Node {
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

func (p *AlphaBetaIAPlayer) alphabeta(n *gameboard.Node, player int, alpha int, beta int) (int, int) {
	var score int = 0
	var bestMove int = n.Move
	if len(n.Childs) == 0 {
		return bestMove, n.Data.GetEvaluation(p.value)
	} else {
		nextPlayer := (player % 2) + 1
		if nextPlayer != p.value {
			score = math.MinInt
			for _, child := range n.Childs {
				_, eval := p.alphabeta(child, nextPlayer, alpha, beta)
				if eval >= beta {
					score = beta
					bestMove = child.Move
					break
				}
				if eval > score {
					score = eval
					bestMove = child.Move
				}
			}
		} else {
			score = math.MaxInt
			for _, child := range n.Childs {
				_, eval := p.alphabeta(child, nextPlayer, alpha, beta)
				if eval <= alpha {
					score = alpha
					bestMove = child.Move
					break
				}
				if eval < score {
					score = eval
					bestMove = child.Move
				}
			}
		}
		return bestMove, score
	}
}

func (p *AlphaBetaIAPlayer) Play(b *gameboard.Board) int {
	time.Sleep(3 * time.Second)
	startTime := time.Now()
	depth := 0
	tree := gameboard.NewNode(b, depth, 0)
	p.construct(tree, depth, p.value)
	fmt.Println("Time to construct : ", time.Since(startTime))

	move, _ := p.alphabeta(tree, p.value, -20000, 4000)
	fmt.Println("Time to play : ", time.Since(startTime))
	return move
}

func (p *AlphaBetaIAPlayer) GetColor() color.RGBA {
	return p.color
}

func (p *AlphaBetaIAPlayer) GetValue() int {
	return p.value
}

func (p *AlphaBetaIAPlayer) GetName() string {
	return p.name
}

func (p *AlphaBetaIAPlayer) GetType() iatype {
	return p.iatype
}
