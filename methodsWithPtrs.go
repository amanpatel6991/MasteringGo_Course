package main

import "fmt"

type level int

func main() {
	//lev := new(level)
	var lev level = 1             //it automatically converts lev.raiseLev(2) to &lev.raiseLev(2)

	lev.raiseLev(2)
	lev.raiseLev(4)

	fmt.Println(lev)
}

func(lev *level) raiseLev(i int) {
	if *lev == 0 {
		*lev =1
	}

	*lev = (*lev) * level(i)
}