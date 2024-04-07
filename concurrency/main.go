package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth, screenHeight = 640, 360
	birdCount                 = 500
)

var (
	green = color.RGBA{R: 10, G: 255, B: 50, A: 255}
	birds [birdCount]*Bird
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, bird := range birds {
		screen.Set(int(bird.position.x+1), int(bird.position.y), green)
		screen.Set(int(bird.position.x-1), int(bird.position.y), green)
		screen.Set(int(bird.position.x), int(bird.position.y-1), green)
		screen.Set(int(bird.position.x), int(bird.position.y+1), green)
	}
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return screenWidth, screenHeight
}

func main() {
	for i := 0; i < birdCount; i++ {
		createBird(i)
	}
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Birds in the box.")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
