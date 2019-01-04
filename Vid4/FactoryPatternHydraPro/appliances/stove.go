package appliances

//Concrete Stove type
type Stove struct {
	typeName string
}

func(stove *Stove) Start() {                //Implementing Start() method of Appliance interface
	stove.typeName = "Stove "
}

func(stove *Stove) GetPurpose() string{                //Implementing GetPurpose() method of Appliance interface
	return "I am a stove :" + stove.typeName + " I warmp up and even burn things sometimes"
}
