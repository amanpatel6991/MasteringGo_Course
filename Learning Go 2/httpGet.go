package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	const sampleUrl string  = "https://samples.openweathermap.org/data/2.5/weather?q=London,uk&appid=b6907d289e10d714a6e88b30761fae22"

	resp , err :=http.Get(sampleUrl)
	if err !=nil {
		fmt.Println(err)
	}

	respBody , _ := ioutil.ReadAll(resp.Body)

	fmt.Println("RESPONSe ::" , string(respBody))

}
