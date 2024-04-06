package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Input uint

const (
	UNKNOWNINPUT Input = iota
	LEFT
	RIGHT
	DOWN
)

const (
	LeftString  = "left"
	RightString = "right"
	DownString  = "down"
)

func (i Input) String() string {
	switch i {
	case LEFT:
		return LeftString
	case RIGHT:
		return RightString
	case DOWN:
		return DownString
	default:
		panic("unhandled default case")
	}
}

func ParseInput(input string) (Input, error) {
	switch input {
	case LeftString:
		return LEFT, nil
	case RightString:
		return RIGHT, nil
	case DownString:
		return DOWN, nil
	default:
		return UNKNOWNINPUT, fmt.Errorf("unknown input: %s", input)
	}
}

func HandleInput() Input {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		return LEFT
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		return RIGHT
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		return DOWN
	}
	return UNKNOWNINPUT
}
