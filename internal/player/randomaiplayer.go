package player

import (
	"image/color"
	"math/rand/v2"
	"time"
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

func (p *RandomAIPlayer) Play() int {
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
