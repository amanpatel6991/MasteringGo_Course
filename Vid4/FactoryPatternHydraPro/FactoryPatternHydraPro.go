//Building system to control appliances of Hydra
//Start appliance and get there description

package main

import (
	"fmt"
	"MasteringGo_Course/Vid4/FactoryPatternHydraPro/appliance_factory"
)

func main() {
	//Use the  factory method now to crate appliance objects automatically
	//instead of creating them manually here

	fmt.Println("Select appliance type")
	fmt.Println("0 : Stove")
	fmt.Println("1 : Fridge")

	var myType int
	fmt.Scan(&myType)

	myApp, err := appliance_factory.CreateAppliance(myType)

	if err == nil {
		myApp.Start()
		fmt.Println(myApp.GetPurpose())
	} else {
		fmt.Println(err)
	}

}
