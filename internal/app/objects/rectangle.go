package objects

import "github.com/hajimehoshi/ebiten/v2"

type Rectangle struct {
	X, Y, Width, Height float64
	Boundaries          [4]*Boundary
}

func NewRectangle(x, y, width, height float64) *Rectangle {
	return &Rectangle{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
		Boundaries: [4]*Boundary{
			NewBoundary(x, y, x+width, y),
			NewBoundary(x, y, x, y+height),
			NewBoundary(x+width, y, x+width, y+height),
			NewBoundary(x, y+height, x+width, y+height),
		},
	}
}

func (r *Rectangle) GetBoundaries() [4]*Boundary {
	return r.Boundaries
}

func (r *Rectangle) Show(screen *ebiten.Image) {
	for _, b := range r.Boundaries {
		b.Show(screen)
	}
}
