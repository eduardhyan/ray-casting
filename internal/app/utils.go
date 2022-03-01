package app

import (
	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Circle(x, y, r float64) {
	dc := gg.NewContext(g.width, g.height)
	dc.DrawCircle(x, y, r)
	dc.SetRGB(1, 1, 1)
	dc.Fill()
	dc.Stroke()

	img := ebiten.NewImageFromImage(dc.Image())
	g.screen.DrawImage(img, &ebiten.DrawImageOptions{})
}
