package elevator

type Button interface {
	IsPressed() bool
	PressDown()
}

type DoorButton struct {
	Status bool
}

func (db *DoorButton) NewDoorButton() DoorButton {
	return DoorButton{
		Status: false,
	}
}

func (db *DoorButton) IsPressed() bool {
	return db.Status
}

func (db *DoorButton) PressDown() {
	db.Status = true
}

type ElevatorButton struct {
	FloorNumber int
	Status      bool
}

func (eb *ElevatorButton) NewElevatorButton(floorNumber int) ElevatorButton {
	return ElevatorButton{
		Status:      false,
		FloorNumber: floorNumber,
	}
}

func (eb *ElevatorButton) IsPressed() bool {
	return eb.Status
}

func (eb *ElevatorButton) PressDown() {
	eb.Status = true
}

func (eb *ElevatorButton) GiveFloorNumber() int {
	return eb.FloorNumber
}

type HallButton struct {
	Status    bool
	Direction Direction
}

func (db *HallButton) NewHallButton(direction string) HallButton {
	return HallButton{
		Status:    false,
		Direction: Direction(direction),
	}
}

func (db *HallButton) IsPressed() bool {
	return db.Status
}

func (db *HallButton) PressDown() {
	db.Status = true
}
