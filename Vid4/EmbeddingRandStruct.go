package main

import (
	"math/rand"
	"fmt"
	"time"
)

type customRand struct {
	*rand.Rand
	count int
}

func newRand(i int64) *customRand{
	return &customRand{
		Rand : rand.New(rand.NewSource(i)),
		count: 0,
	}
}

func(cRand *customRand) rangeRand(min , max int) int{           //Remove * in Receiver and see the change in count
	cRand.count++                   //Go takes care that cRand is pointer and we are still using cRand.count
	return cRand.Rand.Intn(max - min) + min                      //Generates random no. between max and min
}

func(cRand *customRand) getCount() int{
	return cRand.count
}

//Overrides the Inner type's method Intn
//To access Intn() of inner type in this case , use custRand.Rand.Intn()instead of custRand.Intn()
//func(cRand *customRand) Intn(in int) int{
//	return 10+in
//}

func main() {
	custRand := newRand(time.Now().UnixNano())
	fmt.Println(custRand.Intn(5))
	fmt.Println(custRand.rangeRand(5 , 20))
	fmt.Println(custRand.rangeRand(5 , 20))
	fmt.Println(custRand.getCount())

	//var h Human
	//h = Man{Human : Human{"Aman" , 21} , Voice: "Good"}  // Gives Error -> cannot use Man literal (type Man) as type Human in assignment (unlike java)
	//
	//fmt.Println(h)

}

type Human struct {
	Name string
	Age int

}

type Man struct {
	Human
	Voice string
}