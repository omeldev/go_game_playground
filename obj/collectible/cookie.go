package collectible

import (
	"go_game_playground/obj/sprite"
)

type Cookie struct {
	*Collectible
}

// NewCookie erstellt ein neues Cookie-Collectible
func NewCookie(x, y float64, animation *sprite.Animation) *Cookie {
	return &Cookie{
		Collectible: NewCollectible(x, y, animation),
	}
}
