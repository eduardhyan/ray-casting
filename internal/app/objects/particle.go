package objects

import (
	"image/color"
	"math"
	"ray-casting/pkg/geometry"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Particle struct {
	Heading float64
	Pos     Vector
	Rays    []*Ray
}

func NewParticle(pos Vector) *Particle {
	p := &Particle{
		Heading: -30,
		Pos:     pos,
	}

	return p
}

func (p *Particle) Show(screen *ebiten.Image) {
	p.InitRays()

	for _, ray := range p.Rays {
		ray.X = p.Pos.X
		ray.Y = p.Pos.Y

		ray.Show(screen)
	}
}

func (p *Particle) Update(x, y float64) {
	p.Pos.X = x
	p.Pos.Y = y
}

func (p *Particle) InitRays() {
	p.Rays = make([]*Ray, 0)

	for angle := p.Heading; angle <= p.Heading+90; angle += 0.5 {
		ray := NewRay(p.Pos, *VectorFromAngle(geometry.Radian(float64(angle))))
		p.Rays = append(p.Rays, ray)
	}
}

func (p *Particle) Turn(angle float64) {
	p.Heading += angle
}

func (p *Particle) Look(screen *ebiten.Image, walls []*Boundary) {

	for _, ray := range p.Rays {
		minDistance := math.Inf(1)
		var closest *Vector

		for _, wall := range walls {
			pt, ok := ray.GetIntersection(wall)

			if ok {
				d := geometry.Distance(ray.X, ray.Y, pt.X, pt.Y)

				if d < minDistance {
					minDistance = d
					closest = pt
				}
			}

		}

		if closest != nil {
			ebitenutil.DrawLine(screen, ray.X, ray.Y, closest.X, closest.Y, color.RGBA{0, 255, 0, 100})
		}
	}
}
