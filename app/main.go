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
	playerSprite rl.Texture2D; 
	playerSrc rl.Rectangle;
	playerDest rl.Rectangle;

	playerSpeed float32 = 3;
)

// Init. 
func init() {
	rl.InitWindow(screenWidth, screenHeight, titleName);
	rl.SetExitKey(0);
	rl.SetTargetFPS(fps); 

	grassTile = rl.LoadTexture("../assets/pack/Sprout_Lands/Tilesets/Grass.png");
	playerSprite = rl.LoadTexture("../assets/pack/Toon_Character/Robot/Tilesheet/character_robot_sheet.png"); 

	playerSrc = rl.NewRectangle(0, 0, 100, 160);
	playerDest = rl.NewRectangle(80, 100, 100, 100);
}

// Input. 
func input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		playerDest.Y -= playerSpeed;
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		playerDest.Y += playerSpeed;
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		playerDest.X -= playerSpeed;
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		playerDest.X += playerSpeed;
	}
}

// Game Loop. 
func update() {
	running = !rl.WindowShouldClose();
}

// Draw. 
func drawScene() {
	rl.DrawTexture(grassTile, 100, 100, rl.White);
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0, rl.White); 
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

	rl.UnloadTexture(playerSprite);  

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