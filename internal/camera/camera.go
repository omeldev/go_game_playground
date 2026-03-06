package camera

import (
	"go_game_playground/obj/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const GroundScreenPadding = 150

type Camera struct {
	Camera2D rl.Camera2D
}

func NewPlayerFollowCamera(player *entity.Player, screenWidth, screenHeight int) *Camera {
	playerCenter := getPlayerCenter(player)

	return &Camera{
		Camera2D: rl.Camera2D{
			Target: playerCenter,
			Offset: rl.Vector2{
				X: float32(screenWidth) / 2,
				Y: getGroundBoundOffsetY(player, playerCenter, screenHeight),
			},
			Rotation: 0,
			Zoom:     1,
		},
	}
}

func (c *Camera) Update(player *entity.Player) {
	c.Camera2D.Target.X = getPlayerCenter(player).X
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

func getGroundBoundOffsetY(player *entity.Player, playerCenter rl.Vector2, screenHeight int) float32 {
	groundScreenY := float32(screenHeight - GroundScreenPadding)
	// Keep the player's ground point on a stable screen line near the bottom.
	return groundScreenY + (playerCenter.Y - float32(player.Position.Y))
}
