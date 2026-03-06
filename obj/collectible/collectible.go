package collectible

import (
	"go_game_playground/math"
	"go_game_playground/obj/sprite"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Collectible struct {
	Sprite *sprite.Sprite
	Active bool
	Width  float32
	Height float32
}

func NewCollectible(x, y float64, animation *sprite.Animation) *Collectible {
	return &Collectible{
		Sprite: &sprite.Sprite{
			Position:  math.Position{X: x, Y: y},
			Velocity:  math.Velocity{X: 0, Y: 0},
			Animation: animation,
		},
		Active: true,
		Width:  32,
		Height: 32,
	}
}

func (c *Collectible) Update(dt float64) {
	if !c.Active {
		return
	}
	c.Sprite.Update(dt)
}

func (c *Collectible) Draw() {
	if !c.Active {
		return
	}
	c.Sprite.Draw()
}

func (c *Collectible) GetBounds() rl.Rectangle {
	return rl.NewRectangle(
		float32(c.Sprite.Position.X),
		float32(c.Sprite.Position.Y),
		c.Width,
		c.Height,
	)
}

func (c *Collectible) Collect() {
	c.Active = false
}
