package main

import (
	"image/color"
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

// Constants defining the screen dimensions, number of birds, and simulation parameters.
const (
	screenWidth, screenHeight = 640, 360
	birdCount                 = 500
	viewRadius                = 13
	adjRate                   = 0.015
)

var (
	// Green color used to represent birds on the screen.
	green = color.RGBA{R: 10, G: 255, B: 50, A: 255}
	// Array to store bird objects.
	birds [birdCount]*Bird
	// Map to track birds' positions on the screen.
	birdsMap [screenWidth + 1][screenHeight + 1]int

	lock = sync.Mutex{}
)

// Game struct represents the game state.
type Game struct{}

// Update method updates the game state.
func (g *Game) Update() error {
	return nil
}

// Draw method draws the game state on the screen.
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw each bird on the screen.
	for _, bird := range birds {
		screen.Set(int(bird.position.x+1), int(bird.position.y), green)
		screen.Set(int(bird.position.x-1), int(bird.position.y), green)
		screen.Set(int(bird.position.x), int(bird.position.y-1), green)
		screen.Set(int(bird.position.x), int(bird.position.y+1), green)
	}
}

// Layout method returns the screen dimensions.
func (g *Game) Layout(_, _ int) (w, h int) {
	return screenWidth, screenHeight
}

// main function initializes the game and starts the event loop.
func main() {
	// Initialize the birds map with default values.
	for i, row := range birdsMap {
		for j := range row {
			birdsMap[i][j] = -1
		}
	}
	// Create birds and start their movement simulations.
	for i := 0; i < birdCount; i++ {
		createBird(i)
	}
	// Set the window size and title.
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Birds in the box.")
	// Start the game event loop.
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
