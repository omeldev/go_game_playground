package main

import (
	utils "super-gopher-bros/internal"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	utils.Foo()
	utils.Bar()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		if rl.IsKeyDown(rl.KeyA) {
			rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.Green)
		} else if rl.IsKeyDown(rl.KeyD) {
			rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.Red)
		} else {
			rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
		}

		rl.EndDrawing()
	}
}
