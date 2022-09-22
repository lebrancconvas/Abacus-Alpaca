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
	running = true; 
	bgColor = rl.NewColor(255, 196, 196, 0);
	grassTile rl.Texture2D; 
)

// Init. 
func init() {
	rl.InitWindow(screenWidth, screenHeight, titleName);
	rl.SetExitKey(0);
	rl.SetTargetFPS(fps); 

	grassTile = rl.LoadTexture("../assets/pack/Sprout_Lands/Tilesets/Grass.png");
}

// Input. 
func input() {

}

// Game Loop. 
func update() {
	running = !rl.WindowShouldClose();
}

// Draw. 
func drawScene() {
	rl.DrawTexture(grassTile, 100, 100, rl.White);
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
	rl.UnloadTexture(grassTile); 
	rl.CloseWindow();
}

// Main Function. 
func main() {
	for running {
		input();
		update();
		render();
	}

	quit();
}