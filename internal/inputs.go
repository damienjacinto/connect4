package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Inputs struct {
	direction Direction
	pressed   bool
}

type Direction uint

const (
	UNKNOWNDIRECTION Direction = iota
	NONE
	LEFT
	RIGHT
	QUIT
	DOWN
)

const (
	LeftString  = "left"
	RightString = "right"
	QuitString  = "quit"
	DownString  = "down"
	NoneString  = "none"
)

func (d Direction) String() string {
	switch d {
	case LEFT:
		return LeftString
	case RIGHT:
		return RightString
	case DOWN:
		return DownString
	case QUIT:
		return QuitString
	case NONE:
		return NoneString
	default:
		panic("unhandled default case")
	}
}

func ParseDirection(direction string) (Direction, error) {
	switch direction {
	case LeftString:
		return LEFT, nil
	case RightString:
		return RIGHT, nil
	case DownString:
		return DOWN, nil
	case QuitString:
		return QUIT, nil
	case NoneString:
		return NONE, nil
	default:
		return UNKNOWNDIRECTION, fmt.Errorf("unknown direction: %s", direction)
	}
}

func NewInput() *Inputs {
	return &Inputs{
		direction: NONE,
		pressed:   false,
	}
}

func (i *Inputs) HandleInput() Direction {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) && !i.pressed {
		i.pressed = true
		return LEFT
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) && !i.pressed {
		i.pressed = true
		return RIGHT
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) && !i.pressed {
		i.pressed = true
		return QUIT
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && !i.pressed {
		i.pressed = true
		return DOWN
	}

	i.pressed = false
	return NONE
}
