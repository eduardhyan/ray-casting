package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ray struct {
	x, y, dx, dy float64
}

func NewRay(pos, dir Vector) *Ray {
	return &Ray{
		x:  pos.x,
		y:  pos.y,
		dx: dir.x,
		dy: dir.y,
	}
}

func (r *Ray) SetDir(x, y float64) {
	r.dx = x - r.x
	r.dy = y - r.y
}

func (r *Ray) Show(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, r.x, r.y, r.x+(r.dx*10), r.y+(r.dy*10), color.White)
}

func (r *Ray) GetIntersection(wall *Boundary) (*Vector, bool) {
	x1 := wall.x1
	y1 := wall.y1
	x2 := wall.x2
	y2 := wall.y2

	x3 := r.x
	y3 := r.y
	x4 := r.x + r.dx
	y4 := r.y + r.dy

	denominator := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)
	if denominator == 0 {
		return nil, false
	}

	t := ((x1-x3)*(y3-y4) - (y1-y3)*(x3-x4)) / denominator
	u := -((x1-x2)*(y1-y3) - (y1-y2)*(x1-x3)) / denominator

	if t > 0 && t < 1 && u > 0 {
		pt := &Vector{}

		pt.x = x1 + t*(x2-x1)
		pt.y = y1 + t*(y2-y1)

		return pt, true
	}

	return nil, false
}
