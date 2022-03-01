package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1200
	screenHeight = 800
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ray casting")

	g := NewGame()
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
