package objects

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ray struct {
	X, Y, Dx, Dy float64
}

func NewRay(pos, dir Vector) *Ray {
	return &Ray{
		X:  pos.X,
		Y:  pos.Y,
		Dx: dir.X,
		Dy: dir.Y,
	}
}

func (r *Ray) SetDir(x, y float64) {
	r.Dx = x - r.X
	r.Dy = y - r.Y
}

func (r *Ray) Show(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, r.X, r.Y, r.X+(r.Dx*10), r.Y+(r.Dy*10), color.White)
}

func (r *Ray) GetIntersection(wall *Boundary) (*Vector, bool) {
	x1 := wall.X1
	y1 := wall.Y1
	x2 := wall.X2
	y2 := wall.Y2

	x3 := r.X
	y3 := r.Y
	x4 := r.X + r.Dx
	y4 := r.Y + r.Dy

	denominator := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)
	if denominator == 0 {
		return nil, false
	}

	t := ((x1-x3)*(y3-y4) - (y1-y3)*(x3-x4)) / denominator
	u := -((x1-x2)*(y1-y3) - (y1-y2)*(x1-x3)) / denominator

	if t > 0 && t < 1 && u > 0 {
		pt := &Vector{}

		pt.X = x1 + t*(x2-x1)
		pt.Y = y1 + t*(y2-y1)

		return pt, true
	}

	return nil, false
}
