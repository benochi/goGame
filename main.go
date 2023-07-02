package main

import (
	"./game"
	"time"
)

func main() {
	window := NewWindow(1920, 1080)
	window.Open()

	game := &Game{}

	input := NewInput()
	graphics := NewGraphics()

	for {
		// Display splash screen while the game loads
		graphics.Render()

		// Simulate game loading time
		// You can replace this with actual game loading logic
		time.Sleep(3 * time.Second)

		// Display start screen
		graphics.Render()

		// Process user input
		input.ProcessInput()

		// Start the game
		game.Run()

		// Cleanup and exit the game
		window.Close()
		break
	}
}
