package player

import (
	"image/color"
	"math/rand/v2"
	"time"

	"github.com/damienjacinto/connect4/internal/board"
)

type MinMaxIAPlayer struct {
	Player
	iatype iatype
}

func NewMinMaxIAPlayer(color color.RGBA, value int) IPlayer {
	return &RandomAIPlayer{
		Player{
			color: color,
			value: value,
			name:  MINMAX.String(),
		},
		MINMAX,
	}
}

func (p *MinMaxIAPlayer) Play(b ...board.Board) int {
	time.Sleep(1000 * time.Millisecond)
	return rand.IntN(6)
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
