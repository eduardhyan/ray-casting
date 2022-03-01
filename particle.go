package main

import (
	"image/color"
	"ray-casting/pkg/geometry"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Particle struct {
	pos  Vector
	rays []*Ray
}

func NewParticle() *Particle {
	p := &Particle{}
	p.pos = Vector{
		x: screenWidth / 2,
		y: screenHeight / 2,
	}

	for angle := 0; angle < 360; angle += 1 {
		ray := NewRay(p.pos, *VectorFromAngle(float64(angle)))
		p.rays = append(p.rays, ray)
	}

	return p
}

func (p *Particle) Show(screen *ebiten.Image) {
	for _, ray := range p.rays {
		ray.x = p.pos.x
		ray.y = p.pos.y

		ray.Show(screen)
	}
}

func (p *Particle) Update(x, y float64) {
	p.pos.x = x
	p.pos.y = y
}

func (p *Particle) Look(screen *ebiten.Image, walls []*Boundary) {

	for _, ray := range p.rays {
		minDistance := float64(screenWidth)
		var closest *Vector

		for _, wall := range walls {
			pt, ok := ray.GetIntersection(wall)

			if ok {
				d := geometry.Distance(ray.x, ray.y, pt.x, pt.y)

				if d < minDistance {
					minDistance = d
					closest = pt
				}
			}

		}

		if closest != nil {
			ebitenutil.DrawLine(screen, ray.x, ray.y, closest.x, closest.y, color.RGBA{255, 255, 255, 123})
			// ray.SetDir(intersection.x, intersection.y)
		}
	}
}
