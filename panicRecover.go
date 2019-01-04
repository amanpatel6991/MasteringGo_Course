package main

import (
	"fmt"
)

func main() {

	defer func() {
		fmt.Print("hii in defer !")
		if err:=recover(); err!=nil {
			fmt.Print("panice error recovered !" , err)
			for i:=0;i<10;i++ {
				fmt.Println(i)
				//time.Sleep(time.Second*1)
			}
		}
	}()

	panic("sample panic !")




}
