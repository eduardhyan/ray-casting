package app

import (
	"math/rand"
	"ray-casting/internal/app/objects"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var walls = []*objects.Boundary{}
var particle = objects.NewParticle(objects.Vector{X: 300, Y: 600})

var eventListener = &EventListener{}
var speed float64 = 5

type Game struct {
	screen *ebiten.Image
	width  int
	height int
}

func NewGame(width, height int) *Game {
	rand.Seed(time.Now().UnixNano())

	boxes := [...]*objects.Rectangle{
		objects.NewRectangle(50, 50, 100, 100),
		objects.NewRectangle(500, 50, 100, 50),
		objects.NewRectangle(800, 50, 300, 80),
		objects.NewRectangle(300, 200, 100, 200),
		objects.NewRectangle(float64(width/2), float64(height/2), 50, 50),
		objects.NewRectangle(750, 300, 100, 400),
		objects.NewRectangle(50, 550, 100, 200),
		objects.NewRectangle(0, 0, float64(width), float64(height)),
	}

	for _, box := range boxes {
		boundaries := box.GetBoundaries()
		walls = append(walls, boundaries[:]...)
	}

	eventListener.Add(ebiten.KeyA, func() {
		particle.Turn(-1)
	})

	eventListener.Add(ebiten.KeyD, func() {
		particle.Turn(1)
	})

	eventListener.Add(ebiten.KeyArrowLeft, func() {
		particle.Update(particle.Pos.X-speed, particle.Pos.Y)

		if particle.Pos.X < 0 {
			particle.Pos.X = float64(width)
		}
	})

	eventListener.Add(ebiten.KeyArrowRight, func() {
		particle.Update(particle.Pos.X+speed, particle.Pos.Y)

		if particle.Pos.X > float64(width) {
			particle.Pos.X = 0
		}
	})

	eventListener.Add(ebiten.KeyArrowUp, func() {
		particle.Update(particle.Pos.X, particle.Pos.Y-speed)

		if particle.Pos.Y < 0 {
			particle.Pos.Y = float64(height)
		}
	})

	eventListener.Add(ebiten.KeyArrowDown, func() {
		particle.Update(particle.Pos.X, particle.Pos.Y+speed)

		if particle.Pos.Y > float64(height) {
			particle.Pos.Y = 0
		}
	})

	return &Game{
		width:  width,
		height: height,
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}

func (g *Game) Update() error {
	// mouseX, mouseY := ebiten.CursorPosition()
	// particle.Update(float64(mouseX), float64(mouseY))

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.screen = screen

	particle.Show(screen)

	eventListener.Handle()

	for _, wall := range walls {
		wall.Show(screen)
	}

	particle.Look(screen, walls)
}
