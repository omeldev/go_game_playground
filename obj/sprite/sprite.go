package sprite

import (
	"go_game_playground/math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	Position  math.Position
	Velocity  math.Velocity
	Texture   rl.Texture2D
	Animation *Animation
}

// dt := rl.GetFrameTime()
func (s *Sprite) Update(dt float64) {
	s.Position.X += s.Velocity.X * dt
	s.Position.Y += s.Velocity.Y * dt
}

func (s *Sprite) Draw() {
	texture := s.Texture
	if s.Animation != nil {
		texture = s.Animation.GetCurrentTexture()
	}

	rl.DrawTextureEx(
		texture,
		rl.Vector2{X: float32(s.Position.X), Y: float32(s.Position.Y)},
		0.0,
		0.2,
		rl.White,
	)
}
