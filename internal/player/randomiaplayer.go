package player

import (
	"image/color"
	"math/rand/v2"
	"time"

	"github.com/damienjacinto/connect4/internal/board"
)

type RandomAIPlayer struct {
	Player
	iatype iatype
}

func NewRandomAIPlayer(color color.RGBA, value int) IPlayer {
	return &RandomAIPlayer{
		Player{
			color: color,
			value: value,
			name:  RANDOM.String(),
		},
		RANDOM,
	}
}

func (p *RandomAIPlayer) Play(b ...board.Board) int {
	time.Sleep(1000 * time.Millisecond)
	return rand.IntN(6)
}

func (p *RandomAIPlayer) GetColor() color.RGBA {
	return p.color
}

func (p *RandomAIPlayer) GetValue() int {
	return p.value
}

func (p *RandomAIPlayer) GetName() string {
	return p.name
}

func (p *RandomAIPlayer) GetType() iatype {
	return p.iatype
}
