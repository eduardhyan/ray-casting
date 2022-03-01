package objects

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Boundary struct {
	X1, Y1, X2, Y2 float64
}

func NewBoundary(x1, y1, x2, y2 float64) *Boundary {
	return &Boundary{
		X1: x1,
		Y1: y1,
		X2: x2,
		Y2: y2,
	}
}

func (b *Boundary) Show(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, b.X1, b.Y1, b.X2, b.Y2, color.White)
}
