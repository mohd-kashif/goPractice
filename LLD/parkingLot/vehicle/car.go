package vehicle

import "goPractice/LLD/parkingLot/constant"

type Car struct {
	registrationNumber string
	color              string
	vehicleType        string
}

func NewCar(registrationNumber string, color string) Car {
	return Car{
		registrationNumber: registrationNumber,
		color:              color,
		vehicleType:        constant.Car,
	}
}

func (c *Car) AssignTicket() {

}
