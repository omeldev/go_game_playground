package entity

import (
	"go_game_playground/math"
	"go_game_playground/obj/sprite"
	"go_game_playground/obj/state"
)

type Player struct {
	*sprite.Sprite
	StateMachine *state.StateMachine
}

func NewPlayer(x, y float64, animation *sprite.Animation) *Player {
	player := &Player{
		Sprite: &sprite.Sprite{
			Position:  math.Position{X: x, Y: y},
			Velocity:  math.Velocity{X: 0, Y: 0},
			Animation: animation,
		},
		StateMachine: state.NewStateMachine(),
	}

	player.StateMachine.AddState("idle", &IdleState{player: player})
	player.StateMachine.AddState("walkLeft", &WalkLeftState{player: player})
	player.StateMachine.AddState("walkRight", &WalkRightState{player: player})
	player.StateMachine.ChangeState("idle")

	return player
}

func (p *Player) Update(dt float64) {
	p.HandleInput()
	p.StateMachine.Update(dt)
	p.Sprite.Update(dt)
}
