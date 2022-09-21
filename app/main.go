package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth, screenHeight = 640, 480;
	titleName = "Abacus Alpaca";
	fps = 60;
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, titleName);
	rl.SetTargetFPS(fps); 

	for !rl.WindowShouldClose() {
		// Begin Draw 
		rl.BeginDrawing();

		rl.ClearBackground(rl.Green); // Background Color. 
		rl.DrawText("Alpaca", 100, 100, 20, rl.Black);

		// End Draw 
		rl.EndDrawing();
	}

	rl.CloseWindow();
}