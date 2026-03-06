package entity

import (
	"go_game_playground/math"
)

type IdleState struct {
	player *Player
}

func (s *IdleState) Enter() {
	s.player.Animation.SetAnimation("idle")
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
}

func (s *WalkRightState) Update(dt float64) {
}

func (s *WalkRightState) Exit() {
}

type JumpState struct {
	player *Player
}

func (s *JumpState) Enter() {
	s.player.Animation.SetAnimation("jump")
	s.player.Velocity.Y = math.JumpVelocity
}

func (s *JumpState) Update(dt float64) {
	if math.IsOnGround(s.player.Position.Y) && s.player.Velocity.Y >= 0 {
		s.player.StateMachine.ChangeState("idle")
	}
}

func (s *JumpState) Exit() {
}

type AngryState struct {
	player   *Player
	duration float64
	elapsed  float64
}

func (s *AngryState) Enter() {
	s.player.Animation.SetAnimation("angry")
	s.elapsed = 0
	s.duration = 0.3 // 0.3 Sekunden angry - schnelle Reaktion
}

func (s *AngryState) Update(dt float64) {
	s.elapsed += dt
	if s.elapsed >= s.duration {
		s.player.StateMachine.ChangeState("idle")
	}
}

func (s *AngryState) Exit() {
}
