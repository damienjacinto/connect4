package player

import "image/color"

type IPlayer interface {
	GetColor() color.RGBA
	GetValue() int
	GetName() string
	Play() int
}

type Player struct {
	color color.RGBA
	value int
	name  string
}

type HumanPlayer struct {
	player Player
	move   int
}

func NewHumanPlayer(color color.RGBA, value int, name string) *HumanPlayer {
	return &HumanPlayer{
		player: Player{
			color: color,
			value: value,
			name:  name,
		},
		move: 0,
	}
}

func (p *HumanPlayer) SetMove(move int) {
	p.move = move
}

func (p *HumanPlayer) Play() int {
	return p.move
}

func (p *HumanPlayer) GetColor() color.RGBA {
	return p.player.color
}

func (p *HumanPlayer) GetValue() int {
	return p.player.value
}

func (p *HumanPlayer) GetName() string {
	return p.player.name
}