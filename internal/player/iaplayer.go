package player

import (
	"image/color"
)

type IAPlayer interface {
	IPlayer
	GetType() iatype
}

func NewAIPlayer(color color.RGBA, value int, ia iatype) IAPlayer {
	switch ia {
	case RANDOM:
		return NewRandomAIPlayer(color, value).(IAPlayer)
	case MINMAX:
		return NewMinMaxIAPlayer(color, value).(IAPlayer)
	default:
		return NewRandomAIPlayer(color, value).(IAPlayer)
	}
}
