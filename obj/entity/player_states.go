package entity

import rl "github.com/gen2brain/raylib-go/raylib"

type IdleState struct {
	player *Player
}

func (s *IdleState) Enter() {
	s.player.Animation.SetAnimation("idle")
	s.player.Velocity.X = 0
}

func (s *IdleState) Update(dt float64) {
}

func (s *IdleState) Exit() {
}

type WalkLeftState struct {
	player *Player
}

func (s *WalkLeftState) Enter() {
	s.player.Animation.SetAnimation("walkLeft")
	s.player.Velocity.X = -300
}

func (s *WalkLeftState) Update(dt float64) {
}

func (s *WalkLeftState) Exit() {
}

type WalkRightState struct {
	player *Player
}

func (s *WalkRightState) Enter() {
	s.player.Animation.SetAnimation("walkRight")
	s.player.Velocity.X = 300
}

func (s *WalkRightState) Update(dt float64) {
}

func (s *WalkRightState) Exit() {
}

func (p *Player) HandleInput() {
	if rl.IsKeyDown(rl.KeyA) {
		p.StateMachine.ChangeState("walkLeft")
	} else if rl.IsKeyDown(rl.KeyD) {
		p.StateMachine.ChangeState("walkRight")
	} else {
		p.StateMachine.ChangeState("idle")
	}
}
