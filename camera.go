package main

import (
	"go_game_playground/obj/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const cameraGroundScreenPadding = 150

func NewPlayerFollowCamera(player *entity.Player) rl.Camera2D {
	playerCenter := getPlayerCenter(player)

	return rl.Camera2D{
		Target: playerCenter,
		Offset: rl.Vector2{
			X: float32(windowWidth) / 2,
			Y: getGroundBoundOffsetY(player, playerCenter),
		},
		Rotation: 0,
		Zoom:     1,
	}
}

func UpdatePlayerFollowCamera(camera *rl.Camera2D, player *entity.Player) {
	camera.Target.X = getPlayerCenter(player).X
}

func getPlayerCenter(player *entity.Player) rl.Vector2 {
	texture := player.Animation.GetCurrentTexture()
	halfWidth := float64(texture.Width) * 0.1
	halfHeight := float64(texture.Height) * 0.1

	return rl.Vector2{
		X: float32(player.Position.X + halfWidth),
		Y: float32(player.Position.Y + halfHeight),
	}
}

func getGroundBoundOffsetY(player *entity.Player, playerCenter rl.Vector2) float32 {
	groundScreenY := float32(windowHeight - cameraGroundScreenPadding)
	// Keep the player's ground point on a stable screen line near the bottom.
	return groundScreenY + (playerCenter.Y - float32(player.Position.Y))
}
