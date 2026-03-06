package math

const (
	Gravity       float64 = 2200.0
	GroundEpsilon float64 = 0.0001
	JumpVelocity  float64 = -900.0
	MoveSpeed     float64 = 300.0
)

// GroundBaseline ist die globale Bodenhöhe für die Gravitation.
var groundBaseline float64

func SetGroundBaseline(y float64) {
	groundBaseline = y
}

func IsOnGround(positionY float64) bool {
	return positionY >= groundBaseline-GroundEpsilon
}

func ApplyGravity(velocityY, dt float64, onGround bool) float64 {
	if onGround {
		return velocityY
	}
	return velocityY + Gravity*dt
}

func ResolveGroundCollision(positionY, velocityY float64) (float64, float64) {
	if positionY >= groundBaseline {
		positionY = groundBaseline
		if velocityY > 0 {
			velocityY = 0
		}
	}
	return positionY, velocityY
}
