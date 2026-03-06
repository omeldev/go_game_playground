package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
    rl.InitWindow(800, 450, "raylib [core] example - basic window")
    defer rl.CloseWindow()

    rl.SetTargetFPS(60)

    for !rl.WindowShouldClose() {
        rl.BeginDrawing()

        rl.ClearBackground(rl.Black)

        if rl.IsKeyPressed(rl.KeyD) {
            rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.Green)
        } else {
            rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
        }

        rl.EndDrawing()
    }
}
