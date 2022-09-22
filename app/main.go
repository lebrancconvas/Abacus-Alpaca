package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth, screenHeight = 640, 480;
	titleName = "Abacus Alpaca";
	fps = 60;
)

var (
	bgColor = rl.NewColor(255, 196, 196, 0);
)

// Init. 
func init() {
	rl.InitWindow(screenWidth, screenHeight, titleName);
	rl.SetExitKey(0);
	rl.SetTargetFPS(fps); 
}

// Input. 
func input() {

}

// Game Loop. 
func update() {

}

// Draw. 
func drawScene() {
	rl.DrawText("Alpaca", 100, 100, 20, rl.Black);
}

// Rendering. 
func render() {
			// Begin Draw 
			rl.BeginDrawing();

			rl.ClearBackground(bgColor); // Background Color. 
			drawScene();
	
			// End Draw 
			rl.EndDrawing();
}

// Quit. 
func quit() {
	rl.CloseWindow();
}

// Main Function. 
func main() {
	for !rl.WindowShouldClose() {
		input();
		update();
		render();
	}

	quit();
}