package player

import (
	"image/color"
)

type IAPlayer interface {
	IPlayer
}

func NewAIPlayer(color color.RGBA, value int, ia iatype) IAPlayer {
	switch ia {
	case RANDOM:
		return NewRandomAIPlayer(color, value)
	default:
		return NewRandomAIPlayer(color, value)
	}
}
