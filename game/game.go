package game

import (
	"go_game_playground/math"
	"go_game_playground/obj/collectible"
	"go_game_playground/obj/entity"
	"go_game_playground/obj/sprite"
	mathRand "math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Player           *entity.Player
	Collectibles     []*collectible.Collectible
	CollectibleCount int
	MapWidth         float64
	MapHeight        float64
	MaxCollectibles  int
}

func NewGame(player *entity.Player, mapWidth, mapHeight float64, maxCollectibles int) *Game {
	return &Game{
		Player:           player,
		Collectibles:     make([]*collectible.Collectible, 0),
		CollectibleCount: 0,
		MapWidth:         mapWidth,
		MapHeight:        mapHeight,
		MaxCollectibles:  maxCollectibles,
	}
}

func (g *Game) SpawnCookie(cookieAnimation *sprite.Animation) {
	activeCount := 0
	for _, c := range g.Collectibles {
		if c.Active {
			activeCount++
		}
	}

	if activeCount >= g.MaxCollectibles {
		return
	}

	var x, y float64
	for {
		// Spawn zwischen Player.X und Player.X + MapWidth
		offset := float64(mathRand.Intn(int(g.MapWidth)))
		x = g.Player.Position.X + offset
		y = g.Player.Position.Y // Bodenhöhe vom Player

		dx := x - g.Player.Position.X
		if dx*dx > math.CollectibleSpawnDistance*math.CollectibleSpawnDistance {
			break
		}
	}

	c := collectible.NewCookie(x, y, cookieAnimation)
	g.Collectibles = append(g.Collectibles, c.Collectible)
}

// SpawnCollectible deprecated: Use SpawnCookie instead
func (g *Game) SpawnCollectible(cookieAnimation *sprite.Animation) {
	g.SpawnCookie(cookieAnimation)
}

func (g *Game) Update(dt float64) {
	g.Player.Update(dt)

	for _, c := range g.Collectibles {
		c.Update(dt)
	}

	g.checkCollisions()
}

func (g *Game) checkCollisions() {
	playerBounds := rl.NewRectangle(
		float32(g.Player.Position.X),
		float32(g.Player.Position.Y),
		32,
		32,
	)

	for _, c := range g.Collectibles {
		if !c.Active {
			continue
		}

		collectibleBounds := c.GetBounds()
		if rl.CheckCollisionRecs(playerBounds, collectibleBounds) {
			c.Collect()
			g.CollectibleCount++
			g.Player.ShowAngryReaction()
		}
	}
}

func (g *Game) Draw() {
	g.Player.Sprite.Draw()

	for _, c := range g.Collectibles {
		c.Draw()
	}
}
