package main

import (
	"go_game_playground/obj/entity"
	"go_game_playground/obj/sprite"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const windowWidth = 800
const windowHeight = 600

func main() {

	rl.InitWindow(windowWidth, windowHeight, "Super Gopher Bros")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	imgBg := rl.LoadTexture("resources/bg.png")
	imgWalkLeft := rl.LoadTexture("resources/gophers/walk_left.png")
	imgWalkRight := rl.LoadTexture("resources/gophers/walk_right.png")
	imgIdle := rl.LoadTexture("resources/gophers/walk_right.png")

	anim := sprite.NewAnimation()
	anim.AddFrame("idle", imgIdle)
	anim.AddFrame("walkLeft", imgWalkLeft)
	anim.AddFrame("walkRight", imgWalkRight)
	anim.SetAnimation("idle")

	player := entity.NewPlayer(0, 450, anim)

	for !rl.WindowShouldClose() {

		dt := rl.GetFrameTime()

		player.Update(float64(dt))

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.DrawTextureEx(
			imgBg,
			rl.Vector2{X: 0, Y: 0},
			0.0,
			1.0,
			rl.White,
		)

		player.Draw()

		rl.DrawText("Congrats! You created your first window!", windowWidth/3, windowHeight/4, 20, rl.LightGray)

		rl.EndDrawing()

	}

}
