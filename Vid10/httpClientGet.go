package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type IP struct {
	Ip  string   `json:"ip"`
}

func main() {

	resp , err :=http.Get("https://api.ipify.org?format=json")
	if err !=nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	var ip IP
	err = json.NewDecoder(resp.Body).Decode(&ip)
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println(ip)

}
