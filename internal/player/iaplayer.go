package player

import (
	"fmt"
	"image/color"
)

type IAPlayer interface {
	IPlayer
	GetType() iatype
}

func NewAIPlayer(color color.RGBA, value int, ia iatype) IAPlayer {
	switch ia {
	case RANDOM:
		fmt.Println("Random computer")
		return NewRandomAIPlayer(color, value)
	case MINMAX:
		fmt.Println("Minmax computer")
		return NewMinMaxIAPlayer(color, value)
	default:
		fmt.Println("Default computer")
		return NewRandomAIPlayer(color, value)
	}
}
