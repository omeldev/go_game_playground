package main

import (
	"fmt"
	"go_game_playground/game"
	"go_game_playground/internal/camera"
	"go_game_playground/internal/rendering"
	"go_game_playground/obj/entity"
	"go_game_playground/obj/sprite"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	windowWidth  = 800
	windowHeight = 600
)

func main() {
	rl.InitWindow(windowWidth, windowHeight, "Super Gopher Bros")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	// Load textures
	imgBg := rl.LoadTexture("resources/bg.png")
	defer rl.UnloadTexture(imgBg)

	imgWalkLeft := rl.LoadTexture("resources/gophers/walk_left.png")
	defer rl.UnloadTexture(imgWalkLeft)

	imgWalkRight := rl.LoadTexture("resources/gophers/walk_right.png")
	defer rl.UnloadTexture(imgWalkRight)

	imgIdle := rl.LoadTexture("resources/gophers/walk_right.png")
	defer rl.UnloadTexture(imgIdle)

	imgAngryRight := rl.LoadTexture("resources/gophers/angry_right.png")
	defer rl.UnloadTexture(imgAngryRight)

	imgCookie := rl.LoadTexture("resources/cookie/cookie.png")
	defer rl.UnloadTexture(imgCookie)

	// Setup player animation
	playerAnim := sprite.NewAnimation()
	playerAnim.AddFrame("idle", imgIdle)
	playerAnim.AddFrame("walkLeft", imgWalkLeft)
	playerAnim.AddFrame("walkRight", imgWalkRight)
	playerAnim.AddFrame("angry", imgAngryRight)
	playerAnim.SetAnimation("idle")

	// Setup cookie animation
	cookieAnim := sprite.NewAnimation()
	cookieAnim.AddFrame("default", imgCookie)
	cookieAnim.SetAnimation("default")

	// Initialize game entities
	player := entity.NewPlayer(0, 450, playerAnim)
	gameInstance := game.NewGame(player, windowWidth, windowHeight, 10)
	playerCamera := camera.NewPlayerFollowCamera(player, windowWidth, windowHeight)

	spawnTimer := 0.0

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()

		// Spawn cookies periodically
		spawnTimer += float64(dt)
		if spawnTimer > 2.0 {
			gameInstance.SpawnCookie(cookieAnim)
			spawnTimer = 0.0
		}

		// Update game state
		gameInstance.Update(float64(dt))
		playerCamera.Update(player)

		// Render
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(playerCamera.Camera2D)
		rendering.DrawInfiniteBackground(imgBg, playerCamera.Camera2D, windowWidth, windowHeight)
		gameInstance.Draw()
		rl.EndMode2D()

		// UI - Cookie counter
		counterText := fmt.Sprintf("Cookies: %d", gameInstance.CollectibleCount)
		textWidth := rl.MeasureText(counterText, 20)
		rl.DrawText(counterText, int32(windowWidth)-textWidth-16, 16, 20, rl.Black)

		// UI - Controls and FPS
		rl.DrawText("WASD bewegen, SPACE/W springen", 16, 16, 20, rl.LightGray)
		rl.DrawFPS(10, 30)
		rl.EndDrawing()
	}
}
