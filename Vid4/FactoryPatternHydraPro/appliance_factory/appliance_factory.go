package appliance_factory

import (
	"MasteringGo_Course/Vid4/FactoryPatternHydraPro/appliances"
	"github.com/pkg/errors"
)

type Appliance interface {
	Start()
	GetPurpose() string
}

const(
	STOVE = iota                                 //counts increments by 1 (as we used iota)
	FRIDGE
)

//Factory Method to create the appliance
func CreateAppliance(myType int) (Appliance , error){

	switch myType {
	case STOVE:
		return new(appliances.Stove) , nil
	case FRIDGE:
		return new(appliances.Fridge) , nil
	default:
		return nil , errors.New("Error crating appliance !")
	}

}