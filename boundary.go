package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Boundary struct {
	x1, y1, x2, y2 float64
}

func NewBoundary(x1, y1, x2, y2 float64) *Boundary {
	return &Boundary{
		x1: x1,
		y1: y1,
		x2: x2,
		y2: y2,
	}
}

func (b *Boundary) Show(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, b.x1, b.y1, b.x2, b.y2, color.White)
}
