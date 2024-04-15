package player

import (
	"fmt"
	"image/color"
)

const maxDepth = 5

type IAPlayer interface {
	IPlayer
	GetType() iatype
}

func NewAIPlayer(color color.RGBA, value int, ia iatype) IAPlayer {
	fmt.Printf("%s computer", ia.String())
	switch ia {
	case RANDOM:
		return NewRandomAIPlayer(color, value)
	case MINMAX:
		return NewMinMaxIAPlayer(color, value)
	case ALPHABETA:
		return NewAlphaBetaIAPlayer(color, value)
	default:
		return NewRandomAIPlayer(color, value)
	}
}
