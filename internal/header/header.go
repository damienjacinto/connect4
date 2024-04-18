package header

import (
	"image/color"

	"github.com/damienjacinto/connect4/internal/player"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	paddingHeader float32 = 40.0
)

type Header struct {
	positionPiece int
	sizeElement   float32
	boardPadding  float32
	width         int
	display       string
}

func NewHeader(size float32, boardPadding float32, widthBoard int) *Header {
	return &Header{
		sizeElement:   size / 1.5,
		positionPiece: widthBoard / 2,
		width:         widthBoard,
		boardPadding:  boardPadding,
		display:       "",
	}
}

func (h *Header) SetDisplay(display string) {
	h.display = display
}

func (h *Header) Draw(screen *ebiten.Image, initx float32, currentPlayer player.IPlayer) {
	x := initx + float32(h.positionPiece-(h.width/2))*(h.sizeElement*1.5+h.boardPadding)
	if h.display != "" {
		ebitenutil.DebugPrint(screen, h.display)
	} else {
		ebitenutil.DebugPrint(screen, currentPlayer.GetName()+" turn")
	}
	drawPiece(screen, x, paddingHeader, h.sizeElement, currentPlayer.GetColor())
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

func drawPiece(screen *ebiten.Image, x, y, size float32, currentPlayerColor color.RGBA) {
	vector.DrawFilledCircle(screen, x, y, size/2, currentPlayerColor, true)
}

func (h *Header) Reset() {
	h.positionPiece = h.width / 2
}

func (h *Header) GetPositionPiece() int {
	return h.positionPiece
}
