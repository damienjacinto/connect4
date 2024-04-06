package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Header struct {
	positionPiece int
	sizeElement   float32
	boardPadding  float32
	width         int
}

func NewHeader(size float32, boardPadding float32, widthBoard int) *Header {
	return &Header{
		sizeElement:   size / 1.5,
		positionPiece: widthBoard / 2,
		width:         widthBoard,
		boardPadding:  boardPadding,
	}
}

func (h *Header) Draw(screen *ebiten.Image, initx float32, inity float32, currentPlayer int) {
	x := initx + float32(h.positionPiece-(h.width/2))*(h.sizeElement*1.5+h.boardPadding)
	y := inity
	drawPiece(screen, x, y, h.sizeElement, currentPlayer)
}

func (h *Header) MoveLeft() {
	if h.positionPiece > 0 {
		h.positionPiece--
	}
}

func (h *Header) MoveRight() {
	if h.positionPiece < h.width-1 {
		h.positionPiece++
	}
}

func drawPiece(screen *ebiten.Image, x, y, size float32, currentPlayer int) {
	switch currentPlayer {
	case 1:
		vector.DrawFilledCircle(screen, x, y, size/2, player1, true)
	case 2:
		vector.DrawFilledCircle(screen, x, y, size/2, player2, true)
	}
}

func (h *Header) Reset() {
	h.positionPiece = h.width / 2
}
