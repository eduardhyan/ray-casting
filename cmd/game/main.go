package main

import (
	"ray-casting/internal/app"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1200
	screenHeight = 800
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ray casting")

	g := app.NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
