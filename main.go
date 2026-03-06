package main

import (
	"fmt"
	"go_game_playground/game"
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
	defer rl.UnloadTexture(imgBg)

	imgWalkLeft := rl.LoadTexture("resources/gophers/walk_left.png")
	defer rl.UnloadTexture(imgWalkLeft)

	imgWalkRight := rl.LoadTexture("resources/gophers/walk_right.png")
	defer rl.UnloadTexture(imgWalkRight)

	imgIdle := rl.LoadTexture("resources/gophers/walk_right.png")
	defer rl.UnloadTexture(imgIdle)

	imgAngryLeft := rl.LoadTexture("resources/gophers/angry_left.png")
	defer rl.UnloadTexture(imgAngryLeft)

	imgAngryRight := rl.LoadTexture("resources/gophers/angry_right.png")
	defer rl.UnloadTexture(imgAngryRight)

	imgCookie := rl.LoadTexture("resources/cookie/cookie.png")
	defer rl.UnloadTexture(imgCookie)

	anim := sprite.NewAnimation()
	anim.AddFrame("idle", imgIdle)
	anim.AddFrame("walkLeft", imgWalkLeft)
	anim.AddFrame("walkRight", imgWalkRight)
	anim.AddFrame("angry", imgAngryRight)
	anim.SetAnimation("idle")

	cookieAnim := sprite.NewAnimation()
	cookieAnim.AddFrame("default", imgCookie)
	cookieAnim.SetAnimation("default")

	player := entity.NewPlayer(0, 450, anim)
	gameInstance := game.NewGame(player, windowWidth, windowHeight, 10)
	camera := NewPlayerFollowCamera(player)

	spawnTimer := 0.0

	for !rl.WindowShouldClose() {

		dt := rl.GetFrameTime()

		spawnTimer += float64(dt)
		if spawnTimer > 2.0 {
			gameInstance.SpawnCookie(cookieAnim)
			spawnTimer = 0.0
		}

		gameInstance.Update(float64(dt))
		UpdatePlayerFollowCamera(&camera, player)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)
		DrawInfiniteBackground(imgBg, camera, windowWidth, windowHeight)
		gameInstance.Draw()
		rl.EndMode2D()

		// Cookie counter - fix oben rechts auf dem Bildschirm in schwarz
		counterText := fmt.Sprintf("Cookies: %d", gameInstance.CollectibleCount)
		textWidth := rl.MeasureText(counterText, 20)
		rl.DrawText(counterText, int32(windowWidth)-textWidth-16, 16, 20, rl.Black)

		rl.DrawText("WASD bewegen, SPACE/W springen", 16, 16, 20, rl.LightGray)
		rl.DrawFPS(10, 30)
		rl.EndDrawing()
	}
}
