package vehicle

import "goPractice/LLD/parkingLot/constant"

type Bike struct {
	registrationNumber string
	color              string
	vehicleType        string
}

func NewBike(registrationNumber string, color string) Bike {
	return Bike{
		registrationNumber: registrationNumber,
		color:              color,
		vehicleType:        constant.Bike,
	}
}

func (c *Bike) AssignTicket() {

}
