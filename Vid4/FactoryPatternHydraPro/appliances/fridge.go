package appliances

//Concrete Fridge type
type Fridge struct {
	typeName string
}

func(fridge *Fridge) Start() {                //Implementing Start() method of Appliance interface
	fridge.typeName = "Fridge "
}

func(fridge *Fridge) GetPurpose() string{                //Implementing GetPurpose() method of Appliance interface
	return "I am a fridge :" + fridge.typeName + " I cool down stuff"
}
