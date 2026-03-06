package rendering

import (
	stdmath "math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawInfiniteBackground(texture rl.Texture2D, camera rl.Camera2D, screenWidth, screenHeight int) {
	tileWidth := float64(texture.Width)
	tileHeight := float64(texture.Height)
	if tileWidth <= 0 || tileHeight <= 0 || camera.Zoom == 0 {
		return
	}

	zoom := float64(camera.Zoom)
	left := float64(camera.Target.X) - float64(camera.Offset.X)/zoom
	right := left + float64(screenWidth)/zoom
	top := float64(camera.Target.Y) - float64(camera.Offset.Y)/zoom
	bottom := top + float64(screenHeight)/zoom

	startX := int(stdmath.Floor(left/tileWidth)) - 1
	endX := int(stdmath.Floor(right/tileWidth)) + 1
	startY := int(stdmath.Floor(top/tileHeight)) - 1
	endY := int(stdmath.Floor(bottom/tileHeight)) + 1

	for tileX := startX; tileX <= endX; tileX++ {
		x := float32(float64(tileX) * tileWidth)
		for tileY := startY; tileY <= endY; tileY++ {
			y := float32(float64(tileY) * tileHeight)
			rl.DrawTextureEx(texture, rl.Vector2{X: x, Y: y}, 0, 1, rl.White)
		}
	}
}
