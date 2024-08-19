package elevator

type ElevatorState string

const (
	IDLE        ElevatorState = "idle"
	MOVING_UP   ElevatorState = "up"
	MOVING_DOWN ElevatorState = "down"
)

type Direction string

const (
	DirectionUp   Direction = "up"
	DirectionDown Direction = "down"
)

type DoorState string

const (
	OPEN  DoorState = "open"
	CLOSE DoorState = "close"
)
