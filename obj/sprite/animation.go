package sprite

import rl "github.com/gen2brain/raylib-go/raylib"

type Animation struct {
	Frames           map[string]rl.Texture2D
	CurrentAnimation string
}

func NewAnimation() *Animation {
	return &Animation{
		Frames: make(map[string]rl.Texture2D),
	}
}

func (a *Animation) AddFrame(key string, texture rl.Texture2D) {
	a.Frames[key] = texture
}

func (a *Animation) SetAnimation(key string) {
	if _, exists := a.Frames[key]; exists {
		a.CurrentAnimation = key
	}
}

func (a *Animation) GetCurrentTexture() rl.Texture2D {
	return a.Frames[a.CurrentAnimation]
}
