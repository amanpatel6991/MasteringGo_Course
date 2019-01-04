package main

import (
	"fmt"
	"MasteringGo_Course/Vid5/customConfigrator/configrator"
)

type ConfS struct {
	TS      string  `name:"testString"`
	TB      bool    `name:"testBool"`
	TF      float64 `name:"testFloat"`
	TestInt int
}

func main() {                          //RUN FROM CMD or OPEN THIS FOLDER AND RUN
	configstruct := new(ConfS)
	err := configrator.GetConf(configrator.CUSTOM, configstruct, "config.conf")
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(*configstruct)

	if configstruct.TB {
		fmt.Println("bool is true")
	}

	fmt.Println(float64(4.8 * configstruct.TF))

	fmt.Println(5 * configstruct.TestInt)

	fmt.Println(configstruct.TS)
}
