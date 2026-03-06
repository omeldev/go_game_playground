package main

import (
	utils "super-gopher-bros/internal"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	rl.InitWindow(800, 600, "Super Gopher Bros")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	imgBg := rl.LoadTexture("resources/bg.png")
	//imgAngryLeft := rl.LoadTexture("resources/gophers/angry_left.png")
	//imgAngryRight := rl.LoadTexture("resources/gophers/angry_right.png")
	//imgRunLeft := rl.LoadTexture("resources/gophers/run_left.png")
	//imgRunRight := rl.LoadTexture("resources/gophers/run_right.png")
	imgWalkLeft := rl.LoadTexture("resources/gophers/walk_left.png")
	imgWalkRight := rl.LoadTexture("resources/gophers/walk_right.png")

	playerPos := rl.Vector2{X: 0, Y: 450}
	imgWalk := imgWalkRight

	utils.Foo()
	utils.Bar()

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.DrawTextureEx(
			imgBg,
			rl.Vector2{
				X: 0,
				Y: 0,
			},
			0.0,
			1.0,
			rl.White,
		)

		rl.DrawTextureEx(
			imgWalk,
			playerPos,
			0.0,
			0.2,
			rl.White,
		)

		if rl.IsKeyDown(rl.KeyA) {
			rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.Green)
			playerPos.X = playerPos.X - 1
			imgWalk = imgWalkLeft
		} else if rl.IsKeyDown(rl.KeyD) {
			rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.Red)
			playerPos.X = playerPos.X + 1
			imgWalk = imgWalkRight
		} else {
			rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
		}

		rl.EndDrawing()

	}

}
