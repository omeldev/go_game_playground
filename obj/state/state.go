package state

type State interface {
	Enter()
	Update(dt float64)
	Exit()
}
