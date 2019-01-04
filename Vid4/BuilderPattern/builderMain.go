package main

import (
	"MasteringGo_Course/Vid4/BuilderPattern/shieldBuilder"
	"fmt"
)

func main() {

	var builder shieldBuilder.ShieldBuilder
	shield := builder.RaiseBack().RaiseFront().Build() // Collecting all info and building object at once

	fmt.Printf("%+v" , *shield)
	fmt.Println()
	shield.UseShield()

}
