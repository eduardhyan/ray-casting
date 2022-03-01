package main

import (
	"math/rand"
	"time"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

var walls = []*Boundary{}

// var ray = NewRay(300, 300, 10, 0)
var particle = NewParticle()

type Game struct {
	screen *ebiten.Image
}

func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		x := rand.Float64() * float64(screenWidth)
		y := rand.Float64() * float64(screenHeight)
		width := rand.Float64() * float64(screenWidth)
		height := rand.Float64() * float64(screenHeight)

		walls = append(walls, NewBoundary(x, y, x+width, y+height))
	}

	walls = append(walls, NewBoundary(0, 0, screenWidth, 0))
	walls = append(walls, NewBoundary(0, 0, 0, screenHeight))
	walls = append(walls, NewBoundary(0, screenHeight, screenWidth, screenHeight))
	walls = append(walls, NewBoundary(screenWidth, 0, screenWidth, screenHeight))

	return &Game{}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
	mouseX, mouseY := ebiten.CursorPosition()
	particle.Update(float64(mouseX), float64(mouseY))
	// ray.SetDir(float64(mouseX), float64(mouseY))

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.screen = screen
	// ebitenutil.DebugPrint(screen, "Hello, World!")

	particle.Show(screen)

	for _, wall := range walls {
		wall.Show(screen)
	}

	particle.Look(screen, walls)

	// ray.Show(screen)

	// pt, isIntersecting := ray.GetIntersection(wall)

	// if isIntersecting {
	// 	// ebitenutil.DrawRect(screen, pt.x, pt.y, 5, 5, color.White)
	// 	// screen.Set(int(pt.x), int(pt.y), color.White)
	// 	g.Circle(pt.x, pt.y, 5)
	// }

}

func (g *Game) Circle(x, y, r float64) {
	dc := gg.NewContext(screenWidth, screenHeight)
	dc.DrawCircle(x, y, r)
	dc.SetRGB(1, 1, 1)
	dc.Fill()
	dc.Stroke()

	img := ebiten.NewImageFromImage(dc.Image())
	g.screen.DrawImage(img, &ebiten.DrawImageOptions{})
}
