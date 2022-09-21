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
		rl.BeginDrawing();
		rl.ClearBackground(rl.RayWhite);
		rl.DrawText("Alpaca", 100, 100, 20, rl.Black);
		rl.EndDrawing();
	}

	rl.CloseWindow();
}