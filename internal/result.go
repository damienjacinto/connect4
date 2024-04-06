package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Result uint

const (
	UNKNOWNRESULT Result = iota
	PLAYER1
	PLAYER2
	DRAW
)

const (
	DrawString     = "drawn"
	Player1tString = "player1"
	Player2String  = "player2"
)

func (r Result) String() string {
	switch r {
	case DRAW:
		return DrawString
	case PLAYER1:
		return Player1tString
	case PLAYER2:
		return Player2String
	default:
		panic("unhandled default case")
	}
}

func ParseResult(result string) (Result, error) {
	switch result {
	case DrawString:
		return DRAW, nil
	case Player1tString:
		return PLAYER1, nil
	case Player2String:
		return PLAYER2, nil
	default:
		return UNKNOWNRESULT, fmt.Errorf("unknown result: %s", result)
	}
}

func (r Result) Draw(screen *ebiten.Image) {
	switch r {
	case PLAYER1:
		ebitenutil.DebugPrint(screen, "Game over, player 1 wins !")
	case PLAYER2:
		ebitenutil.DebugPrint(screen, "Game over, player 2 wins !")
	case DRAW:
		ebitenutil.DebugPrint(screen, "Game over, draw !")
	}
}
