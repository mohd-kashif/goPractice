package vehicle

import "goPractice/LLD/parkingLot/constant"

type Truck struct {
	registrationNumber string
	color              string
	vehicleType        string
}

func NewTruck(registrationNumber string, color string) Truck {
	return Truck{
		registrationNumber: registrationNumber,
		color:              color,
		vehicleType:        constant.Truck,
	}
}

func (c *Truck) AssignTicket() {

}
