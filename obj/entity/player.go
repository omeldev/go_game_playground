package entity

import (
	"go_game_playground/math"
	"go_game_playground/obj/sprite"
	"go_game_playground/obj/state"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	*sprite.Sprite
	StateMachine *state.StateMachine
	angryTimer   float64
}

func NewPlayer(x, y float64, animation *sprite.Animation) *Player {
	math.SetGroundBaseline(y)

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
	player.StateMachine.AddState("jump", &JumpState{player: player})
	player.StateMachine.AddState("angry", &AngryState{player: player})
	player.StateMachine.ChangeState("idle")

	return player
}

func (p *Player) Update(dt float64) {
	p.HandleInput()
	p.StateMachine.Update(dt)
	p.applyGravity(dt)
	p.Sprite.Update(dt)
	p.resolveGroundCollision()

	// Wenn wir noch im angry Timer sind und landen, gehe zurück zur angry Animation
	if p.angryTimer > 0 {
		p.angryTimer -= dt
		if math.IsOnGround(p.Position.Y) && p.angryTimer > 0 && p.StateMachine.CurrentState != p.StateMachine.States["angry"] {
			p.StateMachine.ChangeState("angry")
			p.angryTimer = 0 // Reset nach Rückkehr
		}
	}
}

func (p *Player) applyGravity(dt float64) {
	onGround := math.IsOnGround(p.Position.Y)
	p.Velocity.Y = math.ApplyGravity(p.Velocity.Y, dt, onGround)
}

func (p *Player) resolveGroundCollision() {
	p.Position.Y, p.Velocity.Y = math.ResolveGroundCollision(p.Position.Y, p.Velocity.Y)
}

func (p *Player) Jump() {
	if !math.IsOnGround(p.Position.Y) {
		return
	}
	p.StateMachine.ChangeState("jump")
}

func (p *Player) HandleInput() {
	if rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeyW) {
		p.Jump()
	}

	if rl.IsKeyDown(rl.KeyA) {
		p.Velocity.X = -math.MoveSpeed
	} else if rl.IsKeyDown(rl.KeyD) {
		p.Velocity.X = math.MoveSpeed
	} else {
		p.Velocity.X = 0
	}

	// State-Wechsel ignorieren während angry State aktiv ist
	if p.StateMachine.CurrentState == p.StateMachine.States["angry"] {
		return
	}

	if !math.IsOnGround(p.Position.Y) {
		return
	}

	if p.Velocity.X < 0 {
		p.StateMachine.ChangeState("walkLeft")
	} else if p.Velocity.X > 0 {
		p.StateMachine.ChangeState("walkRight")
	} else {
		p.StateMachine.ChangeState("idle")
	}
}

func (p *Player) ShowAngryReaction() {
	// Nur zu angry State wechseln, wenn nicht bereits im angry State
	if p.StateMachine.CurrentState == p.StateMachine.States["angry"] {
		return
	}
	p.angryTimer = 1.5 // 1.5 Sekunden angry Countdown
	p.StateMachine.ChangeState("angry")
}
